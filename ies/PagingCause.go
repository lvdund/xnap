package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PagingCauseVoice int64 = 0
)

var pagingCauseConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type PagingCause struct {
	Value int64
}

func (ie *PagingCause) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pagingCauseConstraints)
}

func (ie *PagingCause) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pagingCauseConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
