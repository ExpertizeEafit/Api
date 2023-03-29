package entities

type (
	SeniorityInfo struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		PriorTo     []int  `json:"priorTo"`
	}
	RequirementInfo struct {
		Name           string `json:"name"`
		Description    string `json:"description"`
		Recommendation string `json:"recommendation"`
		Status         string `json:"status"`
	}
	SeniorityInfoExtended struct {
		SeniorityInfo
		Requirements map[int]RequirementInfo `json:"requirements"`
	}
	Path         map[int]SeniorityInfo
	PathExtended map[int]SeniorityInfoExtended
)
