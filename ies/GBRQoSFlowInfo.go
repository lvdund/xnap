package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	GBRQoSFlowInfoNotificationControlNotificationRequested int64 = 0
)

var gBRQoSFlowInfoNotificationControlConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type GBRQoSFlowInfoNotificationControl struct {
	Value int64
}

func (ie *GBRQoSFlowInfoNotificationControl) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, gBRQoSFlowInfoNotificationControlConstraints)
}

func (ie *GBRQoSFlowInfoNotificationControl) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(gBRQoSFlowInfoNotificationControlConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var gBRQoSFlowInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maxFlowBitRateDL"},
		{Name: "maxFlowBitRateUL"},
		{Name: "guaranteedFlowBitRateDL"},
		{Name: "guaranteedFlowBitRateUL"},
		{Name: "notificationControl", Optional: true},
		{Name: "maxPacketLossRateDL", Optional: true},
		{Name: "maxPacketLossRateUL", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type GBRQoSFlowInfo struct {
	MaxFlowBitRateDL        BitRate
	MaxFlowBitRateUL        BitRate
	GuaranteedFlowBitRateDL BitRate
	GuaranteedFlowBitRateUL BitRate
	NotificationControl     *GBRQoSFlowInfoNotificationControl
	MaxPacketLossRateDL     *PacketLossRate
	MaxPacketLossRateUL     *PacketLossRate
	IEExtensions            []byte
}

func (ie *GBRQoSFlowInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gBRQoSFlowInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NotificationControl != nil, ie.MaxPacketLossRateDL != nil, ie.MaxPacketLossRateUL != nil, false}); err != nil {
		return err
	}
	if err := ie.MaxFlowBitRateDL.Encode(e); err != nil {
		return err
	}
	if err := ie.MaxFlowBitRateUL.Encode(e); err != nil {
		return err
	}
	if err := ie.GuaranteedFlowBitRateDL.Encode(e); err != nil {
		return err
	}
	if err := ie.GuaranteedFlowBitRateUL.Encode(e); err != nil {
		return err
	}
	if ie.NotificationControl != nil {
		if err := ie.NotificationControl.Encode(e); err != nil {
			return err
		}
	}
	if ie.MaxPacketLossRateDL != nil {
		if err := ie.MaxPacketLossRateDL.Encode(e); err != nil {
			return err
		}
	}
	if ie.MaxPacketLossRateUL != nil {
		if err := ie.MaxPacketLossRateUL.Encode(e); err != nil {
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

func (ie *GBRQoSFlowInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gBRQoSFlowInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.MaxFlowBitRateDL.Decode(d); err != nil {
		return err
	}
	if err := ie.MaxFlowBitRateUL.Decode(d); err != nil {
		return err
	}
	if err := ie.GuaranteedFlowBitRateDL.Decode(d); err != nil {
		return err
	}
	if err := ie.GuaranteedFlowBitRateUL.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(4) {
		ie.NotificationControl = new(GBRQoSFlowInfoNotificationControl)
		if err := ie.NotificationControl.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.MaxPacketLossRateDL = new(PacketLossRate)
		if err := ie.MaxPacketLossRateDL.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.MaxPacketLossRateUL = new(PacketLossRate)
		if err := ie.MaxPacketLossRateUL.Decode(d); err != nil {
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
