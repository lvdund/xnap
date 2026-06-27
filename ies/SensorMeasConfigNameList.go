package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var sensorMeasConfigNameListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofSensorName)),
}

type SensorMeasConfigNameList struct {
	Value []*SensorName
}

func (ie *SensorMeasConfigNameList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sensorMeasConfigNameListConstraints)
	if err := seqOf.EncodeLength(int64(len(ie.Value))); err != nil {
		return err
	}
	for _, item := range ie.Value {
		if err := item.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SensorMeasConfigNameList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sensorMeasConfigNameListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*SensorName, n)
	for i := range ie.Value {
		ie.Value[i] = new(SensorName)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
