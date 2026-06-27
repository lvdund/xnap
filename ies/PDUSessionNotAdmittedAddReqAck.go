package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionNotAdmittedAddReqAckConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionResourcesNotAdmitted-SNterminated", Optional: true},
		{Name: "pduSessionResourcesNotAdmitted-MNterminated", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionNotAdmittedAddReqAck struct {
	PduSessionResourcesNotAdmittedSNterminated *PDUSessionResourcesNotAdmittedList
	PduSessionResourcesNotAdmittedMNterminated *PDUSessionResourcesNotAdmittedList
	IEExtensions                               []byte
}

func (ie *PDUSessionNotAdmittedAddReqAck) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionNotAdmittedAddReqAckConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PduSessionResourcesNotAdmittedSNterminated != nil, ie.PduSessionResourcesNotAdmittedMNterminated != nil, false}); err != nil {
		return err
	}
	if ie.PduSessionResourcesNotAdmittedSNterminated != nil {
		if err := ie.PduSessionResourcesNotAdmittedSNterminated.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSessionResourcesNotAdmittedMNterminated != nil {
		if err := ie.PduSessionResourcesNotAdmittedMNterminated.Encode(e); err != nil {
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

func (ie *PDUSessionNotAdmittedAddReqAck) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionNotAdmittedAddReqAckConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.PduSessionResourcesNotAdmittedSNterminated = new(PDUSessionResourcesNotAdmittedList)
		if err := ie.PduSessionResourcesNotAdmittedSNterminated.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.PduSessionResourcesNotAdmittedMNterminated = new(PDUSessionResourcesNotAdmittedList)
		if err := ie.PduSessionResourcesNotAdmittedMNterminated.Decode(d); err != nil {
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
