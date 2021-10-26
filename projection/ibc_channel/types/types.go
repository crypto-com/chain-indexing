package types

type Status = string

const (
	STATUS_NOT_ESTABLISHED Status = "NOT_ESTABLISHED"
	STATUS_OPENED          Status = "OPENED"
	STATUS_CLOSED          Status = "CLOSED"
)
