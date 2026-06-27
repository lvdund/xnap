package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionAdmittedSNModResponseConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionResourcesAdmittedToBeAdded", Optional: true},
		{Name: "pduSessionResourcesAdmittedToBeModified", Optional: true},
		{Name: "pduSessionResourcesAdmittedToBeReleased", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionAdmittedSNModResponse struct {
	PduSessionResourcesAdmittedToBeAdded    *PDUSessionAdmittedToBeAddedSNModResponse
	PduSessionResourcesAdmittedToBeModified *PDUSessionAdmittedToBeModifiedSNModResponse
	PduSessionResourcesAdmittedToBeReleased *PDUSessionAdmittedToBeReleasedSNModResponse
	IEExtensions                            []byte
}

func (ie *PDUSessionAdmittedSNModResponse) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionAdmittedSNModResponseConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PduSessionResourcesAdmittedToBeAdded != nil, ie.PduSessionResourcesAdmittedToBeModified != nil, ie.PduSessionResourcesAdmittedToBeReleased != nil, false}); err != nil {
		return err
	}
	if ie.PduSessionResourcesAdmittedToBeAdded != nil {
		if err := ie.PduSessionResourcesAdmittedToBeAdded.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSessionResourcesAdmittedToBeModified != nil {
		if err := ie.PduSessionResourcesAdmittedToBeModified.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSessionResourcesAdmittedToBeReleased != nil {
		if err := ie.PduSessionResourcesAdmittedToBeReleased.Encode(e); err != nil {
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

func (ie *PDUSessionAdmittedSNModResponse) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionAdmittedSNModResponseConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.PduSessionResourcesAdmittedToBeAdded = new(PDUSessionAdmittedToBeAddedSNModResponse)
		if err := ie.PduSessionResourcesAdmittedToBeAdded.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.PduSessionResourcesAdmittedToBeModified = new(PDUSessionAdmittedToBeModifiedSNModResponse)
		if err := ie.PduSessionResourcesAdmittedToBeModified.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.PduSessionResourcesAdmittedToBeReleased = new(PDUSessionAdmittedToBeReleasedSNModResponse)
		if err := ie.PduSessionResourcesAdmittedToBeReleased.Decode(d); err != nil {
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
