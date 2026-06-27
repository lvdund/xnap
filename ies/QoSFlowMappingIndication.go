package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	QoSFlowMappingIndicationUl int64 = 0
	QoSFlowMappingIndicationDl int64 = 1
)

var qoSFlowMappingIndicationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type QoSFlowMappingIndication struct {
	Value int64
}

func (ie *QoSFlowMappingIndication) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, qoSFlowMappingIndicationConstraints)
}

func (ie *QoSFlowMappingIndication) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(qoSFlowMappingIndicationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
