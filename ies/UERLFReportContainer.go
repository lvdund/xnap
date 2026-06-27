package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	UERLFReportContainerChNRUERLFReportContainer  = 0
	UERLFReportContainerChLTEUERLFReportContainer = 1
	UERLFReportContainerChChoiceExtension         = 2
)

var uERLFReportContainerConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nR-UERLFReportContainer"},
		{Name: "lTE-UERLFReportContainer"},
		{Name: "choice-Extension"},
	},
	ExtAlternatives: nil,
}

type UERLFReportContainer struct {
	Choice                  int
	NRUERLFReportContainer  *UERLFReportContainerNR
	LTEUERLFReportContainer *UERLFReportContainerLTE
	ChoiceExtension         []byte
}

func (ie *UERLFReportContainer) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(uERLFReportContainerConstraints)
	switch ie.Choice {
	case 0: // nR-UERLFReportContainer
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.NRUERLFReportContainer.Encode(e); err != nil {
			return err
		}
	case 1: // lTE-UERLFReportContainer
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.LTEUERLFReportContainer.Encode(e); err != nil {
			return err
		}
	case 2: // choice-Extension
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *UERLFReportContainer) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(uERLFReportContainerConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // nR-UERLFReportContainer
		ie.NRUERLFReportContainer = new(UERLFReportContainerNR)
		if err := ie.NRUERLFReportContainer.Decode(d); err != nil {
			return err
		}
	case 1: // lTE-UERLFReportContainer
		ie.LTEUERLFReportContainer = new(UERLFReportContainerLTE)
		if err := ie.LTEUERLFReportContainer.Decode(d); err != nil {
			return err
		}
	case 2: // choice-Extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
