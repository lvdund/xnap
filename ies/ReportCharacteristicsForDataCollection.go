package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var reportCharacteristicsForDataCollectionConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(32)),
	Max:        common.Ptr(int64(32)),
}

type ReportCharacteristicsForDataCollection struct {
	Value per.BitString
}

func (ie *ReportCharacteristicsForDataCollection) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, reportCharacteristicsForDataCollectionConstraints)
}

func (ie *ReportCharacteristicsForDataCollection) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(reportCharacteristicsForDataCollectionConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
