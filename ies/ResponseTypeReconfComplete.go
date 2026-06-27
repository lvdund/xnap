package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ResponseTypeReconfCompleteChConfigurationSuccessfullyApplied  = 0
	ResponseTypeReconfCompleteChConfigurationRejectedByMNGRANNode = 1
	ResponseTypeReconfCompleteChChoiceExtension                   = 2
)

var responseTypeReconfCompleteConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "configuration-successfully-applied"},
		{Name: "configuration-rejected-by-M-NG-RANNode"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type ResponseTypeReconfComplete struct {
	Choice                            int
	ConfigurationSuccessfullyApplied  *ConfigurationSuccessfullyApplied
	ConfigurationRejectedByMNGRANNode *ConfigurationRejectedByMNGRANNode
	ChoiceExtension                   []byte
}

func (ie *ResponseTypeReconfComplete) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(responseTypeReconfCompleteConstraints)
	switch ie.Choice {
	case 0: // configuration-successfully-applied
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.ConfigurationSuccessfullyApplied.Encode(e); err != nil {
			return err
		}
	case 1: // configuration-rejected-by-M-NG-RANNode
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.ConfigurationRejectedByMNGRANNode.Encode(e); err != nil {
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

func (ie *ResponseTypeReconfComplete) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(responseTypeReconfCompleteConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // configuration-successfully-applied
		ie.ConfigurationSuccessfullyApplied = new(ConfigurationSuccessfullyApplied)
		if err := ie.ConfigurationSuccessfullyApplied.Decode(d); err != nil {
			return err
		}
	case 1: // configuration-rejected-by-M-NG-RANNode
		ie.ConfigurationRejectedByMNGRANNode = new(ConfigurationRejectedByMNGRANNode)
		if err := ie.ConfigurationRejectedByMNGRANNode.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
