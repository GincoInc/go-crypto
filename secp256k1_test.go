package crypto

import (
	"crypto/elliptic"
	"math/big"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSecp256k1(t *testing.T) {
	Convey("returns the secp256k1 curve", t, func() {
		So(Secp256k1(), ShouldEqual, secp256k1)
		_, ok := Secp256k1().(elliptic.Curve)
		So(ok, ShouldBeTrue)
	})
}

func TestCurveParams_Params(t *testing.T) {
	Convey("returns the parameters for the curve", t, func() {
		So(Secp256k1().Params().P, ShouldEqual, p)
		So(Secp256k1().Params().N, ShouldEqual, n)
		So(Secp256k1().Params().B, ShouldEqual, b)
		So(Secp256k1().Params().Gx, ShouldEqual, gx)
		So(Secp256k1().Params().Gy, ShouldEqual, gy)
		So(Secp256k1().Params().BitSize, ShouldEqual, bitSize)
		So(Secp256k1().Params().Name, ShouldEqual, name)
	})
}

func TestCurveParams_IsOnCurve(t *testing.T) {
	Convey("reports whether the given (x,y) lies on the curve", t, func() {
		for _, vector := range testVectors {
			x, _ := new(big.Int).SetString(vector.x, 16)
			y, _ := new(big.Int).SetString(vector.y, 16)
			So(Secp256k1().IsOnCurve(x, y), ShouldBeTrue)
		}
	})
}

func TestCurveParams_Add_And_Double(t *testing.T) {
	Convey("return the sum of 2*(x1,y1)", t, func() {
		for _, vector := range testVectors {
			k, _ := new(big.Int).SetString(vector.k, 10)
			x, y := Secp256k1().ScalarBaseMult(k.Bytes())
			x1, y1 := Secp256k1().Add(x, y, x, y)
			x2, y2 := Secp256k1().Double(x, y)
			So(x1, ShouldResemble, x2)
			So(y1, ShouldResemble, y2)
		}
	})
}

func TestCurveParams_ScalarBaseMult(t *testing.T) {
	Convey("returns k*G, where G is the base point of the group and k is an integer in big-endian form", t, func() {
		for _, vector := range testVectors {
			k, _ := new(big.Int).SetString(vector.k, 10)
			x, _ := new(big.Int).SetString(vector.x, 16)
			y, _ := new(big.Int).SetString(vector.y, 16)
			kx, ky := Secp256k1().ScalarBaseMult(k.Bytes())
			So(kx, ShouldResemble, x)
			So(ky, ShouldResemble, y)
		}
	})
}
