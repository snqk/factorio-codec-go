package blueprint

type InfinityFilter struct {
	Name  string             `json:"name"`
	Count uint32             `json:"count"`
	Mode  InfinityFilterMode `json:"mode"`
	Index int                `json:"index"`
}
