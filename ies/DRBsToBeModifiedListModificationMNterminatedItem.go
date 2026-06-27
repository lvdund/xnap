package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dRBsToBeModifiedListModificationMNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "mN-UL-PDCP-UP-TNLInfo", Optional: true},
		{Name: "dRB-QoS", Optional: true},
		{Name: "secondary-MN-UL-PDCP-UP-TNLInfo", Optional: true},
		{Name: "uL-Configuration", Optional: true},
		{Name: "pdcpDuplicationConfiguration", Optional: true},
		{Name: "duplicationActivation", Optional: true},
		{Name: "qoSFlowsMappedtoDRB-Setup-MNterminated", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBsToBeModifiedListModificationMNterminatedItem struct {
	DrbID                                DRBID
	MNULPDCPUPTNLInfo                    *UPTransportParameters
	DRBQoS                               *QoSFlowLevelQoSParameters
	SecondaryMNULPDCPUPTNLInfo           *UPTransportParameters
	ULConfiguration                      *ULConfiguration
	PdcpDuplicationConfiguration         *PDCPDuplicationConfiguration
	DuplicationActivation                *DuplicationActivation
	QoSFlowsMappedtoDRBSetupMNterminated *QoSFlowsMappedtoDRBSetupMNterminated
	IEExtensions                         []byte
}

func (ie *DRBsToBeModifiedListModificationMNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBsToBeModifiedListModificationMNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MNULPDCPUPTNLInfo != nil, ie.DRBQoS != nil, ie.SecondaryMNULPDCPUPTNLInfo != nil, ie.ULConfiguration != nil, ie.PdcpDuplicationConfiguration != nil, ie.DuplicationActivation != nil, ie.QoSFlowsMappedtoDRBSetupMNterminated != nil, false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if ie.MNULPDCPUPTNLInfo != nil {
		if err := ie.MNULPDCPUPTNLInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.DRBQoS != nil {
		if err := ie.DRBQoS.Encode(e); err != nil {
			return err
		}
	}
	if ie.SecondaryMNULPDCPUPTNLInfo != nil {
		if err := ie.SecondaryMNULPDCPUPTNLInfo.Encode(e); err != nil {
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
	if ie.QoSFlowsMappedtoDRBSetupMNterminated != nil {
		if err := ie.QoSFlowsMappedtoDRBSetupMNterminated.Encode(e); err != nil {
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

func (ie *DRBsToBeModifiedListModificationMNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBsToBeModifiedListModificationMNterminatedItemConstraints)
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
		ie.MNULPDCPUPTNLInfo = new(UPTransportParameters)
		if err := ie.MNULPDCPUPTNLInfo.Decode(d); err != nil {
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
		ie.SecondaryMNULPDCPUPTNLInfo = new(UPTransportParameters)
		if err := ie.SecondaryMNULPDCPUPTNLInfo.Decode(d); err != nil {
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
		ie.QoSFlowsMappedtoDRBSetupMNterminated = new(QoSFlowsMappedtoDRBSetupMNterminated)
		if err := ie.QoSFlowsMappedtoDRBSetupMNterminated.Decode(d); err != nil {
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
