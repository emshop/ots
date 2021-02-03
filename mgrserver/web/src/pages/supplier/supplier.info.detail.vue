<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="供货商信息" name="SupplierInfoDetail">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.spp_no | fltrEmpty }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">供货商名称:</div>
                  </el-col>
                  <el-col :span="6">
                    <el-tooltip class="item" v-if="info.spp_name && info.spp_name.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.spp_name}}</div>
                      <div >{{ info.spp_name | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div>{{ info.spp_name}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">所属公司:</div>
                  </el-col>
                  <el-col :span="6">
                    <el-tooltip class="item" v-if="info.mer_crop && info.mer_crop.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.mer_crop}}</div>
                      <div >{{ info.mer_crop | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div>{{ info.mer_crop}}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商务人员:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.bd_uid |  fltrNumberFormat(0)}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.status|fltrTextColor">{{ info.status | fltrEnum("status") }}</div>
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
        tabName: "SupplierInfoDetail",
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
        this.info = await this.$http.xget("/supplier/info",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>