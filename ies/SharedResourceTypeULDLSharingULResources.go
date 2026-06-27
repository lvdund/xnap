package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	SharedResourceTypeULDLSharingULResourcesChUnchanged       = 0
	SharedResourceTypeULDLSharingULResourcesChChanged         = 1
	SharedResourceTypeULDLSharingULResourcesChChoiceExtension = 2
)

var sharedResourceTypeULDLSharingULResourcesConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "unchanged"},
		{Name: "changed"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type SharedResourceTypeULDLSharingULResources struct {
	Choice          int
	Unchanged       common.NULL
	Changed         *SharedResourceTypeULDLSharingULResourcesChanged
	ChoiceExtension []byte
}

func (ie *SharedResourceTypeULDLSharingULResources) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(sharedResourceTypeULDLSharingULResourcesConstraints)
	switch ie.Choice {
	case 0: // unchanged
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case 1: // changed
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.Changed.Encode(e); err != nil {
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

func (ie *SharedResourceTypeULDLSharingULResources) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(sharedResourceTypeULDLSharingULResourcesConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // unchanged
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case 1: // changed
		ie.Changed = new(SharedResourceTypeULDLSharingULResourcesChanged)
		if err := ie.Changed.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
