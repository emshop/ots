<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body" id="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-select size="medium" v-model="queryData.account_id" clearable filterable class="input-cos" placeholder="请选择帐户编号">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in accountID" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-input clearable size="medium" v-model="queryData.trade_no" placeholder="请输入交易编号">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.trade_type" clearable filterable class="input-cos" placeholder="请选择交易类型">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in tradeType" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.change_type" clearable filterable class="input-cos" placeholder="请选择变动类型">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in changeType" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-button  type="primary" @click="query" size="medium">查询</el-button>
				</el-form-item>
				
			</el-form>
		</div>
    	<!-- query end -->

    	<!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" stripe style="width: 100%" :max-height="maxHeight">
				
				<el-table-column   prop="record_id" label="变动编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.record_id}}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="account_id" label="帐户编号" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.account_id | fltrEnum("account_info")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="trade_no" label="交易编号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.trade_no}}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="trade_type" label="交易类型" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.trade_type | fltrEnum("trade_type")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="change_type" label="变动类型" align="center">
					<template slot-scope="scope">
						<span >{{scope.row.change_type | fltrEnum("change_type")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="amount" label="变动金额" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.amount | fltrNumberFormat(2)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="balance" label="帐户余额" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.balance | fltrNumberFormat(2)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="create_time" label="创建时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.create_time | fltrDate("yyyy-MM-dd") }}</div>
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
			accountID: this.$enum.get("account_info"),
			tradeType: this.$enum.get("trade_type"),
			changeType: this.$enum.get("change_type"),
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
      let res = this.$http.xpost("/account/record/query",this.$utility.delEmptyProperty(this.queryData))
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
        record_id: val.record_id,
      }
      this.$emit("addTab","详情"+val.record_id,"/account/record/detail",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
