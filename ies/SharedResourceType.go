package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SharedResourceTypeChUlOnlySharing   = 0
	SharedResourceTypeChUlAndDlSharing  = 1
	SharedResourceTypeChChoiceExtension = 2
)

var sharedResourceTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ul-onlySharing"},
		{Name: "ul-and-dl-Sharing"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type SharedResourceType struct {
	Choice          int
	UlOnlySharing   *SharedResourceTypeULOnlySharing
	UlAndDlSharing  *SharedResourceTypeULDLSharing
	ChoiceExtension []byte
}

func (ie *SharedResourceType) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(sharedResourceTypeConstraints)
	switch ie.Choice {
	case 0: // ul-onlySharing
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.UlOnlySharing.Encode(e); err != nil {
			return err
		}
	case 1: // ul-and-dl-Sharing
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.UlAndDlSharing.Encode(e); err != nil {
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

func (ie *SharedResourceType) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(sharedResourceTypeConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // ul-onlySharing
		ie.UlOnlySharing = new(SharedResourceTypeULOnlySharing)
		if err := ie.UlOnlySharing.Decode(d); err != nil {
			return err
		}
	case 1: // ul-and-dl-Sharing
		ie.UlAndDlSharing = new(SharedResourceTypeULDLSharing)
		if err := ie.UlAndDlSharing.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
