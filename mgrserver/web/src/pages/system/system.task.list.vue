<template>
	<div class="panel panel-default">
    <!-- query start -->
		<div class="panel-body" id="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-input clearable size="medium" v-model="queryData.order_no" placeholder="请输入订单号">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.name"  clearable filterable class="input-cos" placeholder="请选择流程名称">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in name" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item label="创建时间:">
						<el-date-picker size="medium" class="input-cos" v-model="createTime" type="date" value-format="yyyy-MM-dd"  placeholder="选择日期"></el-date-picker>
				</el-form-item>
			
				<el-form-item>
					<el-button  type="primary" @click="queryDatas" size="medium">查询</el-button>
				</el-form-item>
				
			</el-form>
		</div>
    <!-- query end -->

    <!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" stripe style="width: 100%" :height="maxHeight">
				
				<el-table-column   prop="task_id" label="编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.task_id | fltrEmpty }}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="order_no" label="订单号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.order_no | fltrEmpty }}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="name" label="流程名称" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.name | fltrEnum("flow_tag")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="create_time" label="创建时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.create_time | fltrDate("MM/dd HH:mm:ss") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="last_execute_time" label="上次执行时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.last_execute_time | fltrDate("MM/dd HH:mm:ss") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="next_execute_time" label="下次执行时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.next_execute_time | fltrDate("MM/dd HH:mm:ss") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="count" label="执行次数" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.count | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="queue_name" label="消息队列" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.queue_name && scope.row.queue_name.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.queue_name}}</div>
							<span>{{scope.row.queue_name | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.queue_name | fltrEmpty }}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="status" label="状态" align="center">
					<template slot-scope="scope">
						<span :class="scope.row.status|fltrTextColor">{{scope.row.status | fltrEnum("process_status")}}</span>
					</template>
				</el-table-column>
				<el-table-column  label="操作" align="center">
					<template slot-scope="scope">
						<el-button type="text" size="mini" @click="showDetail(scope.row)">详情</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-scrollbar>
		<!-- list end-->

		

		

		<!-- pagination start -->
		<div class="page-pagination">
		<el-pagination
			@size-change="pageSizeChange"
			@current-change="pageIndexChange"
			:current-page="paging.pi"
			:page-size="paging.ps"
			:page-sizes="paging.sizes"
			layout="total, sizes, prev, pager, next, jumper"
			:total="dataList.count">
		</el-pagination>
		</div>
		<!-- pagination end -->

	</div>
</template>


<script>
export default {
  components: {
  },
  data () {
		return {
			paging: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
			editData:{},                //编辑数据对象
			addData:{},                 //添加数据对象 
      queryData:{},               //查询数据对象
			name: this.$enum.get("flow_tag"),
			createTime: this.$utility.dateFormat(new Date(),"yyyy-MM-dd"),
			dataList: {count: 0,items: []}, //表单数据对象,
			maxHeight: 0
		}
  },
  created(){
  },
  mounted(){
		this.$nextTick(()=>{
			this.maxHeight = this.$utility.getTableHeight("panel-body")
		})
    this.init()
  },
	methods:{
    /**初始化操作**/
    init(){
      this.query()
		},
    /**查询数据并赋值*/
		queryDatas() {
      this.paging.pi = 1
      this.query()
    },
    query(){
      this.queryData.pi = this.paging.pi
			this.queryData.ps = this.paging.ps
			this.queryData.create_time = this.$utility.dateFormat(this.createTime,"yyyy-MM-dd")
      let res = this.$http.xpost("/system/task/query",this.$utility.delEmptyProperty(this.queryData))
			this.dataList.items = res.items || []
			this.dataList.count = res.count
    },
    /**改变页容量*/
		pageSizeChange(val) {
      this.paging.ps = val
      this.query()
    },
    /**改变当前页码*/
    pageIndexChange(val) {
      this.paging.pi = val
      this.query()
    },
    /**重置添加表单*/
    resetForm(formName) {
      this.dialogAddVisible = false
      this.$refs[formName].resetFields();
		},
		showDetail(val){
			var data = {
        task_id: val.task_id,
      }
      this.$emit("addTab","详情"+val.task_id,"/system/task/detail",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
