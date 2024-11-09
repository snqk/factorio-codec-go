package blueprint

import (
	"encoding/json"
	"fmt"
	"strings"
)

type WaitConditionType uint

const (
	WaitConditionTypeUnknown WaitConditionType = iota
	WaitConditionTypeTime
	WaitConditionTypeInactivity
	WaitConditionTypeFull
	WaitConditionTypeEmpty
	WaitConditionTypeItemCount
	WaitConditionTypeCircuit
	WaitConditionTypeRobotsInactive
	WaitConditionTypeFluidCount
	WaitConditionTypePassengerPresent
	WaitConditionTypePassengerNotPresent
)

func (s WaitConditionType) String() string {
	return map[WaitConditionType]string{
		WaitConditionTypeTime:                "time",
		WaitConditionTypeInactivity:          "inactivity",
		WaitConditionTypeFull:                "full",
		WaitConditionTypeEmpty:               "empty",
		WaitConditionTypeItemCount:           "item_count",
		WaitConditionTypeCircuit:             "circuit",
		WaitConditionTypeRobotsInactive:      "robots_inactive",
		WaitConditionTypeFluidCount:          "fluid_count",
		WaitConditionTypePassengerPresent:    "passenger_present",
		WaitConditionTypePassengerNotPresent: "passenger_not_present",
	}[s]
}

func ParseWaitConditionType(s string) (WaitConditionType, error) {
	switch strings.ToLower(s) {
	case WaitConditionTypeTime.String():
		return WaitConditionTypeTime, nil
	case WaitConditionTypeInactivity.String():
		return WaitConditionTypeInactivity, nil
	case WaitConditionTypeFull.String():
		return WaitConditionTypeFull, nil
	case WaitConditionTypeEmpty.String():
		return WaitConditionTypeEmpty, nil
	case WaitConditionTypeItemCount.String():
		return WaitConditionTypeItemCount, nil
	case WaitConditionTypeCircuit.String():
		return WaitConditionTypeCircuit, nil
	case WaitConditionTypeRobotsInactive.String():
		return WaitConditionTypeRobotsInactive, nil
	case WaitConditionTypeFluidCount.String():
		return WaitConditionTypeFluidCount, nil
	case WaitConditionTypePassengerPresent.String():
		return WaitConditionTypePassengerPresent, nil
	case WaitConditionTypePassengerNotPresent.String():
		return WaitConditionTypePassengerNotPresent, nil
	}

	return WaitConditionTypeUnknown, fmt.Errorf("unknown WaitConditionType: %q", s)
}

func (s WaitConditionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *WaitConditionType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	p, err := ParseWaitConditionType(str)
	if err != nil {
		return err
	}

	*s = p
	return nil
}
