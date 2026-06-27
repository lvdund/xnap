package ies

import (
	"github.com/lvdund/asn1go/per"
)

var trafficNotAddedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "trafficIndex"},
		{Name: "casue", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type TrafficNotAddedItem struct {
	TrafficIndex TrafficIndex
	Casue        *Cause
	IEExtensions []byte
}

func (ie *TrafficNotAddedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(trafficNotAddedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.Casue != nil, false}); err != nil {
		return err
	}
	if err := ie.TrafficIndex.Encode(e); err != nil {
		return err
	}
	if ie.Casue != nil {
		if err := ie.Casue.Encode(e); err != nil {
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

func (ie *TrafficNotAddedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(trafficNotAddedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.TrafficIndex.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.Casue = new(Cause)
		if err := ie.Casue.Decode(d); err != nil {
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
