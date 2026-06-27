package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	TNLAssociationUsageUe    int64 = 0
	TNLAssociationUsageNonUe int64 = 1
	TNLAssociationUsageBoth  int64 = 2
)

var tNLAssociationUsageConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2},
	ExtValues:  nil,
}

type TNLAssociationUsage struct {
	Value int64
}

func (ie *TNLAssociationUsage) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, tNLAssociationUsageConstraints)
}

func (ie *TNLAssociationUsage) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(tNLAssociationUsageConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
