package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/xnap/common"
)

var coverageModificationListItemConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "globalNG-RANCell-ID"},
		{Name: "cellCoverageState"},
		{Name: "cellDeploymentStatusIndicator", Optional: true},
		{Name: "cellReplacingInfo", Optional: true},
		{Name: "sSB-Coverage-Modification-List"},
		{Name: "iE-Extension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "iE-Extensions"},
	},
}

type CoverageModificationListItem struct {
	GlobalNGRANCellID             GlobalCellID
	CellCoverageState             int64
	CellDeploymentStatusIndicator *CellDeploymentStatusIndicator
	CellReplacingInfo             *CellReplacingInfo
	SSBCoverageModificationList   SSBCoverageModificationList
	IEExtensions                  []byte
}

func (ie *CoverageModificationListItem) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(coverageModificationListItemConstraints)
	hasExt := len(ie.IEExtensions) > 0
	if err := seq.EncodeExtensionBit(hasExt); err != nil {
		return err
	}
	if err := seq.EncodePreamble([]bool{ie.CellDeploymentStatusIndicator != nil, ie.CellReplacingInfo != nil, false}); err != nil {
		return err
	}
	if err := ie.GlobalNGRANCellID.Encode(e); err != nil {
		return err
	}
	if err := e.EncodeInteger(ie.CellCoverageState, per.IntegerConstraints{
		Extensible: false,
		LowerBound: common.Ptr(int64(0)),
		UpperBound: common.Ptr(int64(63)),
	}); err != nil {
		return err
	}
	if ie.CellDeploymentStatusIndicator != nil {
		if err := ie.CellDeploymentStatusIndicator.Encode(e); err != nil {
			return err
		}
	}
	if ie.CellReplacingInfo != nil {
		if err := ie.CellReplacingInfo.Encode(e); err != nil {
			return err
		}
	}
	if err := ie.SSBCoverageModificationList.Encode(e); err != nil {
		return err
	}
	if hasExt {
		if err := seq.EncodeExtensionAdditions([]bool{true}, [][]byte{ie.IEExtensions}); err != nil {
			return err
		}
	}
	return nil
}

func (ie *CoverageModificationListItem) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(coverageModificationListItemConstraints)
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}
	if err := seq.DecodePreamble(); err != nil {
		return err
	}
	if err := ie.GlobalNGRANCellID.Decode(d); err != nil {
		return err
	}
	{
		val, err := d.DecodeInteger(per.IntegerConstraints{
			Extensible: false,
			LowerBound: common.Ptr(int64(0)),
			UpperBound: common.Ptr(int64(63)),
		})
		if err != nil {
			return err
		}
		ie.CellCoverageState = val
	}
	if seq.IsComponentPresent(2) {
		ie.CellDeploymentStatusIndicator = new(CellDeploymentStatusIndicator)
		if err := ie.CellDeploymentStatusIndicator.Decode(d); err != nil {
			return err
		}
	}
	if seq.IsComponentPresent(3) {
		ie.CellReplacingInfo = new(CellReplacingInfo)
		if err := ie.CellReplacingInfo.Decode(d); err != nil {
			return err
		}
	}
	if err := ie.SSBCoverageModificationList.Decode(d); err != nil {
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
