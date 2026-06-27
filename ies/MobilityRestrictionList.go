package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var mobilityRestrictionListConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "serving-PLMN"},
		{Name: "equivalent-PLMNs", Optional: true},
		{Name: "rat-Restrictions", Optional: true},
		{Name: "forbiddenAreaInformation", Optional: true},
		{Name: "serviceAreaInformation", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type MobilityRestrictionList struct {
	ServingPLMN              PLMNIdentity
	EquivalentPLMNs          []*PLMNIdentity
	RatRestrictions          *RATRestrictionsList
	ForbiddenAreaInformation *ForbiddenAreaList
	ServiceAreaInformation   *ServiceAreaList
	IEExtensions             []byte
}

func (ie *MobilityRestrictionList) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mobilityRestrictionListConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{len(ie.EquivalentPLMNs) > 0, ie.RatRestrictions != nil, ie.ForbiddenAreaInformation != nil, ie.ServiceAreaInformation != nil, false}); err != nil {
		return err
	}
	if err := ie.ServingPLMN.Encode(e); err != nil {
		return err
	}
	if len(ie.EquivalentPLMNs) > 0 {
		soEnc := e.NewSequenceOfEncoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofEPLMNs)),
		})
		if err := soEnc.EncodeLength(int64(len(ie.EquivalentPLMNs))); err != nil {
			return err
		}
		for _, item := range ie.EquivalentPLMNs {
			if err := item.Encode(e); err != nil {
				return err
			}
		}
	}
	if ie.RatRestrictions != nil {
		if err := ie.RatRestrictions.Encode(e); err != nil {
			return err
		}
	}
	if ie.ForbiddenAreaInformation != nil {
		if err := ie.ForbiddenAreaInformation.Encode(e); err != nil {
			return err
		}
	}
	if ie.ServiceAreaInformation != nil {
		if err := ie.ServiceAreaInformation.Encode(e); err != nil {
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

func (ie *MobilityRestrictionList) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mobilityRestrictionListConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.ServingPLMN.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		soDec := d.NewSequenceOfDecoder(per.SizeConstraints{
			Extensible: false,
			Min:        common.Ptr(int64(1)),
			Max:        common.Ptr(int64(common.MaxnoofEPLMNs)),
		})
		n, err := soDec.DecodeLength()
		if err != nil {
			return err
		}
		ie.EquivalentPLMNs = make([]*PLMNIdentity, n)
		for i := range ie.EquivalentPLMNs {
			ie.EquivalentPLMNs[i] = new(PLMNIdentity)
			if err := ie.EquivalentPLMNs[i].Decode(d); err != nil {
				return err
			}
		}
	}
	if seq.IsComponentPresent(2) {
		ie.RatRestrictions = new(RATRestrictionsList)
		if err := ie.RatRestrictions.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.ForbiddenAreaInformation = new(ForbiddenAreaList)
		if err := ie.ForbiddenAreaInformation.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.ServiceAreaInformation = new(ServiceAreaList)
		if err := ie.ServiceAreaInformation.Decode(d); err != nil {
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
