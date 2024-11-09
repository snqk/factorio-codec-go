package blueprint

import (
	"encoding/json"
	"fmt"
	"strings"
)

type InfinityFilterMode uint

const (
	InfinityFilterModeUnknown InfinityFilterMode = iota
	InfinityFilterModeAtLeast
	InfinityFilterModeExactly
	InfinityFilterModeAtMost
)

func (s InfinityFilterMode) String() string {
	return map[InfinityFilterMode]string{
		InfinityFilterModeAtLeast: "at-least",
		InfinityFilterModeExactly: "exactly",
		InfinityFilterModeAtMost:  "at-most",
	}[s]
}

func ParseInfinityFilterMode(s string) (InfinityFilterMode, error) {
	switch strings.ToLower(s) {
	case InfinityFilterModeAtLeast.String():
		return InfinityFilterModeAtLeast, nil
	case InfinityFilterModeExactly.String():
		return InfinityFilterModeExactly, nil
	case InfinityFilterModeAtMost.String():
		return InfinityFilterModeAtMost, nil
	}

	return InfinityFilterModeUnknown, fmt.Errorf("unknown InfinityFilterMode: %q", s)
}

func (s InfinityFilterMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *InfinityFilterMode) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	p, err := ParseInfinityFilterMode(str)
	if err != nil {
		return err
	}

	*s = p
	return nil
}
