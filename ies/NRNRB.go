package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NRNRBNrb11  int64 = 0
	NRNRBNrb18  int64 = 1
	NRNRBNrb24  int64 = 2
	NRNRBNrb25  int64 = 3
	NRNRBNrb31  int64 = 4
	NRNRBNrb32  int64 = 5
	NRNRBNrb38  int64 = 6
	NRNRBNrb51  int64 = 7
	NRNRBNrb52  int64 = 8
	NRNRBNrb65  int64 = 9
	NRNRBNrb66  int64 = 10
	NRNRBNrb78  int64 = 11
	NRNRBNrb79  int64 = 12
	NRNRBNrb93  int64 = 13
	NRNRBNrb106 int64 = 14
	NRNRBNrb107 int64 = 15
	NRNRBNrb121 int64 = 16
	NRNRBNrb132 int64 = 17
	NRNRBNrb133 int64 = 18
	NRNRBNrb135 int64 = 19
	NRNRBNrb160 int64 = 20
	NRNRBNrb162 int64 = 21
	NRNRBNrb189 int64 = 22
	NRNRBNrb216 int64 = 23
	NRNRBNrb217 int64 = 24
	NRNRBNrb245 int64 = 25
	NRNRBNrb264 int64 = 26
	NRNRBNrb270 int64 = 27
	NRNRBNrb273 int64 = 28
	NRNRBNrb33  int64 = 29
	NRNRBNrb62  int64 = 30
	NRNRBNrb124 int64 = 31
	NRNRBNrb148 int64 = 32
	NRNRBNrb248 int64 = 33
	NRNRBNrb44  int64 = 34
	NRNRBNrb58  int64 = 35
	NRNRBNrb92  int64 = 36
	NRNRBNrb119 int64 = 37
	NRNRBNrb188 int64 = 38
	NRNRBNrb242 int64 = 39
	NRNRBNrb15  int64 = 40
)

var nRNRBConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28},
	ExtValues:  []int64{29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
}

type NRNRB struct {
	Value int64
}

func (ie *NRNRB) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, nRNRBConstraints)
}

func (ie *NRNRB) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(nRNRBConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
