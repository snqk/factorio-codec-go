package blueprint

type ScheduleRecord struct {
	Station        string          `json:"station"`
	WaitConditions []WaitCondition `json:"wait_conditions"`
	Temporary      *bool           `json:"temporary,omitempty"`
}
