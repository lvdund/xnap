package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dRBsToBeModifiedListModRqdMNterminatedItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "sN-DL-SCG-UP-TNLInfo"},
		{Name: "secondary-SN-DL-SCG-UP-TNLInfo", Optional: true},
		{Name: "lCID", Optional: true},
		{Name: "rlc-status", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DRBsToBeModifiedListModRqdMNterminatedItem struct {
	DrbID                     DRBID
	SNDLSCGUPTNLInfo          UPTransportLayerInformation
	SecondarySNDLSCGUPTNLInfo *UPTransportLayerInformation
	LCID                      *LCID
	RlcStatus                 *RLCStatus
	IEExtensions              []byte
}

func (ie *DRBsToBeModifiedListModRqdMNterminatedItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRBsToBeModifiedListModRqdMNterminatedItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.SecondarySNDLSCGUPTNLInfo != nil, ie.LCID != nil, ie.RlcStatus != nil, false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if err := ie.SNDLSCGUPTNLInfo.Encode(e); err != nil {
		return err
	}
	if ie.SecondarySNDLSCGUPTNLInfo != nil {
		if err := ie.SecondarySNDLSCGUPTNLInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.LCID != nil {
		if err := ie.LCID.Encode(e); err != nil {
			return err
		}
	}
	if ie.RlcStatus != nil {
		if err := ie.RlcStatus.Encode(e); err != nil {
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

func (ie *DRBsToBeModifiedListModRqdMNterminatedItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRBsToBeModifiedListModRqdMNterminatedItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DrbID.Decode(d); err != nil {
		return err
	}
	if err := ie.SNDLSCGUPTNLInfo.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(2) {
		ie.SecondarySNDLSCGUPTNLInfo = new(UPTransportLayerInformation)
		if err := ie.SecondarySNDLSCGUPTNLInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.LCID = new(LCID)
		if err := ie.LCID.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.RlcStatus = new(RLCStatus)
		if err := ie.RlcStatus.Decode(d); err != nil {
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
