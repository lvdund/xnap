package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dRBsToBeSetupListModRqdSNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "pDCP-SNLength", Optional: true},
		{Name: "sn-UL-PDCP-UPTNLinfo"},
		{Name: "dRB-QoS"},
		{Name: "secondary-SN-UL-PDCP-UP-TNLInfo", Optional: true},
		{Name: "duplicationActivation", Optional: true},
		{Name: "uL-Configuration", Optional: true},
		{Name: "qoSFlowsMappedtoDRB-ModRqd-SNterminated"},
		{Name: "rLC-Mode"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBsToBeSetupListModRqdSNterminatedItem struct {
	DrbID                                 DRBID
	PDCPSNLength                          *PDCPSNLength
	SnULPDCPUPTNLinfo                     UPTransportParameters
	DRBQoS                                QoSFlowLevelQoSParameters
	SecondarySNULPDCPUPTNLInfo            *UPTransportParameters
	DuplicationActivation                 *DuplicationActivation
	ULConfiguration                       *ULConfiguration
	QoSFlowsMappedtoDRBModRqdSNterminated QoSFlowsSetupMappedtoDRBModRqdSNterminated
	RLCMode                               RLCMode
	IEExtensions                          []byte
}

func (ie *DRBsToBeSetupListModRqdSNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBsToBeSetupListModRqdSNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.PDCPSNLength != nil, ie.SecondarySNULPDCPUPTNLInfo != nil, ie.DuplicationActivation != nil, ie.ULConfiguration != nil, false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if ie.PDCPSNLength != nil {
		if err := ie.PDCPSNLength.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.SnULPDCPUPTNLinfo.Encode(e); err != nil {
		return err
	}
	if err := ie.DRBQoS.Encode(e); err != nil {
		return err
	}
	if ie.SecondarySNULPDCPUPTNLInfo != nil {
		if err := ie.SecondarySNULPDCPUPTNLInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.DuplicationActivation != nil {
		if err := ie.DuplicationActivation.Encode(e); err != nil {
			return err
		}
	}
	if ie.ULConfiguration != nil {
		if err := ie.ULConfiguration.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.QoSFlowsMappedtoDRBModRqdSNterminated.Encode(e); err != nil {
		return err
	}
	if err := ie.RLCMode.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *DRBsToBeSetupListModRqdSNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBsToBeSetupListModRqdSNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DrbID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.PDCPSNLength = new(PDCPSNLength)
		if err := ie.PDCPSNLength.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.SnULPDCPUPTNLinfo.Decode(d); err != nil {
		return err
	}
	if err := ie.DRBQoS.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(4) {
		ie.SecondarySNULPDCPUPTNLInfo = new(UPTransportParameters)
		if err := ie.SecondarySNULPDCPUPTNLInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.DuplicationActivation = new(DuplicationActivation)
		if err := ie.DuplicationActivation.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.ULConfiguration = new(ULConfiguration)
		if err := ie.ULConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.QoSFlowsMappedtoDRBModRqdSNterminated.Decode(d); err != nil {
		return err
	}
	if err := ie.RLCMode.Decode(d); err != nil {
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
