package plan

import "pc-monitoring/models/config"

type Floor struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Width     float64     `json:"width"`
	Height    float64     `json:"height"`
	Grid      float64     `json:"grid"`
	Rooms     []Room      `json:"rooms"`
}

type Room struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	PCs       []config.PCs       `json:"pcs"`

    // TODO: Make both config.Printers and config.APs match the same type
	Printers  []config.Printers  `json:"printers"`
	APs       []config.Printers  `json:"access_points"` // []config.APs
}
