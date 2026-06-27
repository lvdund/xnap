package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SliceAvailabilityListChUnavailableCellList = 0
	SliceAvailabilityListChAvailableCellList   = 1
	SliceAvailabilityListChChoiceExtension     = 2
	SliceAvailabilityListExtension             = -1
)

var sliceAvailabilityListConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "unavailableCellList"},
		{Name: "availableCellList"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type SliceAvailabilityList struct {
	Choice              int
	UnavailableCellList *UnavailableCellList
	AvailableCellList   *AvailableCellList
	ChoiceExtension     []byte
}

func (ie *SliceAvailabilityList) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(sliceAvailabilityListConstraints)
	switch ie.Choice {
	case 0: // unavailableCellList
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.UnavailableCellList.Encode(e); err != nil {
			return err
		}
	case 1: // availableCellList
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.AvailableCellList.Encode(e); err != nil {
			return err
		}
	case 2: // choice-extension
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	default:
		if err := choice.EncodeChoice(0, true, ie.ChoiceExtension); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SliceAvailabilityList) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(sliceAvailabilityListConstraints)
	idx, isExt, extBytes, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	if isExt {
		ie.Choice = SliceAvailabilityListExtension
		ie.ChoiceExtension = extBytes
		return nil
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // unavailableCellList
		ie.UnavailableCellList = new(UnavailableCellList)
		if err := ie.UnavailableCellList.Decode(d); err != nil {
			return err
		}
	case 1: // availableCellList
		ie.AvailableCellList = new(AvailableCellList)
		if err := ie.AvailableCellList.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
