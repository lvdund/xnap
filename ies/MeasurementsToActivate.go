package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var measurementsToActivateConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(8)),
	Max:        common.Ptr(int64(8)),
}

type MeasurementsToActivate struct {
	Value per.BitString
}

func (ie *MeasurementsToActivate) Encode(e *per.Encoder) error {
	return e.EncodeBitString(ie.Value, measurementsToActivateConstraints)
}

func (ie *MeasurementsToActivate) Decode(d *per.Decoder) error {
	val, err := d.DecodeBitString(measurementsToActivateConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
