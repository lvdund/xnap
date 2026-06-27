package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var volumeTimedReportListConstraints = per.SizeConstraints{
	Extensible: false,
	Min:        common.Ptr(int64(1)),
	Max:        common.Ptr(int64(common.Maxnooftimeperiods)),
}

type VolumeTimedReportList struct {
	Value []*VolumeTimedReportItem
}

func (ie *VolumeTimedReportList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(volumeTimedReportListConstraints)
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

func (ie *VolumeTimedReportList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(volumeTimedReportListConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	ie.Value = make([]*VolumeTimedReportItem, n)
	for i := range ie.Value {
		ie.Value[i] = new(VolumeTimedReportItem)
		if err := ie.Value[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
