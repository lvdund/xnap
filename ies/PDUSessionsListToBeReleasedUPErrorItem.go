package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionsListToBeReleasedUPErrorItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionId"},
		{Name: "userPlaneErrorIndicator"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionsListToBeReleasedUPErrorItem struct {
	PduSessionId            PDUSessionID
	UserPlaneErrorIndicator UserPlaneErrorIndicator
	IEExtensions            []byte
}

func (ie *PDUSessionsListToBeReleasedUPErrorItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionsListToBeReleasedUPErrorItemConstraints)
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
	if err := ie.UserPlaneErrorIndicator.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSessionsListToBeReleasedUPErrorItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionsListToBeReleasedUPErrorItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionId.Decode(d); err != nil {
		return err
	}
	if err := ie.UserPlaneErrorIndicator.Decode(d); err != nil {
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
