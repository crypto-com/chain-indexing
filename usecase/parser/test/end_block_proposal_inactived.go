package usecase_parser_test

const END_BLOCK_PROPOSAL_INACTIVED_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "21541",
    "txs_results": null,
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
            "value": "MTI4MWJhc2Vjcm8=",
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
            "value": "MC4wMzIyNDM3MzgwOTI4MzI4ODg=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMzA0MjIzMjg0MTE0NjI1NTA=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "ODA4OTc3NzE4Ni4yNTMyNzIyNzM5NzE2MzEwMDA=",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTI4MQ==",
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
            "value": "MTI4MWJhc2Vjcm8=",
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
            "value": "NjQuMDUwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDF2OXE0ZHpzemphbXJtNWp0Y3AzNDZ0bTlsOGFlcThrbHVsbDUycw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Ni40MDUwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDF2OXE0ZHpzemphbXJtNWp0Y3AzNDZ0bTlsOGFlcThrbHVsbDUycw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjQuMDUwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDF2OXE0ZHpzemphbXJtNWp0Y3AzNDZ0bTlsOGFlcThrbHVsbDUycw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTkuNTY2NTAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFrOXozNjdmdDBkanV2eDN2eHRyOTR6bHRuNWY4dWd3aGFld255OA==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTk1LjY2NTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFrOXozNjdmdDBkanV2eDN2eHRyOTR6bHRuNWY4dWd3aGFld255OA==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTkuNTY2NTAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDF2OXE0ZHpzemphbXJtNWp0Y3AzNDZ0bTlsOGFlcThrbHVsbDUycw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTk1LjY2NTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDF2OXE0ZHpzemphbXJtNWp0Y3AzNDZ0bTlsOGFlcThrbHVsbDUycw==",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": [
      {
        "type": "inactive_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "Mg==",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfcmVzdWx0",
            "value": "cHJvcG9zYWxfZHJvcHBlZA==",
            "index": true
          }
        ]
      }
    ],
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
}`
