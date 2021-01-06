package pkg

import (
	"fmt"

	"github.com/emshop/ots/modules/const/db/biz"
	"github.com/emshop/ots/modules/const/enums"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask"
)

//ProductFlow 业务流程
type ProductFlow struct {

	//FlowID 流程编号
	FlowID int `json:"flow_id"`

	//PlID 产品线编号
	PlID int `json:"pl_id"`

	//FlowName 流程名称
	FlowName string `json:"flow_name"`

	//QueueName 队列名称
	QueueName string `json:"queue_name"`

	//ScanInterval 超时时长
	ScanInterval int `json:"scan_interval"`

	//Delay 延后处理时长
	Delay int `json:"delay"`

	//Timeout 超时时长
	Timeout int `json:"timeout"`

	//MaxCount 最大执行次数
	MaxCount int `json:"max_count"`
}

//NextFlow 处理后续任务
func (p *ProductFlow) NextFlow(obj interface{}, kv ...interface{}) error {
	input := types.NewXMap()
	input.SetValue(biz.FieldIDFlowID, p.FlowID)
	input.Append(kv)
	_, callback, err := qtask.Create(obj, p.FlowName, input, p.ScanInterval, p.QueueName)
	if err != nil {
		return err
	}
	return callback(obj)
}

//NextFlow 处理后续任务
func NextFlow(plid int, s enums.FlowStatus, ctx interface{}, kv ...interface{}) error {
	flow, err := GetFirstFlow(plid, s)
	if err != nil || flow == nil {
		return err
	}
	return flow.NextFlow(ctx, kv...)
}

//GetFirstFlow 获取产品线的首要流程
func GetFirstFlow(plid int, s enums.FlowStatus) (*ProductFlow, error) {
	return GetFlow(plid, 0, s)
}

//GetFlow 根据产品线，父母流程编号、本次处理结果获取后续处理流程配置
func GetFlow(plid int, parentFlowID int, s enums.FlowStatus) (p *ProductFlow, err error) {

	//根据产品线父级流程编号获取流程配置
	rows, err := hydra.C.DB().GetRegularDB().Query(biz.SelectProductFlowByPlID, map[string]interface{}{
		biz.FieldTradeOrderPlID:     plid,
		biz.FieldFlowIDParentFlowID: parentFlowID,
	})
	if err != nil {
		return nil, err
	}
	var flowID = "0"
	switch s {
	case enums.Success:
		flowID = rows.Get(0).GetString(biz.FieldFlowIDSuccessFlowID)
	case enums.Failed:
		flowID = rows.Get(0).GetString(biz.FieldFlowIDFailedFlowID)
	case enums.Unknown:
		flowID = rows.Get(0).GetString(biz.FieldFlowIDUnknownFlowID)
	default:
		return nil, fmt.Errorf("产品线(%d)对应的流程未配置", plid)
	}

	//无后续处理流程
	if flowID == "0" {
		return nil, nil
	}

	//查询对应的后续处理流程
	rows, err = hydra.C.DB().GetRegularDB().Query(biz.SelectProductFlowByFlowID, map[string]interface{}{
		biz.FieldIDFlowID: flowID,
	})
	if err != nil {
		return nil, err
	}
	if rows.Len() == 0 {
		return nil, fmt.Errorf("产品线(%d)的后续处理流程未配置", plid)
	}
	err = rows.Get(0).ToAnyStruct(p)
	return p, err
}
