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
                    <el-tooltip class="item" v-if="info.mer_name && info.mer_name.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.mer_name}}</div>
                      <div >{{ info.mer_name | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div>{{ info.mer_name}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">公司:</div>
                  </el-col>
                  <el-col :span="6">
                    <el-tooltip class="item" v-if="info.mer_crop && info.mer_crop.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.mer_crop}}</div>
                      <div >{{ info.mer_crop | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div>{{ info.mer_crop}}</div>
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