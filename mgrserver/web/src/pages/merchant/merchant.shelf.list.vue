<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-input clearable v-model="queryData.mer_shelf_name" placeholder="请输入货架名称">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.mer_no" class="input-cos" placeholder="请选择商户编号">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in merNo" :key="index" :value="item.value" :label="item.name"></el-option>
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
					<span>{{scope.row.mer_shelf_id | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column prop="mer_shelf_name" label="货架名称" >
					<template slot-scope="scope">
						<span>{{scope.row.mer_shelf_name | fltrSubstr(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="mer_no" label="商户编号" >
					<template slot-scope="scope">
						<span >{{scope.row.mer_no | fltrEnum("merchant_info")}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="mer_fee_discount" label="商户佣金" >
				<template slot-scope="scope">
					<span>{{scope.row.mer_fee_discount | fltrNumberFormat(2)}}</span>
				</template>
				</el-table-column>
				<el-table-column prop="trade_fee_discount" label="交易服务费" >
				<template slot-scope="scope">
					<span>{{scope.row.trade_fee_discount | fltrNumberFormat(2)}}</span>
				</template>
				</el-table-column>
				<el-table-column prop="payment_fee_discount" label="支付手续费" >
				<template slot-scope="scope">
					<span>{{scope.row.payment_fee_discount | fltrNumberFormat(2)}}</span>
				</template>
				</el-table-column>
				<el-table-column prop="order_timeout" label="订单超时时长" >
				<template slot-scope="scope">
					<span>{{scope.row.order_timeout | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column prop="can_split_order" label="允许拆单" >
					<template slot-scope="scope">
						<span :class="scope.row.can_split_order|fltrTextColor">{{scope.row.can_split_order | fltrEnum("bool")}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="status" label="状态" >
					<template slot-scope="scope">
						<span :class="scope.row.status|fltrTextColor">{{scope.row.status | fltrEnum("status")}}</span>
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
			merNo: this.$enum.get("merchant_info"),
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
      let res = await this.$http.xpost("/merchant/shelf/query",this.queryData)
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
        mer_shelf_id: val.mer_shelf_id,
      }
      this.$emit("addTab","详情"+val.mer_shelf_id,"/merchant/shelf/detail",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
