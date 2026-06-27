package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	MeasurementThresholdA2ChThresholdRSRP   = 0
	MeasurementThresholdA2ChThresholdRSRQ   = 1
	MeasurementThresholdA2ChThresholdSINR   = 2
	MeasurementThresholdA2ChChoiceExtension = 3
)

var measurementThresholdA2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "threshold-RSRP"},
		{Name: "threshold-RSRQ"},
		{Name: "threshold-SINR"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type MeasurementThresholdA2 struct {
	Choice          int
	ThresholdRSRP   *ThresholdRSRP
	ThresholdRSRQ   *ThresholdRSRQ
	ThresholdSINR   *ThresholdSINR
	ChoiceExtension []byte
}

func (ie *MeasurementThresholdA2) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(measurementThresholdA2Constraints)
	switch ie.Choice {
	case 0: // threshold-RSRP
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.ThresholdRSRP.Encode(e); err != nil {
			return err
		}
	case 1: // threshold-RSRQ
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.ThresholdRSRQ.Encode(e); err != nil {
			return err
		}
	case 2: // threshold-SINR
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := ie.ThresholdSINR.Encode(e); err != nil {
			return err
		}
	case 3: // choice-extension
		if err := choice.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *MeasurementThresholdA2) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(measurementThresholdA2Constraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // threshold-RSRP
		ie.ThresholdRSRP = new(ThresholdRSRP)
		if err := ie.ThresholdRSRP.Decode(d); err != nil {
			return err
		}
	case 1: // threshold-RSRQ
		ie.ThresholdRSRQ = new(ThresholdRSRQ)
		if err := ie.ThresholdRSRQ.Decode(d); err != nil {
			return err
		}
	case 2: // threshold-SINR
		ie.ThresholdSINR = new(ThresholdSINR)
		if err := ie.ThresholdSINR.Decode(d); err != nil {
			return err
		}
	case 3: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
