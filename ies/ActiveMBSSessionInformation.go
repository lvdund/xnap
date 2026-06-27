package ies

import (
	"github.com/lvdund/asn1go/per"
)

var activeMBSSessionInformationConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mBS-QoSFlowsToAdd-List"},
		{Name: "mBS-ServiceArea", Optional: true},
		{Name: "mBS-MappingandDataForwardingRequestInfofromSource", Optional: true},
		{Name: "iE-Extensions", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type ActiveMBSSessionInformation struct {
	MBSQoSFlowsToAddList                             MBSQoSFlowsToAddList
	MBSServiceArea                                   *MBSServiceArea
	MBSMappingandDataForwardingRequestInfofromSource *MBSMappingandDataForwardingRequestInfofromSource
	IEExtensions                                     []byte
}

func (ie *ActiveMBSSessionInformation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(activeMBSSessionInformationConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.MBSServiceArea != nil, ie.MBSMappingandDataForwardingRequestInfofromSource != nil, false}); err != nil {
		return err
	}
	if err := ie.MBSQoSFlowsToAddList.Encode(e); err != nil {
		return err
	}
	if ie.MBSServiceArea != nil {
		if err := ie.MBSServiceArea.Encode(e); err != nil {
			return err
		}
	}
	if ie.MBSMappingandDataForwardingRequestInfofromSource != nil {
		if err := ie.MBSMappingandDataForwardingRequestInfofromSource.Encode(e); err != nil {
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

func (ie *ActiveMBSSessionInformation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(activeMBSSessionInformationConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.MBSQoSFlowsToAddList.Decode(d); err != nil {
		return err
	}
	if seq.IsComponentPresent(1) {
		ie.MBSServiceArea = new(MBSServiceArea)
		if err := ie.MBSServiceArea.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(2) {
		ie.MBSMappingandDataForwardingRequestInfofromSource = new(MBSMappingandDataForwardingRequestInfofromSource)
		if err := ie.MBSMappingandDataForwardingRequestInfofromSource.Decode(d); err != nil {
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
