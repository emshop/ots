<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body" id="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-input clearable v-model="queryData.flow_name" placeholder="请输入流程名称">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-input clearable v-model="queryData.tag_name" placeholder="请输入tag标签">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.pl_id" clearable filterable class="input-cos" placeholder="请选择产品线">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-button type="primary" @click="query" size="small">查询</el-button>
				</el-form-item>
				
				<el-form-item>
					<el-button type="success" size="small" @click="showAdd">添加</el-button>
				</el-form-item>
				
			</el-form>
		</div>
    	<!-- query end -->

    	<!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" stripe style="width: 100%" :max-height="maxHeight">
				
				<el-table-column   prop="flow_id" label="流程编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.flow_id}}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="flow_name" label="流程名称" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.flow_name && scope.row.flow_name.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.flow_name}}</div>
							<span>{{scope.row.flow_name | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.flow_name}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="tag_name" label="tag标签" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.tag_name && scope.row.tag_name.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.tag_name}}</div>
							<span>{{scope.row.tag_name | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.tag_name}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="pl_id" label="产品线" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.pl_id | fltrEnum("product_line")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="success_flow_id" label="成功流程" align="center">
					<template slot-scope="scope">
						<span :class="scope.row.success_flow_id|fltrTextColor">{{scope.row.success_flow_id | fltrEnum("product_flow")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="failed_flow_id" label="失败流程" align="center">
					<template slot-scope="scope">
						<span :class="scope.row.failed_flow_id|fltrTextColor">{{scope.row.failed_flow_id | fltrEnum("product_flow")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="unknown_flow_id" label="未知流程" align="center">
					<template slot-scope="scope">
						<span :class="scope.row.unknown_flow_id|fltrTextColor">{{scope.row.unknown_flow_id | fltrEnum("product_flow")}}</span>
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
			plID: this.$enum.get("product_line"),
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
      let res = this.$http.xpost("/product/flow/query",this.$utility.delEmptyProperty(this.queryData))
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
