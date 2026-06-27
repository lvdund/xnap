package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	IABTNLAddressChIPv4Address     = 0
	IABTNLAddressChIPv6Address     = 1
	IABTNLAddressChIPv6Prefix      = 2
	IABTNLAddressChChoiceExtension = 3
)

var iABTNLAddressConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "iPv4Address"},
		{Name: "iPv6Address"},
		{Name: "iPv6Prefix"},
		{Name: "choice-extension"},
	},
	ExtAlternatives: nil,
}

type IABTNLAddress struct {
	Choice          int
	IPv4Address     *per.BitString
	IPv6Address     *per.BitString
	IPv6Prefix      *per.BitString
	ChoiceExtension []byte
}

func (ie *IABTNLAddress) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(iABTNLAddressConstraints)
	switch ie.Choice {
	case 0: // iPv4Address
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.IPv4Address, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(32)),
			Max:        common.Ptr(int64(32)),
		}); err != nil {
			return err
		}
	case 1: // iPv6Address
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.IPv6Address, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(128)),
			Max:        common.Ptr(int64(128)),
		}); err != nil {
			return err
		}
	case 2: // iPv6Prefix
		if err := choice.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err := e.EncodeBitString(*ie.IPv6Prefix, per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(64)),
			Max:        common.Ptr(int64(64)),
		}); err != nil {
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

func (ie *IABTNLAddress) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(iABTNLAddressConstraints)
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0: // iPv4Address
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(32)),
			Max:        common.Ptr(int64(32)),
		})
		if err != nil {
			return err
		}
		ie.IPv4Address = &val
	case 1: // iPv6Address
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(128)),
			Max:        common.Ptr(int64(128)),
		})
		if err != nil {
			return err
		}
		ie.IPv6Address = &val
	case 2: // iPv6Prefix
		val, err := d.DecodeBitString(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(64)),
			Max:        common.Ptr(int64(64)),
		})
		if err != nil {
			return err
		}
		ie.IPv6Prefix = &val
	case 3: // choice-extension
		// TODO decode field ChoiceExtension (kind=ext)
	}
	return nil
}
