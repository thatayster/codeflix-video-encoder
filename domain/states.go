package domain

type OperationStatus int8

const (
	SUCCEEDED OperationStatus = iota
	FAILED
	COMPLETED
	STARTING
)

func (status OperationStatus) String() string {
	switch status {
	case SUCCEEDED:
		return "succeeded"
	case FAILED:
		return "failed"
	case COMPLETED:
		return "completed"
	case STARTING:
		return "starting"
	}
	return "unknown"
}
