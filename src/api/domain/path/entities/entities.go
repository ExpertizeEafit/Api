package entities

type (
	SeniorityInfo struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		PriorTo     []int  `json:"priorTo"`
	}
	SeniorityInfoExtended struct {
		SeniorityInfo
	}
	Path map[int]SeniorityInfo
)
