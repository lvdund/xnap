package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionsToBeAddedSNModRequestItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionId"},
		{Name: "s-NSSAI"},
		{Name: "sN-PDUSessionAMBR", Optional: true},
		{Name: "sn-terminated", Optional: true},
		{Name: "mn-terminated", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionsToBeAddedSNModRequestItem struct {
	PduSessionId     PDUSessionID
	SNSSAI           SNSSAI
	SNPDUSessionAMBR *PDUSessionAggregateMaximumBitRate
	SnTerminated     *PDUSessionResourceSetupInfoSNterminated
	MnTerminated     *PDUSessionResourceSetupInfoMNterminated
	IEExtensions     []byte
}

func (ie *PDUSessionsToBeAddedSNModRequestItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionsToBeAddedSNModRequestItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SNPDUSessionAMBR != nil, ie.SnTerminated != nil, ie.MnTerminated != nil, false}); err != nil {
		return err
	}
	if err := ie.PduSessionId.Encode(e); err != nil {
		return err
	}
	if err := ie.SNSSAI.Encode(e); err != nil {
		return err
	}
	if ie.SNPDUSessionAMBR != nil {
		if err := ie.SNPDUSessionAMBR.Encode(e); err != nil {
			return err
		}
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

func (ie *PDUSessionsToBeAddedSNModRequestItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionsToBeAddedSNModRequestItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionId.Decode(d); err != nil {
		return err
	}
	if err := ie.SNSSAI.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.SNPDUSessionAMBR = new(PDUSessionAggregateMaximumBitRate)
		if err := ie.SNPDUSessionAMBR.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.SnTerminated = new(PDUSessionResourceSetupInfoSNterminated)
		if err := ie.SnTerminated.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.MnTerminated = new(PDUSessionResourceSetupInfoMNterminated)
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
