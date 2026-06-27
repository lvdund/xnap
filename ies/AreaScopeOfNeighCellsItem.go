package ies

import (
	"github.com/lvdund/asn1go/per"
)

var areaScopeOfNeighCellsItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrFrequencyInfo"},
		{Name: "pciListForMDT", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type AreaScopeOfNeighCellsItem struct {
	NrFrequencyInfo NRFrequencyInfo
	PciListForMDT   *PCIListForMDT
	IEExtensions    []byte
}

func (ie *AreaScopeOfNeighCellsItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(areaScopeOfNeighCellsItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PciListForMDT != nil, false}); err != nil {
		return err
	}
	if err := ie.NrFrequencyInfo.Encode(e); err != nil {
		return err
	}
	if ie.PciListForMDT != nil {
		if err := ie.PciListForMDT.Encode(e); err != nil {
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

func (ie *AreaScopeOfNeighCellsItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(areaScopeOfNeighCellsItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NrFrequencyInfo.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.PciListForMDT = new(PCIListForMDT)
		if err := ie.PciListForMDT.Decode(d); err != nil {
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
