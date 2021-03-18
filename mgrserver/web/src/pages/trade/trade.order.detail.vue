
<template>
  <div>
    <div>
      <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="订单记录" name="TradeOrderDetail">
          <div class="table-responsive">
            <table :date="info" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.order_id | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.mer_no | fltrEnum("merchant_info") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户订单:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="info.mer_order_no && info.mer_order_no.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{info.mer_order_no}}</div>
                        <div >{{ info.mer_order_no | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ info.mer_order_no | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户商品:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.mer_product_id | fltrEnum("merchant_product") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户货架:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ info.mer_shelf_id | fltrEnum("merchant_shelf") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">外部商品编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.mer_product_no | fltrEmpty }}</div>
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
                      <div class="pull-right" style="margin-right: 10px">数量:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.num |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商品总面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.total_face |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
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
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发票:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.invoice_type|fltrTextColor">{{ info.invoice_type | fltrEnum("invoice_type") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">销售折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.sell_discount |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">总销售金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.sell_amount |  fltrNumberFormat(5)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户佣金折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.mer_fee_discount |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">交易服务折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.trade_fee_discount |  fltrNumberFormat(5)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支付手续费折扣:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.payment_fee_discount |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">是否拆单:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.can_split_order|fltrTextColor">{{ info.can_split_order | fltrEnum("bool") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.create_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">完成时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.finish_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单超时时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.order_timeout | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">支付超时时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.payment_timeout | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">绑定面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.bind_face |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">发货暂停:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.delivery_pause|fltrTextColor">{{ info.delivery_pause | fltrEnum("bool") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.order_status|fltrTextColor">{{ info.order_status | fltrEnum("order_status") }}</div>
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
                      <div class="pull-right" style="margin-right: 10px">发货状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.delivery_status|fltrTextColor">{{ info.delivery_status | fltrEnum("process_status") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">退款状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.refund_status|fltrTextColor">{{ info.refund_status | fltrEnum("process_status") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">通知状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.notify_status|fltrTextColor">{{ info.notify_status | fltrEnum("process_status") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">用户退款:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="info.is_refund|fltrTextColor">{{ info.is_refund | fltrEnum("bool") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">成功总面值:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_face |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">成功销售金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_sell_amount |  fltrNumberFormat(2)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户佣金金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_mer_fee |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户服务费金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_mer_trade_fee |  fltrNumberFormat(5)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户手续费金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_mer_payment_fee |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">成本金额:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_cost_amount |  fltrNumberFormat(5)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商佣金:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_spp_fee |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商服务费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_spp_trade_fee |  fltrNumberFormat(5)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">供货商手续费:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.success_spp_payment_fee |  fltrNumberFormat(5)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">利润:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ info.profit |  fltrNumberFormat(5)}}</div>
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
        <el-tab-pane label="订单通知表" name="NotifyInfoDetail">
          <div class="table-responsive">
            <table :date="notifyInfoInfo" class="table table-striped m-b-none">
              <tbody class="table-border">
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ notifyInfoInfo.order_id | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">商户名称:</div>
                    </el-col>
                    <el-col :span="6">
                      <div >{{ notifyInfoInfo.mer_no | fltrEnum("merchant_info") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">订单编号:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="notifyInfoInfo.mer_order_no && notifyInfoInfo.mer_order_no.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{ notifyInfoInfo.mer_order_no}}</div>
                        <div >{{ notifyInfoInfo.mer_order_no | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ notifyInfoInfo.mer_order_no | fltrEmpty }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">通知地址:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="notifyInfoInfo.notify_url && notifyInfoInfo.notify_url.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{ notifyInfoInfo.notify_url}}</div>
                        <div >{{ notifyInfoInfo.notify_url | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ notifyInfoInfo.notify_url | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">创建时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ notifyInfoInfo.create_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">开始时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ notifyInfoInfo.start_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">结束时间:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ notifyInfoInfo.end_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">最大通知次数:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ notifyInfoInfo.max_count |  fltrNumberFormat(0)}}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">通知次数:</div>
                    </el-col>
                    <el-col :span="6">
                      <div>{{ notifyInfoInfo.notify_count |  fltrNumberFormat(0)}}</div>
                    </el-col>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">通知状态:</div>
                    </el-col>
                    <el-col :span="6">
                      <div :class="notifyInfoInfo.notify_status|fltrTextColor">{{ notifyInfoInfo.notify_status | fltrEnum("process_status") }}</div>
                    </el-col>
                  </td>
                </tr>
                <tr>
                  <td>                 
                    <el-col :span="6">
                      <div class="pull-right" style="margin-right: 10px">通知结果:</div>
                    </el-col>
                    <el-col :span="6">
                      <el-tooltip class="item" v-if="notifyInfoInfo.notify_msg && notifyInfoInfo.notify_msg.length > 50" effect="dark" placement="top">
                        <div slot="content" style="width: 110px">{{ notifyInfoInfo.notify_msg}}</div>
                        <div >{{ notifyInfoInfo.notify_msg | fltrSubstr(50) }}</div>
                      </el-tooltip>
                      <div v-else>{{ notifyInfoInfo.notify_msg | fltrEmpty }}</div>
                    </el-col>
                  </td>
                </tr>            
              </tbody>
            </table>
          </div>
        </el-tab-pane>
        <el-tab-pane label="订单发货表" name="TradeDeliveryDetail">
          <el-scrollbar style="height:100%" id="panel-body">
            <el-table :data="TradeDeliveryList.items" stripe style="width: 100%" :height="maxHeight">
              
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
                  <el-tooltip class="item" v-if="scope.row.account_name && scope.row.account_name.length > 20" effect="dark" placement="top">
                    <div slot="content" style="width: 110px">{{scope.row.account_name}}</div>
                    <span>{{scope.row.account_name | fltrSubstr(20) }}</span>
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
                  <el-tooltip class="item" v-if="scope.row.return_msg && scope.row.return_msg.length > 20" effect="dark" placement="top">
                    <div slot="content" style="width: 110px">{{scope.row.return_msg}}</div>
                    <span>{{scope.row.return_msg | fltrSubstr(20) }}</span>
                  </el-tooltip>
                  <span v-else>{{scope.row.return_msg | fltrEmpty }}</span>
                </template>
              </el-table-column>
            </el-table>
          </el-scrollbar>
        </el-tab-pane>
        <el-tab-pane label="资金变动" name="AccountRecordDetail">
          <el-scrollbar style="height:100%" id="panel-body">
            <el-table :data="AccountRecordList.items" stripe style="width: 100%" :height="maxHeight">
              
              <el-table-column   prop="record_id" label="变动编号" align="center">
              <template slot-scope="scope">
                <span>{{scope.row.record_id | fltrEmpty }}</span>
              </template>
              
              </el-table-column>
              <el-table-column   prop="account_id" label="帐户编号" align="center">
                <template slot-scope="scope">
                  <span >{{scope.row.account_id | fltrEnum("account_info")}}</span>
                </template>
              </el-table-column>
              <el-table-column   prop="trade_no" label="交易编号" align="center">
              <template slot-scope="scope">
                <span>{{scope.row.trade_no | fltrEmpty }}</span>
              </template>
              
              </el-table-column>
              <el-table-column   prop="trade_type" label="交易类型" align="center">
                <template slot-scope="scope">
                  <span :class="scope.row.trade_type|fltrTextColor">{{scope.row.trade_type | fltrEnum("trade_type")}}</span>
                </template>
              </el-table-column>
              <el-table-column   prop="change_type" label="变动类型" align="center">
                <template slot-scope="scope">
                  <span :class="scope.row.change_type|fltrTextColor">{{scope.row.change_type | fltrEnum("change_type")}}</span>
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
                <div>{{scope.row.create_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
              </template>
              </el-table-column>
            </el-table>
          </el-scrollbar>
        </el-tab-pane>
        <el-tab-pane label="生命周期" name="LifeTimeDetail">
          <el-scrollbar style="height:100%" id="panel-body">
            <el-table :data="LifeTimeList.items" stripe style="width: 100%" :height="maxHeight">
              
              <el-table-column   prop="id" label="序号" align="center">
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
                  <el-tooltip class="item" v-if="scope.row.content && scope.row.content.length > 48" effect="dark" placement="top">
                    <div slot="content" style="width: 110px">{{scope.row.content}}</div>
                    <span>{{scope.row.content | fltrSubstr(48) }}</span>
                  </el-tooltip>
                  <span v-else>{{scope.row.content | fltrEmpty }}</span>
                </template>
              </el-table-column>
              <el-table-column   prop="create_time" label="创建时间" align="center">
              <template slot-scope="scope">
                <div>{{scope.row.create_time | fltrDate("yyyy-MM-dd HH:mm:ss") }}</div>
              </template>
              </el-table-column>
            </el-table>
          </el-scrollbar>
        </el-tab-pane>
        <el-tab-pane label="任务表" name="SystemTaskDetail">
          <el-scrollbar style="height:100%" id="panel-body">
            <el-table :data="SystemTaskList.items" stripe style="width: 100%" :height="maxHeight">
              
              <el-table-column   prop="task_id" label="编号" align="center">
              <template slot-scope="scope">
                <span>{{scope.row.task_id | fltrEmpty }}</span>
              </template>
              
              </el-table-column>
              <el-table-column   prop="order_no" label="订单号" align="center">
              <template slot-scope="scope">
                <span>{{scope.row.order_no | fltrEmpty }}</span>
              </template>
              
              </el-table-column>
              <el-table-column   prop="name" label="流程名称" align="center">
                <template slot-scope="scope">
                  <span >{{scope.row.name | fltrEnum("flow_tag")}}</span>
                </template>
              </el-table-column>
              <el-table-column   prop="create_time" label="创建时间" align="center">
              <template slot-scope="scope">
                <div>{{scope.row.create_time | fltrDate("MM/dd HH:mm:ss") }}</div>
              </template>
              </el-table-column>
              <el-table-column   prop="last_execute_time" label="上次执行时间" align="center">
              <template slot-scope="scope">
                <div>{{scope.row.last_execute_time | fltrDate("MM/dd HH:mm:ss") }}</div>
              </template>
              </el-table-column>
              <el-table-column   prop="next_execute_time" label="下次执行时间" align="center">
              <template slot-scope="scope">
                <div>{{scope.row.next_execute_time | fltrDate("MM/dd HH:mm:ss") }}</div>
              </template>
              </el-table-column>
              <el-table-column   prop="count" label="执行次数" align="center">
              <template slot-scope="scope">
                <span>{{scope.row.count | fltrNumberFormat(0)}}</span>
              </template>
              </el-table-column>
              <el-table-column   prop="queue_name" label="消息队列" align="center">
                <template slot-scope="scope">
                  <el-tooltip class="item" v-if="scope.row.queue_name && scope.row.queue_name.length > 20" effect="dark" placement="top">
                    <div slot="content" style="width: 110px">{{scope.row.queue_name}}</div>
                    <span>{{scope.row.queue_name | fltrSubstr(20) }}</span>
                  </el-tooltip>
                  <span v-else>{{scope.row.queue_name | fltrEmpty }}</span>
                </template>
              </el-table-column>
              <el-table-column   prop="status" label="状态" align="center">
                <template slot-scope="scope">
                  <span :class="scope.row.status|fltrTextColor">{{scope.row.status | fltrEnum("process_status")}}</span>
                </template>
              </el-table-column>
            </el-table>
          </el-scrollbar>
        </el-tab-pane>
        
      </el-tabs>
    </div>   
    <div class="page-pagination" v-show="tabName =='TradeDeliveryDetail'">
    <el-pagination
      @size-change="pageTradeDeliverySizeChange"
      @current-change="pageTradeDeliveryIndexChange"
      :current-page="pagingTradeDelivery.pi"
      :page-size="pagingTradeDelivery.ps"
      :page-sizes="pagingTradeDelivery.sizes"
      layout="total, sizes, prev, pager, next, jumper"
      :total="TradeDeliveryList.count">
    </el-pagination>
    </div>   
    <div class="page-pagination" v-show="tabName =='AccountRecordDetail'">
    <el-pagination
      @size-change="pageAccountRecordSizeChange"
      @current-change="pageAccountRecordIndexChange"
      :current-page="pagingAccountRecord.pi"
      :page-size="pagingAccountRecord.ps"
      :page-sizes="pagingAccountRecord.sizes"
      layout="total, sizes, prev, pager, next, jumper"
      :total="AccountRecordList.count">
    </el-pagination>
    </div>   
    <div class="page-pagination" v-show="tabName =='LifeTimeDetail'">
    <el-pagination
      @size-change="pageLifeTimeSizeChange"
      @current-change="pageLifeTimeIndexChange"
      :current-page="pagingLifeTime.pi"
      :page-size="pagingLifeTime.ps"
      :page-sizes="pagingLifeTime.sizes"
      layout="total, sizes, prev, pager, next, jumper"
      :total="LifeTimeList.count">
    </el-pagination>
    </div>   
    <div class="page-pagination" v-show="tabName =='SystemTaskDetail'">
    <el-pagination
      @size-change="pageSystemTaskSizeChange"
      @current-change="pageSystemTaskIndexChange"
      :current-page="pagingSystemTask.pi"
      :page-size="pagingSystemTask.ps"
      :page-sizes="pagingSystemTask.sizes"
      layout="total, sizes, prev, pager, next, jumper"
      :total="SystemTaskList.count">
    </el-pagination>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return {
      tabName: "TradeOrderDetail",
      info: {},
      notifyInfoInfo:{},
      pagingTradeDelivery: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
      TradeDeliveryList: {count: 0,items: []}, //表单数据对象,
      queryTradeDeliveryParams:{},  //查询数据对象
      pagingAccountRecord: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
      AccountRecordList: {count: 0,items: []}, //表单数据对象,
      queryAccountRecordParams:{},  //查询数据对象
      pagingLifeTime: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
      LifeTimeList: {count: 0,items: []}, //表单数据对象,
      queryLifeTimeParams:{},  //查询数据对象
      pagingSystemTask: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
      SystemTaskList: {count: 0,items: []}, //表单数据对象,
      querySystemTaskParams:{},  //查询数据对象
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
      this.info = this.$http.xget("/trade/order",this.$route.query)
    },
    queryNotifyInfoData() {
      this.notifyInfoInfo = this.$http.xget("/notify/info/detail",{ order_id: this.info.order_id })
    },
    /**查询数据并赋值*/
		queryTradeDeliveryDatas() {
      this.pagingTradeDelivery.pi = 1
      this.queryTradeDeliveryData()
    },
    queryTradeDeliveryData(){
      this.queryTradeDeliveryParams.pi = this.pagingTradeDelivery.pi
			this.queryTradeDeliveryParams.ps = this.pagingTradeDelivery.ps
      var data = this.$utility.delEmptyProperty(this.queryTradeDeliveryParams)
      data.order_id = this.info.order_id || ""
      let res = this.$http.xpost("/trade/delivery/querydetail", data)
			this.TradeDeliveryList.items = res.items || []
			this.TradeDeliveryList.count = res.count
    },
    /**改变页容量*/
		pageTradeDeliverySizeChange(val) {
      this.pagingTradeDelivery.ps = val
      this.queryTradeDeliveryData()
    },
    /**改变当前页码*/
    pageTradeDeliveryIndexChange(val) {
      this.pagingTradeDelivery.pi = val
      this.queryTradeDeliveryData()
    },
    /**查询数据并赋值*/
		queryAccountRecordDatas() {
      this.pagingAccountRecord.pi = 1
      this.queryAccountRecordData()
    },
    queryAccountRecordData(){
      this.queryAccountRecordParams.pi = this.pagingAccountRecord.pi
			this.queryAccountRecordParams.ps = this.pagingAccountRecord.ps
      var data = this.$utility.delEmptyProperty(this.queryAccountRecordParams)
      data.trade_no = this.info.order_id || ""
      let res = this.$http.xpost("/account/record/querydetail", data)
			this.AccountRecordList.items = res.items || []
			this.AccountRecordList.count = res.count
    },
    /**改变页容量*/
		pageAccountRecordSizeChange(val) {
      this.pagingAccountRecord.ps = val
      this.queryAccountRecordData()
    },
    /**改变当前页码*/
    pageAccountRecordIndexChange(val) {
      this.pagingAccountRecord.pi = val
      this.queryAccountRecordData()
    },
    /**查询数据并赋值*/
		queryLifeTimeDatas() {
      this.pagingLifeTime.pi = 1
      this.queryLifeTimeData()
    },
    queryLifeTimeData(){
      this.queryLifeTimeParams.pi = this.pagingLifeTime.pi
			this.queryLifeTimeParams.ps = this.pagingLifeTime.ps
      var data = this.$utility.delEmptyProperty(this.queryLifeTimeParams)
      data.order_no = this.info.order_id || ""
      let res = this.$http.xpost("/life/time/querydetail", data)
			this.LifeTimeList.items = res.items || []
			this.LifeTimeList.count = res.count
    },
    /**改变页容量*/
		pageLifeTimeSizeChange(val) {
      this.pagingLifeTime.ps = val
      this.queryLifeTimeData()
    },
    /**改变当前页码*/
    pageLifeTimeIndexChange(val) {
      this.pagingLifeTime.pi = val
      this.queryLifeTimeData()
    },
    /**查询数据并赋值*/
		querySystemTaskDatas() {
      this.pagingSystemTask.pi = 1
      this.querySystemTaskData()
    },
    querySystemTaskData(){
      this.querySystemTaskParams.pi = this.pagingSystemTask.pi
			this.querySystemTaskParams.ps = this.pagingSystemTask.ps
      var data = this.$utility.delEmptyProperty(this.querySystemTaskParams)
      data.order_no = this.info.order_id || ""
      let res = this.$http.xpost("/system/task/querydetail", data)
			this.SystemTaskList.items = res.items || []
			this.SystemTaskList.count = res.count
    },
    /**改变页容量*/
		pageSystemTaskSizeChange(val) {
      this.pagingSystemTask.ps = val
      this.querySystemTaskData()
    },
    /**改变当前页码*/
    pageSystemTaskIndexChange(val) {
      this.pagingSystemTask.pi = val
      this.querySystemTaskData()
    },
    handleClick(tab) {
      switch (tab.name) {
        case "TradeOrderDetail":
          this.queryDetailData();
          break;
        case "NotifyInfoDetail":
          this.queryNotifyInfoData();
          break;
        case "TradeDeliveryDetail":
          this.queryTradeDeliveryData();
          break;
        case "AccountRecordDetail":
          this.queryAccountRecordData();
          break;
        case "LifeTimeDetail":
          this.queryLifeTimeData();
          break;
        case "SystemTaskDetail":
          this.querySystemTaskData();
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
