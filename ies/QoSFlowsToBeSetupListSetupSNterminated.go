package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qoSFlowsToBeSetupListSetupSNterminatedConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofQoSFlows)),
}

type QoSFlowsToBeSetupListSetupSNterminated struct {
	Value []*QoSFlowsToBeSetupListSetupSNterminatedItem
}

func (ie *QoSFlowsToBeSetupListSetupSNterminated) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(qoSFlowsToBeSetupListSetupSNterminatedConstraints)
	if err := seqOf.EncodeLength(int64(len(ie.Value))); err != nil {
		return err
	}
	for _, item := range ie.Value {
		if err := item.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *QoSFlowsToBeSetupListSetupSNterminated) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(qoSFlowsToBeSetupListSetupSNterminatedConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*QoSFlowsToBeSetupListSetupSNterminatedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(QoSFlowsToBeSetupListSetupSNterminatedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
