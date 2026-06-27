package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	ConfigurationUpdateInitiatingNodeChoiceChGNB             = 0
	ConfigurationUpdateInitiatingNodeChoiceChNgENB           = 1
	ConfigurationUpdateInitiatingNodeChoiceChChoiceExtension = 2
)

var configurationUpdateInitiatingNodeChoiceConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "gNB"},
		{Name: "ng-eNB"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type ConfigurationUpdateInitiatingNodeChoice struct {
	Choice          int
	GNB             []byte // opaque
	NgENB           []byte // opaque
	ChoiceExtension []byte
}

func (ie *ConfigurationUpdateInitiatingNodeChoice) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(configurationUpdateInitiatingNodeChoiceConstraints)
	switch ie.Choice {
	case 0: // gNB
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		// TODO encode field GNB (kind=opaque)
	case 1: // ng-eNB
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		// TODO encode field NgENB (kind=opaque)
	case 2: // choice-extension
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		// TODO encode field ChoiceExtension (kind=ext)
	}
	return nil
}

func (ie *ConfigurationUpdateInitiatingNodeChoice) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(configurationUpdateInitiatingNodeChoiceConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // gNB
		// TODO decode field GNB (kind=opaque)
	case 1: // ng-eNB
		// TODO decode field NgENB (kind=opaque)
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
