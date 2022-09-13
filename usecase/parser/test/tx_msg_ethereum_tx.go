package usecase_parser_test

const TX_MSG_ETHEREUM_TX_BLOCK_RESP = `{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "block_id": {
            "hash": "970D335ED369B2A53A58CC03647723133B2BC00D3CC84B79A83D966CCB1F7C35",
            "parts": {
                "total": 1,
                "hash": "B5306CCC162EAFDFE3658FE3278A2A4E42DED019416B22C22B8D1B8DDAE2327A"
            }
        },
        "block": {
            "header": {
                "version": {
                    "block": "11"
                },
                "chain_id": "cronostestnet_338-3",
                "height": "83178",
                "time": "2021-10-18T19:52:20.259005853Z",
                "last_block_id": {
                    "hash": "AE5259DC2C2D115254285E258956F2F54E8EE269A13889F155A73A9A2114B390",
                    "parts": {
                        "total": 1,
                        "hash": "06031C4D8A5D64AEAECC4E4655CC8D2540DDA7152857914815FCE4363E4FEBC4"
                    }
                },
                "last_commit_hash": "AFA474D967FC70BA9D3F33DEBE2C0E147213DC70C11D4C8AF9F131C486C7AF8E",
                "data_hash": "E4FBD0A69662DCEE6B21EDA7F954C73F727896A91B9E244A4B9CBB13864B017C",
                "validators_hash": "E719BAD04CC344817536E818E5396E604131C85BA8BA0EB20BD7662818DA30F3",
                "next_validators_hash": "E719BAD04CC344817536E818E5396E604131C85BA8BA0EB20BD7662818DA30F3",
                "consensus_hash": "252FE7CF36DD1BB85DAFC47A08961DF0CFD8C027DEFA5E01E958BE121599DB9D",
                "app_hash": "6F7C5A1D0F108AD6248F9F31BD85C70F7AC6CB601512CFEAAA2065EAE9B6634F",
                "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "proposer_address": "0421821D46C46F86F1EAD79EEB31CFA88A5578CB"
            },
            "data": {
                "txs": [
                    "CroDCoYDCh8vZXRoZXJtaW50LmV2bS52MS5Nc2dFdGhlcmV1bVR4EuICCpICChovZXRoZXJtaW50LmV2bS52MS5MZWdhY3lUeBLzAQiCARINNTAwMDAwMDAwMDAwMBib3gQiKjB4QWE1M0RkNkQyMzRBMGM0MzFiMzlCOUU5MDQ1NDY2NjQzMjg2OWRjOSoBMDJkk4YIwgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABphbdB3AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGFt0HcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABXPXkUYDoCAsdCIBlg1/pB3FTVSqeStpRvv8xwDiSPMkepWGdhluJGIBIRSiA+prAX+yCgIqI6eo31qX3H0ABY5La4ppbkOr5yDvguCBEAAAAAAABqQBpCMHgzMTE4NTgzYjZmNzFlYmVkOTI0MTBhZmJkYzA2OWZhY2I5ZTk0MTY5YmQ3NjQ3MTFkNThjYTFmMTMxZDYzZmZm+j8uCiwvZXRoZXJtaW50LmV2bS52MS5FeHRlbnNpb25PcHRpb25zRXRoZXJldW1UeBImEiQKHgoIYmFzZXRjcm8SEjM4Nzk3NTAwMDAwMDAwMDAwMBCb3gQ="
                ]
            },
            "evidence": {
                "evidence": []
            },
            "last_commit": {
                "height": "83177",
                "round": 0,
                "block_id": {
                    "hash": "AE5259DC2C2D115254285E258956F2F54E8EE269A13889F155A73A9A2114B390",
                    "parts": {
                        "total": 1,
                        "hash": "06031C4D8A5D64AEAECC4E4655CC8D2540DDA7152857914815FCE4363E4FEBC4"
                    }
                },
                "signatures": [
                    {
                        "block_id_flag": 2,
                        "validator_address": "0421821D46C46F86F1EAD79EEB31CFA88A5578CB",
                        "timestamp": "2021-10-18T19:52:20.183678475Z",
                        "signature": "Myfldqx5a8MkQzYdVJ778y0t165SJHgu5nod7IeNyMHz8jke2tn7zCPWpyjNdzZFasjejUefpfUPKUr+ZZSqBQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "A3F7397F81D4CC82D566633F9728DFB50A35D965",
                        "timestamp": "2021-10-18T19:52:20.259005853Z",
                        "signature": "Zlt3udQndSjA34Yms3mO+Ms6hKF0cmPE6Agv7+qAFBdd5UaPIalPezrftdA/yW4yOug7e13kHegCU5E61E/0BQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "1473CC22B535A206B956CB47C65C272EE5324927",
                        "timestamp": "2021-10-18T19:52:20.545795303Z",
                        "signature": "L7Jc9NNz89WISlHnlgpItgQ0fML0jSpn/4BnjrIUuYJQdTXpyI4Xi/CMWmniCaNvMmAXTGl2DmpsLwsSSlxxAA=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "1E036C0FE3472F33A9FFC9F3ED14B56DE85BE389",
                        "timestamp": "2021-10-18T19:52:20.521107078Z",
                        "signature": "WEKX7qGtIEwmIkERbOIntZTwxL5WGSW9TEovBsdjMz4QgBk5r3CcorFs9s8jMiOoaWTMCPAws8X6G6HRhr1hCQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "559AB20018F2E53A89533C390DCA2CEA6721D869",
                        "timestamp": "2021-10-18T19:52:20.618105238Z",
                        "signature": "0ZWQmcjVgPuGZNbgPdfPN6DjEtrRvqGpAIqLw0k6PQ22oEjwHpcl8W2kdWy+cpwUIJja5S698puWL8eQkjusDA=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "5600A9A2FF66520305EC37F1F698FEE236FACF98",
                        "timestamp": "2021-10-18T19:52:20.536050607Z",
                        "signature": "kdv6BIiY73DyjZye+o3aWUYfj5cLPascYDYZTr2d31z1c7VpsCOBAPkmx+rCpvDzWfj0xxurvtnzrYlUiG8XAQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "72A52EBF69F6670C05A3EB2BA831ADBF629B8948",
                        "timestamp": "2021-10-18T19:52:20.658532552Z",
                        "signature": "QkA9DtEzrjX8nFzXjDpMi4Y0u2/8ZIDCCUSIT5cPJ0zNy8vLZsGmrcXeNl5TgiIVfuHQigiCs66yH/v1wlrdCg=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "739109BF2993601BEF624283C9A572DF3175CBA8",
                        "timestamp": "2021-10-18T19:52:20.4372106Z",
                        "signature": "4sewEyE/IMfcPXFpY+g4L7EGM4OwYxRi4lots8XOAwc1Z2jmvfwPwC/Yf09Lz1xMx1e8yW1FFLA0PYPR8mgjDg=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "A6C944F5ECC995C358572B7FF1B769B6DB60000D",
                        "timestamp": "2021-10-18T19:52:20.536371595Z",
                        "signature": "r+nfjdBr9p2A2XMJG//H5SRvXBPKtpcTVE8v0hNzXtAIf/2untIUgFsw/4d+zoV8rRRbCEBDWHJ8UUfRz0bsDQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "A72AA9230A83D4D6D18AD07EC03B5168DEA702AE",
                        "timestamp": "2021-10-18T19:52:20.536839936Z",
                        "signature": "ewSyVEHjudQ/26QLCbYkz6cWozXQGF3ISBv3cNuEbd9pMdf07Wgwj5zNIpkKI1OvYQBBPNnxT/71lmFqSbhpCg=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "A7E012CB6FFABA17C6C98600F3F7402DC1A06C8C",
                        "timestamp": "2021-10-18T19:52:20.536250727Z",
                        "signature": "iNJI44LHaZDl7jBbhYWDcFRtAyI9DUNJ71rn1p938cBaLOQkgowqVxZzC8VDh5+qsDms2twPhW7z+QXeTyx5CQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "AD582CC7756A0C2F3803840997BAF10075D6D005",
                        "timestamp": "2021-10-18T19:52:20.527174233Z",
                        "signature": "E2+3R2B3bxLlGqzsnKMopDMXsW6SYt/OLXNgwKPDHSmGP8vqdBGS9PRJr4eX6CbGYb4bsBMCrA3f2+0YZGseCw=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "B6581676389C2D34EC256DA777F3D9D743AD10AB",
                        "timestamp": "2021-10-18T19:52:20.653699905Z",
                        "signature": "dCN6XL2ke+yvKghrMhVv2kf25+oWbLRWuJvkkM35zIOL+pzPcuauQcBP8baCY73SfWaj962wI6RGmojoFZgJAg=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "B8735B4BBA3AD413B7DBB3A74ABC75A5907BE7A5",
                        "timestamp": "2021-10-18T19:52:20.564010383Z",
                        "signature": "s81PlF0qVa4n6UAP9u4wwTmu0HXcoPwtk7jTVaBKLjiEFajxzwNKiPx3FPEKWiKPA6k1UpUBn9nDozn4nbFtBg=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "BE937B566A3EBBD72E6553F3D216E2869D7FD1DD",
                        "timestamp": "2021-10-18T19:52:20.650446947Z",
                        "signature": "Ozat+GYFGu5ykLZCTHFKzeFtb7SYUdeJGwdoMSMPhmD6hV7JOx8iT7CSBV4ZmeY2JxwTCbIPh+5aeyVKEpYWDw=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "D03C1628B633072EF67E42141110CCD6AA420456",
                        "timestamp": "2021-10-18T19:52:20.536362626Z",
                        "signature": "OEksFHL61EdLQGczPD3hlnwv8JBIH8vIKJcVkaKBJ/KHLjU41T7WC+xWEfEn7OCzoHKfTkQ/8byWgER5wJRkAA=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "09D283BB0AC4B6A8BA05F6600E018E1D4DD25C12",
                        "timestamp": "2021-10-18T19:52:20.4609217Z",
                        "signature": "z0DLo2nFN8bcwLnPgpzwYMOQuJlvxfj5uyDFhhhDnh/Lmc2W+awa4Zh0652ewFWpHctcOu7JPyyuVSLESpUjCg=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "B359B56836AF6117E780045985D45F988065B746",
                        "timestamp": "2021-10-18T19:52:20.4687103Z",
                        "signature": "MwjlhlFfNGsZuWBHbfncG+2TM5NEtOF360i1Jct4osPdc0S85DELo3s9pBwFNNs4ealPVqKern9wLinWQr/BAg=="
                    }
                ]
            }
        }
    }
}`

