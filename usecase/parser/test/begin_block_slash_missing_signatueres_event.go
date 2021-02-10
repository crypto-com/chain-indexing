package usecase_parser_test

var BEGIN_BLOCK_SLASH_MISSING_SIGNATURES_EVENT_BLOCK_RESULTS_RESP = `
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "168481",
    "txs_results": [],
    "begin_block_events": [
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNTAyNTl3MDhkOXZlM2Fra2Q2dnMwN2w5cTZyMjc1dWEyNjBlZDg=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OTIz",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxOHZ5cmozc2U2Y3ZkcnlrM3ZwMDl4ODhuNDY4OTR1a2p5d25mbXk=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "MTAwMQ==",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "slash",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxOHZ5cmozc2U2Y3ZkcnlrM3ZwMDl4ODhuNDY4OTR1a2p5d25mbXk=",
            "index": true
          },
          {
            "key": "cG93ZXI=",
            "value": "MTcyNzQ2MTc=",
            "index": true
          },
          {
            "key": "cmVhc29u",
            "value": "bWlzc2luZ19zaWduYXR1cmU=",
            "index": true
          },
          {
            "key": "amFpbGVk",
            "value": "Y3JvY25jbGNvbnMxOHZ5cmozc2U2Y3ZkcnlrM3ZwMDl4ODhuNDY4OTR1a2p5d25mbXk=",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxc3BqaHVqNnQ4N2d5MmpqOGR0cGZ2bWo1c2tyMjBkbmNhazN6bW4=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NDkz",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxd3R4ano2OTMzdnFmcDI2d2Z3ZGF2NnV0eDNucjJ3dHd0bjJndzM=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NDY3",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxaHA3N3N5dTgzc3c1Mmh6bG42dWxkYXRjdjJhZHF6dHNhNWdxdnI=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "ODg5",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxdTI5OXpzOXBrOXpjeXp2Y2pwbWhrdGF3ZnJ4c3djc3U3dTBreHM=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NDkw",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNHpzMjlkNTh6cTg1cTRmYXl4ODl5Z2FxNzkzbnprdTc3emVoZXk=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "ODM0",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxM3o5eThsY2h6cHcybXoycmM3MjM2bGEydDc1OGhxcjc5dGtqY2o=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NDc1",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbG00Z3RxdG5od2ZlbDJtMDhuYTV6dGt3ZWY0dnIzZHlnMnlqdHE=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NTM4",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxejlnZGV5MDNwZzBlZng5YXQzYzZrbWp1NzBrem15emwwc2g3YXc=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NjA3",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxOWxudWo1eGF2Y2xneXg1anZ4ajRoN2hrdHcwOXIwZzU5cDVtZ20=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NzE0",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNjhzZjlwd3VmamZsM3BjbmVkcmF2dHF3eTNhazczOHhlZWQwOGQ=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NTM3",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbmZ4Yzk4ZDJjZHZkZHR1ZG55cDR5a3YzbHE0OXN3OG02M3NlbDY=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NDg1",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxMndsOW0yNWR6bTYzd3U2NndjdjR6ODVlODJjY201Zm16NXF2bDQ=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NTY=",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxN3VlZXU4dXFubGwyaHQ2dGF4cjZjdXdmZ3p5bXY1NTc4ZDZ4azU=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NTQ2",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbHd4aHVrZjRwdTc4cmg5MGd3Z2g3NXVkY2gzMzlqYTh1M3UwcHY=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "ODkw",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxeGxyejg4aGZkOTBrcjJ6N3d2eXByeDhkODduMHl4Nzllc2FwNTk=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NTcy",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxaG43bW51OXhld3htZGM1ZHA1d2p2MGt1eWpqbHNyZWt2cW1ld24=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NTI4",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxZmw0enh0MDhycjZ3aGV0ZTh1cDVueWR5MzZtaGF1ZXJxeDJwemo=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "Mzgy",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNGQ1ODRseHFhZ3h3a2M1cnJwajRwbjIwZmQydHB6Zzd1azAzNHM=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NDU1",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxeWFnbnU5dmdxanh1aDU3cWZteTZuaDNzYXRsNTNoMHp4djNxaHM=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "ODIy",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxajR0ZXFlcjJlMnBheDludGhzN3RyYTJxMnhzeHNjY3lnNnhqcGc=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "Njk3",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxc2tlZzQyY2QwcDRlM2NrNTBmZW54aDIzd2ZtM2Nqajd5aDI4eDk=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "ODI4",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNzB0bGM3NHduc3o4cTlrZnBhZm5zbGZrZjh1Z2w1c3o4NXY2enk=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "ODIz",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNXI0ZXN4Njk5ZGFmM2VtaHdycWdwdnZocGxkdXJ0M3VmOHhhOHg=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NTIz",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxN215cXNqaGVzZHU1bXIzN3l3azkycHF6emxnMjBnZm5jNDIybm4=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OTc=",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxd2Nndndnam1tMDlmMjJuanlsajdhZGU1N3lsNXJzazN1OHNlNmM=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "ODQ2",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxZnZ2am52eXM2aHFyaDd5NGR3amNnOHl5dnd1aGFkOTQwMDhhdDU=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "MTg5",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxcms0ZmY0ODB6MGcyNDNkdHZlZXc2eWtqZjVnN3llMHU2cnUyanc=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OTg1",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxZ3EzdHo1ejA5dHR1Y3hodTc5dmt5OWc2bTR6eGdjdjUycTN4ODg=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NzM1",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbjRnZGoyMDdjc3hua3IzZmZmMGxsY3Z6azkzeTI5d3h4OWFsbTQ=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OTc2",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxZTRhOWF1MHBoeHNqZ3VjY2h6a3F6NnBweDZrNXVmOGFjam40aGo=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NzY2",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxODM0NXVmazN2ZjI1MmU4a21nanh6MzllcDkzdDR0MmZuM2xyeGo=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "ODc1",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxYWN0Njd4Y3A5ZXp2eDNwdWh1engwdjJxd2FsdnRydzJtaDJucTU=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NTky",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxZ2dnbDc1Y2g4MmV1dDk5eG12MHB1cmh0azU2dHp6ZW1sdjZkbnM=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OTIw",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbWc5OHg2bXpzZzNqbGxhNHUzNDRkNXltdjAzejVwZ214dDYzOG0=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "Njgz",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxNzB3cHhjdzczdWw4ZGgwczhqazQycTlmOGF4MnRxcjZnODVkNmo=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "ODAw",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxZjV4OG5yOXA3eHprbDY1OHVxNmU1cnVmbnJwZmt1d3F5MnBsZjQ=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "ODUw",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxMHd5NGszZGtqZDNodGdsZWF4bHptYWtoMGs0aDhxbDIzY3B1c3g=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "MTAwMQ==",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "slash",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxMHd5NGszZGtqZDNodGdsZWF4bHptYWtoMGs0aDhxbDIzY3B1c3g=",
            "index": true
          },
          {
            "key": "cG93ZXI=",
            "value": "OTkwMjAzMg==",
            "index": true
          },
          {
            "key": "cmVhc29u",
            "value": "bWlzc2luZ19zaWduYXR1cmU=",
            "index": true
          },
          {
            "key": "amFpbGVk",
            "value": "Y3JvY25jbGNvbnMxMHd5NGszZGtqZDNodGdsZWF4bHptYWtoMGs0aDhxbDIzY3B1c3g=",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxcTZhZHJlM2FzazI0NmdrcGplamtkdHRzdWE5NmV1NzU5emdtMmg=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OTUy",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbnR1eDNmMmdla2FydmttcjN4cWhyN2p6YWpnajBoNHRrZnRwOXk=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "NzM0",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "Y3JvY25jbGNvbnMxbXZoenN6ajN1dWU3dndnYWc1ZjRha3ZxZWx5cWEwY2R4ZWV6OHI=",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "OTE2",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTY4NDgx",
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
            "value": "MjcxNDk0NzM1NDMwMzViYXNldGNybw==",
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
              "ed25519": "0ONjQW1vbJ1wJezekfvKplMH0SDbbUoPzaCm6W6I53c="
            }
          }
        }
      },
      {
        "pub_key": {
          "Sum": {
            "type": "tendermint.crypto.PublicKey_Ed25519",
            "value": {
              "ed25519": "YRBTvcFy5dPrJwFBkVR9Rj5tEHoTn56EOtzvlb/jUJY="
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
