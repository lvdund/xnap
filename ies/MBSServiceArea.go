package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MBSServiceAreaChLocationindependent = 0
	MBSServiceAreaChLocationdependent   = 1
	MBSServiceAreaChChoiceExtension     = 2
)

var mBSServiceAreaConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "locationindependent"},
		{Name: "locationdependent"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type MBSServiceArea struct {
	Choice              int
	Locationindependent *MBSServiceAreaInformation
	Locationdependent   *MBSServiceAreaInformationList
	ChoiceExtension     []byte
}

func (ie *MBSServiceArea) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(mBSServiceAreaConstraints)
	switch ie.Choice {
	case 0: // locationindependent
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.Locationindependent.Encode(e); err != nil {
			return err
		}
	case 1: // locationdependent
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.Locationdependent.Encode(e); err != nil {
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

func (ie *MBSServiceArea) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(mBSServiceAreaConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // locationindependent
		ie.Locationindependent = new(MBSServiceAreaInformation)
		if err := ie.Locationindependent.Decode(d); err != nil {
			return err
		}
	case 1: // locationdependent
		ie.Locationdependent = new(MBSServiceAreaInformationList)
		if err := ie.Locationdependent.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
