package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SrcSNToTgtSNQMCInfoInquiryTrue int64 = 0
)

var srcSNToTgtSNQMCInfoInquiryConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type SrcSNToTgtSNQMCInfoInquiry struct {
	Value int64
}

func (ie *SrcSNToTgtSNQMCInfoInquiry) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, srcSNToTgtSNQMCInfoInquiryConstraints)
}

func (ie *SrcSNToTgtSNQMCInfoInquiry) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(srcSNToTgtSNQMCInfoInquiryConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
