package ies

import (
	"github.com/lvdund/asn1go/per"
)

var xnUAddressInfoperPDUSessionItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSession-ID"},
		{Name: "dataForwardingInfoFromTargetNGRANnode", Optional: true},
		{Name: "pduSessionResourceSetupCompleteInfo-SNterm", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type XnUAddressInfoperPDUSessionItem struct {
	PduSessionID                              PDUSessionID
	DataForwardingInfoFromTargetNGRANnode     *DataForwardingInfoFromTargetNGRANnode
	PduSessionResourceSetupCompleteInfoSNterm *PDUSessionResourceBearerSetupCompleteInfoSNterminated
	IEExtensions                              []byte
}

func (ie *XnUAddressInfoperPDUSessionItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(xnUAddressInfoperPDUSessionItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DataForwardingInfoFromTargetNGRANnode != nil, ie.PduSessionResourceSetupCompleteInfoSNterm != nil, false}); err != nil {
		return err
	}
	if err := ie.PduSessionID.Encode(e); err != nil {
		return err
	}
	if ie.DataForwardingInfoFromTargetNGRANnode != nil {
		if err := ie.DataForwardingInfoFromTargetNGRANnode.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSessionResourceSetupCompleteInfoSNterm != nil {
		if err := ie.PduSessionResourceSetupCompleteInfoSNterm.Encode(e); err != nil {
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

func (ie *XnUAddressInfoperPDUSessionItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(xnUAddressInfoperPDUSessionItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionID.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.DataForwardingInfoFromTargetNGRANnode = new(DataForwardingInfoFromTargetNGRANnode)
		if err := ie.DataForwardingInfoFromTargetNGRANnode.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.PduSessionResourceSetupCompleteInfoSNterm = new(PDUSessionResourceBearerSetupCompleteInfoSNterminated)
		if err := ie.PduSessionResourceSetupCompleteInfoSNterm.Decode(d); err != nil {
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
