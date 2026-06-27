package ies

import (
	"github.com/lvdund/asn1go/per"
)

var dataForwardingResponseDRBItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "dlForwardingUPTNL", Optional: true},
		{Name: "ulForwardingUPTNL", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type DataForwardingResponseDRBItem struct {
	DrbID             DRBID
	DlForwardingUPTNL *UPTransportLayerInformation
	UlForwardingUPTNL *UPTransportLayerInformation
	IEExtensions      []byte
}

func (ie *DataForwardingResponseDRBItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dataForwardingResponseDRBItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DlForwardingUPTNL != nil, ie.UlForwardingUPTNL != nil, false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if ie.DlForwardingUPTNL != nil {
		if err := ie.DlForwardingUPTNL.Encode(e); err != nil {
			return err
		}
	}
	if ie.UlForwardingUPTNL != nil {
		if err := ie.UlForwardingUPTNL.Encode(e); err != nil {
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

func (ie *DataForwardingResponseDRBItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dataForwardingResponseDRBItemConstraints)
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
		ie.DlForwardingUPTNL = new(UPTransportLayerInformation)
		if err := ie.DlForwardingUPTNL.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.UlForwardingUPTNL = new(UPTransportLayerInformation)
		if err := ie.UlForwardingUPTNL.Decode(d); err != nil {
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
