package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	Dynamic5QIDescriptorDelayCriticalDelayCritical    int64 = 0
	Dynamic5QIDescriptorDelayCriticalNonDelayCritical int64 = 1
)

var dynamic5QIDescriptorDelayCriticalConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type Dynamic5QIDescriptorDelayCritical struct {
	Value int64
}

func (ie *Dynamic5QIDescriptorDelayCritical) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, dynamic5QIDescriptorDelayCriticalConstraints)
}

func (ie *Dynamic5QIDescriptorDelayCritical) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(dynamic5QIDescriptorDelayCriticalConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var dynamic5QIDescriptorConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "priorityLevelQoS"},
		{Name: "packetDelayBudget"},
		{Name: "packetErrorRate"},
		{Name: "fiveQI", Optional: true},
		{Name: "delayCritical", Optional: true},
		{Name: "averagingWindow", Optional: true},
		{Name: "maximumDataBurstVolume", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type Dynamic5QIDescriptor struct {
	PriorityLevelQoS       PriorityLevelQoS
	PacketDelayBudget      PacketDelayBudget
	PacketErrorRate        PacketErrorRate
	FiveQI                 *FiveQI
	DelayCritical          *Dynamic5QIDescriptorDelayCritical
	AveragingWindow        *AveragingWindow
	MaximumDataBurstVolume *MaximumDataBurstVolume
	IEExtensions           []byte
}

func (ie *Dynamic5QIDescriptor) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dynamic5QIDescriptorConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.FiveQI != nil, ie.DelayCritical != nil, ie.AveragingWindow != nil, ie.MaximumDataBurstVolume != nil, false}); err != nil {
		return err
	}
	if err := ie.PriorityLevelQoS.Encode(e); err != nil {
		return err
	}
	if err := ie.PacketDelayBudget.Encode(e); err != nil {
		return err
	}
	if err := ie.PacketErrorRate.Encode(e); err != nil {
		return err
	}
	if ie.FiveQI != nil {
		if err := ie.FiveQI.Encode(e); err != nil {
			return err
		}
	}
	if ie.DelayCritical != nil {
		if err := ie.DelayCritical.Encode(e); err != nil {
			return err
		}
	}
	if ie.AveragingWindow != nil {
		if err := ie.AveragingWindow.Encode(e); err != nil {
			return err
		}
	}
	if ie.MaximumDataBurstVolume != nil {
		if err := ie.MaximumDataBurstVolume.Encode(e); err != nil {
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

func (ie *Dynamic5QIDescriptor) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dynamic5QIDescriptorConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PriorityLevelQoS.Decode(d); err != nil {
		return err
	}
	if err := ie.PacketDelayBudget.Decode(d); err != nil {
		return err
	}
	if err := ie.PacketErrorRate.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(3) {
		ie.FiveQI = new(FiveQI)
		if err := ie.FiveQI.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.DelayCritical = new(Dynamic5QIDescriptorDelayCritical)
		if err := ie.DelayCritical.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.AveragingWindow = new(AveragingWindow)
		if err := ie.AveragingWindow.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.MaximumDataBurstVolume = new(MaximumDataBurstVolume)
		if err := ie.MaximumDataBurstVolume.Decode(d); err != nil {
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
