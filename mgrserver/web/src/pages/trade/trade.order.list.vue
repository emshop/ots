<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body" id="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-input clearable size="medium" v-model="queryData.order_id" placeholder="请输入订单编号">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.mer_no" clearable filterable class="input-cos" placeholder="请选择商户编号">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in merNo" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.pl_id" clearable filterable class="input-cos" placeholder="请选择产品线">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.province_no" @change="setCityNo(queryData.province_no)" clearable filterable class="input-cos" placeholder="请选择省份">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in provinceNo" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.city_no" clearable filterable class="input-cos" placeholder="请选择城市">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in cityNo" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item label="创建时间:">
						<el-date-picker size="medium" class="input-cos" v-model="createTime" type="date" value-format="yyyy-MM-dd"  placeholder="选择日期"></el-date-picker>
				</el-form-item>
			
				<el-form-item>
					<el-button  type="primary" @click="query" size="medium">查询</el-button>
				</el-form-item>
				
			</el-form>
		</div>
    	<!-- query end -->

    	<!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" stripe style="width: 100%" :height="maxHeight">
				
				<el-table-column fixed sortable prop="order_id" label="订单编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.order_id}}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="mer_no" label="商户编号" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.mer_no | fltrEnum("merchant_info")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="mer_order_no" label="商户订单编号" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.mer_order_no && scope.row.mer_order_no.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.mer_order_no}}</div>
							<span>{{scope.row.mer_order_no | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.mer_order_no}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="pl_id" label="产品线" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.pl_id | fltrEnum("product_line")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="brand_no" label="品牌" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.brand_no | fltrEnum("brand")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="province_no" label="省份" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.province_no | fltrEnum("province")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="face" label="面值" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.face | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="num" label="数量" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.num | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="account_name" label="用户账户" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.account_name && scope.row.account_name.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.account_name}}</div>
							<span>{{scope.row.account_name | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.account_name}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="sell_discount" label="销售折扣" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.sell_discount | fltrNumberFormat(2)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="create_time" label="创建时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.create_time | fltrDate("MM/dd HH:mm:ss") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="bind_face" label="绑定面值" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.bind_face | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="order_status" label="订单状态" align="center">
					<template slot-scope="scope">
						<span :class="scope.row.order_status|fltrTextColor">{{scope.row.order_status | fltrEnum("order_status")}}</span>
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
			plID: this.$enum.get("product_line"),
			provinceNo: this.$enum.get("province"),
			cityNo: [],
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
		setCityNo(pid){
			this.cityNo=[];
			this.queryData.city_no = ""
			this.cityNo=this.$enum.get("city",pid)
		},
    /**查询数据并赋值*/
    query(){
      this.queryData.pi = this.paging.pi
			this.queryData.ps = this.paging.ps
			this.queryData.create_time = this.$utility.dateFormat(this.createTime,"yyyy-MM-dd")
      let res = this.$http.xpost("/trade/order/query",this.$utility.delEmptyProperty(this.queryData))
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
      this.$emit("addTab","详情"+val.order_id,"/trade/order/detail",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
