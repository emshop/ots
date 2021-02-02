package ut

import (
	"fmt"

	"github.com/emshop/ots/otsserver/modules/const/enums"
	"github.com/emshop/ots/otsserver/services/bind"
	"github.com/emshop/ots/otsserver/services/delivery"
	"github.com/emshop/ots/otsserver/services/notify"
	"github.com/emshop/ots/otsserver/services/order"
	"github.com/emshop/ots/otsserver/services/payment"
	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/hydra/mock"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/logger"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/lib4go/utility"
)

var orderCache = &testCache{
	order:             make(chan types.XMap, 1000),
	paidOrder:         make(chan types.XMap, 1000),
	bindedOrder:       make(chan types.XMap, 1000),
	deliveriededOrder: make(chan types.XMap, 1000),
	notifyOrder:       make(chan types.XMap, 1000),
}

type testCache struct {
	order             chan types.XMap
	paidOrder         chan types.XMap
	bindedOrder       chan types.XMap
	deliveriededOrder chan types.XMap
	notifyOrder       chan types.XMap
}

func execteSQL(sql string) error {
	_, err := hydra.C.DB().GetRegularDB().Execute(sql, map[string]interface{}{})
	return err
}

func (t *testCache) makeOrder(count int) error {
	//构建context
	for i := 0; i < count; i++ {
		input := fmt.Sprintf(`mer_no=qx001&mer_order_no=%s&mer_product_id=10000&num=1&account_name=15828680877`, utility.GetGUID())
		ctx := mock.NewContext(input)

		//构建请求处理------------------------
		order := &order.Order{}
		rs := order.RequestHandle(ctx)
		err := ctx.Response().WriteAny(getAny(rs))
		if err != nil {
			return err
		}

		status, _, _ := ctx.Response().GetFinalResponse()
		if status != 200 {
			return fmt.Errorf("订单创建失败:%v", rs)
		}
		t.order <- ctx.Meta().GetValue("flow").(types.XMaps).Get(0)
	}
	return nil
}

func (t *testCache) makePaidOrder(count int) error {
	if err := t.makeOrder(count); err != nil {
		return err
	}
	for {
		if len(t.order) == 0 {
			return nil
		}
		select {
		case order := <-t.order:
			//构建context
			ctx := mock.NewContext(string(order.Marshal()))
			payment := &payment.Payment{}
			ps := payment.PayHandle(ctx)
			ctx.Response().WriteAny(getAny(ps))
			status, _, _ := ctx.Response().GetFinalResponse()
			if status != 200 {
				return fmt.Errorf("订单支付失败:%v", ps)
			}
			t.paidOrder <- ps.(types.XMaps).Get(0)
		}
	}
}
func (t *testCache) makeDeliveryOrder(count int) error {
	err := orderCache.makeBindedOrder(1)
	if err != nil {
		return err
	}

	//开始发货--------------------
	order := <-orderCache.bindedOrder
	input := string(order.Marshal())
	ctx := mock.NewContext(input)

	var dlv = &delivery.Delivery{}
	rs := dlv.StartHandle(ctx)

	if err := errs.GetError(rs); err != nil {
		return err
	}

	//保存请求结果---------------------------
	ctx.Request().SetValue("result_code", "000")
	ctx.Request().SetValue("return_msg", "上游请求成功")
	rs = dlv.SaveSartHandle(ctx)
	if err := errs.GetError(rs); err != nil {
		return err
	}

	//保存发货结果---------------------------
	ctx.Request().SetValue("result_code", "000")
	ctx.Request().SetValue("return_msg", "发货成功000")
	rs = dlv.SaveResultHandle(ctx)
	if err := errs.GetError(rs); err != nil {
		return err
	}
	fmt.Println("delivery.flow:", rs)
	t.deliveriededOrder <- rs.(types.XMaps).Get(0)
	return nil

}
func (t *testCache) makeBindedOrder(count int) error {
	if err := t.makePaidOrder(count); err != nil {
		return err
	}
	for {
		if len(t.paidOrder) == 0 {
			return nil
		}
		select {
		case order := <-t.paidOrder:
			//构建context
			input := string(order.Marshal())
			ctx := mock.NewContext(input)

			//构建请求处理-----------------
			bind := &bind.Bind{}

			ps := bind.StartHandle(ctx)
			ctx.Response().WriteAny(getAny(ps))
			status, _, _ := ctx.Response().GetFinalResponse()
			if status != 200 {
				return fmt.Errorf("绑定订单失败:%v", ps)
			}
			t.bindedOrder <- ps.(types.XMaps).Get(0)
		}
	}
}
func (t *testCache) makeNotifyOrder(count int) error {
	err := orderCache.makeDeliveryOrder(1)
	if err != nil {
		return err
	}

	for {
		if len(t.deliveriededOrder) == 0 {
			return nil
		}
		select {
		case order := <-t.deliveriededOrder:

			input := string(order.Marshal())
			ctx := mock.NewContext(input)

			//开始通知
			ntf := &notify.Notify{}
			rs := ntf.StartHandle(ctx)
			if err := errs.GetError(rs); err != nil {
				return err
			}
			//保存通知结果
			rs = ntf.SuccessHandle(ctx)
			if err := errs.GetError(rs); err != nil {
				return err
			}
			t.notifyOrder <- rs.(types.XMaps).Get(0)
		}
	}

}
func getAny(ps interface{}) interface{} {
	switch vs := ps.(type) {
	case types.XMaps:
		return vs.Get(0)
	default:
		return ps
	}
}

func preper() error {
	logger.Pause()
	defer logger.Close()
	hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra:123456@tcp(192.168.0.36:3306)/hydra?charset=utf8")
	// hydra.Conf.Vars().DB().MySQLByConnStr("db", "hydra:123456@tcp(222.209.84.37:10036)/hydra?charset=utf8")
	hydra.Conf.Vars().Queue().LMQ("queue")

	mock.NewAPPConf() //构建app配置

	//创建支付账户
	baccount := beanpay.GetAccount(global.Def.PlatName, string(enums.AccountMerchant))
	if _, err := baccount.CreateAccount(nil, "qx001", "qx001"); err != nil {
		return err
	}

	//加款
	_, err := baccount.AddAmount(nil, "qx001", "0001", float64(100000), "测试加款")
	return err

}
