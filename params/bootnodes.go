// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// add enode
var MainnetBootnodes = []string{
	"enode://c00b9f537c2b009228f95bf5a84a5f0c0e0407675efeea136e57cd8a67e16c8abdaa16661a4933bc852cc2424603d7d8f57a1199847cc0f8064dcc58814fe174@172.18.188.158:21211", //
	"enode://c84c9860a017cadda359c2b63c29555811d02bf5839938107878ccce856447f67cc72adbad6837c18f823f4ee0a29d48405082ffb47fd490fa5d8d9f80b8ae78@172.18.188.158:21212", //
	"enode://7428587a4839162af1c2bc93629bcf45162abae8772ad93027dfdd94fddc7c653e085b5f58fd0def2533ca61a3e380f4f5e9e891d26eea8cc11df7dcf4b77189@172.18.188.158:21213", //
	"enode://85cb1e6fb7ae05238ad01d3459c28237ce970e847a1e12408e9c8c7e7edfbedc355d467d3e4ab39cad998c1d389dac3d96fa29859b501bf3cfdf38c054542680@172.18.188.158:21214", //
	"enode://29416b59c5ce177413ab98f03fb307928e477c5437ec658cdd127633b19dc3968399990a58557e62d6ee843bb93f0d18639e42f3ec7ba0ec0869507d8b6ea551@172.18.188.158:21215",
}

var MainnetMasternodes = []string{
	"enode://c00b9f537c2b009228f95bf5a84a5f0c0e0407675efeea136e57cd8a67e16c8abdaa16661a4933bc852cc2424603d7d8f57a1199847cc0f8064dcc58814fe174", //
	"enode://c84c9860a017cadda359c2b63c29555811d02bf5839938107878ccce856447f67cc72adbad6837c18f823f4ee0a29d48405082ffb47fd490fa5d8d9f80b8ae78", //
	"enode://7428587a4839162af1c2bc93629bcf45162abae8772ad93027dfdd94fddc7c653e085b5f58fd0def2533ca61a3e380f4f5e9e891d26eea8cc11df7dcf4b77189", //
	"enode://85cb1e6fb7ae05238ad01d3459c28237ce970e847a1e12408e9c8c7e7edfbedc355d467d3e4ab39cad998c1d389dac3d96fa29859b501bf3cfdf38c054542680", //
	"enode://29416b59c5ce177413ab98f03fb307928e477c5437ec658cdd127633b19dc3968399990a58557e62d6ee843bb93f0d18639e42f3ec7ba0ec0869507d8b6ea551",
}

