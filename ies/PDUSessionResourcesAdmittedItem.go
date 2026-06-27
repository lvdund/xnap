package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourcesAdmittedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionId"},
		{Name: "pduSessionResourceAdmittedInfo"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourcesAdmittedItem struct {
	PduSessionId                   PDUSessionID
	PduSessionResourceAdmittedInfo PDUSessionResourceAdmittedInfo
	IEExtensions                   []byte
}

func (ie *PDUSessionResourcesAdmittedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourcesAdmittedItemConstraints)
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
	if err := ie.PduSessionResourceAdmittedInfo.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSessionResourcesAdmittedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourcesAdmittedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionId.Decode(d); err != nil {
		return err
	}
	if err := ie.PduSessionResourceAdmittedInfo.Decode(d); err != nil {
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
