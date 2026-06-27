package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SensorNameUncompensatedBarometricConfigTrue int64 = 0
)

var sensorNameUncompensatedBarometricConfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SensorNameUncompensatedBarometricConfig struct {
	Value int64
}

func (ie *SensorNameUncompensatedBarometricConfig) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sensorNameUncompensatedBarometricConfigConstraints)
}

func (ie *SensorNameUncompensatedBarometricConfig) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sensorNameUncompensatedBarometricConfigConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SensorNameUeSpeedConfigTrue int64 = 0
)

var sensorNameUeSpeedConfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SensorNameUeSpeedConfig struct {
	Value int64
}

func (ie *SensorNameUeSpeedConfig) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sensorNameUeSpeedConfigConstraints)
}

func (ie *SensorNameUeSpeedConfig) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sensorNameUeSpeedConfigConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	SensorNameUeOrientationConfigTrue int64 = 0
)

var sensorNameUeOrientationConfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SensorNameUeOrientationConfig struct {
	Value int64
}

func (ie *SensorNameUeOrientationConfig) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sensorNameUeOrientationConfigConstraints)
}

func (ie *SensorNameUeOrientationConfig) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sensorNameUeOrientationConfigConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var sensorNameConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uncompensatedBarometricConfig", Optional: true},
		{Name: "ueSpeedConfig", Optional: true},
		{Name: "ueOrientationConfig", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SensorName struct {
	UncompensatedBarometricConfig *SensorNameUncompensatedBarometricConfig
	UeSpeedConfig                 *SensorNameUeSpeedConfig
	UeOrientationConfig           *SensorNameUeOrientationConfig
	IEExtensions                  []byte
}

func (ie *SensorName) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sensorNameConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.UncompensatedBarometricConfig != nil, ie.UeSpeedConfig != nil, ie.UeOrientationConfig != nil, false}); err != nil {
		return err
	}
	if ie.UncompensatedBarometricConfig != nil {
		if err := ie.UncompensatedBarometricConfig.Encode(e); err != nil {
			return err
		}
	}
	if ie.UeSpeedConfig != nil {
		if err := ie.UeSpeedConfig.Encode(e); err != nil {
			return err
		}
	}
	if ie.UeOrientationConfig != nil {
		if err := ie.UeOrientationConfig.Encode(e); err != nil {
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

func (ie *SensorName) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sensorNameConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.UncompensatedBarometricConfig = new(SensorNameUncompensatedBarometricConfig)
		if err := ie.UncompensatedBarometricConfig.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.UeSpeedConfig = new(SensorNameUeSpeedConfig)
		if err := ie.UeSpeedConfig.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.UeOrientationConfig = new(SensorNameUeOrientationConfig)
		if err := ie.UeOrientationConfig.Decode(d); err != nil {
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
