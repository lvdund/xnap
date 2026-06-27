package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	LocalNGRANNodeIdentifierChFullIRNTIProfileList  = 0
	LocalNGRANNodeIdentifierChShortIRNTIProfileList = 1
	LocalNGRANNodeIdentifierChChoiceExtension       = 2
)

var localNGRANNodeIdentifierConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "full-I-RNTI-Profile-List"},
		{Name: "short-I-RNTI-Profile-List"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type LocalNGRANNodeIdentifier struct {
	Choice                int
	FullIRNTIProfileList  *FullIRNTIProfileList
	ShortIRNTIProfileList *ShortIRNTIProfileList
	ChoiceExtension       []byte
}

func (ie *LocalNGRANNodeIdentifier) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(localNGRANNodeIdentifierConstraints)
	switch ie.Choice {
	case 0: // full-I-RNTI-Profile-List
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.FullIRNTIProfileList.Encode(e); err != nil {
			return err
		}
	case 1: // short-I-RNTI-Profile-List
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.ShortIRNTIProfileList.Encode(e); err != nil {
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

func (ie *LocalNGRANNodeIdentifier) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(localNGRANNodeIdentifierConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // full-I-RNTI-Profile-List
		ie.FullIRNTIProfileList = new(FullIRNTIProfileList)
		if err := ie.FullIRNTIProfileList.Decode(d); err != nil {
			return err
		}
	case 1: // short-I-RNTI-Profile-List
		ie.ShortIRNTIProfileList = new(ShortIRNTIProfileList)
		if err := ie.ShortIRNTIProfileList.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
