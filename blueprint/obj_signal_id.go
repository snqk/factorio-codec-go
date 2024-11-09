package blueprint

type SignalID struct {
	Name string        `json:"name"`
	Type *SignalIDType `json:"type,omitempty"`
}
