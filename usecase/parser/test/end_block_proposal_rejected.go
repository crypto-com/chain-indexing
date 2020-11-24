package usecase_parser_test

const END_BLOCK_PROPOSAL_REJECTED_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "8CC241A96C52C137928DBE8CA7E290CF94A46B5653C022ED2663810811F4CEC7",
      "parts": {
        "total": 1,
        "hash": "086DADFD6C3B05C6563063B4FCCBF24090F66AB2C1D5F4816626B5E6C1B65DF2"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "chainmaind",
        "height": "480",
        "time": "2020-11-21T13:40:00.028278Z",
        "last_block_id": {
          "hash": "47F87F7C8BAB6286940868BBC3241399C1D4C75A550F95DA2709465B3C521BEE",
          "parts": {
            "total": 1,
            "hash": "4AFB8C0A0D8AC7EF340FCBF440875E17F5FA7BB60BABE3EF4A63707404E328BB"
          }
        },
        "last_commit_hash": "AD424BA6306B43CFDE3DC955F8628040701AB1426B9EC62202D1DA1252E70A1C",
        "data_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "validators_hash": "C3713F6B3AA91FCEF36844A73764052CBF38A4D6081775E9F46789F84011A2DA",
        "next_validators_hash": "C3713F6B3AA91FCEF36844A73764052CBF38A4D6081775E9F46789F84011A2DA",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "0256598B84DD30DAFE873390E4DE6A2226E2791FBFB5A575F69C38CC0A6A09DA",
        "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "EE2934A38A8066EB27C831ADD91B17263D77EEA6"
      },
      "data": {
        "txs": []
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "479",
        "round": 0,
        "block_id": {
          "hash": "47F87F7C8BAB6286940868BBC3241399C1D4C75A550F95DA2709465B3C521BEE",
          "parts": {
            "total": 1,
            "hash": "4AFB8C0A0D8AC7EF340FCBF440875E17F5FA7BB60BABE3EF4A63707404E328BB"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "3A4721A86B536F867A728CF1948E264BF451C708",
            "timestamp": "2020-11-21T13:40:00.028278Z",
            "signature": "Qh+QlAZYOhGQzNZgWSipewU15F2fC0D5TW98CBTafL1foNuSrQ5ua7VpDntLU5IMDFvY15Yz2tixL4tynH2KBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "EE2934A38A8066EB27C831ADD91B17263D77EEA6",
            "timestamp": "2020-11-21T13:40:00.068027Z",
            "signature": "O/AQbyZ6tR5KlT7iBCFxUQMkW8dglu4sizeJZKP1zlAtekhi+mWn1JMgbK59CQLB+5L+hsRsgZn6Mnxm/nRxDg=="
          }
        ]
      }
    }
  }
}`
