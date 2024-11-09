package blueprint

import (
	"encoding/json"
	"fmt"
	"strings"
)

type WaitConditionCompareType uint

const (
	WaitConditionCompareTypeUnknown WaitConditionCompareType = iota
	WaitConditionCompareTypeAnd
	WaitConditionCompareTypeOr
)

func (s WaitConditionCompareType) String() string {
	return map[WaitConditionCompareType]string{
		WaitConditionCompareTypeAnd: "and",
		WaitConditionCompareTypeOr:  "or",
	}[s]
}

func ParseWaitConditionCompareType(s string) (WaitConditionCompareType, error) {
	switch strings.ToLower(s) {
	case WaitConditionCompareTypeAnd.String():
		return WaitConditionCompareTypeAnd, nil
	case WaitConditionCompareTypeOr.String():
		return WaitConditionCompareTypeOr, nil
	}

	return WaitConditionCompareTypeUnknown, fmt.Errorf("unknown WaitConditionCompareType: %q", s)
}

func (s WaitConditionCompareType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *WaitConditionCompareType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	p, err := ParseWaitConditionCompareType(str)
	if err != nil {
		return err
	}

	*s = p
	return nil
}
