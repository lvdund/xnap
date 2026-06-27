package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sDTPartialUEContextInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dRBsToBeSetup", Optional: true},
		{Name: "sRBsToBeSetup"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SDTPartialUEContextInfo struct {
	DRBsToBeSetup *SDTDRBsToBeSetupList
	SRBsToBeSetup SDTSRBsToBeSetupList
	IEExtensions  []byte
}

func (ie *SDTPartialUEContextInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDTPartialUEContextInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DRBsToBeSetup != nil, false}); err != nil {
		return err
	}
	if ie.DRBsToBeSetup != nil {
		if err := ie.DRBsToBeSetup.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.SRBsToBeSetup.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SDTPartialUEContextInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDTPartialUEContextInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.DRBsToBeSetup = new(SDTDRBsToBeSetupList)
		if err := ie.DRBsToBeSetup.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.SRBsToBeSetup.Decode(d); err != nil {
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
