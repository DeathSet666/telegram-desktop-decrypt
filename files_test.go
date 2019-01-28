package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

const hexSettings0 string = "54444624ca550f0000000020b0ca901027e08bc58296fa67d2e2383f2e62ff5feecc38fb0939d2a9cc2d4561000003c04fb0453b3a675ab56d5e5e2088c5dad7c03ef053cd2e56d2afa5f6582f3ed8be97ccead39c64cc27167eb61f319b73a1b614a97f6f82742090065389cae892f94694c1e3527cdd8e64a2830641bedd4b8af2971a0cf93213182677d0d54dba58a96cbb6e0e279c6160d1d434f7146fd9e4bb8e3c75908b52243db395de506566957da30a2023499f500146975ecac7196a115e5703444d946bcdb5702c221225ef6c3d12f05569b9ed48eda0b4c734232a1420b9dd2d4c5538bc764635fdffe6c0565353a8a1c9abece4ac9dbbd653c3fc83e5359b0c960bc429acb191e03641c3c52502e1a6e2dcf5f4c645c165fe82df7c4f19bb7b3321bde735d6b3d12e307c632cd4c1025bcc53cde0105bf763ee2971d7f9046a788bec6a37d433d140f2c949141dcc2a47ba8dcb3d7a2ff73d3fb3bc3e126584a3287dca4f7c1418c4cf3d7fdabf86effb4ea59480f685ff9ff2c118270a996568065e798b1798418659b081398e40f9b93b2a0d53db12b93d80a4f0ad27ff00be213bdb45e92436abfd6b68761e1d6f5f1db6fa3b093616c5867ae90a27027f408c0a67c501f131a2b548e540bb0eae71558cda2e0c2bf4f79da3928dc52abf529aecc26f088d6c2dcc7b92f660bff54135dd8320c9edcdb2e5894fdce34b4b57c196cb4e68b2e5f957c90406a3e0e580c100a01d462759df684db5eddb04d033964b5aefbaeb7b8c7fa9ab57f321704cf9c927ead77c20a32f9662637ab92b4aee03ebf296fe06670370d90f09c0c0315d2159b9a4660855eef0539a658c02ec911af05a8b98a22bb892920df9096b72427b1031109ca6dac14cb03ed1906cc577793f384a2d58adff3fe73286f26647f79e0b728a320b8a7ffd5a3982f2926893cff7cfa873d2ab498c15435aa09aafbc785a49a4b262f9715dcc4a40946a76168878085a96c21667b801d64bfaa75789cb72901bed7babdbc81c280c90155656bb5a5f43e6d9fdbccf49cc441959dc86fb30386a50807f22598a76fde8734972b3c61e68e654761dbc9969b16b7155564b3327dea3d6a003cf4c351aab48a7e4df6718115cd4f79c3272120d42b661051d30d9ef30525804bd50a71dccd230cf9fcc1425b38b29b450ca324486484e52801d8e2cd4bb30464c1f98b74f2a543c4c4b7b29f093d7d0124669a16fae6f1a9014083a17ab84c8de401edebab4777dacfd0720ab07e19b18e8aa7d91331a8dec5e6ef177eb881e200279db42894d8822262250fbcb770dc9c1016feb30538128d286b6329ff3c8437b81178fdde154333d347186445b19713b4032259191eea6ded44075387d7601452cf59326b9ddedb2d5943939a2eaafcfca85676e873236b98f34c842e76a"

func ExamplePrintTdataFile() {
	settings0, _ := hex.DecodeString(hexSettings0)
	PrintTdataFile(bytes.NewReader(settings0), false)
	// Output:
	// version	1005002
	// partialMD5	afcfca85676e873236b98f34c842e76a
	// correctMD5	true
	// dataLength	1000
	// stream   0	32
	// stream   1	960
}

