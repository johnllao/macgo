package message

import (
	"bytes"
	"encoding/gob"
)

type Operation struct {
	Name string
	Parameter []byte
}

func ToOperation(d []byte) (*Operation, error) {
	var err error

	var r = bytes.NewReader(d)
	var dec = gob.NewDecoder(r)
	var o = new(Operation)
	err = dec.Decode(o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *Operation) ToBytes() ([]byte, error) {
	var w  = bytes.NewBuffer(nil)
	var enc = gob.NewEncoder(w)

	var err = enc.Encode(o)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// SUM

type Sum struct {
	A int
	B int
}

func ToSum(d []byte) (*Sum, error) {
	var err error

	var r = bytes.NewReader(d)
	var dec = gob.NewDecoder(r)
	var o = new(Sum)
	err = dec.Decode(o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *Sum) ToBytes() ([]byte, error) {
	var w  = bytes.NewBuffer(nil)
	var enc = gob.NewEncoder(w)

	var err = enc.Encode(o)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func NewSumOperation(a, b int) (*Operation, error) {

	var err error

	var o = &Sum {
		A: a,
		B: b,
	}

	var d []byte
	d, err = o.ToBytes()
	if err != nil {
		return nil, err
	}

	return &Operation{
		Name: "Sum",
		Parameter: d,
	}, nil
}

// DIFF

type Diff struct {
	A int
	B int
}

func ToDiff(d []byte) (*Diff, error) {
	var err error

	var r = bytes.NewReader(d)
	var dec = gob.NewDecoder(r)
	var o = new(Diff)
	err = dec.Decode(o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *Diff) ToBytes() ([]byte, error) {
	var w  = bytes.NewBuffer(nil)
	var enc = gob.NewEncoder(w)

	var err = enc.Encode(o)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func NewDiffOperation(a, b int) (*Operation, error) {
	var err error

	var o = &Diff {
		A: a,
		B: b,
	}

	var d []byte
	d, err = o.ToBytes()
	if err != nil {
		return nil, err
	}

	return &Operation{
		Name: "Diff",
		Parameter: d,
	}, nil
}