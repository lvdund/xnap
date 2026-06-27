package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	AreaScopeOfMDTEUTRAChCellBased = 0
	AreaScopeOfMDTEUTRAChTABased   = 1
	AreaScopeOfMDTEUTRAChTAIBased  = 2
	AreaScopeOfMDTEUTRAExtension   = -1
)

var areaScopeOfMDTEUTRAConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellBased"},
		{Name: "tABased"},
		{Name: "tAIBased"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "choice-extension"},
	},
}

type AreaScopeOfMDTEUTRA struct {
	Choice          int
	CellBased       *CellBasedMDTEUTRA
	TABased         *TABasedMDT
	TAIBased        *TAIBasedMDT
	ChoiceExtension []byte
}

func (ie *AreaScopeOfMDTEUTRA) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(areaScopeOfMDTEUTRAConstraints)
	switch ie.Choice {
	case 0: // cellBased
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.CellBased.Encode(e); err != nil {
			return err
		}
	case 1: // tABased
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.TABased.Encode(e); err != nil {
			return err
		}
	case 2: // tAIBased
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := ie.TAIBased.Encode(e); err != nil {
			return err
		}
	default: // extension
		if err := choice.EncodeChoice(0, true, ie.ChoiceExtension); err != nil {
			return err
		}
	}
	return nil
}

func (ie *AreaScopeOfMDTEUTRA) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(areaScopeOfMDTEUTRAConstraints)
	idx, isExt, extBytes, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	if isExt {
		ie.Choice = AreaScopeOfMDTEUTRAExtension
		ie.ChoiceExtension = extBytes
		return nil
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // cellBased
		ie.CellBased = new(CellBasedMDTEUTRA)
		if err := ie.CellBased.Decode(d); err != nil {
			return err
		}
	case 1: // tABased
		ie.TABased = new(TABasedMDT)
		if err := ie.TABased.Decode(d); err != nil {
			return err
		}
	case 2: // tAIBased
		ie.TAIBased = new(TAIBasedMDT)
		if err := ie.TAIBased.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
