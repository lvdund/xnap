package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRUChannelInfoItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nR-U-ChannelID"},
		{Name: "nRARFCN"},
		{Name: "bandwidth"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRUChannelInfoItem struct {
	NRUChannelID NRUChannelID
	NRARFCN      NRARFCN
	Bandwidth    Bandwidth
	IEExtensions []byte
}

func (ie *NRUChannelInfoItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRUChannelInfoItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.NRUChannelID.Encode(e); err != nil {
		return err
	}
	if err := ie.NRARFCN.Encode(e); err != nil {
		return err
	}
	if err := ie.Bandwidth.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NRUChannelInfoItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRUChannelInfoItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NRUChannelID.Decode(d); err != nil {
		return err
	}
	if err := ie.NRARFCN.Decode(d); err != nil {
		return err
	}
	if err := ie.Bandwidth.Decode(d); err != nil {
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
