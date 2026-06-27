package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RANPagingAreaChoiceChCellList        = 0
	RANPagingAreaChoiceChRANAreaIDList   = 1
	RANPagingAreaChoiceChChoiceExtension = 2
)

var rANPagingAreaChoiceConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cell-List"},
		{Name: "rANAreaID-List"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type RANPagingAreaChoice struct {
	Choice          int
	CellList        *NGRANCellIdentityListinRANPagingArea
	RANAreaIDList   *RANAreaIDList
	ChoiceExtension []byte
}

func (ie *RANPagingAreaChoice) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(rANPagingAreaChoiceConstraints)
	switch ie.Choice {
	case 0: // cell-List
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.CellList.Encode(e); err != nil {
			return err
		}
	case 1: // rANAreaID-List
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.RANAreaIDList.Encode(e); err != nil {
			return err
		}
	case 2: // choice-extension
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *RANPagingAreaChoice) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(rANPagingAreaChoiceConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // cell-List
		ie.CellList = new(NGRANCellIdentityListinRANPagingArea)
		if err := ie.CellList.Decode(d); err != nil {
			return err
		}
	case 1: // rANAreaID-List
		ie.RANAreaIDList = new(RANAreaIDList)
		if err := ie.RANAreaIDList.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
