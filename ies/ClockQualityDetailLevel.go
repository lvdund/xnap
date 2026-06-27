package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	ClockQualityDetailLevelChClockQualityMetrics  = 0
	ClockQualityDetailLevelChAcceptanceIndication = 1
	ClockQualityDetailLevelChChoiceExtension      = 2
)

var clockQualityDetailLevelConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "clockQualityMetrics"},
		{Name: "acceptanceIndication"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type ClockQualityDetailLevel struct {
	Choice               int
	ClockQualityMetrics  common.NULL
	AcceptanceIndication *ClockQualityAcceptanceCriteria
	ChoiceExtension      []byte
}

func (ie *ClockQualityDetailLevel) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(clockQualityDetailLevelConstraints)
	switch ie.Choice {
	case 0: // clockQualityMetrics
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case 1: // acceptanceIndication
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.AcceptanceIndication.Encode(e); err != nil {
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

func (ie *ClockQualityDetailLevel) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(clockQualityDetailLevelConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // clockQualityMetrics
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case 1: // acceptanceIndication
		ie.AcceptanceIndication = new(ClockQualityAcceptanceCriteria)
		if err := ie.AcceptanceIndication.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
