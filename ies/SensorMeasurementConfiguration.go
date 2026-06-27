package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sensorMeasurementConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sensorMeasConfig"},
		{Name: "sensorMeasConfigNameList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SensorMeasurementConfiguration struct {
	SensorMeasConfig         SensorMeasConfig
	SensorMeasConfigNameList *SensorMeasConfigNameList
	IEExtensions             []byte
}

func (ie *SensorMeasurementConfiguration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sensorMeasurementConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SensorMeasConfigNameList != nil, false}); err != nil {
		return err
	}
	if err := ie.SensorMeasConfig.Encode(e); err != nil {
		return err
	}
	if ie.SensorMeasConfigNameList != nil {
		if err := ie.SensorMeasConfigNameList.Encode(e); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SensorMeasurementConfiguration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sensorMeasurementConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SensorMeasConfig.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.SensorMeasConfigNameList = new(SensorMeasConfigNameList)
		if err := ie.SensorMeasConfigNameList.Decode(d); err != nil {
			return err
		}
	}
	extBytes, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	if len(extBytes) > 0 {
		ie.IEExtensions = extBytes[0]
	}
	return nil
}
