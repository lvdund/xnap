package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var reportCharacteristicsConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(32)),
	Max:        common.Ptr(int64(32)),
}

type ReportCharacteristics struct {
	Value per.BitString
}

func (ie *ReportCharacteristics) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, reportCharacteristicsConstraints)
}

func (ie *ReportCharacteristics) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(reportCharacteristicsConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
