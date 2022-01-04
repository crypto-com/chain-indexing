package types

type ValidatorStatus string

const (
	// UNBONDED defines a validator that is not bonded.
	UNBONDED ValidatorStatus = "Unbonded"
	// UNBONDING defines a validator that is unbonding.
	UNBONDING ValidatorStatus = "Unbonding"
	// BONDED defines a validator that is bonded.
	BONDED ValidatorStatus = "Bonded"
)
