package ssss

import (
	"encoding/hex"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	bitcoinAddress = "36vg9gKtEAVdou1z4ExFNzuXokSWSdoDRB"
	bitcoinPubkey  = "022ac329a4ed93c17ae841d348c76467952be188b40c86059343dbb07e9bc3b9aa"
	bitcoinPrivkey = "Kyc4BGZonrq1cg4jvtKbnM5hXjde7o7Nm3BCs1a7SpAYiFYUEhQX"
)

func TestSplitSecret(t *testing.T) {
	Convey("success split secret 3 of 5", t, func() {
		shares, err := Split(5, 3, []byte(bitcoinPrivkey))
		So(err, ShouldBeNil)
		So(len(shares), ShouldEqual, 5)
	})

	Convey("success restore secert 3 of 5", t, func() {
		share1, _ := hex.DecodeString("a66f642c3250aa733405d2e6969efb3908c30addb2a26a3bccbe47826a2749bb72098e2897f8d676d5c4c68832334551f215be91")
		share2, _ := hex.DecodeString("85f6910a24c4c382b4bafed4981aa5cd9463fc433a64f27be93498e1d2990dc826f2fd59505b9297297b588a22ea1df26078ba41")
		share3, _ := hex.DecodeString("68e0961254d3339eeecd5d036de36a9eead4bdfce68bad287de0bb068fd1733d39c83132b49225d6afcfdf5b799f01f6d7055588")
		//share4, _ := hex.DecodeString("f0c64170a0d8e0d7626111bba6fb5d3bd1c9ea2b37825fac1f4f5929811b769e3448d76ad844b2e4089d435726063a8a54cd241c")
		//share5, _ := hex.DecodeString("1dd04668d0cf10cb3816b26c53029268af7eab94eb6d00ff8b9b7acedc53086b2b721b013c8d05a58e29c4867d73268ee3b0cbd5")

		shares := map[byte][]byte{
			1: share1,
			2: share2,
			3: share3,
		}

		secret := Combine(shares)

		So(string(secret), ShouldEqual, bitcoinPrivkey)
	})

	Convey("success split secret 2 of 2", t, func() {
		shares, err := Split(2, 2, []byte(bitcoinPrivkey))
		So(err, ShouldBeNil)
		So(len(shares), ShouldEqual, 2)
	})

	Convey("success restore secert 2 of 2", t, func() {
		share1, _ := hex.DecodeString("86aa5b5067030b61cb0077345e02e288400cd79927bf826ad401b6ee6e4664660abfda18230849be1cb7e827cd7dfefb9e2eed45")
		share2, _ := hex.DecodeString("cac413fc08cff8733f967d3b19ad83b51a84688ffcb2406c5bbcdb68853d911ea33069f5d343313ecde508a53a300c12e8e43262")

		shares := map[byte][]byte{
			1: share1,
			2: share2,
		}

		secret := Combine(shares)

		So(string(secret), ShouldEqual, bitcoinPrivkey)
	})
}
