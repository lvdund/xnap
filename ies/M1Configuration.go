package ies

import (
	"github.com/lvdund/asn1go/per"
)

var m1ConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "m1reportingTrigger"},
		{Name: "m1thresholdeventA2", Optional: true},
		{Name: "m1periodicReporting", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type M1Configuration struct {
	M1reportingTrigger  M1ReportingTrigger
	M1thresholdeventA2  *M1ThresholdEventA2
	M1periodicReporting *M1PeriodicReporting
	IEExtensions        []byte
}

func (ie *M1Configuration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(m1ConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.M1thresholdeventA2 != nil, ie.M1periodicReporting != nil, false}); err != nil {
		return err
	}
	if err := ie.M1reportingTrigger.Encode(e); err != nil {
		return err
	}
	if ie.M1thresholdeventA2 != nil {
		if err := ie.M1thresholdeventA2.Encode(e); err != nil {
			return err
		}
	}
	if ie.M1periodicReporting != nil {
		if err := ie.M1periodicReporting.Encode(e); err != nil {
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

func (ie *M1Configuration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(m1ConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.M1reportingTrigger.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.M1thresholdeventA2 = new(M1ThresholdEventA2)
		if err := ie.M1thresholdeventA2.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.M1periodicReporting = new(M1PeriodicReporting)
		if err := ie.M1periodicReporting.Decode(d); err != nil {
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
