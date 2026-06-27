package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mDTConfigurationEUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mdt-Activation"},
		{Name: "areaScopeOfMDT-EUTRA", Optional: true},
		{Name: "mDTMode-EUTRA"},
		{Name: "signallingBasedMDTPLMNList"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MDTConfigurationEUTRA struct {
	MdtActivation              MDTActivation
	AreaScopeOfMDTEUTRA        *AreaScopeOfMDTEUTRA
	MDTModeEUTRA               MDTModeEUTRA
	SignallingBasedMDTPLMNList MDTPLMNList
	IEExtensions               []byte
}

func (ie *MDTConfigurationEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mDTConfigurationEUTRAConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.AreaScopeOfMDTEUTRA != nil, false}); err != nil {
		return err
	}
	if err := ie.MdtActivation.Encode(e); err != nil {
		return err
	}
	if ie.AreaScopeOfMDTEUTRA != nil {
		if err := ie.AreaScopeOfMDTEUTRA.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.MDTModeEUTRA.Encode(e); err != nil {
		return err
	}
	if err := ie.SignallingBasedMDTPLMNList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *MDTConfigurationEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mDTConfigurationEUTRAConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.MdtActivation.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.AreaScopeOfMDTEUTRA = new(AreaScopeOfMDTEUTRA)
		if err := ie.AreaScopeOfMDTEUTRA.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.MDTModeEUTRA.Decode(d); err != nil {
		return err
	}
	if err := ie.SignallingBasedMDTPLMNList.Decode(d); err != nil {
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
