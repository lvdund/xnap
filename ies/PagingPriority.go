package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	PagingPriorityPriolevel1 int64 = 0
	PagingPriorityPriolevel2 int64 = 1
	PagingPriorityPriolevel3 int64 = 2
	PagingPriorityPriolevel4 int64 = 3
	PagingPriorityPriolevel5 int64 = 4
	PagingPriorityPriolevel6 int64 = 5
	PagingPriorityPriolevel7 int64 = 6
	PagingPriorityPriolevel8 int64 = 7
)

var pagingPriorityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3, 4, 5, 6, 7},
	ExtValues:  nil,
}

type PagingPriority struct {
	Value int64
}

func (ie *PagingPriority) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, pagingPriorityConstraints)
}

func (ie *PagingPriority) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(pagingPriorityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
