package blueprint

type WaitCondition struct {
	Type        WaitConditionType        `json:"type"`
	CompareType WaitConditionCompareType `json:"compare_type"`
	Ticks       *uint                    `json:"ticks,omitempty"`
	Condition   any                      `json:"condition,omitempty"` // FIXME
}
