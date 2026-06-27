package ies

import (
	"github.com/lvdund/asn1go/per"
)

var cHOCPACInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cHO-CPAC-config-indicator", Optional: true},
		{Name: "cHO-target-SN-node-list"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CHOCPACInformation struct {
	CHOCPACConfigIndicator *CHOCPACConfigIndicator
	CHOTargetSNNodeList    CHOTargetSNNodeList
	IEExtensions           []byte
}

func (ie *CHOCPACInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cHOCPACInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.CHOCPACConfigIndicator != nil, false}); err != nil {
		return err
	}
	if ie.CHOCPACConfigIndicator != nil {
		if err := ie.CHOCPACConfigIndicator.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.CHOTargetSNNodeList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CHOCPACInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cHOCPACInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.CHOCPACConfigIndicator = new(CHOCPACConfigIndicator)
		if err := ie.CHOCPACConfigIndicator.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.CHOTargetSNNodeList.Decode(d); err != nil {
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
