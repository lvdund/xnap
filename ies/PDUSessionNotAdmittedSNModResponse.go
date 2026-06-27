package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionNotAdmittedSNModResponseConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pdu-Session-List", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionNotAdmittedSNModResponse struct {
	PduSessionList *PDUSessionList
	IEExtensions   []byte
}

func (ie *PDUSessionNotAdmittedSNModResponse) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionNotAdmittedSNModResponseConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PduSessionList != nil, false}); err != nil {
		return err
	}
	if ie.PduSessionList != nil {
		if err := ie.PduSessionList.Encode(e); err != nil {
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

func (ie *PDUSessionNotAdmittedSNModResponse) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionNotAdmittedSNModResponseConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.PduSessionList = new(PDUSessionList)
		if err := ie.PduSessionList.Decode(d); err != nil {
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
