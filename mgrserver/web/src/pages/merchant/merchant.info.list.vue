<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body" id="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-input clearable size="medium" v-model="queryData.mer_name" placeholder="请输入商户名称">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.status" clearable filterable class="input-cos" placeholder="请选择状态">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-button  type="primary" @click="query" size="medium">查询</el-button>
				</el-form-item>
				
				<el-form-item>
					<el-button type="success" size="medium" @click="showAdd">添加</el-button>
				</el-form-item>
				
			</el-form>
		</div>
    	<!-- query end -->

    	<!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" stripe style="width: 100%" :max-height="maxHeight">
				
				<el-table-column   prop="mer_no" label="编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.mer_no}}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="mer_name" label="商户名称" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.mer_name && scope.row.mer_name.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.mer_name}}</div>
							<span>{{scope.row.mer_name | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.mer_name}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="mer_crop" label="公司名称" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.mer_crop && scope.row.mer_crop.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.mer_crop}}</div>
							<span>{{scope.row.mer_crop | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.mer_crop}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="mer_type" label="类型" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.mer_type | fltrEnum("merchant_type")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="status" label="状态" align="center">
					<template slot-scope="scope">
						<span :class="scope.row.status|fltrTextColor">{{scope.row.status | fltrEnum("status")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="create_time" label="创建时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.create_time | fltrDate("yyyy-MM-dd") }}</div>
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
import Add from "./merchant.info.add"
import Edit from "./merchant.info.edit"
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
			status: this.$enum.get("status"),
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
      let res = this.$http.xpost("/merchant/info/query",this.$utility.delEmptyProperty(this.queryData))
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
        mer_no: val.mer_no,
      }
      this.$emit("addTab","详情"+val.mer_no,"/merchant/info/detail",data);
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
