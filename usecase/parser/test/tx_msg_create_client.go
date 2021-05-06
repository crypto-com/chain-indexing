package usecase_parser_test

const TX_MSG_CREATE_CLIENT_BLOCK_RESP = `{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
      "block_id": {
        "hash": "121B48440580D4384DFC6CF9D17444BEB5C7D7E293902F7B3D7513DA8DC44E39",
        "parts": {
          "total": 1,
          "hash": "CBD380F934F8A87A8DCB8BB6C3F1BFD824FC168DE127EDF951F8C135A91E2560"
        }
      },
      "block": {
        "header": {
          "version": {
            "block": "11"
          },
          "chain_id": "devnet-1",
          "height": "5",
          "time": "2021-05-04T18:02:43.146958Z",
          "last_block_id": {
            "hash": "63D8678B08EF82C9348063E96A56D5460A777E4466A21F92DD8AEB6E3C86A485",
            "parts": {
              "total": 1,
              "hash": "1A0B7FC37A15A32F07139A122585392AB1BFBFCB1DCEDCCF82654255B7A82B88"
            }
          },
          "last_commit_hash": "B6561327E9076BCBAD0D2FE2401BB20D61A8DFD7CD0F24BB5D406D97D20309DD",
          "data_hash": "7F9FFF9C83858216169E31902C6B656244A067BF4C0DF65BC23958A99E8A226A",
          "validators_hash": "3F4AEF8603CD6582C072B7EACBBF1BDD9591996AF659617CC72683B969BE8DF5",
          "next_validators_hash": "3F4AEF8603CD6582C072B7EACBBF1BDD9591996AF659617CC72683B969BE8DF5",
          "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
          "app_hash": "CDAA14AFC04657F25998983D02CBEC6A1A4CFBDCC8B44A363F3B3552B291EECD",
          "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
          "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
          "proposer_address": "CD1AC8BFC2C01FC3C2DAF26628D1CC820FD0CF6C"
        },
        "data": {
          "txs": [
            "CooDCocDCiMvaWJjLmNvcmUuY2xpZW50LnYxLk1zZ0NyZWF0ZUNsaWVudBLfAgqoAQorL2liYy5saWdodGNsaWVudHMudGVuZGVybWludC52MS5DbGllbnRTdGF0ZRJ5CghkZXZuZXQtMhIECAEQAxoECIDqSSIECIDfbioCCAUyADoECAIQAkIZCgkIARgBIAEqAQASDAoCAAEQIRgEIAwwAUIZCgkIARgBIAEqAQASDAoCAAEQIBgBIAEwAUoHdXBncmFkZUoQdXBncmFkZWRJQkNTdGF0ZRKFAQouL2liYy5saWdodGNsaWVudHMudGVuZGVybWludC52MS5Db25zZW5zdXNTdGF0ZRJTCgsIvJjGhAYQ8KzTKhIiCiBtWOJDb35XnbY1UV3HUGJUmcgElXxjWlh72my/FkUBaBog494NKzI3oC6cIMNPnuBPafWGH7wuJyKgEcqQN/xnp+waKmNybzFnZHN3cm13dHpndjNrdmYyOGx2dHQ3cXY3cTdteXptbjQ2NnIzZhJoCk4KRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDLcXkU5CJ/wyuELKl1oyeLMLDjMqHdZkwHo5h0PTj54gSBAoCCAESFgoPCgdiYXNlY3JvEgQxMDAwEMCNtwEaQPsQZNEJFh4OrCaN9BDEgw6ZvtnNluD6+euZLAW0PCQLflcG2HSCAiNJkFF5Ebvb3gjKK30jSz3LLkR9LPtU5+k="
          ]
        },
        "evidence": {
          "evidence": []
        },
        "last_commit": {
          "height": "4",
          "round": 0,
          "block_id": {
            "hash": "63D8678B08EF82C9348063E96A56D5460A777E4466A21F92DD8AEB6E3C86A485",
            "parts": {
              "total": 1,
              "hash": "1A0B7FC37A15A32F07139A122585392AB1BFBFCB1DCEDCCF82654255B7A82B88"
            }
          },
          "signatures": [
            {
              "block_id_flag": 2,
              "validator_address": "CD1AC8BFC2C01FC3C2DAF26628D1CC820FD0CF6C",
              "timestamp": "2021-05-04T18:02:43.146958Z",
              "signature": "/y94bsvn2Eg/lVCTs1u9BRmgJMGJz5Pko3Kw6rWRHaBcSiEmR7KSuaXBEeX8rIQhDm5olTrQrS2A8Fa6+EE6AQ=="
            }
          ]
        }
      }
    }
}`

