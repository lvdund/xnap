package ies

import (
	"github.com/lvdund/asn1go/per"
)

var lastVisitedPSCellListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "lastVisitedPSCellInformation"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type LastVisitedPSCellListItem struct {
	LastVisitedPSCellInformation LastVisitedPSCellInformation
	IEExtensions                 []byte
}

func (ie *LastVisitedPSCellListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lastVisitedPSCellListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.LastVisitedPSCellInformation.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *LastVisitedPSCellListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lastVisitedPSCellListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.LastVisitedPSCellInformation.Decode(d); err != nil {
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
