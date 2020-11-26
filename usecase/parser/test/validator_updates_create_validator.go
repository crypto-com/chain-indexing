package usecase_parser_test

var VALIDATOR_UPDATES_CREATE_VALIDATOR_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "17",
    "txs_results": [
      {
        "code": 0,
        "data": "ChIKEGNyZWF0ZV92YWxpZGF0b3I=",
        "log": "[{\"events\":[{\"type\":\"create_validator\",\"attributes\":[{\"key\":\"validator\",\"value\":\"crocncl1r3vase2ug6gwu47sc5mqvdv07gkujfmh5mgtaf\"},{\"key\":\"amount\",\"value\":\"212345678\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"create_validator\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"cro1r3vase2ug6gwu47sc5mqvdv07gkujfmhhktzl4\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "139749",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "Y3JlYXRlX3ZhbGlkYXRvcg==",
                "index": true
              }
            ]
          },
          {
            "type": "create_validator",
            "attributes": [
              {
                "key": "dmFsaWRhdG9y",
                "value": "Y3JvY25jbDFyM3Zhc2UydWc2Z3d1NDdzYzVtcXZkdjA3Z2t1amZtaDVtZ3RhZg==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MjEyMzQ1Njc4",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "c3Rha2luZw==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMXIzdmFzZTJ1ZzZnd3U0N3NjNW1xdmR2MDdna3VqZm1oaGt0emw0",
                "index": true
              }
            ]
          }
        ],
        "codespace": ""
      },
      {
        "code": 0,
        "data": "ChIKEGNyZWF0ZV92YWxpZGF0b3I=",
        "log": "[{\"events\":[{\"type\":\"create_validator\",\"attributes\":[{\"key\":\"validator\",\"value\":\"crocncl1rgefmqw2wufwffsg9528k58edgsgc4xhzsx9l2\"},{\"key\":\"amount\",\"value\":\"312345678\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"create_validator\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"cro1rgefmqw2wufwffsg9528k58edgsgc4xhpa9vak\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "139743",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "YWN0aW9u",
                "value": "Y3JlYXRlX3ZhbGlkYXRvcg==",
                "index": true
              }
            ]
          },
          {
            "type": "create_validator",
            "attributes": [
              {
                "key": "dmFsaWRhdG9y",
                "value": "Y3JvY25jbDFyZ2VmbXF3Mnd1ZndmZnNnOTUyOGs1OGVkZ3NnYzR4aHpzeDlsMg==",
                "index": true
              },
              {
                "key": "YW1vdW50",
                "value": "MzEyMzQ1Njc4",
                "index": true
              }
            ]
          },
          {
            "type": "message",
            "attributes": [
              {
                "key": "bW9kdWxl",
                "value": "c3Rha2luZw==",
                "index": true
              },
              {
                "key": "c2VuZGVy",
                "value": "Y3JvMXJnZWZtcXcyd3Vmd2Zmc2c5NTI4azU4ZWRnc2djNHhocGE5dmFr",
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
            "value": "MTI3N2Jhc2Vjcm8=",
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
            "value": "MC4wMzIyNTgwNTM4ODU1Mzk0MDM=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMzAwMDAzMzMyOTQ3NjExNjc=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "ODA2MDAyMzMyMC40NDIwMDIyMzI1NjAxNjQxNDQ=",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTI3Nw==",
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
            "value": "MTI3N2Jhc2Vjcm8=",
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
            "value": "NjMuODUwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFlbWo1Nzg2a2hxaGdoaHd1ejQ2eWdzcTVuNWpzeWczcjc5dzducw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Ni4zODUwMDAwMDAwMDAwMDAwMDBiYXNlY3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFlbWo1Nzg2a2hxaGdoaHd1ejQ2eWdzcTVuNWpzeWczcjc5dzducw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjMuODUwMDAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFlbWo1Nzg2a2hxaGdoaHd1ejQ2eWdzcTVuNWpzeWczcjc5dzducw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTkuMzgwNTAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFlbWo1Nzg2a2hxaGdoaHd1ejQ2eWdzcTVuNWpzeWczcjc5dzducw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTkzLjgwNTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFlbWo1Nzg2a2hxaGdoaHd1ejQ2eWdzcTVuNWpzeWczcjc5dzducw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTkuMzgwNTAwMDAwMDAwMDAwMDAwYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDE1ZXgwZmRtY3FuMGFseHNyZjJlamE3cG5kdzRqeTlsMnh4ajJjOQ==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NTkzLjgwNTAwMDAwMDAwMDAwMDAwMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDE1ZXgwZmRtY3FuMGFseHNyZjJlamE3cG5kdzRqeTlsMnh4ajJjOQ==",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": [
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "Y3JvMWZsNDh2c25tc2R6Y3Y4NXE1ZDJxNHo1YWpkaGE4eXUzZHFwazl4",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "Y3JvMXR5Z21zM3hoaHMzeXY0ODdwaHgzZHc0YTk1am43dDdsZXFhOG5q",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NTI0NjkxMzU2YmFzZWNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y3JvMXR5Z21zM3hoaHMzeXY0ODdwaHgzZHc0YTk1am43dDdsZXFhOG5q",
            "index": true
          }
        ]
      }
    ],
    "validator_updates": [
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "jZmiyA+S/yVqVuN2Px/9OqB/xgMPaj4mPdHpUOg/Kj0="
            }
          }
        },
        "power": "312"
      },
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "LNa+qkaUeJ97z/uLAKv1YTLMspaGxSkQyipkAmtwivo="
            }
          }
        },
        "power": "212"
      }
    ],
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
