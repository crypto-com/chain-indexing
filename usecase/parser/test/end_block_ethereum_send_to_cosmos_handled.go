package usecase_parser_test

const END_BLOCK_ETHEREUM_SEND_TO_COSMOS_HANDLED_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "630",
    "txs_results": [
      {
        "code": 0,
        "data": "CiQKIi9ncmF2aXR5LnYxLk1zZ1N1Ym1pdEV0aGVyZXVtRXZlbnQ=",
        "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/gravity.v1.MsgSubmitEthereumEvent\"},{\"key\":\"module\",\"value\":\"*types.SendToCosmosEvent\"},{\"key\":\"ethereum_event_vote_record_id\",\"value\":\"\\u0005\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0006\\ufffd\\ufffd\\ufffdGO\\ufffdՄ\\ufffdh\\ufffd86\\ufffd%\\ufffd\\ufffd1,;g\\ufffd\\u003e\\ufffdU\\ufffd\\u0010I\\ufffd1\\ufffd\\ufffd\"}]}]}]",
        "info": "",
        "gas_wanted": "500000",
        "gas_used": "92515",
        "events": [
          {
            "type": "tx",
            "attributes": [
              {
                "key": "YWNjX3NlcQ==",
                "value": "dGNyYzEwZTRkc3R5bGFrbmh4eGw4eGZrdHczMnpqOXV2NG53MnQzM2xhMy82",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "c2lnbmF0dXJl",
                "value": "UE1SSTV4T3FpMXNXUmVqdDFhb3B4dVlGSytlVzR0ZFJ2UWhNalVyK1hsTXVBL2Y3Um9Mek5EaWdZREFFV09JcHhWWUY2VU8yRlpGTjkrZThhSUsvc2c9PQ==",
                "index": true
              }
            ]
          },
          {
            "type": "coin_spent",
            "attributes": [
              {
                "key": "c3BlbmRlcg==",
                "value": "dGNyYzEwZTRkc3R5bGFrbmh4eGw4eGZrdHczMnpqOXV2NG53MnQzM2xhMw==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MjUwMDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
                "value": "MjUwMDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
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
                "value": "dGNyYzEwZTRkc3R5bGFrbmh4eGw4eGZrdHczMnpqOXV2NG53MnQzM2xhMw==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MjUwMDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "c2VuZGVy",
                "value": "dGNyYzEwZTRkc3R5bGFrbmh4eGw4eGZrdHczMnpqOXV2NG53MnQzM2xhMw==",
                "index": true
              }
            ]
          },
          {
            "type": "tx",
            "attributes": [
              {
                "key": "ZmVl",
                "value": "MjUwMDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "L2dyYXZpdHkudjEuTXNnU3VibWl0RXRoZXJldW1FdmVudA==",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "KnR5cGVzLlNlbmRUb0Nvc21vc0V2ZW50",
                "index": true
              },
              {
                "key": "ZXRoZXJldW1fZXZlbnRfdm90ZV9yZWNvcmRfaWQ=",
                "value": "BQAAAAAAAAAGwJnOR0/u1YSdaJ44Nv8loZAxLDtntz7EVYgQSYcx8qc=",
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
            "value": "MC4wMDAwMDAwMDAwMDA0OTk5OTk=",
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
            "value": "dGNyY3ZhbG9wZXIxZHpsaGF1cHR5ZHdzcmo0Y3d1Y25mZHhrcHIzcHdwNDAyNDZwNnc=",
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
            "value": "dGNyY3ZhbG9wZXIxZHpsaGF1cHR5ZHdzcmo0Y3d1Y25mZHhrcHIzcHdwNDAyNDZwNnc=",
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
            "value": "dGNyY3ZhbG9wZXIxZHpsaGF1cHR5ZHdzcmo0Y3d1Y25mZHhrcHIzcHdwNDAyNDZwNnc=",
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
            "value": "dGNyY3ZhbG9wZXIxZHpsaGF1cHR5ZHdzcmo0Y3d1Y25mZHhrcHIzcHdwNDAyNDZwNnc=",
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
            "value": "dGNyY3ZhbG9wZXIxZHpsaGF1cHR5ZHdzcmo0Y3d1Y25mZHhrcHIzcHdwNDAyNDZwNnc=",
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
            "value": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNyYzE2bjNsYzdjeXdhNjhtZzUwcWhwODQ3MDM0dzg4cG50cWxhc3U4dQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjUwZ3Jhdml0eTB4NTY0QTFjM0FGMDg5RDAyRDBCNkMzMTFDNjUwZUEzNzY4NDI0Y2JmYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coinbase",
        "attributes": [
          {
            "key": "bWludGVy",
            "value": "dGNyYzE2bjNsYzdjeXdhNjhtZzUwcWhwODQ3MDM0dzg4cG50cWxhc3U4dQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjUwZ3Jhdml0eTB4NTY0QTFjM0FGMDg5RDAyRDBCNkMzMTFDNjUwZUEzNzY4NDI0Y2JmYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNyYzE2bjNsYzdjeXdhNjhtZzUwcWhwODQ3MDM0dzg4cG50cWxhc3U4dQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjUwZ3Jhdml0eTB4NTY0QTFjM0FGMDg5RDAyRDBCNkMzMTFDNjUwZUEzNzY4NDI0Y2JmYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNyYzEzeXV4Nno4bWg2dzV0M3Y0dXE3Y2xld25oMzV6bnJnZGd5ZTBrMg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjUwZ3Jhdml0eTB4NTY0QTFjM0FGMDg5RDAyRDBCNkMzMTFDNjUwZUEzNzY4NDI0Y2JmYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNyYzEzeXV4Nno4bWg2dzV0M3Y0dXE3Y2xld25oMzV6bnJnZGd5ZTBrMg==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNyYzE2bjNsYzdjeXdhNjhtZzUwcWhwODQ3MDM0dzg4cG50cWxhc3U4dQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjUwZ3Jhdml0eTB4NTY0QTFjM0FGMDg5RDAyRDBCNkMzMTFDNjUwZUEzNzY4NDI0Y2JmYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNyYzE2bjNsYzdjeXdhNjhtZzUwcWhwODQ3MDM0dzg4cG50cWxhc3U4dQ==",
            "index": true
          }
        ]
      },
      {
        "type": "ethereum_send_to_cosmos_handled",
        "attributes": [
          {
            "key": "bW9kdWxl",
            "value": "Z3Jhdml0eQ==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "MHg1RTQ0RDQzRjRBYTBCM0VEMDA0ZWFhRDRlRjIxYTgzREZGMmVmNkU1",
            "index": true
          },
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNyYzEzeXV4Nno4bWg2dzV0M3Y0dXE3Y2xld25oMzV6bnJnZGd5ZTBrMg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjUwZ3Jhdml0eTB4NTY0QTFjM0FGMDg5RDAyRDBCNkMzMTFDNjUwZUEzNzY4NDI0Y2JmYQ==",
            "index": true
          },
          {
            "key": "YnJpZGdlX2NoYWluX2lk",
            "value": "NDI=",
            "index": true
          },
          {
            "key": "ZXRoZXJldW1fdG9rZW5fY29udHJhY3Q=",
            "value": "MHg1NjRBMWMzQUYwODlEMDJEMEI2QzMxMUM2NTBlQTM3Njg0MjRjYmZh",
            "index": true
          },
          {
            "key": "bm9uY2U=",
            "value": "Ng==",
            "index": true
          },
          {
            "key": "ZXRoZXJldW1fZXZlbnRfdm90ZV9yZWNvcmRfaWQ=",
            "value": "BQAAAAAAAAAGwJnOR0/u1YSdaJ44Nv8loZAxLDtntz7EVYgQSYcx8qc=",
            "index": true
          }
        ]
      },
      {
        "type": "coin_spent",
        "attributes": [
          {
            "key": "c3BlbmRlcg==",
            "value": "dGNyYzEzeXV4Nno4bWg2dzV0M3Y0dXE3Y2xld25oMzV6bnJnZGd5ZTBrMg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjUwZ3Jhdml0eTB4NTY0QTFjM0FGMDg5RDAyRDBCNkMzMTFDNjUwZUEzNzY4NDI0Y2JmYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "coin_received",
        "attributes": [
          {
            "key": "cmVjZWl2ZXI=",
            "value": "dGNyYzE2d3c1eWh1MDJzZGM3NGE0MjZndTB3ODAwbGVxa2d3a2xqcHNwZg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjUwZ3Jhdml0eTB4NTY0QTFjM0FGMDg5RDAyRDBCNkMzMTFDNjUwZUEzNzY4NDI0Y2JmYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNyYzE2d3c1eWh1MDJzZGM3NGE0MjZndTB3ODAwbGVxa2d3a2xqcHNwZg==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNyYzEzeXV4Nno4bWg2dzV0M3Y0dXE3Y2xld25oMzV6bnJnZGd5ZTBrMg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjUwZ3Jhdml0eTB4NTY0QTFjM0FGMDg5RDAyRDBCNkMzMTFDNjUwZUEzNzY4NDI0Y2JmYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNyYzEzeXV4Nno4bWg2dzV0M3Y0dXE3Y2xld25oMzV6bnJnZGd5ZTBrMg==",
            "index": true
          }
        ]
      },
      {
        "type": "observation",
        "attributes": [
          {
            "key": "bW9kdWxl",
            "value": "Z3Jhdml0eQ==",
            "index": true
          },
          {
            "key": "ZXRoZXJldW1fZXZlbnRfdHlwZQ==",
            "value": "KnR5cGVzLlNlbmRUb0Nvc21vc0V2ZW50",
            "index": true
          },
          {
            "key": "YnJpZGdlX2NvbnRyYWN0",
            "value": "MHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAw",
            "index": true
          },
          {
            "key": "YnJpZGdlX2NoYWluX2lk",
            "value": "NDI=",
            "index": true
          },
          {
            "key": "ZXRoZXJldW1fZXZlbnRfdm90ZV9yZWNvcmRfaWQ=",
            "value": "BQAAAAAAAAAGwJnOR0/u1YSdaJ44Nv8loZAxLDtntz7EVYgQSYcx8qc=",
            "index": true
          },
          {
            "key": "bm9uY2U=",
            "value": "Ng==",
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
