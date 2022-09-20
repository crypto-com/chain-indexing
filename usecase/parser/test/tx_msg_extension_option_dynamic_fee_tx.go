package usecase_parser_test

const TX_MSG_EXTENSION_OPTION_DYNAMIC_FEE_TX_BLOCK_RESP = `{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "block_id": {
            "hash": "BAE99CBA6E908A6AA2FFF5EF9F826AC81A30DDF9717C4966A1C0E9C612B84814",
            "parts": {
                "total": 1,
                "hash": "4B23C9376865D93750C8D99B76626355B7B0A254B4F161DE5F913A374DF3643A"
            }
        },
        "block": {
            "header": {
                "version": {
                    "block": "11"
                },
                "chain_id": "cronostestnet_338-3",
                "height": "5168311",
                "time": "2022-09-16T03:14:55.502206912Z",
                "last_block_id": {
                    "hash": "1BAB2F586A90A244BC2ABDE749C5DA15010451F8EE08969354E9FD7FF114B09E",
                    "parts": {
                        "total": 1,
                        "hash": "1BB8EAE9BE456DE2C3F3F057F5890ED3D009EB5D35AD0D08FCA52DCB0A5568E8"
                    }
                },
                "last_commit_hash": "99238AFBBE27F4044D88333C80F10D3B64E17F06ED99A37DF2BA56AF53B4D79F",
                "data_hash": "6250BA220F846B3CC45103440CB771D26C0485116629EC25ADB28A76ECB91317",
                "validators_hash": "133A3E0C61E2C3F64BB5A8541A03359C99F7CD9567521922ED9DD03E428940E4",
                "next_validators_hash": "133A3E0C61E2C3F64BB5A8541A03359C99F7CD9567521922ED9DD03E428940E4",
                "consensus_hash": "5B1778CC6A8A0D4443986BB557A8AB5231AA1DE78E4748B9731C8647D5616985",
                "app_hash": "9FADC232331354963F288185BC5CCF43372C5A535EF696F9B5769190B27334BD",
                "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                "proposer_address": "0421821D46C46F86F1EAD79EEB31CFA88A5578CB"
            },
            "data": {
                "txs": [
                    "CqQDCvACCh8vZXRoZXJtaW50LmV2bS52MS5Nc2dFdGhlcmV1bVR4EswCCoUCCh4vZXRoZXJtaW50LmV2bS52MS5EeW5hbWljRmVlVHgS4gEKAzMzOBBzGg01MDAwMDAwMDAwMDAwIg01MDAwMDAwMDAwMDAwKML9AjIqMHg5MDRCZDVhNUFBQzBCOWQ4OEEwRDQ3ODY0NzI0MjE4OTg2QWQ0YTNhOgEwQkQJXqezAAAAAAAAAAAAAAAA9WTvADS7fXzYRBInXh5RNk1NdzT//////////////////////////////////////////1ogHHigR02IN48gvEM991C/MOCmbVfOw9Sj804IhtI5z2RiIErKR5ojgKDOG0HbwIoPSPD0My7+Rk8xJYXETqaTtFqVGkIweGQ1NmRiNjU2NzJhNDhkYWFmMTRmNzQzZjUzYzMxNmEzZWJhMTMwZDcyYmViNzAzNDI5NmY0ZWEzZjdlNDliNTP6Py4KLC9ldGhlcm1pbnQuZXZtLnYxLkV4dGVuc2lvbk9wdGlvbnNFdGhlcmV1bVR4EiYSJAoeCghiYXNldGNybxISMjQ0MTcwMDAwMDAwMDAwMDAwEML9Ag=="
                ]
            },
            "evidence": {
                "evidence": []
            },
            "last_commit": {
                "height": "5168310",
                "round": 0,
                "block_id": {
                    "hash": "1BAB2F586A90A244BC2ABDE749C5DA15010451F8EE08969354E9FD7FF114B09E",
                    "parts": {
                        "total": 1,
                        "hash": "1BB8EAE9BE456DE2C3F3F057F5890ED3D009EB5D35AD0D08FCA52DCB0A5568E8"
                    }
                },
                "signatures": [
                    {
                        "block_id_flag": 2,
                        "validator_address": "0421821D46C46F86F1EAD79EEB31CFA88A5578CB",
                        "timestamp": "2022-09-16T03:14:55.412522098Z",
                        "signature": "T9VpDeEJVBQGyDZOlnhP8Rij8FNlOvnCEcpeWLUEZzrIXKGYeMfd3OhJHrdqYvbPBAx8dZOwRwuImCC+OC+cCw=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "A3F7397F81D4CC82D566633F9728DFB50A35D965",
                        "timestamp": "2022-09-16T03:14:55.502206912Z",
                        "signature": "c6wzGeXjsmITKQvXZT97atZj3n4tQHwZY3x47/3V3VONWUkLNA47CFVfZbJcVC9Cd7CPewTeg0nrwHLaDdCzAQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "739109BF2993601BEF624283C9A572DF3175CBA8",
                        "timestamp": "2022-09-16T03:14:55.904650283Z",
                        "signature": "w5hHQg0iURWtrUV/zoyRvWId6Cs7cmzOZxP6oVGXYbNDNiQ/ky6kdOlQUmttGihtptOTw7H5iFJgz62ZFNMtAg=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "A7E012CB6FFABA17C6C98600F3F7402DC1A06C8C",
                        "timestamp": "2022-09-16T03:14:56.133612785Z",
                        "signature": "1N9rYVPlbnzcqk9PA18AKRByaF8D1Uw9UenkDE9w3Kb5WJ4Mvw4H/jjuVQnsErtFi0UogjEPyfSFfslT03m2Aw=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "B8735B4BBA3AD413B7DBB3A74ABC75A5907BE7A5",
                        "timestamp": "2022-09-16T03:14:55.922103446Z",
                        "signature": "tv/YzR+ifULm4wH2LuolJeocN4a0rV0mOJpT2siQb3Yh1HHDHySyowzzUbQ+K66tz7EPzn2StD+CuCwLjtPIBQ=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "09D283BB0AC4B6A8BA05F6600E018E1D4DD25C12",
                        "timestamp": "2022-09-16T03:14:55.9156978Z",
                        "signature": "ovtvbZBC+LxFzFDHF/rE06GQ5YHuceGK5JPk3uJqZgn7GLa6j5NibOxVcy+3Fvr6WJUgdDjcJZt2jQMKhoe2BA=="
                    },
                    {
                        "block_id_flag": 2,
                        "validator_address": "3ECB552B1E6651DD70CE27923099F96408A7703D",
                        "timestamp": "2022-09-16T03:14:55.926903388Z",
                        "signature": "JoItuXg74Cp3aLlmofHVkcqe3+45OhH1ryXjSB8ns+ccqxjE9aqxRS+AMvfMxWgc3miZNPCDhb8DqWeI6qwDCw=="
                    },
                    {
                        "block_id_flag": 3,
                        "validator_address": "B359B56836AF6117E780045985D45F988065B746",
                        "timestamp": "2022-09-16T03:14:55.9657163Z",
                        "signature": "BhDo8ceW2/AFaxeBqDmt3lUNoTZaWVikOSYpmYPeTEN2deHaqZNHM+bFhseN/vfIzCB51VhUjLGXQ7LWQ+gbBw=="
                    }
                ]
            }
        }
    }
}`

