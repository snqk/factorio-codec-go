package blueprint

type Schedule struct {
	Schedule    []ScheduleRecord `json:"schedule"`
	Locomotives []int            `json:"locomotives"`
}
