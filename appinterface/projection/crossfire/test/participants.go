package crossfire_test

// 3 participants for unit test
const THREE_PARTICIPANTS_SAMPLE_JSON = `[
  {
	"operatorAddress": "tcrocncl14m5a4kxt2e82uqqs5gtqza29dm5wqzyalddug5",
	"primaryAddress": "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
	"moniker": "Bambarello"
  },
  {
	"operatorAddress": "tcrocncl1n4t5q77kn9vf73s7ljs96m85jgg49yqpg0chrj",
	"primaryAddress": "tcro1f6qcvp33dc79xzpuwll7mln5lnepuqv8d7led9",
	"moniker": "eric-node"
  },
  {
	"operatorAddress": "tcrocncl1f6qcvp33dc79xzpuwll7mln5lnepuqv8cpuq4x",
	"primaryAddress": "tcro1432x4lc5mrgm30c9xx35unmn9ultemm5nt40vq",
	"moniker": "erasde"
  }
]`

// sample participants with corresponding tendermint address
// 1. tcrocncl1n4t5q77kn9vf73s7ljs96m85jgg49yqpg0chrj: 703B26AEA0867B03572719D22F4B8E6D93CA838C
// 2. tcrocncl15xr8daqzpu0wf8t6hx95zlxmqwzmf4ea5gja60: AEA0F558C9616A7089791D1AE4C08DC5F69A0A0B
// 3. tcrocncl197ujxhaeyyv309f39c0s2gn0af0pps5pcxsr0a: 0FB7AE9AC2E3F148CA130341B6CD4DB3682E2D54
const COMMIT_RANK_PARTICIPANTS_SAMPLE_JSON = `[
  {
	"operatorAddress": "tcrocncl1n4t5q77kn9vf73s7ljs96m85jgg49yqpg0chrj",
	"primaryAddress": "tcro1n4t5q77kn9vf73s7ljs96m85jgg49yqpasmwm3",
	"moniker": "node1"
  },
  {
	"operatorAddress": "tcrocncl15xr8daqzpu0wf8t6hx95zlxmqwzmf4ea5gja60",
	"primaryAddress": "tcro15xr8daqzpu0wf8t6hx95zlxmqwzmf4eaph3yzv",
	"moniker": "node2"
  },
  {
	"operatorAddress": "tcrocncl197ujxhaeyyv309f39c0s2gn0af0pps5pcxsr0a",
	"primaryAddress": "tcro197ujxhaeyyv309f39c0s2gn0af0pps5pden6h7",
	"moniker": "node3"
  }
]`

