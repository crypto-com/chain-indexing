package usecase_parser_test

const TX_MSG_EXEC_MSG_WITHDRAW_VALIDATOR_COMMISSION_BLOCK_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "block_id": {
            "hash": "94727699F5F489DEE28C94D6B0E3C8B503EA742B5307B187C30DF233FD6E334B",
            "parts": {
                "total": 1,
                "hash": "D22B5D67F5F71D8996DAD52D74F1064314A02F44F64C9CD1950FE05F089AB92F"
            }
        },
        "block": {
            "header": {
                "version": {
                    "block": "11"
                },
                "chain_id": "chaintest",
                "height": "5991",
                "time": "2022-08-12T02:59:57.09639Z",
                "last_block_id": {
                    "hash": "632C790AB185E8D5C6D3129579D49D56BCED17E1D47C9CFA88A4BE295D231148",
                    "parts": {
                        "total": 1,
                        "hash": "18D6C5A6872B6C62D504C5FBCCE53812EA911AC56CDDAC12DFCD05E165611CC6"
                    }
                },
                "last_commit_hash": "D4570C1D0AFFE51426301B0B4CB6907B03074827CCB1BA5B24ED4B868DC746FC",
                "data_hash": "43B7683878C0235CEE8F88DE67660ECE8B940FCDEA4F3AA3ED88DF6BC2E762E9",
                "validators_hash": "9797875DD992FE368DA8CB7A1A6C04F8F09B5F667332108571700DBDE676B8D8",
                "next_validators_hash": "9797875DD992FE368DA8CB7A1A6C04F8F09B5F667332108571700DBDE676B8D8",
                "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
                "app_hash": "A69E1D8D43AF188CD99FC9E6031D85BEDF30F26A66EB9CFFCB6616150DDA67CB",
                "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "proposer_address": "42404550EC84BDC92FE4A309740FE93DA1E98A1B"
            },
            "data": {
                "txs": [
                    "CtwCCtkCCh0vY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnRXhlYxK3AgoqY3JvMTQwNnkwMDluNDNhd2EwZTVuNnQydDA2YTh0dHk5YXpwdTlhdDI1EpcBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkElwKKmNybzFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN2RkdjUwbhIuY3JvY25jbDFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN3dxMGFkMBJvCjsvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3VmFsaWRhdG9yQ29tbWlzc2lvbhIwCi5jcm9jbmNsMWh0cXN4Zmo0azloaGFndHZsbXF4Nmw0ajU5M3B6ZGs3d3EwYWQwElgKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQNrfGTd70aQIuyCXaBOVHUNvHddYe0krKAWqPLnwNefrhIECgIIARgQEgQQwJoMGkDMONvP4mTQflkVF6GtC7TvzPHDSBwt9OoYzYgd78EAPzvsQb0HTzf7JNTcFltR/900Mo1c5k3FDWkIFxK9ztBY"
                ]
            },
            "evidence": {
                "evidence": []
            },
            "last_commit": {
                "height": "5990",
                "round": 0,
                "block_id": {
                    "hash": "632C790AB185E8D5C6D3129579D49D56BCED17E1D47C9CFA88A4BE295D231148",
                    "parts": {
                        "total": 1,
                        "hash": "18D6C5A6872B6C62D504C5FBCCE53812EA911AC56CDDAC12DFCD05E165611CC6"
                    }
                },
                "signatures": [
                    {
                        "block_id_flag": 2,
                        "validator_address": "42404550EC84BDC92FE4A309740FE93DA1E98A1B",
                        "timestamp": "2022-08-12T02:59:57.09639Z",
                        "signature": "yNz6LZdB/uvsn6JmMWl3dvI5Lz+OeiOc7Vu98pL6vnbMCIIFpjmKh/NGIAyDf4bFdPNGi4i8us7s39XO2LfUBg=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "FB8249644B36E76759F42EAE764F5842E779D1EE",
                        "timestamp": "2022-08-12T02:59:57.180172Z",
                        "signature": "cnpZUOaMD92bFykK0ISQriFiAK6iMVCeK1YFpM7B1NR0X6oxlXeGtleDyxQHkBi1Kk0KqomxliTteEsArYW5Bg=="
                    }
                ]
            }
        }
    }
}
`

const TX_MSG_EXEC_MSG_WITHDRAW_VALIDATOR_COMMISSION_BLOCK_RESULTS_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "height": "5991",
        "txs_results": [
            {
                "code": 0,
                "data": "CiUKHS9jb3Ntb3MuYXV0aHoudjFiZXRhMS5Nc2dFeGVjEgQKAAoA",
                "log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"amount\",\"value\":\"1285329basecro\"},{\"key\":\"receiver\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"amount\",\"value\":\"142814basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"amount\",\"value\":\"1285329basecro\"},{\"key\":\"spender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"amount\",\"value\":\"142814basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.authz.v1beta1.MsgExec\"},{\"key\":\"sender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"sender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"crocncl1htqsxfj4k9hhagtvlmqx6l4j593pzdk7wq0ad0\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"sender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"amount\",\"value\":\"1285329basecro\"},{\"key\":\"recipient\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"sender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"amount\",\"value\":\"142814basecro\"}]},{\"type\":\"withdraw_commission\",\"attributes\":[{\"key\":\"amount\",\"value\":\"142814basecro\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"1285329basecro\"},{\"key\":\"validator\",\"value\":\"crocncl1htqsxfj4k9hhagtvlmqx6l4j593pzdk7wq0ad0\"}]}]}]",
                "info": "",
                "gas_wanted": "200000",
                "gas_used": "130101",
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
                                "value": "Y3JvMTQwNnkwMDluNDNhd2EwZTVuNnQydDA2YTh0dHk5YXpwdTlhdDI1LzE2",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "c2lnbmF0dXJl",
                                "value": "ekRqYnorSmswSDVaRlJlaHJRdTA3OHp4dzBnY0xmVHFHTTJJSGUvQkFEODc3RUc5QjA4Myt5VFUzQlpiVWYvZE5ES05YT1pOeFExcENCY1N2YzdRV0E9PQ==",
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
                        "type": "coin_spent",
                        "attributes": [
                            {
                                "key": "c3BlbmRlcg==",
                                "value": "Y3JvMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4bHl2OTR3",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MTI4NTMyOWJhc2Vjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "coin_received",
                        "attributes": [
                            {
                                "key": "cmVjZWl2ZXI=",
                                "value": "Y3JvMWh0cXN4Zmo0azloaGFndHZsbXF4Nmw0ajU5M3B6ZGs3ZGR2NTBu",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MTI4NTMyOWJhc2Vjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "Y3JvMWh0cXN4Zmo0azloaGFndHZsbXF4Nmw0ajU5M3B6ZGs3ZGR2NTBu",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4bHl2OTR3",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MTI4NTMyOWJhc2Vjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4bHl2OTR3",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "withdraw_rewards",
                        "attributes": [
                            {
                                "key": "YW1vdW50",
                                "value": "MTI4NTMyOWJhc2Vjcm8=",
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
                    },
                    {
                        "type": "coin_spent",
                        "attributes": [
                            {
                                "key": "c3BlbmRlcg==",
                                "value": "Y3JvMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4bHl2OTR3",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MTQyODE0YmFzZWNybw==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "coin_received",
                        "attributes": [
                            {
                                "key": "cmVjZWl2ZXI=",
                                "value": "Y3JvMWh0cXN4Zmo0azloaGFndHZsbXF4Nmw0ajU5M3B6ZGs3ZGR2NTBu",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MTQyODE0YmFzZWNybw==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "Y3JvMWh0cXN4Zmo0azloaGFndHZsbXF4Nmw0ajU5M3B6ZGs3ZGR2NTBu",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4bHl2OTR3",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MTQyODE0YmFzZWNybw==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "c2VuZGVy",
                                "value": "Y3JvMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4bHl2OTR3",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "withdraw_commission",
                        "attributes": [
                            {
                                "key": "YW1vdW50",
                                "value": "MTQyODE0YmFzZWNybw==",
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
                                "value": "Y3JvY25jbDFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN3dxMGFkMA==",
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
                        "value": "MzQyNjliYXNlY3Jv",
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
                        "value": "MzQyNjliYXNlY3Jv",
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
                        "value": "MzQyNjliYXNlY3Jv",
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
                        "value": "MzQyNjliYXNlY3Jv",
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
                        "value": "MzQyNjliYXNlY3Jv",
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
                        "value": "MC4wMDEyMDMyMjA5MDY2MTk1NzU=",
                        "index": true
                    },
                    {
                        "key": "aW5mbGF0aW9u",
                        "value": "MC4xMzAxMjMxNzY1NDg2MjM5MDU=",
                        "index": true
                    },
                    {
                        "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
                        "value": "MjE2MjkxNDE1NTM3LjY1NDI5NDA1MTM5NTY4ODY1MA==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyNjk=",
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
                        "value": "MzQyNjliYXNlY3Jv",
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
                        "value": "MzQyNjliYXNlY3Jv",
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
                        "value": "MzQyNjliYXNlY3Jv",
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
                        "value": "MTcxMy40NTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": "MTcxLjM0NTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
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
                        "value": "MTcxMy40NTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": "MTU5My41MDg1MDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                        "value": "MTU5MzUuMDg1MDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
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
                        "value": "MTU5My41MDg1MDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                        "value": "MTU5MzUuMDg1MDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
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
