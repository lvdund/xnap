package ies

import (
	"github.com/lvdund/asn1go/per"
)

var expectedUEBehaviourConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "expectedUEActivityBehaviour", Optional: true},
		{Name: "expectedHOInterval", Optional: true},
		{Name: "expectedUEMobility", Optional: true},
		{Name: "expectedUEMovingTrajectory", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ExpectedUEBehaviour struct {
	ExpectedUEActivityBehaviour *ExpectedUEActivityBehaviour
	ExpectedHOInterval          *ExpectedHOInterval
	ExpectedUEMobility          *ExpectedUEMobility
	ExpectedUEMovingTrajectory  *ExpectedUEMovingTrajectory
	IEExtensions                []byte
}

func (ie *ExpectedUEBehaviour) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(expectedUEBehaviourConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ExpectedUEActivityBehaviour != nil, ie.ExpectedHOInterval != nil, ie.ExpectedUEMobility != nil, ie.ExpectedUEMovingTrajectory != nil, false}); err != nil {
		return err
	}
	if ie.ExpectedUEActivityBehaviour != nil {
		if err := ie.ExpectedUEActivityBehaviour.Encode(e); err != nil {
			return err
		}
	}
	if ie.ExpectedHOInterval != nil {
		if err := ie.ExpectedHOInterval.Encode(e); err != nil {
			return err
		}
	}
	if ie.ExpectedUEMobility != nil {
		if err := ie.ExpectedUEMobility.Encode(e); err != nil {
			return err
		}
	}
	if ie.ExpectedUEMovingTrajectory != nil {
		if err := ie.ExpectedUEMovingTrajectory.Encode(e); err != nil {
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

func (ie *ExpectedUEBehaviour) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(expectedUEBehaviourConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.ExpectedUEActivityBehaviour = new(ExpectedUEActivityBehaviour)
		if err := ie.ExpectedUEActivityBehaviour.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.ExpectedHOInterval = new(ExpectedHOInterval)
		if err := ie.ExpectedHOInterval.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.ExpectedUEMobility = new(ExpectedUEMobility)
		if err := ie.ExpectedUEMobility.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.ExpectedUEMovingTrajectory = new(ExpectedUEMovingTrajectory)
		if err := ie.ExpectedUEMovingTrajectory.Decode(d); err != nil {
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
