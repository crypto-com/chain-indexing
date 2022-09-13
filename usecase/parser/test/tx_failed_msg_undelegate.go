package usecase_parser_test

const TX_FAILED_MSG_UNDELEGATE_BLOCK_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "block_id": {
      "hash": "9578A8F24E6F7820DAAB9DDEF241F0E1B5178028B82850F5C7A10A678BF84173",
      "parts": {
        "total": 1,
        "hash": "C58CAC6C7539940D8EA51976979732995402B2EA079A9BA581D250939883F511"
      }
    },
    "block": {
      "header": {
        "version": {
          "block": "11"
        },
        "chain_id": "testnet-croeseid-1",
        "height": "184399",
        "time": "2020-10-28T12:21:43.05321162Z",
        "last_block_id": {
          "hash": "13A6E4FC56F827443F4F6D2009624DA9C48078BA8919F4C15921E4AC8DF6AB7D",
          "parts": {
            "total": 1,
            "hash": "48206C6A0537A6147B69A0D7F901862A96C568BE59E0AF4B73EB1161E0B2BE2B"
          }
        },
        "last_commit_hash": "9E5A1B4B1BFCE6ADDBE7D4C0C486A860B29C4FEC0CBD370A5AD28FBE5AABD67C",
        "data_hash": "4AA4B9F725C2F5662D4E324BE3249659A6D2015A770DD0CC933B921353658F22",
        "validators_hash": "72EA2A7DCC4962BE35EEE017D5D9251631C0AAC05D83AE7E7CD1F2588C8259DB",
        "next_validators_hash": "72EA2A7DCC4962BE35EEE017D5D9251631C0AAC05D83AE7E7CD1F2588C8259DB",
        "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
        "app_hash": "A124B6020267C98A2BB40818A69D9697B3EAB2CA8963B49C13AA1E76C29BFA80",
        "last_results_hash": "3321C2E9EE05506EF7E9982EDDDC3DFE358D5BD9BA8644DDE92058837F4544E7",
        "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
        "proposer_address": "0F7479FBE0193349D3FD22D4F5ADD2BFFFAB1DA1"
      },
      "data": {
        "txs": [
          "CqYBCqMBCiUvY29zbW9zLnN0YWtpbmcudjFiZXRhMS5Nc2dVbmRlbGVnYXRlEnoKK3Rjcm8xbGxzdDBjZ3VoNWF6bDl0OHdyNm16NXl6anV3dWt6N2Y2N3o3ZjYSL3Rjcm9jbmNsMTVlNjlrZHJ0Y3phampkbHp5dDJxZ3M1cTJhbmM1cXBtazJjNjh6GhoKCGJhc2V0Y3JvEg41MDAwMDAwMDAwMDAwMBJqClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDxgP6Hyqdhy1LijIg6NnLf3X2m9cRn55aA/DWuNmCSusSBAoCCAEYAhIWChAKCGJhc2V0Y3JvEgQ1MDAwEMCaDBpA0aXlcyCxEWPiHMA0aTBa2bBDZ76M1oHeIPoIO5wzoDItcENuhe/WSJmWa1NLcLOIggz9I12XmGCnXND8G/Kolw=="
        ]
      },
      "evidence": {
        "evidence": []
      },
      "last_commit": {
        "height": "184398",
        "round": 0,
        "block_id": {
          "hash": "13A6E4FC56F827443F4F6D2009624DA9C48078BA8919F4C15921E4AC8DF6AB7D",
          "parts": {
            "total": 1,
            "hash": "48206C6A0537A6147B69A0D7F901862A96C568BE59E0AF4B73EB1161E0B2BE2B"
          }
        },
        "signatures": [
          {
            "block_id_flag": 2,
            "validator_address": "504C0C3FE72728946911C7956E1B012784446B64",
            "timestamp": "2020-10-28T12:21:42.935586448Z",
            "signature": "O4jyTJW9lurn5kCPSF+P4isU6qorkpK78MMGzjH0brMSpsrcCbHgqnsHc5LuuqS4H9flOZncdRYRyJUgxnd4Cw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A1E8AAEBBC82929B852748734BA39D67A62F201B",
            "timestamp": "2020-10-28T12:21:42.983810251Z",
            "signature": "mGby0rgMI4a5w3tRurigYIyo+MqRMn3pNdf2ZLnVApKSpeD04fjCa4uMqeTrQ56ZKgWAxiAfHlgywt1bwU6vAg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3705DA4F2E53A09025DAA8E6581EFCE851811914",
            "timestamp": "2020-10-28T12:21:43.051500249Z",
            "signature": "YFLdv+BDuu4aQzSTqjKLviH+XwqUuCX9DAJuE7rVSO7/DCTwrhZAcXU8GqCH5ROzZCmrfsMutqqvKy3UNdMzBw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4B68F098199E7F565B02EF58115FB3CB9BAD52B0",
            "timestamp": "2020-10-28T12:21:43.084724153Z",
            "signature": "vDQtP/yMwbiBg6UXhWJ8su2BxC2+Ku/TqWku9VhTxWQcFbQsPecgrZcxNNgZPWuaAWTZtBt0Z7xn18kLdaWzDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "95CDD1C2F0E79F62745D17A90D9A7B138DC8F922",
            "timestamp": "2020-10-28T12:21:43.165886176Z",
            "signature": "HzXioUTqUm17tysOFfWGTs0Pf/ZrC56XfDGj6/B1IEsbS0S1VuYB2RmKN/4s9yaaXIxbccSy5VMh3BL1BgfkDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "A3DC2532609F7C2E6BBAB2809080F9EA1EA183C0",
            "timestamp": "2020-10-28T12:21:43.106854548Z",
            "signature": "dmJlNaGRJS/kdScv4jnwMsNmoCqsV/IKac/FJ+map12M0zZpECcHnK7PjAFtM8wbEAwEcR1Yw5aYshWSIWtXBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "EDB80454CA0459949E6E564A979E2A5C20BE2DAC",
            "timestamp": "2020-10-28T12:21:43.084001989Z",
            "signature": "VUvmKDLw2IgeW7YzFW/gkaOv5aYn8WMtK3sumKtZ7rIfUjy7LfgkeXZLVeSq6AxzqW/O7RyYLvMFt8a+3MCkDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "6E66EC70E6A243F69BC82AEF3783E63E237A6FE1",
            "timestamp": "2020-10-28T12:21:43.05321162Z",
            "signature": "TkBg5CIDAtNzMjUYOT1hW5KXa/3g93meTcs59UR0DRwBPT3Dwd9SZN7RZUVGTt50npoN0rgYbSxOSVDjHOHLAg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "5D7B2D44CDAC899167C44AD3E3286D6DA80A31DE",
            "timestamp": "2020-10-28T12:21:43.096808846Z",
            "signature": "mS6GHEymcPCJyFgzHBwIWu2i+F2cZzb+KneabywCU80HCNn4AxGBcO6xdXCIlKQ85IMuoD3BseOSOA6CJrCCBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "0F7479FBE0193349D3FD22D4F5ADD2BFFFAB1DA1",
            "timestamp": "2020-10-28T12:21:43.075824844Z",
            "signature": "n3OxWisoHzW2T8glJWlQ7WErUugur/ZJVCsusF13OJHCC8DQJnEExPO/8x8c0J/SWO7mzP/xu1yGlA/gUqgFAw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "725BFF0129AE55F287366DE568940FD76280F180",
            "timestamp": "2020-10-28T12:21:43.038632719Z",
            "signature": "M2ky3XyRNiWJ0bbc9x21aqtuuUP7Aw75lYrWmRcmVQZtBdGhY8CHvEbfleH8pFYXjh6NUp4p1i2WWejaE9RgDA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E23F5C63108A87067478C426B7CE3B8DABFC1ED4",
            "timestamp": "2020-10-28T12:21:43.057960227Z",
            "signature": "G/QbMZbMcByggTYbX0cde1f8szuz5S3SxMZ3z27IVmYQ3dbNeRhzP4BjpA+vmB8bHbCe89Xi0dhIfA8gGbToDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "7CF7D85DEF3B52A521B0EAC4AF3F154E5A1D888A",
            "timestamp": "2020-10-28T12:21:43.094082242Z",
            "signature": "xn0i6Wclaqu/CKz9UKq4MnjPPEtgVYz8ye3uWPmL7qqxRDFsubLkUoklZAH3wC8BajEhtVNPCBfEjbLTfw5sAA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "D3A418F7C83C7DE225D6F5C9E8E97CC054FD1469",
            "timestamp": "2020-10-28T12:21:43.596455211Z",
            "signature": "JEnKg8OJLHENpmthLuzBISbuRYg01dEih6CXNtGY/Ur2g/21wuKmq7PKMtM0oBOIq4/vIgLYAPb1BWGtBJjQBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "80FB6FE1A78531BABD26642F1D30B498D9D9A894",
            "timestamp": "2020-10-28T12:21:43.080875954Z",
            "signature": "nA/x4v3fUQ35nWJ8Z3xN3Ds3CWFl3TG8ib6oa90aDhKF9WgqpA6HIPzXYaKcfwwcTmrI3I2Bzx7igtjtnADzBw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "1C289E675E313B3233365F212F860053E8E54242",
            "timestamp": "2020-10-28T12:21:43.092705337Z",
            "signature": "BQnLiaZnNIvsfD/TAd2pOIfqZBBy9cF2R4aVKCgb9rARWxWl+g+t2jK3nJGmAsO9S5Hxc5sDrNO8nqnl+MlZBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "D899A023FD783DA554F12D9CA37C9B3CCAFD6087",
            "timestamp": "2020-10-28T12:21:43.11115708Z",
            "signature": "FJMEVyuEi6iFP69l7gxK8qQWkNTbU4aGnoopm4+73ZquS4IWX5Igz3nbgHYDLFiSl70BJ45Z3WLAHqRg43yNDw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "064D3F6A5737D3605A3AEFFE30BA8BEE2ACBDF7E",
            "timestamp": "2020-10-28T12:21:43.021681498Z",
            "signature": "Dr/e5hyyZS4WF9/WmgDX2I9uH/B4nNNR4NGZMWETtmyuFqXJ1A+/I0i9AYu/66K3Ld6UY/3PuqraBlJFYbm4CQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "D2B48A6A01CD6B3BBBA3DD4D0CAC047CA38C80C0",
            "timestamp": "2020-10-28T12:21:43.068996887Z",
            "signature": "pKKOj2DwXrynqKzZxQmNpJzZQoUsQUHSDFyKMj9bXTc3jORk/IsICAY4iWtaMUxQhsYIdgYuLmdVsXZweS8bDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "D62A1C65C25591E0E1C7FEAC1F860C8CFDDDE6B2",
            "timestamp": "2020-10-28T12:21:43.107034348Z",
            "signature": "1m8kpI8SFmomWio5Qx4osoa/ih6C9nB8855mYhrZ2YZdf6SgkBVCVhDS5eRpbr0dOqjys29QEKurW6XTTUHABQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "F7EE6CD99D77C737E66A033F7E1BE460A03B2533",
            "timestamp": "2020-10-28T12:21:42.986135597Z",
            "signature": "whdLEWt2AJs3PkRm45T7Z2BKdAo3GdZoK/G1JROri0IPlPJfwb1bBK+YE94SuyfCuzju689XCW7leoOLJ+xsBg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "EEB931B95F19BC8190146F87C2B771845C2C84B6",
            "timestamp": "2020-10-28T12:21:43.079102009Z",
            "signature": "5hnQRlOBzUILzHhPCuphF1rILsY2lpRfgZO8mQMYcu6YnYiIiehHnDJPireDPRJm3OlfboV/ESgEXuJSbylJBw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "62E0A1625735852A49D0D42BE3D11EE6EE4EB198",
            "timestamp": "2020-10-28T12:21:43.141304015Z",
            "signature": "d85WTFqrufLjAzhgzzGAoViyA2Hk1TfzL0tyGqIrWbZ63p/7ZV0df8/7R4k4+KQSV42VXWJOzDetAEpBnNOdBA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "879944D7386D753A883DEED60C29F5487AA0BDFE",
            "timestamp": "2020-10-28T12:21:42.977577713Z",
            "signature": "FF4vuU13OtCRcO94UGM7D7OywfAvUCY/40tqxau7qaPYydtmjyARNMNxqpuhwqQkDuszrDoEX9cxVUZMA3Q0DQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "0643A20176ADE79C775D40A87E27EF5A401E982E",
            "timestamp": "2020-10-28T12:21:42.972468618Z",
            "signature": "Y1QR26/ySojFjS3uhYuqSAI2tcgeN3J632wtLC/quwz2E62TTFPSfWQvk0u/O9QxYieHJRN0zfNa8yoXqw+/Bg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "CA721C3A05F500838DDD1B16F4E2D2D09E463218",
            "timestamp": "2020-10-28T12:21:43.076982631Z",
            "signature": "jV6RvzyPxf94pEYdOPrYd/32CCa9suWuzC3EBqitzDgNgghEuCxK56+t2guqg+eWKpim7wg8oySa5ZQf3nkwAg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "5CECCE87B39619DFE859E9A01D78391155BC19C3",
            "timestamp": "2020-10-28T12:21:43.060924365Z",
            "signature": "dAkGlJ+hxnUdBJKwe+DvnUD/Tsbkhz9oor9yFKzo77AFYBVRRokOK0P0dPA9+G48VMu0UHNSkS5kkNipr1I7DQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "B4955A72BEE29AF36BC48496E7E7AFD63FE8FC4C",
            "timestamp": "2020-10-28T12:21:43.081032317Z",
            "signature": "LikEMZUP59FYZbF60ttNEipvgt4lk6R8ZIRgwyE3rQq6i8knMFox16h6tjz+l5JQ9Zf2yvmM4dWzxU+YXmcVCQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "3881E24B92FF4C9128E3F2CD05DFC7170B69428D",
            "timestamp": "2020-10-28T12:21:42.957670473Z",
            "signature": "cyWFBKpHljxHsczH1FjDWfvRXg6i3YENy+Wr7p9QudsszXm2sorHCOysf6iIFEPgcs+ZnRX+/rBBMZ3HK/QrCw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "6D9E4B4995037D608E365CE90436C24580ABCC33",
            "timestamp": "2020-10-28T12:21:42.832266593Z",
            "signature": "yVg18SFyQY2m1tUXR/29r8KXbiFNSQJXCOla98TZJO90VyAqOQi16fgnqJffspm6M983Ib2t+4rbrF4abcFuBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4D9F47C5A19D5550685D2D55DF2FF8C5D7504CEB",
            "timestamp": "2020-10-28T12:21:42.897727952Z",
            "signature": "rS7fTOLAn3yqXoV6i8wEQ4J+k0WQQ4EExKADSCYlffqMzvC1oOjsPE6lndGyboDtczW4yx0oisWafpwZba07AQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "6ED4882CF1C6C4E0E30C26BC964785E79390F15E",
            "timestamp": "2020-10-28T12:21:42.986527546Z",
            "signature": "p52ej7hv33vrmzB0suFVZbD75kHDwDsn+TH3MbzyN8lyKdbGJ+HF057xSQOkTwUmFZyOHgjT7KwRom5fuGqDAw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "679424928F3AB39FD197BEA70896D3D4E107D9FC",
            "timestamp": "2020-10-28T12:21:43.001891237Z",
            "signature": "uz9hmMgiA6izUp5PWP2IlE3QuiPWH7aAlFtlGFiwTvRMYPFM49qHovQ16njPlwJgUghx6rAaZB6Q1UcxxbZXCQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "0A530588FE95A76359D2C68E0FB70F0F48CBB858",
            "timestamp": "2020-10-28T12:21:42.971592485Z",
            "signature": "LNurb8xqaQczGv+raNjgEO7MRR2C7yPsEgKfJDQmwWp+fHig18s5tyNvhIMPaqOj+E5bRUIGd1/WzEmJvallCg=="
          },
          {
            "block_id_flag": 1,
            "validator_address": "",
            "timestamp": "0001-01-01T00:00:00Z",
            "signature": null
          },
          {
            "block_id_flag": 2,
            "validator_address": "12C28BD07CEB00778ADD9C879CA96A6BFF3B9F3C",
            "timestamp": "2020-10-28T12:21:43.048489628Z",
            "signature": "OQ90ISqd2nccWSQPq95hvC3p6Rpt3Pr9XlXzs3/qwkr9erVJB8D9d7eyGD24acJzdu0Ofuy3ejQ+0tjjo+gRDQ=="
          },
          {
            "block_id_flag": 1,
            "validator_address": "",
            "timestamp": "0001-01-01T00:00:00Z",
            "signature": null
          },
          {
            "block_id_flag": 2,
            "validator_address": "421966D9096396595CE3BBFDBA0A178E86671849",
            "timestamp": "2020-10-28T12:21:42.984586565Z",
            "signature": "dwCXbZzXpo+Sr7+jtsy0BHNt5wPsFnhQB0at0RBqiJ5OYri/plB4EF7F98P/ZYOCgF+RiqJTcBLHM/fqyN8OBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4F72618591F29B8B560471058C72B2042C3EF6AB",
            "timestamp": "2020-10-28T12:21:42.991807792Z",
            "signature": "t8ITiDfiSDX2D9DYOkKmx8liC4krGh892KIqF4HocvMtuauVgHBw99yg8pkJddzVaA+SesVK1uXRGyK/zYFMAQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "4FF7CC1247C075775995AAE3AEDA28F847581DEE",
            "timestamp": "2020-10-28T12:21:43.071663829Z",
            "signature": "tx1yyNLMfKaN++/2aFRI0Lc5RVN4ZZ5Ms7oOhNvLT9V3ag//8IoE7k5IknEi+eiu+cJPlhg8CQBn0T0xIC3gBw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "5640383EEFEEFAF39B1D6DDE1828088A9D878380",
            "timestamp": "2020-10-28T12:21:42.974974302Z",
            "signature": "9VVVZ4RzJQLfHeCtzNv/h9LWFy15ickYQGK03hTuRrvy0tR3dqvflU0Xp2fvELlG/jVqIWftyeel15RGEY/wDg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "657D474F1632C422A5593AAB4BDBF92AADAE6039",
            "timestamp": "2020-10-28T12:21:43.116798084Z",
            "signature": "aLkqXgiQe4WSn9+z9X84caS0zZv5T2d9bE31GXL/ReRkczrNToKlzwQWJ+SuCZcFTo5CrX4NWR1LpxrIcUkPBQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "72CDB6AE1290693343AD491883D445648556CABC",
            "timestamp": "2020-10-28T12:21:42.938686952Z",
            "signature": "VNkNR/ZfJUegd3nES3WX3MdxpGPIY8infb7Cyp3RxL0353Zs17aYeUw2xrBfK3hmq+OvPiLnq1jAUP0Qwyp2AA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "73F3DAAE50A59F67990065ECAA33BF89F49C44B0",
            "timestamp": "2020-10-28T12:21:54.420983297Z",
            "signature": "X6L/by+R1GCStKVOKI6OnYo+fULjBjTJz2Q68BnCzD7Dg/S8VqzTsShnIw1+ShAlP6y0f+7LGZypJ7NZtl9OCg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "8E6DAE6FBC033A6FE2119BC8DDE375E827387979",
            "timestamp": "2020-10-28T12:21:43.016882214Z",
            "signature": "sJBN06p6aZV4juqROHQEi/0gsTE2QFFcWi8Z42rDkBFGkPO1g81Op+4HC3x9jlyRD+85W4dmp2Q85jPBtkP3DQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "8F1C21289B3E146B0914A0235F8D39B1CCBFD068",
            "timestamp": "2020-10-28T12:21:43.005846918Z",
            "signature": "zSllYGiRpmPNJ7708Cq6rRt8pzF1OV0G7axMiEsiCJXEzFQwHilTSiNKOizrDWhmWocSdalKK1G1UXUN7cMhAg=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "ACE4DFDD67C5EA4A85D43C98FA716B72AB380D8E",
            "timestamp": "2020-10-28T12:21:42.995512614Z",
            "signature": "OrUqIjM5KfYd1lq7tR4N9m2gL1peXVYW0mo19fojjLeVwSx7UPU7dsBigyCqbyt5Mz4UlsY1M4By3AVQvtJ0AQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "B05BB1117138585B0FB415CC045E34CFA3413CA2",
            "timestamp": "2020-10-28T12:21:43.01405356Z",
            "signature": "BtYLrD/qne3EpA5fRhbWK/wrU9gHb3cxbDYAoJbSF0gB4lvpg0uBFJZwG5c2J0KMyKQAgP2UrYZjPHWd4FnjAw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "D1500A980A7F132A701407D00B1301D5F6B9D86F",
            "timestamp": "2020-10-28T12:21:43.004846859Z",
            "signature": "Yxy2Ig+qX83m0R/r1vZdz8sOK/bjhpuJMT0xRtlG95HnE5jmRGjsW9c2LxsDLKP5Sk49tDz1YScz13PEntX8Bw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E067FCE33F7FDBD0CE4872F8E240A7AD6E654726",
            "timestamp": "2020-10-28T12:21:43.131755766Z",
            "signature": "YmxbajYTH1QtqvkZTsnsC8nb5A8e2r7PcscXvfIlf71mCqI9pFrjD3fCC0T05VZTD9crXl0kEoLvJ3y7GM2hCQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E1D1889790E3502571F3CE6B347E90110D20D1FC",
            "timestamp": "2020-10-28T12:21:42.960018481Z",
            "signature": "T7vMp/kM1PUBRMboLi+ijs81MgL4sBYa74zphmyKp//+doFhjxrLDwaWyCTiqBHA4B1Gt/y5NvBM9rPcEwsPDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "E9F8AD395510009C0C5DABB67664B9E300F3AB30",
            "timestamp": "2020-10-28T12:21:42.968296135Z",
            "signature": "M5cIgyNTQhl9XSeVwF5twKp8M95wdTIkUJcyGi0DIr9r1a5KMWFWMH6sbuHJ+PHvIPyAqRm3Yp/3HUjfBQSsCw=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "FF87F5F1D995291C0439C6690CC79314E61DD762",
            "timestamp": "2020-10-28T12:21:43.065974533Z",
            "signature": "lUAcAflWGRwW4SQzsUSQc06Ua53+vQhzGbig5r/wRflzHLj4oRfkx8hl7rsr4XLXBSWgkJSb2Hs7fagOUhOGDQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "DB15AEDC38A039940170D75A8920255FD8C9AAFF",
            "timestamp": "2020-10-28T12:21:43.0628969Z",
            "signature": "JJmhRuENYT+7INznqYx3gugqipJm+xXI0fKL/dtmMJ8B9em2ue3ltgc1js/R5ER/+o+Zo4kcBXt7oC6lCnA/CQ=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "9B3C4955B744627F7D6B5CA7B3FBC5EC64014E11",
            "timestamp": "2020-10-28T12:21:42.969381682Z",
            "signature": "1NiUaiGAPBGxciukhBat9R/UpdcvHiCyqFd1nQjoyGHkU9nFDJvSIXihExt5lcCZOGSEbNJpdRu99PruvHoIAA=="
          },
          {
            "block_id_flag": 2,
            "validator_address": "2E6E44CE33AF193544CFCB92D25CD6983B21F9E2",
            "timestamp": "2020-10-28T12:21:37.311828929Z",
            "signature": "U9K5L6uAdtWgOh6hElSD3/togm7ANPXh+N+CNVcBV1WCyFyHMpEnkB1N4cOxyej17VwDM39F99GviUCfggpFCw=="
          }
        ]
      }
    }
  }
}`

const TX_FAILED_MSG_UNDELEGATE_BLOCK_RESULTS_RESP = `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "height": "184399",
    "txs_results": [
      {
        "code": 24,
        "data": null,
        "log": "failed to execute message; message index: 0: no delegation for (address, validator) tuple",
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "52906",
        "events": [],
        "codespace": "staking"
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
            "value": "MTY5NjU2OTczNzZiYXNldGNybw==",
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
            "value": "MC4wMDA1MzY1NTg5NzQ0MTIxNjU=",
            "index": true
          },
          {
            "key": "aW5mbGF0aW9u",
            "value": "MC4wMTMzNzk2NjEzMzY0MzI0NDU=",
            "index": true
          },
          {
            "key": "YW5udWFsX3Byb3Zpc2lvbnM=",
            "value": "MTA3MDc5MzM4MzAzNTQ3NzM4LjY0NTYwMzI5NTUyODcxMTYxMA==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTY5NjU2OTczNzY=",
            "index": true
          }
        ]
      },
      {
        "type": "transfer",
        "attributes": [
          {
            "key": "cmVjaXBpZW50",
            "value": "dGNybzFqdjY1czNncnFmNnY2amwzZHA0dDZjOXQ5cms5OWNkODMzOXA0bA==",
            "index": true
          },
          {
            "key": "c2VuZGVy",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          },
          {
            "key": "YW1vdW50",
            "value": "MTY5NjU3MTczNzZiYXNldGNybw==",
            "index": true
          }
        ]
      },
      {
        "type": "message",
        "attributes": [
          {
            "key": "c2VuZGVy",
            "value": "dGNybzE3eHBmdmFrbTJhbWc5NjJ5bHM2Zjg0ejNrZWxsOGM1bHhoemFoYQ==",
            "index": true
          }
        ]
      },
      {
        "type": "proposer_reward",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODMyNDgyMzM1LjM3Nzk4OTY4MjQ0NjIwMDU3NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOHlsY2hnbXh5cGh3M2N0c2w3NW41M3VqZXF1a21tYWcybjZ4M2Y=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODMyNDgyMzMuNTM3Nzk4OTY4MjQ0NjIwMDU4YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOHlsY2hnbXh5cGh3M2N0c2w3NW41M3VqZXF1a21tYWcybjZ4M2Y=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "ODMyNDgyMzM1LjM3Nzk4OTY4MjQ0NjIwMDU3NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOHlsY2hnbXh5cGh3M2N0c2w3NW41M3VqZXF1a21tYWcybjZ4M2Y=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NzcwMjI2NTUuMjk0NDEzMTM0OTA0NTcxMDQ4YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxajdwZWo4a3BsZW00d3Q1MHA0aGZ2bmRodXc1anByeHh4dGVudnI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NzcwMjI2NTUyLjk0NDEzMTM0OTA0NTcxMDQ3NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxajdwZWo4a3BsZW00d3Q1MHA0aGZ2bmRodXc1anByeHh4dGVudnI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mzc1MzM5NTMxLjMxODIyNzIxNTM2ODczMjcxMWJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeHdkM2s4eHRlcmRlZnQzbnhxZzkyc3pocHo2dng0M3FzcGRwdzY=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NzUwNjc5MDYyLjYzNjQ1NDQzMDczNzQ2NTQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeHdkM2s4eHRlcmRlZnQzbnhxZzkyc3pocHo2dng0M3FzcGRwdzY=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjgxNzcxMDIuODI5NTgyNzQ1Mzc3MDQxODI2YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNWdyZnRnODhsMGdkdzRtZzl0OXB3bmwwcGRlMmFzanpla3owZWs=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjgxNzcxMDI4LjI5NTgyNzQ1Mzc3MDQxODI1NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNWdyZnRnODhsMGdkdzRtZzl0OXB3bmwwcGRlMmFzanpla3owZWs=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjMxODQwODUuMjQ5MTIzNTU5ODEzOTg4NjM1YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOHlsY2hnbXh5cGh3M2N0c2w3NW41M3VqZXF1a21tYWcybjZ4M2Y=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NjMxODQwODUyLjQ5MTIzNTU5ODEzOTg4NjM1MGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOHlsY2hnbXh5cGh3M2N0c2w3NW41M3VqZXF1a21tYWcybjZ4M2Y=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjgyMjcwNTg4Ljc1MjA3NDU2NzQ0NzQ2MTMyMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxY3VxMmpoZGhnaHV4d3BmOXQyZDAzdmxjZW1tNG5mdjA4cjRxZ2w=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mzc2MzYwNzg1LjAwMjc2NjA4OTkyOTk0ODQyN2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxY3VxMmpoZGhnaHV4d3BmOXQyZDAzdmxjZW1tNG5mdjA4cjRxZ2w=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NzUyNjA1MzIuNzgyNDQyODM2MjIyMDc5NDM3YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOTIzcHowM21oamF6dGdjdjNnZXkwaGowYW13eDAyZHlza2F1NTI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mzc2MzAyNjYzLjkxMjIxNDE4MTExMDM5NzE4M2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxOTIzcHowM21oamF6dGdjdjNnZXkwaGowYW13eDAyZHlza2F1NTI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzczMTk4MjUzLjc4Njg5NTc5MTU5NzA4NjExOGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxd3lwazB1bmhnOTQzMmtkejZobXVtcXFqZDBsejgzcDNtYzQydHk=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzczMTk4MjUzLjc4Njg5NTc5MTU5NzA4NjExOGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxd3lwazB1bmhnOTQzMmtkejZobXVtcXFqZDBsejgzcDNtYzQydHk=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDA5OTgyOTkuMDkzMzM3MjgyODM3MDUzOTk0YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZnM4cjZ6eG1yNW5jODZqOGNwY21qbWNjZjhzMmNhZnh6dDVhbHE=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzcyNzExODA5LjkzOTQyOTg0Mzk3MzIxODEyN2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZnM4cjZ6eG1yNW5jODZqOGNwY21qbWNjZjhzMmNhZnh6dDVhbHE=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg2MjAwMjg5LjIyNTAwNTE3MDAxNDI5NDk0MGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjk4Z3RsNjlxYXc2ODh1ZXd0Z2FoanZkMHBjZnQ2eGo1MzJjOXI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzcyNDAwNTc4LjQ1MDAxMDM0MDAyODU4OTg4MGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjk4Z3RsNjlxYXc2ODh1ZXd0Z2FoanZkMHBjZnQ2eGo1MzJjOXI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTExNDM3NDYyLjcwNzE2Nzc0MTgyMTcyMjU4N2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNzJ2OWFnYTZrNW5scnc2dWMzODdlZ3psczA4bmhsNGN5em5jbW4=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzcxNDU4MjA5LjAyMzg5MjQ3MjczOTA3NTI4OWJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNzJ2OWFnYTZrNW5scnc2dWMzODdlZ3psczA4bmhsNGN5em5jbW4=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzcwMzg5ODcwLjQ5MzM3MTU4ODI5MTY3MTU3M2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZGVzY244aDdrajUyZW44Z245ajlkcXd5eTQ5NW1ueHowbnU2Zms=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzcwMzg5ODcwLjQ5MzM3MTU4ODI5MTY3MTU3M2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZGVzY244aDdrajUyZW44Z245ajlkcXd5eTQ5NW1ueHowbnU2Zms=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzYzMjU5MzU5LjQ1MTU0ODkxNDk5ODc3MDI3MmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZmdhZTg1cmd6djU3a2QyM2hrdXg0a3RqcXRzY2o3NWs0cnk1NmU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzYzMjU5MzU5LjQ1MTU0ODkxNDk5ODc3MDI3MmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZmdhZTg1cmd6djU3a2QyM2hrdXg0a3RqcXRzY2o3NWs0cnk1NmU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzUzNTQ1NjcuMjEzMTc3NzUwOTE5ODU1MTc4YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdXJtcnJtbXQ2Z2RmMDc3ZG1ndDk1Y21qNnRjMHo5MDRwamhscmQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzUzNTQ1NjcyLjEzMTc3NzUwOTE5ODU1MTc3NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdXJtcnJtbXQ2Z2RmMDc3ZG1ndDk1Y21qNnRjMHo5MDRwamhscmQ=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTQxMTM2Mzg3LjE5OTA2MzQxMzY4MTkyMTQ1OGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdWV2bXMybnY0ZjJkaHZtNXU3c2d1czJ5bmNnaDdnZHd4OWw2azY=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzUyODQwOTY3Ljk5NzY1ODUzNDIwNDgwMzY0NGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdWV2bXMybnY0ZjJkaHZtNXU3c2d1czJ5bmNnaDdnZHd4OWw2azY=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzMxODUyODguNjE2MDQ5MDg3ODExNDczNDA3YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxcmV5c2hmZHlnZjc2NzN4bTlwOHYweHZ0ZDk2bTZjZDZjYW5odTM=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzMxODUyODg2LjE2MDQ5MDg3ODExNDczNDA2OWJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxcmV5c2hmZHlnZjc2NzN4bTlwOHYweHZ0ZDk2bTZjZDZjYW5odTM=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI3NzQwNjguMTg3MjczNDU5Nzg5MDM5MDQ5YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMGo0NW1xY3g5bXM4aHB4MzM0bGZhdzlyeXkydXNwYWNscHo3YzI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI3NzQwNjgxLjg3MjczNDU5Nzg5MDM5MDQ5MGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMGo0NW1xY3g5bXM4aHB4MzM0bGZhdzlyeXkydXNwYWNscHo3YzI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI3MjE5NDEuMDQ1NTE3MDYzNDIzNDY4MzQ1YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxa3NjNDd1dGEwMjIza2hsanNqemd0dnpqOGdmbWtleHk2cjQyazk=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI3MjE5NDEwLjQ1NTE3MDYzNDIzNDY4MzQ1NGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxa3NjNDd1dGEwMjIza2hsanNqemd0dnpqOGdmbWtleHk2cjQyazk=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI2OTkxNDcuNTg0MzQ5MDM4ODI1NTkwOTk5YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMjJ3OWZoYzBwdTNleTlyNmhla3puZDJma2w1anN3bDVhcXN2Z3k=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI2OTkxNDc1Ljg0MzQ5MDM4ODI1NTkwOTk4OGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMjJ3OWZoYzBwdTNleTlyNmhla3puZDJma2w1anN3bDVhcXN2Z3k=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI2OTQzMTguMzExNTc2MjI1ODUwOTY2NTgxYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdHh0OTMweHV4bGZrd2Y4a25laDV6eXRlMmNoN3dwdjczc3d4eTI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI2OTQzMTgzLjExNTc2MjI1ODUwOTY2NTgwN2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdHh0OTMweHV4bGZrd2Y4a25laDV6eXRlMmNoN3dwdjczc3d4eTI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI2ODg4MDMuMzgyNzU3Nzk0MDYyODM0MDgxYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMHcycWYyOWYwODc3OWwyNHoydjM5cm52cGhuZ3Fma2x1cnY3ZWg=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI2ODg4MDMzLjgyNzU3Nzk0MDYyODM0MDgxM2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMHcycWYyOWYwODc3OWwyNHoydjM5cm52cGhuZ3Fma2x1cnY3ZWg=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI2NjkyMTYuMTYxOTY5NjY2MTExNzQ4NzE5YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNTJlbmE3NWdoNW5xbnUybmxhcndtcHp4YTJjenhzOHlzeGpmODU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI2NjkyMTYxLjYxOTY5NjY2MTExNzQ4NzE5MWJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNTJlbmE3NWdoNW5xbnUybmxhcndtcHp4YTJjenhzOHlzeGpmODU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI2NjI3OTYuOTE4NTc0OTAyNDYxMTg2ODkzYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeGdkMDV2dWZuY2FmeDh0Y25zdjc3dWN1bWhoMHV6OHh0N2Q1N2c=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzI2NjI3OTY5LjE4NTc0OTAyNDYxMTg2ODkyN2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeGdkMDV2dWZuY2FmeDh0Y25zdjc3dWN1bWhoMHV6OHh0N2Q1N2c=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTYwNzExMzEzLjE0MTExMzY1MTUxMjk3NzAzM2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNHpxbjVtNnEyZXhsbTI5Zmg4ajVyc2FjbDg2ajRtcXBhYTNseXg=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzIxNDIyNjI2LjI4MjIyNzMwMzAyNTk1NDA2NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNHpxbjVtNnEyZXhsbTI5Zmg4ajVyc2FjbDg2ajRtcXBhYTNseXg=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzE5NzQ5NTAuNTMwNjI5NTgxMDM4NTk3Mzg5YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxc3J1emQ1MjlsaGpqdTZoZmN3ZDJmeHAzdjBlN3AwdnFxdG1lNzY=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzE5NzQ5NTA1LjMwNjI5NTgxMDM4NTk3Mzg5M2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxc3J1emQ1MjlsaGpqdTZoZmN3ZDJmeHAzdjBlN3AwdnFxdG1lNzY=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzE5MjE1MTUuNTcwNDEwMjgxNzI0MjkxMTY0YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxN3hmdjByZjdsZ2xjZ3FodnV1cDZubDlwYWpqcWpsdm0ydW11ZGw=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzE5MjE1MTU1LjcwNDEwMjgxNzI0MjkxMTY0MmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxN3hmdjByZjdsZ2xjZ3FodnV1cDZubDlwYWpqcWpsdm0ydW11ZGw=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjkzNDA0NTIuOTkxNTkxNzgyODg1Mzg5MDI5YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxcG0yN2RqY3M1ZGp4anN4dzN1bnJrdjNtM2p0eGRleGt0dzVlcHU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjkzNDA0NTI5LjkxNTkxNzgyODg1Mzg5MDI5MmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxcG0yN2RqY3M1ZGp4anN4dzN1bnJrdjNtM2p0eGRleGt0dzVlcHU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTAyMDEzNDMwLjY4OTk0MjUxMTg5MDczOTQ2NGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZ2RmejN3Y3ZxYXlrbGZyYzZqdzN3ZzY0ZHhybnJtNzlndHE1Z3I=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjkxNDY2OTQ0LjgyODQwNzE3NjgzMDY4NDE4NGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZ2RmejN3Y3ZxYXlrbGZyYzZqdzN3ZzY0ZHhybnJtNzlndHE1Z3I=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NzI3NDk2MTMuOTQ0MTk2OTY0NDI4NTg2NzAzYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdDcweGozNGhtdWQ0a2F1dHE3djZnazNyZ3Z6NmgzN254bGZtNTc=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MjkwOTk4NDU1Ljc3Njc4Nzg1NzcxNDM0NjgxM2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdDcweGozNGhtdWQ0a2F1dHE3djZnazNyZ3Z6NmgzN254bGZtNTc=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mjc4MjUwNzQuOTg0NzcwMTk4ODUzNTM1MjEwYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdGgydjh6cmtsdTB0djBxY2hxMnZzMmp5aHB6YWRlMjJscnlybmo=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mjc4MjUwNzQ5Ljg0NzcwMTk4ODUzNTM1MjA5OGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdGgydjh6cmtsdTB0djBxY2hxMnZzMmp5aHB6YWRlMjJscnlybmo=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg4MjY0MjQuMzUzOTQ4NzkzODgxNjQ1NzIyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNnl6Y3ozdHk5NGF3cjducjJ0eGVrOWRwMmtscDJhdjl2aDQzN3M=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg4MjY0MjQzLjUzOTQ4NzkzODgxNjQ1NzIyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNnl6Y3ozdHk5NGF3cjducjJ0eGVrOWRwMmtscDJhdjl2aDQzN3M=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3OTI2MTEuMTc5MDk2MzYwNDM2MzA4NDExYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNmtxcjAwOXB0Z2tlbjZxc3huemZueWpmc3E2cTk3ZzN1ZWRjZXI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3OTI2MTExLjc5MDk2MzYwNDM2MzA4NDExM2Jhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNmtxcjAwOXB0Z2tlbjZxc3huemZueWpmc3E2cTk3ZzN1ZWRjZXI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODkyMjkuODYxNjExMTE3NDE0NDM5MzgxYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxd3Z3NjQwNWFueWF2ZzlhOHh0YWF4Z2UwZmU4c2V0cWtuY3B5Z2Q=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODkyMjk4LjYxNjExMTE3NDE0NDM5MzgxMGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxd3Z3NjQwNWFueWF2ZzlhOHh0YWF4Z2UwZmU4c2V0cWtuY3B5Z2Q=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODg4NTQuMTU5NjY4MzEyMjc1NzE1Mzc3YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxanQ1dWNlcjU5MjU3bWprcjg2bDVwZ2MyN3o2aHY4N2NjZHBoZGU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODg4NTQxLjU5NjY4MzEyMjc1NzE1Mzc2OGJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxanQ1dWNlcjU5MjU3bWprcjg2bDVwZ2MyN3o2aHY4N2NjZHBoZGU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxbnZzazJoOTdxbHJ0c3pqM3V0MDZkdDhkeHN3MjVsYWU1a3U5amQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxbnZzazJoOTdxbHJ0c3pqM3V0MDZkdDhkeHN3MjVsYWU1a3U5amQ=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNWU2OWtkcnRjemFqamRsenl0MnFnczVxMmFuYzVxcG1rMmM2OHo=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNWU2OWtkcnRjemFqamRsenl0MnFnczVxMmFuYzVxcG1rMmM2OHo=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mzc1NzAxOTQuMjgwNDgwNTI4MjMwMjQ0Njg0YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZ3MwNGNjeWozYTR5cDNxMGozZXEwMmdsbXpxeGthZDQ0eHRjdTI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZ3MwNGNjeWozYTR5cDNxMGozZXEwMmdsbXpxeGthZDQ0eHRjdTI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxbGxzdDBjZ3VoNWF6bDl0OHdyNm16NXl6anV3dWt6N2YwcHA4M2U=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxbGxzdDBjZ3VoNWF6bDl0OHdyNm16NXl6anV3dWt6N2YwcHA4M2U=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxM2hqYzAzZnZ2Z2gwbXAzcWF2cHBqd2Zqd3ZubndjYzMzMmp1NHQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxM2hqYzAzZnZ2Z2gwbXAzcWF2cHBqd2Zqd3ZubndjYzMzMmp1NHQ=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNnAwdW04ZjIwc3E3N3hxbHBxcW14NHQ1cXd5Y3pnam1qdDY5bjU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNnAwdW04ZjIwc3E3N3hxbHBxcW14NHQ1cXd5Y3pnam1qdDY5bjU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdzg3cGpxMHl5eXV6dTZqOXZlamY3NXVhcXh6ZXp2ZWRnNXkzN20=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdzg3cGpxMHl5eXV6dTZqOXZlamY3NXVhcXh6ZXp2ZWRnNXkzN20=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdXZ2bXplczlrYXpwa3QzNTlleG02N3FxajM4NGw3YzdxeTMzbXE=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdXZ2bXplczlrYXpwa3QzNTlleG02N3FxajM4NGw3YzdxeTMzbXE=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxbXA0MGFnNnhncHpxZzJkbWhmZ3I4NzJ1czd0Mnl3cnI3eXhyOHU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxbXA0MGFnNnhncHpxZzJkbWhmZ3I4NzJ1czd0Mnl3cnI3eXhyOHU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNGx6ZDZxNzNwdjhmam1lYXFybjN0ZWM3ZTg5MzB1dTdxOGRlZWY=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNGx6ZDZxNzNwdjhmam1lYXFybjN0ZWM3ZTg5MzB1dTdxOGRlZWY=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeG1mM2U0dWE1dGhmZXNlbTh0eWpkeDM4cmdrNnVrZHIwNDY5NnI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeG1mM2U0dWE1dGhmZXNlbTh0eWpkeDM4cmdrNnVrZHIwNDY5NnI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeTBlZTdrNzU3dWZ6bm44ZXk0NDUzd2Q5MmV0ejY1emx3ZTVxYXg=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxeTBlZTdrNzU3dWZ6bm44ZXk0NDUzd2Q5MmV0ejY1emx3ZTVxYXg=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdGpyMDhjcmU3OTl1anczOWYzZ3drdjlsczloMjIyY2F2cHY3OWY=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdGpyMDhjcmU3OTl1anczOWYzZ3drdjlsczloMjIyY2F2cHY3OWY=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNmNlNDRleTh6M3Q3cjl3YzA1enA5NXVnN2NhcDZwZjVnOTkyeng=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNmNlNDRleTh6M3Q3cjl3YzA1enA5NXVnN2NhcDZwZjVnOTkyeng=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxaGo3d2Z6eXl3a3pscDNtMHUycHR2enlwanNsNXg1c2VhNHN3dnM=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxaGo3d2Z6eXl3a3pscDNtMHUycHR2enlwanNsNXg1c2VhNHN3dnM=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZXpwejRhODhjcTdxdmthZGR3bDdoc3BocTc0eG56NHNkc2d6ZzI=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZXpwejRhODhjcTdxdmthZGR3bDdoc3BocTc0eG56NHNkc2d6ZzI=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mzc1NzAxOTQuMjgwNDgwNTI4MjMwMjQ0Njg0YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMGdzcXM4anpkbHJlbTgwc2hwMHg2d3gwanc3cXU3bThjZDI5eTU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMGdzcXM4anpkbHJlbTgwc2hwMHg2d3gwanc3cXU3bThjZDI5eTU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "NDY5NjI3NDIuODUwNjAwNjYwMjg3ODA1ODU2YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMzVxZ2wyaHJxenQ0N2g5ZTI4ODR5eGVzN2pteTJrcnNkaHJ1NHQ=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxMzVxZ2wyaHJxenQ0N2g5ZTI4ODR5eGVzN2pteTJrcnNkaHJ1NHQ=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdTk2cHdoejI5dDR6ZDIyeDl1eGd0YTczY3M4djhkeWFhc2VraGo=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxdTk2cHdoejI5dDR6ZDIyeDl1eGd0YTczY3M4djhkeWFhc2VraGo=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcuMTQwMjQwMjY0MTE1MTIyMzQyYmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjhldDY5Mmd4aHJ2cGNqcGRqMmRuN3RzemM4amN1dDZ2MzJ1amU=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3ODUwOTcxLjQwMjQwMjY0MTE1MTIyMzQyMmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxNjhldDY5Mmd4aHJ2cGNqcGRqMmRuN3RzemM4amN1dDZ2MzJ1amU=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3NzAwNjUuMzA1NTA4NjQ0MDM0NDI1ODk2YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZHMyYTJzZXRrdm14ajVzbHg4YXkycDk0bnQ2cmVrbjB4cmZseGw=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg3NzAwNjUzLjA1NTA4NjQ0MDM0NDI1ODk1NmJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZHMyYTJzZXRrdm14ajVzbHg4YXkycDk0bnQ2cmVrbjB4cmZseGw=",
            "index": true
          }
        ]
      },
      {
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MTg0MDkzOTUuMTk3NDM1NDU5MzQ5MDgzNDE3YmFzZXRjcm8=",
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
            "value": "MTg0MDkzOTUxLjk3NDM1NDU5MzQ5MDgzNDE2N2Jhc2V0Y3Jv",
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
        "type": "commission",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "MzcuNTcwMTk0Mjc5MjIzMjEzNTk3YmFzZXRjcm8=",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZjBxMGs0eXlzYXZreHM3NWE4M3c3MDM4NGRxdTd2bnh4a2Q4Mjg=",
            "index": true
          }
        ]
      },
      {
        "type": "rewards",
        "attributes": [
          {
            "key": "YW1vdW50",
            "value": "Mzc1LjcwMTk0Mjc5MjIzMjEzNTk3MWJhc2V0Y3Jv",
            "index": true
          },
          {
            "key": "dmFsaWRhdG9y",
            "value": "dGNyb2NuY2wxZjBxMGs0eXlzYXZreHM3NWE4M3c3MDM4NGRxdTd2bnh4a2Q4Mjg=",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "dGNyb2NuY2xjb25zMXA2eGF6bmMwNGxjMG12dW0wdm5wNzNza2Y4cmM4MDU3cnE1NHpr",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "Mjg3",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTg0Mzk5",
            "index": true
          }
        ]
      },
      {
        "type": "liveness",
        "attributes": [
          {
            "key": "YWRkcmVzcw==",
            "value": "dGNyb2NuY2xjb25zMTgyZHllNnhndXpoeGg3czB4eDRneHk1dG5ud3JzcWRzNXV6M3dn",
            "index": true
          },
          {
            "key": "bWlzc2VkX2Jsb2Nrcw==",
            "value": "MTI=",
            "index": true
          },
          {
            "key": "aGVpZ2h0",
            "value": "MTg0Mzk5",
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

const TX_FAILED_MSG_UNDELEGATE_TXS_RESP = `{
    "tx": {
        "body": {
            "messages": [
                {
                    "@type": "/cosmos.staking.v1beta1.MsgUndelegate",
                    "delegator_address": "tcro1llst0cguh5azl9t8wr6mz5yzjuwukz7f67z7f6",
                    "validator_address": "tcrocncl15e69kdrtczajjdlzyt2qgs5q2anc5qpmk2c68z",
                    "amount": {
                        "denom": "basetcro",
                        "amount": "50000000000000"
                    }
                }
            ],
            "memo": "",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
        },
        "auth_info": {
            "signer_infos": [
                {
                    "public_key": {
                        "@type": "/cosmos.crypto.secp256k1.PubKey",
                        "key": "A8YD+h8qnYctS4oyIOjZy3919pvXEZ+eWgPw1rjZgkrr"
                    },
                    "mode_info": {
                        "single": {
                            "mode": "SIGN_MODE_DIRECT"
                        }
                    },
                    "sequence": "2"
                }
            ],
            "fee": {
                "amount": [
                    {
                        "denom": "basetcro",
                        "amount": "5000"
                    }
                ],
                "gas_limit": "200000",
                "payer": "",
                "granter": ""
            }
        },
        "signatures": [
            "0aXlcyCxEWPiHMA0aTBa2bBDZ76M1oHeIPoIO5wzoDItcENuhe/WSJmWa1NLcLOIggz9I12XmGCnXND8G/Kolw=="
        ]
    },
    "tx_response": {
        "height": "184399",
        "txhash": "5285A9B475157E01540536299A2B5F505AC900159C268B3D90652557F9ACDE1E",
        "codespace": "staking",
        "code": 24,
        "data": "",
        "raw_log": "failed to execute message; message index: 0: no delegation for (address, validator) tuple",
        "logs": [
            {
                "msg_index": 0,
                "log": "",
                "events": []
            }
        ],
        "info": "",
        "gas_wanted": "200000",
        "gas_used": "52906",
        "tx": {
            "@type": "/cosmos.tx.v1beta1.Tx",
            "body": {
                "messages": [
                    {
                        "@type": "/cosmos.staking.v1beta1.MsgUndelegate",
                        "delegator_address": "tcro1llst0cguh5azl9t8wr6mz5yzjuwukz7f67z7f6",
                        "validator_address": "tcrocncl15e69kdrtczajjdlzyt2qgs5q2anc5qpmk2c68z",
                        "amount": {
                            "denom": "basetcro",
                            "amount": "50000000000000"
                        }
                    }
                ],
                "memo": "",
                "timeout_height": "0",
                "extension_options": [],
                "non_critical_extension_options": []
            },
            "auth_info": {
                "signer_infos": [
                    {
                        "public_key": {
                            "@type": "/cosmos.crypto.secp256k1.PubKey",
                            "key": "A8YD+h8qnYctS4oyIOjZy3919pvXEZ+eWgPw1rjZgkrr"
                        },
                        "mode_info": {
                            "single": {
                                "mode": "SIGN_MODE_DIRECT"
                            }
                        },
                        "sequence": "2"
                    }
                ],
                "fee": {
                    "amount": [
                        {
                            "denom": "basetcro",
                            "amount": "5000"
                        }
                    ],
                    "gas_limit": "200000",
                    "payer": "",
                    "granter": ""
                }
            },
            "signatures": [
                "0aXlcyCxEWPiHMA0aTBa2bBDZ76M1oHeIPoIO5wzoDItcENuhe/WSJmWa1NLcLOIggz9I12XmGCnXND8G/Kolw=="
            ]
        },
        "timestamp": "2021-08-29T17:15:46Z"
    }
}`
