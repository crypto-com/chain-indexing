package usecase_parser_test

const TX_MSG_EXEC_MSG_SET_WITHDRAW_ADDRESS_BLOCK_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "block_id": {
            "hash": "03263B69D5AAF9BC9952477DF72BAAF2B578E0CCC5E184103148B13B69B3389D",
            "parts": {
                "total": 1,
                "hash": "6501FF7A2E2B1270188B6100914734A10FAC6475FFB98562BE2D37F74D29D5A7"
            }
        },
        "block": {
            "header": {
                "version": {
                    "block": "11"
                },
                "chain_id": "chaintest",
                "height": "292",
                "time": "2022-08-11T08:26:54.360681Z",
                "last_block_id": {
                    "hash": "84BB867F07C7A6EC81B20599F2F20E9A795257335B120CEBECB0BB5FDF5E3CF2",
                    "parts": {
                        "total": 1,
                        "hash": "007061C1BC5B26B8D6A99D7C3E4133F6B5D1DBA637FEB95CEFEDDA37180BD4E1"
                    }
                },
                "last_commit_hash": "66BFF84634FEAD8DB8AB76CEEDC65919513B935DC2E1D3BA1F57DFC8ABE54302",
                "data_hash": "677F2B86E71F84545D73EBF8449794CB04895411B2EFF78EB31F4EAE2AFADE55",
                "validators_hash": "D179F82943E9BA4C392F1A3A75F75670098576F035DC973B1FBC21F816A84CC5",
                "next_validators_hash": "D179F82943E9BA4C392F1A3A75F75670098576F035DC973B1FBC21F816A84CC5",
                "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
                "app_hash": "29B03993F27E581FC1C0650B2E3E264342574C44612F5FB9458F117B1797EF34",
                "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "proposer_address": "F288243EE77CDC90BF8BBDA4F3D75163A8C8E589"
            },
            "data": {
                "txs": [
                    "CuIBCt8BCh0vY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnRXhlYxK9AQoqY3JvMTQwNnkwMDluNDNhd2EwZTVuNnQydDA2YTh0dHk5YXpwdTlhdDI1Eo4BCjIvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1NldFdpdGhkcmF3QWRkcmVzcxJYCipjcm8xaHRxc3hmajRrOWhoYWd0dmxtcXg2bDRqNTkzcHpkazdkZHY1MG4SKmNybzE0MDZ5MDA5bjQzYXdhMGU1bjZ0MnQwNmE4dHR5OWF6cHU5YXQyNRJYClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDa3xk3e9GkCLsgl2gTlR1Dbx3XWHtJKygFqjy58DXn64SBAoCCAEYAxIEEMCaDBpALNJps5yQJ8AJXnf53XJcrN2DbXSbWMxE/e82rhm0as8S/vTIWZUCVIaDxSCDKt09mlYRwwDSAZOfW6XGh/6p5Q=="
                ]
            },
            "evidence": {
                "evidence": []
            },
            "last_commit": {
                "height": "291",
                "round": 0,
                "block_id": {
                    "hash": "84BB867F07C7A6EC81B20599F2F20E9A795257335B120CEBECB0BB5FDF5E3CF2",
                    "parts": {
                        "total": 1,
                        "hash": "007061C1BC5B26B8D6A99D7C3E4133F6B5D1DBA637FEB95CEFEDDA37180BD4E1"
                    }
                },
                "signatures": [
                    {
                        "block_id_flag": 2,
                        "validator_address": "0737362FA0CB60BBEC339BBA31F25FE6BD363ED0",
                        "timestamp": "2022-08-11T08:26:54.360681Z",
                        "signature": "xDkTXjfirZahCf0c5a4ls3qhgVheQKsAQdKMwlT7HoaLSCcgctFaeI+Z3zcqGKXAn0yTIbVSmjIJ4xRFDa1GBQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "F288243EE77CDC90BF8BBDA4F3D75163A8C8E589",
                        "timestamp": "2022-08-11T08:26:54.3967Z",
                        "signature": "r2AgkWmmMr4X0FT1f0/7k2aqX9Sp5dGN4zBO5LTYeAQ9ca0G2HNO7PNdSISU0we7WMi1XxHSiWdvV/G51U88Bw=="
                    }
                ]
            }
        }
    }
}
`

const TX_MSG_EXEC_MSG_SET_WITHDRAW_ADDRESS_BLOCK_RESULTS_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "height": "292",
        "txs_results": [
            {
                "code": 0,
                "data": "CiMKHS9jb3Ntb3MuYXV0aHoudjFiZXRhMS5Nc2dFeGVjEgIKAA==",
                "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.authz.v1beta1.MsgExec\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"}]},{\"type\":\"set_withdraw_address\",\"attributes\":[{\"key\":\"withdraw_address\",\"value\":\"cro1406y009n43awa0e5n6t2t06a8tty9azpu9at25\"}]}]}]",
                "info": "",
                "gas_wanted": "200000",
                "gas_used": "47394",
                "events": [
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "ZmVl",
                                "value": null,
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "YWNjX3NlcQ==",
                                "value": "Y3JvMTQwNnkwMDluNDNhd2EwZTVuNnQydDA2YTh0dHk5YXpwdTlhdDI1LzM=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "c2lnbmF0dXJl",
                                "value": "TE5KcHM1eVFKOEFKWG5mNTNYSmNyTjJEYlhTYldNeEUvZTgycmhtMGFzOFMvdlRJV1pVQ1ZJYUR4U0NES3QwOW1sWVJ3d0RTQVpPZlc2WEdoLzZwNVE9PQ==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "YWN0aW9u",
                                "value": "L2Nvc21vcy5hdXRoei52MWJldGExLk1zZ0V4ZWM=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "set_withdraw_address",
                        "attributes": [
                            {
                                "key": "d2l0aGRyYXdfYWRkcmVzcw==",
                                "value": "Y3JvMTQwNnkwMDluNDNhd2EwZTVuNnQydDA2YTh0dHk5YXpwdTlhdDI1",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "bW9kdWxl",
                                "value": "ZGlzdHJpYnV0aW9u",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMWh0cXN4Zmo0azloaGFndHZsbXF4Nmw0ajU5M3B6ZGs3ZGR2NTBu",
                                "index": true
                            }
                        ]
                    }
                ],
                "codespace": ""
            }
        ],
        "begin_block_events": [
            {
                "type": "coin_received",
                "attributes": [
                    {
                        "key": "cmVjZWl2ZXI=",
                        "value": "Y3JvMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04czIwcG0z",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzRiYXNlY3Jv",
                        "index": true
                    }
                ]
            },
            {
                "type": "coinbase",
                "attributes": [
                    {
                        "key": "bWludGVy",
                        "value": "Y3JvMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04czIwcG0z",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzRiYXNlY3Jv",
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_spent",
                "attributes": [
                    {
                        "key": "c3BlbmRlcg==",
                        "value": "Y3JvMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04czIwcG0z",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzRiYXNlY3Jv",
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_received",
                "attributes": [
                    {
                        "key": "cmVjZWl2ZXI=",
                        "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzRiYXNlY3Jv",
                        "index": true
                    }
                ]
            },
            {
                "type": "transfer",
                "attributes": [
                    {
                        "key": "cmVjaXBpZW50",
                        "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
                        "index": true
                    },
                    {
                        "key": "c2VuZGVy",
                        "value": "Y3JvMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04czIwcG0z",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzRiYXNlY3Jv",
                        "index": true
                    }
                ]
            },
            {
                "type": "message",
                "attributes": [
                    {
                        "key": "c2VuZGVy",
                        "value": "Y3JvMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04czIwcG0z",
                        "index": true
                    }
                ]
            },
            {
                "type": "mint",
                "attributes": [
                    {
                        "key": "Ym9uZGVkX3JhdGlv",
                        "value": "MC4wMDEyMDMzNjIyMjE2MTgxNTU=",
                        "index": true
                    },
                    {
                        "key": "aW5mbGF0aW9u",
                        "value": "MC4xMzAwMDYwMDM1OTY3OTQ5OTQ=",
                        "index": true
                    },
                    {
                        "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
                        "value": "MjE2MDcxMjczMDczLjQ5OTk5NDYzMDcyNjM1OTExNg==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzQ=",
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_spent",
                "attributes": [
                    {
                        "key": "c3BlbmRlcg==",
                        "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzRiYXNlY3Jv",
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_received",
                "attributes": [
                    {
                        "key": "cmVjZWl2ZXI=",
                        "value": "Y3JvMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4bHl2OTR3",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzRiYXNlY3Jv",
                        "index": true
                    }
                ]
            },
            {
                "type": "transfer",
                "attributes": [
                    {
                        "key": "cmVjaXBpZW50",
                        "value": "Y3JvMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4bHl2OTR3",
                        "index": true
                    },
                    {
                        "key": "c2VuZGVy",
                        "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzRiYXNlY3Jv",
                        "index": true
                    }
                ]
            },
            {
                "type": "message",
                "attributes": [
                    {
                        "key": "c2VuZGVy",
                        "value": "Y3JvMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsZ3p0ZWh2",
                        "index": true
                    }
                ]
            },
            {
                "type": "proposer_reward",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": "MTcxMS43MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "Y3JvY25jbDE3Y3llNWdtZWFsbXBnejN0Y2xmamdhN3VyZGp4ajZuYWgzdTh3Mg==",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": "MTcxLjE3MDAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "Y3JvY25jbDE3Y3llNWdtZWFsbXBnejN0Y2xmamdhN3VyZGp4ajZuYWgzdTh3Mg==",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": "MTcxMS43MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "Y3JvY25jbDE3Y3llNWdtZWFsbXBnejN0Y2xmamdhN3VyZGp4ajZuYWgzdTh3Mg==",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": "MTU5MS44ODEwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "Y3JvY25jbDE3Y3llNWdtZWFsbXBnejN0Y2xmamdhN3VyZGp4ajZuYWgzdTh3Mg==",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": "MTU5MTguODEwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "Y3JvY25jbDE3Y3llNWdtZWFsbXBnejN0Y2xmamdhN3VyZGp4ajZuYWgzdTh3Mg==",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": "MTU5MS44ODEwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "Y3JvY25jbDFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN3dxMGFkMA==",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": "MTU5MTguODEwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "Y3JvY25jbDFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN3dxMGFkMA==",
                        "index": true
                    }
                ]
            }
        ],
        "end_block_events": null,
        "validator_updates": null,
        "consensus_param_updates": {
            "block": {
                "max_bytes": "22020096",
                "max_gas": "-1"
            },
            "evidence": {
                "max_age_num_blocks": "100000",
                "max_age_duration": "172800000000000",
                "max_bytes": "1048576"
            },
            "validator": {
                "pub_key_types": [
                    "ed25519"
                ]
            }
        }
    }
}
`
