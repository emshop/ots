<template>
	<div class="panel panel-default">
    <!-- query start -->
		<div class="panel-body" id="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-select size="medium" v-model="queryData.flow_tag"  clearable filterable class="input-cos" placeholder="请选择流程名称">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in flowTag" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.pl_id"  clearable filterable class="input-cos" placeholder="请选择产品线">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.status"  clearable filterable class="input-cos" placeholder="请选择状态">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-button  type="primary" @click="queryDatas" size="medium">查询</el-button>
				</el-form-item>
				
				<el-form-item>
					<el-button type="success" size="medium" @click="showAdd">添加</el-button>
				</el-form-item>
				
			</el-form>
		</div>
    <!-- query end -->

    <!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" stripe style="width: 100%" :height="maxHeight" >
				
				<el-table-column   prop="flow_id" label="流程编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.flow_id | fltrEmpty }}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="flow_tag" label="流程名称" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.flow_tag | fltrEnum("flow_tag")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="pl_id" label="产品线" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.pl_id | fltrEnum("product_line")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="queue_name" label="队列名称" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.queue_name && scope.row.queue_name.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.queue_name}}</div>
							<span>{{scope.row.queue_name | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.queue_name | fltrEmpty }}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="scan_interval" label="执行间隔" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.scan_interval | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="delay" label="延后时长" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.delay | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="timeout" label="超时时长" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.timeout | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="max_count" label="最大次数" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.max_count | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="status" label="状态" align="center">
					<template slot-scope="scope">
						<span :class="scope.row.status|fltrTextColor">{{scope.row.status | fltrEnum("status")}}</span>
					</template>
				</el-table-column>
				<el-table-column  label="操作" align="center">
					<template slot-scope="scope">
						<el-button type="text" size="mini" @click="showEdit(scope.row)">编辑</el-button>
						<el-button type="text" size="mini" @click="showDetail(scope.row)">详情</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-scrollbar>
		<!-- list end-->

		<!-- Add Form -->
		<Add ref="Add" :refresh="query"></Add>
		<!--Add Form -->

		<!-- edit Form start-->
		<Edit ref="Edit" :refresh="query"></Edit>
		<!-- edit Form end-->

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
import Add from "./product.flow.add"
import Edit from "./product.flow.edit"
export default {
  components: {
		Add,
		Edit
  },
  data () {
		return {
			paging: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
			editData:{},                //编辑数据对象
			addData:{},                 //添加数据对象 
      queryData:{},               //查询数据对象
			flowTag: this.$enum.get("flow_tag"),
			plID: this.$enum.get("product_line"),
			status: this.$enum.get("status"),
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
      let res = this.$http.xpost("/product/flow/query",this.$utility.delEmptyProperty(this.queryData))
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
        flow_id: val.flow_id,
      }
      this.$emit("addTab","详情"+val.flow_id,"/product/flow/detail",data);
		},
    showAdd(){
      this.$refs.Add.show();
		},
    showEdit(val) {
      this.$refs.Edit.editData = val
      this.$refs.Edit.show();
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