const TX_MSG_ETHEREUM_TX_BLOCK_RESULTS_RESP = `{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "height": "83178",
        "txs_results": [
            {
                "code": 0,
                "data": "CrcECh8vZXRoZXJtaW50LmV2bS52MS5Nc2dFdGhlcmV1bVR4EpMECkIweDMxMTg1ODNiNmY3MWViZWQ5MjQxMGFmYmRjMDY5ZmFjYjllOTQxNjliZDc2NDcxMWQ1OGNhMWYxMzFkNjNmZmYSyAMKKjB4QWE1M0RkNkQyMzRBMGM0MzFiMzlCOUU5MDQ1NDY2NjQzMjg2OWRjORJCMHhkOThiYjRmZWNlMjRjMWViMzA2ZjgxZDhmZmZkMjEyNGRlOTg2OTRlNTUyZWMxYjFiMTA2YjNmYzY5ZDVlNTFhEkIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDFhNjE2ZGQwNzcSQjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA2MTZkZDA3NxJCMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA1NzNkNzkxNDYwIOqJBSpCMHgzMTE4NTgzYjZmNzFlYmVkOTI0MTBhZmJkYzA2OWZhY2I5ZTk0MTY5YmQ3NjQ3MTFkNThjYTFmMTMxZDYzZmZmOkIweDk3MGQzMzVlZDM2OWIyYTUzYTU4Y2MwMzY0NzcyMzEzM2IyYmMwMGQzY2M4NGI3OWE4M2Q5NjZjY2IxZjdjMzUom94E",
                "log": "[{\"events\":[{\"type\":\"ethereum_tx\",\"attributes\":[{\"key\":\"amount\",\"value\":\"0\"},{\"key\":\"ethereumTxHash\",\"value\":\"0x3118583b6f71ebed92410afbdc069facb9e94169bd764711d58ca1f131d63fff\"},{\"key\":\"txHash\",\"value\":\"2678437368AFC7E0E6D891D858F17B9C05CFEE850A786592A11992813D6A89FD\"},{\"key\":\"recipient\",\"value\":\"0xAa53Dd6D234A0c431b39B9E90454666432869dc9\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"ethereum_tx\"},{\"key\":\"module\",\"value\":\"evm\"},{\"key\":\"sender\",\"value\":\"0x6F966DA8f83ac4b4ae3DFbD2da1aDa7f333967b1\"},{\"key\":\"txType\",\"value\":\"0\"}]},{\"type\":\"tx_log\",\"attributes\":[{\"key\":\"txLog\",\"value\":\"{\\\"address\\\":\\\"0xAa53Dd6D234A0c431b39B9E90454666432869dc9\\\",\\\"topics\\\":[\\\"0xd98bb4fece24c1eb306f81d8fffd2124de98694e552ec1b1b106b3fc69d5e51a\\\",\\\"0x0000000000000000000000000000000000000000000000000000001a616dd077\\\",\\\"0x00000000000000000000000000000000000000000000000000000000616dd077\\\",\\\"0x000000000000000000000000000000000000000000000000000000573d791460\\\"],\\\"blockNumber\\\":83178,\\\"transactionHash\\\":\\\"0x3118583b6f71ebed92410afbdc069facb9e94169bd764711d58ca1f131d63fff\\\",\\\"transactionIndex\\\":0,\\\"blockHash\\\":\\\"0x970d335ed369b2a53a58cc03647723133b2bc00d3cc84b79a83d966ccb1f7c35\\\",\\\"logIndex\\\":0}\"}]}]}]",
                "info": "",
                "gas_wanted": "0",
                "gas_used": "77595",
                "events": [
                    {
                        "type": "coin_spent",
                        "attributes": [
                            {
                                "key": "c3BlbmRlcg==",
                                "value": "dGNyYzFkN3R4bTI4Yzh0enRmdDNhbDBmZDV4azYwdWVuamVhMzBrMnA1Zw==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "Mzg3OTc1MDAwMDAwMDAwMDAwYmFzZXRjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "coin_received",
                        "attributes": [
                            {
                                "key": "cmVjZWl2ZXI=",
                                "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "Mzg3OTc1MDAwMDAwMDAwMDAwYmFzZXRjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "dGNyYzFkN3R4bTI4Yzh0enRmdDNhbDBmZDV4azYwdWVuamVhMzBrMnA1Zw==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "Mzg3OTc1MDAwMDAwMDAwMDAwYmFzZXRjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "c2VuZGVy",
                                "value": "dGNyYzFkN3R4bTI4Yzh0enRmdDNhbDBmZDV4azYwdWVuamVhMzBrMnA1Zw==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "ZmVl",
                                "value": "Mzg3OTc1MDAwMDAwMDAwMDAwYmFzZXRjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "YWN0aW9u",
                                "value": "ZXRoZXJldW1fdHg=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "ethereum_tx",
                        "attributes": [
                            {
                                "key": "YW1vdW50",
                                "value": "MA==",
                                "index": true
                            },
                            {
                                "key": "ZXRoZXJldW1UeEhhc2g=",
                                "value": "MHgzMTE4NTgzYjZmNzFlYmVkOTI0MTBhZmJkYzA2OWZhY2I5ZTk0MTY5YmQ3NjQ3MTFkNThjYTFmMTMxZDYzZmZm",
                                "index": true
                            },
                            {
                                "key": "dHhIYXNo",
                                "value": "MjY3ODQzNzM2OEFGQzdFMEU2RDg5MUQ4NThGMTdCOUMwNUNGRUU4NTBBNzg2NTkyQTExOTkyODEzRDZBODlGRA==",
                                "index": true
                            },
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "MHhBYTUzRGQ2RDIzNEEwYzQzMWIzOUI5RTkwNDU0NjY2NDMyODY5ZGM5",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx_log",
                        "attributes": [
                            {
                                "key": "dHhMb2c=",
                                "value": "eyJhZGRyZXNzIjoiMHhBYTUzRGQ2RDIzNEEwYzQzMWIzOUI5RTkwNDU0NjY2NDMyODY5ZGM5IiwidG9waWNzIjpbIjB4ZDk4YmI0ZmVjZTI0YzFlYjMwNmY4MWQ4ZmZmZDIxMjRkZTk4Njk0ZTU1MmVjMWIxYjEwNmIzZmM2OWQ1ZTUxYSIsIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMWE2MTZkZDA3NyIsIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA2MTZkZDA3NyIsIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwNTczZDc5MTQ2MCJdLCJibG9ja051bWJlciI6ODMxNzgsInRyYW5zYWN0aW9uSGFzaCI6IjB4MzExODU4M2I2ZjcxZWJlZDkyNDEwYWZiZGMwNjlmYWNiOWU5NDE2OWJkNzY0NzExZDU4Y2ExZjEzMWQ2M2ZmZiIsInRyYW5zYWN0aW9uSW5kZXgiOjAsImJsb2NrSGFzaCI6IjB4OTcwZDMzNWVkMzY5YjJhNTNhNThjYzAzNjQ3NzIzMTMzYjJiYzAwZDNjYzg0Yjc5YTgzZDk2NmNjYjFmN2MzNSIsImxvZ0luZGV4IjowfQ==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "bW9kdWxl",
                                "value": "ZXZt",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "MHg2Rjk2NkRBOGY4M2FjNGI0YWUzREZiRDJkYTFhRGE3ZjMzMzk2N2Ix",
                                "index": true
                            },
                            {
                                "key": "dHhUeXBl",
                                "value": "MA==",
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
                "type": "coin_spent",
                "attributes": [
                    {
                        "key": "c3BlbmRlcg==",
                        "value": "dGNyYzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODM2NTI0MA==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_received",
                "attributes": [
                    {
                        "key": "cmVjZWl2ZXI=",
                        "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    }
                ]
            },
            {
                "type": "transfer",
                "attributes": [
                    {
                        "key": "cmVjaXBpZW50",
                        "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
                        "index": true
                    },
                    {
                        "key": "c2VuZGVy",
                        "value": "dGNyYzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODM2NTI0MA==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    }
                ]
            },
            {
                "type": "message",
                "attributes": [
                    {
                        "key": "c2VuZGVy",
                        "value": "dGNyYzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODM2NTI0MA==",
                        "index": true
                    }
                ]
            },
            {
                "type": "mint",
                "attributes": [
                    {
                        "key": "Ym9uZGVkX3JhdGlv",
                        "value": "MC4wMDAwMDAwMDAxMDcyOTk5OTg=",
                        "index": true
                    },
                    {
                        "key": "aW5mbGF0aW9u",
                        "value": "MC4wMDAwMDAwMDAwMDAwMDAwMDA=",
                        "index": true
                    },
                    {
                        "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
                        "value": "MC4wMDAwMDAwMDAwMDAwMDAwMDA=",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MA==",
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_spent",
                "attributes": [
                    {
                        "key": "c3BlbmRlcg==",
                        "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_received",
                "attributes": [
                    {
                        "key": "cmVjZWl2ZXI=",
                        "value": "dGNyYzFqdjY1czNncnFmNnY2amwzZHA0dDZjOXQ5cms5OWNkODc1aHdtcw==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    }
                ]
            },
            {
                "type": "transfer",
                "attributes": [
                    {
                        "key": "cmVjaXBpZW50",
                        "value": "dGNyYzFqdjY1czNncnFmNnY2amwzZHA0dDZjOXQ5cms5OWNkODc1aHdtcw==",
                        "index": true
                    },
                    {
                        "key": "c2VuZGVy",
                        "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    }
                ]
            },
            {
                "type": "message",
                "attributes": [
                    {
                        "key": "c2VuZGVy",
                        "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
                        "index": true
                    }
                ]
            },
            {
                "type": "proposer_reward",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxczllM2Y1eTR0dDRma3ozanlqODY1cWF1ZDJjcWhzNjZ5bm1qNGY=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxczllM2Y1eTR0dDRma3ozanlqODY1cWF1ZDJjcWhzNjZ5bm1qNGY=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxczllM2Y1eTR0dDRma3ozanlqODY1cWF1ZDJjcWhzNjZ5bm1qNGY=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxdjc2cjd1NHV5cjNld2RrczhjcW11dzdjYTRsZWp2Yzh1bmFnMm0=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxdjc2cjd1NHV5cjNld2RrczhjcW11dzdjYTRsZWp2Yzh1bmFnMm0=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxczllM2Y1eTR0dDRma3ozanlqODY1cWF1ZDJjcWhzNjZ5bm1qNGY=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxczllM2Y1eTR0dDRma3ozanlqODY1cWF1ZDJjcWhzNjZ5bm1qNGY=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxc2h5OHY0ZnQ3OG16dnNzZDkwYWw1Y2x6MzNydTZhZ3JsYW04Z2U=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxc2h5OHY0ZnQ3OG16dnNzZDkwYWw1Y2x6MzNydTZhZ3JsYW04Z2U=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxMDhnODBoZmcwcWt6d3U1bm4wN3Bkcnhqc2R4ZmthY2g5c2N0cjQ=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxMDhnODBoZmcwcWt6d3U1bm4wN3Bkcnhqc2R4ZmthY2g5c2N0cjQ=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxbGxkamhqbm4zMmU4dmVrN2N4ZTlnMDVuZjhqNzR5MHhsejVnMjM=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxbGxkamhqbm4zMmU4dmVrN2N4ZTlnMDVuZjhqNzR5MHhsejVnMjM=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxbXRjbjI1MDVrMzdtbHp0eXdmOGVnOHNwdjBrcG5zcWFwYWZ6c3o=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxbXRjbjI1MDVrMzdtbHp0eXdmOGVnOHNwdjBrcG5zcWFwYWZ6c3o=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxejI3bG4zM3FsZG45NnVhMmxkZDRkdnk3cDhlMjlwN3dsbXhzdGw=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxejI3bG4zM3FsZG45NnVhMmxkZDRkdnk3cDhlMjlwN3dsbXhzdGw=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxNXE4Zmx2bTVxZnp0dDV6dHY5bXkyaHQyeGwyeGFra3g5c2F4eDc=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxNXE4Zmx2bTVxZnp0dDV6dHY5bXkyaHQyeGwyeGFra3g5c2F4eDc=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxdW1zaGY4dDZjY3NkZHpua2cwOXNheHlycDIzNXF6c2Z2MDBlNWs=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxdW1zaGY4dDZjY3NkZHpua2cwOXNheHlycDIzNXF6c2Z2MDBlNWs=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxM2F3dm0zNnVsZTl0NXNkank4YXZ5cndkd2t2bHZtaDNzcHE2eDU=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxM2F3dm0zNnVsZTl0NXNkank4YXZ5cndkd2t2bHZtaDNzcHE2eDU=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxcTBnY3EwMjVqeHM4ZnJzNnZhZGc4cnphNjUwNzI0MzU2eWFzYXU=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxcTBnY3EwMjVqeHM4ZnJzNnZhZGc4cnphNjUwNzI0MzU2eWFzYXU=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxM2s2ejYyY2c2MHhmcW10MGFqbGU2dTRmemc0dWU0MmZncWM3bno=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxM2s2ejYyY2c2MHhmcW10MGFqbGU2dTRmemc0dWU0MmZncWM3bno=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxeHlodjJnYzZ1dnR1am0zNWN0eTh0Z2czeDBudjdyanMwZWd3NG4=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxeHlodjJnYzZ1dnR1am0zNWN0eTh0Z2czeDBudjdyanMwZWd3NG4=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxa2Q3bGRoZHh6cjM4ZWt5bnFranJwM3A5bHhwZGN6eHhzeWFqYXE=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxa2Q3bGRoZHh6cjM4ZWt5bnFranJwM3A5bHhwZGN6eHhzeWFqYXE=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxZ3l3eGZkcjRuZXRzODdjenhscGNhenBsYTVxY2tlZmpkczNjeHU=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxZ3l3eGZkcjRuZXRzODdjenhscGNhenBsYTVxY2tlZmpkczNjeHU=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxNXV0dTJjMjN2Mmg0d3RkZXV3bXI0eXZjaHFseGphZnY4dTgzcnI=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxNXV0dTJjMjN2Mmg0d3RkZXV3bXI0eXZjaHFseGphZnY4dTgzcnI=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxaDBqY2cwM3hqZnE0bjl3NmwybHh1dHA4bDU1bm1jejl3enkwMHY=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxaDBqY2cwM3hqZnE0bjl3NmwybHh1dHA4bDU1bm1jejl3enkwMHY=",
                        "index": true
                    }
                ]
            },
            {
                "type": "commission",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxOWE2cjc0ZHZmeGp5dmp6ZjNwZzl5M3k1cmhrNnJkczJwaDM5OHk=",
                        "index": true
                    }
                ]
            },
            {
                "type": "rewards",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": null,
                        "index": true
                    },
                    {
                        "key": "dmFsaWRhdG9y",
                        "value": "dGNyY3ZhbG9wZXIxOWE2cjc0ZHZmeGp5dmp6ZjNwZzl5M3k1cmhrNnJkczJwaDM5OHk=",
                        "index": true
                    }
                ]
            }
        ],
        "end_block_events": [
            {
                "type": "block_bloom",
                "attributes": [
                    {
                        "key": "Ymxvb20=",
                        "value": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAIAEAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAQAACAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAA==",
                        "index": true
                    }
                ]
            }
        ],
        "validator_updates": null,
        "consensus_param_updates": {
            "block": {
                "max_bytes": "1048576",
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
        }
    }
}`

