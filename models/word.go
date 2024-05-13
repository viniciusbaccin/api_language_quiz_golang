package models

type Word struct {
	Word string `json:"word"`
	Translation string `json:"translation"`
	Options []string `json:"options"`
}
