package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var tAISupportItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tac"},
		{Name: "broadcastPLMNs"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TAISupportItem struct {
	Tac            TAC
	BroadcastPLMNs []*BroadcastPLMNinTAISupportItem
	IEExtensions   []byte
}

func (ie *TAISupportItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tAISupportItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.Tac.Encode(e); err != nil {
		return err
	}
	soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
		Extensible: false,
		Min:        common.Ptr(int64(1)),
		Max:        common.Ptr(int64(common.MaxnoofsupportedPLMNs)),
	})
	if err := soEnc.EncodeLength(int64(len(ie.BroadcastPLMNs))); err != nil {
		return err
	}
	for _, item := range ie.BroadcastPLMNs {
		if err := item.Encode(e); err != nil {
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

func (ie *TAISupportItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tAISupportItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.Tac.Decode(d); err != nil {
		return err
	}
	{
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofsupportedPLMNs)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.BroadcastPLMNs = make([]*BroadcastPLMNinTAISupportItem, n)
		for i := range ie.BroadcastPLMNs {
			ie.BroadcastPLMNs[i] = new(BroadcastPLMNinTAISupportItem)
			if err := ie.BroadcastPLMNs[i].Decode(d); err != nil {
				return err
			}
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
