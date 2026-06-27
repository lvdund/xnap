package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

const (
	AllocationandRetentionPriorityPreEmptionCapabilityShallNotTriggerPreemption int64 = 0
	AllocationandRetentionPriorityPreEmptionCapabilityMayTriggerPreemption      int64 = 1
)

var allocationandRetentionPriorityPreEmptionCapabilityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type AllocationandRetentionPriorityPreEmptionCapability struct {
	Value int64
}

func (ie *AllocationandRetentionPriorityPreEmptionCapability) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, allocationandRetentionPriorityPreEmptionCapabilityConstraints)
}

func (ie *AllocationandRetentionPriorityPreEmptionCapability) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(allocationandRetentionPriorityPreEmptionCapabilityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

const (
	AllocationandRetentionPriorityPreEmptionVulnerabilityNotPreemptable int64 = 0
	AllocationandRetentionPriorityPreEmptionVulnerabilityPreemptable    int64 = 1
)

var allocationandRetentionPriorityPreEmptionVulnerabilityConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type AllocationandRetentionPriorityPreEmptionVulnerability struct {
	Value int64
}

func (ie *AllocationandRetentionPriorityPreEmptionVulnerability) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, allocationandRetentionPriorityPreEmptionVulnerabilityConstraints)
}

func (ie *AllocationandRetentionPriorityPreEmptionVulnerability) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(allocationandRetentionPriorityPreEmptionVulnerabilityConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var allocationandRetentionPriorityConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "priorityLevel"},
		{Name: "pre-emption-capability"},
		{Name: "pre-emption-vulnerability"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type AllocationandRetentionPriority struct {
	PriorityLevel           int64
	PreEmptionCapability    AllocationandRetentionPriorityPreEmptionCapability
	PreEmptionVulnerability AllocationandRetentionPriorityPreEmptionVulnerability
	IEExtensions            []byte
}

func (ie *AllocationandRetentionPriority) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(allocationandRetentionPriorityConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.PriorityLevel, per.IntegerConstraints{
		Extensible: true,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(15)),
	}); err != nil {
		return err
	}
	if err := ie.PreEmptionCapability.Encode(e); err != nil {
		return err
	}
	if err := ie.PreEmptionVulnerability.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *AllocationandRetentionPriority) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(allocationandRetentionPriorityConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: true,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(15)),
		})
		if err != nil {
			return err
		}
		ie.PriorityLevel = val
	}
	if err := ie.PreEmptionCapability.Decode(d); err != nil {
		return err
	}
	if err := ie.PreEmptionVulnerability.Decode(d); err != nil {
		return err
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
