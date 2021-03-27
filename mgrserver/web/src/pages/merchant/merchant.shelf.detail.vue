
<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="商户货架" name="MerchantShelfDetail">
          <div class="table-responsive">
            <table :date="info" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.mer_shelf_id | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">货架名称:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.mer_shelf_name && info.mer_shelf_name.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.mer_shelf_name}}</div>
                        <div >{{ info.mer_shelf_name | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.mer_shelf_name | fltrEmpty }}</div>
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
                      <div class="pull-right" style="margin-right: 10px">商户佣金:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.mer_fee_discount |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">交易服务费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.trade_fee_discount |  fltrNumberFormat(5)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支付手续费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.payment_fee_discount |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单超时时长:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.order_timeout |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支付超时时长:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.payment_timeout |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">单次最大数量:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.limit_count |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">允许开票:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.invoice_type|fltrTextColor">{{ info.invoice_type | fltrEnum("invoice_type") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">允许退款:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.can_refund|fltrTextColor">{{ info.can_refund | fltrEnum("bool") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">指定上游:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.assign_upstream|fltrTextColor">{{ info.assign_upstream | fltrEnum("bool") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.status|fltrTextColor">{{ info.status | fltrEnum("status") }}</div>
                    </el-col>                 
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
        <el-tab-pane label="商户上游" name="MerchantUpstreamDetail">
          <el-scrollbar style="height:100%" id="panel-body">
            <el-table :data="MerchantUpstreamList.items" stripe style="width: 100%" :height="maxHeight">
              
              <el-table-column   prop="mu_id" label="编号" align="center">
              <template slot-scope="scope">
                <span>{{scope.row.mu_id | fltrEmpty }}</span>
              </template>
              
              </el-table-column>
              <el-table-column   prop="mer_shelf_id" label="货架" align="center">
                <template slot-scope="scope">
                  <span >{{scope.row.mer_shelf_id | fltrEnum("merchant_shelf")}}</span>
                </template>
              </el-table-column>
              <el-table-column   prop="spp_no" label="供货商" align="center">
                <template slot-scope="scope">
                  <span >{{scope.row.spp_no | fltrEnum("supplier_info")}}</span>
                </template>
              </el-table-column>
              <el-table-column   prop="status" label="状态" align="center">
                <template slot-scope="scope">
                  <span :class="scope.row.status|fltrTextColor">{{scope.row.status | fltrEnum("status")}}</span>
                </template>
              </el-table-column>
              <el-table-column   prop="create_time" label="创建时间" align="center">
              <template slot-scope="scope">
                <div>{{scope.row.create_time | fltrDate("yyyy-MM-dd") }}</div>
              </template>
              </el-table-column>
            </el-table>
          </el-scrollbar>
        </el-tab-pane>
        
      </el-tabs>
    </div>   
    <div class="page-pagination" v-show="tabName =='MerchantUpstreamDetail'">
    <el-pagination
      @size-change="pageMerchantUpstreamSizeChange"
      @current-change="pageMerchantUpstreamIndexChange"
      :current-page="pagingMerchantUpstream.pi"
      :page-size="pagingMerchantUpstream.ps"
      :page-sizes="pagingMerchantUpstream.sizes"
      layout="total, sizes, prev, pager, next, jumper"
      :total="MerchantUpstreamList.count">
    </el-pagination>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return {
      tabName: "MerchantShelfDetail",
      info: {},
      pagingMerchantUpstream: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
      MerchantUpstreamList: {count: 0,items: []}, //表单数据对象,
      queryMerchantUpstreamParams:{},  //查询数据对象
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
      this.info = this.$http.xget("/merchant/shelf",this.$route.query)
    },
    /**查询数据并赋值*/
		queryMerchantUpstreamDatas() {
      this.pagingMerchantUpstream.pi = 1
      this.queryMerchantUpstreamData()
    },
    queryMerchantUpstreamData(){
      this.queryMerchantUpstreamParams.pi = this.pagingMerchantUpstream.pi
			this.queryMerchantUpstreamParams.ps = this.pagingMerchantUpstream.ps
      var data = this.$utility.delEmptyProperty(this.queryMerchantUpstreamParams)
      data.mer_shelf_id = this.info.mer_shelf_id || ""
      let res = this.$http.xpost("/merchant/upstream/querydetail", data)
			this.MerchantUpstreamList.items = res.items || []
			this.MerchantUpstreamList.count = res.count
    },
    /**改变页容量*/
		pageMerchantUpstreamSizeChange(val) {
      this.pagingMerchantUpstream.ps = val
      this.queryMerchantUpstreamData()
    },
    /**改变当前页码*/
    pageMerchantUpstreamIndexChange(val) {
      this.pagingMerchantUpstream.pi = val
      this.queryMerchantUpstreamData()
    },
    handleClick(tab) {
      switch (tab.name) {
        case "MerchantShelfDetail":
          this.queryDetailData();
          break;
        case "MerchantUpstreamDetail":
          this.queryMerchantUpstreamData();
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
