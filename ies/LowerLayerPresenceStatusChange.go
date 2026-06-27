package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	LowerLayerPresenceStatusChangeReleaseLowerLayers     int64 = 0
	LowerLayerPresenceStatusChangeReEstablishLowerLayers int64 = 1
	LowerLayerPresenceStatusChangeSuspendLowerLayers     int64 = 2
	LowerLayerPresenceStatusChangeResumeLowerLayers      int64 = 3
)

var lowerLayerPresenceStatusChangeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  []int64{2, 3},
}

type LowerLayerPresenceStatusChange struct {
	Value int64
}

func (ie *LowerLayerPresenceStatusChange) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, lowerLayerPresenceStatusChangeConstraints)
}

func (ie *LowerLayerPresenceStatusChange) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(lowerLayerPresenceStatusChangeConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}
