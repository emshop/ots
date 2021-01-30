<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="供货商错误码" name="SupplierErrorCodeDetail">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.id | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.spp_no | fltrEmpty }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">产品线:</div>
                  </el-col>
                  <el-col :span="6">
                    <div >{{ info.pl_id | fltrEnum("product_line") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">分类:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.category | fltrEmpty }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">处理码:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.deal_code | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">错误码:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.error_code | fltrEmpty }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.status | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">错误码描述:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.error_desc | fltrEnum("error_desc") }}</div>
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
        tabName: "SupplierErrorCodeDetail",
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
        this.info = await this.$http.xget("/supplier/error/code",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>