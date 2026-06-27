package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	IABIPv6RequestTypeChIPv6Address     = 0
	IABIPv6RequestTypeChIPv6Prefix      = 1
	IABIPv6RequestTypeChChoiceExtension = 2
)

var iABIPv6RequestTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "iPv6Address"},
		{Name: "iPv6Prefix"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type IABIPv6RequestType struct {
	Choice          int
	IPv6Address     *IABTNLAddressesRequested
	IPv6Prefix      *IABTNLAddressesRequested
	ChoiceExtension []byte
}

func (ie *IABIPv6RequestType) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(iABIPv6RequestTypeConstraints)
	switch ie.Choice {
	case 0: // iPv6Address
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := ie.IPv6Address.Encode(e); err != nil {
			return err
		}
	case 1: // iPv6Prefix
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.IPv6Prefix.Encode(e); err != nil {
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

func (ie *IABIPv6RequestType) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(iABIPv6RequestTypeConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // iPv6Address
		ie.IPv6Address = new(IABTNLAddressesRequested)
		if err := ie.IPv6Address.Decode(d); err != nil {
			return err
		}
	case 1: // iPv6Prefix
		ie.IPv6Prefix = new(IABTNLAddressesRequested)
		if err := ie.IPv6Prefix.Decode(d); err != nil {
			return err
		}
	case 2: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
