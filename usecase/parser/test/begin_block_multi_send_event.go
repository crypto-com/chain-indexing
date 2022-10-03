package usecase_parser_test

var BEGIN_BLOCK_MULTI_SEND_EVENT_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "165487",
    "txs_results": [],
    "begin_block_events": [
      
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjAwMDBiYXNldGNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "YWN0aW9u",
            "value": "bXVsdGlzZW5k",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFmbXBybTBzank2bHo5bGx2N3JsdG4wdjJhenp3Y3d6dmsybHN5bg==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzE0bTVhNGt4dDJlODJ1cXFzNWd0cXphMjlkbTV3cXp5YTJqdzlzaA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MWJhc2V0Y3Jv",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzE0bTVhNGt4dDJlODJ1cXFzNWd0cXphMjlkbTV3cXp5YTJqdzlzaA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjBiYXNldGNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzE0bTVhNGt4dDJlODJ1cXFzNWd0cXphMjlkbTV3cXp5YTJqdzlzaA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MzBiYXNldGNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "bW9kdWxl",
            "value": "YmFuaw==",
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
              "ed25519": "irHM4J0M4AktkI3vTkNemLs/LeO3EVi/Ps/y0mn+/c4="
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
}
`
