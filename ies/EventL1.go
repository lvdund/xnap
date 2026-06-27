package ies

import (
	"github.com/lvdund/asn1go/per"
)

var eventL1Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "l1Threshold"},
		{Name: "hysteresis"},
		{Name: "timeToTrigger"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type EventL1 struct {
	L1Threshold   MeasurementThresholdL1LoggedMDT
	Hysteresis    Hysteresis
	TimeToTrigger TimeToTrigger
	IEExtensions  []byte
}

func (ie *EventL1) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eventL1Constraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.L1Threshold.Encode(e); err != nil {
		return err
	}
	if err := ie.Hysteresis.Encode(e); err != nil {
		return err
	}
	if err := ie.TimeToTrigger.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *EventL1) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eventL1Constraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.L1Threshold.Decode(d); err != nil {
		return err
	}
	if err := ie.Hysteresis.Decode(d); err != nil {
		return err
	}
	if err := ie.TimeToTrigger.Decode(d); err != nil {
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
