package ies

import (
	"github.com/lvdund/asn1go/per"
)

var nRPagingLongeDRXInformationforRRCINACTIVEConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nRPaging-long-eDRX-Cycle-Inactive"},
		{Name: "nRPaging-Time-Window-Inactive"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type NRPagingLongeDRXInformationforRRCINACTIVE struct {
	NRPagingLongEDRXCycleInactive NRPagingLongEDRXCycleInactive
	NRPagingTimeWindowInactive    NRPagingTimeWindowInactive
	IEExtensions                  []byte
}

func (ie *NRPagingLongeDRXInformationforRRCINACTIVE) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRPagingLongeDRXInformationforRRCINACTIVEConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.NRPagingLongEDRXCycleInactive.Encode(e); err != nil {
		return err
	}
	if err := ie.NRPagingTimeWindowInactive.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *NRPagingLongeDRXInformationforRRCINACTIVE) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRPagingLongeDRXInformationforRRCINACTIVEConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NRPagingLongEDRXCycleInactive.Decode(d); err != nil {
		return err
	}
	if err := ie.NRPagingTimeWindowInactive.Decode(d); err != nil {
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
