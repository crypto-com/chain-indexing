package usecase_parser_test

const TX_MSG_DELEGATE_TX_BLOCK_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "block_id": {
            "hash": "68C2916D2A4B8E398AEAE831984A14D011F83E5B11B9BE871D1825455725DB54",
            "parts": {
                "total": 1,
                "hash": "E81AC01184D90B0F96E5DF74FFDB9EDA8353FDBF8E4822547BB06B8DC4F99420"
            }
        },
        "block": {
            "header": {
                "version": {
                    "block": "11"
                },
                "chain_id": "testnet-croeseid-4",
                "height": "22232949",
                "time": "2025-06-16T06:10:00.644379482Z",
                "last_block_id": {
                    "hash": "22D263EBE9D5BF1CC257F24FCA293E3AAFEE9FB1B2BA74EF04237BE90A65D0B2",
                    "parts": {
                        "total": 1,
                        "hash": "8B65BAB84CF025913E0FC5299B160EA9A2BBB1D8B59319DE907742271A4A05CB"
                    }
                },
                "last_commit_hash": "627F877A45E9E75BCEAB028BA6966393F0F37A25B224A9FEFB005A0197E34375",
                "data_hash": "DAD4585CE91813054C7B8766294148A125E83FACB3B0DD67A271CA97353F212A",
                "validators_hash": "7C0AC8CB45AB80871E4119EE3E0C459C3B2DF754760AE29ECBF68FD483F280F5",
                "next_validators_hash": "7C0AC8CB45AB80871E4119EE3E0C459C3B2DF754760AE29ECBF68FD483F280F5",
                "consensus_hash": "372B4AE845086C837EFEF79A189B085B1FD6610C53F3BEB17EE0E27B347C06DE",
                "app_hash": "7176CD3CBE42E0106441366EA7C99FFCCB0A7F2609D304E7D834848B5BCA6EDB",
                "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "proposer_address": "7E2B455523A35AE8F4064C00AF01515FEF673079"
            },
            "data": {
                "txs": [
                    "CqIBCp8BCiMvY29zbW9zLnN0YWtpbmcudjFiZXRhMS5Nc2dEZWxlZ2F0ZRJ4Cit0Y3JvMWd5Y3p1azR5bDJqYWU2cGd1MHFjMnR6dGM4ODltbjd2eDl5ajgzEi90Y3JvY25jbDFyZDh5ZTI3a2R5Mmg5MmhsazZ4enFhYTh0cDJ3ZDl2ZnRwemhlaxoYCghiYXNldGNybxIMMjc4ODM5MzU3ODg3EmsKUQpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQP6UH0bEygBke70Fnk0zRnT/7E+pe5KtB2Og/vGikJeIhIECgIIARjnOxIWChAKCGJhc2V0Y3JvEgQ3MDAwEMCLERpAsmusQ1riqxyNKNU9xskGCffk+PcBLIllKs99MZ/dJ40zDFLilpJNwo/2FrlBjR8HKo2GCGQ2jFUIpKfM7lku4w=="
                ]
            },
            "evidence": {
                "evidence": []
            },
            "last_commit": {
                "height": "22232948",
                "round": 0,
                "block_id": {
                    "hash": "22D263EBE9D5BF1CC257F24FCA293E3AAFEE9FB1B2BA74EF04237BE90A65D0B2",
                    "parts": {
                        "total": 1,
                        "hash": "8B65BAB84CF025913E0FC5299B160EA9A2BBB1D8B59319DE907742271A4A05CB"
                    }
                },
                "signatures": [
                    {
                        "block_id_flag": 2,
                        "validator_address": "7E2B455523A35AE8F4064C00AF01515FEF673079",
                        "timestamp": "2025-06-16T06:10:00.644379482Z",
                        "signature": "vHEqT8Q3il0n/znAuHcsALnjVwb+DjyhFw7Tq3oJe9g0HZ+6cY0TPDyB/kXmaUpOIKJ5oUAalVZdGcjGmzP9Cg=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "F9A4582896E5F2692BEDDD7C251755E33D2DE013",
                        "timestamp": "2025-06-16T06:10:00.657702374Z",
                        "signature": "L7u4vGqZPHjYkSnEzDmQfCplSRpPi1LqGvsJR/SaHU4LqSG4RFUw2N7rtt8unf/jc4aLQA4UjNGzzOAskdElDQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "2ACA46D6C1BBECE7176107DFF4C7D14AD18FF54D",
                        "timestamp": "2025-06-16T06:10:00.644240502Z",
                        "signature": "oxp2XJ4taWz5ilc7iBRvjtcDr1ene51B/F9CAV+V9FvIf4VovlZ2gG1R1oTyCwMid4psePuqrUj77FBLCTZUAg=="
                    },
                    {
                        "block_id_flag": 1,
                        "validator_address": "",
                        "timestamp": "0001-01-01T00:00:00Z",
                        "signature": null
                    }
                ]
            }
        }
    }
}`

const TX_MSG_DELEGATE_TX_BLOCK_RESULTS_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "height": "22232949",
        "txs_results": [
            {
                "code": 0,
                "data": "Ei0KKy9jb3Ntb3Muc3Rha2luZy52MWJldGExLk1zZ0RlbGVnYXRlUmVzcG9uc2U=",
                "log": "",
                "info": "",
                "gas_wanted": "280000",
                "gas_used": "174260",
                "events": [
                    {
                        "type": "coin_spent",
                        "attributes": [
                            {
                                "key": "spender",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83",
                                "index": true
                            },
                            {
                                "key": "amount",
                                "value": "7000basetcro",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "coin_received",
                        "attributes": [
                            {
                                "key": "receiver",
                                "value": "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
                                "index": true
                            },
                            {
                                "key": "amount",
                                "value": "7000basetcro",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "recipient",
                                "value": "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
                                "index": true
                            },
                            {
                                "key": "sender",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83",
                                "index": true
                            },
                            {
                                "key": "amount",
                                "value": "7000basetcro",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "sender",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "fee",
                                "value": "7000basetcro",
                                "index": true
                            },
                            {
                                "key": "fee_payer",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "acc_seq",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83/7655",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "signature",
                                "value": "smusQ1riqxyNKNU9xskGCffk+PcBLIllKs99MZ/dJ40zDFLilpJNwo/2FrlBjR8HKo2GCGQ2jFUIpKfM7lku4w==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "action",
                                "value": "/cosmos.staking.v1beta1.MsgDelegate",
                                "index": true
                            },
                            {
                                "key": "sender",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83",
                                "index": true
                            },
                            {
                                "key": "module",
                                "value": "staking",
                                "index": true
                            },
                            {
                                "key": "msg_index",
                                "value": "0",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "coin_spent",
                        "attributes": [
                            {
                                "key": "spender",
                                "value": "tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l",
                                "index": true
                            },
                            {
                                "key": "amount",
                                "value": "267926095292basetcro",
                                "index": true
                            },
                            {
                                "key": "msg_index",
                                "value": "0",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "coin_received",
                        "attributes": [
                            {
                                "key": "receiver",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83",
                                "index": true
                            },
                            {
                                "key": "amount",
                                "value": "267926095292basetcro",
                                "index": true
                            },
                            {
                                "key": "msg_index",
                                "value": "0",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "recipient",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83",
                                "index": true
                            },
                            {
                                "key": "sender",
                                "value": "tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l",
                                "index": true
                            },
                            {
                                "key": "amount",
                                "value": "267926095292basetcro",
                                "index": true
                            },
                            {
                                "key": "msg_index",
                                "value": "0",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "sender",
                                "value": "tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l",
                                "index": true
                            },
                            {
                                "key": "msg_index",
                                "value": "0",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "withdraw_rewards",
                        "attributes": [
                            {
                                "key": "amount",
                                "value": "267926095292basetcro",
                                "index": true
                            },
                            {
                                "key": "validator",
                                "value": "tcrocncl1rd8ye27kdy2h92hlk6xzqaa8tp2wd9vftpzhek",
                                "index": true
                            },
                            {
                                "key": "delegator",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83",
                                "index": true
                            },
                            {
                                "key": "msg_index",
                                "value": "0",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "coin_spent",
                        "attributes": [
                            {
                                "key": "spender",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83",
                                "index": true
                            },
                            {
                                "key": "amount",
                                "value": "278839357887basetcro",
                                "index": true
                            },
                            {
                                "key": "msg_index",
                                "value": "0",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "coin_received",
                        "attributes": [
                            {
                                "key": "receiver",
                                "value": "tcro1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3r4gj9h",
                                "index": true
                            },
                            {
                                "key": "amount",
                                "value": "278839357887basetcro",
                                "index": true
                            },
                            {
                                "key": "msg_index",
                                "value": "0",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "delegate",
                        "attributes": [
                            {
                                "key": "validator",
                                "value": "tcrocncl1rd8ye27kdy2h92hlk6xzqaa8tp2wd9vftpzhek",
                                "index": true
                            },
                            {
                                "key": "delegator",
                                "value": "tcro1gyczuk4yl2jae6pgu0qc2tztc889mn7vx9yj83",
                                "index": true
                            },
                            {
                                "key": "amount",
                                "value": "278839357887basetcro",
                                "index": true
                            },
                            {
                                "key": "new_shares",
                                "value": "278839357887.000000000000000000",
                                "index": true
                            },
                            {
                                "key": "msg_index",
                                "value": "0",
                                "index": true
                            }
                        ]
                    }
                ],
                "codespace": ""
            }
        ],
        "finalize_block_events": [
            {
                "type": "coin_received",
                "attributes": [
                    {
                        "key": "receiver",
                        "value": "tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq",
                        "index": true
                    },
                    {
                        "key": "amount",
                        "value": "15448024645basetcro",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "coinbase",
                "attributes": [
                    {
                        "key": "minter",
                        "value": "tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq",
                        "index": true
                    },
                    {
                        "key": "amount",
                        "value": "15448024645basetcro",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_spent",
                "attributes": [
                    {
                        "key": "spender",
                        "value": "tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq",
                        "index": true
                    },
                    {
                        "key": "amount",
                        "value": "15448024645basetcro",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_received",
                "attributes": [
                    {
                        "key": "receiver",
                        "value": "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
                        "index": true
                    },
                    {
                        "key": "amount",
                        "value": "15448024645basetcro",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "transfer",
                "attributes": [
                    {
                        "key": "recipient",
                        "value": "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
                        "index": true
                    },
                    {
                        "key": "sender",
                        "value": "tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq",
                        "index": true
                    },
                    {
                        "key": "amount",
                        "value": "15448024645basetcro",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "message",
                "attributes": [
                    {
                        "key": "sender",
                        "value": "tcro1m3h30wlvsf8llruxtpukdvsy0km2kum87lx9mq",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "mint",
                "attributes": [
                    {
                        "key": "bonded_ratio",
                        "value": "0.066383588729985860",
                        "index": true
                    },
                    {
                        "key": "inflation",
                        "value": "0.010000000000000000",
                        "index": true
                    },
                    {
                        "key": "annual_provisions",
                        "value": "97500516510890338.160000000000000000",
                        "index": true
                    },
                    {
                        "key": "amount",
                        "value": "15448024645",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_spent",
                "attributes": [
                    {
                        "key": "spender",
                        "value": "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
                        "index": true
                    },
                    {
                        "key": "amount",
                        "value": "15448024645basetcro",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_received",
                "attributes": [
                    {
                        "key": "receiver",
                        "value": "tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l",
                        "index": true
                    },
                    {
                        "key": "amount",
                        "value": "15448024645basetcro",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "transfer",
                "attributes": [
                    {
                        "key": "recipient",
                        "value": "tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l",
                        "index": true
                    },
                    {
                        "key": "sender",
                        "value": "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
                        "index": true
                    },
                    {
                        "key": "amount",
                        "value": "15448024645basetcro",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "message",
                "attributes": [
                    {
                        "key": "sender",
                        "value": "tcro17xpfvakm2amg962yls6f84z3kell8c5lxhzaha",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "amount",
                        "value": "633149867.253093370121380725basetcro",
                        "index": true
                    },
                    {
                        "key": "validator",
                        "value": "tcrocncl1s4ggq2zuzvwg5k8vnx2xfwtdm4cz6wtnuqkl7a",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "amount",
                        "value": "6331498672.530933701213807252basetcro",
                        "index": true
                    },
                    {
                        "key": "validator",
                        "value": "tcrocncl1s4ggq2zuzvwg5k8vnx2xfwtdm4cz6wtnuqkl7a",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "amount",
                        "value": "633085272.530721704642863140basetcro",
                        "index": true
                    },
                    {
                        "key": "validator",
                        "value": "tcrocncl163tv59yzgeqcap8lrsa2r4zk580h8ddr5a0sdd",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "amount",
                        "value": "6330852725.307217046428631397basetcro",
                        "index": true
                    },
                    {
                        "key": "validator",
                        "value": "tcrocncl163tv59yzgeqcap8lrsa2r4zk580h8ddr5a0sdd",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "amount",
                        "value": "45715602.071381722970467740basetcro",
                        "index": true
                    },
                    {
                        "key": "validator",
                        "value": "tcrocncl1rd8ye27kdy2h92hlk6xzqaa8tp2wd9vftpzhek",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "amount",
                        "value": "457156020.713817229704677396basetcro",
                        "index": true
                    },
                    {
                        "key": "validator",
                        "value": "tcrocncl1rd8ye27kdy2h92hlk6xzqaa8tp2wd9vftpzhek",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "amount",
                        "value": "5656764.849015998195621028basetcro",
                        "index": true
                    },
                    {
                        "key": "validator",
                        "value": "tcrocncl1uw77lwy0zfwuuk6a7w6vre4wccweeevu025kny",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "amount",
                        "value": "11313529.698031996391242056basetcro",
                        "index": true
                    },
                    {
                        "key": "validator",
                        "value": "tcrocncl1uw77lwy0zfwuuk6a7w6vre4wccweeevu025kny",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            },
            {
                "type": "liveness",
                "attributes": [
                    {
                        "key": "address",
                        "value": "tcrocnclcons19v0javd6k6n74aastnd05vvzw55t2u82w6skd8",
                        "index": true
                    },
                    {
                        "key": "missed_blocks",
                        "value": "76",
                        "index": true
                    },
                    {
                        "key": "height",
                        "value": "22232949",
                        "index": true
                    },
                    {
                        "key": "mode",
                        "value": "BeginBlock",
                        "index": true
                    }
                ]
            }
        ],
        "validator_updates": [
            {
                "pub_key": {
                    "Sum": {
                        "type": "tendermint.crypto.PublicKey_Ed25519",
                        "value": {
                            "ed25519": "YNZKR5IF9Q2PpBPXePZ8yZ7AeYpClF6Cg4rOlnDG4gg="
                        }
                    }
                },
                "power": "22534378344"
            }
        ],
        "consensus_param_updates": {
            "block": {
                "max_bytes": "500000",
                "max_gas": "81500000"
            },
            "evidence": {
                "max_age_num_blocks": "403200",
                "max_age_duration": "2419200000000000",
                "max_bytes": "150000"
            },
            "validator": {
                "pub_key_types": [
                    "ed25519"
                ]
            }
        },
        "app_hash": "OEiYz/PCehPPM7G3rIdVJSBDkGNQUnD/0r/a4ydxt+w="
    }
}`
