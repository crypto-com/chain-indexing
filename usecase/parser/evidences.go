package parser

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

func ParseEvidencesCommands(
	blockHeight int64,
	evidences []model.BlockEvidence,
) ([]command.Command, error) {
	commands := make([]command.Command, 0)

	for i := 0; i < len(evidences); i++ {
		evidence := evidences[i]

		infractionHeight, err := strconv.ParseInt(evidence.Value.VoteA.Height, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error in parsing evidence.Value.VoteA.Height to int64: %v", err)
		}

		commands = append(commands, command_usecase.NewCreateEvidence(blockHeight, model.EvidenceParams{
			TendermintAddress: evidence.Value.VoteA.ValidatorAddress, // the same as evidence.Value.VoteB.ValidatorAddress
			InfractionHeight:  infractionHeight,                      // the same as evidence.Value.VoteB.Height
			RawEvidence:       evidence,
		}))
	}

	return commands, nil
}
