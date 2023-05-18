package entities

import "mime/multipart"

type RequirementFile struct {
	RequirementId int                   `form:"requirement_id"`
	UserId        int                   `form:"user_id"`
	File          *multipart.FileHeader `form:"file"`
}

type UserRequirementStatus struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type (
	Status     string
	StatusList []Status
)

const (
	PendingStatus   Status = "PENDING"
	CompletedStatus Status = "COMPLETED"
	RejectedStatus  Status = "REJECTED"
)

var PossibleStatus = StatusList{PendingStatus, CompletedStatus, RejectedStatus}

func (status StatusList) Contains(s Status) bool {
	for _, st := range status {
		if st == s {
			return true
		}
	}
	return false
}
