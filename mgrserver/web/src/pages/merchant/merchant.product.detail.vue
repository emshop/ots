
<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="商户商品" name="MerchantProductDetail">
          <div class="table-responsive">
            <table :date="info" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商品编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.mer_product_id | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">货架名称:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.mer_shelf_id | fltrEnum("merchant_shelf") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户名称:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.mer_no | fltrEnum("merchant_info") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商品名称:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.prod_name && info.prod_name.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.prod_name}}</div>
                        <div >{{ info.prod_name | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.prod_name | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">产品线:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.pl_id | fltrEnum("product_line") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">品牌:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.brand_no | fltrEnum("brand") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">省份:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.province_no | fltrEnum("province") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">城市:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.city_no | fltrEnum("city") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.face |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户商品:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.mer_product_no | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">销售折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.discount |  fltrNumberFormat(5)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.status|fltrTextColor">{{ info.status | fltrEnum("status") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.create_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>            
              </tbody>
            </table>
          </div>
        </el-tab-pane>
        <el-tab-pane label="组合商品" name="MerchantPackageDetail">
          <el-scrollbar style="height:100%" id="panel-body">
            <el-table :data="MerchantPackageList.items" stripe style="width: 100%" :height="maxHeight">
              
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
              <el-table-column  sortable prop="brand_no" label="品牌" align="center">
                <template slot-scope="scope">
                  <span >{{scope.row.brand_no | fltrEnum("brand")}}</span>
                </template>
              </el-table-column>
              <el-table-column  sortable prop="province_no" label="省份" align="center">
                <template slot-scope="scope">
                  <span >{{scope.row.province_no | fltrEnum("province")}}</span>
                </template>
              </el-table-column>
              <el-table-column  sortable prop="city_no" label="城市" align="center">
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
              <el-table-column   prop="status" label="状态" align="center">
                <template slot-scope="scope">
                  <span :class="scope.row.status|fltrTextColor">{{scope.row.status | fltrEnum("status")}}</span>
                </template>
              </el-table-column>
            </el-table>
          </el-scrollbar>
        </el-tab-pane>
        
      </el-tabs>
    </div>   
    <div class="page-pagination" v-show="tabName =='MerchantPackageDetail'">
    <el-pagination
      @size-change="pageMerchantPackageSizeChange"
      @current-change="pageMerchantPackageIndexChange"
      :current-page="pagingMerchantPackage.pi"
      :page-size="pagingMerchantPackage.ps"
      :page-sizes="pagingMerchantPackage.sizes"
      layout="total, sizes, prev, pager, next, jumper"
      :total="MerchantPackageList.count">
    </el-pagination>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return {
      tabName: "MerchantProductDetail",
      info: {},
      pagingMerchantPackage: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
      MerchantPackageList: {count: 0,items: []}, //表单数据对象,
      queryMerchantPackageParams:{},  //查询数据对象
			maxHeight: 0
    }
  },
  mounted() {
    this.$nextTick(()=>{
			this.maxHeight = this.$utility.getTableHeight("panel-body")
		})
    this.init();
  },
  created(){
  },
  methods: {
    init(){
      this.queryDetailData()
    },
    queryDetailData() {
      this.info = this.$http.xget("/merchant/product",this.$route.query)
    },
    /**查询数据并赋值*/
		queryMerchantPackageDatas() {
      this.pagingMerchantPackage.pi = 1
      this.queryMerchantPackageData()
    },
    queryMerchantPackageData(){
      this.queryMerchantPackageParams.pi = this.pagingMerchantPackage.pi
			this.queryMerchantPackageParams.ps = this.pagingMerchantPackage.ps
      var data = this.$utility.delEmptyProperty(this.queryMerchantPackageParams)
      data.mer_product_id = this.info.mer_product_id || ""
      let res = this.$http.xpost("/merchant/package/querydetail", data)
			this.MerchantPackageList.items = res.items || []
			this.MerchantPackageList.count = res.count
    },
    /**改变页容量*/
		pageMerchantPackageSizeChange(val) {
      this.pagingMerchantPackage.ps = val
      this.queryMerchantPackageData()
    },
    /**改变当前页码*/
    pageMerchantPackageIndexChange(val) {
      this.pagingMerchantPackage.pi = val
      this.queryMerchantPackageData()
    },
    handleClick(tab) {
      switch (tab.name) {
        case "MerchantProductDetail":
          this.queryDetailData();
          break;
        case "MerchantPackageDetail":
          this.queryMerchantPackageData();
          break;
        default:
          this.$notify({
            title: "警告",
            message: "选项卡错误！"
          });
          return;
      }
    }
  },
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
