package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dRBsToBeSetupListSetupMNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "mN-UL-PDCP-UP-TNLInfo"},
		{Name: "rLC-Mode"},
		{Name: "uL-Configuration", Optional: true},
		{Name: "dRB-QoS"},
		{Name: "pDCP-SNLength", Optional: true},
		{Name: "secondary-MN-UL-PDCP-UP-TNLInfo", Optional: true},
		{Name: "duplicationActivation", Optional: true},
		{Name: "qoSFlowsMappedtoDRB-Setup-MNterminated"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBsToBeSetupListSetupMNterminatedItem struct {
	DrbID                                DRBID
	MNULPDCPUPTNLInfo                    UPTransportParameters
	RLCMode                              RLCMode
	ULConfiguration                      *ULConfiguration
	DRBQoS                               QoSFlowLevelQoSParameters
	PDCPSNLength                         *PDCPSNLength
	SecondaryMNULPDCPUPTNLInfo           *UPTransportParameters
	DuplicationActivation                *DuplicationActivation
	QoSFlowsMappedtoDRBSetupMNterminated QoSFlowsMappedtoDRBSetupMNterminated
	IEExtensions                         []byte
}

func (ie *DRBsToBeSetupListSetupMNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBsToBeSetupListSetupMNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.ULConfiguration != nil, ie.PDCPSNLength != nil, ie.SecondaryMNULPDCPUPTNLInfo != nil, ie.DuplicationActivation != nil, false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if err := ie.MNULPDCPUPTNLInfo.Encode(e); err != nil {
		return err
	}
	if err := ie.RLCMode.Encode(e); err != nil {
		return err
	}
	if ie.ULConfiguration != nil {
		if err := ie.ULConfiguration.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.DRBQoS.Encode(e); err != nil {
		return err
	}
	if ie.PDCPSNLength != nil {
		if err := ie.PDCPSNLength.Encode(e); err != nil {
			return err
		}
	}
	if ie.SecondaryMNULPDCPUPTNLInfo != nil {
		if err := ie.SecondaryMNULPDCPUPTNLInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.DuplicationActivation != nil {
		if err := ie.DuplicationActivation.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.QoSFlowsMappedtoDRBSetupMNterminated.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DRBsToBeSetupListSetupMNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBsToBeSetupListSetupMNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DrbID.Decode(d); err != nil {
		return err
	}
	if err := ie.MNULPDCPUPTNLInfo.Decode(d); err != nil {
		return err
	}
	if err := ie.RLCMode.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(3) {
		ie.ULConfiguration = new(ULConfiguration)
		if err := ie.ULConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.DRBQoS.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(5) {
		ie.PDCPSNLength = new(PDCPSNLength)
		if err := ie.PDCPSNLength.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.SecondaryMNULPDCPUPTNLInfo = new(UPTransportParameters)
		if err := ie.SecondaryMNULPDCPUPTNLInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.DuplicationActivation = new(DuplicationActivation)
		if err := ie.DuplicationActivation.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.QoSFlowsMappedtoDRBSetupMNterminated.Decode(d); err != nil {
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
