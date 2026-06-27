package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	NGRANNodeResourceCoordinationInfoChEutraResourceCoordinationInfo = 0
	NGRANNodeResourceCoordinationInfoChNrResourceCoordinationInfo    = 1
)

var nGRANNodeResourceCoordinationInfoConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "eutra-resource-coordination-info"},
		{Name: "nr-resource-coordination-info"},
	},
	ExtAlternatives: nil,
}

type NGRANNodeResourceCoordinationInfo struct {
	Choice                        int
	EutraResourceCoordinationInfo *EUTRAResourceCoordinationInfo
	NrResourceCoordinationInfo    *NRResourceCoordinationInfo
}

func (ie *NGRANNodeResourceCoordinationInfo) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(nGRANNodeResourceCoordinationInfoConstraints)
	switch ie.Choice {
	case 0: // eutra-resource-coordination-info
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.EutraResourceCoordinationInfo.Encode(e); err != nil {
			return err
		}
	case 1: // nr-resource-coordination-info
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.NrResourceCoordinationInfo.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NGRANNodeResourceCoordinationInfo) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(nGRANNodeResourceCoordinationInfoConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // eutra-resource-coordination-info
		ie.EutraResourceCoordinationInfo = new(EUTRAResourceCoordinationInfo)
		if err := ie.EutraResourceCoordinationInfo.Decode(d); err != nil {
			return err
		}
	case 1: // nr-resource-coordination-info
		ie.NrResourceCoordinationInfo = new(NRResourceCoordinationInfo)
		if err := ie.NrResourceCoordinationInfo.Decode(d); err != nil {
			return err
		}
	}
	return nil
}
