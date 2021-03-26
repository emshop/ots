<template>
	<div class="panel panel-default">
    <!-- query start -->
		<div class="panel-body" id="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-input clearable size="medium" v-model="queryData.pg_name" placeholder="请输入包名称">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.brand_no"  clearable filterable class="input-cos" placeholder="请选择品牌">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in brandNo" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-input clearable size="medium" v-model="queryData.num" placeholder="请输入数量">
					</el-input>
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
			<el-table :data="dataList.items" stripe style="width: 100%" :height="maxHeight" @sort-change="sort">
				
				<el-table-column   prop="pg_id" label="包编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.pg_id | fltrEmpty }}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="pg_name" label="包名称" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.pg_name && scope.row.pg_name.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.pg_name}}</div>
							<span>{{scope.row.pg_name | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.pg_name | fltrEmpty }}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="pl_id" label="产品线" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.pl_id | fltrEnum("product_line")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="mer_product_id" label="商品编号" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.mer_product_id | fltrEnum("merchant_product")}}</span>
					</template>
				</el-table-column>
				<el-table-column  sortable="custom" prop="brand_no" label="品牌" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.brand_no | fltrEnum("brand")}}</span>
					</template>
				</el-table-column>
				<el-table-column  sortable="custom" prop="province_no" label="省份" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.province_no | fltrEnum("province")}}</span>
					</template>
				</el-table-column>
				<el-table-column  sortable="custom" prop="city_no" label="城市" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.city_no | fltrEnum("city")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="face" label="商品面值" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.face | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="num" label="数量" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.num | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column  sortable="custom" prop="discount" label="销售折扣" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.discount | fltrNumberFormat(5)}}</span>
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
import Add from "./merchant.package.add"
import Edit from "./merchant.package.edit"
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
			brandNo: this.$enum.get("brand"),
			status: this.$enum.get("status"),
			order: "t.brand_no desc, t.province_no desc, t.city_no desc, t.discount desc",
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
		sort(column) {
      if (column.order == "ascending") {
        this.order ="t." +  column.prop + " " + "asc"
      } else if (column.order == "descending") {
        this.order ="t." +  column.prop + " " + "desc"
      } else {
        this.order = ""
      }
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
			this.queryData.order_by = this.order
      let res = this.$http.xpost("/merchant/package/query",this.$utility.delEmptyProperty(this.queryData))
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
        pg_id: val.pg_id,
      }
      this.$emit("addTab","详情"+val.pg_id,"/merchant/package/detail",data);
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
