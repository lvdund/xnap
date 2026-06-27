package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionToBeReleasedListRelRqdConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionsToBeReleasedList-SNterminated", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionToBeReleasedListRelRqd struct {
	PduSessionsToBeReleasedListSNterminated *PDUSessionListWithDataForwardingRequest
	IEExtensions                            []byte
}

func (ie *PDUSessionToBeReleasedListRelRqd) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionToBeReleasedListRelRqdConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PduSessionsToBeReleasedListSNterminated != nil, false}); err != nil {
		return err
	}
	if ie.PduSessionsToBeReleasedListSNterminated != nil {
		if err := ie.PduSessionsToBeReleasedListSNterminated.Encode(e); err != nil {
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

func (ie *PDUSessionToBeReleasedListRelRqd) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionToBeReleasedListRelRqdConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.PduSessionsToBeReleasedListSNterminated = new(PDUSessionListWithDataForwardingRequest)
		if err := ie.PduSessionsToBeReleasedListSNterminated.Decode(d); err != nil {
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
