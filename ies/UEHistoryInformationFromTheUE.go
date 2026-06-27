package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	UEHistoryInformationFromTheUEChNR              = 0
	UEHistoryInformationFromTheUEChChoiceExtension = 1
)

var uEHistoryInformationFromTheUEConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nR"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type UEHistoryInformationFromTheUE struct {
	Choice          int
	NR              *NRMobilityHistoryReport
	ChoiceExtension []byte
}

func (ie *UEHistoryInformationFromTheUE) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(uEHistoryInformationFromTheUEConstraints)
	switch ie.Choice {
	case 0: // nR
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.NR.Encode(e); err != nil {
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

func (ie *UEHistoryInformationFromTheUE) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(uEHistoryInformationFromTheUEConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // nR
		ie.NR = new(NRMobilityHistoryReport)
		if err := ie.NR.Decode(d); err != nil {
			return err
		}
	case 1: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
