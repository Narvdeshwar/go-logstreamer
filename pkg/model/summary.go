package model

type Summary struct {
	TotalLines   int            `json:"total_lines"`
	LevelCount   map[string]int `json:"levels"`
	ServiceCount map[string]int `json:"services"`
}
