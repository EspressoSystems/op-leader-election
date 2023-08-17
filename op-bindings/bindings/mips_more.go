// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const MIPSStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/cannon/MIPS.sol:MIPS\",\"label\":\"oracle\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_contract(IPreimageOracle)1001\"}],\"types\":{\"t_contract(IPreimageOracle)1001\":{\"encoding\":\"inplace\",\"label\":\"contract IPreimageOracle\",\"numberOfBytes\":\"20\"}}}"

var MIPSStorageLayout = new(solc.StorageLayout)

var MIPSDeployedBin = "0x608060405234801561001057600080fd5b50600436106100415760003560e01c8063155633fe146100465780637dc0d1d01461006b578063f8e0cb96146100b0575b600080fd5b610051634000000081565b60405163ffffffff90911681526020015b60405180910390f35b60005461008b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610062565b6100c36100be366004611baa565b6100d1565b604051908152602001610062565b60006100db611ad7565b608081146100e857600080fd5b604051610600146100f857600080fd5b6064861461010557600080fd5b610184841461011357600080fd5b8535608052602086013560a052604086013560e090811c60c09081526044880135821c82526048880135821c61010052604c880135821c610120526050880135821c61014052605488013590911c61016052605887013560f890811c610180526059880135901c6101a052605a870135901c6101c0526102006101e0819052606287019060005b60208110156101be57823560e01c825260049092019160209091019060010161019a565b505050806101200151156101dc576101d4610612565b91505061060a565b6101408101805160010167ffffffffffffffff169052606081015160009061020490826106ba565b9050603f601a82901c16600281148061022357508063ffffffff166003145b15610270576102668163ffffffff1660021461024057601f610243565b60005b60ff166002610259856303ffffff16601a610776565b63ffffffff16901b6107e9565b935050505061060a565b6101608301516000908190601f601086901c81169190601587901c166020811061029c5761029c611c16565b602002015192508063ffffffff851615806102bd57508463ffffffff16601c145b156102f4578661016001518263ffffffff16602081106102df576102df611c16565b6020020151925050601f600b86901c166103b0565b60208563ffffffff161015610356578463ffffffff16600c148061031e57508463ffffffff16600d145b8061032f57508463ffffffff16600e145b15610340578561ffff1692506103b0565b61034f8661ffff166010610776565b92506103b0565b60288563ffffffff1610158061037257508463ffffffff166022145b8061038357508463ffffffff166026145b156103b0578661016001518263ffffffff16602081106103a5576103a5611c16565b602002015192508190505b60048563ffffffff16101580156103cd575060088563ffffffff16105b806103de57508463ffffffff166001145b156103fd576103ef858784876108e3565b97505050505050505061060a565b63ffffffff60006020878316106104625761041d8861ffff166010610776565b9095019463fffffffc86166104338160016106ba565b915060288863ffffffff161015801561045357508763ffffffff16603014155b1561046057809250600093505b505b600061047089888885610af3565b63ffffffff9081169150603f8a16908916158015610495575060088163ffffffff1610155b80156104a75750601c8163ffffffff16105b15610583578063ffffffff16600814806104c757508063ffffffff166009145b156104fe576104ec8163ffffffff166008146104e357856104e6565b60005b896107e9565b9b50505050505050505050505061060a565b8063ffffffff16600a0361051e576104ec858963ffffffff8a1615611196565b8063ffffffff16600b0361053f576104ec858963ffffffff8a161515611196565b8063ffffffff16600c03610555576104ec61127c565b60108163ffffffff16101580156105725750601c8163ffffffff16105b15610583576104ec81898988611790565b8863ffffffff16603814801561059e575063ffffffff861615155b156105d35760018b61016001518763ffffffff16602081106105c2576105c2611c16565b63ffffffff90921660209290920201525b8363ffffffff1663ffffffff146105f0576105f08460018461198a565b6105fc85836001611196565b9b5050505050505050505050505b949350505050565b60408051608051815260a051602082015260dc519181019190915260fc51604482015261011c51604882015261013c51604c82015261015c51605082015261017c51605482015261019f5160588201526101bf5160598201526101d851605a8201526000906102009060628101835b60208110156106a557601c8401518252602090930192600490910190600101610681565b506000815281810382a0819003902092915050565b6000806106c683611a2e565b905060038416156106d657600080fd5b6020810190358460051c8160005b601b81101561073c5760208501943583821c600116801561070c576001811461072157610732565b60008481526020839052604090209350610732565b600082815260208590526040902093505b50506001016106e4565b50608051915081811461075757630badf00d60005260206000fd5b5050601f94909416601c0360031b9390931c63ffffffff169392505050565b600063ffffffff8381167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80850183169190911c821615159160016020869003821681901b830191861691821b92911b01826107d35760006107d5565b815b90861663ffffffff16179250505092915050565b60006107f3611ad7565b60809050806060015160040163ffffffff16816080015163ffffffff161461087c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6a756d7020696e2064656c617920736c6f74000000000000000000000000000060448201526064015b60405180910390fd5b60608101805160808301805163ffffffff9081169093528583169052908516156108d257806008018261016001518663ffffffff16602081106108c1576108c1611c16565b63ffffffff90921660209290920201525b6108da610612565b95945050505050565b60006108ed611ad7565b608090506000816060015160040163ffffffff16826080015163ffffffff1614610973576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6272616e636820696e2064656c617920736c6f740000000000000000000000006044820152606401610873565b8663ffffffff166004148061098e57508663ffffffff166005145b15610a0a5760008261016001518663ffffffff16602081106109b2576109b2611c16565b602002015190508063ffffffff168563ffffffff161480156109da57508763ffffffff166004145b80610a0257508063ffffffff168563ffffffff1614158015610a0257508763ffffffff166005145b915050610a87565b8663ffffffff16600603610a275760008460030b13159050610a87565b8663ffffffff16600703610a435760008460030b139050610a87565b8663ffffffff16600103610a8757601f601087901c166000819003610a6c5760008560030b1291505b8063ffffffff16600103610a855760008560030b121591505b505b606082018051608084015163ffffffff169091528115610acd576002610ab28861ffff166010610776565b63ffffffff90811690911b8201600401166080840152610adf565b60808301805160040163ffffffff1690525b610ae7610612565b98975050505050505050565b6000603f601a86901c81169086166020821015610eb75760088263ffffffff1610158015610b275750600f8263ffffffff16105b15610bc7578163ffffffff16600803610b4257506020610bc2565b8163ffffffff16600903610b5857506021610bc2565b8163ffffffff16600a03610b6e5750602a610bc2565b8163ffffffff16600b03610b845750602b610bc2565b8163ffffffff16600c03610b9a57506024610bc2565b8163ffffffff16600d03610bb057506025610bc2565b8163ffffffff16600e03610bc2575060265b600091505b8163ffffffff16600003610e0b57601f600688901c16602063ffffffff83161015610ce55760088263ffffffff1610610c055786935050505061060a565b8163ffffffff16600003610c285763ffffffff86811691161b925061060a915050565b8163ffffffff16600203610c4b5763ffffffff86811691161c925061060a915050565b8163ffffffff16600303610c75576102668163ffffffff168763ffffffff16901c82602003610776565b8163ffffffff16600403610c98575050505063ffffffff8216601f84161b61060a565b8163ffffffff16600603610cbb575050505063ffffffff8216601f84161c61060a565b8163ffffffff16600703610ce5576102668763ffffffff168763ffffffff16901c88602003610776565b8163ffffffff1660201480610d0057508163ffffffff166021145b15610d1257858701935050505061060a565b8163ffffffff1660221480610d2d57508163ffffffff166023145b15610d3f57858703935050505061060a565b8163ffffffff16602403610d5a57858716935050505061060a565b8163ffffffff16602503610d7557858717935050505061060a565b8163ffffffff16602603610d9057858718935050505061060a565b8163ffffffff16602703610dab57505050508282171961060a565b8163ffffffff16602a03610ddd578560030b8760030b12610dcd576000610dd0565b60015b60ff16935050505061060a565b8163ffffffff16602b03610e05578563ffffffff168763ffffffff1610610dcd576000610dd0565b50611134565b8163ffffffff16600f03610e2d5760108563ffffffff16901b9250505061060a565b8163ffffffff16601c03610eb2578063ffffffff16600203610e545750505082820261060a565b8063ffffffff1660201480610e6f57508063ffffffff166021145b15610eb2578063ffffffff16602003610e86579419945b60005b6380000000871615610ea8576401fffffffe600197881b169601610e89565b925061060a915050565b611134565b60288263ffffffff16101561101a578163ffffffff16602003610f0357610efa8660031660080260180363ffffffff168563ffffffff16901c60ff166008610776565b9250505061060a565b8163ffffffff16602103610f3857610efa8660021660080260100363ffffffff168563ffffffff16901c61ffff166010610776565b8163ffffffff16602203610f685750505063ffffffff60086003851602811681811b198416918316901b1761060a565b8163ffffffff16602303610f8057839250505061060a565b8163ffffffff16602403610fb3578560031660080260180363ffffffff168463ffffffff16901c60ff169250505061060a565b8163ffffffff16602503610fe7578560021660080260100363ffffffff168463ffffffff16901c61ffff169250505061060a565b8163ffffffff16602603610eb25750505063ffffffff60086003851602601803811681811c198416918316901c1761060a565b8163ffffffff166028036110515750505060ff63ffffffff60086003861602601803811682811b9091188316918416901b1761060a565b8163ffffffff166029036110895750505061ffff63ffffffff60086002861602601003811682811b9091188316918416901b1761060a565b8163ffffffff16602a036110b95750505063ffffffff60086003851602811681811c198316918416901c1761060a565b8163ffffffff16602b036110d157849250505061060a565b8163ffffffff16602e036111045750505063ffffffff60086003851602601803811681811b198316918416901b1761060a565b8163ffffffff1660300361111c57839250505061060a565b8163ffffffff1660380361113457849250505061060a565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f696e76616c696420696e737472756374696f6e000000000000000000000000006044820152606401610873565b60006111a0611ad7565b506080602063ffffffff861610611213576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f76616c69642072656769737465720000000000000000000000000000000000006044820152606401610873565b63ffffffff8516158015906112255750825b1561125957838161016001518663ffffffff166020811061124857611248611c16565b63ffffffff90921660209290920201525b60808101805163ffffffff808216606085015260049091011690526108da610612565b6000611286611ad7565b506101e051604081015160808083015160a084015160c09094015191936000928392919063ffffffff8616610ffa036113005781610fff8116156112cf57610fff811661100003015b8363ffffffff166000036112f65760e08801805163ffffffff8382011690915295506112fa565b8395505b5061174f565b8563ffffffff16610fcd0361131b576340000000945061174f565b8563ffffffff1661101803611333576001945061174f565b8563ffffffff166110960361136857600161012088015260ff831661010088015261135c610612565b97505050505050505090565b8563ffffffff16610fa3036115b25763ffffffff83161561174f577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb63ffffffff84160161156c5760006113c38363fffffffc1660016106ba565b60208901519091508060001a6001036114305761142d81600090815233602052604090207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001790565b90505b6000805460408b81015190517fe03110e10000000000000000000000000000000000000000000000000000000081526004810185905263ffffffff9091166024820152829173ffffffffffffffffffffffffffffffffffffffff169063e03110e1906044016040805180830381865afa1580156114b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114d59190611c45565b915091506003861680600403828110156114ed578092505b50818610156114fa578591505b8260088302610100031c9250826008828460040303021b9250600180600883600403021b036001806008858560040303021b039150811981169050838119871617955050506115518663fffffffc1660018661198a565b60408b018051820163ffffffff16905297506115ad92505050565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd63ffffffff8416016115a15780945061174f565b63ffffffff9450600993505b61174f565b8563ffffffff16610fa4036116a35763ffffffff8316600114806115dc575063ffffffff83166002145b806115ed575063ffffffff83166004145b156115fa5780945061174f565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa63ffffffff8416016115a157600061163a8363fffffffc1660016106ba565b60208901519091506003841660040383811015611655578093505b83900360089081029290921c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600193850293841b0116911b1760208801526000604088015293508361174f565b8563ffffffff16610fd70361174f578163ffffffff166003036117435763ffffffff831615806116d9575063ffffffff83166005145b806116ea575063ffffffff83166003145b156116f8576000945061174f565b63ffffffff831660011480611713575063ffffffff83166002145b80611724575063ffffffff83166006145b80611735575063ffffffff83166004145b156115a1576001945061174f565b63ffffffff9450601693505b6101608701805163ffffffff808816604090920191909152905185821660e09091015260808801805180831660608b0152600401909116905261135c610612565b600061179a611ad7565b506080600063ffffffff87166010036117b8575060c0810151611921565b8663ffffffff166011036117d75763ffffffff861660c0830152611921565b8663ffffffff166012036117f0575060a0810151611921565b8663ffffffff1660130361180f5763ffffffff861660a0830152611921565b8663ffffffff166018036118435763ffffffff600387810b9087900b02602081901c821660c08501521660a0830152611921565b8663ffffffff166019036118745763ffffffff86811681871602602081901c821660c08501521660a0830152611921565b8663ffffffff16601a036118ca578460030b8660030b8161189757611897611c69565b0763ffffffff1660c0830152600385810b9087900b816118b9576118b9611c69565b0563ffffffff1660a0830152611921565b8663ffffffff16601b03611921578463ffffffff168663ffffffff16816118f3576118f3611c69565b0663ffffffff90811660c08401528581169087168161191457611914611c69565b0463ffffffff1660a08301525b63ffffffff84161561195c57808261016001518563ffffffff166020811061194b5761194b611c16565b63ffffffff90921660209290920201525b60808201805163ffffffff8082166060860152600490910116905261197f610612565b979650505050505050565b600061199583611a2e565b905060038416156119a557600080fd5b6020810190601f8516601c0360031b83811b913563ffffffff90911b1916178460051c60005b601b811015611a235760208401933582821c60011680156119f35760018114611a0857611a19565b60008581526020839052604090209450611a19565b600082815260208690526040902094505b50506001016119cb565b505060805250505050565b60ff811661038002610184810190369061050401811015611ad1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f636865636b207468617420746865726520697320656e6f7567682063616c6c6460448201527f61746100000000000000000000000000000000000000000000000000000000006064820152608401610873565b50919050565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101919091526101608101611b3d611b42565b905290565b6040518061040001604052806020906020820280368337509192915050565b60008083601f840112611b7357600080fd5b50813567ffffffffffffffff811115611b8b57600080fd5b602083019150836020828501011115611ba357600080fd5b9250929050565b60008060008060408587031215611bc057600080fd5b843567ffffffffffffffff80821115611bd857600080fd5b611be488838901611b61565b90965094506020870135915080821115611bfd57600080fd5b50611c0a87828801611b61565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008060408385031215611c5857600080fd5b505080516020909101519092909150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fdfea164736f6c634300080f000a"

