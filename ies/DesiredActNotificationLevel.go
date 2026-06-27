package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	DesiredActNotificationLevelNone       int64 = 0
	DesiredActNotificationLevelQosFlow    int64 = 1
	DesiredActNotificationLevelPduSession int64 = 2
	DesiredActNotificationLevelUeLevel    int64 = 3
)

var desiredActNotificationLevelConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1, 2, 3},
	ExtValues:  nil,
}

type DesiredActNotificationLevel struct {
	Value int64
}

func (ie *DesiredActNotificationLevel) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, desiredActNotificationLevelConstraints)
}

func (ie *DesiredActNotificationLevel) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(desiredActNotificationLevelConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
