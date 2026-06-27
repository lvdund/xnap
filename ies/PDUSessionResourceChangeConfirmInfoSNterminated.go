package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceChangeConfirmInfoSNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dataforwardinginfoTarget", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceChangeConfirmInfoSNterminated struct {
	DataforwardinginfoTarget *DataForwardingInfoFromTargetNGRANnode
	IEExtensions             []byte
}

func (ie *PDUSessionResourceChangeConfirmInfoSNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceChangeConfirmInfoSNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DataforwardinginfoTarget != nil, false}); err != nil {
		return err
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

func (ie *PDUSessionResourceChangeConfirmInfoSNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceChangeConfirmInfoSNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
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
