package domain

import (
	"fmt"
)

type OperationStatus string

const (
	SUCCEEDED OperationStatus = "SUCCEEDED"
	FAILED    OperationStatus = "FAILED"
	COMPLETED OperationStatus = "COMPLETED"
	STARTING  OperationStatus = "STARTING"
	DOWNLOADING OperationStatus = "DOWNLOADING"
	FRAGMENTING OperationStatus = "FRAGMENTING"
	ENCODING OperationStatus = "ENCODING"
	UPLOADING OperationStatus = "UPLOADING"
	FINISHNING OperationStatus = "FINISHING"
)

var stringToOperationStatus = map[string]OperationStatus{
	"SUCCEEDED": SUCCEEDED,
	"FAILED": FAILED,
	"COMPLETED": COMPLETED,
	"STARTING": STARTING,
	"DOWNLOADING": DOWNLOADING,
	"FRAGMENTING": FRAGMENTING,
	"ENCODING": ENCODING,
	"UPLOADING": UPLOADING,
	"FINISHNING": FINISHNING,
}

func (os OperationStatus) ToString() string {
	return string(os)
}

func OperationStatusFromString(s string) (OperationStatus, error) {
	status, ok := stringToOperationStatus[s]
	if !ok {
		return "", fmt.Errorf("unknown OperationStatus: %s", s)
	}	
	return status, nil
}

