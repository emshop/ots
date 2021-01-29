<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-select size="medium" v-model="queryData.mer_shelf_id" class="input-cos" placeholder="请选择货架编号">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in merShelfId" :key="index" :value="item.value" :label="item.name"></el-option>
						</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.mer_no" class="input-cos" placeholder="请选择商户编号">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in merNo" :key="index" :value="item.value" :label="item.name"></el-option>
						</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.pl_id" class="input-cos" placeholder="请选择产品线">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in plId" :key="index" :value="item.value" :label="item.name"></el-option>
						</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.brand_no" class="input-cos" placeholder="请选择品牌">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in brandNo" :key="index" :value="item.value" :label="item.name"></el-option>
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
			<el-table :data="dataList.items" border style="width: 100%">
				<el-table-column prop="mer_shelf_id" label="货架编号" >
					<template slot-scope="scope">
						<span>{{scope.row.mer_shelf_id | fltrEnum("merchant_info")}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="mer_no" label="商户编号" >
					<template slot-scope="scope">
						<span>{{scope.row.mer_no | fltrEnum("merchant_info")}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="pl_id" label="产品线" >
					<template slot-scope="scope">
						<span>{{scope.row.pl_id | fltrEnum("product_line")}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="brand_no" label="品牌" >
					<template slot-scope="scope">
						<span>{{scope.row.brand_no | fltrEnum("brand")}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="province_no" label="省份" >
					<template slot-scope="scope">
						<span>{{scope.row.province_no | fltrEnum("provice")}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="city_no" label="城市" >
					<template slot-scope="scope">
						<span>{{scope.row.city_no | fltrEnum("city")}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="face" label="面值" >
				<template slot-scope="scope">
					<span>{{scope.row.face | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column prop="mer_product_no" label="商户商品编号" >
				<template slot-scope="scope">
					<span>{{scope.row.mer_product_no}}</span>
				</template>
				
				</el-table-column>
				<el-table-column prop="discount" label="销售折扣" >
				<template slot-scope="scope">
					<span>{{scope.row.discount | fltrNumberFormat(2)}}</span>
				</template>
				</el-table-column>
				<el-table-column prop="status" label="状态" >
					<template slot-scope="scope">
						<span>{{scope.row.status | fltrEnum("status")}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="create_time" label="创建时间" >
				<template slot-scope="scope">
					<span>{{scope.row.create_time | fltrDate }}</span>
				</template>
				</el-table-column>
				<el-table-column  label="操作">
					<template slot-scope="scope">
						<el-button type="text" size="small" @click="showDetail(scope.row)">详情</el-button>
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
			merShelfId: this.$enum.get("merchant_info"),
			merNo: this.$enum.get("merchant_info"),
			plId: this.$enum.get("product_line"),
			brandNo: this.$enum.get("brand"),
			dataList: {count: 0,items: []}, //表单数据对象
		}
  },
  created(){
  },
  mounted(){
    this.init()
  },
	methods:{
    /**初始化操作**/
    init(){
      this.query()
		},
    /**查询数据并赋值*/
    query:async function(){
      this.queryData.pi = this.paging.pi
			this.queryData.ps = this.paging.ps
      let res = await this.$http.xpost("/merchant/product/query",this.queryData)
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
        mer_product_id: val.mer_product_id,
      }
      this.$emit("addTab","详情"+val.mer_product_id,"/merchant/product/detail",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
