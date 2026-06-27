package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ReportTypeChPeriodical     = 0
	ReportTypeChEventTriggered = 1
	ReportTypeExtension        = -1
)

var reportTypeConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "periodical"},
		{Name: "eventTriggered"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "choice-extension"},
	},
}

type ReportType struct {
	Choice          int
	Periodical      *Periodical
	EventTriggered  *EventTriggered
	ChoiceExtension []byte
}

func (ie *ReportType) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(reportTypeConstraints)
	switch ie.Choice {
	case 0: // periodical
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.Periodical.Encode(e); err != nil {
			return err
		}
	case 1: // eventTriggered
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.EventTriggered.Encode(e); err != nil {
			return err
		}
	default: // extension
		if err := choice.EncodeChoice(0, true, ie.ChoiceExtension); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ReportType) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(reportTypeConstraints)
	idx, isExt, extBytes, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	if isExt {
		ie.Choice = ReportTypeExtension
		ie.ChoiceExtension = extBytes
		return nil
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // periodical
		ie.Periodical = new(Periodical)
		if err := ie.Periodical.Decode(d); err != nil {
			return err
		}
	case 1: // eventTriggered
		ie.EventTriggered = new(EventTriggered)
		if err := ie.EventTriggered.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
