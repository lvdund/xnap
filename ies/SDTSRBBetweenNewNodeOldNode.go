package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sDTSRBBetweenNewNodeOldNodeConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rrcContainer"},
		{Name: "srb-ID"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SDTSRBBetweenNewNodeOldNode struct {
	RrcContainer []byte
	SrbID        SRBID
	IEExtensions []byte
}

func (ie *SDTSRBBetweenNewNodeOldNode) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDTSRBBetweenNewNodeOldNodeConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.RrcContainer, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if err := ie.SrbID.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SDTSRBBetweenNewNodeOldNode) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDTSRBBetweenNewNodeOldNodeConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.RrcContainer = val
	}
	if err := ie.SrbID.Decode(d); err != nil {
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
