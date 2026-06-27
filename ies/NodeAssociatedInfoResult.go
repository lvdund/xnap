package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nodeAssociatedInfoResultConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "energyCost", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NodeAssociatedInfoResult struct {
	EnergyCost   *EnergyCost
	IEExtensions []byte
}

func (ie *NodeAssociatedInfoResult) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nodeAssociatedInfoResultConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.EnergyCost != nil, false}); err != nil {
		return err
	}
	if ie.EnergyCost != nil {
		if err := ie.EnergyCost.Encode(e); err != nil {
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

func (ie *NodeAssociatedInfoResult) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nodeAssociatedInfoResultConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.EnergyCost = new(EnergyCost)
		if err := ie.EnergyCost.Decode(d); err != nil {
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
