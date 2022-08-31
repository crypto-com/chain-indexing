package usecase_parser_test

const TX_MSG_EXEC_MSG_WITHDRAW_DELEGATOR_REWARD_BLOCK_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "block_id": {
            "hash": "ABADA3C9E4E43B019EEE4CD2A7FD9EFEA0F8968C9D62915615E0C20F3D8CBEFC",
            "parts": {
                "total": 1,
                "hash": "FEB80A021E4361105437B636AC163F2F3D6E2312E194F06C90F44384B6D80610"
            }
        },
        "block": {
            "header": {
                "version": {
                    "block": "11"
                },
                "chain_id": "chaintest",
                "height": "313",
                "time": "2022-08-11T09:10:07.111359Z",
                "last_block_id": {
                    "hash": "7D37C7B0D7951DF33C9FB8325FBD3195273283844F58DD635EE6670CCEF07699",
                    "parts": {
                        "total": 1,
                        "hash": "E7B1758C470956AB30ADD1EEF6D05628FF53F9CC2A609F69B620634FCDA2F31B"
                    }
                },
                "last_commit_hash": "978B717F13BF769FB299BE09244B5E954627E41F6E0BF57CADD57C1728BA5285",
                "data_hash": "6EDB8920F3922910E0351E4B4CBA40DFBA493584608BF6B06F5C7DF9C34C3701",
                "validators_hash": "3CC6A79299D004B413B5701492B3CC3BCB0A44DC245AB40E92E4D3A7F38D44C9",
                "next_validators_hash": "3CC6A79299D004B413B5701492B3CC3BCB0A44DC245AB40E92E4D3A7F38D44C9",
                "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
                "app_hash": "56F6964C4476F01DEC0E546A701302896469017E1BF46E60AD39D96ADD63C0DB",
                "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "proposer_address": "1BA1E260F60A3A76B51CD36CB92026402D49A031"
            },
            "data": {
                "txs": [
                    "CoUDCoIDCh0vY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnRXhlYxLgAgoqY3JvMWh0cXN4Zmo0azloaGFndHZsbXF4Nmw0ajU5M3B6ZGs3ZGR2NTBuEpcBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkElwKKmNybzFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN2RkdjUwbhIuY3JvY25jbDFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN3dxMGFkMBKXAQo3L2Nvc21vcy5kaXN0cmlidXRpb24udjFiZXRhMS5Nc2dXaXRoZHJhd0RlbGVnYXRvclJld2FyZBJcCipjcm8xaHRxc3hmajRrOWhoYWd0dmxtcXg2bDRqNTkzcHpkazdkZHY1MG4SLmNyb2NuY2wxN2N5ZTVnbWVhbG1wZ3ozdGNsZmpnYTd1cmRqeGo2bmFoM3U4dzISWApQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAzwKY/mjHju6d7GSvY/oTu62ILJwQ+ZdCiAOadX5F366EgQKAggBGAMSBBDAmgwaQCKdcfyXIPQOrVuiaZZXswnVj5NvAQLYTU33iIH3kAVoMQ7QC03X0FjeMvVLhwzt5cVzjaSdGXERFRXZ8N4sc2Y="
                ]
            },
            "evidence": {
                "evidence": []
            },
            "last_commit": {
                "height": "312",
                "round": 0,
                "block_id": {
                    "hash": "7D37C7B0D7951DF33C9FB8325FBD3195273283844F58DD635EE6670CCEF07699",
                    "parts": {
                        "total": 1,
                        "hash": "E7B1758C470956AB30ADD1EEF6D05628FF53F9CC2A609F69B620634FCDA2F31B"
                    }
                },
                "signatures": [
                    {
                        "block_id_flag": 2,
                        "validator_address": "1BA1E260F60A3A76B51CD36CB92026402D49A031",
                        "timestamp": "2022-08-11T09:10:07.111359Z",
                        "signature": "PQCoHaYCXLYgLoYYLAy49SL6BvGLvYV5wwdvvfUqNzCAY85fPvQoTOnTrg31Bo6QvXpv5P8aIXuzzFWZAJJHBw=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "D1483ACF4613CF4450BC0FCEB07F84FE2634A5B1",
                        "timestamp": "2022-08-11T09:10:07.129718Z",
                        "signature": "SrzBacorSQbUnovRRcI1R3rbYyRgvh2lUks3FCX5McyDZvRvzLXgd14fNDUAdAMK0mrGGN96M8eOAG9QZXQMBQ=="
                    }
                ]
            }
        }
    }
}
`

const TX_MSG_EXEC_MSG_WITHDRAW_DELEGATOR_REWARD_BLOCK_RESULTS_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "height": "313",
        "txs_results": [
            {
                "code": 0,
                "data": "CiUKHS9jb3Ntb3MuYXV0aHoudjFiZXRhMS5Nc2dFeGVjEgQKAAoA",
                "log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"amount\",\"value\":\"362332basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"amount\",\"value\":\"362332basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.authz.v1beta1.MsgExec\"},{\"key\":\"sender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"sender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"amount\",\"value\":\"362332basecro\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"362332basecro\"},{\"key\":\"validator\",\"value\":\"crocncl1htqsxfj4k9hhagtvlmqx6l4j593pzdk7wq0ad0\"},{\"key\":\"amount\"},{\"key\":\"validator\",\"value\":\"crocncl17cye5gmealmpgz3tclfjga7urdjxj6nah3u8w2\"}]}]}]",
                "info": "",
                "gas_wanted": "200000",
                "gas_used": "153624",
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
                                "value": "SXAxeC9KY2c5QTZ0VzZKcGxsZXpDZFdQazI4QkF0aE5UZmVJZ2ZlUUJXZ3hEdEFMVGRmUVdONHk5VXVIRE8zbHhYT05wSjBaY1JFVkZkbnczaXh6Wmc9PQ==",
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
                                "value": "MzYyMzMyYmFzZWNybw==",
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
                                "value": "MzYyMzMyYmFzZWNybw==",
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
                                "value": "MzYyMzMyYmFzZWNybw==",
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
                                "value": "MzYyMzMyYmFzZWNybw==",
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
                        "value": "MC4wMDEyMDMzNjE3MDE2OTczMzk=",
                        "index": true
                    },
                    {
                        "key": "aW5mbGF0aW9u",
                        "value": "MC4xMzAwMDY0MzUzNjIzMjA0MjM=",
                        "index": true
                    },
                    {
                        "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
                        "value": "MjE2MDcyMDg0MTM1LjU1MDg5NzU1MTg4NjkwNzk0NA==",
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
