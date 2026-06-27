package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RRCReestabInitiatedReportingChRRCReestabReportingWoUERLFReport   = 0
	RRCReestabInitiatedReportingChRRCReestabReportingWithUERLFReport = 1
	RRCReestabInitiatedReportingChChoiceExtension                    = 2
)

var rRCReestabInitiatedReportingConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "rRCReestab-reporting-wo-UERLFReport"},
		{Name: "rRCReestab-reporting-with-UERLFReport"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type RRCReestabInitiatedReporting struct {
	Choice                             int
	RRCReestabReportingWoUERLFReport   *RRCReestabInitiatedReportingWoUERLFReport
	RRCReestabReportingWithUERLFReport *RRCReestabInitiatedReportingWithUERLFReport
	ChoiceExtension                    []byte
}

func (ie *RRCReestabInitiatedReporting) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(rRCReestabInitiatedReportingConstraints)
	switch ie.Choice {
	case 0: // rRCReestab-reporting-wo-UERLFReport
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.RRCReestabReportingWoUERLFReport.Encode(e); err != nil {
			return err
		}
	case 1: // rRCReestab-reporting-with-UERLFReport
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.RRCReestabReportingWithUERLFReport.Encode(e); err != nil {
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

func (ie *RRCReestabInitiatedReporting) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(rRCReestabInitiatedReportingConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // rRCReestab-reporting-wo-UERLFReport
		ie.RRCReestabReportingWoUERLFReport = new(RRCReestabInitiatedReportingWoUERLFReport)
		if err := ie.RRCReestabReportingWoUERLFReport.Decode(d); err != nil {
			return err
		}
	case 1: // rRCReestab-reporting-with-UERLFReport
		ie.RRCReestabReportingWithUERLFReport = new(RRCReestabInitiatedReportingWithUERLFReport)
		if err := ie.RRCReestabReportingWithUERLFReport.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
