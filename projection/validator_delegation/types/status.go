package types

type ValidatorStatus string

const (
	// UNBONDED defines a validator that is not bonded.
	Unbonded ValidatorStatus = "UNBONDED"
	// UNBONDING defines a validator that is unbonding.
	Unbonding ValidatorStatus = "UNBONDING"
	// BONDED defines a validator that is bonded.
	Bonded ValidatorStatus = "BONDED"
)
