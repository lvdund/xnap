package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sDTDataForwardingDRBListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "dL-TNLInfo", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SDTDataForwardingDRBListItem struct {
	DrbID        DRBID
	DLTNLInfo    *UPTransportLayerInformation
	IEExtensions []byte
}

func (ie *SDTDataForwardingDRBListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDTDataForwardingDRBListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DLTNLInfo != nil, false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if ie.DLTNLInfo != nil {
		if err := ie.DLTNLInfo.Encode(e); err != nil {
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

func (ie *SDTDataForwardingDRBListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDTDataForwardingDRBListItemConstraints)
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
		ie.DLTNLInfo = new(UPTransportLayerInformation)
		if err := ie.DLTNLInfo.Decode(d); err != nil {
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
