package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uEContextInfoHORequestConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ng-c-UE-reference"},
		{Name: "cp-TNL-info-source"},
		{Name: "ueSecurityCapabilities"},
		{Name: "securityInformation"},
		{Name: "indexToRatFrequencySelectionPriority", Optional: true},
		{Name: "ue-AMBR"},
		{Name: "pduSessionResourcesToBeSetup-List"},
		{Name: "rrc-Context"},
		{Name: "locationReportingInformation", Optional: true},
		{Name: "mrl", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UEContextInfoHORequest struct {
	NgCUEReference                       AMFUENGAPID
	CpTNLInfoSource                      CPTransportLayerInformation
	UeSecurityCapabilities               UESecurityCapabilities
	SecurityInformation                  ASSecurityInformation
	IndexToRatFrequencySelectionPriority *RFSPIndex
	UeAMBR                               UEAggregateMaximumBitRate
	PduSessionResourcesToBeSetupList     PDUSessionResourcesToBeSetupList
	RrcContext                           []byte
	LocationReportingInformation         *LocationReportingInformation
	Mrl                                  *MobilityRestrictionList
	IEExtensions                         []byte
}

func (ie *UEContextInfoHORequest) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEContextInfoHORequestConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.IndexToRatFrequencySelectionPriority != nil, ie.LocationReportingInformation != nil, ie.Mrl != nil, false}); err != nil {
		return err
	}
	if err := ie.NgCUEReference.Encode(e); err != nil {
		return err
	}
	if err := ie.CpTNLInfoSource.Encode(e); err != nil {
		return err
	}
	if err := ie.UeSecurityCapabilities.Encode(e); err != nil {
		return err
	}
	if err := ie.SecurityInformation.Encode(e); err != nil {
		return err
	}
	if ie.IndexToRatFrequencySelectionPriority != nil {
		if err := ie.IndexToRatFrequencySelectionPriority.Encode(e); err != nil {
			return err
		}
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
	if ie.LocationReportingInformation != nil {
		if err := ie.LocationReportingInformation.Encode(e); err != nil {
			return err
		}
	}
	if ie.Mrl != nil {
		if err := ie.Mrl.Encode(e); err != nil {
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

func (ie *UEContextInfoHORequest) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEContextInfoHORequestConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.NgCUEReference.Decode(d); err != nil {
		return err
	}
	if err := ie.CpTNLInfoSource.Decode(d); err != nil {
		return err
	}
	if err := ie.UeSecurityCapabilities.Decode(d); err != nil {
		return err
	}
	if err := ie.SecurityInformation.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(4) {
		ie.IndexToRatFrequencySelectionPriority = new(RFSPIndex)
		if err := ie.IndexToRatFrequencySelectionPriority.Decode(d); err != nil {
			return err
		}
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
	if seq.IsComponentPresent(8) {
		ie.LocationReportingInformation = new(LocationReportingInformation)
		if err := ie.LocationReportingInformation.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(9) {
		ie.Mrl = new(MobilityRestrictionList)
		if err := ie.Mrl.Decode(d); err != nil {
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
