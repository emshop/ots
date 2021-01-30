<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="退款申请" name="RefundApplyDetail">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">申请编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.apply_id | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">订单编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.order_id | fltrEmpty }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商户编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div >{{ info.mer_no | fltrEnum("merchant_info") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商户订单号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.mer_order_no | fltrEmpty }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">退款原因:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.refund_cause | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">申请状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div >{{ info.apply_status | fltrEnum("process_status") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">退款金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.refund_amount | fltrNumberFormat(2) }}</div>
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
        tabName: "RefundApplyDetail",
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
        this.info = await this.$http.xget("/refund/apply",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>