package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourcesNotAdmittedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionId"},
		{Name: "cause", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourcesNotAdmittedItem struct {
	PduSessionId PDUSessionID
	Cause        *Cause
	IEExtensions []byte
}

func (ie *PDUSessionResourcesNotAdmittedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourcesNotAdmittedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.Cause != nil, false}); err != nil {
		return err
	}
	if err := ie.PduSessionId.Encode(e); err != nil {
		return err
	}
	if ie.Cause != nil {
		if err := ie.Cause.Encode(e); err != nil {
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

func (ie *PDUSessionResourcesNotAdmittedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourcesNotAdmittedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionId.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.Cause = new(Cause)
		if err := ie.Cause.Decode(d); err != nil {
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
