<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="商户货架" name="MerchantShelfDetail">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.mer_shelf_id | fltrEmpty }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">货架名称:</div>
                  </el-col>
                  <el-col :span="6">
                    <el-tooltip class="item" v-if="info.mer_shelf_name && info.mer_shelf_name.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.mer_shelf_name}}</div>
                      <div >{{ info.mer_shelf_name | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div>{{ info.mer_shelf_name | fltrEmpty }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商户名称:</div>
                  </el-col>
                  <el-col :span="6">
                    <div >{{ info.mer_no | fltrEnum("merchant_info") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商户佣金:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.mer_fee_discount |  fltrNumberFormat(5)}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">交易服务费:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.trade_fee_discount |  fltrNumberFormat(5)}}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">支付手续费:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.payment_fee_discount |  fltrNumberFormat(5)}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">订单超时时长:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.order_timeout |  fltrNumberFormat(0)}}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">支付超时时长:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.payment_timeout |  fltrNumberFormat(0)}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">开票方式:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.invoice_type|fltrTextColor">{{ info.invoice_type | fltrEnum("invoice_type") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">允许退款:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.can_refund|fltrTextColor">{{ info.can_refund | fltrEnum("bool") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">单次购买数量:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.limit_count |  fltrNumberFormat(0)}}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">允许拆单:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.can_split_order|fltrTextColor">{{ info.can_split_order | fltrEnum("bool") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.status|fltrTextColor">{{ info.status | fltrEnum("status") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.create_time | fltrDate }}</div>
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
        tabName: "MerchantShelfDetail",
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
      queryData() {
        this.info = this.$http.xget("/merchant/shelf",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>