package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceModConfirmInfoSNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uL-NG-U-TNLatUPF", Optional: true},
		{Name: "dRBsAdmittedList"},
		{Name: "dRBsNotAdmittedSetupModifyList", Optional: true},
		{Name: "dataforwardinginfoTarget", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceModConfirmInfoSNterminated struct {
	ULNGUTNLatUPF                  *UPTransportLayerInformation
	DRBsAdmittedList               DRBsAdmittedListModConfirmSNterminated
	DRBsNotAdmittedSetupModifyList *DRBListWithCause
	DataforwardinginfoTarget       *DataForwardingInfoFromTargetNGRANnode
	IEExtensions                   []byte
}

func (ie *PDUSessionResourceModConfirmInfoSNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceModConfirmInfoSNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ULNGUTNLatUPF != nil, ie.DRBsNotAdmittedSetupModifyList != nil, ie.DataforwardinginfoTarget != nil, false}); err != nil {
		return err
	}
	if ie.ULNGUTNLatUPF != nil {
		if err := ie.ULNGUTNLatUPF.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.DRBsAdmittedList.Encode(e); err != nil {
		return err
	}
	if ie.DRBsNotAdmittedSetupModifyList != nil {
		if err := ie.DRBsNotAdmittedSetupModifyList.Encode(e); err != nil {
			return err
		}
	}
	if ie.DataforwardinginfoTarget != nil {
		if err := ie.DataforwardinginfoTarget.Encode(e); err != nil {
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

func (ie *PDUSessionResourceModConfirmInfoSNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceModConfirmInfoSNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.ULNGUTNLatUPF = new(UPTransportLayerInformation)
		if err := ie.ULNGUTNLatUPF.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.DRBsAdmittedList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.DRBsNotAdmittedSetupModifyList = new(DRBListWithCause)
		if err := ie.DRBsNotAdmittedSetupModifyList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.DataforwardinginfoTarget = new(DataForwardingInfoFromTargetNGRANnode)
		if err := ie.DataforwardinginfoTarget.Decode(d); err != nil {
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
