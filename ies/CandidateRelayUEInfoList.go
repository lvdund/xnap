package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var candidateRelayUEInfoListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofCandidateRelayUEs)),
}

type CandidateRelayUEInfoList struct {
	Value []*CandidateRelayUEInfoItem
}

func (ie *CandidateRelayUEInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(candidateRelayUEInfoListConstraints)
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

func (ie *CandidateRelayUEInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(candidateRelayUEInfoListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CandidateRelayUEInfoItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CandidateRelayUEInfoItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
