package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mDTConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mDT-Configuration-NR", Optional: true},
		{Name: "mDT-Configuration-EUTRA", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MDTConfiguration struct {
	MDTConfigurationNR    *MDTConfigurationNR
	MDTConfigurationEUTRA *MDTConfigurationEUTRA
	IEExtensions          []byte
}

func (ie *MDTConfiguration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mDTConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MDTConfigurationNR != nil, ie.MDTConfigurationEUTRA != nil, false}); err != nil {
		return err
	}
	if ie.MDTConfigurationNR != nil {
		if err := ie.MDTConfigurationNR.Encode(e); err != nil {
			return err
		}
	}
	if ie.MDTConfigurationEUTRA != nil {
		if err := ie.MDTConfigurationEUTRA.Encode(e); err != nil {
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

func (ie *MDTConfiguration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mDTConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.MDTConfigurationNR = new(MDTConfigurationNR)
		if err := ie.MDTConfigurationNR.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.MDTConfigurationEUTRA = new(MDTConfigurationEUTRA)
		if err := ie.MDTConfigurationEUTRA.Decode(d); err != nil {
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
