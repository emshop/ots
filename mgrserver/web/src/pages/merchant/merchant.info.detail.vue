<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="商户信息" name="MerchantInfoDetail">
          <div class="table-responsive">
            <table :date="info" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.mer_no | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户名称:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.mer_name && info.mer_name.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.mer_name}}</div>
                        <div >{{ info.mer_name | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.mer_name | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">公司名称:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.mer_crop && info.mer_crop.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.mer_crop}}</div>
                        <div >{{ info.mer_crop | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.mer_crop | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">类型:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.mer_type | fltrEnum("merchant_type") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商务人员:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.bd_uid | fltrEnum("user_info") }}</div>
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
                      <div>{{ info.create_time | fltrDate("yyyy-MM-dd") }}</div>
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
      tabName: "MerchantInfoDetail",
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
      this.info = this.$http.xget("/merchant/info",this.$route.query)
    },
    handleClick(tab) {
      switch (tab.name) {
        case "MerchantInfoDetail":
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
