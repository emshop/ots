<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="任务表" name="SystemTaskDetail">
          <div class="table-responsive">
            <table :date="info" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.task_id | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.order_no | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">流程名称:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.name | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.create_time | fltrDate("MM/dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">上次执行时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.last_execute_time | fltrDate("yyyy-MM-dd") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">下次执行时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.next_execute_time | fltrDate("MM/dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">执行期限:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.max_execute_time | fltrDate("yyyy-MM-dd") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">时间间隔:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.next_interval |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">删除间隔:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.delete_interval |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">删除期限:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.delete_time | fltrDate("yyyy-MM-dd") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">执行次数:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.count |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">最大执行次数:</div>
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
                      <div :class="info.status|fltrTextColor">{{ info.status | fltrEnum("process_status") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">执行批次号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.batch_id |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">消息队列:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.queue_name && info.queue_name.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.queue_name}}</div>
                        <div >{{ info.queue_name | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.queue_name | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">消息内容:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.msg_content && info.msg_content.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.msg_content}}</div>
                        <div >{{ info.msg_content | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.msg_content | fltrEmpty }}</div>
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
      tabName: "SystemTaskDetail",
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
      this.info = this.$http.xget("/system/task",this.$route.query)
    },
    handleClick(tab) {
      switch (tab.name) {
        case "SystemTaskDetail":
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
