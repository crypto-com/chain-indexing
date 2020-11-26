package usecase_parser_test

var VALIDATOR_UPDATES_VALIDATOR_SLASHED_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "28",
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
            "value": "MC4wNDA3MjA4MDU2NzcyMTUxNzg=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xMzAwMDA1NDYwOTQyODAyNTY=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "ODA2MDAzODM0MC4xMzQyMDQ2NTY2ODg5NDY2MjQ=",
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
            "value": "NTMuMjQ1NDM1ODE2MTY0ODE2OTMyYmFzZWNybw==",
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
            "value": "NS4zMjQ1NDM1ODE2MTY0ODE2OTNiYXNlY3Jv",
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
            "value": "NTMuMjQ1NDM1ODE2MTY0ODE2OTMyYmFzZWNybw==",
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
            "value": "NDcuNDcyODQzMjcxOTQyNzU2ODU2YmFzZWNybw==",
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
            "value": "NDc0LjcyODQzMjcxOTQyNzU2ODU2NWJhc2Vjcm8=",
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
            "value": "NDcuNDcyODQzMjcxOTQyNzU2ODU2YmFzZWNybw==",
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
            "value": "NDc0LjcyODQzMjcxOTQyNzU2ODU2NWJhc2Vjcm8=",
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
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTQuODExNTI3MTAwODQ2MTQwMDMyYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFyZ2VmbXF3Mnd1ZndmZnNnOTUyOGs1OGVkZ3NnYzR4aHpzeDlsMg==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTQ4LjExNTI3MTAwODQ2MTQwMDMxOGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFyZ2VmbXF3Mnd1ZndmZnNnOTUyOGs1OGVkZ3NnYzR4aHpzeDlsMg==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAuMDY0MjQyNzczNjUxODY0NDQyYmFzZWNybw==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFyM3Zhc2UydWc2Z3d1NDdzYzVtcXZkdjA3Z2t1amZtaDVtZ3RhZg==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAwLjY0MjQyNzczNjUxODY0NDQyMGJhc2Vjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y3JvY25jbDFyM3Zhc2UydWc2Z3d1NDdzYzVtcXZkdjA3Z2t1amZtaDVtZ3RhZg==",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNTQ4ZjVoeWRkZGcwZWE0c2RneHNlN3Q3ajRqbjg0enA2eTk1NGU=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OQ==",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "Mjg=",
            "index": true
          }
        ]
      },
      {
        "type": "slash",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNTQ4ZjVoeWRkZGcwZWE0c2RneHNlN3Q3ajRqbjg0enA2eTk1NGU=",
            "index": true
          },
          {
            "key": "cG93ZXI=",
            "value": "MzEy",
            "index": true
          },
          {
            "key": "cmVhc29u",
            "value": "bWlzc2luZ19zaWduYXR1cmU=",
            "index": true
          },
          {
            "key": "amFpbGVk",
            "value": "Y3JvY25jbGNvbnMxNTQ4ZjVoeWRkZGcwZWE0c2RneHNlN3Q3ajRqbjg0enA2eTk1NGU=",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbmZ0ZzJuOWd6anIybDdsZW1jc2hrMHY4d2RtdXV6cThudWxudnM=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OQ==",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "Mjg=",
            "index": true
          }
        ]
      },
      {
        "type": "slash",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbmZ0ZzJuOWd6anIybDdsZW1jc2hrMHY4d2RtdXV6cThudWxudnM=",
            "index": true
          },
          {
            "key": "cG93ZXI=",
            "value": "MjEy",
            "index": true
          },
          {
            "key": "cmVhc29u",
            "value": "bWlzc2luZ19zaWduYXR1cmU=",
            "index": true
          },
          {
            "key": "amFpbGVk",
            "value": "Y3JvY25jbGNvbnMxbmZ0ZzJuOWd6anIybDdsZW1jc2hrMHY4d2RtdXV6cThudWxudnM=",
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
            "value": "Y3JvMXR5Z21zM3hoaHMzeXY0ODdwaHgzZHc0YTk1am43dDdsZXFhOG5q",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "Y3JvMWZsNDh2c25tc2R6Y3Y4NXE1ZDJxNHo1YWpkaGE4eXUzZHFwazl4",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "NTE5NDUxMzU2YmFzZWNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y3JvMWZsNDh2c25tc2R6Y3Y4NXE1ZDJxNHo1YWpkaGE4eXUzZHFwazl4",
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
        }
      },
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "LNa+qkaUeJ97z/uLAKv1YTLMspaGxSkQyipkAmtwivo="
            }
          }
        }
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
