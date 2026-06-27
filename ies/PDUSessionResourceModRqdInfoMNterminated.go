package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceModRqdInfoMNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dRBsToBeModified", Optional: true},
		{Name: "dRBsToBeReleased", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceModRqdInfoMNterminated struct {
	DRBsToBeModified *DRBsToBeModifiedListModRqdMNterminated
	DRBsToBeReleased *DRBListWithCause
	IEExtensions     []byte
}

func (ie *PDUSessionResourceModRqdInfoMNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceModRqdInfoMNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DRBsToBeModified != nil, ie.DRBsToBeReleased != nil, false}); err != nil {
		return err
	}
	if ie.DRBsToBeModified != nil {
		if err := ie.DRBsToBeModified.Encode(e); err != nil {
			return err
		}
	}
	if ie.DRBsToBeReleased != nil {
		if err := ie.DRBsToBeReleased.Encode(e); err != nil {
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

func (ie *PDUSessionResourceModRqdInfoMNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceModRqdInfoMNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.DRBsToBeModified = new(DRBsToBeModifiedListModRqdMNterminated)
		if err := ie.DRBsToBeModified.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.DRBsToBeReleased = new(DRBListWithCause)
		if err := ie.DRBsToBeReleased.Decode(d); err != nil {
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
