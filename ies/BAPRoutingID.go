package ies

import (
	"github.com/lvdund/asn1go/per"
)

var bAPRoutingIDConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bAPAddress"},
		{Name: "bAPPathID"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type BAPRoutingID struct {
	BAPAddress   BAPAddress
	BAPPathID    BAPPathID
	IEExtensions []byte
}

func (ie *BAPRoutingID) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bAPRoutingIDConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.BAPAddress.Encode(e); err != nil {
		return err
	}
	if err := ie.BAPPathID.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *BAPRoutingID) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bAPRoutingIDConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.BAPAddress.Decode(d); err != nil {
		return err
	}
	if err := ie.BAPPathID.Decode(d); err != nil {
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
