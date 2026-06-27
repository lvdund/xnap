package ies

import (
	"github.com/lvdund/asn1go/per"
)

const (
	RLCDuplicationStateItemDuplicationStateActive   int64 = 0
	RLCDuplicationStateItemDuplicationStateInactive int64 = 1
)

var rLCDuplicationStateItemDuplicationStateConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{0, 1},
	ExtValues:  nil,
}

type RLCDuplicationStateItemDuplicationState struct {
	Value int64
}

func (ie *RLCDuplicationStateItemDuplicationState) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(ie.Value, rLCDuplicationStateItemDuplicationStateConstraints)
}

func (ie *RLCDuplicationStateItemDuplicationState) Decode(d *per.Decoder) error {
	val, err := d.DecodeEnumerated(rLCDuplicationStateItemDuplicationStateConstraints)
	if err != nil {
		return err
	}
	ie.Value = val
	return nil
}

var rLCDuplicationStateItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "duplicationState"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type RLCDuplicationStateItem struct {
	DuplicationState RLCDuplicationStateItemDuplicationState
	IEExtensions     []byte
}

func (ie *RLCDuplicationStateItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rLCDuplicationStateItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DuplicationState.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *RLCDuplicationStateItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rLCDuplicationStateItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DuplicationState.Decode(d); err != nil {
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
