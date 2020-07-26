package main

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

type Prize struct {
	GroupNumber int `json:"GroupNumber"`
	ShareAmount int `json:"ShareAmount"`
}

type Result struct {
	Prizes []Prize `json:"Prizes"`
	WinningNumbers []string `json:"WinningNumbers"`
	AdditionalNumber string `json:""`
}

type Raw struct {
	D string `json:"d"`
}

func ToRaw(r io.Reader) (*Raw, error) {
	var err error
	var raw Raw

	var d = json.NewDecoder(r)
	err = d.Decode(&raw)
	if err != nil {
		return nil, errors.New("Failed to decode JSON reeponse. " + err.Error())
	}
	return &raw, nil
}

func (raw *Raw) Result() (*Result, error) {
	var err error
	var result Result

	var r = strings.NewReader(raw.D)
	var d = json.NewDecoder(r)
	d.Decode(&result)
	if err != nil {
		return nil, errors.New("Failed to decode JSON result. " + err.Error())
	}
	return &result, nil
}

