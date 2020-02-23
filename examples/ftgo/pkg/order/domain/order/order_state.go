package order

type OrderState int

const (
	OrderState_Unknown OrderState = iota
	OrderState_Approved
	OrderState_ApprovalPending
	OrderState_Canceled
	OrderState_CancelPending
	OrderState_Rejected
	OrderState_RevisionPending
)
