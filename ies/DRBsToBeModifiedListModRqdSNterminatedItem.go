package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dRBsToBeModifiedListModRqdSNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "sN-UL-PDCP-UP-TNLInfo", Optional: true},
		{Name: "dRB-QoS", Optional: true},
		{Name: "secondary-SN-UL-PDCP-UP-TNLInfo", Optional: true},
		{Name: "uL-Configuration", Optional: true},
		{Name: "pdcpDuplicationConfiguration", Optional: true},
		{Name: "duplicationActivation", Optional: true},
		{Name: "qoSFlowsMappedtoDRB-ModRqd-SNterminated", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBsToBeModifiedListModRqdSNterminatedItem struct {
	DrbID                                 DRBID
	SNULPDCPUPTNLInfo                     *UPTransportParameters
	DRBQoS                                *QoSFlowLevelQoSParameters
	SecondarySNULPDCPUPTNLInfo            *UPTransportParameters
	ULConfiguration                       *ULConfiguration
	PdcpDuplicationConfiguration          *PDCPDuplicationConfiguration
	DuplicationActivation                 *DuplicationActivation
	QoSFlowsMappedtoDRBModRqdSNterminated *QoSFlowsModifiedMappedtoDRBModRqdSNterminated
	IEExtensions                          []byte
}

func (ie *DRBsToBeModifiedListModRqdSNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBsToBeModifiedListModRqdSNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SNULPDCPUPTNLInfo != nil, ie.DRBQoS != nil, ie.SecondarySNULPDCPUPTNLInfo != nil, ie.ULConfiguration != nil, ie.PdcpDuplicationConfiguration != nil, ie.DuplicationActivation != nil, ie.QoSFlowsMappedtoDRBModRqdSNterminated != nil, false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if ie.SNULPDCPUPTNLInfo != nil {
		if err := ie.SNULPDCPUPTNLInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.DRBQoS != nil {
		if err := ie.DRBQoS.Encode(e); err != nil {
			return err
		}
	}
	if ie.SecondarySNULPDCPUPTNLInfo != nil {
		if err := ie.SecondarySNULPDCPUPTNLInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.ULConfiguration != nil {
		if err := ie.ULConfiguration.Encode(e); err != nil {
			return err
		}
	}
	if ie.PdcpDuplicationConfiguration != nil {
		if err := ie.PdcpDuplicationConfiguration.Encode(e); err != nil {
			return err
		}
	}
	if ie.DuplicationActivation != nil {
		if err := ie.DuplicationActivation.Encode(e); err != nil {
			return err
		}
	}
	if ie.QoSFlowsMappedtoDRBModRqdSNterminated != nil {
		if err := ie.QoSFlowsMappedtoDRBModRqdSNterminated.Encode(e); err != nil {
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

func (ie *DRBsToBeModifiedListModRqdSNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBsToBeModifiedListModRqdSNterminatedItemConstraints)
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
		ie.SNULPDCPUPTNLInfo = new(UPTransportParameters)
		if err := ie.SNULPDCPUPTNLInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.DRBQoS = new(QoSFlowLevelQoSParameters)
		if err := ie.DRBQoS.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.SecondarySNULPDCPUPTNLInfo = new(UPTransportParameters)
		if err := ie.SecondarySNULPDCPUPTNLInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.ULConfiguration = new(ULConfiguration)
		if err := ie.ULConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.PdcpDuplicationConfiguration = new(PDCPDuplicationConfiguration)
		if err := ie.PdcpDuplicationConfiguration.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.DuplicationActivation = new(DuplicationActivation)
		if err := ie.DuplicationActivation.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.QoSFlowsMappedtoDRBModRqdSNterminated = new(QoSFlowsModifiedMappedtoDRBModRqdSNterminated)
		if err := ie.QoSFlowsMappedtoDRBModRqdSNterminated.Decode(d); err != nil {
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
