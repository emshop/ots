<template>
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
                    <div class="pull-right" style="margin-right: 10px">面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.face |  fltrNumberFormat(0)}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">商户商品编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.mer_product_no | fltrEmpty }}</div>
                  </el-col>                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right: 10px">销售折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div>{{ info.discount |  fltrNumberFormat(2)}}</div>
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
	  </el-tabs>
	</div>
</template>

<script>
	export default {
    data(){
      return {
        tabName: "MerchantProductDetail",
        info: {},
      }
    },
    mounted() {
      this.init();
    },
    created(){
    },
    methods: {
      init(){
        this.queryData()
      },
      queryData() {
        this.info = this.$http.xget("/merchant/product",this.$route.query)
      },
      handleClick(tab) {}
    },
	}
</script>