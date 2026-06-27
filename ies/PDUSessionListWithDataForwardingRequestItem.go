package ies

import (
	"github.com/lvdund/asn1go/per"
)

var pDUSessionListWithDataForwardingRequestItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pduSessionId"},
		{Name: "dataforwardingInfofromSource", Optional: true},
		{Name: "dRBtoBeReleasedList", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type PDUSessionListWithDataForwardingRequestItem struct {
	PduSessionId                 PDUSessionID
	DataforwardingInfofromSource *DataforwardingandOffloadingInfofromSource
	DRBtoBeReleasedList          *DRBToQoSFlowMappingList
	IEExtensions                 []byte
}

func (ie *PDUSessionListWithDataForwardingRequestItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDUSessionListWithDataForwardingRequestItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.DataforwardingInfofromSource != nil, ie.DRBtoBeReleasedList != nil, false}); err != nil {
		return err
	}
	if err := ie.PduSessionId.Encode(e); err != nil {
		return err
	}
	if ie.DataforwardingInfofromSource != nil {
		if err := ie.DataforwardingInfofromSource.Encode(e); err != nil {
			return err
		}
	}
	if ie.DRBtoBeReleasedList != nil {
		if err := ie.DRBtoBeReleasedList.Encode(e); err != nil {
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

func (ie *PDUSessionListWithDataForwardingRequestItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDUSessionListWithDataForwardingRequestItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.PduSessionId.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.DataforwardingInfofromSource = new(DataforwardingandOffloadingInfofromSource)
		if err := ie.DataforwardingInfofromSource.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.DRBtoBeReleasedList = new(DRBToQoSFlowMappingList)
		if err := ie.DRBtoBeReleasedList.Decode(d); err != nil {
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
