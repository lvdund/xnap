package ies

import (
	"github.com/lvdund/asn1go/per"
)

var sDTDRBsToBeSetupListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "drb-ID"},
		{Name: "uL-TNLInfo"},
		{Name: "dRB-RLC-Bearer-Configuration"},
		{Name: "dRB-QoS"},
		{Name: "rLC-Mode"},
		{Name: "s-nssai"},
		{Name: "pDCP-SNLength"},
		{Name: "flows-Mapped-To-DRB-List"},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type SDTDRBsToBeSetupListItem struct {
	DrbID                     DRBID
	ULTNLInfo                 UPTransportLayerInformation
	DRBRLCBearerConfiguration []byte
	DRBQoS                    QoSFlowLevelQoSParameters
	RLCMode                   RLCMode
	SNssai                    SNSSAI
	PDCPSNLength              PDCPSNLength
	FlowsMappedToDRBList      FlowsMappedToDRBList
	IEExtensions              []byte
}

func (ie *SDTDRBsToBeSetupListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDTDRBsToBeSetupListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{false}); err != nil {
		return err
	}
	if err := ie.DrbID.Encode(e); err != nil {
		return err
	}
	if err := ie.ULTNLInfo.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.DRBRLCBearerConfiguration, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if err := ie.DRBQoS.Encode(e); err != nil {
		return err
	}
	if err := ie.RLCMode.Encode(e); err != nil {
		return err
	}
	if err := ie.SNssai.Encode(e); err != nil {
		return err
	}
	if err := ie.PDCPSNLength.Encode(e); err != nil {
		return err
	}
	if err := ie.FlowsMappedToDRBList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SDTDRBsToBeSetupListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDTDRBsToBeSetupListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.DrbID.Decode(d); err != nil {
		return err
	}
	if err := ie.ULTNLInfo.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeOctetString(per.SizeConstraints{
			Extensible: false,
			Min:        nil,
			Max:        nil,
		})
		if err != nil {
			return err
		}
		ie.DRBRLCBearerConfiguration = val
	}
	if err := ie.DRBQoS.Decode(d); err != nil {
		return err
	}
	if err := ie.RLCMode.Decode(d); err != nil {
		return err
	}
	if err := ie.SNssai.Decode(d); err != nil {
		return err
	}
	if err := ie.PDCPSNLength.Decode(d); err != nil {
		return err
	}
	if err := ie.FlowsMappedToDRBList.Decode(d); err != nil {
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
