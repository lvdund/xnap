package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	TrafficProfileChUPTraffic       = 0
	TrafficProfileChNonUPTraffic    = 1
	TrafficProfileChChoiceExtension = 2
)

var trafficProfileConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "uPTraffic"},
		{Name: "nonUPTraffic"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type TrafficProfile struct {
	Choice          int
	UPTraffic       *QoSFlowLevelQoSParameters
	NonUPTraffic    *NonUPTraffic
	ChoiceExtension []byte
}

func (ie *TrafficProfile) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(trafficProfileConstraints)
	switch ie.Choice {
	case 0: // uPTraffic
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.UPTraffic.Encode(e); err != nil {
			return err
		}
	case 1: // nonUPTraffic
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.NonUPTraffic.Encode(e); err != nil {
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

func (ie *TrafficProfile) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(trafficProfileConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // uPTraffic
		ie.UPTraffic = new(QoSFlowLevelQoSParameters)
		if err := ie.UPTraffic.Decode(d); err != nil {
			return err
		}
	case 1: // nonUPTraffic
		ie.NonUPTraffic = new(NonUPTraffic)
		if err := ie.NonUPTraffic.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
