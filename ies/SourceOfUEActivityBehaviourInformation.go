package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SourceOfUEActivityBehaviourInformationSubscriptionInformation int64 = 0
	SourceOfUEActivityBehaviourInformationStatistics              int64 = 1
)

var sourceOfUEActivityBehaviourInformationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type SourceOfUEActivityBehaviourInformation struct {
	Value int64
}

func (ie *SourceOfUEActivityBehaviourInformation) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sourceOfUEActivityBehaviourInformationConstraints)
}

func (ie *SourceOfUEActivityBehaviourInformation) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sourceOfUEActivityBehaviourInformationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
