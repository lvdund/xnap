package ies

import (
	"github.com/lvdund/asn1go/per"
)

var expectedUEActivityBehaviourConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "expectedActivityPeriod", Optional: true},
		{Name: "expectedIdlePeriod", Optional: true},
		{Name: "sourceOfUEActivityBehaviourInformation", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ExpectedUEActivityBehaviour struct {
	ExpectedActivityPeriod                 *ExpectedActivityPeriod
	ExpectedIdlePeriod                     *ExpectedIdlePeriod
	SourceOfUEActivityBehaviourInformation *SourceOfUEActivityBehaviourInformation
	IEExtensions                           []byte
}

func (ie *ExpectedUEActivityBehaviour) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(expectedUEActivityBehaviourConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ExpectedActivityPeriod != nil, ie.ExpectedIdlePeriod != nil, ie.SourceOfUEActivityBehaviourInformation != nil, false}); err != nil {
		return err
	}
	if ie.ExpectedActivityPeriod != nil {
		if err := ie.ExpectedActivityPeriod.Encode(e); err != nil {
			return err
		}
	}
	if ie.ExpectedIdlePeriod != nil {
		if err := ie.ExpectedIdlePeriod.Encode(e); err != nil {
			return err
		}
	}
	if ie.SourceOfUEActivityBehaviourInformation != nil {
		if err := ie.SourceOfUEActivityBehaviourInformation.Encode(e); err != nil {
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

func (ie *ExpectedUEActivityBehaviour) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(expectedUEActivityBehaviourConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.ExpectedActivityPeriod = new(ExpectedActivityPeriod)
		if err := ie.ExpectedActivityPeriod.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.ExpectedIdlePeriod = new(ExpectedIdlePeriod)
		if err := ie.ExpectedIdlePeriod.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.SourceOfUEActivityBehaviourInformation = new(SourceOfUEActivityBehaviourInformation)
		if err := ie.SourceOfUEActivityBehaviourInformation.Decode(d); err != nil {
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
