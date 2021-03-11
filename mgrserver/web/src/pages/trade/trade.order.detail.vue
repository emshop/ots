<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="订单记录" name="TradeOrderDetail">
          <div class="table-responsive">
            <table :date="info" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.order_id | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.mer_no | fltrEnum("merchant_info") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户订单编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.mer_order_no && info.mer_order_no.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.mer_order_no}}</div>
                        <div >{{ info.mer_order_no | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.mer_order_no | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户商品:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.mer_product_id | fltrEnum("merchant_product") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户货架:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.mer_shelf_id | fltrEnum("merchant_shelf") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">外部商品编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.mer_product_no | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">产品线:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.pl_id | fltrEnum("product_line") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">品牌:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.brand_no | fltrEnum("brand") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">省份:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.province_no | fltrEnum("province") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">城市:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.city_no | fltrEnum("city") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.face |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">数量:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.num |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商品总面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.total_face |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">用户账户:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.account_name && info.account_name.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.account_name}}</div>
                        <div >{{ info.account_name | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.account_name | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发票:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.invoice_type|fltrTextColor">{{ info.invoice_type | fltrEnum("invoice_type") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">销售折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.sell_discount |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">总销售金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.sell_amount |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户佣金折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.mer_fee_discount |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">交易服务折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.trade_fee_discount |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支付手续费折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.payment_fee_discount |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">是否拆单:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.can_split_order|fltrTextColor">{{ info.can_split_order | fltrEnum("bool") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.create_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">完成时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.finish_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单超时时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.order_timeout | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支付超时时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.payment_timeout | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">绑定面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.bind_face |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货暂停:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.delivery_pause|fltrTextColor">{{ info.delivery_pause | fltrEnum("bool") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.order_status|fltrTextColor">{{ info.order_status | fltrEnum("order_status") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支付状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.payment_status|fltrTextColor">{{ info.payment_status | fltrEnum("process_status") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.delivery_status|fltrTextColor">{{ info.delivery_status | fltrEnum("process_status") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">退款状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.refund_status|fltrTextColor">{{ info.refund_status | fltrEnum("process_status") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">通知状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.notify_status|fltrTextColor">{{ info.notify_status | fltrEnum("process_status") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">用户退款:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.is_refund|fltrTextColor">{{ info.is_refund | fltrEnum("bool") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">成功总面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_face |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">成功销售金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_sell_amount |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户佣金金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_mer_fee |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户服务费金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_mer_trade_fee |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户手续费金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_mer_payment_fee |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">成本金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_cost_amount |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商佣金:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_spp_fee |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商服务费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_spp_trade_fee |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商手续费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_spp_payment_fee |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">利润:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.profit |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>            
              </tbody>
            </table>
          </div>
        </el-tab-pane>
        <el-tab-pane label="订单发货表" name="TradeDeliveryDetail">
          <div class="table-responsive">
            <table :date="tradeDeliveryInfo" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.delivery_id | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.order_id | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ tradeDeliveryInfo.spp_no | fltrEnum("ots_supplier_info") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商商品编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.spp_product_id |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户商品编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.mer_product_id |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">产品线:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ tradeDeliveryInfo.pl_id | fltrEnum("ots_product_line") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">品牌:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ tradeDeliveryInfo.brand_no | fltrEnum("brand") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">省份:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ tradeDeliveryInfo.province_no | fltrEnum("province") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">城市:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ tradeDeliveryInfo.city_no | fltrEnum("city") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支持开票:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="tradeDeliveryInfo.invoice_type|fltrTextColor">{{ tradeDeliveryInfo.invoice_type | fltrEnum("invoice_type") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">用户账户:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="tradeDeliveryInfo.account_name && tradeDeliveryInfo.account_name.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{ tradeDeliveryInfo.account_name}}</div>
                        <div >{{ tradeDeliveryInfo.account_name | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ tradeDeliveryInfo.account_name | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="tradeDeliveryInfo.delivery_status|fltrTextColor">{{ tradeDeliveryInfo.delivery_status | fltrEnum("delivery_status") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支付状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="tradeDeliveryInfo.payment_status|fltrTextColor">{{ tradeDeliveryInfo.payment_status | fltrEnum("process_status") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.create_time | fltrDate("yyyy-MM-dd") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商品面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.face |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货数量:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.num |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货总面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.total_face |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">扣款折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.cost_discount |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">实际折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.real_discount |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户佣金:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.spp_fee_discount |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">交易服务费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.trade_fee_discount |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支付手续费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.payment_fee_discount |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货成本:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.cost_amount |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商服务费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.trade_fee_amount |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">成功面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.succ_face |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">开始时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.start_time | fltrDate("yyyy-MM-dd") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">结束时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.end_time | fltrDate("yyyy-MM-dd") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商发货编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.spp_delivery_no | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商商品编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.spp_product_no | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货结果:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="tradeDeliveryInfo.return_msg && tradeDeliveryInfo.return_msg.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{ tradeDeliveryInfo.return_msg}}</div>
                        <div >{{ tradeDeliveryInfo.return_msg | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ tradeDeliveryInfo.return_msg | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">扩展参数json:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="tradeDeliveryInfo.request_params && tradeDeliveryInfo.request_params.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{ tradeDeliveryInfo.request_params}}</div>
                        <div >{{ tradeDeliveryInfo.request_params | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ tradeDeliveryInfo.request_params | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">结果来源:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="tradeDeliveryInfo.result_source|fltrTextColor">{{ tradeDeliveryInfo.result_source | fltrEnum("result_source") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货结果码:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.result_code | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">最后更新时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ tradeDeliveryInfo.last_update_time | fltrDate("yyyy-MM-dd") }}</div>
                    </el-col>
                  </td>
                </tr>            
              </tbody>
            </table>
          </div>
        </el-tab-pane>
        
      </el-tabs>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return {
      tabName: "TradeOrderDetail",
      info: {},
      tradeDeliveryInfo:{},
			maxHeight: 0
    }
  },
  mounted() {
    this.$nextTick(()=>{
			this.maxHeight = this.$utility.getTableHeight("panel-body")
		})
    this.init();
  },
  created(){
  },
  methods: {
    init(){
      this.queryDetailData()
    },
    queryDetailData() {
      this.info = this.$http.xget("/trade/order",this.$route.query)
    },
    queryTradeDeliveryData() {
      this.tradeDeliveryInfo = this.$http.xget("/trade/delivery",{ delivery_id: this.info.delivery_id })
    },
    handleClick(tab) {
      switch (tab.name) {
        case "TradeOrderDetail":
          this.queryDetailData();
          break;
        case "TradeDeliveryDetail":
          this.queryTradeDeliveryData();
          break;
        default:
          this.$notify({
            title: "警告",
            message: "选项卡错误！"
          });
          return;
      }
    }
  },
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
