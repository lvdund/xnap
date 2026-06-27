package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	CoverageModificationCauseCoverage            int64 = 0
	CoverageModificationCauseCellEdgeCapacity    int64 = 1
	CoverageModificationCauseNetworkEnergySaving int64 = 2
)

var coverageModificationCauseConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  []int64{2},
}

type CoverageModificationCause struct {
	Value int64
}

func (ie *CoverageModificationCause) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, coverageModificationCauseConstraints)
}

func (ie *CoverageModificationCause) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(coverageModificationCauseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