const TX_MSG_EXTENSION_OPTION_DYNAMIC_FEE_TX_BLOCK_RESULTS_RESP = `{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
        "height": "5168311",
        "txs_results": [
            {
                "code": 0,
                "data": "EsAECicvZXRoZXJtaW50LmV2bS52MS5Nc2dFdGhlcmV1bVR4UmVzcG9uc2USlAQKQjB4ZDU2ZGI2NTY3MmE0OGRhYWYxNGY3NDNmNTNjMzE2YTNlYmExMzBkNzJiZWI3MDM0Mjk2ZjRlYTNmN2U0OWI1MxKnAwoqMHg5MDRCZDVhNUFBQzBCOWQ4OEEwRDQ3ODY0NzI0MjE4OTg2QWQ0YTNhEkIweDhjNWJlMWU1ZWJlYzdkNWJkMTRmNzE0MjdkMWU4NGYzZGQwMzE0YzBmN2IyMjkxZTViMjAwYWM4YzdjM2I5MjUSQjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwZTA4OWQxYTBkMTUxYTEzZDgxMmVhNGU4ZDY0OGYxZGM1OGI0MWNhORJCMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDBmNTY0ZWYwMDM0YmI3ZDdjZDg0NDEyMjc1ZTFlNTEzNjRkNGQ3NzM0GiD//////////////////////////////////////////yC3ubsCKkIweGQ1NmRiNjU2NzJhNDhkYWFmMTRmNzQzZjUzYzMxNmEzZWJhMTMwZDcyYmViNzAzNDI5NmY0ZWEzZjdlNDliNTM6QjB4YmFlOTljYmE2ZTkwOGE2YWEyZmZmNWVmOWY4MjZhYzgxYTMwZGRmOTcxN2M0OTY2YTFjMGU5YzYxMmI4NDgxNBogAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEo69oC",
                "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"tcrc1uzyargx32xsnmqfw5n5dvj83m3vtg89fhjges3\"},{\"key\":\"amount\",\"value\":\"22195000000000000basetcro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"tcrc17xpfvakm2amg962yls6f84z3kell8c5lfjsjej\"},{\"key\":\"amount\",\"value\":\"22195000000000000basetcro\"}]},{\"type\":\"ethereum_tx\",\"attributes\":[{\"key\":\"amount\",\"value\":\"0\"},{\"key\":\"ethereumTxHash\",\"value\":\"0xd56db65672a48daaf14f743f53c316a3eba130d72beb7034296f4ea3f7e49b53\"},{\"key\":\"txIndex\",\"value\":\"0\"},{\"key\":\"txGasUsed\",\"value\":\"44395\"},{\"key\":\"txHash\",\"value\":\"3F1B98E6A43A3666699CA28DA2A0872E1478E092F7EC3FCF8A94202BB8B330D2\"},{\"key\":\"recipient\",\"value\":\"0x904Bd5a5AAC0B9d88A0D47864724218986Ad4a3a\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/ethermint.evm.v1.EthermintLegacyTx\"},{\"key\":\"sender\",\"value\":\"tcrc17xpfvakm2amg962yls6f84z3kell8c5lfjsjej\"},{\"key\":\"module\",\"value\":\"evm\"},{\"key\":\"sender\",\"value\":\"0xE089d1a0d151A13D812eA4e8D648f1Dc58B41ca9\"},{\"key\":\"txType\",\"value\":\"2\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcrc1uzyargx32xsnmqfw5n5dvj83m3vtg89fhjges3\"},{\"key\":\"sender\",\"value\":\"tcrc17xpfvakm2amg962yls6f84z3kell8c5lfjsjej\"},{\"key\":\"amount\",\"value\":\"22195000000000000basetcro\"}]},{\"type\":\"tx_log\",\"attributes\":[{\"key\":\"txLog\",\"value\":\"{\\\"address\\\":\\\"0x904Bd5a5AAC0B9d88A0D47864724218986Ad4a3a\\\",\\\"topics\\\":[\\\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\\\",\\\"0x000000000000000000000000e089d1a0d151a13d812ea4e8d648f1dc58b41ca9\\\",\\\"0x000000000000000000000000f564ef0034bb7d7cd84412275e1e51364d4d7734\\\"],\\\"data\\\":\\\"//////////////////////////////////////////8=\\\",\\\"blockNumber\\\":5168311,\\\"transactionHash\\\":\\\"0xd56db65672a48daaf14f743f53c316a3eba130d72beb7034296f4ea3f7e49b53\\\",\\\"transactionIndex\\\":0,\\\"blockHash\\\":\\\"0xbae99cba6e908a6aa2fff5ef9f826ac81a30ddf9717c4966a1c0e9c612b84814\\\",\\\"logIndex\\\":0}\"}]}]}]",
                "info": "",
                "gas_wanted": "48834",
                "gas_used": "44395",
                "events": [
                    {
                        "type": "coin_spent",
                        "attributes": [
                            {
                                "key": "c3BlbmRlcg==",
                                "value": "dGNyYzF1enlhcmd4MzJ4c25tcWZ3NW41ZHZqODNtM3Z0Zzg5ZmhqZ2VzMw==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MjQ0MTcwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
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
                                "value": "MjQ0MTcwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
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
                                "value": "dGNyYzF1enlhcmd4MzJ4c25tcWZ3NW41ZHZqODNtM3Z0Zzg5ZmhqZ2VzMw==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MjQ0MTcwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "c2VuZGVy",
                                "value": "dGNyYzF1enlhcmd4MzJ4c25tcWZ3NW41ZHZqODNtM3Z0Zzg5ZmhqZ2VzMw==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx",
                        "attributes": [
                            {
                                "key": "ZmVl",
                                "value": "MjQ0MTcwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "ethereum_tx",
                        "attributes": [
                            {
                                "key": "ZXRoZXJldW1UeEhhc2g=",
                                "value": "MHhkNTZkYjY1NjcyYTQ4ZGFhZjE0Zjc0M2Y1M2MzMTZhM2ViYTEzMGQ3MmJlYjcwMzQyOTZmNGVhM2Y3ZTQ5YjUz",
                                "index": true
                            },
                            {
                                "key": "dHhJbmRleA==",
                                "value": "MA==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "YWN0aW9u",
                                "value": "L2V0aGVybWludC5ldm0udjEuTXNnRXRoZXJldW1UeA==",
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
                                "value": "MjIxOTUwMDAwMDAwMDAwMDBiYXNldGNybw==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "coin_received",
                        "attributes": [
                            {
                                "key": "cmVjZWl2ZXI=",
                                "value": "dGNyYzF1enlhcmd4MzJ4c25tcWZ3NW41ZHZqODNtM3Z0Zzg5ZmhqZ2VzMw==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MjIxOTUwMDAwMDAwMDAwMDBiYXNldGNybw==",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "dGNyYzF1enlhcmd4MzJ4c25tcWZ3NW41ZHZqODNtM3Z0Zzg5ZmhqZ2VzMw==",
                                "index": true
                            },
                            {
                                "key": "c2VuZGVy",
                                "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
                                "index": true
                            },
                            {
                                "key": "YW1vdW50",
                                "value": "MjIxOTUwMDAwMDAwMDAwMDBiYXNldGNybw==",
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
                        "type": "ethereum_tx",
                        "attributes": [
                            {
                                "key": "YW1vdW50",
                                "value": "MA==",
                                "index": true
                            },
                            {
                                "key": "ZXRoZXJldW1UeEhhc2g=",
                                "value": "MHhkNTZkYjY1NjcyYTQ4ZGFhZjE0Zjc0M2Y1M2MzMTZhM2ViYTEzMGQ3MmJlYjcwMzQyOTZmNGVhM2Y3ZTQ5YjUz",
                                "index": true
                            },
                            {
                                "key": "dHhJbmRleA==",
                                "value": "MA==",
                                "index": true
                            },
                            {
                                "key": "dHhHYXNVc2Vk",
                                "value": "NDQzOTU=",
                                "index": true
                            },
                            {
                                "key": "dHhIYXNo",
                                "value": "M0YxQjk4RTZBNDNBMzY2NjY5OUNBMjhEQTJBMDg3MkUxNDc4RTA5MkY3RUMzRkNGOEE5NDIwMkJCOEIzMzBEMg==",
                                "index": true
                            },
                            {
                                "key": "cmVjaXBpZW50",
                                "value": "MHg5MDRCZDVhNUFBQzBCOWQ4OEEwRDQ3ODY0NzI0MjE4OTg2QWQ0YTNh",
                                "index": true
                            }
                        ]
                    },
                    {
                        "type": "tx_log",
                        "attributes": [
                            {
                                "key": "dHhMb2c=",
                                "value": "eyJhZGRyZXNzIjoiMHg5MDRCZDVhNUFBQzBCOWQ4OEEwRDQ3ODY0NzI0MjE4OTg2QWQ0YTNhIiwidG9waWNzIjpbIjB4OGM1YmUxZTVlYmVjN2Q1YmQxNGY3MTQyN2QxZTg0ZjNkZDAzMTRjMGY3YjIyOTFlNWIyMDBhYzhjN2MzYjkyNSIsIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwZTA4OWQxYTBkMTUxYTEzZDgxMmVhNGU4ZDY0OGYxZGM1OGI0MWNhOSIsIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwZjU2NGVmMDAzNGJiN2Q3Y2Q4NDQxMjI3NWUxZTUxMzY0ZDRkNzczNCJdLCJkYXRhIjoiLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vOD0iLCJibG9ja051bWJlciI6NTE2ODMxMSwidHJhbnNhY3Rpb25IYXNoIjoiMHhkNTZkYjY1NjcyYTQ4ZGFhZjE0Zjc0M2Y1M2MzMTZhM2ViYTEzMGQ3MmJlYjcwMzQyOTZmNGVhM2Y3ZTQ5YjUzIiwidHJhbnNhY3Rpb25JbmRleCI6MCwiYmxvY2tIYXNoIjoiMHhiYWU5OWNiYTZlOTA4YTZhYTJmZmY1ZWY5ZjgyNmFjODFhMzBkZGY5NzE3YzQ5NjZhMWMwZTljNjEyYjg0ODE0IiwibG9nSW5kZXgiOjB9",
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
                                "value": "MHhFMDg5ZDFhMGQxNTFBMTNEODEyZUE0ZThENjQ4ZjFEYzU4QjQxY2E5",
                                "index": true
                            },
                            {
                                "key": "dHhUeXBl",
                                "value": "Mg==",
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
                        "value": "MC4wMDAwMDAwMDAwMTE4NTQ5OTk=",
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
                        "value": "dGNyY3ZhbG9wZXIxMjc3dGpmMnB3MmtkNmsyNXgzc3loZms3NXltdmV1NWgzcHJtd3M=",
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
                        "value": "dGNyY3ZhbG9wZXIxMjc3dGpmMnB3MmtkNmsyNXgzc3loZms3NXltdmV1NWgzcHJtd3M=",
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
            },
            {
                "type": "fee_market",
                "attributes": [
                    {
                        "key": "YmFzZV9mZWU=",
                        "value": "MTk1NTM0NzQzODEzMA==",
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
                        "value": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABQAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAEAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
                        "index": true
                    }
                ]
            },
            {
                "type": "block_gas",
                "attributes": [
                    {
                        "key": "aGVpZ2h0",
                        "value": "NTE2ODMxMQ==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "NDQzOTU=",
                        "index": true
                    }
                ]
            }
        ],
        "validator_updates": null,
        "consensus_param_updates": {
            "block": {
                "max_bytes": "1048576",
                "max_gas": "40000000"
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
}
`

