package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uEContextInfoSNModRequestConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ueSecurityCapabilities", Optional: true},
		{Name: "s-ng-RANnode-SecurityKey", Optional: true},
		{Name: "s-ng-RANnodeUE-AMBR", Optional: true},
		{Name: "indexToRatFrequencySelectionPriority", Optional: true},
		{Name: "lowerLayerPresenceStatusChange", Optional: true},
		{Name: "pduSessionResourceToBeAdded", Optional: true},
		{Name: "pduSessionResourceToBeModified", Optional: true},
		{Name: "pduSessionResourceToBeReleased", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UEContextInfoSNModRequest struct {
	UeSecurityCapabilities               *UESecurityCapabilities
	SNgRANnodeSecurityKey                *SNGRANnodeSecurityKey
	SNgRANnodeUEAMBR                     *UEAggregateMaximumBitRate
	IndexToRatFrequencySelectionPriority *RFSPIndex
	LowerLayerPresenceStatusChange       *LowerLayerPresenceStatusChange
	PduSessionResourceToBeAdded          *PDUSessionsToBeAddedSNModRequestList
	PduSessionResourceToBeModified       *PDUSessionsToBeModifiedSNModRequestList
	PduSessionResourceToBeReleased       *PDUSessionsToBeReleasedSNModRequestList
	IEExtensions                         []byte
}

func (ie *UEContextInfoSNModRequest) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEContextInfoSNModRequestConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.UeSecurityCapabilities != nil, ie.SNgRANnodeSecurityKey != nil, ie.SNgRANnodeUEAMBR != nil, ie.IndexToRatFrequencySelectionPriority != nil, ie.LowerLayerPresenceStatusChange != nil, ie.PduSessionResourceToBeAdded != nil, ie.PduSessionResourceToBeModified != nil, ie.PduSessionResourceToBeReleased != nil, false}); err != nil {
		return err
	}
	if ie.UeSecurityCapabilities != nil {
		if err := ie.UeSecurityCapabilities.Encode(e); err != nil {
			return err
		}
	}
	if ie.SNgRANnodeSecurityKey != nil {
		if err := ie.SNgRANnodeSecurityKey.Encode(e); err != nil {
			return err
		}
	}
	if ie.SNgRANnodeUEAMBR != nil {
		if err := ie.SNgRANnodeUEAMBR.Encode(e); err != nil {
			return err
		}
	}
	if ie.IndexToRatFrequencySelectionPriority != nil {
		if err := ie.IndexToRatFrequencySelectionPriority.Encode(e); err != nil {
			return err
		}
	}
	if ie.LowerLayerPresenceStatusChange != nil {
		if err := ie.LowerLayerPresenceStatusChange.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSessionResourceToBeAdded != nil {
		if err := ie.PduSessionResourceToBeAdded.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSessionResourceToBeModified != nil {
		if err := ie.PduSessionResourceToBeModified.Encode(e); err != nil {
			return err
		}
	}
	if ie.PduSessionResourceToBeReleased != nil {
		if err := ie.PduSessionResourceToBeReleased.Encode(e); err != nil {
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

func (ie *UEContextInfoSNModRequest) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEContextInfoSNModRequestConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if seq.IsComponentPresent(0) {
		ie.UeSecurityCapabilities = new(UESecurityCapabilities)
		if err := ie.UeSecurityCapabilities.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(1) {
		ie.SNgRANnodeSecurityKey = new(SNGRANnodeSecurityKey)
		if err := ie.SNgRANnodeSecurityKey.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.SNgRANnodeUEAMBR = new(UEAggregateMaximumBitRate)
		if err := ie.SNgRANnodeUEAMBR.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.IndexToRatFrequencySelectionPriority = new(RFSPIndex)
		if err := ie.IndexToRatFrequencySelectionPriority.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.LowerLayerPresenceStatusChange = new(LowerLayerPresenceStatusChange)
		if err := ie.LowerLayerPresenceStatusChange.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.PduSessionResourceToBeAdded = new(PDUSessionsToBeAddedSNModRequestList)
		if err := ie.PduSessionResourceToBeAdded.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.PduSessionResourceToBeModified = new(PDUSessionsToBeModifiedSNModRequestList)
		if err := ie.PduSessionResourceToBeModified.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.PduSessionResourceToBeReleased = new(PDUSessionsToBeReleasedSNModRequestList)
		if err := ie.PduSessionResourceToBeReleased.Decode(d); err != nil {
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
