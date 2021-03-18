
<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="资金变动" name="AccountRecordDetail">
          <div class="table-responsive">
            <table :date="info" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">变动编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.record_id | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">帐户编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.account_id | fltrEnum("account_info") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">交易编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.trade_no | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">拓展编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.ext_no | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">交易类型:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.trade_type|fltrTextColor">{{ info.trade_type | fltrEnum("trade_type") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">变动类型:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.change_type|fltrTextColor">{{ info.change_type | fltrEnum("change_type") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">变动金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.amount |  fltrNumberFormat(2)}}</div>
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
                      <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.create_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">交易说明:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.memo && info.memo.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.memo}}</div>
                        <div >{{ info.memo | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.memo | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">扩展字段:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.ext && info.ext.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.ext}}</div>
                        <div >{{ info.ext | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.ext | fltrEmpty }}</div>
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
      tabName: "AccountRecordDetail",
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
      this.info = this.$http.xget("/account/record",this.$route.query)
    },
    handleClick(tab) {
      switch (tab.name) {
        case "AccountRecordDetail":
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
