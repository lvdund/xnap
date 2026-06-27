package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sliceAvailableCapacityItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pLMNIdentity"},
		{Name: "sNSSAIAvailableCapacity-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SliceAvailableCapacityItem struct {
	PLMNIdentity                PLMNIdentity
	SNSSAIAvailableCapacityList SNSSAIAvailableCapacityList
	IEExtensions                []byte
}

func (ie *SliceAvailableCapacityItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sliceAvailableCapacityItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PLMNIdentity.Encode(e); err != nil {
		return err
	}
	if err := ie.SNSSAIAvailableCapacityList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SliceAvailableCapacityItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sliceAvailableCapacityItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PLMNIdentity.Decode(d); err != nil {
		return err
	}
	if err := ie.SNSSAIAvailableCapacityList.Decode(d); err != nil {
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
