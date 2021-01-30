<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="供货商货架" name="SupplierShelfDetail">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">货架编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.spp_shelf_id | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">货架名称:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.spp_shelf_name | fltrEnum("spp_shelf_name") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">供货商:</div>
                  </el-col>
                  <el-col :span="6">
                    <div >{{ info.spp_no | fltrEnum("supplier_info") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">请求地址:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.req_url | fltrEnum("req_url") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">查询地址:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.query_url | fltrEnum("query_url") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">回调地址:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.notify_url | fltrEnum("notify_url") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">开发票:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.invoice_type|fltrTextColor">{{ info.invoice_type | fltrEnum("bool") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商户佣金:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.spp_fee_discount | fltrNumberFormat(2) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">交易服务费:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.trade_fee_discount | fltrNumberFormat(2) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">支付手续费:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.payment_fee_discount | fltrNumberFormat(2) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">支持退货:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.can_refund|fltrTextColor">{{ info.can_refund | fltrEnum("bool") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">货架状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.status|fltrTextColor">{{ info.status | fltrEnum("status") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">单次最大发货数量:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.limit_count | fltrNumberFormat(0) }}</div>
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
        tabName: "SupplierShelfDetail",
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
        this.info = await this.$http.xget("/supplier/shelf",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>