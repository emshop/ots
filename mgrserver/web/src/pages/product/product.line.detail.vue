<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="产品线" name="ProductLineDetail">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">产品线编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.pl_id | fltrNumberFormat(0) }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">产品线名称:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.pl_name | fltrEnum("pl_name") }}</div>
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
        tabName: "ProductLineDetail",
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
        this.info = await this.$http.xget("/product/line",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>