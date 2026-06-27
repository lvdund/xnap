package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ProcedureStageChoiceChFirstDlCount    = 0
	ProcedureStageChoiceChDlDiscarding    = 1
	ProcedureStageChoiceChChoiceExtension = 2
)

var procedureStageChoiceConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "first-dl-count"},
		{Name: "dl-discarding"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type ProcedureStageChoice struct {
	Choice          int
	FirstDlCount    *FirstDLCount
	DlDiscarding    *DLDiscarding
	ChoiceExtension []byte
}

func (ie *ProcedureStageChoice) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(procedureStageChoiceConstraints)
	switch ie.Choice {
	case 0: // first-dl-count
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.FirstDlCount.Encode(e); err != nil {
			return err
		}
	case 1: // dl-discarding
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.DlDiscarding.Encode(e); err != nil {
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

func (ie *ProcedureStageChoice) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(procedureStageChoiceConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // first-dl-count
		ie.FirstDlCount = new(FirstDLCount)
		if err := ie.FirstDlCount.Decode(d); err != nil {
			return err
		}
	case 1: // dl-discarding
		ie.DlDiscarding = new(DLDiscarding)
		if err := ie.DlDiscarding.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
