package ies

import (
	"github.com/lvdund/asn1go/per"
)

var neighbourNGRANNodeItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "globalNG-RANNodeID"},
		{Name: "local-NG-RAN-Node-Identifier"},
		{Name: "ie-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NeighbourNGRANNodeItem struct {
	GlobalNGRANNodeID        GlobalNGRANNodeID
	LocalNGRANNodeIdentifier LocalNGRANNodeIdentifier
	IEExtensions             []byte
}

func (ie *NeighbourNGRANNodeItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(neighbourNGRANNodeItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.GlobalNGRANNodeID.Encode(e); err != nil {
		return err
	}
	if err := ie.LocalNGRANNodeIdentifier.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NeighbourNGRANNodeItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(neighbourNGRANNodeItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.GlobalNGRANNodeID.Decode(d); err != nil {
		return err
	}
	if err := ie.LocalNGRANNodeIdentifier.Decode(d); err != nil {
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
