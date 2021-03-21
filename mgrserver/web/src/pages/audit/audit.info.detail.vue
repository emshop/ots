
<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="发货人工审核表" name="AuditInfoDetail">
          <div class="table-responsive">
            <table :date="info" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.delivery_id | fltrEmpty }}</div>
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
                      <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.create_time | fltrDate("yyyy-MM-dd") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货结果:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.delivery_status |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">是否终结订单:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.end_order|fltrTextColor">{{ info.end_order | fltrEnum("bool") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">是否加入黑名单:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.add_to_blacklist|fltrTextColor">{{ info.add_to_blacklist | fltrEnum("bool") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">审核状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.audit_status|fltrTextColor">{{ info.audit_status | fltrEnum("status") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">审核人:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.audit_by |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">审核时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.audit_time | fltrDate("yyyy-MM-dd") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">审核信息:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.audit_msg && info.audit_msg.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.audit_msg}}</div>
                        <div >{{ info.audit_msg | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.audit_msg | fltrEmpty }}</div>
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
      tabName: "AuditInfoDetail",
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
      this.info = this.$http.xget("/audit/info",this.$route.query)
    },
    handleClick(tab) {
      switch (tab.name) {
        case "AuditInfoDetail":
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
