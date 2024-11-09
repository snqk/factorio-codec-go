package blueprint

type Entity struct {
	EntityNumber int      `json:"entity_number"`
	Name         string   `json:"name"`
	Position     Position `json:"position"`
	Direction    *int     `json:"direction,omitempty"`
	Type         *string  `json:"type,omitempty"` // Optional ("input" or "output")
}
