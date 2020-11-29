package infrastructure_tendermint_test

const BLOCK_RESULTS_EMPTY_EVENTS_JSON = `{
	"jsonrpc": "2.0",
	"id": -1,
	"result": {
		"height": "1",
		"txs_results": null,
		"begin_block_events": null,
		"end_block_events": null,
		"validator_updates": null,
		"consensus_param_updates": {
		  "block": {
			"max_bytes": "22020096",
			"max_gas": "-1"
		  },
		  "evidence": {
			"max_age_num_blocks": "100000",
			"max_age_duration": "172800000000000"
		  },
		  "validator": {
			"pub_key_types": [
			  "ed25519"
			]
		  }
		}
	}
}`

const BLOCK_RESULTS_JSON = `
	{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "367216",
    "txs_results": [
      {
        "code": 0,
        "data": "CgoKCGRlbGVnYXRl",
        "log": "[{\"events\":[{\"type\":\"delegate\",\"attributes\":[{\"key\":\"validator\",\"value\":\"tcrocncl1pm27djcs5djxjsxw3unrkv3m3jtxdexktw5epu\"},{\"key\":\"amount\",\"value\":\"19302674761\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"delegate\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro1pm27djcs5djxjsxw3unrkv3m3jtxdexk73hqel\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"amount\",\"value\":\"1913979901basetcro\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "143179",
        "events": [
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
                "value": "dGNybzFwbTI3ZGpjczVkanhqc3h3M3Vucmt2M20zanR4ZGV4azczaHFlbA==",
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
                "value": "dGNybzFwbTI3ZGpjczVkanhqc3h3M3Vucmt2M20zanR4ZGV4azczaHFlbA==",
                "index": true
              }
            ]
          }
        ],
        "codespace": ""
      },
      {
        "code": 0,
        "data": "CgoKCGRlbGVnYXRl",
        "log": "[{\"events\":[{\"type\":\"delegate\",\"attributes\":[{\"key\":\"validator\",\"value\":\"tcrocncl10gsqs8jzdlrem80shp0x6wx0jw7qu7m8cd29y5\"},{\"key\":\"amount\",\"value\":\"17338013566\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"delegate\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tcro10gsqs8jzdlrem80shp0x6wx0jw7qu7m8djfuuh\"},{\"key\":\"sender\",\"value\":\"tcro1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8339p4l\"},{\"key\":\"amount\",\"value\":\"2310941249basetcro\"}]}]}]",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "143737",
        "events": [
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
                "value": "dGNybzEwZ3NxczhqemRscmVtODBzaHAweDZ3eDBqdzdxdTdtOGRqZnV1aA==",
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
                "value": "dGNybzEwZ3NxczhqemRscmVtODBzaHAweDZ3eDBqdzdxdTdtOGRqZnV1aA==",
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
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc0NDk1MjgzMjFiYXNldGNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzFtM2gzMHdsdnNmOGxscnV4dHB1a2R2c3kwa20ya3VtODdseDltcQ==",
            "index": true
          }
        ]
      },
      {
        "type": "mint",
        "attributes": [
          {
            "key": "Ym9uZGVkX3JhdGlv",
            "value": "MC4wMDA4MDkxOTYwNTQzNzY2NDQ=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTM3NTU4MjE5MzY4NTUxODQ=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTEwMTMzMDQ2OTk0MjA0NTc2LjEzODAxNjUyNjU3OTM4NjI4OA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTc0NDk1MjgzMjE=",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": [
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODcyNDc4NDEuNjA1MDAwMDAwMDAwMDAwMDAwYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOHAwN3l2bXBoeW1zY3o2dGw0YTd6bWg5M2cwazZ2eTcyd3c0czQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODcyNDc4NDE2LjA1MDAwMDAwMDAwMDAwMDAwMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOHAwN3l2bXBoeW1zY3o2dGw0YTd6bWg5M2cwazZ2eTcyd3c0czQ=",
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
              "ed25519": "SE5zeTjcYPXVrfcOva61QWokSZFfQu2h316fR6bB2dY="
            }
          }
        },
        "power": "138525202"
      },
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "Epmo3U6yXlxSDQzWZ8yBPOMHw2R85lc26RK98Rlo0oM="
            }
          }
        },
        "power": "112904113"
      }
    ],
    "consensus_param_updates": {
      "block": {
        "max_bytes": "22020096",
        "max_gas": "-1"
      },
      "evidence": {
        "max_age_num_blocks": "100000",
        "max_age_duration": "172800000000000"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      }
    }
  }
}`

const BLOCK_RESULTS_VALIDATOR_UPDATES_WITHOUT_POWER_JSON = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "4794",
    "txs_results": null,
    "begin_block_events": null,
    "end_block_events": null,
    "validator_updates": [
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "CpCz+c19SHaNWW31P+7blzyHo0sQMn4uk8gIej+pXW8="
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
        "max_age_duration": "172800000000000"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      }
    }
  }
}`
