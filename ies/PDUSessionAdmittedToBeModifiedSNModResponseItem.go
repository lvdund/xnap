package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionAdmittedToBeModifiedSNModResponseItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionId"},
		{Name: "sn-terminated", Optional: true},
		{Name: "mn-terminated", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionAdmittedToBeModifiedSNModResponseItem struct {
	PduSessionId PDUSessionID
	SnTerminated *PDUSessionResourceModificationResponseInfoSNterminated
	MnTerminated *PDUSessionResourceModificationResponseInfoMNterminated
	IEExtensions []byte
}

func (ie *PDUSessionAdmittedToBeModifiedSNModResponseItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionAdmittedToBeModifiedSNModResponseItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SnTerminated != nil, ie.MnTerminated != nil, false}); err != nil {
		return err
	}
	if err := ie.PduSessionId.Encode(e); err != nil {
		return err
	}
	if ie.SnTerminated != nil {
		if err := ie.SnTerminated.Encode(e); err != nil {
			return err
		}
	}
	if ie.MnTerminated != nil {
		if err := ie.MnTerminated.Encode(e); err != nil {
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

func (ie *PDUSessionAdmittedToBeModifiedSNModResponseItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionAdmittedToBeModifiedSNModResponseItemConstraints)
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
		ie.SnTerminated = new(PDUSessionResourceModificationResponseInfoSNterminated)
		if err := ie.SnTerminated.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.MnTerminated = new(PDUSessionResourceModificationResponseInfoMNterminated)
		if err := ie.MnTerminated.Decode(d); err != nil {
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
