<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body" id="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-select size="medium" v-model="queryData.mer_no" clearable filterable class="input-cos" placeholder="请选择商户名称">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in merNo" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.notify_status" clearable filterable class="input-cos" placeholder="请选择通知状态">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in notifyStatus" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-button type="primary" @click="query" size="small">查询</el-button>
				</el-form-item>
				
			</el-form>
		</div>
    	<!-- query end -->

    	<!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" stripe style="width: 100%" :max-height="maxHeight">
				
				<el-table-column   prop="order_id" label="订单编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.order_id}}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="mer_no" label="商户名称" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.mer_no | fltrEnum("merchant_info")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="mer_order_no" label="订单编号" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.mer_order_no && scope.row.mer_order_no.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.mer_order_no}}</div>
							<span>{{scope.row.mer_order_no | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.mer_order_no}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="notify_status" label="通知状态" align="center">
					<template slot-scope="scope">
						<span :class="scope.row.notify_status|fltrTextColor">{{scope.row.notify_status | fltrEnum("process_status")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="notify_count" label="通知次数" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.notify_count | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="create_time" label="创建时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.create_time | fltrDate("yyyy-MM-dd") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="start_time" label="开始时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.start_time | fltrDate("yyyy-MM-dd") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="end_time" label="结束时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.end_time | fltrDate("yyyy-MM-dd") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="notify_msg" label="通知结果" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.notify_msg && scope.row.notify_msg.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.notify_msg}}</div>
							<span>{{scope.row.notify_msg | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.notify_msg}}</span>
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
			merNo: this.$enum.get("merchant_info"),
			notifyStatus: this.$enum.get("process_status"),
			dataList: {count: 0,items: []}, //表单数据对象,
			maxHeight: document.body.clientHeight
		}
  },
  created(){
  },
  mounted(){
		this.maxHeight = this.$utility.getTableHeight("panel-body")
    this.init()
  },
	methods:{
    /**初始化操作**/
    init(){
      this.query()
		},
    /**查询数据并赋值*/
    query(){
      this.queryData.pi = this.paging.pi
			this.queryData.ps = this.paging.ps
      let res = this.$http.xpost("/notify/info/query",this.$utility.delEmptyProperty(this.queryData))
			this.dataList.items = res.items
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
        order_id: val.order_id,
      }
      this.$emit("addTab","详情"+val.order_id,"/notify/info/detail",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
