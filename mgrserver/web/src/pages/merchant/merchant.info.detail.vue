<template>
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
                    <div class="pull-right" style="margin-right: 10px">名称:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.mer_name | fltrEnum("mer_name") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">公司:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.mer_crop | fltrEnum("mer_crop") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">类型:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.mer_type | fltrNumberFormat(0) }}</div>
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
        tabName: "MerchantInfoDetail",
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
        this.info = await this.$http.xget("/merchant/info",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>