package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	InitiatingNodeTypeResourceCoordRequestChNgENB           = 0
	InitiatingNodeTypeResourceCoordRequestChGNB             = 1
	InitiatingNodeTypeResourceCoordRequestChChoiceExtension = 2
)

var initiatingNodeTypeResourceCoordRequestConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ng-eNB"},
		{Name: "gNB"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type InitiatingNodeTypeResourceCoordRequest struct {
	Choice          int
	NgENB           *ResourceCoordRequestNgENBInitiated
	GNB             *ResourceCoordRequestGNBInitiated
	ChoiceExtension []byte
}

func (ie *InitiatingNodeTypeResourceCoordRequest) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(initiatingNodeTypeResourceCoordRequestConstraints)
	switch ie.Choice {
	case 0: // ng-eNB
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.NgENB.Encode(e); err != nil {
			return err
		}
	case 1: // gNB
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.GNB.Encode(e); err != nil {
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

func (ie *InitiatingNodeTypeResourceCoordRequest) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(initiatingNodeTypeResourceCoordRequestConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // ng-eNB
		ie.NgENB = new(ResourceCoordRequestNgENBInitiated)
		if err := ie.NgENB.Decode(d); err != nil {
			return err
		}
	case 1: // gNB
		ie.GNB = new(ResourceCoordRequestGNBInitiated)
		if err := ie.GNB.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
