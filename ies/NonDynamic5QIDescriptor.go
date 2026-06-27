package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nonDynamic5QIDescriptorConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "fiveQI"},
		{Name: "priorityLevelQoS", Optional: true},
		{Name: "averagingWindow", Optional: true},
		{Name: "maximumDataBurstVolume", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NonDynamic5QIDescriptor struct {
	FiveQI                 FiveQI
	PriorityLevelQoS       *PriorityLevelQoS
	AveragingWindow        *AveragingWindow
	MaximumDataBurstVolume *MaximumDataBurstVolume
	IEExtensions           []byte
}

func (ie *NonDynamic5QIDescriptor) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nonDynamic5QIDescriptorConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PriorityLevelQoS != nil, ie.AveragingWindow != nil, ie.MaximumDataBurstVolume != nil, false}); err != nil {
		return err
	}
	if err := ie.FiveQI.Encode(e); err != nil {
		return err
	}
	if ie.PriorityLevelQoS != nil {
		if err := ie.PriorityLevelQoS.Encode(e); err != nil {
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

func (ie *NonDynamic5QIDescriptor) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nonDynamic5QIDescriptorConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.FiveQI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.PriorityLevelQoS = new(PriorityLevelQoS)
		if err := ie.PriorityLevelQoS.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.AveragingWindow = new(AveragingWindow)
		if err := ie.AveragingWindow.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
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
