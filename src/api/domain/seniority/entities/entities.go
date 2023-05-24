package entities

type SeniorityRequest struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Seniority string `json:"seniority"`
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