// full participants sample
const PARTICIPANTS_SAMPLE_JSON = `[
  {
    "operatorAddress": "tcrocncl197ujxhaeyyv309f39c0s2gn0af0pps5pcxsr0a",
    "primaryAddress": "tcro197ujxhaeyyv309f39c0s2gn0af0pps5pden6h7",
    "moniker": "node1"
  },
  {
    "operatorAddress": "tcrocncl1f6qcvp33dc79xzpuwll7mln5lnepuqv8cpuq4x",
    "primaryAddress": "tcro1f6qcvp33dc79xzpuwll7mln5lnepuqv8d7led9",
    "moniker": "nebkas.ro"
  },
  {
    "operatorAddress": "tcrocncl18ylchgmxyphw3ctsl75n53ujequkmmag2n6x3f",
    "primaryAddress": "tcro18ylchgmxyphw3ctsl75n53ujequkmmaglvelf2",
    "moniker": "Spanish_Team_OLDMAN"
  },
  {
    "operatorAddress": "tcrocncl1n4t5q77kn9vf73s7ljs96m85jgg49yqpg0chrj",
    "primaryAddress": "tcro1n4t5q77kn9vf73s7ljs96m85jgg49yqpasmwm3",
    "moniker": "node0"
  },
  {
    "operatorAddress": "tcrocncl15xr8daqzpu0wf8t6hx95zlxmqwzmf4ea5gja60",
    "primaryAddress": "tcro15xr8daqzpu0wf8t6hx95zlxmqwzmf4eaph3yzv",
    "moniker": "node2"
  },
  {
    "operatorAddress": "tcrocncl1war4r2rvy0nhyxhel8du6mrsly0nzgqd8r9wvt",
    "primaryAddress": "tcro1war4r2rvy0nhyxhel8du6mrsly0nzgqdjuxh5g",
    "moniker": "hello123-node"
  },
  {
    "operatorAddress": "tcrocncl1xmf3e4ua5thfesem8tyjdx38rgk6ukdr04696r",
    "primaryAddress": "tcro1xmf3e4ua5thfesem8tyjdx38rgk6ukdr62euzq",
    "moniker": "RoHoss"
  },
  {
    "operatorAddress": "tcrocncl14fnzv5g92s6f8dg534lccp4x5tylkvthtamehl",
    "primaryAddress": "tcro14fnzv5g92s6f8dg534lccp4x5tylkvth7zcq0u",
    "moniker": "crypto.bzh"
  },
  {
    "operatorAddress": "tcrocncl19s6xmujtpv408uwvm7f76mu56elax7m45ufrn6",
    "primaryAddress": "tcro19s6xmujtpv408uwvm7f76mu56elax7m4pr26te",
    "moniker": "EvLeoNode"
  },
  {
    "operatorAddress": "tcrocncl1sedlpyy3ff7y943gdq02v7rffdf0u8ejm3e5u4",
    "primaryAddress": "tcro1sedlpyy3ff7y943gdq02v7rffdf0u8ejww6dyk",
    "moniker": "cdcpde"
  },
  {
    "operatorAddress": "tcrocncl1sunjpn9yy930sf9p5gmgygft4gfa4gdqlry097",
    "primaryAddress": "tcro1sunjpn9yy930sf9p5gmgygft4gfa4gdq2u8kaa",
    "moniker": "sheepmaan"
  },
  {
    "operatorAddress": "tcrocncl13m8uf6409w7euqle2hw3rrx07txh5pzs9ewwsa",
    "primaryAddress": "tcro13m8uf6409w7euqle2hw3rrx07txh5pzssxdhg7",
    "moniker": "Pilotcoin-Node2"
  },
  {
    "operatorAddress": "tcrocncl1qnc9xflghhg4eh76yu789hxuflp2pskg758cjn",
    "primaryAddress": "tcro1qnc9xflghhg4eh76yu789hxuflp2pskgttyp2s",
    "moniker": "Primary Node-RR"
  },
  {
    "operatorAddress": "tcrocncl1mrhp0spxjlh8hdx77z70snkp6452fpxyqjdgj6",
    "primaryAddress": "tcro1mrhp0spxjlh8hdx77z70snkp6452fpxy4dw32e",
    "moniker": "theBro3"
  },
  {
    "operatorAddress": "tcrocncl1m80c9cfmfxx4gcnza4q0tycc3etm0crmdclmjz",
    "primaryAddress": "tcro1m80c9cfmfxx4gcnza4q0tycc3etm0crmc8uz2p",
    "moniker": "Jasper"
  },
  {
    "operatorAddress": "tcrocncl1mnz68z72pwvndpdyy9rpqet8xr8wtmhvv2tnkr",
    "primaryAddress": "tcro1mnz68z72pwvndpdyy9rpqet8xr8wtmhve4g2wq",
    "moniker": "Jayjay"
  },
  {
    "operatorAddress": "tcrocncl1yr0p3t3uz9fxjzgle3m69u52gkmqmg6n5w2ssq",
    "primaryAddress": "tcro1yr0p3t3uz9fxjzgle3m69u52gkmqmg6np3ffgr",
    "moniker": "wilevalidator2"
  },
  {
    "operatorAddress": "tcrocncl14tfhw845xpdgpvhc6f7rlsmshzgj3tr4rrl4gu",
    "primaryAddress": "tcro14tfhw845xpdgpvhc6f7rlsmshzgj3tr4kuuvsl",
    "moniker": "GreenTara"
  },
  {
    "operatorAddress": "tcrocncl1zjd2ewqn055ggs889yg00cxs2npyfd6glpjp0a",
    "primaryAddress": "tcro1zjd2ewqn055ggs889yg00cxs2npyfd6g273ch7",
    "moniker": "dhjensen-test"
  },
  {
    "operatorAddress": "tcrocncl1yrju459530e854883w0vcfq838ea8dgmtkufm0",
    "primaryAddress": "tcro1yrju459530e854883w0vcfq838ea8dgm7flsrv",
    "moniker": "fethiye-node"
  },
  {
    "operatorAddress": "tcrocncl1l82wd3ujk2v7yhd66fdzmnhhdekeawp8fyzt0f",
    "primaryAddress": "tcro1l82wd3ujk2v7yhd66fdzmnhhdekeawp8umpjh2",
    "moniker": "villagelife-1"
  },
  {
    "operatorAddress": "tcrocncl12pfdjkey50vwwv0lnn4ttgupfld9c387prva44",
    "primaryAddress": "tcro12pfdjkey50vwwv0lnn4ttgupfld9c3875u0ydk",
    "moniker": "zdeadex"
  },
  {
    "operatorAddress": "tcrocncl1ld2e06tpzkv0gudtd3fe8xdl6wxql29g40ytrc",
    "primaryAddress": "tcro1ld2e06tpzkv0gudtd3fe8xdl6wxql29gqs8jmm",
    "moniker": "DavletunerValidatorTest"
  },
  {
    "operatorAddress": "tcrocncl1z8al8uxaykh5ytaj8v2g3ham0kavmk6xvxnyyn",
    "primaryAddress": "tcro1z8al8uxaykh5ytaj8v2g3ham0kavmk6xeesaus",
    "moniker": "anna-node"
  },
  {
    "operatorAddress": "tcrocncl1t9lu27kdmyee82knq6pqy2xeqkc59cryxeeghm",
    "primaryAddress": "tcro1t9lu27kdmyee82knq6pqy2xeqkc59crynx630c",
    "moniker": "SebastianSVK"
  },
  {
    "operatorAddress": "tcrocncl1n0t4pxyxusg5a6uz5xnwmhcjfludy7zfjhfgg0",
    "primaryAddress": "tcro1n0t4pxyxusg5a6uz5xnwmhcjfludy7zf8g23sv",
    "moniker": "xeset-node"
  },
  {
    "operatorAddress": "tcrocncl1x344af78plh9r52xr8l2r74eyx0a8z2y3xfl2d",
    "primaryAddress": "tcro1x344af78plh9r52xr8l2r74eyx0a8z2yye2xjw",
    "moniker": "Vision"
  },
  {
    "operatorAddress": "tcrocncl1reyshfdygf7673xm9p8v0xvtd96m6cd6canhu3",
    "primaryAddress": "tcro1reyshfdygf7673xm9p8v0xvtd96m6cd6dzswyj",
    "moniker": "Spanish_Team_JORDI"
  },
  {
    "operatorAddress": "tcrocncl17ftaz7ac29gtv6nqycw2042ygm05u88z59wvhc",
    "primaryAddress": "tcro17ftaz7ac29gtv6nqycw2042ygm05u88zp6d40m",
    "moniker": "MallorcaChain"
  },
  {
    "operatorAddress": "tcrocncl1ly70gar6tpxgh4ntn6ryfhmeeper48w0w06jm7",
    "primaryAddress": "tcro1ly70gar6tpxgh4ntn6ryfhmeeper48w0msetra",
    "moniker": "wormchester-cro-node"
  },
  {
    "operatorAddress": "tcrocncl16ce44ey8z3t7r9wc05zp95ug7cap6pf5g992zx",
    "primaryAddress": "tcro16ce44ey8z3t7r9wc05zp95ug7cap6pf5a6xn69",
    "moniker": "Nazar"
  },
  {
    "operatorAddress": "tcrocncl1zg8d42mwc4fzspaakcuygfjhhtjenxxyacrsd0",
    "primaryAddress": "tcro1zg8d42mwc4fzspaakcuygfjhhtjenxxyg8qf4v",
    "moniker": "wildoperator"
  },
  {
    "operatorAddress": "tcrocncl1sruzd529lhjju6hfcwd2fxp3v0e7p0vqqtme76",
    "primaryAddress": "tcro1sruzd529lhjju6hfcwd2fxp3v0e7p0vq45cqxe",
    "moniker": "Spanish_Team_ATOLAMO"
  },
  {
    "operatorAddress": "tcrocncl1rdrdg2n2hwmjscajf99fkv7sancexfnj8m9z0a",
    "primaryAddress": "tcro1rdrdg2n2hwmjscajf99fkv7sancexfnjjyxmh7",
    "moniker": "xlr8tor"
  },
  {
    "operatorAddress": "tcrocncl1mzug729mjl9f0kqv3p8f39tmw2wgkls09tg25k",
    "primaryAddress": "tcro1mzug729mjl9f0kqv3p8f39tmw2wgkls0s5tnv4",
    "moniker": "rainbow"
  },
  {
    "operatorAddress": "tcrocncl1uyre8mjsvclg9z4syqshawumwucmm2hermtjrx",
    "primaryAddress": "tcro1uyre8mjsvclg9z4syqshawumwucmm2hekygtm9",
    "moniker": "santoshi-node"
  },
  {
    "operatorAddress": "tcrocncl1f2nkkr8h4k0sfqwj0geyk285s7smqlpayfnukw",
    "primaryAddress": "tcro1f2nkkr8h4k0sfqwj0geyk285s7smqlpa3ks9wd",
    "moniker": "JM-node"
  },
  {
    "operatorAddress": "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek",
    "primaryAddress": "tcro15grftg88l0gdw4mg9t9pwnl0pde2asjzvfpkp4",
    "moniker": "Spanish_Team_IDARRIB"
  },
  {
    "operatorAddress": "tcrocncl1tytfek05p73yylhap2gy2k5svlhrnke5s73kcl",
    "primaryAddress": "tcro1tytfek05p73yylhap2gy2k5svlhrnke59pj0qu",
    "moniker": "bananapi"
  },
  {
    "operatorAddress": "tcrocncl1pf6u7cu4ukt6e7hhpg98k99g0c4ea32zwx0c09",
    "primaryAddress": "tcro1pf6u7cu4ukt6e7hhpg98k99g0c4ea32zmevphx",
    "moniker": "abel-croeseid"
  },
  {
    "operatorAddress": "tcrocncl13vh39gf8qa0mjzwpmttpwqlzmz2hdq02kf52lr",
    "primaryAddress": "tcro13vh39gf8qa0mjzwpmttpwqlzmz2hdq02rkhn8q",
    "moniker": "seashells"
  },
  {
    "operatorAddress": "tcrocncl17mw82r342a4mgg6ppyd55sztrvyhpmvn2mv43e",
    "primaryAddress": "tcro17mw82r342a4mgg6ppyd55sztrvyhpmvnly0vf6",
    "moniker": "cromagnon"
  },
  {
    "operatorAddress": "tcrocncl1xwd3k8xterdeft3nxqg92szhpz6vx43qspdpw6",
    "primaryAddress": "tcro1xwd3k8xterdeft3nxqg92szhpz6vx43q97wcke",
    "moniker": "Bambarello"
  },
  {
    "operatorAddress": "tcrocncl1432x4lc5mrgm30c9xx35unmn9ultemm5x5kk5r",
    "primaryAddress": "tcro1432x4lc5mrgm30c9xx35unmn9ultemm5nt40vq",
    "moniker": "eric-node"
  },
  {
    "operatorAddress": "tcrocncl1urmrrmmt6gdf077dmgt95cmj6tc0z904pjhlrd",
    "primaryAddress": "tcro1urmrrmmt6gdf077dmgt95cmj6tc0z9045d5xmw",
    "moniker": "1NP_Raul"
  },
  {
    "operatorAddress": "tcrocncl1gs60vktclkdf004regmezdqxu29knm8u8k7hfq",
    "primaryAddress": "tcro1gs60vktclkdf004regmezdqxu29knm8ujfaw3r",
    "moniker": "huglester"
  },
  {
    "operatorAddress": "tcrocncl1arz0nncekfdsj6e7jhy3pcjthze2wpchx6a2vs",
    "primaryAddress": "tcro1arz0nncekfdsj6e7jhy3pcjthze2wpchn97n5n",
    "moniker": "mas_node"
  },
  {
    "operatorAddress": "tcrocncl1mz5rdtf9wufwkh8te2zww7twtmna6rhllluw8m",
    "primaryAddress": "tcro1mz5rdtf9wufwkh8te2zww7twtmna6rhl2qlhlc",
    "moniker": "arg-node"
  },
  {
    "operatorAddress": "tcrocncl1wypk0unhg9432kdz6hmumqqjd0lz83p3mc42ty",
    "primaryAddress": "tcro1wypk0unhg9432kdz6hmumqqjd0lz83p3w8knn8",
    "moniker": "MakingCash"
  },
  {
    "operatorAddress": "tcrocncl1tncpcuhh7cedmp9jc6qap4vqm5kxulfa5xg5z8",
    "primaryAddress": "tcro1tncpcuhh7cedmp9jc6qap4vqm5kxulfapetd6y",
    "moniker": "sagepoint-node"
  },
  {
    "operatorAddress": "tcrocncl1698gtl69qaw688uewtgahjvd0pcft6xj532c9r",
    "primaryAddress": "tcro1698gtl69qaw688uewtgahjvd0pcft6xjpwfpaq",
    "moniker": "Perfect Stake"
  },
  {
    "operatorAddress": "tcrocncl1434qa8s7vcw9jgln8lxzf8dvxedattumv33vgh",
    "primaryAddress": "tcro1434qa8s7vcw9jgln8lxzf8dvxedattumewj4s5",
    "moniker": "CroBaba"
  },
  {
    "operatorAddress": "tcrocncl1fgae85rgzv57kd23hkux4ktjqtscj75k4ry56e",
    "primaryAddress": "tcro1fgae85rgzv57kd23hkux4ktjqtscj75kqu8dz6",
    "moniker": "WS-Totti Validator"
  },
  {
    "operatorAddress": "tcrocncl18s4fkkfxw9parlgd2v554qcxuac8fkta0yjx4c",
    "primaryAddress": "tcro18s4fkkfxw9parlgd2v554qcxuac8fkta6m3ldm",
    "moniker": "ilCompare"
  },
  {
    "operatorAddress": "tcrocncl1d2fvfvnff9pj3zv9mnlnv50vuqg0d3vczmkk2s",
    "primaryAddress": "tcro1d2fvfvnff9pj3zv9mnlnv50vuqg0d3vchy40jn",
    "moniker": "santoae-node"
  },
  {
    "operatorAddress": "tcrocncl1fs8r6zxmr5nc86j8cpcmjmccf8s2cafxzt5alq",
    "primaryAddress": "tcro1fs8r6zxmr5nc86j8cpcmjmccf8s2cafxh5hy8r",
    "moniker": "GRom81"
  },
  {
    "operatorAddress": "tcrocncl1923pz03mhjaztgcv3gey0hj0amwx02dyskau52",
    "primaryAddress": "tcro1923pz03mhjaztgcv3gey0hj0amwx02dy9f79vf",
    "moniker": "Code-Breader"
  },
  {
    "operatorAddress": "tcrocncl172v9aga6k5nlrw6uc387egzls08nhl4cyzncmn",
    "primaryAddress": "tcro172v9aga6k5nlrw6uc387egzls08nhl4c3asprs",
    "moniker": "StakeService"
  },
  {
    "operatorAddress": "tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu",
    "primaryAddress": "tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel",
    "moniker": "NP_SIRJN"
  },
  {
    "operatorAddress": "tcrocncl1descn8h7kj52en8gn9j9dqwyy495mnxz0nu6fk",
    "primaryAddress": "tcro1descn8h7kj52en8gn9j9dqwyy495mnxz6vlr34",
    "moniker": "Tazman Validator"
  },
  {
    "operatorAddress": "tcrocncl1223xh0cgfy3tak22f8wf6zt0kj7wph638zcguy",
    "primaryAddress": "tcro1223xh0cgfy3tak22f8wf6zt0kj7wph63jam3y8",
    "moniker": "captainStart"
  },
  {
    "operatorAddress": "tcrocncl1pe3lqu8w9laa2euce7zudeu44w49m43ed3vs4t",
    "primaryAddress": "tcro1pe3lqu8w9laa2euce7zudeu44w49m43ecw0fdg",
    "moniker": "safuguard"
  },
  {
    "operatorAddress": "tcrocncl12mamjtrddwldc0n9474cx900ey7vzn5qsa03w0",
    "primaryAddress": "tcro12mamjtrddwldc0n9474cx900ey7vzn5q9zvgkv",
    "moniker": "node-corenting"
  },
  {
    "operatorAddress": "tcrocncl14qc4fecpt7updwlg4v0cpntetuy7w8wvkq2pez",
    "primaryAddress": "tcro14qc4fecpt7updwlg4v0cpntetuy7w8wvrlfcpp",
    "moniker": "otis"
  },
  {
    "operatorAddress": "tcrocncl1u7lj7vhxncelcf90dwz3sf5pnc07jk9sdavypr",
    "primaryAddress": "tcro1u7lj7vhxncelcf90dwz3sf5pnc07jk9scz0aeq",
    "moniker": "Jorgeminator"
  },
  {
    "operatorAddress": "tcrocncl1cdaml3s69j6u4nawlvd7jffhsuzqjy8j7mr3v3",
    "primaryAddress": "tcro1cdaml3s69j6u4nawlvd7jffhsuzqjy8jtyqg5j",
    "moniker": "isillien"
  },
  {
    "operatorAddress": "tcrocncl1ncagf3j5d2mz9m7f5jz7zymynlm6gc8mxy678z",
    "primaryAddress": "tcro1ncagf3j5d2mz9m7f5jz7zymynlm6gc8mnme8lp",
    "moniker": "thundercrypto-node"
  },
  {
    "operatorAddress": "tcrocncl1phw7nqk08kacpw76rg6d4m33dx7kkmcel9uwkt",
    "primaryAddress": "tcro1phw7nqk08kacpw76rg6d4m33dx7kkmce26lhwg",
    "moniker": "bbpool"
  },
  {
    "operatorAddress": "tcrocncl18ekp7lpd2e0t9duaedv3x94xssrp5uw386qyms",
    "primaryAddress": "tcro18ekp7lpd2e0t9duaedv3x94xssrp5uw3j9rarn",
    "moniker": "twsonaws-one"
  },
  {
    "operatorAddress": "tcrocncl1zf6clmtf9lsxz748rcqlsuvs5tcjmv50a6cw7r",
    "primaryAddress": "tcro1zf6clmtf9lsxz748rcqlsuvs5tcjmv50g9mhxq",
    "moniker": "monikerliudas"
  },
  {
    "operatorAddress": "tcrocncl13tt3jw6klqza3xd92t89h82r79de0f89tqvlk5",
    "primaryAddress": "tcro13tt3jw6klqza3xd92t89h82r79de0f897l0xwh",
    "moniker": "falcon-node"
  },
  {
    "operatorAddress": "tcrocncl1qvcw58u9vqnc3eymqfgs9ztdcmzxj8qu3rjjpj",
    "primaryAddress": "tcro1qvcw58u9vqnc3eymqfgs9ztdcmzxj8quyu3te3",
    "moniker": "vipnamai"
  },
  {
    "operatorAddress": "tcrocncl1t6u0xmthzh753lwqxpuzmadykv0g2hmx7d796c",
    "primaryAddress": "tcro1t6u0xmthzh753lwqxpuzmadykv0g2hmxtjauzm",
    "moniker": "sopi"
  },
  {
    "operatorAddress": "tcrocncl1avac3ekwhhls6apz5naqm88h5gxz5uavjwfa4k",
    "primaryAddress": "tcro1avac3ekwhhls6apz5naqm88h5gxz5uav832yd4",
    "moniker": "temurblocks"
  },
  {
    "operatorAddress": "tcrocncl1046hpqlc2aakwg9qfzn000fp2n4rzfwxvk03du",
    "primaryAddress": "tcro1046hpqlc2aakwg9qfzn000fp2n4rzfwxefvg4l",
    "moniker": "DarkStar"
  },
  {
    "operatorAddress": "tcrocncl1cuq2jhdhghuxwpf9t2d03vlcemm4nfv08r4qgl",
    "primaryAddress": "tcro1cuq2jhdhghuxwpf9t2d03vlcemm4nfv0jukesu",
    "moniker": "Wanderer"
  },
  {
    "operatorAddress": "tcrocncl13afx0q4w9j5s93skdewq79fgnp7rqswl0phf40",
    "primaryAddress": "tcro13afx0q4w9j5s93skdewq79fgnp7rqswl675sdv",
    "moniker": "ErikFin-node"
  },
  {
    "operatorAddress": "tcrocncl1uevms2nv4f2dhvm5u7sgus2yncgh7gdwx9l6k6",
    "primaryAddress": "tcro1uevms2nv4f2dhvm5u7sgus2yncgh7gdwn6urwe",
    "moniker": "Julia"
  },
  {
    "operatorAddress": "tcrocncl1ttw599wan2jvwshczsverpjfpwtt75cpdvhz0x",
    "primaryAddress": "tcro1ttw599wan2jvwshczsverpjfpwtt75cpcn5mh9",
    "moniker": "CryptoHank"
  },
  {
    "operatorAddress": "tcrocncl1upjwcer0hv00vs8qszj954x75znvm46nyvd5pw",
    "primaryAddress": "tcro1upjwcer0hv00vs8qszj954x75znvm46n3nwded",
    "moniker": "Primary Node"
  },
  {
    "operatorAddress": "tcrocncl1h583nq3w5x26datkhdmt7y6pc50f74yytcn0mq",
    "primaryAddress": "tcro1h583nq3w5x26datkhdmt7y6pc50f74yy78skrr",
    "moniker": "aurelius"
  },
  {
    "operatorAddress": "tcrocncl1l388chd86uwuf0xjdz8n7489whvv64tcez2nvj",
    "primaryAddress": "tcro1l388chd86uwuf0xjdz8n7489whvv64tcvaf253",
    "moniker": "Seb"
  },
  {
    "operatorAddress": "tcrocncl1whw6cvv044q2y2eam8kp4hgnhgugr0amvt0eu3",
    "primaryAddress": "tcro1whw6cvv044q2y2eam8kp4hgnhgugr0ame5vqyj",
    "moniker": "Terminator"
  },
  {
    "operatorAddress": "tcrocncl1d8lkhl8a9hchjk98t4my2c88srq6f6h2khjwr6",
    "primaryAddress": "tcro1d8lkhl8a9hchjk98t4my2c88srq6f6h2rg3hme",
    "moniker": "siddha-node"
  },
  {
    "operatorAddress": "tcrocncl1fv49fg30nm37gsk058vt4gmjy4dqm9tvh9gad0",
    "primaryAddress": "tcro1fv49fg30nm37gsk058vt4gmjy4dqm9tvz6ty4v",
    "moniker": "tdtest-node"
  },
  {
    "operatorAddress": "tcrocncl14aekczvm4gedj8au288h739eqp4mx7t2dnves5",
    "primaryAddress": "tcro14aekczvm4gedj8au288h739eqp4mx7t2cv0qgh",
    "moniker": "koshmanakov"
  },
  {
    "operatorAddress": "tcrocncl17wggxq5d5fqp2xz03rgxuyw2al3k552hy8wfvr",
    "primaryAddress": "tcro17wggxq5d5fqp2xz03rgxuyw2al3k552h3cds5q",
    "moniker": "jokerD"
  },
  {
    "operatorAddress": "tcrocncl1kvv44a8ccrashw5taeprx9pzwlwp8z50333028",
    "primaryAddress": "tcro1kvv44a8ccrashw5taeprx9pzwlwp8z50ywjkjy",
    "moniker": "demonical-node"
  },
  {
    "operatorAddress": "tcrocncl1387sysfz49qa6ywevr9zdf4myzl8y64zf0hyxs",
    "primaryAddress": "tcro1387sysfz49qa6ywevr9zdf4myzl8y64zus5a7n",
    "moniker": "laptop-node"
  },
  {
    "operatorAddress": "tcrocncl1a9usdrhtx4uz9d7luupwfmeuuqzw5r7x30j8q6",
    "primaryAddress": "tcro1a9usdrhtx4uz9d7luupwfmeuuqzw5r7xys37ce",
    "moniker": "sales1man"
  },
  {
    "operatorAddress": "tcrocncl1xhvzy9083jh8mn6k49ckuhv5r9m7jr9w7hnkdk",
    "primaryAddress": "tcro1xhvzy9083jh8mn6k49ckuhv5r9m7jr9wtgs044",
    "moniker": "OptimusValidator"
  },
  {
    "operatorAddress": "tcrocncl1gqkl50u3hy6yey46yrexcxphfnf00cqtmxmety",
    "primaryAddress": "tcro1gqkl50u3hy6yey46yrexcxphfnf00cqtwecqn8",
    "moniker": "interPolator"
  },
  {
    "operatorAddress": "tcrocncl1nj3vx6yjysgvzxcrsycnafnhhs4ptjjwg9qlxq",
    "primaryAddress": "tcro1nj3vx6yjysgvzxcrsycnafnhhs4ptjjwa6rx7r",
    "moniker": "keti"
  },
  {
    "operatorAddress": "tcrocncl1zmspcr80mrng8flddv7luknuhgvjaf4j7ae0xl",
    "primaryAddress": "tcro1zmspcr80mrng8flddv7luknuhgvjaf4jtz6k7u",
    "moniker": "GreenMile"
  },
  {
    "operatorAddress": "tcrocncl13plny9ushrdlq4pxglmlmyrca85lcn3zkv4rv2",
    "primaryAddress": "tcro13plny9ushrdlq4pxglmlmyrca85lcn3zrnk65f",
    "moniker": "sosocovidoloco"
  },
  {
    "operatorAddress": "tcrocncl14zqn5m6q2exlm29fh8j5rsacl86j4mqpaa3lyx",
    "primaryAddress": "tcro14zqn5m6q2exlm29fh8j5rsacl86j4mqpgzjxu9",
    "moniker": "Surg"
  },
  {
    "operatorAddress": "tcrocncl1tpj45q5569y69fntmrgetxd7ptvfs6s5krchqr",
    "primaryAddress": "tcro1tpj45q5569y69fntmrgetxd7ptvfs6s5rumwcq",
    "moniker": "quacc-node"
  },
  {
    "operatorAddress": "tcrocncl1cjsg4n242qwc0dj9tnfnv7yfh85sk3h9ufalcq",
    "primaryAddress": "tcro1cjsg4n242qwc0dj9tnfnv7yfh85sk3h9fk7xqr",
    "moniker": "Maramicky"
  },
  {
    "operatorAddress": "tcrocncl1qrxlpmych4ayllsqlac3dk8ut3rce3u83akf3q",
    "primaryAddress": "tcro1qrxlpmych4ayllsqlac3dk8ut3rce3u8yz4sfr",
    "moniker": "dh"
  },
  {
    "operatorAddress": "tcrocncl1qjw67mxlqzvzvhhwe04rwunpyj2wcka8h83gvq",
    "primaryAddress": "tcro1qjw67mxlqzvzvhhwe04rwunpyj2wcka8zcj35r",
    "moniker": "eric-node2"
  },
  {
    "operatorAddress": "tcrocncl19vwrgyzczl765k3ujjycua6zl76whu3w6rgr9l",
    "primaryAddress": "tcro19vwrgyzczl765k3ujjycua6zl76whu3w0ut6au",
    "moniker": "STRM"
  },
  {
    "operatorAddress": "tcrocncl1nz6jujxv2px3nq5gvw9cfguumwtwhsmwvrkc7e",
    "primaryAddress": "tcro1nz6jujxv2px3nq5gvw9cfguumwtwhsmweu4px6",
    "moniker": "Fenix"
  },
  {
    "operatorAddress": "tcrocncl1acred059mrqg50hw7aqmdv5m5l27kc4g2t8jp4",
    "primaryAddress": "tcro1acred059mrqg50hw7aqmdv5m5l27kc4gl5ytek",
    "moniker": "Vanlee"
  },
  {
    "operatorAddress": "tcrocncl1xqqg3z70gz0lcyk9h8s0cm9w9gx4r7lgeega9d",
    "primaryAddress": "tcro1xqqg3z70gz0lcyk9h8s0cm9w9gx4r7lgvxtyaw",
    "moniker": "team-thrive-node"
  },
  {
    "operatorAddress": "tcrocncl1dpadgw8rfaxrevt07r7azkmen8fymvmdxnfek2",
    "primaryAddress": "tcro1dpadgw8rfaxrevt07r7azkmen8fymvmdnv2qwf",
    "moniker": "Primary Node - Primes"
  },
  {
    "operatorAddress": "tcrocncl1plfra5vy3lj0ar37hm04p0u9n6teh7epr03vlx",
    "primaryAddress": "tcro1plfra5vy3lj0ar37hm04p0u9n6teh7epksj489",
    "moniker": "Primary Node - TCoSM"
  },
  {
    "operatorAddress": "tcrocncl146cnm6f2fa2gehg5zwnf7fdztctl0fv5840q3w",
    "primaryAddress": "tcro146cnm6f2fa2gehg5zwnf7fdztctl0fv5j2vefd",
    "moniker": "Primary Node - Countries"
  },
  {
    "operatorAddress": "tcrocncl152yy0jyjptnrvkyn4tcgpv4yguq4xj0xhqmzm0",
    "primaryAddress": "tcro152yy0jyjptnrvkyn4tcgpv4yguq4xj0xzlcmrv",
    "moniker": "Primary Node - Comics"
  },
  {
    "operatorAddress": "tcrocncl1rpl6tnu986ph3284k0qdkpneuggdfy2v6xk8c2",
    "primaryAddress": "tcro1rpl6tnu986ph3284k0qdkpneuggdfy2v0e47qf",
    "moniker": "Primary Node - Stars"
  },
  {
    "operatorAddress": "tcrocncl1a3xv6a2kl0rtxuc56v7hlt0eruqe0lyq9v7qdf",
    "primaryAddress": "tcro1a3xv6a2kl0rtxuc56v7hlt0eruqe0lyqsnae42",
    "moniker": "Primary Node - Knock Knock"
  },
  {
    "operatorAddress": "tcrocncl1tqaqlt49uqkwx0zrhcudmtsj57uhcxmrgklcqk",
    "primaryAddress": "tcro1tqaqlt49uqkwx0zrhcudmtsj57uhcxmrafupc4",
    "moniker": "Primary Node - Dieties"
  },
  {
    "operatorAddress": "tcrocncl1tt2yhw293zwqazp94d07n33p7wqjf0g9e52gwa",
    "primaryAddress": "tcro1tt2yhw293zwqazp94d07n33p7wqjf0g9vtf3k7",
    "moniker": "Primary Node - Songs"
  },
  {
    "operatorAddress": "tcrocncl1x5dtk7fus777m72hca8qa8pcgk6ezkdhaddhn9",
    "primaryAddress": "tcro1x5dtk7fus777m72hca8qa8pcgk6ezkdhgjwwtx",
    "moniker": "Primary Node - Video Games"
  },
  {
    "operatorAddress": "tcrocncl1p4fzn6ta24c6ek4v2qls6y5uug44ku9txmzp92",
    "primaryAddress": "tcro1p4fzn6ta24c6ek4v2qls6y5uug44ku9tnypcaf",
    "moniker": "Primary Node - Culture"
  },
  {
    "operatorAddress": "tcrocncl1rl3sgwzrf6qt8xrucexx0970f9ytve4x5403d3",
    "primaryAddress": "tcro1rl3sgwzrf6qt8xrucexx0970f9ytve4xp2vg4j",
    "moniker": "stef18"
  },
  {
    "operatorAddress": "tcrocncl1jz30jstc0qup7hunxarp5z636va6dk5zfua6wj",
    "primaryAddress": "tcro1jz30jstc0qup7hunxarp5z636va6dk5zur7rk3",
    "moniker": "TOBG-ND01"
  },
  {
    "operatorAddress": "tcrocncl1843dda5rqu68ypa5aggsvc3pcagnjdfhmwxkcj",
    "primaryAddress": "tcro1843dda5rqu68ypa5aggsvc3pcagnjdfhw390q3",
    "moniker": "[baronek-node]"
  },
  {
    "operatorAddress": "tcrocncl10ufcsvyln95xaw96twenswep5a5wt0e6yht95x",
    "primaryAddress": "tcro10ufcsvyln95xaw96twenswep5a5wt0e63gguv9",
    "moniker": "Primary Node - Rocks"
  },
  {
    "operatorAddress": "tcrocncl1znthgf6k9n33pfcpymlnr6ezw9v7chv4vxeq3g",
    "primaryAddress": "tcro1znthgf6k9n33pfcpymlnr6ezw9v7chv4ee6eft",
    "moniker": "Primary Node - Oceans"
  },
  {
    "operatorAddress": "tcrocncl18n7h83tr3d6nygr3089u022uqnhclwdrw0wvj4",
    "primaryAddress": "tcro18n7h83tr3d6nygr3089u022uqnhclwdrmsd42k",
    "moniker": "Primary Node - States"
  },
  {
    "operatorAddress": "tcrocncl1rjr6yekcqzle3k4w60pax6tq02l4fphp5hx67k",
    "primaryAddress": "tcro1rjr6yekcqzle3k4w60pax6tq02l4fphppg9rx4",
    "moniker": "misakas-node"
  },
  {
    "operatorAddress": "tcrocncl1ws7zlvktf9avv0vj6jejfz0egudpjlvysfz245",
    "primaryAddress": "tcro1ws7zlvktf9avv0vj6jejfz0egudpjlvy9kpndh",
    "moniker": "madmax"
  },
  {
    "operatorAddress": "tcrocncl1dgl5gtayw4uwpvcmfsgnfxlwagct0lp7knr3l7",
    "primaryAddress": "tcro1dgl5gtayw4uwpvcmfsgnfxlwagct0lp7rvqg8a",
    "moniker": "ivanodes"
  },
  {
    "operatorAddress": "tcrocncl1fnhd627p5rh58898mswkmam7cj3srgahknftvu",
    "primaryAddress": "tcro1fnhd627p5rh58898mswkmam7cj3srgahrv2j5l",
    "moniker": "cdc-testnet-node"
  },
  {
    "operatorAddress": "tcrocncl1q726n4j8rtd0putajp2mw82da0etdnvmm4l99g",
    "primaryAddress": "tcro1q726n4j8rtd0putajp2mw82da0etdnvmw2uuat",
    "moniker": "allright"
  },
  {
    "operatorAddress": "tcrocncl1a4g54l95ruq7psxv4sa22gw3skmgtua5vhh5ra",
    "primaryAddress": "tcro1a4g54l95ruq7psxv4sa22gw3skmgtua5eg5dm7",
    "moniker": "alitaangel"
  },
  {
    "operatorAddress": "tcrocncl197hqjnvgw85nerrheg0g4u6ylfukv6np2a7vsf",
    "primaryAddress": "tcro197hqjnvgw85nerrheg0g4u6ylfukv6nplza4g2",
    "moniker": "thecrown"
  },
  {
    "operatorAddress": "tcrocncl125v62z943x7jt0smzsz8qcaapesql7tuwaxcyc",
    "primaryAddress": "tcro125v62z943x7jt0smzsz8qcaapesql7tumz9pum",
    "moniker": "sinitsy"
  },
  {
    "operatorAddress": "tcrocncl12vefvd03zpa4z6pej55rx5f9kvghfd2mrwq33n",
    "primaryAddress": "tcro12vefvd03zpa4z6pej55rx5f9kvghfd2mk3rgfs",
    "moniker": "snegiri"
  },
  {
    "operatorAddress": "tcrocncl1utw780mmmylrrqxagvkp74vcqc6n5zuzj54ffe",
    "primaryAddress": "tcro1utw780mmmylrrqxagvkp74vcqc6n5zuz8tks36",
    "moniker": "numstake"
  },
  {
    "operatorAddress": "tcrocncl1ye9upz62pmfru8m6x4xhe7cyw0xwutt96qz4l4",
    "primaryAddress": "tcro1ye9upz62pmfru8m6x4xhe7cyw0xwutt90lpv8k",
    "moniker": "evgenfarm"
  },
  {
    "operatorAddress": "tcrocncl1yp07r5v78fukq8nq789rhdcqhfyxryesv5ka6u",
    "primaryAddress": "tcro1yp07r5v78fukq8nq789rhdcqhfyxryeset4yzl",
    "moniker": "dieboss"
  },
  {
    "operatorAddress": "tcrocncl1g8u0gw4ttg86suvwm9lzs32afph7lfacr9kukr",
    "primaryAddress": "tcro1g8u0gw4ttg86suvwm9lzs32afph7lfack649wq",
    "moniker": "blizardinio"
  },
  {
    "operatorAddress": "tcrocncl1auvfj444nfw7993kgl005pz5622uq0d8u8rta3",
    "primaryAddress": "tcro1auvfj444nfw7993kgl005pz5622uq0d8fcqj9j",
    "moniker": "iLucky"
  },
  {
    "operatorAddress": "tcrocncl1f4ndrukrzhzg32ch75kfh237shz88shlvuqld6",
    "primaryAddress": "tcro1f4ndrukrzhzg32ch75kfh237shz88shlerrx4e",
    "moniker": "circle"
  },
  {
    "operatorAddress": "tcrocncl10mzlhga4h8ycjh9p4z9ktpada3vymd75lh29sy",
    "primaryAddress": "tcro10mzlhga4h8ycjh9p4z9ktpada3vymd752gfug8",
    "moniker": "smorodina"
  },
  {
    "operatorAddress": "tcrocncl1h73zkwmstye2twcuj37nrlv2na8y0dr9jt8c8r",
    "primaryAddress": "tcro1h73zkwmstye2twcuj37nrlv2na8y0dr985yplq",
    "moniker": "DrDre"
  },
  {
    "operatorAddress": "tcrocncl1gzqqecshr520pzffsgaa5ne5d9ce23euwkk4jc",
    "primaryAddress": "tcro1gzqqecshr520pzffsgaa5ne5d9ce23eumf4v2m",
    "moniker": "zverev"
  },
  {
    "operatorAddress": "tcrocncl107sjzaxjlwqa5l85am9zr2lldsrsnc3ap8qd53",
    "primaryAddress": "tcro107sjzaxjlwqa5l85am9zr2lldsrsnc3a5cr5vj",
    "moniker": "gitbash"
  },
  {
    "operatorAddress": "tcrocncl1dhan0e7d8lfghr5qcw5zwvkmy55cydjw953g3w",
    "primaryAddress": "tcro1dhan0e7d8lfghr5qcw5zwvkmy55cydjwstj3fd",
    "moniker": "zankinpro"
  },
  {
    "operatorAddress": "tcrocncl1htn47ayr6fnzyrwczah2cyt2fn2cljgp4n5juk",
    "primaryAddress": "tcro1htn47ayr6fnzyrwczah2cyt2fn2cljgpqvhty4",
    "moniker": "basliking"
  },
  {
    "operatorAddress": "tcrocncl1gwjlsn3n7ww96x6c69e3t34rm7rztm526nkejy",
    "primaryAddress": "tcro1gwjlsn3n7ww96x6c69e3t34rm7rztm520v4q28",
    "moniker": "maximka"
  },
  {
    "operatorAddress": "tcrocncl1acy63xj9y9ac7w9kp5nk400ht5w6tk6kf46xf7",
    "primaryAddress": "tcro1acy63xj9y9ac7w9kp5nk400ht5w6tk6ku2el3a",
    "moniker": "MambaTrex"
  },
  {
    "operatorAddress": "tcrocncl1yn3hffqlqkzahjcqwqygmvyg8amwz4870tafhy",
    "primaryAddress": "tcro1yn3hffqlqkzahjcqwqygmvyg8amwz487657s08",
    "moniker": "alelagu"
  },
  {
    "operatorAddress": "tcrocncl12ezeg6e2gjv70qf5y6rvckjmnctzk33e3z9n72",
    "primaryAddress": "tcro12ezeg6e2gjv70qf5y6rvckjmnctzk33eyax2xf",
    "moniker": "mmuscle"
  },
  {
    "operatorAddress": "tcrocncl189wqe9d97c2f5cfpvduy4jyxlvy2qetfldcdgq",
    "primaryAddress": "tcro189wqe9d97c2f5cfpvduy4jyxlvy2qetf2jm5sr",
    "moniker": "submarine"
  },
  {
    "operatorAddress": "tcrocncl14fuqs856uglxrzms2ed4tfx2693zl838mkrp7c",
    "primaryAddress": "tcro14fuqs856uglxrzms2ed4tfx2693zl838wfqcxm",
    "moniker": "FrodoBeginz"
  },
  {
    "operatorAddress": "tcrocncl1kphhg84kkty64v2ldr9cw2ehzprdjkd4jtzalx",
    "primaryAddress": "tcro1kphhg84kkty64v2ldr9cw2ehzprdjkd485py89",
    "moniker": "afisport"
  },
  {
    "operatorAddress": "tcrocncl1vjeaqmrufrzx0g5e4ejnuevs2gzaljzd4df525",
    "primaryAddress": "tcro1vjeaqmrufrzx0g5e4ejnuevs2gzaljzdqj2djh",
    "moniker": "greshnikKing"
  },
  {
    "operatorAddress": "tcrocncl1pnk3qkasc42v4cl906l9ut7erftcp23vqrrm8m",
    "primaryAddress": "tcro1pnk3qkasc42v4cl906l9ut7erftcp23v4uqzlc",
    "moniker": "gummybear"
  },
  {
    "operatorAddress": "tcrocncl1mt4askyztg4nmw7kvtmf6nmcssd30aszc2a8ml",
    "primaryAddress": "tcro1mt4askyztg4nmw7kvtmf6nmcssd30aszd477ru",
    "moniker": "HyperionPrime_HP"
  },
  {
    "operatorAddress": "tcrocncl1kphwd20ugw9753pq6hnf5yge0q2w7zsazv4ffk",
    "primaryAddress": "tcro1kphwd20ugw9753pq6hnf5yge0q2w7zsahnks34",
    "moniker": "yanci-node"
  },
  {
    "operatorAddress": "tcrocncl1cfdcu4j5as9s0jltjvk098qud8254jughrv6he",
    "primaryAddress": "tcro1cfdcu4j5as9s0jltjvk098qud8254jugzu0r06",
    "moniker": "Vladimir"
  },
  {
    "operatorAddress": "tcrocncl1cnst3q53z5rkf5hal4rqtejwtza9f482aj232g",
    "primaryAddress": "tcro1cnst3q53z5rkf5hal4rqtejwtza9f482gdfgjt",
    "moniker": "cronode01"
  },
  {
    "operatorAddress": "tcrocncl178x2mszn6j4tq7x0erp8l48x0aaedf3qnkc694",
    "primaryAddress": "tcro178x2mszn6j4tq7x0erp8l48x0aaedf3qxfmrak",
    "moniker": "sap"
  },
  {
    "operatorAddress": "tcrocncl1tlmf7783fdd9u2kgvwkx7pukply60dfu28a4w7",
    "primaryAddress": "tcro1tlmf7783fdd9u2kgvwkx7pukply60dfulc7vka",
    "moniker": "Staking Fund"
  },
  {
    "operatorAddress": "tcrocncl1prqwler9xp47zct4l57sqttyqu88j9f4kf9alt",
    "primaryAddress": "tcro1prqwler9xp47zct4l57sqttyqu88j9f4rkxy8g",
    "moniker": "ArcticNode"
  },
  {
    "operatorAddress": "tcrocncl15lkku5g5ggaazld60ryn8tseepf8ckx3vumk8y",
    "primaryAddress": "tcro15lkku5g5ggaazld60ryn8tseepf8ckx3erc0l8",
    "moniker": "Ubik Capital"
  },
  {
    "operatorAddress": "tcrocncl1elpwuqhr5chnhx85xc6lt3jv0yngghrzrnchwz",
    "primaryAddress": "tcro1elpwuqhr5chnhx85xc6lt3jv0yngghrzkvmwkp",
    "moniker": "Nuggets"
  },
  {
    "operatorAddress": "tcrocncl1msmlnzyqwdnmzq0seznk2qn5gds5fh4g9n6exm",
    "primaryAddress": "tcro1msmlnzyqwdnmzq0seznk2qn5gds5fh4gsveq7c",
    "moniker": "Cookies-node"
  },
  {
    "operatorAddress": "tcrocncl1tjr08cre799ujw39f3gwkv9ls9h222cavpv79f",
    "primaryAddress": "tcro1tjr08cre799ujw39f3gwkv9ls9h222cae708a2",
    "moniker": "SNK"
  },
  {
    "operatorAddress": "tcrocncl1t5nkrt4yq4vud6durp5mn03sujgsgku7lsnndj",
    "primaryAddress": "tcro1t5nkrt4yq4vud6durp5mn03sujgsgku720s243",
    "moniker": "papsan"
  },
  {
    "operatorAddress": "tcrocncl129xlqcjasyzqsc4364x7hc6m6t5w33ruqp53qx",
    "primaryAddress": "tcro129xlqcjasyzqsc4364x7hc6m6t5w33ru47hgc9",
    "moniker": "web34ever"
  },
  {
    "operatorAddress": "tcrocncl1uy43rscp9x924vfdrg26l5w4pfkhjrnm6rstya",
    "primaryAddress": "tcro1uy43rscp9x924vfdrg26l5w4pfkhjrnm0unju7",
    "moniker": "GoldenCoin"
  },
  {
    "operatorAddress": "tcrocncl1ezpz4a88cq7qvkaddwl7hsphq74xnz4sdsgzg2",
    "primaryAddress": "tcro1ezpz4a88cq7qvkaddwl7hsphq74xnz4sc0tmsf",
    "moniker": "validation-capital"
  },
  {
    "operatorAddress": "tcrocncl18lnzn7wsunj0p3lsvy4kx4xk9h0esuupnx0pyy",
    "primaryAddress": "tcro18lnzn7wsunj0p3lsvy4kx4xk9h0esuupxevcu8",
    "moniker": "burbank"
  },
  {
    "operatorAddress": "tcrocncl1hysfxsrk4vmzaumd0quc30nrcz7y06vy2382qw",
    "primaryAddress": "tcro1hysfxsrk4vmzaumd0quc30nrcz7y06vylwyncd",
    "moniker": "CroStarNode"
  },
  {
    "operatorAddress": "tcrocncl1cfj0v88dg9zwy93pncxxrg5auefu9sztpkna3j",
    "primaryAddress": "tcro1cfj0v88dg9zwy93pncxxrg5auefu9szt5fsyf3",
    "moniker": "Manris"
  },
  {
    "operatorAddress": "tcrocncl1e0xzzy5mpjgv6t2d3ej499mm68llg8q2qqgxyz",
    "primaryAddress": "tcro1e0xzzy5mpjgv6t2d3ej499mm68llg8q24ltlup",
    "moniker": "Miromice"
  },
  {
    "operatorAddress": "tcrocncl1u572s2suvqcrpqzcuyj9xsqy3qlg7twlyqjdf9",
    "primaryAddress": "tcro1u572s2suvqcrpqzcuyj9xsqy3qlg7twl3l353x",
    "moniker": "Kison"
  },
  {
    "operatorAddress": "tcrocncl14ntlx0pzfdjysfa6ldmzsh8uwkm2sawpfvcmf6",
    "primaryAddress": "tcro14ntlx0pzfdjysfa6ldmzsh8uwkm2sawpunmz3e",
    "moniker": "Legnum"
  },
  {
    "operatorAddress": "tcrocncl1c2kkvsj44pej54n0q24f9cumff7tmce4tjjyz0",
    "primaryAddress": "tcro1c2kkvsj44pej54n0q24f9cumff7tmce47d3a6v",
    "moniker": "angristan"
  },
  {
    "operatorAddress": "tcrocncl178acj8stdm6a5p5a22pn2k9aspqsvfd6lp0r92",
    "primaryAddress": "tcro178acj8stdm6a5p5a22pn2k9aspqsvfd627v6af",
    "moniker": "COVID1980"
  },
  {
    "operatorAddress": "tcrocncl10ddjxnjy6fk8l9pxgzt52yvrd0670ytklx06fl",
    "primaryAddress": "tcro10ddjxnjy6fk8l9pxgzt52yvrd0670ytk2evr3u",
    "moniker": "krtchain"
  },
  {
    "operatorAddress": "tcrocncl1w6flwh952flt48ad6dts8fstemlugzc3af5u5w",
    "primaryAddress": "tcro1w6flwh952flt48ad6dts8fstemlugzc3gkh9vd",
    "moniker": "Joy INA"
  },
  {
    "operatorAddress": "tcrocncl1e77t46hwyzym5tjwmhaccqdezh89g40ur08fzd",
    "primaryAddress": "tcro1e77t46hwyzym5tjwmhaccqdezh89g40uksys6w",
    "moniker": "MotherTech"
  },
  {
    "operatorAddress": "tcrocncl1fevd9hvdw7yzz45mjhzca2cmhyq7jtpnvwqnrs",
    "primaryAddress": "tcro1fevd9hvdw7yzz45mjhzca2cmhyq7jtpne3r2mn",
    "moniker": "Army IDs"
  },
  {
    "operatorAddress": "tcrocncl1yrjyu94xeqrdgkk6czs0xhxypq6gru83xvp5zl",
    "primaryAddress": "tcro1yrjyu94xeqrdgkk6czs0xhxypq6gru83nnzd6u",
    "moniker": "Moonshine"
  },
  {
    "operatorAddress": "tcrocncl1q982erd578mvkr36psc4pakshvmpy793wfv4gp",
    "primaryAddress": "tcro1q982erd578mvkr36psc4pakshvmpy793mk0vsz",
    "moniker": "vemulam"
  },
  {
    "operatorAddress": "tcrocncl1zyy48y6fdvnrs9na5dsdn84839yjrsj5q80dzd",
    "primaryAddress": "tcro1zyy48y6fdvnrs9na5dsdn84839yjrsj54cv56w",
    "moniker": "alexil-node"
  },
  {
    "operatorAddress": "tcrocncl1kfjsq9xjk5mwnf02meygqcyngdnak3q5zdajye",
    "primaryAddress": "tcro1kfjsq9xjk5mwnf02meygqcyngdnak3q5hj7tu6",
    "moniker": "Proxenet"
  },
  {
    "operatorAddress": "tcrocncl1h7xcd8qs2nq4zkw2clyyv3sgyvku68eca5c6le",
    "primaryAddress": "tcro1h7xcd8qs2nq4zkw2clyyv3sgyvku68ecgtmr86",
    "moniker": "keitaj"
  },
  {
    "operatorAddress": "tcrocncl1dc4g5wecgmq7nhfq294w6r5k3v5fcu77gdd7qj",
    "primaryAddress": "tcro1dc4g5wecgmq7nhfq294w6r5k3v5fcu77ajw8c3",
    "moniker": "Viking"
  },
  {
    "operatorAddress": "tcrocncl13hmw5uryuwu2rrlp8maswgfzfrclv87zyx6cly",
    "primaryAddress": "tcro13hmw5uryuwu2rrlp8maswgfzfrclv87z3eep88",
    "moniker": "bristly0range-node"
  },
  {
    "operatorAddress": "tcrocncl1j66cts730l0dxszuclr5d4dkh72zhta6rmwvap",
    "primaryAddress": "tcro1j66cts730l0dxszuclr5d4dkh72zhta6kyd49z",
    "moniker": "frenchchocolatine-node"
  },
  {
    "operatorAddress": "tcrocncl12mq9gaxay66u5pea3qjdw28wsq7jk3gmfgdphe",
    "primaryAddress": "tcro12mq9gaxay66u5pea3qjdw28wsq7jk3gmuhwc06",
    "moniker": "sage-node"
  },
  {
    "operatorAddress": "tcrocncl1kvmcqussxv7usl6wldysd62hudnhffgyzj0679",
    "primaryAddress": "tcro1kvmcqussxv7usl6wldysd62hudnhffgyhdvrxx",
    "moniker": "ShadowValidator"
  },
  {
    "operatorAddress": "tcrocncl1tpp982scz3m5asw05h9zm6s3dw3sql6v2z2z8v",
    "primaryAddress": "tcro1tpp982scz3m5asw05h9zm6s3dw3sql6vlafml0",
    "moniker": "vanek_1"
  },
  {
    "operatorAddress": "tcrocncl1g23sp6mq386fjmqaklkk4qmlmpus9nn3q8mdxa",
    "primaryAddress": "tcro1g23sp6mq386fjmqaklkk4qmlmpus9nn34cc577",
    "moniker": "Konung"
  },
  {
    "operatorAddress": "tcrocncl1ejp5mkshacdkkmg9umemgqcxxz84favxldz0jl",
    "primaryAddress": "tcro1ejp5mkshacdkkmg9umemgqcxxz84favx2jpk2u",
    "moniker": "nextyp"
  },
  {
    "operatorAddress": "tcrocncl17afqrf9zy7p0gvlrdalcexxkhzuqxfvvcuq3l0",
    "primaryAddress": "tcro17afqrf9zy7p0gvlrdalcexxkhzuqxfvvdrrg8v",
    "moniker": "KANDUKURUKISHORE"
  },
  {
    "operatorAddress": "tcrocncl1s2jv8j77gqsh08xfj0pcry5tp2v6jwf2cqs4tm",
    "primaryAddress": "tcro1s2jv8j77gqsh08xfj0pcry5tp2v6jwf2dlnvnc",
    "moniker": "duncanski"
  },
  {
    "operatorAddress": "tcrocncl1fmyhsn9j2769agkljf608dycy3p07sj82j8wqf",
    "primaryAddress": "tcro1fmyhsn9j2769agkljf608dycy3p07sj8ldyhc2",
    "moniker": "AsiaGodTone-gg3be0"
  },
  {
    "operatorAddress": "tcrocncl1za5e5guwhgfwuuqgq8vf884rw2pt7n02a4n0c7",
    "primaryAddress": "tcro1za5e5guwhgfwuuqgq8vf884rw2pt7n02g2skqa",
    "moniker": "prlx"
  },
  {
    "operatorAddress": "tcrocncl10nlv439f2fn35krxulcm0dnz5x5k2tsej2gvpe",
    "primaryAddress": "tcro10nlv439f2fn35krxulcm0dnz5x5k2tse84t4e6",
    "moniker": "cnewex-node"
  },
  {
    "operatorAddress": "tcrocncl1dxy4gn4lhngxp7kk306qmnnzj4qvdlvl5httlr",
    "primaryAddress": "tcro1dxy4gn4lhngxp7kk306qmnnzj4qvdlvlpggj8q",
    "moniker": "Cryptonodeab1"
  },
  {
    "operatorAddress": "tcrocncl1gdxltu8a3ulwkmvcslcwpujf4u2v0ada9yhh3k",
    "primaryAddress": "tcro1gdxltu8a3ulwkmvcslcwpujf4u2v0adasm5wf4",
    "moniker": "zeus-node-0"
  },
  {
    "operatorAddress": "tcrocncl1x8sk5th32e75g98szmwyjsvk07nu5gpk6hgxyw",
    "primaryAddress": "tcro1x8sk5th32e75g98szmwyjsvk07nu5gpk0gtlud",
    "moniker": "Staking_Hub"
  },
  {
    "operatorAddress": "tcrocncl109ww3ss92v4vsaq470vvgw528mtqp98m4s04ex",
    "primaryAddress": "tcro109ww3ss92v4vsaq470vvgw528mtqp98mq0vvp9",
    "moniker": "leo-node"
  },
  {
    "operatorAddress": "tcrocncl1d678ylljqv8zsmktf64pmh90sg5ypewc83fjvu",
    "primaryAddress": "tcro1d678ylljqv8zsmktf64pmh90sg5ypewcjw2t5l",
    "moniker": "unit1"
  }
]`
