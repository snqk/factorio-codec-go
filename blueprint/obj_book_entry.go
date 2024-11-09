package blueprint

import (
	"encoding/json"
	"fmt"
)

type IsBlueprint interface {
	isBlueprint()
}

type BookEntry struct {
	Index *int
	Entry IsBlueprint
}

func (be BookEntry) MarshalJSON() ([]byte, error) {
	type idx struct {
		Index *int `json:"index,omitempty"`
	}
	var entry idx
	entry.Index = be.Index

	switch v := be.Entry.(type) {
	case Blueprint:
		return json.Marshal(struct {
			idx
			Blueprint Blueprint `json:"blueprint"`
		}{entry, v})
	case Book:
		return json.Marshal(struct {
			idx
			BlueprintBook Book `json:"blueprint_book"`
		}{entry, v})
	default:
		return nil, fmt.Errorf("unsupported type for BookEntry: %T", v)
	}
}

func (be *BookEntry) UnmarshalJSON(in []byte) error {
	var m map[string]any
	if err := json.Unmarshal(in, &m); err != nil {
		return err
	}

	nbe := BookEntry{}
	if v, ok := m["index"]; ok {
		idx := int(v.(float64))
		nbe.Index = &idx
	}

	if _, ok := m["blueprint"]; ok {
		var s struct {
			Blueprint Blueprint `json:"blueprint"`
		}
		if err := json.Unmarshal(in, &s); err != nil {
			return err
		}
		nbe.Entry = s.Blueprint
	} else if _, ok := m["blueprint_book"]; ok {
		var s struct {
			Book Book `json:"blueprint_book"`
		}
		if err := json.Unmarshal(in, &s); err != nil {
			return err
		}
		nbe.Entry = s.Book
	} else {
		return fmt.Errorf("missing 'blueprint' or 'blueprint_book' field")
	}

	*be = nbe
	return nil
}
