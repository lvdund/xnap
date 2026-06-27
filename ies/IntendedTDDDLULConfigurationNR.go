package ies

import (
	"github.com/lvdund/asn1go/per"
)

var intendedTDDDLULConfigurationNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrscs"},
		{Name: "nrCyclicPrefix"},
		{Name: "nrDL-ULTransmissionPeriodicity"},
		{Name: "slotConfiguration-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type IntendedTDDDLULConfigurationNR struct {
	Nrscs                         NRSCS
	NrCyclicPrefix                NRCyclicPrefix
	NrDLULTransmissionPeriodicity NRDLULTransmissionPeriodicity
	SlotConfigurationList         SlotConfigurationList
	IEExtensions                  []byte
}

func (ie *IntendedTDDDLULConfigurationNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(intendedTDDDLULConfigurationNRConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.Nrscs.Encode(e); err != nil {
		return err
	}
	if err := ie.NrCyclicPrefix.Encode(e); err != nil {
		return err
	}
	if err := ie.NrDLULTransmissionPeriodicity.Encode(e); err != nil {
		return err
	}
	if err := ie.SlotConfigurationList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *IntendedTDDDLULConfigurationNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(intendedTDDDLULConfigurationNRConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.Nrscs.Decode(d); err != nil {
		return err
	}
	if err := ie.NrCyclicPrefix.Decode(d); err != nil {
		return err
	}
	if err := ie.NrDLULTransmissionPeriodicity.Decode(d); err != nil {
		return err
	}
	if err := ie.SlotConfigurationList.Decode(d); err != nil {
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