const TX_MSG_ETHEREUM_TX_TXS_RESP = `{
  "tx": {
    "body": {
      "messages": [
        {
          "@type": "/ethermint.evm.v1.MsgEthereumTx",
          "data": {
            "@type": "/ethermint.evm.v1.LegacyTx",
            "nonce": "130",
            "gas_price": "5000000000000",
            "gas": "77595",
            "to": "0xAa53Dd6D234A0c431b39B9E90454666432869dc9",
            "value": "0",
            "data": "k4YIwgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABphbdB3AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGFt0HcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABXPXkUYA==",
            "v": "Asc=",
            "r": "GWDX+kHcVNVKp5K2lG+/zHAOJI8yR6lYZ2GW4kYgEhE=",
            "s": "PqawF/sgoCKiOnqN9al9x9AAWOS2uKaW5Dq+cg74Lgg="
          },
          "size": 208,
          "hash": "0x3118583b6f71ebed92410afbdc069facb9e94169bd764711d58ca1f131d63fff",
          "from": ""
        }
      ],
      "memo": "",
      "timeout_height": "0",
      "extension_options": [
        {
          "@type": "/ethermint.evm.v1.ExtensionOptionsEthereumTx"
        }
      ],
      "non_critical_extension_options": [
      ]
    },
    "auth_info": {
      "signer_infos": [
      ],
      "fee": {
        "amount": [
          {
            "denom": "basetcro",
            "amount": "387975000000000000"
          }
        ],
        "gas_limit": "77595",
        "payer": "",
        "granter": ""
      }
    },
    "signatures": [
    ]
  },
  "tx_response": {
    "height": "83178",
    "txhash": "2678437368AFC7E0E6D891D858F17B9C05CFEE850A786592A11992813D6A89FD",
    "codespace": "",
    "code": 0,
    "data": "0AB7040A1F2F65746865726D696E742E65766D2E76312E4D7367457468657265756D54781293040A4230783331313835383362366637316562656439323431306166626463303639666163623965393431363962643736343731316435386361316631333164363366666612C8030A2A307841613533446436443233344130633433316233394239453930343534363636343332383639646339124230786439386262346665636532346331656233303666383164386666666432313234646539383639346535353265633162316231303662336663363964356535316112423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030316136313664643037371242307830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303631366464303737124230783030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303537336437393134363020EA89052A423078333131383538336236663731656265643932343130616662646330363966616362396539343136396264373634373131643538636131663133316436336666663A42307839373064333335656433363962326135336135386363303336343737323331333362326263303064336363383462373961383364393636636362316637633335289BDE04",
    "raw_log": "[{\"events\":[{\"type\":\"ethereum_tx\",\"attributes\":[{\"key\":\"amount\",\"value\":\"0\"},{\"key\":\"ethereumTxHash\",\"value\":\"0x3118583b6f71ebed92410afbdc069facb9e94169bd764711d58ca1f131d63fff\"},{\"key\":\"txHash\",\"value\":\"2678437368AFC7E0E6D891D858F17B9C05CFEE850A786592A11992813D6A89FD\"},{\"key\":\"recipient\",\"value\":\"0xAa53Dd6D234A0c431b39B9E90454666432869dc9\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"ethereum_tx\"},{\"key\":\"module\",\"value\":\"evm\"},{\"key\":\"sender\",\"value\":\"0x6F966DA8f83ac4b4ae3DFbD2da1aDa7f333967b1\"},{\"key\":\"txType\",\"value\":\"0\"}]},{\"type\":\"tx_log\",\"attributes\":[{\"key\":\"txLog\",\"value\":\"{\\\"address\\\":\\\"0xAa53Dd6D234A0c431b39B9E90454666432869dc9\\\",\\\"topics\\\":[\\\"0xd98bb4fece24c1eb306f81d8fffd2124de98694e552ec1b1b106b3fc69d5e51a\\\",\\\"0x0000000000000000000000000000000000000000000000000000001a616dd077\\\",\\\"0x00000000000000000000000000000000000000000000000000000000616dd077\\\",\\\"0x000000000000000000000000000000000000000000000000000000573d791460\\\"],\\\"blockNumber\\\":83178,\\\"transactionHash\\\":\\\"0x3118583b6f71ebed92410afbdc069facb9e94169bd764711d58ca1f131d63fff\\\",\\\"transactionIndex\\\":0,\\\"blockHash\\\":\\\"0x970d335ed369b2a53a58cc03647723133b2bc00d3cc84b79a83d966ccb1f7c35\\\",\\\"logIndex\\\":0}\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "log": "",
        "events": [
          {
            "type": "ethereum_tx",
            "attributes": [
              {
                "key": "amount",
                "value": "0"
              },
              {
                "key": "ethereumTxHash",
                "value": "0x3118583b6f71ebed92410afbdc069facb9e94169bd764711d58ca1f131d63fff"
              },
              {
                "key": "txHash",
                "value": "2678437368AFC7E0E6D891D858F17B9C05CFEE850A786592A11992813D6A89FD"
              },
              {
                "key": "recipient",
                "value": "0xAa53Dd6D234A0c431b39B9E90454666432869dc9"
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "ethereum_tx"
              },
              {
                "key": "module",
                "value": "evm"
              },
              {
                "key": "sender",
                "value": "0x6F966DA8f83ac4b4ae3DFbD2da1aDa7f333967b1"
              },
              {
                "key": "txType",
                "value": "0"
              }
            ]
          },
          {
            "type": "tx_log",
            "attributes": [
              {
                "key": "txLog",
                "value": "{\"address\":\"0xAa53Dd6D234A0c431b39B9E90454666432869dc9\",\"topics\":[\"0xd98bb4fece24c1eb306f81d8fffd2124de98694e552ec1b1b106b3fc69d5e51a\",\"0x0000000000000000000000000000000000000000000000000000001a616dd077\",\"0x00000000000000000000000000000000000000000000000000000000616dd077\",\"0x000000000000000000000000000000000000000000000000000000573d791460\"],\"blockNumber\":83178,\"transactionHash\":\"0x3118583b6f71ebed92410afbdc069facb9e94169bd764711d58ca1f131d63fff\",\"transactionIndex\":0,\"blockHash\":\"0x970d335ed369b2a53a58cc03647723133b2bc00d3cc84b79a83d966ccb1f7c35\",\"logIndex\":0}"
              }
            ]
          }
        ]
      }
    ],
    "info": "",
    "gas_wanted": "0",
    "gas_used": "77595",
    "tx": {
      "@type": "/cosmos.tx.v1beta1.Tx",
      "body": {
        "messages": [
          {
            "@type": "/ethermint.evm.v1.MsgEthereumTx",
            "data": {
              "@type": "/ethermint.evm.v1.LegacyTx",
              "nonce": "130",
              "gas_price": "5000000000000",
              "gas": "77595",
              "to": "0xAa53Dd6D234A0c431b39B9E90454666432869dc9",
              "value": "0",
              "data": "k4YIwgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABphbdB3AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGFt0HcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABXPXkUYA==",
              "v": "Asc=",
              "r": "GWDX+kHcVNVKp5K2lG+/zHAOJI8yR6lYZ2GW4kYgEhE=",
              "s": "PqawF/sgoCKiOnqN9al9x9AAWOS2uKaW5Dq+cg74Lgg="
            },
            "size": 208,
            "hash": "0x3118583b6f71ebed92410afbdc069facb9e94169bd764711d58ca1f131d63fff",
            "from": ""
          }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [
          {
            "@type": "/ethermint.evm.v1.ExtensionOptionsEthereumTx"
          }
        ],
        "non_critical_extension_options": [
        ]
      },
      "auth_info": {
        "signer_infos": [
        ],
        "fee": {
          "amount": [
            {
              "denom": "basetcro",
              "amount": "387975000000000000"
            }
          ],
          "gas_limit": "77595",
          "payer": "",
          "granter": ""
        }
      },
      "signatures": [
      ]
    },
    "timestamp": "2021-10-18T19:52:20Z",
    "events": [
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNyYzFkN3R4bTI4Yzh0enRmdDNhbDBmZDV4azYwdWVuamVhMzBrMnA1Zw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "Mzg3OTc1MDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "Mzg3OTc1MDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNyYzFkN3R4bTI4Yzh0enRmdDNhbDBmZDV4azYwdWVuamVhMzBrMnA1Zw==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "Mzg3OTc1MDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNyYzFkN3R4bTI4Yzh0enRmdDNhbDBmZDV4azYwdWVuamVhMzBrMnA1Zw==",
            "index": true
          }
        ]
      },
      {
        "type": "tx",
        "attributes": [
          {
            "key": "ZmVl",
            "value": "Mzg3OTc1MDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "ZXRoZXJldW1fdHg=",
            "index": true
          }
        ]
      },
      {
        "type": "ethereum_tx",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MA==",
            "index": true
          },
          {
            "key": "ZXRoZXJldW1UeEhhc2g=",
            "value": "MHgzMTE4NTgzYjZmNzFlYmVkOTI0MTBhZmJkYzA2OWZhY2I5ZTk0MTY5YmQ3NjQ3MTFkNThjYTFmMTMxZDYzZmZm",
            "index": true
          },
          {
            "key": "dHhIYXNo",
            "value": "MjY3ODQzNzM2OEFGQzdFMEU2RDg5MUQ4NThGMTdCOUMwNUNGRUU4NTBBNzg2NTkyQTExOTkyODEzRDZBODlGRA==",
            "index": true
          },
          {
            "key": "cmVjaXBpZW50",
            "value": "MHhBYTUzRGQ2RDIzNEEwYzQzMWIzOUI5RTkwNDU0NjY2NDMyODY5ZGM5",
            "index": true
          }
        ]
      },
      {
        "type": "tx_log",
        "attributes": [
          {
            "key": "dHhMb2c=",
            "value": "eyJhZGRyZXNzIjoiMHhBYTUzRGQ2RDIzNEEwYzQzMWIzOUI5RTkwNDU0NjY2NDMyODY5ZGM5IiwidG9waWNzIjpbIjB4ZDk4YmI0ZmVjZTI0YzFlYjMwNmY4MWQ4ZmZmZDIxMjRkZTk4Njk0ZTU1MmVjMWIxYjEwNmIzZmM2OWQ1ZTUxYSIsIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMWE2MTZkZDA3NyIsIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA2MTZkZDA3NyIsIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwNTczZDc5MTQ2MCJdLCJibG9ja051bWJlciI6ODMxNzgsInRyYW5zYWN0aW9uSGFzaCI6IjB4MzExODU4M2I2ZjcxZWJlZDkyNDEwYWZiZGMwNjlmYWNiOWU5NDE2OWJkNzY0NzExZDU4Y2ExZjEzMWQ2M2ZmZiIsInRyYW5zYWN0aW9uSW5kZXgiOjAsImJsb2NrSGFzaCI6IjB4OTcwZDMzNWVkMzY5YjJhNTNhNThjYzAzNjQ3NzIzMTMzYjJiYzAwZDNjYzg0Yjc5YTgzZDk2NmNjYjFmN2MzNSIsImxvZ0luZGV4IjowfQ==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "bW9kdWxl",
            "value": "ZXZt",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "MHg2Rjk2NkRBOGY4M2FjNGI0YWUzREZiRDJkYTFhRGE3ZjMzMzk2N2Ix",
            "index": true
          },
          {
            "key": "dHhUeXBl",
            "value": "MA==",
            "index": true
          }
        ]
      }
    ]
  }
}`
