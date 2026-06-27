package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SharedResourceTypeULDLSharingChUlResources     = 0
	SharedResourceTypeULDLSharingChDlResources     = 1
	SharedResourceTypeULDLSharingChChoiceExtension = 2
)

var sharedResourceTypeULDLSharingConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ul-resources"},
		{Name: "dl-resources"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type SharedResourceTypeULDLSharing struct {
	Choice          int
	UlResources     *SharedResourceTypeULDLSharingULResources
	DlResources     *SharedResourceTypeULDLSharingDLResources
	ChoiceExtension []byte
}

func (ie *SharedResourceTypeULDLSharing) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(sharedResourceTypeULDLSharingConstraints)
	switch ie.Choice {
	case 0: // ul-resources
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.UlResources.Encode(e); err != nil {
			return err
		}
	case 1: // dl-resources
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.DlResources.Encode(e); err != nil {
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

func (ie *SharedResourceTypeULDLSharing) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(sharedResourceTypeULDLSharingConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // ul-resources
		ie.UlResources = new(SharedResourceTypeULDLSharingULResources)
		if err := ie.UlResources.Decode(d); err != nil {
			return err
		}
	case 1: // dl-resources
		ie.DlResources = new(SharedResourceTypeULDLSharingDLResources)
		if err := ie.DlResources.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
