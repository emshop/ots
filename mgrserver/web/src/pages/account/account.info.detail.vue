<template>
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
                    <div class="pull-right" style="margin-right: 10px">用户分组:</div>
                  </el-col>
                  <el-col :span="6">
                    <div >{{ info.groups | fltrEnum("account_group") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商户信息:</div>
                  </el-col>
                  <el-col :span="6">
                    <div >{{ info.eid | fltrEnum("merchant_info") }}</div>
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
        tabName: "AccountInfoDetail",
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
        this.info = this.$http.xget("/account/info",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>