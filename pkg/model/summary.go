package model

type Summary struct {
	TotalLines   int
	LevelCount   map[string]int
	ServiceCount map[string]int
}
