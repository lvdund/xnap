package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRPagingeDRXInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nRPaging-eDRX-Cycle"},
		{Name: "nRPaging-Time-Window", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRPagingeDRXInformation struct {
	NRPagingEDRXCycle  NRPagingEDRXCycle
	NRPagingTimeWindow *NRPagingTimeWindow
	IEExtensions       []byte
}

func (ie *NRPagingeDRXInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRPagingeDRXInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.NRPagingTimeWindow != nil, false}); err != nil {
		return err
	}
	if err := ie.NRPagingEDRXCycle.Encode(e); err != nil {
		return err
	}
	if ie.NRPagingTimeWindow != nil {
		if err := ie.NRPagingTimeWindow.Encode(e); err != nil {
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

func (ie *NRPagingeDRXInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRPagingeDRXInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NRPagingEDRXCycle.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.NRPagingTimeWindow = new(NRPagingTimeWindow)
		if err := ie.NRPagingTimeWindow.Decode(d); err != nil {
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
