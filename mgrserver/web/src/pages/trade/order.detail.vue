<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="订单记录" name="OtsTradeOrder">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">订单编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.order_id}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">商户编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.mer_no | fltrEnum("MerNo")}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">商户订单编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.mer_order_no}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">商品编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.mer_product_id}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">货架编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.mer_shelf_id}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">外部商品编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.mer_product_no}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">产品线:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.pl_id | fltrEnum("PlID")}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">品牌:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.brand_no}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">省份:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.province_no}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">城市:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.city_no}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.face}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">数量:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.num}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">商品总面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.total_face}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">用户账户:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.account_name}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">开票方式:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.invoice_type}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">销售折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.sell_discount}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">总销售金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.sell_amount}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">商户佣金折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.mer_fee_discount}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">交易服务折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.trade_fee_discount}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">支付手续费折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.payment_fee_discount}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">是否拆单:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.can_split_order}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">创建时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.create_time|fltrDateTime}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">订单超时时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.order_timeout|fltrDateTime}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">支付超时时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.payment_timeout|fltrDateTime}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货暂停:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.delivery_pause}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">订单状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.order_status}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">支付状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.payment_status}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.delivery_status}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">退款状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.refund_status}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">通知状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.notify_status}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">用户退款:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.is_refund}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">成功绑定总面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.bind_face}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际成功总面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.success_face}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际成功总销售金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.success_sell_amount}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际成功总佣金金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.success_mer_fee}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际成功总服务费金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.success_mer_trade_fee}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际成功总手续费金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.success_mer_payment_fee}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际发货成功总成本:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.success_cost_amount}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际发货成功总供货商佣金:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.success_spp_fee}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际发货成功总供货商服务费:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.success_spp_trade_fee}}</div>
                  </el-col>
                </td>
              </tr>
            
              <tr>
                <td>               
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际成功总手续费金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.success_spp_payment_fee}}</div>
                  </el-col>
                           
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">利润:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.profit}}</div>
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
        tabName:"OtsTradeOrder",
        info:{},
      }
    },
    created(){
       this.queryData()
    },
    methods: {
      queryData:async function() {
        console.log("this.$route.query:",this.$route.query)
        this.info = await this.$http.xget("/trade/order/single",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>