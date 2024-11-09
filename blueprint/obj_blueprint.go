package blueprint

import (
	"snqk.dev/factorio/codec/version"
)

type Blueprint struct {
	Item                   string          `json:"item"`
	Label                  string          `json:"label"`
	LabelColor             *Color          `json:"label_color,omitempty"`
	Entities               []Entity        `json:"entities,omitempty"`
	Tiles                  []Tile          `json:"tiles,omitempty"`
	Icons                  []Icon          `json:"icons,omitempty"`
	Schedules              []Schedule      `json:"schedules,omitempty"`
	Description            *string         `json:"description,omitempty"`
	SnapToGrid             *Position       `json:"snap-to-grid,omitempty"`
	AbsoluteSnapping       *bool           `json:"absolute-snapping,omitempty"`
	PositionRelativeToGrid *Position       `json:"position-relative-to-grid,omitempty"`
	Version                version.Version `json:"version"`
}

func (b Blueprint) isBlueprint() {}
