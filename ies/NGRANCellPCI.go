package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NGRANCellPCIChNr              = 0
	NGRANCellPCIChEUtra           = 1
	NGRANCellPCIChChoiceExtension = 2
)

var nGRANCellPCIConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr"},
		{Name: "e-utra"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type NGRANCellPCI struct {
	Choice          int
	Nr              *NRPCI
	EUtra           *EUTRAPCI
	ChoiceExtension []byte
}

func (ie *NGRANCellPCI) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(nGRANCellPCIConstraints)
	switch ie.Choice {
	case 0: // nr
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.Nr.Encode(e); err != nil {
			return err
		}
	case 1: // e-utra
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.EUtra.Encode(e); err != nil {
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

func (ie *NGRANCellPCI) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(nGRANCellPCIConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // nr
		ie.Nr = new(NRPCI)
		if err := ie.Nr.Decode(d); err != nil {
			return err
		}
	case 1: // e-utra
		ie.EUtra = new(EUTRAPCI)
		if err := ie.EUtra.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
