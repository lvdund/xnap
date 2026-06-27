package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SymbolAllocationInSlotChAllDL           = 0
	SymbolAllocationInSlotChAllUL           = 1
	SymbolAllocationInSlotChBothDLandUL     = 2
	SymbolAllocationInSlotChChoiceExtension = 3
)

var symbolAllocationInSlotConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "allDL"},
		{Name: "allUL"},
		{Name: "bothDLandUL"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type SymbolAllocationInSlot struct {
	Choice          int
	AllDL           *SymbolAllocationInSlotAllDL
	AllUL           *SymbolAllocationInSlotAllUL
	BothDLandUL     *SymbolAllocationInSlotBothDLandUL
	ChoiceExtension []byte
}

func (ie *SymbolAllocationInSlot) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(symbolAllocationInSlotConstraints)
	switch ie.Choice {
	case 0: // allDL
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.AllDL.Encode(e); err != nil {
			return err
		}
	case 1: // allUL
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.AllUL.Encode(e); err != nil {
			return err
		}
	case 2: // bothDLandUL
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := ie.BothDLandUL.Encode(e); err != nil {
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

func (ie *SymbolAllocationInSlot) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(symbolAllocationInSlotConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // allDL
		ie.AllDL = new(SymbolAllocationInSlotAllDL)
		if err := ie.AllDL.Decode(d); err != nil {
			return err
		}
	case 1: // allUL
		ie.AllUL = new(SymbolAllocationInSlotAllUL)
		if err := ie.AllUL.Decode(d); err != nil {
			return err
		}
	case 2: // bothDLandUL
		ie.BothDLandUL = new(SymbolAllocationInSlotBothDLandUL)
		if err := ie.BothDLandUL.Decode(d); err != nil {
			return err
		}
	case 3: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
