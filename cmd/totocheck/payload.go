package main

import (
	"encoding/json"
)

type Payload struct {
	Numbers        string `json:"numbers"`
	DrawNo         string `json:"drawNumber"`
	HalfBet        string `json:"isHalfBet"`
	NumberOfParts  string `json:"totalNumberOfParts"`
	PartsPurchased string `json:"partsPurchased"`
}

func NewPayload(numbers, drawno string) *Payload {
	return &Payload{
		Numbers:        numbers,
		DrawNo:         drawno,
		HalfBet:        "false",
		NumberOfParts:  "1",
		PartsPurchased: "1",
	}
}

func (p *Payload) ToJSON() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}