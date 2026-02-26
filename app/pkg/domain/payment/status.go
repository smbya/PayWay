package payment

type Status string

const (
	STATUS_NEW        Status = "NEW"
	STATUS_PROCESSING Status = "PROCESSING"
	STATUS_COMPLETED  Status = "COMPLETED"
	STATUS_FAILED     Status = "FAILED"
	STATUS_CANCELED   Status = "CANCELED"
)
