
<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="账户信息" name="AccountInfoDetail">
          <div class="table-responsive">
            <table :date="info" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">帐户编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.account_id | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">帐户名称:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.account_name | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">系统标识:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.ident | fltrEnum("ident") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">账户类型:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.groups | fltrEnum("account_type") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户信息:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.eid | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">帐户余额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.balance |  fltrNumberFormat(2)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">信用余额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.credit |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.status|fltrTextColor">{{ info.status | fltrEnum("status") }}</div>
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
      tabName: "AccountInfoDetail",
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
      this.info = this.$http.xget("/account/info",this.$route.query)
    },
    handleClick(tab) {
      switch (tab.name) {
        case "AccountInfoDetail":
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
