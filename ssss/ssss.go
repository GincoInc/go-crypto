package ssss

import (
	"crypto/rand"
	"errors"
	"io"
)

var (
	// ErrInvalidCount is returned when the count parameter is invalid.
	ErrInvalidCount = errors.New("N must be >= K")
	// ErrInvalidThreshold is returned when the threshold parameter is invalid.
	ErrInvalidThreshold = errors.New("K must be > 1")
)

type pair struct {
	x byte
	y byte
}

// Split ...
func Split(n, k byte, secret []byte) (map[byte][]byte, error) {
	if k <= 1 {
		return nil, ErrInvalidThreshold
	}
	if n < k {
		return nil, ErrInvalidCount
	}
	shares := make(map[byte][]byte, n)
	for _, b := range secret {
		p, err := genPoly(k, b, rand.Reader)
		if err != nil {
			return nil, err
		}
		for x := byte(1); x <= n; x++ {
			shares[x] = append(shares[x], calcPoint(p, x))
		}
	}
	return shares, nil
}

// Combine ...
func Combine(shares map[byte][]byte) []byte {
	var secret []byte
	for _, v := range shares {
		secret = make([]byte, len(v))
		break
	}
	points := make([]pair, len(shares))
	for i := range secret {
		p := 0
		for k, v := range shares {
			points[p] = pair{x: k, y: v[i]}
			p++
		}
		secret[i] = interpolate(points, 0)
	}
	return secret
}

func gdiv(a byte, b byte) byte {
	for i := byte(0); i <= 0xff; i++ {
		if gmul(i, b) == a {
			return i
		}
	}
	return 0
}

func gmul(a byte, b byte) byte {
	var counter, hiBitSet byte
	p := byte(0)
	for counter = 0; counter < 8; counter++ {
		if (b & 1) == 1 {
			p ^= a
		}
		hiBitSet = a & 0x80
		a <<= 1
		if hiBitSet == 0x80 {
			a ^= 0x1b
		}
		b >>= 1
	}
	return p
}

func genPoly(k byte, secret byte, rand io.Reader) ([]byte, error) {
	result := make([]byte, k)
	result[0] = secret
	buf := make([]byte, k-1)
	_, err := io.ReadFull(rand, buf)
	if err != nil {
		return nil, err
	}
	for i := byte(1); i < k; i++ {
		result[i] = buf[i-1]
	}
	return result, nil
}

func calcPoint(p []byte, x byte) byte {
	y := byte(0)
	for i := 1; i <= len(p); i++ {
		y = gmul(y, x) ^ p[len(p)-i]
	}
	return y
}

func interpolate(points []pair, x byte) (value byte) {
	for i, a := range points {
		weight := byte(1)
		for j, b := range points {
			if i != j {
				top := x ^ b.x
				bottom := a.x ^ b.x
				factor := gdiv(top, bottom)
				weight = gmul(weight, factor)
			}
		}
		value = value ^ gmul(weight, a.y)
	}
	return
}
