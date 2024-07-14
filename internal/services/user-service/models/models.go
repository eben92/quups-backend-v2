package models

type Role string

const (
	ADMIN_ROLE      Role = "ADMIN"
	DISPATCHER_ROLE Role = "DISPATCHER"
	WAITER_ROLE     Role = "WAITER"
	AGENT_ROLE      Role = "AGENT"
)

type Status string

const (
	ACTIVE_STATUS   Status = "ACTIVE"
	INACTIVE_STATUS Status = "INACTIVE"
)
