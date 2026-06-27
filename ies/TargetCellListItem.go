package ies

import (
	"github.com/lvdund/asn1go/per"
)

var targetCellListItemConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "target-cell"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type TargetCellListItem struct {
	TargetCell   TargetCGI
	IEExtensions []byte
}

func (ie *TargetCellListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(targetCellListItemConstraints)
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.TargetCell.Encode(e); err != nil {
		return err
	}
	return nil
}

func (ie *TargetCellListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(targetCellListItemConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TargetCell.Decode(d); err != nil {
		return err
	}
	return nil
}
