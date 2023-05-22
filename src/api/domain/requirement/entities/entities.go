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

type RequirementData struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Recommendation string `json:"recommendation"`
	Points         int    `json:"points"`
}

type RequirementBasicData struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
