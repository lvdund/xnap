package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var iABTNLAddressesRequestedConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tNLAddressesOrPrefixesRequestedAllTraffic", Optional: true},
		{Name: "tNLAddressesOrPrefixesRequestedF1-C", Optional: true},
		{Name: "tNLAddressesOrPrefixesRequestedF1-U", Optional: true},
		{Name: "tNLAddressesOrPrefixesRequestedNoNF1", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: nil,
}

type IABTNLAddressesRequested struct {
	TNLAddressesOrPrefixesRequestedAllTraffic *int64
	TNLAddressesOrPrefixesRequestedF1C        *int64
	TNLAddressesOrPrefixesRequestedF1U        *int64
	TNLAddressesOrPrefixesRequestedNoNF1      *int64
	IEExtensions                              []byte
}

func (ie *IABTNLAddressesRequested) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABTNLAddressesRequestedConstraints)
	if err := seq.EncodePreamble([]bool{ie.TNLAddressesOrPrefixesRequestedAllTraffic != nil, ie.TNLAddressesOrPrefixesRequestedF1C != nil, ie.TNLAddressesOrPrefixesRequestedF1U != nil, ie.TNLAddressesOrPrefixesRequestedNoNF1 != nil, false}); err != nil {
		return err
	}
	if ie.TNLAddressesOrPrefixesRequestedAllTraffic != nil {
		if err := e.EncodeInteger(*ie.TNLAddressesOrPrefixesRequestedAllTraffic, per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(256)),
		}); err != nil {
			return err
		}
	}
	if ie.TNLAddressesOrPrefixesRequestedF1C != nil {
		if err := e.EncodeInteger(*ie.TNLAddressesOrPrefixesRequestedF1C, per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(256)),
		}); err != nil {
			return err
		}
	}
	if ie.TNLAddressesOrPrefixesRequestedF1U != nil {
		if err := e.EncodeInteger(*ie.TNLAddressesOrPrefixesRequestedF1U, per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(256)),
		}); err != nil {
			return err
		}
	}
	if ie.TNLAddressesOrPrefixesRequestedNoNF1 != nil {
		if err := e.EncodeInteger(*ie.TNLAddressesOrPrefixesRequestedNoNF1, per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(256)),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *IABTNLAddressesRequested) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABTNLAddressesRequestedConstraints)
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(256)),
		})
		if err != nil {
			return err
		}
		ie.TNLAddressesOrPrefixesRequestedAllTraffic = &val
	}
	if seq.IsComponentPresent(1) {
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(256)),
		})
		if err != nil {
			return err
		}
		ie.TNLAddressesOrPrefixesRequestedF1C = &val
	}
	if seq.IsComponentPresent(2) {
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(256)),
		})
		if err != nil {
			return err
		}
		ie.TNLAddressesOrPrefixesRequestedF1U = &val
	}
	if seq.IsComponentPresent(3) {
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(1)),
			UpperBound: common.Ptr(int64(256)),
		})
		if err != nil {
			return err
		}
		ie.TNLAddressesOrPrefixesRequestedNoNF1 = &val
	}
	return nil
}