const TX_MSG_EXTENSION_OPTION_DYNAMIC_FEE_TXS_RESP = `
{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/ethermint.evm.v1.MsgEthereumTx",
                    "data": {
                        "@type": "/ethermint.evm.v1.DynamicFeeTx",
                        "chain_id": "338",
                        "nonce": "115",
                        "gas_tip_cap": "5000000000000",
                        "gas_fee_cap": "5000000000000",
                        "gas": "48834",
                        "to": "0x904Bd5a5AAC0B9d88A0D47864724218986Ad4a3a",
                        "value": "0",
                        "data": "CV6nswAAAAAAAAAAAAAAAPVk7wA0u3182EQSJ14eUTZNTXc0//////////////////////////////////////////8=",
                        "accesses": [],
                        "v": null,
                        "r": "HHigR02IN48gvEM991C/MOCmbVfOw9Sj804IhtI5z2Q=",
                        "s": "SspHmiOAoM4bQdvAig9I8PQzLv5GTzElhcROppO0WpU="
                    },
                    "size": 0,
                    "hash": "0xd56db65672a48daaf14f743f53c316a3eba130d72beb7034296f4ea3f7e49b53",
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
            "non_critical_extension_options": []
        },
        "auth_info": {
            "signer_infos": [],
            "fee": {
                "amount": [
                    {
                        "denom": "basetcro",
                        "amount": "244170000000000000"
                    }
                ],
                "gas_limit": "48834",
                "payer": "",
                "granter": ""
            },
            "tip": null
        },
        "signatures": []
    },
    "tx_response": {
        "height": "5168311",
        "txhash": "3F1B98E6A43A3666699CA28DA2A0872E1478E092F7EC3FCF8A94202BB8B330D2",
        "codespace": "",
        "code": 0,
        "data": "12C0040A272F65746865726D696E742E65766D2E76312E4D7367457468657265756D5478526573706F6E73651294040A4230786435366462363536373261343864616166313466373433663533633331366133656261313330643732626562373033343239366634656133663765343962353312A7030A2A3078393034426435613541414330423964383841304434373836343732343231383938364164346133611242307838633562653165356562656337643562643134663731343237643165383466336464303331346330663762323239316535623230306163386337633362393235124230783030303030303030303030303030303030303030303030306530383964316130643135316131336438313265613465386436343866316463353862343163613912423078303030303030303030303030303030303030303030303030663536346566303033346262376437636438343431323237356531653531333634643464373733341A20FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF20B7B9BB022A423078643536646236353637326134386461616631346637343366353363333136613365626131333064373262656237303334323936663465613366376534396235333A423078626165393963626136653930386136616132666666356566396638323661633831613330646466393731376334393636613163306539633631326238343831341A20000000000000000000000000000000000000000000000000000000000000000128EBDA02",
        "raw_log": "[{\"msg_index\":0,\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"tcrc1uzyargx32xsnmqfw5n5dvj83m3vtg89fhjges3\"},{\"key\":\"amount\",\"value\":\"22195000000000000basetcro\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"tcrc17xpfvakm2amg962yls6f84z3kell8c5lfjsjej\"},{\"key\":\"amount\",\"value\":\"22195000000000000basetcro\"}]},{\"type\":\"ethereum_tx\",\"attributes\":[{\"key\":\"amount\",\"value\":\"0\"},{\"key\":\"ethereumTxHash\",\"value\":\"0xd56db65672a48daaf14f743f53c316a3eba130d72beb7034296f4ea3f7e49b53\"},{\"key\":\"txIndex\",\"value\":\"0\"},{\"key\":\"txGasUsed\",\"value\":\"44395\"},{\"key\":\"txHash\",\"value\":\"3F1B98E6A43A3666699CA28DA2A0872E1478E092F7EC3FCF8A94202BB8B330D2\"},{\"key\":\"recipient\",\"value\":\"0x904Bd5a5AAC0B9d88A0D47864724218986Ad4a3a\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/ethermint.evm.v1.EthermintLegacyTx\"},{\"key\":\"sender\",\"value\":\"tcrc17xpfvakm2amg962yls6f84z3kell8c5lfjsjej\"},{\"key\":\"module\",\"value\":\"evm\"},{\"key\":\"sender\",\"value\":\"0xE089d1a0d151A13D812eA4e8D648f1Dc58B41ca9\"},{\"key\":\"txType\",\"value\":\"2\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcrc1uzyargx32xsnmqfw5n5dvj83m3vtg89fhjges3\"},{\"key\":\"sender\",\"value\":\"tcrc17xpfvakm2amg962yls6f84z3kell8c5lfjsjej\"},{\"key\":\"amount\",\"value\":\"22195000000000000basetcro\"}]},{\"type\":\"tx_log\",\"attributes\":[{\"key\":\"txLog\",\"value\":\"{\\\"address\\\":\\\"0x904Bd5a5AAC0B9d88A0D47864724218986Ad4a3a\\\",\\\"topics\\\":[\\\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\\\",\\\"0x000000000000000000000000e089d1a0d151a13d812ea4e8d648f1dc58b41ca9\\\",\\\"0x000000000000000000000000f564ef0034bb7d7cd84412275e1e51364d4d7734\\\"],\\\"data\\\":\\\"//////////////////////////////////////////8=\\\",\\\"blockNumber\\\":5168311,\\\"transactionHash\\\":\\\"0xd56db65672a48daaf14f743f53c316a3eba130d72beb7034296f4ea3f7e49b53\\\",\\\"transactionIndex\\\":0,\\\"blockHash\\\":\\\"0xbae99cba6e908a6aa2fff5ef9f826ac81a30ddf9717c4966a1c0e9c612b84814\\\",\\\"logIndex\\\":0}\"}]}]}]",
        "logs": [
            {
                "msg_index": 0,
                "log": "",
                "events": [
                    {
                        "type": "coin_received",
                        "attributes": [
                            {
                                "key": "receiver",
                                "value": "tcrc1uzyargx32xsnmqfw5n5dvj83m3vtg89fhjges3"
                            },
                            {
                                "key": "amount",
                                "value": "22195000000000000basetcro"
                            }
                        ]
                    },
                    {
                        "type": "coin_spent",
                        "attributes": [
                            {
                                "key": "spender",
                                "value": "tcrc17xpfvakm2amg962yls6f84z3kell8c5lfjsjej"
                            },
                            {
                                "key": "amount",
                                "value": "22195000000000000basetcro"
                            }
                        ]
                    },
                    {
                        "type": "ethereum_tx",
                        "attributes": [
                            {
                                "key": "amount",
                                "value": "0"
                            },
                            {
                                "key": "ethereumTxHash",
                                "value": "0xd56db65672a48daaf14f743f53c316a3eba130d72beb7034296f4ea3f7e49b53"
                            },
                            {
                                "key": "txIndex",
                                "value": "0"
                            },
                            {
                                "key": "txGasUsed",
                                "value": "44395"
                            },
                            {
                                "key": "txHash",
                                "value": "3F1B98E6A43A3666699CA28DA2A0872E1478E092F7EC3FCF8A94202BB8B330D2"
                            },
                            {
                                "key": "recipient",
                                "value": "0x904Bd5a5AAC0B9d88A0D47864724218986Ad4a3a"
                            }
                        ]
                    },
                    {
                        "type": "message",
                        "attributes": [
                            {
                                "key": "action",
                                "value": "/ethermint.evm.v1.EthermintLegacyTx"
                            },
                            {
                                "key": "sender",
                                "value": "tcrc17xpfvakm2amg962yls6f84z3kell8c5lfjsjej"
                            },
                            {
                                "key": "module",
                                "value": "evm"
                            },
                            {
                                "key": "sender",
                                "value": "0xE089d1a0d151A13D812eA4e8D648f1Dc58B41ca9"
                            },
                            {
                                "key": "txType",
                                "value": "2"
                            }
                        ]
                    },
                    {
                        "type": "transfer",
                        "attributes": [
                            {
                                "key": "recipient",
                                "value": "tcrc1uzyargx32xsnmqfw5n5dvj83m3vtg89fhjges3"
                            },
                            {
                                "key": "sender",
                                "value": "tcrc17xpfvakm2amg962yls6f84z3kell8c5lfjsjej"
                            },
                            {
                                "key": "amount",
                                "value": "22195000000000000basetcro"
                            }
                        ]
                    },
                    {
                        "type": "tx_log",
                        "attributes": [
                            {
                                "key": "txLog",
                                "value": "{\"address\":\"0x904Bd5a5AAC0B9d88A0D47864724218986Ad4a3a\",\"topics\":[\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\",\"0x000000000000000000000000e089d1a0d151a13d812ea4e8d648f1dc58b41ca9\",\"0x000000000000000000000000f564ef0034bb7d7cd84412275e1e51364d4d7734\"],\"data\":\"//////////////////////////////////////////8=\",\"blockNumber\":5168311,\"transactionHash\":\"0xd56db65672a48daaf14f743f53c316a3eba130d72beb7034296f4ea3f7e49b53\",\"transactionIndex\":0,\"blockHash\":\"0xbae99cba6e908a6aa2fff5ef9f826ac81a30ddf9717c4966a1c0e9c612b84814\",\"logIndex\":0}"
                            }
                        ]
                    }
                ]
            }
        ],
        "info": "",
        "gas_wanted": "48834",
        "gas_used": "44395",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/ethermint.evm.v1.EthermintLegacyTx",
                        "data": {
                            "@type": "/ethermint.evm.v1.DynamicFeeTx",
                            "chain_id": "338",
                            "nonce": "115",
                            "gas_tip_cap": "5000000000000",
                            "gas_fee_cap": "5000000000000",
                            "gas": "48834",
                            "to": "0x904Bd5a5AAC0B9d88A0D47864724218986Ad4a3a",
                            "value": "0",
                            "data": "CV6nswAAAAAAAAAAAAAAAPVk7wA0u3182EQSJ14eUTZNTXc0//////////////////////////////////////////8=",
                            "accesses": [],
                            "v": null,
                            "r": "HHigR02IN48gvEM991C/MOCmbVfOw9Sj804IhtI5z2Q=",
                            "s": "SspHmiOAoM4bQdvAig9I8PQzLv5GTzElhcROppO0WpU="
                        },
                        "size": 0,
                        "hash": "0xd56db65672a48daaf14f743f53c316a3eba130d72beb7034296f4ea3f7e49b53",
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
                "non_critical_extension_options": []
            },
            "auth_info": {
                "signer_infos": [],
                "fee": {
                    "amount": [
                        {
                            "denom": "basetcro",
                            "amount": "244170000000000000"
                        }
                    ],
                    "gas_limit": "48834",
                    "payer": "",
                    "granter": ""
                },
                "tip": null
            },
            "signatures": []
        },
        "timestamp": "2022-09-16T03:14:55Z",
        "events": [
            {
                "type": "coin_spent",
                "attributes": [
                    {
                        "key": "c3BlbmRlcg==",
                        "value": "dGNyYzF1enlhcmd4MzJ4c25tcWZ3NW41ZHZqODNtM3Z0Zzg5ZmhqZ2VzMw==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MjQ0MTcwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
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
                        "value": "MjQ0MTcwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
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
                        "value": "dGNyYzF1enlhcmd4MzJ4c25tcWZ3NW41ZHZqODNtM3Z0Zzg5ZmhqZ2VzMw==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MjQ0MTcwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
                        "index": true
                    }
                ]
            },
            {
                "type": "message",
                "attributes": [
                    {
                        "key": "c2VuZGVy",
                        "value": "dGNyYzF1enlhcmd4MzJ4c25tcWZ3NW41ZHZqODNtM3Z0Zzg5ZmhqZ2VzMw==",
                        "index": true
                    }
                ]
            },
            {
                "type": "tx",
                "attributes": [
                    {
                        "key": "ZmVl",
                        "value": "MjQ0MTcwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
                        "index": true
                    }
                ]
            },
            {
                "type": "ethereum_tx",
                "attributes": [
                    {
                        "key": "ZXRoZXJldW1UeEhhc2g=",
                        "value": "MHhkNTZkYjY1NjcyYTQ4ZGFhZjE0Zjc0M2Y1M2MzMTZhM2ViYTEzMGQ3MmJlYjcwMzQyOTZmNGVhM2Y3ZTQ5YjUz",
                        "index": true
                    },
                    {
                        "key": "dHhJbmRleA==",
                        "value": "MA==",
                        "index": true
                    }
                ]
            },
            {
                "type": "message",
                "attributes": [
                    {
                        "key": "YWN0aW9u",
                        "value": "L2V0aGVybWludC5ldm0udjEuTXNnRXRoZXJldW1UeA==",
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
                        "value": "MjIxOTUwMDAwMDAwMDAwMDBiYXNldGNybw==",
                        "index": true
                    }
                ]
            },
            {
                "type": "coin_received",
                "attributes": [
                    {
                        "key": "cmVjZWl2ZXI=",
                        "value": "dGNyYzF1enlhcmd4MzJ4c25tcWZ3NW41ZHZqODNtM3Z0Zzg5ZmhqZ2VzMw==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MjIxOTUwMDAwMDAwMDAwMDBiYXNldGNybw==",
                        "index": true
                    }
                ]
            },
            {
                "type": "transfer",
                "attributes": [
                    {
                        "key": "cmVjaXBpZW50",
                        "value": "dGNyYzF1enlhcmd4MzJ4c25tcWZ3NW41ZHZqODNtM3Z0Zzg5ZmhqZ2VzMw==",
                        "index": true
                    },
                    {
                        "key": "c2VuZGVy",
                        "value": "dGNyYzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bGZqc2plag==",
                        "index": true
                    },
                    {
                        "key": "YW1vdW50",
                        "value": "MjIxOTUwMDAwMDAwMDAwMDBiYXNldGNybw==",
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
                "type": "ethereum_tx",
                "attributes": [
                    {
                        "key": "YW1vdW50",
                        "value": "MA==",
                        "index": true
                    },
                    {
                        "key": "ZXRoZXJldW1UeEhhc2g=",
                        "value": "MHhkNTZkYjY1NjcyYTQ4ZGFhZjE0Zjc0M2Y1M2MzMTZhM2ViYTEzMGQ3MmJlYjcwMzQyOTZmNGVhM2Y3ZTQ5YjUz",
                        "index": true
                    },
                    {
                        "key": "dHhJbmRleA==",
                        "value": "MA==",
                        "index": true
                    },
                    {
                        "key": "dHhHYXNVc2Vk",
                        "value": "NDQzOTU=",
                        "index": true
                    },
                    {
                        "key": "dHhIYXNo",
                        "value": "M0YxQjk4RTZBNDNBMzY2NjY5OUNBMjhEQTJBMDg3MkUxNDc4RTA5MkY3RUMzRkNGOEE5NDIwMkJCOEIzMzBEMg==",
                        "index": true
                    },
                    {
                        "key": "cmVjaXBpZW50",
                        "value": "MHg5MDRCZDVhNUFBQzBCOWQ4OEEwRDQ3ODY0NzI0MjE4OTg2QWQ0YTNh",
                        "index": true
                    }
                ]
            },
            {
                "type": "tx_log",
                "attributes": [
                    {
                        "key": "dHhMb2c=",
                        "value": "eyJhZGRyZXNzIjoiMHg5MDRCZDVhNUFBQzBCOWQ4OEEwRDQ3ODY0NzI0MjE4OTg2QWQ0YTNhIiwidG9waWNzIjpbIjB4OGM1YmUxZTVlYmVjN2Q1YmQxNGY3MTQyN2QxZTg0ZjNkZDAzMTRjMGY3YjIyOTFlNWIyMDBhYzhjN2MzYjkyNSIsIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwZTA4OWQxYTBkMTUxYTEzZDgxMmVhNGU4ZDY0OGYxZGM1OGI0MWNhOSIsIjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwZjU2NGVmMDAzNGJiN2Q3Y2Q4NDQxMjI3NWUxZTUxMzY0ZDRkNzczNCJdLCJkYXRhIjoiLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vLy8vOD0iLCJibG9ja051bWJlciI6NTE2ODMxMSwidHJhbnNhY3Rpb25IYXNoIjoiMHhkNTZkYjY1NjcyYTQ4ZGFhZjE0Zjc0M2Y1M2MzMTZhM2ViYTEzMGQ3MmJlYjcwMzQyOTZmNGVhM2Y3ZTQ5YjUzIiwidHJhbnNhY3Rpb25JbmRleCI6MCwiYmxvY2tIYXNoIjoiMHhiYWU5OWNiYTZlOTA4YTZhYTJmZmY1ZWY5ZjgyNmFjODFhMzBkZGY5NzE3YzQ5NjZhMWMwZTljNjEyYjg0ODE0IiwibG9nSW5kZXgiOjB9",
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
                        "value": "MHhFMDg5ZDFhMGQxNTFBMTNEODEyZUE0ZThENjQ4ZjFEYzU4QjQxY2E5",
                        "index": true
                    },
                    {
                        "key": "dHhUeXBl",
                        "value": "Mg==",
                        "index": true
                    }
                ]
            }
        ]
    }
}
`