func ExamplePrintTdataSettings() {
	settings0, _ := hex.DecodeString(hexSettings0)
	PrintTdataSettings(bytes.NewReader(settings0))
	// Output:
	// salt	b0ca901027e08bc58296fa67d2e2383f2e62ff5feecc38fb0939d2a9cc2d4561
	// encrypted	4fb0453b3a675ab56d5e5e2088c5dad7c03ef053cd2e56d2afa5f6582f3ed8be97ccead39c64cc27167eb61f319b73a1b614a97f6f82742090065389cae892f94694c1e3527cdd8e64a2830641bedd4b8af2971a0cf93213182677d0d54dba58a96cbb6e0e279c6160d1d434f7146fd9e4bb8e3c75908b52243db395de506566957da30a2023499f500146975ecac7196a115e5703444d946bcdb5702c221225ef6c3d12f05569b9ed48eda0b4c734232a1420b9dd2d4c5538bc764635fdffe6c0565353a8a1c9abece4ac9dbbd653c3fc83e5359b0c960bc429acb191e03641c3c52502e1a6e2dcf5f4c645c165fe82df7c4f19bb7b3321bde735d6b3d12e307c632cd4c1025bcc53cde0105bf763ee2971d7f9046a788bec6a37d433d140f2c949141dcc2a47ba8dcb3d7a2ff73d3fb3bc3e126584a3287dca4f7c1418c4cf3d7fdabf86effb4ea59480f685ff9ff2c118270a996568065e798b1798418659b081398e40f9b93b2a0d53db12b93d80a4f0ad27ff00be213bdb45e92436abfd6b68761e1d6f5f1db6fa3b093616c5867ae90a27027f408c0a67c501f131a2b548e540bb0eae71558cda2e0c2bf4f79da3928dc52abf529aecc26f088d6c2dcc7b92f660bff54135dd8320c9edcdb2e5894fdce34b4b57c196cb4e68b2e5f957c90406a3e0e580c100a01d462759df684db5eddb04d033964b5aefbaeb7b8c7fa9ab57f321704cf9c927ead77c20a32f9662637ab92b4aee03ebf296fe06670370d90f09c0c0315d2159b9a4660855eef0539a658c02ec911af05a8b98a22bb892920df9096b72427b1031109ca6dac14cb03ed1906cc577793f384a2d58adff3fe73286f26647f79e0b728a320b8a7ffd5a3982f2926893cff7cfa873d2ab498c15435aa09aafbc785a49a4b262f9715dcc4a40946a76168878085a96c21667b801d64bfaa75789cb72901bed7babdbc81c280c90155656bb5a5f43e6d9fdbccf49cc441959dc86fb30386a50807f22598a76fde8734972b3c61e68e654761dbc9969b16b7155564b3327dea3d6a003cf4c351aab48a7e4df6718115cd4f79c3272120d42b661051d30d9ef30525804bd50a71dccd230cf9fcc1425b38b29b450ca324486484e52801d8e2cd4bb30464c1f98b74f2a543c4c4b7b29f093d7d0124669a16fae6f1a9014083a17ab84c8de401edebab4777dacfd0720ab07e19b18e8aa7d91331a8dec5e6ef177eb881e200279db42894d8822262250fbcb770dc9c1016feb30538128d286b6329ff3c8437b81178fdde154333d347186445b19713b4032259191eea6ded44075387d7601452cf59326b9ddedb2d5943939a2ea
}

func ExampleDecryptSettings() {
	settings0, _ := hex.DecodeString(hexSettings0)
	settings, err := ReadTdataSettings(bytes.NewReader(settings0))
	if err != nil {
		log.Fatal(err)
	}
	settingsKey := CreateLocalKey([]byte{}, settings.Salt)
	decrypted, err := DecryptLocal(settings.Encrypted, settingsKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(decrypted))
	// Output:
	// a603000000000003000000c800000032000186a000000035000000c800000043000000c80000005000000005000000060000000000000007000000000000001d0000000000000009000000000000000a000000010000000c000000000000000d5823b34900000058000000000000004a000002a0ffffffff0000000f0000000100000000000001bb0000000e3134392e3135342e3137352e3530000000000000000100000001000001bb00000027323030313a306232383a663233643a663030313a303030303a303030303a303030303a30303061000000000000000200000000000001bb0000000e3134392e3135342e3136372e3530000000000000000200000010000001bb0000000e3134392e3135342e3136372e3531000000000000000200000001000001bb00000027323030313a303637633a303465383a663030323a303030303a303030303a303030303a30303061000000000000000300000000000001bb0000000f3134392e3135342e3137352e313030000000000000000300000001000001bb00000027323030313a306232383a663233643a663030333a303030303a303030303a303030303a30303061000000000000000400000000000001bb0000000e3134392e3135342e3136372e3932000000000000000400000010000001bb0000000e3134392e3135342e3136372e3931000000000000000400000001000001bb00000027323030313a303637633a303465383a663030343a303030303a303030303a303030303a30303061000000000000000400000002000001bb0000000f3134392e3135342e3136342e323530000000000000000400000003000001bb00000027323030313a303637633a303465383a663030343a303030303a303030303a303030303a30303062000000000000000500000001000001bb00000027323030313a306232383a663233663a663030353a303030303a303030303a303030303a30303061000000000000000500000010000001bb0000000d39312e3130382e35362e313330000000000000000500000000000001bb0000000d39312e3130382e35362e313731000000000000000000000019ffffffff000000530000001a0061007000760032002e007300740065006c002e0063006f006d00000057000000000000004f000000050000000000000001000000000000000000000028000000000000005400000000000000000000000000000000000000000000004e89f68d97393267c20000000e0000025a000003c20000025400000390c674053f00000000ea56f1191d328106fbd1
}
