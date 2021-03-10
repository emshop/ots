<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="业务流程" name="ProductFlowDetail">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">流程编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.flow_id | fltrEmpty }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">流程名称:</div>
                  </el-col>
                  <el-col :span="6">
                    <div >{{ info.flow_tag | fltrEnum("flow_tag") }}</div>
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
                    <div class="pull-right" style="margin-right: 10px">队列名称:</div>
                  </el-col>
                  <el-col :span="6">
                    <el-tooltip class="item" v-if="info.queue_name && info.queue_name.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.queue_name}}</div>
                      <div >{{ info.queue_name | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div v-else>{{ info.queue_name | fltrEmpty }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">执行间隔:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.scan_interval |  fltrNumberFormat(0)}}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">延后时长:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.delay |  fltrNumberFormat(0)}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">超时时长:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.timeout |  fltrNumberFormat(0)}}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">最大次数:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.max_count |  fltrNumberFormat(0)}}</div>
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
      tabName: "ProductFlowDetail",
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
      this.info = this.$http.xget("/product/flow",this.$route.query)
    },
    handleClick(tab) {
      switch (tab.name) {
      case "ProductFlowDetail":
        this.queryData();
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