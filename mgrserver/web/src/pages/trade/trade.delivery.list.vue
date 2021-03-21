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
					<el-select size="medium" v-model="queryData.spp_no"  clearable filterable class="input-cos" placeholder="请选择供货商">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in sppNo" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.pl_id"  clearable filterable class="input-cos" placeholder="请选择产品线">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.delivery_status"  clearable filterable class="input-cos" placeholder="请选择发货状态">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in deliveryStatus" :key="index" :value="item.value" :label="item.name"></el-option>
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
			<el-table :data="dataList.items" stripe style="width: 100%" :height="maxHeight" >
				
				<el-table-column   prop="delivery_id" label="发货编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.delivery_id | fltrEmpty }}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="order_id" label="订单编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.order_id | fltrEmpty }}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="spp_no" label="供货商" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.spp_no | fltrEnum("supplier_info")}}</span>
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
				<el-table-column   prop="account_name" label="用户账户" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.account_name && scope.row.account_name.length > 8" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.account_name}}</div>
							<span>{{scope.row.account_name | fltrSubstr(8) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.account_name | fltrEmpty }}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="delivery_status" label="发货状态" align="center">
					<template slot-scope="scope">
						<span :class="scope.row.delivery_status|fltrTextColor">{{scope.row.delivery_status | fltrEnum("delivery_status")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="create_time" label="创建时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.create_time | fltrDate("HH:mm:ss") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="face" label="商品面值" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.face | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="end_time" label="完成时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.end_time | fltrDate("MM/dd HH:mm") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="return_msg" label="发货结果" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.return_msg && scope.row.return_msg.length > 6" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.return_msg}}</div>
							<span>{{scope.row.return_msg | fltrSubstr(6) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.return_msg | fltrEmpty }}</span>
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
			sppNo: this.$enum.get("supplier_info"),
			plID: this.$enum.get("product_line"),
			deliveryStatus: this.$enum.get("delivery_status"),
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
      let res = this.$http.xpost("/trade/delivery/query",this.$utility.delEmptyProperty(this.queryData))
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
        delivery_id: val.delivery_id,
      }
      this.$emit("addTab","详情"+val.delivery_id,"/trade/delivery/detail",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
