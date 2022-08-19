package usecase_parser_test

const TX_MSG_EXEC_MULT_WITHDRAW_DELEGATOR_REWARD_BLOCK_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "block_id": {
            "hash": "5C24ADF82FDDA56090BA29B2BCFC2599A21609DE0C04E5DC141B1A3698DBD010",
            "parts": {
                "total": 1,
                "hash": "773673E2B23A74F10012F2718BFAA4A88A26B89DB9D8B472A5F8ACADF455A61C"
            }
        },
        "block": {
            "header": {
                "version": {
                    "block": "11"
                },
                "chain_id": "chaintest",
                "height": "386",
                "time": "2022-08-19T01:59:01.291819Z",
                "last_block_id": {
                    "hash": "C57A8A83A77F2B5C131ABD61F6FCAD2610EB6A6D816CA0BC9C04940DA9FC6775",
                    "parts": {
                        "total": 1,
                        "hash": "C6A5AC9B3878B94C1DEB6FDCF882B2E0E2EC21CEB0958CE57251E05EDA638EDD"
                    }
                },
                "last_commit_hash": "F8DEC4F3397E31FD251B2BB4D9808C47B79A48414E27B0666AB97FB409532041",
                "data_hash": "E9D6612EAB1FC07D4EC3DBCE998796A7BA90115AAAA189FFC38A843AB1C887C8",
                "validators_hash": "61D1EB1DA21B8A0FED97669B5E3CF6BA7EEDE24F0F5D9930F38FE472BF5638F4",
                "next_validators_hash": "61D1EB1DA21B8A0FED97669B5E3CF6BA7EEDE24F0F5D9930F38FE472BF5638F4",
                "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
                "app_hash": "988481D0D12CE8DBDC29AE46CD4B0882C1124BB8D356CA49441709CC34E4C444",
                "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "proposer_address": "7BB7616678951F7A91230C2FDF5DAE8B09E5BB1E"
            },
            "data": {
                "txs": [
                    "CtYDCugBCh0vY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnRXhlYxLGAQoqY3JvMTQwNnkwMDluNDNhd2EwZTVuNnQydDA2YTh0dHk5YXpwdTlhdDI1EpcBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkElwKKmNybzFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN2RkdjUwbhIuY3JvY25jbDFodHFzeGZqNGs5aGhhZ3R2bG1xeDZsNGo1OTNwemRrN3dxMGFkMAroAQodL2Nvc21vcy5hdXRoei52MWJldGExLk1zZ0V4ZWMSxgEKKmNybzE0MDZ5MDA5bjQzYXdhMGU1bjZ0MnQwNmE4dHR5OWF6cHU5YXQyNRKXAQo3L2Nvc21vcy5kaXN0cmlidXRpb24udjFiZXRhMS5Nc2dXaXRoZHJhd0RlbGVnYXRvclJld2FyZBJcCipjcm8xaHRxc3hmajRrOWhoYWd0dmxtcXg2bDRqNTkzcHpkazdkZHY1MG4SLmNyb2NuY2wxN2N5ZTVnbWVhbG1wZ3ozdGNsZmpnYTd1cmRqeGo2bmFoM3U4dzISWApQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohA2t8ZN3vRpAi7IJdoE5UdQ28d11h7SSsoBao8ufA15+uEgQKAggBGAMSBBDAmgwaQGgnvhOU9S+FKjkB9n3LEOTQR+F04JIvgIwSRIqZrQmyHAMyYpvl/Hgs07/jB7msvhJrcZxq5RS/C4KJSr+L4ZM="
                ]
            },
            "evidence": {
                "evidence": []
            },
            "last_commit": {
                "height": "385",
                "round": 0,
                "block_id": {
                    "hash": "C57A8A83A77F2B5C131ABD61F6FCAD2610EB6A6D816CA0BC9C04940DA9FC6775",
                    "parts": {
                        "total": 1,
                        "hash": "C6A5AC9B3878B94C1DEB6FDCF882B2E0E2EC21CEB0958CE57251E05EDA638EDD"
                    }
                },
                "signatures": [
                    {
                        "block_id_flag": 2,
                        "validator_address": "4974ACEAB292467EBD9B5992E6330079F79BB1A7",
                        "timestamp": "2022-08-19T01:59:01.291819Z",
                        "signature": "2J/pttkZwNoakuWyux9mlfBh+Z3qK7Mv2F35kaIujYjLmVvcdWOMx62MdAvwvdQ4SaEm/J63QDqypnXSGYtrAQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "7BB7616678951F7A91230C2FDF5DAE8B09E5BB1E",
                        "timestamp": "2022-08-19T01:59:01.342858Z",
                        "signature": "XHFbCn7wvequM6Wpn8+5tLUM7rid0cNXW4Rz4iLXyyGj617SmFNo0s9YRS6qBffeZjJLZEVBxV5HqzVkJmiwAg=="
                    }
                ]
            }
        }
    }
}
`

const TX_MSG_EXEC_MULT_WITHDRAW_DELEGATOR_REWARD_RESULTS_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "height": "386",
        "txs_results": [
            {
                "code": 0,
                "data": "CiMKHS9jb3Ntb3MuYXV0aHoudjFiZXRhMS5Nc2dFeGVjEgIKAAojCh0vY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnRXhlYxICCgA=",
                "log": "[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"amount\",\"value\":\"513305basecro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"amount\",\"value\":\"513305basecro\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.authz.v1beta1.MsgExec\"},{\"key\":\"sender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"},{\"key\":\"sender\",\"value\":\"cro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8lyv94w\"},{\"key\":\"amount\",\"value\":\"513305basecro\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"513305basecro\"},{\"key\":\"validator\",\"value\":\"crocncl1htqsxfj4k9hhagtvlmqx6l4j593pzdk7wq0ad0\"}]}]},{\"msg_index\":1,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/cosmos.authz.v1beta1.MsgExec\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"cro1htqsxfj4k9hhagtvlmqx6l4j593pzdk7ddv50n\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\"},{\"key\":\"validator\",\"value\":\"crocncl17cye5gmealmpgz3tclfjga7urdjxj6nah3u8w2\"}]}]}]",
                "info": "",
                "gas_wanted": "200000",
                "gas_used": "157871",
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
                                "value": "YUNlK0U1VDFMNFVxT1FIMmZjc1E1TkJINFhUZ2tpK0FqQkpFaXBtdENiSWNBekppbStYOGVDelR2K01IdWF5K0VtdHhuR3JsRkw4TGdvbEt2NHZoa3c9PQ==",
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
                                "value": "NTEzMzA1YmFzZWNybw==",
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
                                "value": "NTEzMzA1YmFzZWNybw==",
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
                                "value": "NTEzMzA1YmFzZWNybw==",
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
                                "value": "NTEzMzA1YmFzZWNybw==",
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
                        "value": "MzQyMzViYXNlY3Jv",
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
                        "value": "MzQyMzViYXNlY3Jv",
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
                        "value": "MzQyMzViYXNlY3Jv",
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
                        "value": "MzQyMzViYXNlY3Jv",
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
                        "value": "MzQyMzViYXNlY3Jv",
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
                        "value": "MC4wMDEyMDMzNTk5NTc4NDg2MTI=",
                        "index": true
                    },
                    {
                        "key": "aW5mbGF0aW9u",
                        "value": "MC4xMzAwMDc5MzYyNjE1MzA0NzI=",
                        "index": true
                    },
                    {
                        "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
                        "value": "MjE2MDc0OTAzNTQ2LjU2MjA2MzU0NTkzMDE3MDMyMA==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MzQyMzU=",
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
                        "value": "MzQyMzViYXNlY3Jv",
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
                        "value": "MzQyMzViYXNlY3Jv",
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
                        "value": "MzQyMzViYXNlY3Jv",
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
                        "value": "MTcxMS43NTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                        "value": "MTcxLjE3NTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
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
                        "value": "MTcxMS43NTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                        "value": "MTU5MS45Mjc1MDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                        "value": "MTU5MTkuMjc1MDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
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
                        "value": "MTU5MS45Mjc1MDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
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
                        "value": "MTU5MTkuMjc1MDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
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
