package ies

import (
	"github.com/lvdund/asn1go/per"
)

var broadcastPLMNinTAISupportItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-id"},
		{Name: "tAISliceSupport-List"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type BroadcastPLMNinTAISupportItem struct {
	PlmnId              PLMNIdentity
	TAISliceSupportList SliceSupportList
	IEExtensions        []byte
}

func (ie *BroadcastPLMNinTAISupportItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(broadcastPLMNinTAISupportItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PlmnId.Encode(e); err != nil {
		return err
	}
	if err := ie.TAISliceSupportList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *BroadcastPLMNinTAISupportItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(broadcastPLMNinTAISupportItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PlmnId.Decode(d); err != nil {
		return err
	}
	if err := ie.TAISliceSupportList.Decode(d); err != nil {
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
