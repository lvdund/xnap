package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	IABTNLAddressUsageF1C   int64 = 0
	IABTNLAddressUsageF1U   int64 = 1
	IABTNLAddressUsageNonF1 int64 = 2
	IABTNLAddressUsageAll   int64 = 3
)

var iABTNLAddressUsageConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  []int64{3},
}

type IABTNLAddressUsage struct {
	Value int64
}

func (ie *IABTNLAddressUsage) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, iABTNLAddressUsageConstraints)
}

func (ie *IABTNLAddressUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(iABTNLAddressUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
