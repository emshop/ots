<template>
	<div class="panel panel-default">
    <!-- query start -->
		<div class="panel-body" id="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
			</el-form>
		</div>
    <!-- query end -->

    <!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" stripe style="width: 100%" :height="maxHeight">
				
				<el-table-column   prop="id" label="id" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.id | fltrEmpty }}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="order_no" label="交易订单号" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.order_no | fltrEmpty }}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="ip" label="用户ip" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.ip | fltrEmpty }}</span>
				</template>
				
				</el-table-column>
				<el-table-column   prop="content" label="内容" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.content && scope.row.content.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.content}}</div>
							<span>{{scope.row.content | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.content | fltrEmpty }}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="create_time" label="创建时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.create_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
				</template>
				</el-table-column>
				<el-table-column  label="操作" align="center">
					<template slot-scope="scope">
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
      let res = this.$http.xpost("/life/time/query",this.$utility.delEmptyProperty(this.queryData))
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
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
