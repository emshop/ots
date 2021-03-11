package flows

import (
	"github.com/emshop/ots/flowserver/modules/const/enums"
	"github.com/emshop/ots/flowserver/modules/const/fields"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask"
)

//ProductFlow 业务流程
type ProductFlow struct {

	//FlowID 流程编号
	FlowID int `json:"flow_id"`

	//FlowTag 流程名称
	FlowTag string `json:"flow_tag"`

	//PlID 产品线
	PlID int `json:"pl_id"`

	//QueueName 队列名称
	QueueName string `json:"queue_name"`

	//ScanInterval 执行间隔
	ScanInterval int `json:"scan_interval"`

	//Delay 延后时长
	Delay int `json:"delay"`

	//Timeout 超时时长
	Timeout int `json:"timeout"`

	//MaxCount 最大次数
	MaxCount int `json:"max_count"`
}

//Next 处理后续任务
func (p *ProductFlow) Next(obj interface{}, input types.IXMap) (types.IXMap, error) {

	p.Timeout = types.DecodeInt(p.Timeout, 0, 86400)
	if p.Delay == 0 {
		_, callback, err := qtask.Create(obj, p.FlowTag, input.ToMap(), p.ScanInterval, p.QueueName,
			qtask.WithOrderNO(input.GetString(fields.FieldOrderID)), qtask.WithDeadline(p.Timeout))
		if err != nil {
			return nil, err
		}
		return input, callback(obj)
	}
	_, err := qtask.Delay(obj, p.FlowTag, input.ToMap(), p.Delay, p.ScanInterval, p.QueueName,
		qtask.WithOrderNO(input.GetString(fields.FieldOrderID)), qtask.WithDeadline(p.Timeout))
	if err != nil {
		return nil, err
	}
	return input, nil
}

//NextByOrderNO 根据订单号处理后续流程
func NextByOrderNO(orderNo string, nextFlowTag enums.FlowName, ctx hydra.IContext, kv ...interface{}) types.IXMap {
	data, err := hydra.C.DB().GetRegularDB().Query(SelectTradeOrderByOrderID, map[string]interface{}{
		fields.FieldOrderID: orderNo,
	})
	if err != nil {
		ctx.Log().Errorf("获取订单信息失败:%v", err)
		return nil
	}
	kv = append(kv, fields.FieldOrderID)
	kv = append(kv, orderNo)
	return Next(nextFlowTag, ctx, data.Get(0), kv...)
}

//NextByDeliveryID 根据发货编号处理后续流程
func NextByDeliveryID(deliveryID string, nextFlowTag enums.FlowName, ctx hydra.IContext, kv ...interface{}) types.IXMap {
	data, err := hydra.C.DB().GetRegularDB().Query(selectDeliveryByDeliveryID, map[string]interface{}{
		fields.FieldDeliveryID: deliveryID,
	})
	if err != nil {
		ctx.Log().Errorf("获取订单信息失败:%v", err)
		return nil
	}
	kv = append(kv, fields.FieldDeliveryID)
	kv = append(kv, deliveryID)
	kv = append(kv, fields.FieldOrderID)
	kv = append(kv, data.Get(0).GetString(fields.FieldOrderID))
	return Next(nextFlowTag, ctx, data.Get(0), kv...)
}

//Next 处理后续任务
func Next(nextFlowTag enums.FlowName, ctx hydra.IContext, input types.IXMap, kv ...interface{}) types.IXMap {
	input.SetValue(fields.FieldFlowTag, nextFlowTag)

	//查询对应的流程配置
	flows, err := hydra.C.DB().GetRegularDB().Query(SelectProductFlowByTag, input.ToMap())
	if err != nil {
		ctx.Log().Errorf("获取处理流程(%s)失败:%v", nextFlowTag, err)
		return nil
	}
	if flows.Len() == 0 {
		ctx.Log().Warnf("未查询到流程:%s", nextFlowTag)
		return nil
	}

	//执行流程处理
	flow := &ProductFlow{}
	err = flows.Get(0).ToAnyStruct(flow)
	if err != nil {
		ctx.Log().Errorf("转换后续流程(%s)失败:%v", nextFlowTag, err)
		return nil
	}
	flows.Get(0).Merge(input)
	flow.QueueName = flows.Get(0).Translate(flow.QueueName)

	ninput := types.NewXMap()
	ninput.Append(kv...)

	mp, err := flow.Next(ctx, ninput)
	if err != nil {
		ctx.Log().Errorf("发送到QTASK失败:%v", err)
		return nil
	}
	return mp
}
