package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	TrafficReleaseTypeChFullRelease     = 0
	TrafficReleaseTypeChPartialRelease  = 1
	TrafficReleaseTypeChChoiceExtension = 2
)

var trafficReleaseTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fullRelease"},
		{Name: "partialRelease"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type TrafficReleaseType struct {
	Choice          int
	FullRelease     *AllTrafficIndication
	PartialRelease  *TrafficToBeReleaseList
	ChoiceExtension []byte
}

func (ie *TrafficReleaseType) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(trafficReleaseTypeConstraints)
	switch ie.Choice {
	case 0: // fullRelease
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.FullRelease.Encode(e); err != nil {
			return err
		}
	case 1: // partialRelease
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.PartialRelease.Encode(e); err != nil {
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

func (ie *TrafficReleaseType) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(trafficReleaseTypeConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // fullRelease
		ie.FullRelease = new(AllTrafficIndication)
		if err := ie.FullRelease.Decode(d); err != nil {
			return err
		}
	case 1: // partialRelease
		ie.PartialRelease = new(TrafficToBeReleaseList)
		if err := ie.PartialRelease.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
