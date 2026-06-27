package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionListWithDataForwardingFromTargetItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionId"},
		{Name: "dataforwardinginfoTarget"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionListWithDataForwardingFromTargetItem struct {
	PduSessionId             PDUSessionID
	DataforwardinginfoTarget DataForwardingInfoFromTargetNGRANnode
	IEExtensions             []byte
}

func (ie *PDUSessionListWithDataForwardingFromTargetItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionListWithDataForwardingFromTargetItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PduSessionId.Encode(e); err != nil {
		return err
	}
	if err := ie.DataforwardinginfoTarget.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSessionListWithDataForwardingFromTargetItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionListWithDataForwardingFromTargetItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionId.Decode(d); err != nil {
		return err
	}
	if err := ie.DataforwardinginfoTarget.Decode(d); err != nil {
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
