package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	LastVisitedCellItemChNGRANCell       = 0
	LastVisitedCellItemChEUTRANCell      = 1
	LastVisitedCellItemChUTRANCell       = 2
	LastVisitedCellItemChGERANCell       = 3
	LastVisitedCellItemChChoiceExtension = 4
)

var lastVisitedCellItemConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nG-RAN-Cell"},
		{Name: "e-UTRAN-Cell"},
		{Name: "uTRAN-Cell"},
		{Name: "gERAN-Cell"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type LastVisitedCellItem struct {
	Choice          int
	NGRANCell       *LastVisitedNGRANCellInformation
	EUTRANCell      *LastVisitedEUTRANCellInformation
	UTRANCell       *LastVisitedUTRANCellInformation
	GERANCell       *LastVisitedGERANCellInformation
	ChoiceExtension []byte
}

func (ie *LastVisitedCellItem) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(lastVisitedCellItemConstraints)
	switch ie.Choice {
	case 0: // nG-RAN-Cell
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.NGRANCell.Encode(e); err != nil {
			return err
		}
	case 1: // e-UTRAN-Cell
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.EUTRANCell.Encode(e); err != nil {
			return err
		}
	case 2: // uTRAN-Cell
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := ie.UTRANCell.Encode(e); err != nil {
			return err
		}
	case 3: // gERAN-Cell
		if err := choice.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		if err := ie.GERANCell.Encode(e); err != nil {
			return err
		}
	case 4: // choice-extension
		if err := choice.EncodeChoice(4, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *LastVisitedCellItem) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(lastVisitedCellItemConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // nG-RAN-Cell
		ie.NGRANCell = new(LastVisitedNGRANCellInformation)
		if err := ie.NGRANCell.Decode(d); err != nil {
			return err
		}
	case 1: // e-UTRAN-Cell
		ie.EUTRANCell = new(LastVisitedEUTRANCellInformation)
		if err := ie.EUTRANCell.Decode(d); err != nil {
			return err
		}
	case 2: // uTRAN-Cell
		ie.UTRANCell = new(LastVisitedUTRANCellInformation)
		if err := ie.UTRANCell.Decode(d); err != nil {
			return err
		}
	case 3: // gERAN-Cell
		ie.GERANCell = new(LastVisitedGERANCellInformation)
		if err := ie.GERANCell.Decode(d); err != nil {
			return err
		}
	case 4: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
