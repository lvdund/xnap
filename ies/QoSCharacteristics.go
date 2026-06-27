package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	QoSCharacteristicsChNonDynamic      = 0
	QoSCharacteristicsChDynamic         = 1
	QoSCharacteristicsChChoiceExtension = 2
)

var qoSCharacteristicsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "non-dynamic"},
		{Name: "dynamic"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type QoSCharacteristics struct {
	Choice          int
	NonDynamic      *NonDynamic5QIDescriptor
	Dynamic         *Dynamic5QIDescriptor
	ChoiceExtension []byte
}

func (ie *QoSCharacteristics) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(qoSCharacteristicsConstraints)
	switch ie.Choice {
	case 0: // non-dynamic
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.NonDynamic.Encode(e); err != nil {
			return err
		}
	case 1: // dynamic
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.Dynamic.Encode(e); err != nil {
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

func (ie *QoSCharacteristics) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(qoSCharacteristicsConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // non-dynamic
		ie.NonDynamic = new(NonDynamic5QIDescriptor)
		if err := ie.NonDynamic.Decode(d); err != nil {
			return err
		}
	case 1: // dynamic
		ie.Dynamic = new(Dynamic5QIDescriptor)
		if err := ie.Dynamic.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
