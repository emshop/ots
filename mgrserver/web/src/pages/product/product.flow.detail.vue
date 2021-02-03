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
                    <el-tooltip class="item" v-if="info.flow_name && info.flow_name.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.flow_name}}</div>
                      <div >{{ info.flow_name | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div>{{ info.flow_name}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">tag标签:</div>
                  </el-col>
                  <el-col :span="6">
                    <el-tooltip class="item" v-if="info.tag_name && info.tag_name.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.tag_name}}</div>
                      <div >{{ info.tag_name | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div>{{ info.tag_name}}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">产品线:</div>
                  </el-col>
                  <el-col :span="6">
                    <div >{{ info.pl_id | fltrEnum("product_line") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">成功流程:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.success_flow_id|fltrTextColor">{{ info.success_flow_id | fltrEnum("product_flow") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">失败流程:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.failed_flow_id|fltrTextColor">{{ info.failed_flow_id | fltrEnum("product_flow") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">未知流程:</div>
                  </el-col>
                  <el-col :span="6">
                    <div :class="info.unknown_flow_id|fltrTextColor">{{ info.unknown_flow_id | fltrEnum("product_flow") }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">队列名称:</div>
                  </el-col>
                  <el-col :span="6">
                    <el-tooltip class="item" v-if="info.queue_name && info.queue_name.length > 50" effect="dark" placement="top">
                      <div slot="content" style="width: 110px">{{info.queue_name}}</div>
                      <div >{{ info.queue_name | fltrSubstr(50) }}</div>
                    </el-tooltip>
                    <div>{{ info.queue_name}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">超时时长:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.scan_interval |  fltrNumberFormat(0)}}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">延后处理时长:</div>
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
                    <div class="pull-right" style="margin-right: 10px">最大执行次数:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.max_count |  fltrNumberFormat(0)}}</div>
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
      queryData:async function() {
        this.info = await this.$http.xget("/product/flow",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>