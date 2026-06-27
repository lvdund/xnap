package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionResourceSetupInfoMNterminatedConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionType"},
		{Name: "dRBsToBeSetup"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionResourceSetupInfoMNterminated struct {
	PduSessionType PDUSessionType
	DRBsToBeSetup  DRBsToBeSetupListSetupMNterminated
	IEExtensions   []byte
}

func (ie *PDUSessionResourceSetupInfoMNterminated) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionResourceSetupInfoMNterminatedConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.PduSessionType.Encode(e); err != nil {
		return err
	}
	if err := ie.DRBsToBeSetup.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *PDUSessionResourceSetupInfoMNterminated) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionResourceSetupInfoMNterminatedConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionType.Decode(d); err != nil {
		return err
	}
	if err := ie.DRBsToBeSetup.Decode(d); err != nil {
		return err
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
