package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	SCGreconfigNotificationExecuted        int64 = 0
	SCGreconfigNotificationExecutedDeleted int64 = 1
	SCGreconfigNotificationDeleted         int64 = 2
)

var sCGreconfigNotificationConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0},
	ExtValues:  []int64{1, 2},
}

type SCGreconfigNotification struct {
	Value int64
}

func (ie *SCGreconfigNotification) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, sCGreconfigNotificationConstraints)
}

func (ie *SCGreconfigNotification) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(sCGreconfigNotificationConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
