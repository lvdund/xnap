package ies

import (
	"github.com/lvdund/asn1go/per"
)

var uEAppLayerMeasConfigInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qOEReference"},
		{Name: "qOEMeasConfigAppLayerID", Optional: true},
		{Name: "serviceType"},
		{Name: "qOEMeasStatus", Optional: true},
		{Name: "containerAppLayerMeasConfig", Optional: true},
		{Name: "mDTAlignmentInfo", Optional: true},
		{Name: "measCollectionEntityIPAddress", Optional: true},
		{Name: "areaScopeOfQMC", Optional: true},
		{Name: "s-NSSAIListQoE", Optional: true},
		{Name: "availableRVQoEMetrics", Optional: true},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type UEAppLayerMeasConfigInfo struct {
	QOEReference                  QOEReference
	QOEMeasConfigAppLayerID       *QOEMeasConfAppLayerID
	ServiceType                   ServiceType
	QOEMeasStatus                 *QOEMeasStatus
	ContainerAppLayerMeasConfig   *ContainerAppLayerMeasConfig
	MDTAlignmentInfo              *MDTAlignmentInfo
	MeasCollectionEntityIPAddress *MeasCollectionEntityIPAddress
	AreaScopeOfQMC                *AreaScopeOfQMC
	SNSSAIListQoE                 *SNSSAIListQoE
	AvailableRVQoEMetrics         *AvailableRVQoEMetrics
	IEExtensions                  []byte
}

func (ie *UEAppLayerMeasConfigInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEAppLayerMeasConfigInfoConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.QOEMeasConfigAppLayerID != nil, ie.QOEMeasStatus != nil, ie.ContainerAppLayerMeasConfig != nil, ie.MDTAlignmentInfo != nil, ie.MeasCollectionEntityIPAddress != nil, ie.AreaScopeOfQMC != nil, ie.SNSSAIListQoE != nil, ie.AvailableRVQoEMetrics != nil, false}); err != nil {
		return err
	}
	if err := ie.QOEReference.Encode(e); err != nil {
		return err
	}
	if ie.QOEMeasConfigAppLayerID != nil {
		if err := ie.QOEMeasConfigAppLayerID.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.ServiceType.Encode(e); err != nil {
		return err
	}
	if ie.QOEMeasStatus != nil {
		if err := ie.QOEMeasStatus.Encode(e); err != nil {
			return err
		}
	}
	if ie.ContainerAppLayerMeasConfig != nil {
		if err := ie.ContainerAppLayerMeasConfig.Encode(e); err != nil {
			return err
		}
	}
	if ie.MDTAlignmentInfo != nil {
		if err := ie.MDTAlignmentInfo.Encode(e); err != nil {
			return err
		}
	}
	if ie.MeasCollectionEntityIPAddress != nil {
		if err := ie.MeasCollectionEntityIPAddress.Encode(e); err != nil {
			return err
		}
	}
	if ie.AreaScopeOfQMC != nil {
		if err := ie.AreaScopeOfQMC.Encode(e); err != nil {
			return err
		}
	}
	if ie.SNSSAIListQoE != nil {
		if err := ie.SNSSAIListQoE.Encode(e); err != nil {
			return err
		}
	}
	if ie.AvailableRVQoEMetrics != nil {
		if err := ie.AvailableRVQoEMetrics.Encode(e); err != nil {
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

func (ie *UEAppLayerMeasConfigInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEAppLayerMeasConfigInfoConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.QOEReference.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.QOEMeasConfigAppLayerID = new(QOEMeasConfAppLayerID)
		if err := ie.QOEMeasConfigAppLayerID.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.ServiceType.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(3) {
		ie.QOEMeasStatus = new(QOEMeasStatus)
		if err := ie.QOEMeasStatus.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(4) {
		ie.ContainerAppLayerMeasConfig = new(ContainerAppLayerMeasConfig)
		if err := ie.ContainerAppLayerMeasConfig.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(5) {
		ie.MDTAlignmentInfo = new(MDTAlignmentInfo)
		if err := ie.MDTAlignmentInfo.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(6) {
		ie.MeasCollectionEntityIPAddress = new(MeasCollectionEntityIPAddress)
		if err := ie.MeasCollectionEntityIPAddress.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(7) {
		ie.AreaScopeOfQMC = new(AreaScopeOfQMC)
		if err := ie.AreaScopeOfQMC.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(8) {
		ie.SNSSAIListQoE = new(SNSSAIListQoE)
		if err := ie.SNSSAIListQoE.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(9) {
		ie.AvailableRVQoEMetrics = new(AvailableRVQoEMetrics)
		if err := ie.AvailableRVQoEMetrics.Decode(d); err != nil {
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
