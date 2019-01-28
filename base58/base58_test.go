package base58

import (
	"encoding/hex"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBitcoinBase58(t *testing.T) {
	bs58 := NewBase58(BITCOIN)
	Convey("CheckEncode", t, func() {
		So(bs58.CheckEncode([]byte("Test data"), 0x00), ShouldEqual, "182iP79GRURMp7oMHDU")
	})

	Convey("CheckDecode", t, func() {
		decoded, version, _ := bs58.CheckDecode("K2RYDcKfupxwXdWhSAxQPCeiULntKm63UXyx5MvEH2")
		So(fmt.Sprintf("%x", decoded), ShouldEqual, fmt.Sprintf("%x", "abcdefghijklmnopqrstuvwxyz"))
		So(version, ShouldEqual, 0x14)
	})
}

func TestXrpBase58(t *testing.T) {
	bs58 := NewBase58(XRP)
	buff, _ := hex.DecodeString("2decab42ca805119a9ba2ff305c9afa12f0b86a1")
	Convey("CheckEncode", t, func() {
		So(bs58.CheckEncode(buff, 0x00), ShouldEqual, "rnBFvgZphmN39GWzUJeUitaP22Fr9be75H")
	})

	Convey("CheckDecode", t, func() {
		decoded, version, err := bs58.CheckDecode("rnBFvgZphmN39GWzUJeUitaP22Fr9be75H")
		if err != nil {
			fmt.Println(err)
		}
		So(fmt.Sprintf("%x", decoded), ShouldEqual, "2decab42ca805119a9ba2ff305c9afa12f0b86a1")
		So(version, ShouldEqual, 0x00)
	})
}

//
//func TestFlickrBase58(t *testing.T) {
//	bs58 := NewBase58(flickr)
//	Convey("Bitcoin", t, func() {
//	})
//
//	Convey("CheckEncode", t, func() {
//
//	})
//
//	Convey("Decode", t, func() {
//
//	})
//
//	Convey("CheckDecode", t, func() {
//
//	})
//}
