package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var qoSFlowsMappedtoDRBSetupResponseMNterminatedConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofQoSFlows)),
}

type QoSFlowsMappedtoDRBSetupResponseMNterminated struct {
	Value []*QoSFlowsMappedtoDRBSetupResponseMNterminatedItem
}

func (ie *QoSFlowsMappedtoDRBSetupResponseMNterminated) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(qoSFlowsMappedtoDRBSetupResponseMNterminatedConstraints)
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

func (ie *QoSFlowsMappedtoDRBSetupResponseMNterminated) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(qoSFlowsMappedtoDRBSetupResponseMNterminatedConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*QoSFlowsMappedtoDRBSetupResponseMNterminatedItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(QoSFlowsMappedtoDRBSetupResponseMNterminatedItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
