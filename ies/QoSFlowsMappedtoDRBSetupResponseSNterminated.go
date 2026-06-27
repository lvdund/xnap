package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qoSFlowsMappedtoDRBSetupResponseSNterminatedConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofQoSFlows)),
}

type QoSFlowsMappedtoDRBSetupResponseSNterminated struct {
	Value []*QoSFlowsMappedtoDRBSetupResponseSNterminatedItem
}

func (ie *QoSFlowsMappedtoDRBSetupResponseSNterminated) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(qoSFlowsMappedtoDRBSetupResponseSNterminatedConstraints)
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

func (ie *QoSFlowsMappedtoDRBSetupResponseSNterminated) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(qoSFlowsMappedtoDRBSetupResponseSNterminatedConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*QoSFlowsMappedtoDRBSetupResponseSNterminatedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(QoSFlowsMappedtoDRBSetupResponseSNterminatedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
