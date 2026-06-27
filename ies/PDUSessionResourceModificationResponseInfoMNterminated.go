package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceModificationResponseInfoMNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dRBsAdmittedList"},
		{Name: "dRBsReleasedList", Optional: true},
		{Name: "dRBsNotAdmittedSetupModifyList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceModificationResponseInfoMNterminated struct {
	DRBsAdmittedList               DRBsAdmittedListModificationResponseMNterminated
	DRBsReleasedList               *DRBList
	DRBsNotAdmittedSetupModifyList *DRBListWithCause
	IEExtensions                   []byte
}

func (ie *PDUSessionResourceModificationResponseInfoMNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceModificationResponseInfoMNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DRBsReleasedList != nil, ie.DRBsNotAdmittedSetupModifyList != nil, false}); err != nil {
		return err
	}
	if err := ie.DRBsAdmittedList.Encode(e); err != nil {
		return err
	}
	if ie.DRBsReleasedList != nil {
		if err := ie.DRBsReleasedList.Encode(e); err != nil {
			return err
		}
	}
	if ie.DRBsNotAdmittedSetupModifyList != nil {
		if err := ie.DRBsNotAdmittedSetupModifyList.Encode(e); err != nil {
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

func (ie *PDUSessionResourceModificationResponseInfoMNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceModificationResponseInfoMNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DRBsAdmittedList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.DRBsReleasedList = new(DRBList)
		if err := ie.DRBsReleasedList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.DRBsNotAdmittedSetupModifyList = new(DRBListWithCause)
		if err := ie.DRBsNotAdmittedSetupModifyList.Decode(d); err != nil {
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
