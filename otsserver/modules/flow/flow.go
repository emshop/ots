package flow

import (
	"fmt"
	"net/http"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/modules/const/sql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
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

//Next 处理后续任务
func (p *ProductFlow) Next(obj interface{}, kv ...interface{}) (types.XMap, error) {
	if p.FlowID == 0 {
		return nil, errs.NewError(http.StatusNoContent, "流程未配置")
	}
	input := types.NewXMap()
	input.SetValue(sql.FieldFlowID, p.FlowID)
	input.Append(kv...)
	_, callback, err := qtask.Create(obj, p.FlowName, input, p.ScanInterval, p.QueueName)
	if err != nil {
		return nil, err
	}
	return input, callback(obj)
}

//NextByFirst 处理后续任务
func NextByFirst(tagName enums.FlowTag, plid int, s enums.FlowStatus, ctx interface{}, kv ...interface{}) (types.XMap, error) {
	flow, err := GetFirst(tagName, plid, s)
	if err != nil || flow == nil {
		return nil, err
	}
	return flow.Next(ctx, kv...)
}

//Next 处理后续任务
func Next(flowID int, s enums.FlowStatus, ctx interface{}, kv ...interface{}) (types.XMap, error) {
	flow, err := Get(flowID, s)
	if err != nil || flow == nil {
		return nil, err
	}
	return flow.Next(ctx, kv...)
}

//GetFirst 获取产品线的首要流程
func GetFirst(tagName enums.FlowTag, plid int, s enums.FlowStatus) (*ProductFlow, error) {
	//根据产品线父级流程编号获取流程配置
	rows, err := hydra.C.DB().GetRegularDB().Query(sql.SelectFlowByTag, map[string]interface{}{
		sql.FieldPlID:    plid,
		sql.FieldTagName: tagName,
	})
	if err != nil {
		return nil, err
	}
	return getChildFlow(rows.Get(0), s)
}

//Get 根据产品线，父母流程编号、本次处理结果获取后续处理流程配置
func Get(flowID int, s enums.FlowStatus) (p *ProductFlow, err error) {
	//根据产品线父级流程编号获取流程配置
	rows, err := hydra.C.DB().GetRegularDB().Query(sql.SelectChildFlowByFlowID, map[string]interface{}{
		sql.FieldFlowID: flowID,
	})
	if err != nil {
		return nil, err
	}
	return getChildFlow(rows.Get(0), s)
}

func getChildFlow(flow types.XMap, s enums.FlowStatus) (p *ProductFlow, err error) {
	if flow.Len() == 0 {
		return &ProductFlow{}, nil
	}

	var flowID = "0"
	switch s {
	case enums.Success:
		flowID = flow.GetString(sql.FieldSuccessFlowID)
	case enums.Failed:
		flowID = flow.GetString(sql.FieldFailedFlowID)
	case enums.Unknown:
		flowID = flow.GetString(sql.FieldUnknownFlowID)
	}

	//无后续处理流程
	if flowID == "0" {
		return &ProductFlow{}, nil
	}

	//查询对应的后续处理流程
	rows, err := hydra.C.DB().GetRegularDB().Query(sql.SelectProductFlowByFlowID, map[string]interface{}{
		sql.FieldFlowID: flowID,
		sql.FieldPlID:   flow.GetString(sql.FieldPlID),
	})
	if err != nil {
		return nil, err
	}
	if rows.Len() == 0 {
		return nil, fmt.Errorf("---产品线(%s)的后续处理流程(%s)未配置", flow.GetString(sql.FieldPlID), flowID)
	}
	p = &ProductFlow{}
	err = rows.Get(0).ToAnyStruct(p)
	return p, err
}
