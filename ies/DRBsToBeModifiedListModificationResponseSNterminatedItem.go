package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dRBsToBeModifiedListModificationResponseSNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "sN-UL-PDCP-UP-TNLInfo", Optional: true},
		{Name: "dRB-QoS", Optional: true},
		{Name: "qoSFlowsMappedtoDRB-SetupResponse-SNterminated", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBsToBeModifiedListModificationResponseSNterminatedItem struct {
	DrbID                                        DRBID
	SNULPDCPUPTNLInfo                            *UPTransportParameters
	DRBQoS                                       *QoSFlowLevelQoSParameters
	QoSFlowsMappedtoDRBSetupResponseSNterminated *QoSFlowsMappedtoDRBSetupResponseSNterminated
	IEExtensions                                 []byte
}

func (ie *DRBsToBeModifiedListModificationResponseSNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBsToBeModifiedListModificationResponseSNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SNULPDCPUPTNLInfo != nil, ie.DRBQoS != nil, ie.QoSFlowsMappedtoDRBSetupResponseSNterminated != nil, false}); err != nil {
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
	if ie.QoSFlowsMappedtoDRBSetupResponseSNterminated != nil {
		if err := ie.QoSFlowsMappedtoDRBSetupResponseSNterminated.Encode(e); err != nil {
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

func (ie *DRBsToBeModifiedListModificationResponseSNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBsToBeModifiedListModificationResponseSNterminatedItemConstraints)
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
		ie.QoSFlowsMappedtoDRBSetupResponseSNterminated = new(QoSFlowsMappedtoDRBSetupResponseSNterminated)
		if err := ie.QoSFlowsMappedtoDRBSetupResponseSNterminated.Decode(d); err != nil {
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
