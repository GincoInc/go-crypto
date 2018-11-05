package crypto

import (
	"encoding/hex"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewHmacDRBG(t *testing.T) {
	Convey("returns the HMAC-DRBG", t, func() {
		entropy, _ := hex.DecodeString("06032cd5eed33f39265f49ecb142c511da9aff2af71203bffaf34a9ca5bd9c0d")
		nonce, _ := hex.DecodeString("0e66f71edc43e42a45ad3c6fc6cdc4df")
		pers := []byte{}
		hmacDRBG := NewHmacDRBG(entropy, nonce, pers)
		k, _ := hex.DecodeString("17dc11c2389f5eeb9d0f6a5148a1ea83ee8a828f4f140ac78272a0da435fa121")
		So(hmacDRBG.k, ShouldResemble, k)
		v, _ := hex.DecodeString("81e0d8830ed2d16f9b288a1cb289c5fab3f3c5c28131be7cafedcc7734604d34")
		So(hmacDRBG.v, ShouldResemble, v)
		So(hmacDRBG.reseedCounter, ShouldEqual, 1)
	})
}

func TestHmacDRBG_Generate(t *testing.T) {
	Convey("returns random number from entropies", t, func() {
		for _, testData := range hmacDRBGTestDataList {
			entropy, _ := hex.DecodeString(testData.entropy)
			nonce, _ := hex.DecodeString(testData.nonce)
			pers, _ := hex.DecodeString(testData.personalization)
			reseedEntropy, _ := hex.DecodeString(testData.reseedEntropy)
			reseedAddInput, _ := hex.DecodeString(testData.reseedAdditionalInput)
			addInput1, _ := hex.DecodeString(testData.additionalInputs[0])
			addInput2, _ := hex.DecodeString(testData.additionalInputs[1])
			returned, _ := hex.DecodeString(testData.returned)

			hmacDRBG := NewHmacDRBG(entropy, nonce, pers)
			hmacDRBG.Reseed(reseedEntropy, reseedAddInput)
			result, err := hmacDRBG.Generate(int32(testData.bitsLen/8), addInput1)
			result, err = hmacDRBG.Generate(int32(testData.bitsLen/8), addInput2)
			So(err, ShouldBeNil)
			So(result, ShouldResemble, returned)
		}
	})
}
