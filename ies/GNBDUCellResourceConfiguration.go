package ies

import (
	"github.com/lvdund/asn1go/per"
)

var gNBDUCellResourceConfigurationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "subcarrierSpacing"},
		{Name: "dUFTransmissionPeriodicity", Optional: true},
		{Name: "dUF-Slot-Config-List", Optional: true},
		{Name: "hSNATransmissionPeriodicity"},
		{Name: "hNSASlotConfigList", Optional: true},
		{Name: "rBsetConfiguration", Optional: true},
		{Name: "freqDomainHSNAconfiguration-List", Optional: true},
		{Name: "nACellResourceConfigurationList", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type GNBDUCellResourceConfiguration struct {
	SubcarrierSpacing               SubcarrierSpacing
	DUFTransmissionPeriodicity      *DUFTransmissionPeriodicity
	DUFSlotConfigList               *DUFSlotConfigList
	HSNATransmissionPeriodicity     HSNATransmissionPeriodicity
	HNSASlotConfigList              *HSNASlotConfigList
	RBsetConfiguration              *RBsetConfiguration
	FreqDomainHSNAconfigurationList *FreqDomainHSNAconfigurationList
	NACellResourceConfigurationList *NACellResourceConfigurationList
	IEExtensions                    []byte
}

func (ie *GNBDUCellResourceConfiguration) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gNBDUCellResourceConfigurationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DUFTransmissionPeriodicity != nil, ie.DUFSlotConfigList != nil, ie.HNSASlotConfigList != nil, ie.RBsetConfiguration != nil, ie.FreqDomainHSNAconfigurationList != nil, ie.NACellResourceConfigurationList != nil, false}); err != nil {
		return err
	}
	if err := ie.SubcarrierSpacing.Encode(e); err != nil {
		return err
	}
	if ie.DUFTransmissionPeriodicity != nil {
		if err := ie.DUFTransmissionPeriodicity.Encode(e); err != nil {
			return err
		}
	}
	if ie.DUFSlotConfigList != nil {
		if err := ie.DUFSlotConfigList.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.HSNATransmissionPeriodicity.Encode(e); err != nil {
		return err
	}
	if ie.HNSASlotConfigList != nil {
		if err := ie.HNSASlotConfigList.Encode(e); err != nil {
			return err
		}
	}
	if ie.RBsetConfiguration != nil {
		if err := ie.RBsetConfiguration.Encode(e); err != nil {
			return err
		}
	}
	if ie.FreqDomainHSNAconfigurationList != nil {
		if err := ie.FreqDomainHSNAconfigurationList.Encode(e); err != nil {
			return err
		}
	}
	if ie.NACellResourceConfigurationList != nil {
		if err := ie.NACellResourceConfigurationList.Encode(e); err != nil {
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

func (ie *GNBDUCellResourceConfiguration) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gNBDUCellResourceConfigurationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.SubcarrierSpacing.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.DUFTransmissionPeriodicity = new(DUFTransmissionPeriodicity)
		if err := ie.DUFTransmissionPeriodicity.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.DUFSlotConfigList = new(DUFSlotConfigList)
		if err := ie.DUFSlotConfigList.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.HSNATransmissionPeriodicity.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(4) {
		ie.HNSASlotConfigList = new(HSNASlotConfigList)
		if err := ie.HNSASlotConfigList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.RBsetConfiguration = new(RBsetConfiguration)
		if err := ie.RBsetConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.FreqDomainHSNAconfigurationList = new(FreqDomainHSNAconfigurationList)
		if err := ie.FreqDomainHSNAconfigurationList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.NACellResourceConfigurationList = new(NACellResourceConfigurationList)
		if err := ie.NACellResourceConfigurationList.Decode(d); err != nil {
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
