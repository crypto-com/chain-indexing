package usecase_parser_test

const TX_MSG_EXEC_NESTED_MSG_EXEC_BLOCK_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "block_id": {
            "hash": "6D58B0914341804F879F343848CC74080086BF9443C8D8CD2DFE8E0CDB839247",
            "parts": {
                "total": 1,
                "hash": "B83125E1E3F48D779C364F7120B1B31A5B6775C690A2C3FEC846B8D6418193AD"
            }
        },
        "block": {
            "header": {
                "version": {
                    "block": "11"
                },
                "chain_id": "chaintest",
                "height": "58",
                "time": "2022-08-19T09:30:27.495942Z",
                "last_block_id": {
                    "hash": "35DDC6635CF00EBABDB3E9376C7B26A42F29E79DAD20B445A28ADE88FE9D9DA6",
                    "parts": {
                        "total": 1,
                        "hash": "2AADC36805CE825FEDBE39DBF98114CEDAC4C4281A4CE71ECB25394085AABDF2"
                    }
                },
                "last_commit_hash": "461EFC6566A7766A1BEE3A6BD94A4E6D7EC57FC02F3923AE22FB7EDE692568B2",
                "data_hash": "F45088B06A9EA1CBCA8AD3A2E78EDA11270607A1BA52F56002A7F4271A45162A",
                "validators_hash": "DB2100932BDA8EC262F0DAD4B9FEB5988D324963E8BB215B0EF24ADD7941CD9C",
                "next_validators_hash": "DB2100932BDA8EC262F0DAD4B9FEB5988D324963E8BB215B0EF24ADD7941CD9C",
                "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
                "app_hash": "54DE68D1DB8FC8163014835704BF74CF8128ED3270742C914E93DFF8E552C69A",
                "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "proposer_address": "CBE0426A5F46596EC0976A56793909C56C0D9A20"
            },
            "data": {
                "txs": [
                    "CtYDCtMDCh0vY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnRXhlYxKxAwoqY3JvMWh0cXN4Zmo0azloaGFndHZsbXF4Nmw0ajU5M3B6ZGs3ZGR2NTBuEoIDCh0vY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnRXhlYxLgAgoqY3JvMWh0cXN4Zmo0azloaGFndHZsbXF4Nmw0ajU5M3B6ZGs3ZGR2NTBuEpcBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkElwKKmNybzFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN2RkdjUwbhIuY3JvY25jbDFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN3dxMGFkMBKXAQo3L2Nvc21vcy5kaXN0cmlidXRpb24udjFiZXRhMS5Nc2dXaXRoZHJhd0RlbGVnYXRvclJld2FyZBJcCipjcm8xaHRxc3hmajRrOWhoYWd0dmxtcXg2bDRqNTkzcHpkazdkZHY1MG4SLmNyb2NuY2wxN2N5ZTVnbWVhbG1wZ3ozdGNsZmpnYTd1cmRqeGo2bmFoM3U4dzISWApQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAzwKY/mjHju6d7GSvY/oTu62ILJwQ+ZdCiAOadX5F366EgQKAggBGAMSBBDAmgwaQKAXbycuiijnQnh6LuiaNusBAAKgbMksOw7VUqhgGrMIETRiLBi8sqFnz8ckWBCfxNPmQ8IVAyEOSo1jE1GsNak="
                ]
            },
            "evidence": {
                "evidence": []
            },
            "last_commit": {
                "height": "57",
                "round": 0,
                "block_id": {
                    "hash": "35DDC6635CF00EBABDB3E9376C7B26A42F29E79DAD20B445A28ADE88FE9D9DA6",
                    "parts": {
                        "total": 1,
                        "hash": "2AADC36805CE825FEDBE39DBF98114CEDAC4C4281A4CE71ECB25394085AABDF2"
                    }
                },
                "signatures": [
                    {
                        "block_id_flag": 2,
                        "validator_address": "AE52FFF6A016B5149347EF576D2892923AB250FC",
                        "timestamp": "2022-08-19T09:30:27.495942Z",
                        "signature": "+9xpBFQIMnW3cbh4wqwe41fUvLB+zYgai6nFXDqb6UQGqqNxAAld5fRUPzSQ8q46T8pUgSG22JjiAcOtzbsMDw=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "CBE0426A5F46596EC0976A56793909C56C0D9A20",
                        "timestamp": "2022-08-19T09:30:27.513927Z",
                        "signature": "3tsHOlBH+2QsKEvIOK7keL99FNWqf4cerZsyrx6EcW5L9iESUVbUSZYRrbuldabLR+KnxY+xkDaQWjP4VEpGDQ=="
                    }
                ]
            }
        }
    }
}
`

const TX_MSG_EXEC_NESTED_MSG_EXEC_RESULTS_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "height": "58",
        "txs_results": [
            {
                "code": 0,
                "data": "CicKHS9jb3Ntb3MuYXV0aHoudjFiZXRhMS5Nc2dFeGVjEgYKBAoACgA=",
                "log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"amount\",\"value\":\"874045basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"amount\",\"value\":\"874045basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.authz.v1beta1.MsgExec\"},{\"key\":\"sender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"sender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"amount\",\"value\":\"874045basecro\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"874045basecro\"},{\"key\":\"validator\",\"value\":\"crocncl1htqsxfj4k9hhagtvlmqx6l4j593pzdk7wq0ad0\"},{\"key\":\"amount\"},{\"key\":\"validator\",\"value\":\"crocncl17cye5gmealmpgz3tclfjga7urdjxj6nah3u8w2\"}]}]}]",
                "info": "",
                "gas_wanted": "200000",
                "gas_used": "152871",
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
                                "value": "Y3JvMWh0cXN4Zmo0azloaGFndHZsbXF4Nmw0ajU5M3B6ZGs3ZGR2NTBuLzM=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "c2lnbmF0dXJl",
                                "value": "b0Jkdkp5NktLT2RDZUhvdTZKbzI2d0VBQXFCc3lTdzdEdFZTcUdBYXN3Z1JOR0lzR0x5eW9XZlB4eVJZRUovRTArWkR3aFVESVE1S2pXTVRVYXcxcVE9PQ==",
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
                                "value": "ODc0MDQ1YmFzZWNybw==",
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
                                "value": "ODc0MDQ1YmFzZWNybw==",
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
                                "value": "ODc0MDQ1YmFzZWNybw==",
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
                                "value": "ODc0MDQ1YmFzZWNybw==",
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
                        "type": "withdraw_rewards",
                        "attributes": [
                            {
                                "key": "YW1vdW50",
                                "value": null,
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
                        "value": "MzQyMzJiYXNlY3Jv",
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
                        "value": "MzQyMzJiYXNlY3Jv",
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
                        "value": "MzQyMzJiYXNlY3Jv",
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
                        "value": "MzQyMzJiYXNlY3Jv",
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
                        "value": "MzQyMzJiYXNlY3Jv",
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
                        "value": "MC4wMDEyMDMzNjgwODcyMTk3OTk=",
                        "index": true
                    },
                    {
                        "key": "aW5mbGF0aW9u",
                        "value": "MC4xMzAwMDExOTI0OTUyNDg1Mjg=",
                        "index": true
                    },
                    {
                        "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
                        "value": "MjE2MDYyMjM1NTg4LjU0OTg3ODg4NDgxMzc5ODI3Mg==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzI=",
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
                        "value": "MzQyMzJiYXNlY3Jv",
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
                        "value": "MzQyMzJiYXNlY3Jv",
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
                        "value": "MzQyMzJiYXNlY3Jv",
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
                        "value": "MTcxMS42MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                        "value": "MTcxLjE2MDAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
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
                        "value": "MTcxMS42MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                        "value": "MTU5MS43ODgwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                        "value": "MTU5MTcuODgwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
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
                        "value": "MTU5MS43ODgwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                        "value": "MTU5MTcuODgwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
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
