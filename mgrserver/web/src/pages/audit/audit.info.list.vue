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
					<el-select size="medium" v-model="queryData.audit_status"  clearable filterable class="input-cos" placeholder="请选择审核状态">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in auditStatus" :key="index" :value="item.value" :label="item.name"></el-option>
					</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-button  type="primary" @click="queryDatas" size="medium">查询</el-button>
				</el-form-item>
				
			</el-form>
		</div>
    <!-- query end -->

    <!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" stripe style="width: 100%" :height="maxHeight">
				
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
				<el-table-column   prop="create_time" label="创建时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.create_time | fltrDate("yyyy-MM-dd") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="delivery_status" label="发货结果" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.delivery_status | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="audit_status" label="审核状态" align="center">
					<template slot-scope="scope">
						<span :class="scope.row.audit_status|fltrTextColor">{{scope.row.audit_status | fltrEnum("status")}}</span>
					</template>
				</el-table-column>
				<el-table-column   prop="audit_by" label="审核人" align="center">
				<template slot-scope="scope">
					<span>{{scope.row.audit_by | fltrNumberFormat(0)}}</span>
				</template>
				</el-table-column>
				<el-table-column   prop="audit_time" label="审核时间" align="center">
				<template slot-scope="scope">
					<div>{{scope.row.audit_time | fltrDate("yyyy-MM-dd") }}</div>
				</template>
				</el-table-column>
				<el-table-column   prop="audit_msg" label="审核信息" align="center">
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.audit_msg && scope.row.audit_msg.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{scope.row.audit_msg}}</div>
							<span>{{scope.row.audit_msg | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{scope.row.audit_msg | fltrEmpty }}</span>
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
			auditStatus: this.$enum.get("status"),
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
      let res = this.$http.xpost("/audit/info/query",this.$utility.delEmptyProperty(this.queryData))
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
      this.$emit("addTab","详情"+val.delivery_id,"/audit/info/detail",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
