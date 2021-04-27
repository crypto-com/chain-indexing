package usecase_parser_test

// Ignore gosec rule because this is CosmosHub on-chain data
// nolint:gosec
const END_BLOCK_PROPOSAL_PASSED_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "520186",
    "txs_results": null,
    "begin_block_events": [
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "Y29zbW9zMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsc2VycXRh",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "Y29zbW9zMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04ZzM4Yzhx",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjI1NDQ5MzQ3MjE0MXVtdW9u",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y29zbW9zMW0zaDMwd2x2c2Y4bGxydXh0cHVrZHZzeTBrbTJrdW04ZzM4Yzhx",
            "index": true
          }
        ]
      },
      {
        "type": "mint",
        "attributes": [
          {
            "key": "Ym9uZGVkX3JhdGlv",
            "value": "MC4wMDAxMDA0MjIyNjY3MTg2MDQ=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4xNDA3MTQyODA2MzcxNzUyNDI=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTQyMjkyODA2MzkyODkwOTM4NzMuMjE2NjM2MDIxOTExNjEyMDY0",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjI1NDQ5MzQ3MjE0MQ==",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "Y29zbW9zMWp2NjVzM2dycWY2djZqbDNkcDR0NmM5dDlyazk5Y2Q4OGx5dWZs",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "Y29zbW9zMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsc2VycXRh",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MjI1NDQ5MzQ3MjE0MXVtdW9u",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "Y29zbW9zMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsc2VycXRh",
            "index": true
          }
        ]
      },
      {
        "type": "proposer_reward",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTEyNzI0NjczNjA3LjA1MDAwMDAwMDAwMDAwMDAwMHVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF4NXdnaDZ2d3llNjB3djNkdHNoczlkbXFnZ3dmeDJsZGs1Y3ZxdQ==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTEyNzI0NjczNjA3LjA1MDAwMDAwMDAwMDAwMDAwMHVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF4NXdnaDZ2d3llNjB3djNkdHNoczlkbXFnZ3dmeDJsZGs1Y3ZxdQ==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTEyNzI0NjczNjA3LjA1MDAwMDAwMDAwMDAwMDAwMHVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF4NXdnaDZ2d3llNjB3djNkdHNoczlkbXFnZ3dmeDJsZGs1Y3ZxdQ==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjIxNjQwMTYzODc2LjU4MTk1MDI2NTUzMzgzOTQyMnVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF4NXdnaDZ2d3llNjB3djNkdHNoczlkbXFnZ3dmeDJsZGs1Y3ZxdQ==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjIxNjQwMTYzODc2LjU4MTk1MDI2NTUzMzgzOTQyMnVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF4NXdnaDZ2d3llNjB3djNkdHNoczlkbXFnZ3dmeDJsZGs1Y3ZxdQ==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDEzNTkyNTE5NTU5LjYwNjA4MTExNTA0MDg0Nzc4MnVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF3OTZycmg5c3gwaDduN3FhazAwbDkwdW4wa3g1d2FsYTJwcm14dA==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDEzNTkyNTE5NTU5LjYwNjA4MTExNTA0MDg0Nzc4MnVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF3OTZycmg5c3gwaDduN3FhazAwbDkwdW4wa3g1d2FsYTJwcm14dA==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDEzNTkyNTExNTA3LjI3MjY4NTE4MTgyOTE3MTUwOXVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjFmanE5YzdkNDV6dTB5cDN6ZGt1MDV6azR2eHQ0ejBoM3Q4dnp1bA==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDEzNTkyNTExNTA3LjI3MjY4NTE4MTgyOTE3MTUwOXVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjFmanE5YzdkNDV6dTB5cDN6ZGt1MDV6azR2eHQ0ejBoM3Q4dnp1bA==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDEzNTg4NTk1NTkuNTYwMDUzNDY3MjY0MzkwMDMzdW11b24=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF1czd2eDZ0dXI5dDAzNjhqMGt6eXY2aGNxNDdkdmd2ODV6dnl0eg==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDEzNTg4NTk1NTk1LjYwMDUzNDY3MjY0MzkwMDMyN3VtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF1czd2eDZ0dXI5dDAzNjhqMGt6eXY2aGNxNDdkdmd2ODV6dnl0eg==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjA3Mzc2MTc2NTQuNjEwODk1MTQzMTQ3MTg3NzExdW11b24=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF0ZmxrMzBtcTV2Z3FqZGx5OTJra2hocTNyYWV2MmhuejZlZXRlMw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjA3Mzc2MTc2NTQ2LjEwODk1MTQzMTQ3MTg3NzExM3VtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF0ZmxrMzBtcTV2Z3FqZGx5OTJra2hocTNyYWV2MmhuejZlZXRlMw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjA2NzkzNzI4OS43MjcxMDI2NTQ1ODgwNDQyNTF1bXVvbg==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjEyazVuenByYzh6cjVyN204Z2E2bHR3ZTZ1Zm1xcDVnMDI4bG15bA==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjA2NzkzNzI4OTcuMjcxMDI2NTQ1ODgwNDQyNTA3dW11b24=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjEyazVuenByYzh6cjVyN204Z2E2bHR3ZTZ1Zm1xcDVnMDI4bG15bA==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mjk3ODY3MTI2LjkyODIwOTE1NTEwMTA0OTA1N3VtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF6a3VwcjgzaHJ6a24zdXA1ZWxrdHpjcTN0dWZ0OG54c213ZHFncA==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mjk3ODY3MTI2OS4yODIwOTE1NTEwMTA0OTA1Njl1bXVvbg==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF6a3VwcjgzaHJ6a24zdXA1ZWxrdHpjcTN0dWZ0OG54c213ZHFncA==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODE4OTk4OTAuNjkxMTMyNDExNTQyNjExNTI3dW11b24=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjFqdGFhcnQ4NjI0bmNtcTcwczVlODM0NnhudWprd2txNmdtODVyeg==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODE4OTk4OTA2LjkxMTMyNDExNTQyNjExNTI3NHVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjFqdGFhcnQ4NjI0bmNtcTcwczVlODM0NnhudWprd2txNmdtODVyeg==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjQ3NzU5MDEuNzIzNDAyNTgxMTkzOTcwMjI3dW11b24=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF4anBhZDJzdHlncXlhZ2M4ZnFwcDB1Y3N1NmxyZXRmaHE2N2s0NQ==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjQ3NzU5MDE3LjIzNDAyNTgxMTkzOTcwMjI3MnVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF4anBhZDJzdHlncXlhZ2M4ZnFwcDB1Y3N1NmxyZXRmaHE2N2s0NQ==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjEwMTg5OTkwLjUzNTUxNjQ0MjcxNjg3ODU5OHVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjE5eGN6eHZ2ZGc4aDY3c2szY2NjcnZ4bGowcnV5dzMzNjByY3RmYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjM2OTM5MzY1LjI1OTE0MDczNTUwNTY5MjcyMXVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjE5eGN6eHZ2ZGc4aDY3c2szY2NjcnZ4bGowcnV5dzMzNjByY3RmYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTIzODkzNjk3LjE3OTYxMTA3MjU1NTg1ODgyN3VtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjEwNDZzMmxqZGR6eGF2MGM0ampmYWNuNGd6bTk0anFycWQ1MHZyZw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjE5NDY4NDg1Ljg5ODA1NTM2Mjc3OTI5NDEzN3VtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjEwNDZzMmxqZGR6eGF2MGM0ampmYWNuNGd6bTk0anFycWQ1MHZyZw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjM5NDEzODM0LjgwMDM1ODY3NzQ1MDA1ODY4MnVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjFjdWNrNTJobDJ1c3ZjY3c2bmxncDQ0ZXEweDhwN3luOGY2Y21lNw==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDc4ODI3NjY5LjYwMDcxNzM1NDkwMDExNzM2M3VtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjFjdWNrNTJobDJ1c3ZjY3c2bmxncDQ0ZXEweDhwN3luOGY2Y21lNw==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTYyNjYzMy4yODcyODc2OTkxMDU1MDA2Nzh1bXVvbg==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjFsbXhuM3R1bWhwNmNwbXY0ZXBxNGhhOXI5aHYwanJ2OW44aGhreQ==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTYyNjYzMzIuODcyODc2OTkxMDU1MDA2NzgwdW11b24=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjFsbXhuM3R1bWhwNmNwbXY0ZXBxNGhhOXI5aHYwanJ2OW44aGhreQ==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTk4OTU5NS4zMTIyMjM3NTMxNzUxNzQ0MTZ1bXVvbg==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF3cDBtcGNjdWduaHBwaHp2YTA4dmNtcHhkOTJtMGQ1aGNjMnI2Ng==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODI4OTk4MC40Njc1OTg5NzE1NjMyMjY3MzV1bXVvbg==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF3cDBtcGNjdWduaHBwaHp2YTA4dmNtcHhkOTJtMGQ1aGNjMnI2Ng==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDE2MDU3OC43MjUzNzUyNzcwMDg1MjE4NjR1bXVvbg==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF3amRnZWVyc252d2Y5ZTd3N2o1NHY3bHU5eWZsdndlNnp5MDRyeQ==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDE2MDU3OC43MjUzNzUyNzcwMDg1MjE4NjR1bXVvbg==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjF3amRnZWVyc252d2Y5ZTd3N2o1NHY3bHU5eWZsdndlNnp5MDRyeQ==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjA2NDcuMDA4NzEwOTQwMjU1MDY0NTUydW11b24=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjE1ZGV1N2hzcjg1Y3U4eDNzMGZtZHZtYXZrOW1yNnFqN2Vsd2hwNg==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjA2NDcwLjA4NzEwOTQwMjU1MDY0NTUxNnVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjE1ZGV1N2hzcjg1Y3U4eDNzMGZtZHZtYXZrOW1yNnFqN2Vsd2hwNg==",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAzLjIzNTA0MzQyNDcwNzE4MTcxOXVtdW9u",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjEwNWd2Y2pnczZzNGo1d3M5c3Jja3gwZHJ0NHg4Y3dneXdwbGg3cA==",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAzMi4zNTA0MzQyNDcwNzE4MTcxOTF1bXVvbg==",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "Y29zbW9zdmFsb3BlcjEwNWd2Y2pnczZzNGo1d3M5c3Jja3gwZHJ0NHg4Y3dneXdwbGg3cA==",
            "index": true
          }
        ]
      }
    ],
    "end_block_events": [
      {
        "type": "active_proposal",
        "attributes": [
          {
            "key": "cHJvcG9zYWxfaWQ=",
            "value": "Nw==",
            "index": true
          },
          {
            "key": "cHJvcG9zYWxfcmVzdWx0",
            "value": "cHJvcG9zYWxfcGFzc2Vk",
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
