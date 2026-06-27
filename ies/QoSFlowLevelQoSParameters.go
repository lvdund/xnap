package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	QoSFlowLevelQoSParametersAdditionalQoSflowInfoMoreLikely int64 = 0
)

var qoSFlowLevelQoSParametersAdditionalQoSflowInfoConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0},
	ExtValues:  nil,
}

type QoSFlowLevelQoSParametersAdditionalQoSflowInfo struct {
	Value int64
}

func (ie *QoSFlowLevelQoSParametersAdditionalQoSflowInfo) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, qoSFlowLevelQoSParametersAdditionalQoSflowInfoConstraints)
}

func (ie *QoSFlowLevelQoSParametersAdditionalQoSflowInfo) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(qoSFlowLevelQoSParametersAdditionalQoSflowInfoConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var qoSFlowLevelQoSParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qos-characteristics"},
		{Name: "allocationAndRetentionPrio"},
		{Name: "gBRQoSFlowInfo", Optional: true},
		{Name: "reflectiveQoS", Optional: true},
		{Name: "additionalQoSflowInfo", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type QoSFlowLevelQoSParameters struct {
	QosCharacteristics         QoSCharacteristics
	AllocationAndRetentionPrio AllocationandRetentionPriority
	GBRQoSFlowInfo             *GBRQoSFlowInfo
	ReflectiveQoS              *ReflectiveQoSAttribute
	AdditionalQoSflowInfo      *QoSFlowLevelQoSParametersAdditionalQoSflowInfo
	IEExtensions               []byte
}

func (ie *QoSFlowLevelQoSParameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qoSFlowLevelQoSParametersConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.GBRQoSFlowInfo != nil, ie.ReflectiveQoS != nil, ie.AdditionalQoSflowInfo != nil, false}); err != nil {
		return err
	}
	if err := ie.QosCharacteristics.Encode(e); err != nil {
		return err
	}
	if err := ie.AllocationAndRetentionPrio.Encode(e); err != nil {
		return err
	}
	if ie.GBRQoSFlowInfo != nil {
		if err := ie.GBRQoSFlowInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.ReflectiveQoS != nil {
		if err := ie.ReflectiveQoS.Encode(e); err != nil {
			return err
		}
	}
	if ie.AdditionalQoSflowInfo != nil {
		if err := ie.AdditionalQoSflowInfo.Encode(e); err != nil {
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

func (ie *QoSFlowLevelQoSParameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qoSFlowLevelQoSParametersConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QosCharacteristics.Decode(d); err != nil {
		return err
	}
	if err := ie.AllocationAndRetentionPrio.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.GBRQoSFlowInfo = new(GBRQoSFlowInfo)
		if err := ie.GBRQoSFlowInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.ReflectiveQoS = new(ReflectiveQoSAttribute)
		if err := ie.ReflectiveQoS.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.AdditionalQoSflowInfo = new(QoSFlowLevelQoSParametersAdditionalQoSflowInfo)
		if err := ie.AdditionalQoSflowInfo.Decode(d); err != nil {
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
