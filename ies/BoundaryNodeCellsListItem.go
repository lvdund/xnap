package ies

import (
	"github.com/lvdund/asn1go/per"
)

var boundaryNodeCellsListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "boundaryNodeCellInformation"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type BoundaryNodeCellsListItem struct {
	BoundaryNodeCellInformation IABCellInformation
	IEExtensions                []byte
}

func (ie *BoundaryNodeCellsListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(boundaryNodeCellsListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.BoundaryNodeCellInformation.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *BoundaryNodeCellsListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(boundaryNodeCellsListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.BoundaryNodeCellInformation.Decode(d); err != nil {
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