var MIPSDeployedSourceMap = "1131:37218:108:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1710:45;;1745:10;1710:45;;;;;188:10:260;176:23;;;158:42;;146:2;131:18;1710:45:108;;;;;;;;2144:29;;;;;;;;;;;;412:42:260;400:55;;;382:74;;370:2;355:18;2144:29:108;211:251:260;24696:6377:108;;;;;;:::i;:::-;;:::i;:::-;;;1687:25:260;;;1675:2;1660:18;24696:6377:108;1541:177:260;24696:6377:108;24774:7;24817:18;;:::i;:::-;24964:4;24957:5;24954:15;24944:134;;25058:1;25055;25048:12;24944:134;25114:4;25108:11;25121;25105:28;25095:137;;25212:1;25209;25202:12;25095:137;25280:3;25262:16;25259:25;25249:150;;25379:1;25376;25369:12;25249:150;25443:3;25429:12;25426:21;25416:145;;25541:1;25538;25531:12;25416:145;25821:24;;26165:4;25867:20;26223:2;25925:21;;25821:24;25983:18;25867:20;25925:21;;;25821:24;25798:21;25794:52;;;25983:18;25867:20;;;25925:21;;;25821:24;25794:52;;25867:20;;25925:21;;;25821:24;25794:52;;25983:18;25867:20;25925:21;;;25821:24;25794:52;;25983:18;25867:20;25925:21;;;25821:24;25794:52;;25983:18;25867:20;25925:21;;;25821:24;25794:52;;;25983:18;25867:20;25925:21;;;25821:24;25798:21;25794:52;;;25983:18;25867:20;25925:21;;;25821:24;25794:52;;25983:18;25867:20;25925:21;;;25821:24;25794:52;;25983:18;25867:20;26841:10;25983:18;26831:21;;;25925;;;;26939:1;26924:77;26949:2;26946:1;26943:9;26924:77;;;25821:24;;25798:21;25794:52;25867:20;;26997:1;25925:21;;;;25809:2;25983:18;;;;26967:1;26960:9;26924:77;;;26928:14;;;27079:5;:12;;;27075:71;;;27118:13;:11;:13::i;:::-;27111:20;;;;;27075:71;27160:10;;;:15;;27174:1;27160:15;;;;;27245:8;;;;-1:-1:-1;;27237:20:108;;-1:-1:-1;27237:7:108;:20::i;:::-;27223:34;-1:-1:-1;27287:10:108;27295:2;27287:10;;;;27364:1;27354:11;;;:26;;;27369:6;:11;;27379:1;27369:11;27354:26;27350:348;;;27619:64;27630:6;:11;;27640:1;27630:11;:20;;27648:2;27630:20;;;27644:1;27630:20;27619:64;;27681:1;27652:25;27655:4;27662:10;27655:17;27674:2;27652;:25::i;:::-;:30;;;;27619:10;:64::i;:::-;27612:71;;;;;;;27350:348;27947:15;;;;27742:9;;;;27879:4;27873:2;27865:10;;;27864:19;;;27947:15;27972:2;27964:10;;;27963:19;27947:36;;;;;;;:::i;:::-;;;;;;-1:-1:-1;28012:5:108;28036:11;;;;;:29;;;28051:6;:14;;28061:4;28051:14;28036:29;28032:832;;;28128:5;:15;;;28144:5;28128:22;;;;;;;;;:::i;:::-;;;;;;-1:-1:-1;;28191:4:108;28185:2;28177:10;;;28176:19;28032:832;;;28229:4;28220:6;:13;;;28216:648;;;28350:6;:13;;28360:3;28350:13;:30;;;;28367:6;:13;;28377:3;28367:13;28350:30;:47;;;;28384:6;:13;;28394:3;28384:13;28350:47;28346:253;;;28460:4;28467:6;28460:13;28455:18;;28216:648;;28346:253;28559:21;28562:4;28569:6;28562:13;28577:2;28559;:21::i;:::-;28554:26;;28216:648;;;28633:4;28623:6;:14;;;;:32;;;;28641:6;:14;;28651:4;28641:14;28623:32;:50;;;;28659:6;:14;;28669:4;28659:14;28623:50;28619:245;;;28743:5;:15;;;28759:5;28743:22;;;;;;;;;:::i;:::-;;;;;28738:27;;28844:5;28836:13;;28619:245;28893:1;28883:6;:11;;;;:25;;;;;28907:1;28898:6;:10;;;28883:25;28882:42;;;;28913:6;:11;;28923:1;28913:11;28882:42;28878:125;;;28951:37;28964:6;28972:4;28978:5;28985:2;28951:12;:37::i;:::-;28944:44;;;;;;;;;;;28878:125;29036:13;29017:16;29188:4;29178:14;;;;29174:446;;29257:21;29260:4;29267:6;29260:13;29275:2;29257;:21::i;:::-;29251:27;;;;29315:10;29310:15;;29349:16;29310:15;29363:1;29349:7;:16::i;:::-;29343:22;;29397:4;29387:6;:14;;;;:32;;;;;29405:6;:14;;29415:4;29405:14;;29387:32;29383:223;;;29484:4;29472:16;;29586:1;29578:9;;29383:223;29194:426;29174:446;29653:10;29666:26;29674:4;29680:2;29684;29688:3;29666:7;:26::i;:::-;29695:10;29666:39;;;;-1:-1:-1;29791:4:108;29784:11;;;29823;;;:24;;;;;29846:1;29838:4;:9;;;;29823:24;:39;;;;;29858:4;29851;:11;;;29823:39;29819:847;;;29886:4;:9;;29894:1;29886:9;:22;;;;29899:4;:9;;29907:1;29899:9;29886:22;29882:144;;;29970:37;29981:4;:9;;29989:1;29981:9;:21;;29997:5;29981:21;;;29993:1;29981:21;30004:2;29970:10;:37::i;:::-;29963:44;;;;;;;;;;;;;;;29882:144;30048:4;:11;;30056:3;30048:11;30044:121;;30118:28;30127:5;30134:2;30138:7;;;;30118:8;:28::i;30044:121::-;30186:4;:11;;30194:3;30186:11;30182:121;;30256:28;30265:5;30272:2;30276:7;;;;;30256:8;:28::i;30182:121::-;30373:4;:11;;30381:3;30373:11;30369:80;;30415:15;:13;:15::i;30369:80::-;30552:4;30544;:12;;;;:27;;;;;30567:4;30560;:11;;;30544:27;30540:112;;;30602:31;30613:4;30619:2;30623;30627:5;30602:10;:31::i;30540:112::-;30726:6;:14;;30736:4;30726:14;:28;;;;-1:-1:-1;30744:10:108;;;;;30726:28;30722:93;;;30799:1;30774:5;:15;;;30790:5;30774:22;;;;;;;;;:::i;:::-;:26;;;;:22;;;;;;:26;30722:93;30861:9;:26;;30874:13;30861:26;30857:92;;30907:27;30916:9;30927:1;30930:3;30907:8;:27::i;:::-;31030:26;31039:5;31046:3;31051:4;31030:8;:26::i;:::-;31023:33;;;;;;;;;;;;;24696:6377;;;;;;;:::o;2858:1709::-;3405:4;3399:11;;3321:4;3124:31;3113:43;;3184:13;3124:31;3523:2;3223:13;;3113:43;3130:24;3124:31;3223:13;;;3113:43;;;;3130:24;3124:31;3223:13;;;3113:43;3130:24;3124:31;3223:13;;;3113:43;3130:24;3124:31;3223:13;;;3113:43;3130:24;3124:31;3223:13;;;3113:43;3130:24;3124:31;3223:13;;;3113:43;3130:24;3124:31;3223:13;;;3113:43;3130:24;3124:31;3223:13;;;3113:43;3130:24;3124:31;3223:13;;;3113:43;2899:12;;4108:13;;3223;;;2899:12;4188:84;4213:2;4210:1;4207:9;4188:84;;;3140:13;3130:24;;3124:31;3113:43;;3144:2;3184:13;;;;4268:1;3223:13;;;;4231:1;4224:9;4188:84;;;4192:14;4335:1;4331:2;4324:13;4430:5;4426:2;4422:14;4415:5;4410:27;4536:14;;;4519:32;;;2858:1709;-1:-1:-1;;2858:1709:108:o;20751:1831::-;20824:11;20935:14;20952:24;20964:11;20952;:24::i;:::-;20935:41;;21084:1;21077:5;21073:13;21070:33;;;21099:1;21096;21089:12;21070:33;21232:2;21220:15;;;21173:20;21662:5;21659:1;21655:13;21697:4;21733:1;21718:343;21743:2;21740:1;21737:9;21718:343;;;21866:2;21854:15;;;21803:20;21901:12;;;21915:1;21897:20;21938:42;;;;22006:1;22001:42;;;;21890:153;;21938:42;21396:1;21389:12;;;21429:2;21422:13;;;21474:2;21461:16;;21947:31;;21938:42;;22001;21396:1;21389:12;;;21429:2;21422:13;;;21474:2;21461:16;;22010:31;;21890:153;-1:-1:-1;;21761:1:108;21754:9;21718:343;;;21722:14;22171:4;22165:11;22150:26;;22257:7;22251:4;22248:17;22238:124;;22299:10;22296:1;22289:21;22341:2;22338:1;22331:13;22238:124;-1:-1:-1;;22489:2:108;22478:14;;;;22466:10;22462:31;22459:1;22455:39;22523:16;;;;22541:10;22519:33;;20751:1831;-1:-1:-1;;;20751:1831:108:o;2416:334::-;2477:6;2536:18;;;;2545:8;;;;2536:18;;;;;;2535:25;;;;;2552:1;2599:2;:9;;;2593:16;;;;;2592:22;;2591:32;;;;;;;2653:9;;2652:15;2535:25;2710:21;;2730:1;2710:21;;;2721:6;2710:21;2695:11;;;;;:37;;-1:-1:-1;;;2416:334:108;;;;:::o;17861:823::-;17930:12;18017:18;;:::i;:::-;18085:4;18076:13;;18137:5;:8;;;18148:1;18137:12;18121:28;;:5;:12;;;:28;;;18117:95;;18169:28;;;;;2114:2:260;18169:28:108;;;2096:21:260;2153:2;2133:18;;;2126:30;2192:20;2172:18;;;2165:48;2230:18;;18169:28:108;;;;;;;;18117:95;18301:8;;;;;18334:12;;;;;18323:23;;;;;;;18360:20;;;;;18301:8;18492:13;;;18488:90;;18553:6;18562:1;18553:10;18525:5;:15;;;18541:8;18525:25;;;;;;;;;:::i;:::-;:38;;;;:25;;;;;;:38;18488:90;18654:13;:11;:13::i;:::-;18647:20;17861:823;-1:-1:-1;;;;;17861:823:108:o;12722:2026::-;12819:12;12905:18;;:::i;:::-;12973:4;12964:13;;13005:17;13065:5;:8;;;13076:1;13065:12;13049:28;;:5;:12;;;:28;;;13045:97;;13097:30;;;;;2461:2:260;13097:30:108;;;2443:21:260;2500:2;2480:18;;;2473:30;2539:22;2519:18;;;2512:50;2579:18;;13097:30:108;2259:344:260;13045:97:108;13212:7;:12;;13223:1;13212:12;:28;;;;13228:7;:12;;13239:1;13228:12;13212:28;13208:947;;;13260:9;13272:5;:15;;;13288:6;13272:23;;;;;;;;;:::i;:::-;;;;;13260:35;;13336:2;13329:9;;:3;:9;;;:25;;;;;13342:7;:12;;13353:1;13342:12;13329:25;13328:58;;;;13367:2;13360:9;;:3;:9;;;;:25;;;;;13373:7;:12;;13384:1;13373:12;13360:25;13313:73;;13242:159;13208:947;;;13498:7;:12;;13509:1;13498:12;13494:661;;13559:1;13551:3;13545:15;;;;13530:30;;13494:661;;;13663:7;:12;;13674:1;13663:12;13659:496;;13723:1;13716:3;13710:14;;;13695:29;;13659:496;;;13844:7;:12;;13855:1;13844:12;13840:315;;13932:4;13926:2;13917:11;;;13916:20;13902:10;13959:8;;;13955:84;;14019:1;14012:3;14006:14;;;13991:29;;13955:84;14060:3;:8;;14067:1;14060:8;14056:85;;14121:1;14113:3;14107:15;;;;14092:30;;14056:85;13858:297;13840:315;14231:8;;;;;14309:12;;;;14298:23;;;;;14465:178;;;;14556:1;14530:22;14533:5;14541:6;14533:14;14549:2;14530;:22::i;:::-;:27;;;;;;;14516:42;;14525:1;14516:42;14501:57;:12;;;:57;14465:178;;;14612:12;;;;;14627:1;14612:16;14597:31;;;;14465:178;14718:13;:11;:13::i;:::-;14711:20;12722:2026;-1:-1:-1;;;;;;;;12722:2026:108:o;31119:7228::-;31206:6;31264:10;31272:2;31264:10;;;;;;31312:11;;31424:4;31415:13;;31411:6876;;;31555:1;31545:6;:11;;;;:27;;;;;31569:3;31560:6;:12;;;31545:27;31541:537;;;31600:6;:11;;31610:1;31600:11;31596:423;;-1:-1:-1;31620:4:108;31596:423;;;31664:6;:11;;31674:1;31664:11;31660:359;;-1:-1:-1;31684:4:108;31660:359;;;31729:6;:13;;31739:3;31729:13;31725:294;;-1:-1:-1;31751:4:108;31725:294;;;31795:6;:13;;31805:3;31795:13;31791:228;;-1:-1:-1;31817:4:108;31791:228;;;31862:6;:13;;31872:3;31862:13;31858:161;;-1:-1:-1;31884:4:108;31858:161;;;31928:6;:13;;31938:3;31928:13;31924:95;;-1:-1:-1;31950:4:108;31924:95;;;31993:6;:13;;32003:3;31993:13;31989:30;;-1:-1:-1;32015:4:108;31989:30;32058:1;32049:10;;31541:537;32139:6;:11;;32149:1;32139:11;32135:3554;;32203:4;32198:1;32190:9;;;32189:18;32240:4;32190:9;32233:11;;;32229:1319;;;32332:4;32324;:12;;;32320:1206;;32375:2;32368:9;;;;;;;32320:1206;32489:4;:12;;32497:4;32489:12;32485:1041;;32540:11;;;;;;;;-1:-1:-1;32533:18:108;;-1:-1:-1;;32533:18:108;32485:1041;32664:4;:12;;32672:4;32664:12;32660:866;;32715:11;;;;;;;;-1:-1:-1;32708:18:108;;-1:-1:-1;;32708:18:108;32660:866;32842:4;:12;;32850:4;32842:12;32838:688;;32893:27;32902:5;32896:11;;:2;:11;;;;32914:5;32909:2;:10;32893:2;:27::i;32838:688::-;33042:4;:12;;33050:4;33042:12;33038:488;;-1:-1:-1;;;;33093:17:108;;;33105:4;33100:9;;33093:17;33086:24;;33038:488;33233:4;:12;;33241:4;33233:12;33229:297;;-1:-1:-1;;;;33284:17:108;;;33296:4;33291:9;;33284:17;33277:24;;33229:297;33427:4;:12;;33435:4;33427:12;33423:103;;33478:21;33487:2;33481:8;;:2;:8;;;;33496:2;33491;:7;33478:2;:21::i;33423:103::-;33708:4;:12;;33716:4;33708:12;:28;;;;33724:4;:12;;33732:4;33724:12;33708:28;33704:1151;;;33776:2;33771;:7;33764:14;;;;;;;33704:1151;33866:4;:12;;33874:4;33866:12;:28;;;;33882:4;:12;;33890:4;33882:12;33866:28;33862:993;;;33934:2;33929;:7;33922:14;;;;;;;33862:993;34016:4;:12;;34024:4;34016:12;34012:843;;34068:2;34063;:7;34056:14;;;;;;;34012:843;34149:4;:12;;34157:4;34149:12;34145:710;;34202:2;34197;:7;34189:16;;;;;;;34145:710;34285:4;:12;;34293:4;34285:12;34281:574;;34338:2;34333;:7;34325:16;;;;;;;34281:574;34421:4;:12;;34429:4;34421:12;34417:438;;-1:-1:-1;;;;34470:7:108;;;34468:10;34461:17;;34417:438;34581:4;:12;;34589:4;34581:12;34577:278;;34646:2;34628:21;;34634:2;34628:21;;;:29;;34656:1;34628:29;;;34652:1;34628:29;34621:36;;;;;;;;;34577:278;34770:4;:12;;34778:4;34770:12;34766:89;;34822:2;34817:7;;:2;:7;;;:15;;34831:1;34817:15;;34766:89;32152:2721;31411:6876;;32135:3554;34944:6;:13;;34954:3;34944:13;34940:749;;34994:2;34988;:8;;;;34981:15;;;;;;34940:749;35069:6;:14;;35079:4;35069:14;35065:624;;35138:4;:9;;35146:1;35138:9;35134:100;;-1:-1:-1;;;35189:21:108;;;35175:36;;35134:100;35286:4;:12;;35294:4;35286:12;:28;;;;35302:4;:12;;35310:4;35302:12;35286:28;35282:389;;;35346:4;:12;;35354:4;35346:12;35342:83;;35395:3;;;35342:83;35450:8;35488:127;35500:10;35495:15;;:20;35488:127;;35580:8;35547:3;35580:8;;;;;35547:3;35488:127;;;35647:1;-1:-1:-1;35640:8:108;;-1:-1:-1;;35640:8:108;35282:389;31411:6876;;;35722:4;35713:6;:13;;;35709:2578;;;35772:6;:14;;35782:4;35772:14;35768:1208;;35817:42;35835:2;35840:1;35835:6;35845:1;35834:12;35829:2;:17;35821:26;;:3;:26;;;;35851:4;35820:35;35857:1;35817:2;:42::i;:::-;35810:49;;;;;;35768:1208;35926:6;:14;;35936:4;35926:14;35922:1054;;35971:45;35989:2;35994:1;35989:6;35999:1;35988:12;35983:2;:17;35975:26;;:3;:26;;;;36005:6;35974:37;36013:2;35971;:45::i;35922:1054::-;36084:6;:14;;36094:4;36084:14;36080:896;;-1:-1:-1;;;36135:21:108;36154:1;36149;36144:6;;36143:12;36135:21;;36192:36;;;36263:5;36258:10;;36135:21;;;;;36257:18;36250:25;;36080:896;36342:6;:14;;36352:4;36342:14;36338:638;;36387:3;36380:10;;;;;;36338:638;36458:6;:14;;36468:4;36458:14;36454:522;;36518:2;36523:1;36518:6;36528:1;36517:12;36512:2;:17;36504:26;;:3;:26;;;;36534:4;36503:35;36496:42;;;;;;36454:522;36606:6;:14;;36616:4;36606:14;36602:374;;36666:2;36671:1;36666:6;36676:1;36665:12;36660:2;:17;36652:26;;:3;:26;;;;36682:6;36651:37;36644:44;;;;;;36602:374;36756:6;:14;;36766:4;36756:14;36752:224;;-1:-1:-1;;;36807:26:108;36831:1;36826;36821:6;;36820:12;36815:2;:17;36807:26;;36869:41;;;36945:5;36940:10;;36807:26;;;;;36939:18;36932:25;;35709:2578;37030:6;:14;;37040:4;37030:14;37026:1261;;-1:-1:-1;;;37083:4:108;37077:34;37109:1;37104;37099:6;;37098:12;37093:2;:17;37077:34;;37163:27;;;37143:48;;;37217:10;;37078:9;;;37077:34;;37216:18;37209:25;;37026:1261;37289:6;:14;;37299:4;37289:14;37285:1002;;-1:-1:-1;;;37342:6:108;37336:36;37370:1;37365;37360:6;;37359:12;37354:2;:17;37336:36;;37424:29;;;37404:50;;;37480:10;;37337:11;;;37336:36;;37479:18;37472:25;;37285:1002;37553:6;:14;;37563:4;37553:14;37549:738;;-1:-1:-1;;;37600:20:108;37618:1;37613;37608:6;;37607:12;37600:20;;37652:36;;;37720:5;37714:11;;37600:20;;;;;37713:19;37706:26;;37549:738;37787:6;:14;;37797:4;37787:14;37783:504;;37828:2;37821:9;;;;;;37783:504;37886:6;:14;;37896:4;37886:14;37882:405;;-1:-1:-1;;;37933:25:108;37956:1;37951;37946:6;;37945:12;37940:2;:17;37933:25;;37990:41;;;38063:5;38057:11;;37933:25;;;;;38056:19;38049:26;;37882:405;38130:6;:14;;38140:4;38130:14;38126:161;;38171:3;38164:10;;;;;;38126:161;38229:6;:14;;38239:4;38229:14;38225:62;;38270:2;38263:9;;;;;;38225:62;38301:29;;;;;2810:2:260;38301:29:108;;;2792:21:260;2849:2;2829:18;;;2822:30;2888:21;2868:18;;;2861:49;2927:18;;38301:29:108;2608:343:260;18965:782:108;19051:12;19138:18;;:::i;:::-;-1:-1:-1;19206:4:108;19313:2;19301:14;;;;19293:41;;;;;;;3158:2:260;19293:41:108;;;3140:21:260;3197:2;3177:18;;;3170:30;3236:16;3216:18;;;3209:44;3270:18;;19293:41:108;2956:338:260;19293:41:108;19430:14;;;;;;;:30;;;19448:12;19430:30;19426:102;;;19509:4;19480:5;:15;;;19496:9;19480:26;;;;;;;;;:::i;:::-;:33;;;;:26;;;;;;:33;19426:102;19583:12;;;;;19572:23;;;;:8;;;:23;19639:1;19624:16;;;19609:31;;;19717:13;:11;:13::i;4608:7728::-;4651:12;4737:18;;:::i;:::-;-1:-1:-1;4915:15:108;;:18;;;;4805:4;5075:18;;;;5119;;;;5163;;;;;4805:4;;4895:17;;;;5075:18;5119;5253;;;5267:4;5253:18;5249:6777;;5303:2;5332:4;5327:9;;:14;5323:144;;5443:4;5438:9;;5430:4;:18;5424:24;5323:144;5488:2;:7;;5494:1;5488:7;5484:161;;5524:10;;;;;5556:16;;;;;;;;5524:10;-1:-1:-1;5484:161:108;;;5624:2;5619:7;;5484:161;5273:386;5249:6777;;;5761:10;:18;;5775:4;5761:18;5757:6269;;1745:10;5799:14;;5757:6269;;;5897:10;:18;;5911:4;5897:18;5893:6133;;5940:1;5935:6;;5893:6133;;;6065:10;:18;;6079:4;6065:18;6061:5965;;6118:4;6103:12;;;:19;6140:26;;;:14;;;:26;6191:13;:11;:13::i;:::-;6184:20;;;;;;;;;4608:7728;:::o;6061:5965::-;6330:10;:18;;6344:4;6330:18;6326:5700;;6481:14;;;6477:2708;6326:5700;6477:2708;6651:22;;;;;6647:2538;;6776:10;6789:27;6797:2;6802:10;6797:15;6814:1;6789:7;:27::i;:::-;6900:17;;;;6776:40;;-1:-1:-1;6900:17:108;6878:19;7050:14;7069:1;7044:26;7040:131;;7112:36;7136:11;1277:21:109;1426:15;;;1467:8;1461:4;1454:22;1595:4;1582:18;;1602:19;1578:44;1624:11;1575:61;;1222:430;7112:36:108;7098:50;;7040:131;7193:11;7224:6;;7257:20;;;;;7224:54;;;;;;;;3472:25:260;;;3545:10;3533:23;;;3513:18;;;3506:51;7193:11:108;;7224:6;;;:19;;3445:18:260;;7224:54:108;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;7192:86;;;;7505:1;7501:2;7497:10;7602:9;7599:1;7595:17;7684:6;7677:5;7674:17;7671:40;;;7704:5;7694:15;;7671:40;;7787:6;7783:2;7780:14;7777:34;;;7807:2;7797:12;;7777:34;7913:3;7908:1;7900:6;7896:14;7891:3;7887:24;7883:34;7876:41;;8013:3;8009:1;7997:9;7988:6;7985:1;7981:14;7977:30;7973:38;7969:48;7962:55;;8168:1;8164;8160;8148:9;8145:1;8141:17;8137:25;8133:33;8129:41;8295:1;8291;8287;8278:6;8266:9;8263:1;8259:17;8255:30;8251:38;8247:46;8243:54;8225:72;;8426:10;8422:15;8416:4;8412:26;8404:34;;8542:3;8534:4;8530:9;8525:3;8521:19;8518:28;8511:35;;;;8688:33;8697:2;8702:10;8697:15;8714:1;8717:3;8688:8;:33::i;:::-;8743:20;;;:38;;;;;;;;;-1:-1:-1;6647:2538:108;;-1:-1:-1;;;6647:2538:108;;8900:18;;;;;8896:289;;9070:2;9065:7;;6326:5700;;8896:289;9124:10;9119:15;;2053:3;9156:10;;8896:289;6326:5700;;;9314:10;:18;;9328:4;9314:18;9310:2716;;9468:15;;;1824:1;9468:15;;:34;;-1:-1:-1;9487:15:108;;;1859:1;9487:15;9468:34;:57;;;-1:-1:-1;9506:19:108;;;1936:1;9506:19;9468:57;9464:1593;;;9554:2;9549:7;;9310:2716;;9464:1593;9680:23;;;;;9676:1381;;9727:10;9740:27;9748:2;9753:10;9748:15;9765:1;9740:7;:27::i;:::-;9843:17;;;;9727:40;;-1:-1:-1;10086:1:108;10078:10;;10180:1;10176:17;10255:13;;;10252:32;;;10277:5;10271:11;;10252:32;10563:14;;;10369:1;10559:22;;;10555:32;;;;10452:26;10476:1;10361:10;;;10456:18;;;10452:26;10551:43;10357:20;;10659:12;10787:17;;;:23;10855:1;10832:20;;;:24;10365:2;-1:-1:-1;10365:2:108;6326:5700;;9310:2716;11259:10;:18;;11273:4;11259:18;11255:771;;11369:2;:7;;11375:1;11369:7;11365:647;;11462:14;;;;;:40;;-1:-1:-1;11480:22:108;;;1978:1;11480:22;11462:40;:62;;;-1:-1:-1;11506:18:108;;;1897:1;11506:18;11462:62;11458:404;;;11557:1;11552:6;;11365:647;;11458:404;11603:15;;;1824:1;11603:15;;:34;;-1:-1:-1;11622:15:108;;;1859:1;11622:15;11603:34;:61;;;-1:-1:-1;11641:23:108;;;2021:1;11641:23;11603:61;:84;;;-1:-1:-1;11668:19:108;;;1936:1;11668:19;11603:84;11599:263;;;11720:1;11715:6;;6326:5700;;11365:647;11913:10;11908:15;;2087:4;11945:11;;11365:647;12101:15;;;;;:23;;;;:18;;;;:23;;;;12138:15;;:23;;;:18;;;;:23;-1:-1:-1;12227:12:108;;;;12216:23;;;:8;;;:23;12283:1;12268:16;12253:31;;;;;12306:13;:11;:13::i;15089:2480::-;15183:12;15269:18;;:::i;:::-;-1:-1:-1;15337:4:108;15369:10;15477:13;;;15486:4;15477:13;15473:1705;;-1:-1:-1;15516:8:108;;;;15473:1705;;;15635:5;:13;;15644:4;15635:13;15631:1547;;15668:14;;;:8;;;:14;15631:1547;;;15798:5;:13;;15807:4;15798:13;15794:1384;;-1:-1:-1;15837:8:108;;;;15794:1384;;;15956:5;:13;;15965:4;15956:13;15952:1226;;15989:14;;;:8;;;:14;15952:1226;;;16130:5;:13;;16139:4;16130:13;16126:1052;;16257:9;16203:17;16183;;;16203;;;;16183:37;16264:2;16257:9;;;;;16239:8;;;:28;16285:22;:8;;;:22;16126:1052;;;16444:5;:13;;16453:4;16444:13;16440:738;;16511:11;16497;;;16511;;;16497:25;16566:2;16559:9;;;;;16541:8;;;:28;16587:22;:8;;;:22;16440:738;;;16768:5;:13;;16777:4;16768:13;16764:414;;16838:3;16819:23;;16825:3;16819:23;;;;;;;:::i;:::-;;16801:42;;:8;;;:42;16879:23;;;;;;;;;;;;;:::i;:::-;;16861:42;;:8;;;:42;16764:414;;;17072:5;:13;;17081:4;17072:13;17068:110;;17122:3;17116:9;;:3;:9;;;;;;;:::i;:::-;;17105:20;;;;:8;;;:20;17154:9;;;;;;;;;;;:::i;:::-;;17143:20;;:8;;;:20;17068:110;17271:14;;;;17267:85;;17334:3;17305:5;:15;;;17321:9;17305:26;;;;;;;;;:::i;:::-;:32;;;;:26;;;;;;:32;17267:85;17406:12;;;;;17395:23;;;;:8;;;:23;17462:1;17447:16;;;17432:31;;;17539:13;:11;:13::i;:::-;17532:20;15089:2480;-1:-1:-1;;;;;;;15089:2480:108:o;22918:1654::-;23094:14;23111:24;23123:11;23111;:24::i;:::-;23094:41;;23243:1;23236:5;23232:13;23229:33;;;23258:1;23255;23248:12;23229:33;23397:2;23591:15;;;23416:2;23405:14;;23393:10;23389:31;23386:1;23382:39;23547:16;;;23332:20;;23532:10;23521:22;;;23517:27;23507:38;23504:60;24033:5;24030:1;24026:13;24104:1;24089:343;24114:2;24111:1;24108:9;24089:343;;;24237:2;24225:15;;;24174:20;24272:12;;;24286:1;24268:20;24309:42;;;;24377:1;24372:42;;;;24261:153;;24309:42;21396:1;21389:12;;;21429:2;21422:13;;;21474:2;21461:16;;24318:31;;24309:42;;24372;21396:1;21389:12;;;21429:2;21422:13;;;21474:2;21461:16;;24381:31;;24261:153;-1:-1:-1;;24132:1:108;24125:9;24089:343;;;-1:-1:-1;;24531:4:108;24524:18;-1:-1:-1;;;;22918:1654:108:o;19951:586::-;20273:20;;;20297:7;20273:32;20266:3;:40;;;20379:14;;20434:17;;20428:24;;;20420:72;;;;;;;4209:2:260;20420:72:108;;;4191:21:260;4248:2;4228:18;;;4221:30;4287:34;4267:18;;;4260:62;4358:5;4338:18;;;4331:33;4381:19;;20420:72:108;4007:399:260;20420:72:108;20506:14;19951:586;;;:::o;-1:-1:-1:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o;467:347:260:-;518:8;528:6;582:3;575:4;567:6;563:17;559:27;549:55;;600:1;597;590:12;549:55;-1:-1:-1;623:20:260;;666:18;655:30;;652:50;;;698:1;695;688:12;652:50;735:4;727:6;723:17;711:29;;787:3;780:4;771:6;763;759:19;755:30;752:39;749:59;;;804:1;801;794:12;749:59;467:347;;;;;:::o;819:717::-;909:6;917;925;933;986:2;974:9;965:7;961:23;957:32;954:52;;;1002:1;999;992:12;954:52;1042:9;1029:23;1071:18;1112:2;1104:6;1101:14;1098:34;;;1128:1;1125;1118:12;1098:34;1167:58;1217:7;1208:6;1197:9;1193:22;1167:58;:::i;:::-;1244:8;;-1:-1:-1;1141:84:260;-1:-1:-1;1332:2:260;1317:18;;1304:32;;-1:-1:-1;1348:16:260;;;1345:36;;;1377:1;1374;1367:12;1345:36;;1416:60;1468:7;1457:8;1446:9;1442:24;1416:60;:::i;:::-;819:717;;;;-1:-1:-1;1495:8:260;-1:-1:-1;;;;819:717:260:o;1723:184::-;1775:77;1772:1;1765:88;1872:4;1869:1;1862:15;1896:4;1893:1;1886:15;3568:245;3647:6;3655;3708:2;3696:9;3687:7;3683:23;3679:32;3676:52;;;3724:1;3721;3714:12;3676:52;-1:-1:-1;;3747:16:260;;3803:2;3788:18;;;3782:25;3747:16;;3782:25;;-1:-1:-1;3568:245:260:o;3818:184::-;3870:77;3867:1;3860:88;3967:4;3964:1;3957:15;3991:4;3988:1;3981:15"

func init() {
	if err := json.Unmarshal([]byte(MIPSStorageLayoutJSON), MIPSStorageLayout); err != nil {
		panic(err)
	}

	layouts["MIPS"] = MIPSStorageLayout
	deployedBytecodes["MIPS"] = MIPSDeployedBin
}
