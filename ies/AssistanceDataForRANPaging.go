package ies

import (
	"github.com/lvdund/asn1go/per"
)

var assistanceDataForRANPagingConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ran-paging-attempt-info", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type AssistanceDataForRANPaging struct {
	RanPagingAttemptInfo *RANPagingAttemptInfo
	IEExtensions         []byte
}

func (ie *AssistanceDataForRANPaging) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(assistanceDataForRANPagingConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.RanPagingAttemptInfo != nil, false}); err != nil {
		return err
	}
	if ie.RanPagingAttemptInfo != nil {
		if err := ie.RanPagingAttemptInfo.Encode(e); err != nil {
			return err
		}
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *AssistanceDataForRANPaging) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(assistanceDataForRANPagingConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.RanPagingAttemptInfo = new(RANPagingAttemptInfo)
		if err := ie.RanPagingAttemptInfo.Decode(d); err != nil {
			return err
		}
	}
	extBytes, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	if len(extBytes) > 0 {
		ie.IEExtensions = extBytes[0]
	}
	return nil
}
