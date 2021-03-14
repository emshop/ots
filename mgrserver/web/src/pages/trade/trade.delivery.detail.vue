<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="订单发货表" name="TradeDeliveryDetail">
          <div class="table-responsive">
            <table :date="info" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.delivery_id | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.order_id | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.spp_no | fltrEnum("supplier_info") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商商品编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.spp_product_id |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户商品编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.mer_product_id |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">产品线:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.pl_id | fltrEnum("product_line") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">品牌:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.brand_no | fltrEnum("brand") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">省份:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.province_no | fltrEnum("province") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">城市:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.city_no | fltrEnum("city") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支持开票:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.invoice_type|fltrTextColor">{{ info.invoice_type | fltrEnum("invoice_type") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">用户账户:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.account_name && info.account_name.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.account_name}}</div>
                        <div >{{ info.account_name | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.account_name | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.delivery_status|fltrTextColor">{{ info.delivery_status | fltrEnum("delivery_status") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支付状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.payment_status|fltrTextColor">{{ info.payment_status | fltrEnum("process_status") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.create_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商品面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.face |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货数量:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.num |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货总面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.total_face |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">扣款折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.cost_discount |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">实际折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.real_discount |  fltrNumberFormat(5)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户佣金:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.spp_fee_discount |  fltrNumberFormat(5)}}</div>
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
                      <div class="pull-right" style="margin-right: 10px">发货成本:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.cost_amount |  fltrNumberFormat(5)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商佣金:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.spp_fee_amount |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商服务费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.trade_fee_amount |  fltrNumberFormat(5)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商手续费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.payment_fee_amount |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">成功面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.succ_face |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">开始时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.start_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">完成时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.end_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商发货编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.spp_delivery_no | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商商品编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.spp_product_no | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货结果:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.return_msg && info.return_msg.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.return_msg}}</div>
                        <div >{{ info.return_msg | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.return_msg | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">扩展参数json:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.request_params && info.request_params.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.request_params}}</div>
                        <div >{{ info.request_params | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.request_params | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">结果来源:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.result_source|fltrTextColor">{{ info.result_source | fltrEnum("result_source") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货结果码:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.result_code | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">最后更新时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.last_update_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>            
              </tbody>
            </table>
          </div>
        </el-tab-pane>
        
      </el-tabs>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return {
      tabName: "TradeDeliveryDetail",
      info: {},
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
      this.info = this.$http.xget("/trade/delivery",this.$route.query)
    },
    handleClick(tab) {
      switch (tab.name) {
        case "TradeDeliveryDetail":
          this.queryDetailData();
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
