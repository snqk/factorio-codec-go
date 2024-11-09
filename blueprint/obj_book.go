package blueprint

import (
	"snqk.dev/factorio/codec/version"
)

type Book struct {
	Item        string          `json:"item"`
	Label       string          `json:"label"`
	LabelColor  *Color          `json:"label_color,omitempty"`
	Blueprints  []BookEntry     `json:"blueprints"`
	ActiveIndex int             `json:"active_index"`
	Icons       []Icon          `json:"icons"`
	Description *string         `json:"description,omitempty"`
	Version     version.Version `json:"version"`
}

func (b Book) isBlueprint() {}
