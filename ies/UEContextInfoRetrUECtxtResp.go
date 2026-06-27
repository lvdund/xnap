package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uEContextInfoRetrUECtxtRespConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ng-c-UE-signalling-ref"},
		{Name: "signalling-TNL-at-source"},
		{Name: "ueSecurityCapabilities"},
		{Name: "securityInformation"},
		{Name: "ue-AMBR"},
		{Name: "pduSessionResourcesToBeSetup-List"},
		{Name: "rrc-Context"},
		{Name: "mobilityRestrictionList", Optional: true},
		{Name: "indexToRatFrequencySelectionPriority", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UEContextInfoRetrUECtxtResp struct {
	NgCUESignallingRef                   AMFUENGAPID
	SignallingTNLAtSource                CPTransportLayerInformation
	UeSecurityCapabilities               UESecurityCapabilities
	SecurityInformation                  ASSecurityInformation
	UeAMBR                               UEAggregateMaximumBitRate
	PduSessionResourcesToBeSetupList     PDUSessionResourcesToBeSetupList
	RrcContext                           []byte
	MobilityRestrictionList              *MobilityRestrictionList
	IndexToRatFrequencySelectionPriority *RFSPIndex
	IEExtensions                         []byte
}

func (ie *UEContextInfoRetrUECtxtResp) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEContextInfoRetrUECtxtRespConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MobilityRestrictionList != nil, ie.IndexToRatFrequencySelectionPriority != nil, false}); err != nil {
		return err
	}
	if err := ie.NgCUESignallingRef.Encode(e); err != nil {
		return err
	}
	if err := ie.SignallingTNLAtSource.Encode(e); err != nil {
		return err
	}
	if err := ie.UeSecurityCapabilities.Encode(e); err != nil {
		return err
	}
	if err := ie.SecurityInformation.Encode(e); err != nil {
		return err
	}
	if err := ie.UeAMBR.Encode(e); err != nil {
		return err
	}
	if err := ie.PduSessionResourcesToBeSetupList.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeOctetString(ie.RrcContext, per.SizeConstraints{
		Extensible: false,
		Min:        nil,
		Max:        nil,
	}); err != nil {
		return err
	}
	if ie.MobilityRestrictionList != nil {
		if err := ie.MobilityRestrictionList.Encode(e); err != nil {
			return err
		}
	}
	if ie.IndexToRatFrequencySelectionPriority != nil {
		if err := ie.IndexToRatFrequencySelectionPriority.Encode(e); err != nil {
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

func (ie *UEContextInfoRetrUECtxtResp) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEContextInfoRetrUECtxtRespConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NgCUESignallingRef.Decode(d); err != nil {
		return err
	}
	if err := ie.SignallingTNLAtSource.Decode(d); err != nil {
		return err
	}
	if err := ie.UeSecurityCapabilities.Decode(d); err != nil {
		return err
	}
	if err := ie.SecurityInformation.Decode(d); err != nil {
		return err
	}
	if err := ie.UeAMBR.Decode(d); err != nil {
		return err
	}
	if err := ie.PduSessionResourcesToBeSetupList.Decode(d); err != nil {
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
		ie.RrcContext = val
	}
	if seq.IsComponentPresent(7) {
		ie.MobilityRestrictionList = new(MobilityRestrictionList)
		if err := ie.MobilityRestrictionList.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(8) {
		ie.IndexToRatFrequencySelectionPriority = new(RFSPIndex)
		if err := ie.IndexToRatFrequencySelectionPriority.Decode(d); err != nil {
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
