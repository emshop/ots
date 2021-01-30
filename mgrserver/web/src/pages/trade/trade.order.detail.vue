<template>
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
                    <div>{{ info.mer_order_no | fltrEnum("mer_order_no") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商品编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.mer_product_id | fltrNumberFormat(0) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">货架编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.mer_shelf_id | fltrNumberFormat(0) }}</div>
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
                    <div>{{ info.city_no | fltrEmpty }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.face | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">数量:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.num | fltrNumberFormat(0) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商品总面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.total_face | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">用户账户:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.account_name | fltrEnum("account_name") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">开票方式:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.invoice_type | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">销售折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.sell_discount | fltrNumberFormat(2) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">总销售金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.sell_amount | fltrNumberFormat(2) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商户佣金折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.mer_fee_discount | fltrNumberFormat(2) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">交易服务折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.trade_fee_discount | fltrNumberFormat(2) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">支付手续费折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.payment_fee_discount | fltrNumberFormat(2) }}</div>
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
                    <div>{{ info.create_time | fltrDate }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">订单超时时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.order_timeout | fltrDate }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">支付超时时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.payment_timeout | fltrDate }}</div>
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
                    <div class="pull-right" style="margin-right: 10px">成功绑定总面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.bind_face | fltrNumberFormat(0) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">实际成功总面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.success_face | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">实际成功总销售金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.success_sell_amount | fltrNumberFormat(2) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">实际成功总佣金金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.success_mer_fee | fltrNumberFormat(2) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">实际成功总服务费金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.success_mer_trade_fee | fltrNumberFormat(2) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">实际成功总手续费金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.success_mer_payment_fee | fltrNumberFormat(2) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">实际发货成功总成本:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.success_cost_amount | fltrNumberFormat(2) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">实际发货成功总供货商佣金:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.success_spp_fee | fltrNumberFormat(2) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">实际发货成功总供货商服务费:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.success_spp_trade_fee | fltrNumberFormat(2) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">实际成功总手续费金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.success_spp_payment_fee | fltrNumberFormat(2) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">利润:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.profit | fltrNumberFormat(2) }}</div>
                  </el-col>
                </td>
              </tr>            
            </tbody>
          </table>
        </div>
	  </el-tab-pane>
	  </el-tabs>
	</div>
</template>

<script>
	export default {
    data(){
      return {
        tabName: "TradeOrderDetail",
        info: {},
      }
    },
    mounted() {
      this.init();
    },
    created(){
    },
    methods: {
      init(){
        this.queryData()
      },
      queryData:async function() {
        this.info = await this.$http.xget("/trade/order",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>