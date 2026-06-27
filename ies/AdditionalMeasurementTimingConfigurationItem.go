package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var additionalMeasurementTimingConfigurationItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "additionalMeasurementTimingConfigurationIndex"},
		{Name: "csi-RS-MTC-Configuration-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type AdditionalMeasurementTimingConfigurationItem struct {
	AdditionalMeasurementTimingConfigurationIndex int64
	CsiRSMTCConfigurationList                     CSIRSMTCConfigurationList
	IEExtensions                                  []byte
}

func (ie *AdditionalMeasurementTimingConfigurationItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(additionalMeasurementTimingConfigurationItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.AdditionalMeasurementTimingConfigurationIndex, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(16)),
	}); err != nil {
		return err
	}
	if err := ie.CsiRSMTCConfigurationList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *AdditionalMeasurementTimingConfigurationItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(additionalMeasurementTimingConfigurationItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(16)),
		})
		if err != nil {
			return err
		}
		ie.AdditionalMeasurementTimingConfigurationIndex = val
	}
	if err := ie.CsiRSMTCConfigurationList.Decode(d); err != nil {
		return err
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
