package blueprint

import (
	"encoding/json"
	"fmt"
	"strings"
)

type SignalIDType uint

const (
	SignalIDTypeUnknown SignalIDType = iota
	SignalIDTypeItem
	SignalIDTypeFluid
	SignalIDTypeVirtual
)

func (s SignalIDType) String() string {
	return map[SignalIDType]string{
		SignalIDTypeItem:    "item",
		SignalIDTypeFluid:   "fluid",
		SignalIDTypeVirtual: "virtual",
	}[s]
}

func ParseSignalIDType(s string) (SignalIDType, error) {
	switch strings.ToLower(s) {
	case SignalIDTypeItem.String():
		return SignalIDTypeItem, nil
	case SignalIDTypeFluid.String():
		return SignalIDTypeFluid, nil
	case SignalIDTypeVirtual.String():
		return SignalIDTypeVirtual, nil
	}

	return SignalIDTypeUnknown, fmt.Errorf("unknown SignalIDType: %q", s)
}

func (s SignalIDType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *SignalIDType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	p, err := ParseSignalIDType(str)
	if err != nil {
		return err
	}

	*s = p
	return nil
}
