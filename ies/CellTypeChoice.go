package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CellTypeChoiceChNgRanEUtra      = 0
	CellTypeChoiceChNgRanNr         = 1
	CellTypeChoiceChEUtran          = 2
	CellTypeChoiceChChoiceExtension = 3
)

var cellTypeChoiceConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ng-ran-e-utra"},
		{Name: "ng-ran-nr"},
		{Name: "e-utran"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type CellTypeChoice struct {
	Choice          int
	NgRanEUtra      *EUTRACellIdentity
	NgRanNr         *NRCellIdentity
	EUtran          *EUTRACellIdentity
	ChoiceExtension []byte
}

func (ie *CellTypeChoice) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(cellTypeChoiceConstraints)
	switch ie.Choice {
	case 0: // ng-ran-e-utra
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.NgRanEUtra.Encode(e); err != nil {
			return err
		}
	case 1: // ng-ran-nr
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.NgRanNr.Encode(e); err != nil {
			return err
		}
	case 2: // e-utran
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := ie.EUtran.Encode(e); err != nil {
			return err
		}
	case 3: // choice-extension
		if err := choice.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *CellTypeChoice) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(cellTypeChoiceConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // ng-ran-e-utra
		ie.NgRanEUtra = new(EUTRACellIdentity)
		if err := ie.NgRanEUtra.Decode(d); err != nil {
			return err
		}
	case 1: // ng-ran-nr
		ie.NgRanNr = new(NRCellIdentity)
		if err := ie.NgRanNr.Decode(d); err != nil {
			return err
		}
	case 2: // e-utran
		ie.EUtran = new(EUTRACellIdentity)
		if err := ie.EUtran.Decode(d); err != nil {
			return err
		}
	case 3: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
