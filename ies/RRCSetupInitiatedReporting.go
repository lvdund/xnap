package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RRCSetupInitiatedReportingChRRCSetupReportingWithUERLFReport = 0
	RRCSetupInitiatedReportingChChoiceExtension                  = 1
)

var rRCSetupInitiatedReportingConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rRCSetup-reporting-with-UERLFReport"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type RRCSetupInitiatedReporting struct {
	Choice                           int
	RRCSetupReportingWithUERLFReport *RRCSetupInitiatedReportingWithUERLFReport
	ChoiceExtension                  []byte
}

func (ie *RRCSetupInitiatedReporting) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(rRCSetupInitiatedReportingConstraints)
	switch ie.Choice {
	case 0: // rRCSetup-reporting-with-UERLFReport
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.RRCSetupReportingWithUERLFReport.Encode(e); err != nil {
			return err
		}
	case 1: // choice-extension
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *RRCSetupInitiatedReporting) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(rRCSetupInitiatedReportingConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // rRCSetup-reporting-with-UERLFReport
		ie.RRCSetupReportingWithUERLFReport = new(RRCSetupInitiatedReportingWithUERLFReport)
		if err := ie.RRCSetupReportingWithUERLFReport.Decode(d); err != nil {
			return err
		}
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
