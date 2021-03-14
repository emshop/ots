<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="订单通知表" name="NotifyInfoDetail">
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
                      <div class="pull-right" style="margin-right: 10px">商户名称:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.mer_no | fltrEnum("merchant_info") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.mer_order_no && info.mer_order_no.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.mer_order_no}}</div>
                        <div >{{ info.mer_order_no | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.mer_order_no | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">通知地址:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.notify_url && info.notify_url.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.notify_url}}</div>
                        <div >{{ info.notify_url | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.notify_url | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.create_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">开始时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.start_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">结束时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.end_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">最大通知次数:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.max_count |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">通知次数:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.notify_count |  fltrNumberFormat(0)}}</div>
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
                      <div class="pull-right" style="margin-right: 10px">通知结果:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.notify_msg && info.notify_msg.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.notify_msg}}</div>
                        <div >{{ info.notify_msg | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.notify_msg | fltrEmpty }}</div>
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
      tabName: "NotifyInfoDetail",
      info: {},
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
      this.info = this.$http.xget("/notify/info",this.$route.query)
    },
    handleClick(tab) {
      switch (tab.name) {
        case "NotifyInfoDetail":
          this.queryDetailData();
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
