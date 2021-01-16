
<template>
  <div ref="main">
    <div class="panel panel-default">
    <div class="panel-body">
    <form class="form-inline" role="form">

 
<div class="form-group select">
	<select
	v-model="query.mer_no"
	name="mer_no"
	class="form-control visible-md visible-lg">
	<option value selected="selected">---请选择商户编号---</option>
  </select>
  </div>
   
<div class="form-group select">
	<select
	v-model="query.pl_id"
	name="pl_id"
	class="form-control visible-md visible-lg">
	<option value selected="selected">---请选择产品线---</option>
  </select>
  </div>
   
<div class="form-group input">
<input
type="text"
class="form-control"
v-model="query.input"
onkeypress="if(event.keyCode == 13) return false;"
placeholder="订单编号/商户订单编号/品牌/省份/"
/>
</div>
<a class="btn btn-success" @click="search">查询</a>
</form>

<el-scrollbar style="height:100%">
<el-table :data="dataList.items" stripe style="width: 100%">

<el-table-column align="center" width="100" prop="order_id" label="订单编号">
  <template slot-scope="scope">
    <a :href="'/trade/delivery?order_id='+scope.row.order_id" v-text="scope.row.order_id"></a>
  </template>

</el-table-column>


<el-table-column align="center" width="100" prop="mer_no" label="商户编号"></el-table-column>


<el-table-column align="center" width="100" prop="pl_id" label="产品线"></el-table-column>


<el-table-column align="center" width="100" prop="brand_no" label="品牌"></el-table-column>


<el-table-column align="center" width="100" prop="province_no" label="省份"></el-table-column>


<el-table-column align="center" width="100" prop="face" label="面值"></el-table-column>


<el-table-column align="center" width="100" prop="num" label="数量"></el-table-column>


<el-table-column align="center" width="100" prop="account_name" label="用户账户"></el-table-column>


<el-table-column align="center" width="100" prop="sell_discount" label="销售折扣"></el-table-column>


  <el-table-column align="center" width="100" prop="create_time" label="创建时间">
  <template slot-scope="scope">
    <i class="el-icon-time"></i>
    <span style="margin-left: 10px" v-text="scope.row.create_time"></span>
  </template>
  </el-table-column>


<el-table-column align="center" width="100" prop="order_status" label="订单状态"></el-table-column>
 
  <el-table-column align="center" label="操作">
	<template slot-scope="scope">
	  <el-button plain type="primary" size="mini" @click="showModal(2,scope.row)">编辑</el-button>
	  <el-button plain type="success" size="mini" @click="userChange(0,scope.row.user_id, scope.row.user_name)" v-if="scope.row.status == 2">启用</el-button>
	</template>
  </el-table-column>
</el-table>
</el-scrollbar>

<div style="padding: 10px 15px;text-align: right;">
<el-pagination
  @size-change="pageSizeChnage"
  @current-change="pageIndexChange"
  :current-page="paging.pi"
  :page-size="paging.ps"
  :page-sizes="paging.sizes"
  layout="total, sizes, prev, pager, next, jumper"
  :total="dataList.count"
></el-pagination>
</div>

</div>
</div>
</div>
</template>


<script>
export default {
  components: {
   // "bootstrap-modal": require("vue2-bootstrap-modal"),
   // pager: require("vue-simple-pager"),
  },
  data() {
    return {
      paging: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
      query:{mer_no:"",pl_id:"" },
      dataList: {count: 0,items: []},
    };
  },
  created() {
    this.queryData()
  },
 mounted() {},
 methods: {
   pageIndexChange:function(val) {
     this.paging.pi = val;
     this.queryData();
   },
   pageSizeChnage:function(val) {
     this.paging.ps = val;
     this.queryData();
   },
   search: function() {
     this.paging.pi = 1;
     this.queryData();
   },
   queryData:async function(){
     Object.assign(this.query, this.paging);  
     let res = await this.$http.xpost("/trade/order/query",this.query)
     this.dataList.items = res.items
     this.dataList.count = res.count
   },

  },
}
</script>
<style>
</style>
