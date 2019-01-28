package base58

import (
	"crypto/sha256"
	"log"
	"math/big"

	"github.com/pkg/errors"
)

const (
	BITCOIN string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	XRP     string = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
	FLICKR  string = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
)

// Base58 ...
type Base58 struct {
	charset string
	basemap [256]byte
}

// NewBase58 ...
func NewBase58(charset string) *Base58 {
	if len(charset) >= 255 {
		log.Fatalln("Alphabet is too long.")
	}
	basemap := [256]byte{}
	for i := 0; i < 256; i++ {
		basemap[i] = byte(255)
	}
	for i := 0; i < len(charset); i++ {
		if basemap[charset[i]] != 255 {
			log.Fatalf("%v is ambiguous", string(charset[i]))
		}
		basemap[charset[i]] = byte(i)
	}
	return &Base58{
		charset: charset,
		basemap: basemap,
	}
}

func checkSum(input []byte) (cksum [4]byte) {
	hash1 := sha256.Sum256(input)
	hash2 := sha256.Sum256(hash1[:])
	copy(cksum[:], hash2[:4])
	return
}

// CheckEncode ...
func (b *Base58) CheckEncode(input []byte, version byte) string {
	data := make([]byte, 0, 1+len(input)+4)
	data = append(data, version)
	data = append(data, input[:]...)
	checksum := checkSum(data)
	data = append(data, checksum[:]...)
	return b.Encode(data)
}

// CheckDecode ...
func (b *Base58) CheckDecode(input string) (result []byte, version byte, err error) {
	decoded := b.Decode(input)
	if len(decoded) < 5 {
		return nil, 0, errors.New("Invalid Format")
	}
	version = decoded[0]
	var checksum [4]byte
	copy(checksum[:], decoded[len(decoded)-4:])
	if checkSum(decoded[:len(decoded)-4]) != checksum {
		return nil, 0, errors.New("Checksum Error")
	}
	payload := decoded[1 : len(decoded)-4]
	result = append(result, payload...)
	return
}

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// Decode ...
func (b *Base58) Decode(input string) []byte {
	answer := big.NewInt(0)
	j := big.NewInt(1)

	scratch := new(big.Int)
	for i := len(input) - 1; i >= 0; i-- {
		tmp := b.basemap[input[i]]
		if tmp == 255 {
			return []byte("")
		}
		scratch.SetInt64(int64(tmp))
		scratch.Mul(j, scratch)
		answer.Add(answer, scratch)
		j.Mul(j, bigRadix)
	}

	tmpval := answer.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(input); numZeros++ {
		if input[numZeros] != b.charset[0] {
			break
		}
	}
	flen := numZeros + len(tmpval)
	val := make([]byte, flen)
	copy(val[numZeros:], tmpval)

	return val
}

// Encode ...
func (b *Base58) Encode(input []byte) string {
	x := new(big.Int)
	x.SetBytes(input)
	answer := make([]byte, 0, len(input)*136/100)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, b.charset[mod.Int64()])
	}

	// leading zero bytes
	for i := 0; input[i] == 0 && i < len(input)-1; i++ {
		answer = append(answer, b.charset[0])
	}

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}
