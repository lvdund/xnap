package ies

import (
	"github.com/lvdund/asn1go/per"
)

var mDTConfigurationNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mdt-Activation"},
		{Name: "areaScopeOfMDT-NR", Optional: true},
		{Name: "mDTMode-NR"},
		{Name: "signallingBasedMDTPLMNList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MDTConfigurationNR struct {
	MdtActivation              MDTActivation
	AreaScopeOfMDTNR           *AreaScopeOfMDTNR
	MDTModeNR                  MDTModeNR
	SignallingBasedMDTPLMNList *MDTPLMNList
	IEExtensions               []byte
}

func (ie *MDTConfigurationNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mDTConfigurationNRConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.AreaScopeOfMDTNR != nil, ie.SignallingBasedMDTPLMNList != nil, false}); err != nil {
		return err
	}
	if err := ie.MdtActivation.Encode(e); err != nil {
		return err
	}
	if ie.AreaScopeOfMDTNR != nil {
		if err := ie.AreaScopeOfMDTNR.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.MDTModeNR.Encode(e); err != nil {
		return err
	}
	if ie.SignallingBasedMDTPLMNList != nil {
		if err := ie.SignallingBasedMDTPLMNList.Encode(e); err != nil {
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

func (ie *MDTConfigurationNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mDTConfigurationNRConstraints)
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
		ie.AreaScopeOfMDTNR = new(AreaScopeOfMDTNR)
		if err := ie.AreaScopeOfMDTNR.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.MDTModeNR.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(3) {
		ie.SignallingBasedMDTPLMNList = new(MDTPLMNList)
		if err := ie.SignallingBasedMDTPLMNList.Decode(d); err != nil {
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
