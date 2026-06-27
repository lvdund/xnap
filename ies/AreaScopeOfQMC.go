package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	AreaScopeOfQMCChCellBased       = 0
	AreaScopeOfQMCChTABased         = 1
	AreaScopeOfQMCChTAIBased        = 2
	AreaScopeOfQMCChPLMNAreaBased   = 3
	AreaScopeOfQMCChChoiceExtension = 4
)

var areaScopeOfQMCConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cellBased"},
		{Name: "tABased"},
		{Name: "tAIBased"},
		{Name: "pLMNAreaBased"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type AreaScopeOfQMC struct {
	Choice          int
	CellBased       *CellBasedQMC
	TABased         *TABasedQMC
	TAIBased        *TAIBasedQMC
	PLMNAreaBased   *PLMNAreaBasedQMC
	ChoiceExtension []byte
}

func (ie *AreaScopeOfQMC) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(areaScopeOfQMCConstraints)
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
	case 3: // pLMNAreaBased
		if err := choice.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		if err := ie.PLMNAreaBased.Encode(e); err != nil {
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

func (ie *AreaScopeOfQMC) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(areaScopeOfQMCConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // cellBased
		ie.CellBased = new(CellBasedQMC)
		if err := ie.CellBased.Decode(d); err != nil {
			return err
		}
	case 1: // tABased
		ie.TABased = new(TABasedQMC)
		if err := ie.TABased.Decode(d); err != nil {
			return err
		}
	case 2: // tAIBased
		ie.TAIBased = new(TAIBasedQMC)
		if err := ie.TAIBased.Decode(d); err != nil {
			return err
		}
	case 3: // pLMNAreaBased
		ie.PLMNAreaBased = new(PLMNAreaBasedQMC)
		if err := ie.PLMNAreaBased.Decode(d); err != nil {
			return err
		}
	case 4: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
