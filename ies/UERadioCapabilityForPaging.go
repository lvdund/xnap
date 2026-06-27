package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uERadioCapabilityForPagingConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uERadioCapabilityForPagingOfNR", Optional: true},
		{Name: "uERadioCapabilityForPagingOfEUTRA", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    *UERadioCapabilityForPagingOfNR
	UERadioCapabilityForPagingOfEUTRA *UERadioCapabilityForPagingOfEUTRA
	IEExtensions                      []byte
}

func (ie *UERadioCapabilityForPaging) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uERadioCapabilityForPagingConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.UERadioCapabilityForPagingOfNR != nil, ie.UERadioCapabilityForPagingOfEUTRA != nil, false}); err != nil {
		return err
	}
	if ie.UERadioCapabilityForPagingOfNR != nil {
		if err := ie.UERadioCapabilityForPagingOfNR.Encode(e); err != nil {
			return err
		}
	}
	if ie.UERadioCapabilityForPagingOfEUTRA != nil {
		if err := ie.UERadioCapabilityForPagingOfEUTRA.Encode(e); err != nil {
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

func (ie *UERadioCapabilityForPaging) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uERadioCapabilityForPagingConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.UERadioCapabilityForPagingOfNR = new(UERadioCapabilityForPagingOfNR)
		if err := ie.UERadioCapabilityForPagingOfNR.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.UERadioCapabilityForPagingOfEUTRA = new(UERadioCapabilityForPagingOfEUTRA)
		if err := ie.UERadioCapabilityForPagingOfEUTRA.Decode(d); err != nil {
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
