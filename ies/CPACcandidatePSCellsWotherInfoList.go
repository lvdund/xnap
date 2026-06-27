package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var cPACcandidatePSCellsWotherInfoListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.MaxnoofPSCellCandidates)),
}

type CPACcandidatePSCellsWotherInfoList struct {
	Value []*CPACcandidatePSCellsWotherInfoItem
}

func (ie *CPACcandidatePSCellsWotherInfoList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cPACcandidatePSCellsWotherInfoListConstraints)
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

func (ie *CPACcandidatePSCellsWotherInfoList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cPACcandidatePSCellsWotherInfoListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*CPACcandidatePSCellsWotherInfoItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(CPACcandidatePSCellsWotherInfoItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