var TestnetMasternodes = []string{
	"enode://f6fc7cc448a2628cd635612118a2de0dd7a41a2f343fc554900d43f38d1b5bee71a46ac5b8fb743fd054477866bf6a9f8de3e4609e2be0fed8c813ba87476d6b@206.189.151.92:20001" ,
	"enode://52e25aa508466c5b61ed84f91929062766980be125ce9e5cdf093a85c73ef3e10cd76299e1f3b0dc93636f4b939fa1826b0971fe1d380af010f8a1e9deddf7a6@206.189.151.92:20002" ,
	"enode://6dd6f3f1fdbf38b75b39c32e981cf44ae91def91de3bef9a0dcbaa412a7909ea4bf7032eb135f4ceca15de0957bb0b0e04563188661c6ee556836172e7df75a9@206.189.151.92:20003" ,
	"enode://a9b83ea9eba7e6b0bc8238c9c0dc0cc8a8df2b6c12efaf112accfacee68d0bf874b680efc0ba9d0802089ca46772aaf4df5675672e16fc1d9d332c28071f94e8@206.189.151.92:20004" ,
	"enode://79585df29910656ace1ad113fc4668e01bbf9fbe0ffaeba126771013dd388169fc7c880ab287a35ae9301bf185abfbd98f3a27757e5f7b64250dc017230be7bf@206.189.151.92:20005" ,
	"enode://99a488ec68a7c57de542f930c6537f96ba1c3db8ac3f6774bb62ae4ba77519a7c4b93a7c80ee99e5d9ac7b329fc3f0ebf80c49e7616ecde29b5e42a9c147bc72@178.128.195.117:20001",
	"enode://14af195555e043442859e768037c761d856067079b0cece9a6e1c0335b6724762776c880ac9f76aa16e782c92797315b0c7394ffae490e208386a4c311f15f76@178.128.195.117:20002",
	"enode://0b7016dd9abad235c009c127b5b2b0d05d2078c579b27e2529212c4bcf34914c9a4fa4ce3273ca05988ea1330bc1b85adde285a9e89ec763b521ed93efb52a78@178.128.195.117:20003",
	"enode://51193b58eebf8b0614a2d05d223d376a8492a329dc1f9d475be2902009887a82e9faa0d059eba2c75c1967488a215e5baf4aa6cea2df22455fbcde90ec9dc6d6@178.128.195.117:20004",
	"enode://7eeb03ba746f2177f71ce6e27331156e4ae04cef7143c01e7daa014e61b20e1b7822b1056af418242ee7805f11c9862282ded65c85d1d87764517458a2070b87@178.128.195.117:20005",
	"enode://3d2b9d5215a2c542339b2dba860b7c962943fde7c960a43414911fcf9d4907da2e53f11bf950d9d74a87c01ce9498b92b96a4b0a9d10b4a1f9c9774c14ce07e8@206.189.105.82:20001" ,
	"enode://5fbf241a0fad4f4452e6ac9e1d1343db707bc09a86f8c3d5ff2ee84c6f33f9843331051b7788d39e17df4ac65d31946c2461215d2bedd13d0bdaeb2147fcb7aa@206.189.105.82:20002" ,
	"enode://2101572dabc573348e3f31a83bcc6bab8531481f8904556716754ca4a57e93cce10f5b2a3ab1f54d21c763e262a88dcef355d03e9de8152f472d2ab1ec0ea284@206.189.105.82:20003" ,
	"enode://7d86aa13716729476692123fcc6f2a23a1abc64eaef8cca87d823465b64ba013d2350c2439c1c91a3bb1bb8ae856f8e15bff8b0ff0378f96c42ade7a710aa9ad@206.189.105.82:20004" ,
	"enode://c325030337695371f3103764e006b164d64839addc3b101bfc936f97e77dd6076022d43b442e75e3ff21d209d6804cf95e54dbdbb9536de5c11704c948358443@206.189.105.82:20005" ,
	"enode://129a841e7c2fef28d5e086fc750b69f538e4c00f1742b0b71dbc7a837dccb39b97624b9b17e602c044940ae84f2d8dd25a7983fde9083ec230dae5e044e5e5f0@142.93.57.240:20001"   ,
	"enode://0125583fa22682d37723b55835db06c7eb934215bff5b8809cda6aa31f6fce48f85bbaf9c7cbc33f1440309a86b2710c758368aa9ccef47a53dc4c0324e25077@142.93.57.240:20002"   ,
	"enode://6c667f2ce11ccde49e868af71cf1ec20149a1edc3e67722896057bbf6356d5f3352b10feb0cb8f3e97ece3c8a15d0493df62bdc87f977a98c45de5f8bef0c0e7@142.93.57.240:20003"   ,
	"enode://17948d4970c290be7b858bda62244bc153ff06e26f6e952bb639d888a371c0b894a8782d3e075da61b65add6b499db6a4b89edfbb28d211fc2ec3c73b2493ac7@142.93.57.240:20004"   ,
	"enode://58291c8c1cddc7f4394f8b2882ce2dc73d77412088d22366389186e006e3df5314e3cc3dc6a1072d221f84ead2ccd2d62a266affff7df5bf4cfc25f4410b4785@142.93.57.240:20005"   ,

}
// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://f6fc7cc448a2628cd635612118a2de0dd7a41a2f343fc554900d43f38d1b5bee71a46ac5b8fb743fd054477866bf6a9f8de3e4609e2be0fed8c813ba87476d6b@206.189.151.92:20001" ,
	"enode://52e25aa508466c5b61ed84f91929062766980be125ce9e5cdf093a85c73ef3e10cd76299e1f3b0dc93636f4b939fa1826b0971fe1d380af010f8a1e9deddf7a6@206.189.151.92:20002" ,
	"enode://6dd6f3f1fdbf38b75b39c32e981cf44ae91def91de3bef9a0dcbaa412a7909ea4bf7032eb135f4ceca15de0957bb0b0e04563188661c6ee556836172e7df75a9@206.189.151.92:20003" ,
	"enode://a9b83ea9eba7e6b0bc8238c9c0dc0cc8a8df2b6c12efaf112accfacee68d0bf874b680efc0ba9d0802089ca46772aaf4df5675672e16fc1d9d332c28071f94e8@206.189.151.92:20004" ,
	"enode://79585df29910656ace1ad113fc4668e01bbf9fbe0ffaeba126771013dd388169fc7c880ab287a35ae9301bf185abfbd98f3a27757e5f7b64250dc017230be7bf@206.189.151.92:20005" ,
	"enode://99a488ec68a7c57de542f930c6537f96ba1c3db8ac3f6774bb62ae4ba77519a7c4b93a7c80ee99e5d9ac7b329fc3f0ebf80c49e7616ecde29b5e42a9c147bc72@178.128.195.117:20001",
	"enode://14af195555e043442859e768037c761d856067079b0cece9a6e1c0335b6724762776c880ac9f76aa16e782c92797315b0c7394ffae490e208386a4c311f15f76@178.128.195.117:20002",
	"enode://0b7016dd9abad235c009c127b5b2b0d05d2078c579b27e2529212c4bcf34914c9a4fa4ce3273ca05988ea1330bc1b85adde285a9e89ec763b521ed93efb52a78@178.128.195.117:20003",
	"enode://51193b58eebf8b0614a2d05d223d376a8492a329dc1f9d475be2902009887a82e9faa0d059eba2c75c1967488a215e5baf4aa6cea2df22455fbcde90ec9dc6d6@178.128.195.117:20004",
	"enode://7eeb03ba746f2177f71ce6e27331156e4ae04cef7143c01e7daa014e61b20e1b7822b1056af418242ee7805f11c9862282ded65c85d1d87764517458a2070b87@178.128.195.117:20005",
	"enode://3d2b9d5215a2c542339b2dba860b7c962943fde7c960a43414911fcf9d4907da2e53f11bf950d9d74a87c01ce9498b92b96a4b0a9d10b4a1f9c9774c14ce07e8@206.189.105.82:20001" ,
	"enode://5fbf241a0fad4f4452e6ac9e1d1343db707bc09a86f8c3d5ff2ee84c6f33f9843331051b7788d39e17df4ac65d31946c2461215d2bedd13d0bdaeb2147fcb7aa@206.189.105.82:20002" ,
	"enode://2101572dabc573348e3f31a83bcc6bab8531481f8904556716754ca4a57e93cce10f5b2a3ab1f54d21c763e262a88dcef355d03e9de8152f472d2ab1ec0ea284@206.189.105.82:20003" ,
	"enode://7d86aa13716729476692123fcc6f2a23a1abc64eaef8cca87d823465b64ba013d2350c2439c1c91a3bb1bb8ae856f8e15bff8b0ff0378f96c42ade7a710aa9ad@206.189.105.82:20004" ,
	"enode://c325030337695371f3103764e006b164d64839addc3b101bfc936f97e77dd6076022d43b442e75e3ff21d209d6804cf95e54dbdbb9536de5c11704c948358443@206.189.105.82:20005" ,
	"enode://129a841e7c2fef28d5e086fc750b69f538e4c00f1742b0b71dbc7a837dccb39b97624b9b17e602c044940ae84f2d8dd25a7983fde9083ec230dae5e044e5e5f0@142.93.57.240:20001"   ,
	"enode://0125583fa22682d37723b55835db06c7eb934215bff5b8809cda6aa31f6fce48f85bbaf9c7cbc33f1440309a86b2710c758368aa9ccef47a53dc4c0324e25077@142.93.57.240:20002"   ,
	"enode://6c667f2ce11ccde49e868af71cf1ec20149a1edc3e67722896057bbf6356d5f3352b10feb0cb8f3e97ece3c8a15d0493df62bdc87f977a98c45de5f8bef0c0e7@142.93.57.240:20003"   ,
	"enode://17948d4970c290be7b858bda62244bc153ff06e26f6e952bb639d888a371c0b894a8782d3e075da61b65add6b499db6a4b89edfbb28d211fc2ec3c73b2493ac7@142.93.57.240:20004"   ,
	"enode://58291c8c1cddc7f4394f8b2882ce2dc73d77412088d22366389186e006e3df5314e3cc3dc6a1072d221f84ead2ccd2d62a266affff7df5bf4cfc25f4410b4785@142.93.57.240:20005"   ,
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://c00b9f537c2b009228f95bf5a84a5f0c0e0407675efeea136e57cd8a67e16c8abdaa16661a4933bc852cc2424603d7d8f57a1199847cc0f8064dcc58814fe174@172.18.188.159:21211", //
	"enode://c84c9860a017cadda359c2b63c29555811d02bf5839938107878ccce856447f67cc72adbad6837c18f823f4ee0a29d48405082ffb47fd490fa5d8d9f80b8ae78@172.18.188.159:21212", //
	"enode://883082b86d3f9224dcf51cd2d577ae049dc63b11d2e25a63ece48b4d27e571320061553bd069210388d23392b98f35e61fcbf1a3d5831a2a187ef979bbf726fe@172.18.188.159:21213", //
	"enode://92170f06c38080f71d2252faeda773348d5d2469893613ae826d03efcf780495c00e312602b9c3be91733898946ed3f6464eb469e47670cf744709da893ac278@172.18.188.159:21214", //
	"enode://a85a8b59e68df8aceb0249ad73f9e19ab707bad3bf0fe7723b9f23e9577902d095231e38053d02880536ce7f43d749e21952bb792e871e3bada99faef9342655@172.18.188.159:21215",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://f6fc7cc448a2628cd635612118a2de0dd7a41a2f343fc554900d43f38d1b5bee71a46ac5b8fb743fd054477866bf6a9f8de3e4609e2be0fed8c813ba87476d6b@206.189.151.92:20001" ,
	"enode://52e25aa508466c5b61ed84f91929062766980be125ce9e5cdf093a85c73ef3e10cd76299e1f3b0dc93636f4b939fa1826b0971fe1d380af010f8a1e9deddf7a6@206.189.151.92:20002" ,
	"enode://6dd6f3f1fdbf38b75b39c32e981cf44ae91def91de3bef9a0dcbaa412a7909ea4bf7032eb135f4ceca15de0957bb0b0e04563188661c6ee556836172e7df75a9@206.189.151.92:20003" ,
	"enode://a9b83ea9eba7e6b0bc8238c9c0dc0cc8a8df2b6c12efaf112accfacee68d0bf874b680efc0ba9d0802089ca46772aaf4df5675672e16fc1d9d332c28071f94e8@206.189.151.92:20004" ,
	"enode://79585df29910656ace1ad113fc4668e01bbf9fbe0ffaeba126771013dd388169fc7c880ab287a35ae9301bf185abfbd98f3a27757e5f7b64250dc017230be7bf@206.189.151.92:20005" ,
	"enode://99a488ec68a7c57de542f930c6537f96ba1c3db8ac3f6774bb62ae4ba77519a7c4b93a7c80ee99e5d9ac7b329fc3f0ebf80c49e7616ecde29b5e42a9c147bc72@178.128.195.117:20001",
	"enode://14af195555e043442859e768037c761d856067079b0cece9a6e1c0335b6724762776c880ac9f76aa16e782c92797315b0c7394ffae490e208386a4c311f15f76@178.128.195.117:20002",
	"enode://0b7016dd9abad235c009c127b5b2b0d05d2078c579b27e2529212c4bcf34914c9a4fa4ce3273ca05988ea1330bc1b85adde285a9e89ec763b521ed93efb52a78@178.128.195.117:20003",
	"enode://51193b58eebf8b0614a2d05d223d376a8492a329dc1f9d475be2902009887a82e9faa0d059eba2c75c1967488a215e5baf4aa6cea2df22455fbcde90ec9dc6d6@178.128.195.117:20004",
	"enode://7eeb03ba746f2177f71ce6e27331156e4ae04cef7143c01e7daa014e61b20e1b7822b1056af418242ee7805f11c9862282ded65c85d1d87764517458a2070b87@178.128.195.117:20005",
	"enode://3d2b9d5215a2c542339b2dba860b7c962943fde7c960a43414911fcf9d4907da2e53f11bf950d9d74a87c01ce9498b92b96a4b0a9d10b4a1f9c9774c14ce07e8@206.189.105.82:20001" ,
	"enode://5fbf241a0fad4f4452e6ac9e1d1343db707bc09a86f8c3d5ff2ee84c6f33f9843331051b7788d39e17df4ac65d31946c2461215d2bedd13d0bdaeb2147fcb7aa@206.189.105.82:20002" ,
	"enode://2101572dabc573348e3f31a83bcc6bab8531481f8904556716754ca4a57e93cce10f5b2a3ab1f54d21c763e262a88dcef355d03e9de8152f472d2ab1ec0ea284@206.189.105.82:20003" ,
	"enode://7d86aa13716729476692123fcc6f2a23a1abc64eaef8cca87d823465b64ba013d2350c2439c1c91a3bb1bb8ae856f8e15bff8b0ff0378f96c42ade7a710aa9ad@206.189.105.82:20004" ,
	"enode://c325030337695371f3103764e006b164d64839addc3b101bfc936f97e77dd6076022d43b442e75e3ff21d209d6804cf95e54dbdbb9536de5c11704c948358443@206.189.105.82:20005" ,
	"enode://129a841e7c2fef28d5e086fc750b69f538e4c00f1742b0b71dbc7a837dccb39b97624b9b17e602c044940ae84f2d8dd25a7983fde9083ec230dae5e044e5e5f0@142.93.57.240:20001"   ,
	"enode://0125583fa22682d37723b55835db06c7eb934215bff5b8809cda6aa31f6fce48f85bbaf9c7cbc33f1440309a86b2710c758368aa9ccef47a53dc4c0324e25077@142.93.57.240:20002"   ,
	"enode://6c667f2ce11ccde49e868af71cf1ec20149a1edc3e67722896057bbf6356d5f3352b10feb0cb8f3e97ece3c8a15d0493df62bdc87f977a98c45de5f8bef0c0e7@142.93.57.240:20003"   ,
	"enode://17948d4970c290be7b858bda62244bc153ff06e26f6e952bb639d888a371c0b894a8782d3e075da61b65add6b499db6a4b89edfbb28d211fc2ec3c73b2493ac7@142.93.57.240:20004"   ,
	"enode://58291c8c1cddc7f4394f8b2882ce2dc73d77412088d22366389186e006e3df5314e3cc3dc6a1072d221f84ead2ccd2d62a266affff7df5bf4cfc25f4410b4785@142.93.57.240:20005"   ,
}