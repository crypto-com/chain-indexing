package usecase_parser_test

const TX_MSG_CONNECTION_OPEN_ACK_BLOCK_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "C01D2D375F17BB5971F8AD477D0A883F268C4B73685B6866679B7CAD298D7DB7",
      "parts": {
        "total": 1,
        "hash": "3518E7A65B9994C72EAD6332BD9946C81C2856BD625F57845653A4224B451CAE"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "devnet-1",
        "height": "11",
        "time": "2021-07-04T09:35:10.510195Z",
        "last_block_id": {
          "hash": "D3E3E37E7C42F96FC609EB2C6918EA34775F14606E4F9C115681FA6924B85A0C",
          "parts": {
            "total": 1,
            "hash": "EB5292ACFCF47BD9F544A81BB877B81DFA22428E320161DF5F0FE6C2E2C3DEA1"
          }
        },
        "last_commit_hash": "5B8F17E19F62E4ED884EF6CA99D0E9F42F6015A2CB9184A90719EECED2149AC2",
        "data_hash": "FC2290C446B154827BB50491D9D37488DA829975D0753E52C1F409F5683741CD",
        "validators_hash": "C3BE664CFB455C44053CDFDFB7DE2A6483B6C196245455BF453D3878C580934B",
        "next_validators_hash": "C3BE664CFB455C44053CDFDFB7DE2A6483B6C196245455BF453D3878C580934B",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "060D1874639EA57BE297193BE17ECBD2A687A8A15199B7DB626F6D33819A50B5",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "FDB7CAC90633FDEC2964A941CC815B5A3B84F0F5"
      },
      "data": {
        "txs": [
          "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"
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "10",
        "round": 0,
        "block_id": {
          "hash": "D3E3E37E7C42F96FC609EB2C6918EA34775F14606E4F9C115681FA6924B85A0C",
          "parts": {
            "total": 1,
            "hash": "EB5292ACFCF47BD9F544A81BB877B81DFA22428E320161DF5F0FE6C2E2C3DEA1"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "FDB7CAC90633FDEC2964A941CC815B5A3B84F0F5",
            "timestamp": "2021-07-04T09:35:10.510195Z",
            "signature": "2fsZovkCWovi7jorK5LlMr4xoh/etxpayuiboqAlLWe7aQ5ruELn3SEhrUIE9z+Kmp9aw6i+2TPmLi58Uk1bAw=="
          }
        ]
      }
    }
  }
}
`

const TX_MSG_CONNECTION_OPEN_ACK_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "11",
    "txs_results": [
      {
        "code": 0,
        "data": "Cg8KDXVwZGF0ZV9jbGllbnQKFQoTY29ubmVjdGlvbl9vcGVuX2Fjaw==",
        "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"update_client\"},{\"key\":\"module\",\"value\":\"ibc_client\"}]},{\"type\":\"update_client\",\"attributes\":[{\"key\":\"client_id\",\"value\":\"07-tendermint-0\"},{\"key\":\"client_type\",\"value\":\"07-tendermint\"},{\"key\":\"consensus_height\",\"value\":\"2-11\"},{\"key\":\"header\",\"value\":\"0a262f6962632e6c69676874636c69656e74732e74656e6465726d696e742e76312e48656164657212ed060aca040a8e030a02080b12086465766e65742d32180b220c08ce818687061080ab9e8a012a480a2050ac16c89ba8f3c77fa60d44d0e86ecfa5555c125d5be6c2cc2ba750da918fe712240801122044236e41aab58422fef1205ea54cedc02954615bc36f3a2a4eddc02af9a16ea63220650f15f47cf48dd1b8219181206ac77ed4e95212fd73c01e9e0d11f19517449f3a20e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85542203275b04ca9c4cf5f20e5f9b4e1eb6b5c2202b4716dfe068e6cc90fb97cf1bbed4a203275b04ca9c4cf5f20e5f9b4e1eb6b5c2202b4716dfe068e6cc90fb97cf1bbed5220048091bc7ddc283f77bfbf91d73c44da58c3df8a9cbc867405d8b7f3daada22f5a20f3a1ebcf149042c7eeac12e26b8715095d009c8ff6e2667f5e9fbe56e2ebf25f6220a98f88d5b3eb8735f84695dbcb1f2944ebbfd902b292403a028bdaf050550dcc6a20e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85572140bbbee47989725c8d25c5a7087a782c7f65a587412b601080b1a480a203acdccc5e1b4512b6129c9c657da1426a27382cba3928b1e6cc2e0166e6ce9c8122408011220c1163e3a7c4422683456fb5760381255175a7ebae198b86c960b4b9b372812002268080212140bbbee47989725c8d25c5a7087a782c7f65a58741a0c08d38186870610e0a0f2bf012240caf72274021456ff6ea71140b8f5f48ac69de24306d43101db43178e3b48e8288574b6d2903abce20c7be571aabac08a01bf7e03cc9fdb489c42d1350da34c03128a010a400a140bbbee47989725c8d25c5a7087a782c7f65a587412220a20e86c93e72292018b7f1f79c16ad3088af625e2fd3104766e7795f533268622f31880c8afa02512400a140bbbee47989725c8d25c5a7087a782c7f65a587412220a20e86c93e72292018b7f1f79c16ad3088af625e2fd3104766e7795f533268622f31880c8afa0251880c8afa0251a0408021006228a010a400a140bbbee47989725c8d25c5a7087a782c7f65a587412220a20e86c93e72292018b7f1f79c16ad3088af625e2fd3104766e7795f533268622f31880c8afa02512400a140bbbee47989725c8d25c5a7087a782c7f65a587412220a20e86c93e72292018b7f1f79c16ad3088af625e2fd3104766e7795f533268622f31880c8afa0251880c8afa025\"}]}]},{\"msg_index\":1,\"events\":[{\"type\":\"connection_open_ack\",\"attributes\":[{\"key\":\"connection_id\",\"value\":\"connection-0\"},{\"key\":\"client_id\",\"value\":\"07-tendermint-0\"},{\"key\":\"counterparty_client_id\",\"value\":\"07-tendermint-0\"},{\"key\":\"counterparty_connection_id\",\"value\":\"connection-0\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"connection_open_ack\"},{\"key\":\"module\",\"value\":\"ibc_connection\"}]}]}]",
        "info": "",
        "gas_wanted": "3000000",
        "gas_used": "117761",
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
                "value": "Y3JvMTBzbmhsdmtwdWM0eGhxODJ1eWc1ZXgyZWV6bW1mNWVkNXRtcXN2",
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
                "value": "Y3JvMTBzbmhsdmtwdWM0eGhxODJ1eWc1ZXgyZWV6bW1mNWVkNXRtcXN2",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "dXBkYXRlX2NsaWVudA==",
                "index": true
              }
            ]
          },
          {
            "type": "update_client",
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
                "value": "Mi0xMQ==",
                "index": true
              },
              {
                "key": "aGVhZGVy",
                "value": "MGEyNjJmNjk2MjYzMmU2YzY5Njc2ODc0NjM2YzY5NjU2ZTc0NzMyZTc0NjU2ZTY0NjU3MjZkNjk2ZTc0MmU3NjMxMmU0ODY1NjE2NDY1NzIxMmVkMDYwYWNhMDQwYThlMDMwYTAyMDgwYjEyMDg2NDY1NzY2ZTY1NzQyZDMyMTgwYjIyMGMwOGNlODE4Njg3MDYxMDgwYWI5ZThhMDEyYTQ4MGEyMDUwYWMxNmM4OWJhOGYzYzc3ZmE2MGQ0NGQwZTg2ZWNmYTU1NTVjMTI1ZDViZTZjMmNjMmJhNzUwZGE5MThmZTcxMjI0MDgwMTEyMjA0NDIzNmU0MWFhYjU4NDIyZmVmMTIwNWVhNTRjZWRjMDI5NTQ2MTViYzM2ZjNhMmE0ZWRkYzAyYWY5YTE2ZWE2MzIyMDY1MGYxNWY0N2NmNDhkZDFiODIxOTE4MTIwNmFjNzdlZDRlOTUyMTJmZDczYzAxZTllMGQxMWYxOTUxNzQ0OWYzYTIwZTNiMGM0NDI5OGZjMWMxNDlhZmJmNGM4OTk2ZmI5MjQyN2FlNDFlNDY0OWI5MzRjYTQ5NTk5MWI3ODUyYjg1NTQyMjAzMjc1YjA0Y2E5YzRjZjVmMjBlNWY5YjRlMWViNmI1YzIyMDJiNDcxNmRmZTA2OGU2Y2M5MGZiOTdjZjFiYmVkNGEyMDMyNzViMDRjYTljNGNmNWYyMGU1ZjliNGUxZWI2YjVjMjIwMmI0NzE2ZGZlMDY4ZTZjYzkwZmI5N2NmMWJiZWQ1MjIwMDQ4MDkxYmM3ZGRjMjgzZjc3YmZiZjkxZDczYzQ0ZGE1OGMzZGY4YTljYmM4Njc0MDVkOGI3ZjNkYWFkYTIyZjVhMjBmM2ExZWJjZjE0OTA0MmM3ZWVhYzEyZTI2Yjg3MTUwOTVkMDA5YzhmZjZlMjY2N2Y1ZTlmYmU1NmUyZWJmMjVmNjIyMGE5OGY4OGQ1YjNlYjg3MzVmODQ2OTVkYmNiMWYyOTQ0ZWJiZmQ5MDJiMjkyNDAzYTAyOGJkYWYwNTA1NTBkY2M2YTIwZTNiMGM0NDI5OGZjMWMxNDlhZmJmNGM4OTk2ZmI5MjQyN2FlNDFlNDY0OWI5MzRjYTQ5NTk5MWI3ODUyYjg1NTcyMTQwYmJiZWU0Nzk4OTcyNWM4ZDI1YzVhNzA4N2E3ODJjN2Y2NWE1ODc0MTJiNjAxMDgwYjFhNDgwYTIwM2FjZGNjYzVlMWI0NTEyYjYxMjljOWM2NTdkYTE0MjZhMjczODJjYmEzOTI4YjFlNmNjMmUwMTY2ZTZjZTljODEyMjQwODAxMTIyMGMxMTYzZTNhN2M0NDIyNjgzNDU2ZmI1NzYwMzgxMjU1MTc1YTdlYmFlMTk4Yjg2Yzk2MGI0YjliMzcyODEyMDAyMjY4MDgwMjEyMTQwYmJiZWU0Nzk4OTcyNWM4ZDI1YzVhNzA4N2E3ODJjN2Y2NWE1ODc0MWEwYzA4ZDM4MTg2ODcwNjEwZTBhMGYyYmYwMTIyNDBjYWY3MjI3NDAyMTQ1NmZmNmVhNzExNDBiOGY1ZjQ4YWM2OWRlMjQzMDZkNDMxMDFkYjQzMTc4ZTNiNDhlODI4ODU3NGI2ZDI5MDNhYmNlMjBjN2JlNTcxYWFiYWMwOGEwMWJmN2UwM2NjOWZkYjQ4OWM0MmQxMzUwZGEzNGMwMzEyOGEwMTBhNDAwYTE0MGJiYmVlNDc5ODk3MjVjOGQyNWM1YTcwODdhNzgyYzdmNjVhNTg3NDEyMjIwYTIwZTg2YzkzZTcyMjkyMDE4YjdmMWY3OWMxNmFkMzA4OGFmNjI1ZTJmZDMxMDQ3NjZlNzc5NWY1MzMyNjg2MjJmMzE4ODBjOGFmYTAyNTEyNDAwYTE0MGJiYmVlNDc5ODk3MjVjOGQyNWM1YTcwODdhNzgyYzdmNjVhNTg3NDEyMjIwYTIwZTg2YzkzZTcyMjkyMDE4YjdmMWY3OWMxNmFkMzA4OGFmNjI1ZTJmZDMxMDQ3NjZlNzc5NWY1MzMyNjg2MjJmMzE4ODBjOGFmYTAyNTE4ODBjOGFmYTAyNTFhMDQwODAyMTAwNjIyOGEwMTBhNDAwYTE0MGJiYmVlNDc5ODk3MjVjOGQyNWM1YTcwODdhNzgyYzdmNjVhNTg3NDEyMjIwYTIwZTg2YzkzZTcyMjkyMDE4YjdmMWY3OWMxNmFkMzA4OGFmNjI1ZTJmZDMxMDQ3NjZlNzc5NWY1MzMyNjg2MjJmMzE4ODBjOGFmYTAyNTEyNDAwYTE0MGJiYmVlNDc5ODk3MjVjOGQyNWM1YTcwODdhNzgyYzdmNjVhNTg3NDEyMjIwYTIwZTg2YzkzZTcyMjkyMDE4YjdmMWY3OWMxNmFkMzA4OGFmNjI1ZTJmZDMxMDQ3NjZlNzc5NWY1MzMyNjg2MjJmMzE4ODBjOGFmYTAyNTE4ODBjOGFmYTAyNQ==",
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
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "Y29ubmVjdGlvbl9vcGVuX2Fjaw==",
                "index": true
              }
            ]
          },
          {
            "type": "connection_open_ack",
            "attributes": [
              {
                "key": "Y29ubmVjdGlvbl9pZA==",
                "value": "Y29ubmVjdGlvbi0w",
                "index": true
              },
              {
                "key": "Y2xpZW50X2lk",
                "value": "MDctdGVuZGVybWludC0w",
                "index": true
              },
              {
                "key": "Y291bnRlcnBhcnR5X2NsaWVudF9pZA==",
                "value": "MDctdGVuZGVybWludC0w",
                "index": true
              },
              {
                "key": "Y291bnRlcnBhcnR5X2Nvbm5lY3Rpb25faWQ=",
                "value": "Y29ubmVjdGlvbi0w",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "aWJjX2Nvbm5lY3Rpb24=",
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
            "value": "MjA2MTc4OTQ1MTYxMDZiYXNlY3Jv",
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
            "value": "MC4wMDAwMDk5OTAwMDc5MzIzNDA=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMzAwMDAyMjY1NjY0NDk2MTI=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTMwMTMwMjUzNTk2Mjk5MjUyNzU1Ljc5MTcwNjQzNzQ1MzQwMTY4MA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjA2MTc4OTQ1MTYxMDY=",
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
            "value": "MjA2MTc4OTQ1MTYxMDZiYXNlY3Jv",
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
            "value": "MTAzMDg5NDcyNTgwNS4zMDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFla3M0c3BuZjZ0dXh5cWhhZzJ6N3M3ZmZ1bHVqOHVmZmV5YWRrcw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAzMDg5NDcyNTgwLjUzMDAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFla3M0c3BuZjZ0dXh5cWhhZzJ6N3M3ZmZ1bHVqOHVmZmV5YWRrcw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAzMDg5NDcyNTgwNS4zMDAwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFla3M0c3BuZjZ0dXh5cWhhZzJ6N3M3ZmZ1bHVqOHVmZmV5YWRrcw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNzQ2NDE4OTk5Ny44NTgwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFla3M0c3BuZjZ0dXh5cWhhZzJ6N3M3ZmZ1bHVqOHVmZmV5YWRrcw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTkxNzQ2NDE4OTk5NzguNTgwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFla3M0c3BuZjZ0dXh5cWhhZzJ6N3M3ZmZ1bHVqOHVmZmV5YWRrcw==",
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

// TODO: Add Txs response