const TX_MSG_CREATE_CLIENT_BLOCK_RESULTS_RESP = `
{
    "jsonrpc": "2.0",
    "id": -1,
    "result": {
      "height": "5",
      "txs_results": [
        {
          "code": 0,
          "data": "Cg8KDWNyZWF0ZV9jbGllbnQ=",
          "log": "[{\"events\":[{\"type\":\"create_client\",\"attributes\":[{\"key\":\"client_id\",\"value\":\"07-tendermint-0\"},{\"key\":\"client_type\",\"value\":\"07-tendermint\"},{\"key\":\"consensus_height\",\"value\":\"2-2\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"create_client\"},{\"key\":\"module\",\"value\":\"ibc_client\"}]}]}]",
          "info": "",
          "gas_wanted": "3000000",
          "gas_used": "78510",
          "events": [
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
                  "value": "Y3JvMWdkc3dybXd0emd2M2t2ZjI4bHZ0dDdxdjdxN215em1uNDY2cjNm",
                  "index": true
                },
                {
                  "key": "YW1vdW50",
                  "value": "MTAwMGJhc2Vjcm8=",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "c2VuZGVy",
                  "value": "Y3JvMWdkc3dybXd0emd2M2t2ZjI4bHZ0dDdxdjdxN215em1uNDY2cjNm",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "YWN0aW9u",
                  "value": "Y3JlYXRlX2NsaWVudA==",
                  "index": true
                }
              ]
            },
            {
              "type": "create_client",
              "attributes": [
                {
                  "key": "Y2xpZW50X2lk",
                  "value": "MDctdGVuZGVybWludC0w",
                  "index": true
                },
                {
                  "key": "Y2xpZW50X3R5cGU=",
                  "value": "MDctdGVuZGVybWludA==",
                  "index": true
                },
                {
                  "key": "Y29uc2Vuc3VzX2hlaWdodA==",
                  "value": "Mi0y",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "bW9kdWxl",
                  "value": "aWJjX2NsaWVudA==",
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
              "value": "MjA2MTc4NzIzNjgxNTBiYXNlY3Jv",
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
              "value": "MC4wMDAwMDk5OTAwMDkxNjY5NDI=",
              "index": true
            },
            {
              "key": "aW5mbGF0aW9u",
              "value": "MC4xMzAwMDAxMDI5ODQ3NDk4MjI=",
              "index": true
            },
            {
              "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
              "value": "MTMwMTMwMTEzODA5MDMxODk3ODM5LjUzOTI4MTcwNDM0MTY4NTUyMg==",
              "index": true
            },
            {
              "key": "YW1vdW50",
              "value": "MjA2MTc4NzIzNjgxNTA=",
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
              "value": "MjA2MTc4NzIzNjgxNTBiYXNlY3Jv",
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
              "value": "MTAzMDg5MzYxODQwNy41MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
              "index": true
            },
            {
              "key": "dmFsaWRhdG9y",
              "value": "Y3JvY25jbDE0OTQzczd3OTBmNno2MzMzeHZ0OTJyYTd3Zjc4NTh6ZnVsNWUydg==",
              "index": true
            }
          ]
        },
        {
          "type": "commission",
          "attributes": [
            {
              "key": "YW1vdW50",
              "value": "MTAzMDg5MzYxODQwLjc1MDAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
              "index": true
            },
            {
              "key": "dmFsaWRhdG9y",
              "value": "Y3JvY25jbDE0OTQzczd3OTBmNno2MzMzeHZ0OTJyYTd3Zjc4NTh6ZnVsNWUydg==",
              "index": true
            }
          ]
        },
        {
          "type": "rewards",
          "attributes": [
            {
              "key": "YW1vdW50",
              "value": "MTAzMDg5MzYxODQwNy41MDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
              "index": true
            },
            {
              "key": "dmFsaWRhdG9y",
              "value": "Y3JvY25jbDE0OTQzczd3OTBmNno2MzMzeHZ0OTJyYTd3Zjc4NTh6ZnVsNWUydg==",
              "index": true
            }
          ]
        },
        {
          "type": "commission",
          "attributes": [
            {
              "key": "YW1vdW50",
              "value": "MTkxNzQ2MjEzMDIzNy45NTAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
              "index": true
            },
            {
              "key": "dmFsaWRhdG9y",
              "value": "Y3JvY25jbDE0OTQzczd3OTBmNno2MzMzeHZ0OTJyYTd3Zjc4NTh6ZnVsNWUydg==",
              "index": true
            }
          ]
        },
        {
          "type": "rewards",
          "attributes": [
            {
              "key": "YW1vdW50",
              "value": "MTkxNzQ2MjEzMDIzNzkuNTAwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
              "index": true
            },
            {
              "key": "dmFsaWRhdG9y",
              "value": "Y3JvY25jbDE0OTQzczd3OTBmNno2MzMzeHZ0OTJyYTd3Zjc4NTh6ZnVsNWUydg==",
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
