package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pEIPSassistanceInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cNsubgroupID"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PEIPSassistanceInformation struct {
	CNsubgroupID CNsubgroupID
	IEExtensions []byte
}

func (ie *PEIPSassistanceInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pEIPSassistanceInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.CNsubgroupID.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PEIPSassistanceInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pEIPSassistanceInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.CNsubgroupID.Decode(d); err != nil {
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
