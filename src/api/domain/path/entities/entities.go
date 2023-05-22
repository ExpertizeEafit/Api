package entities

type (
	SeniorityInfo struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		PriorTo     []int  `json:"priorTo"`
	}
	RequirementInfo struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		Recommendation string `json:"recommendation"`
		Status         string `json:"status"`
	}
	SeniorityInfoExtended struct {
		SeniorityInfo
		Requirements []RequirementInfo `json:"requirements"`
	}
	Path         map[int]SeniorityInfo
	PathExtended map[int]SeniorityInfoExtended
)
