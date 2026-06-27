package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	ServedCellsToActivateChNrCells         = 0
	ServedCellsToActivateChEUtraCells      = 1
	ServedCellsToActivateChChoiceExtension = 2
)

var servedCellsToActivateConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr-cells"},
		{Name: "e-utra-cells"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type ServedCellsToActivate struct {
	Choice          int
	NrCells         []*NRCGI
	EUtraCells      []*EUTRACGI
	ChoiceExtension []byte
}

func (ie *ServedCellsToActivate) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(servedCellsToActivateConstraints)
	switch ie.Choice {
	case 0: // nr-cells
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.NrCells))); err != nil {
			return err
		}
		for _, item := range ie.NrCells {
			if err := item.Encode(e); err != nil {
				return err
			}
		}
	case 1: // e-utra-cells
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.EUtraCells))); err != nil {
			return err
		}
		for _, item := range ie.EUtraCells {
			if err := item.Encode(e); err != nil {
				return err
			}
		}
	case 2: // choice-extension
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *ServedCellsToActivate) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(servedCellsToActivateConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // nr-cells
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.NrCells = make([]*NRCGI, n)
		for i := range ie.NrCells {
			ie.NrCells[i] = new(NRCGI)
			if err := ie.NrCells[i].Decode(d); err != nil {
				return err
			}
		}
	case 1: // e-utra-cells
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofCellsinNG_RANnode)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.EUtraCells = make([]*EUTRACGI, n)
		for i := range ie.EUtraCells {
			ie.EUtraCells[i] = new(EUTRACGI)
			if err := ie.EUtraCells[i].Decode(d); err != nil {
				return err
			}
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
