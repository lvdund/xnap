package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PSCellHistoryInformationRetrieveQuery int64 = 0
)

var pSCellHistoryInformationRetrieveConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type PSCellHistoryInformationRetrieve struct {
	Value int64
}

func (ie *PSCellHistoryInformationRetrieve) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pSCellHistoryInformationRetrieveConstraints)
}

func (ie *PSCellHistoryInformationRetrieve) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pSCellHistoryInformationRetrieveConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
