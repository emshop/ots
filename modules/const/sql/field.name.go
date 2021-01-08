package sql

//ots_product_flow 业务流程的字段信息------------------------------------

//FieldFlowID 字段流程编号的数据库名称
const FieldFlowID = "flow_id"

//FieldFlowName 字段流程名称的数据库名称
const FieldFlowName = "flow_Name"

//FieldPlID 字段产品线编号的数据库名称
const FieldPlID = "pl_id"

//FieldParentFlowID 字段父级流程编号的数据库名称
const FieldParentFlowID = "parent_flow_id"

//FieldSuccessFlowID 字段成功后续流程的数据库名称
const FieldSuccessFlowID = "success_flow_id"

//FieldFailedFlowID 字段失败后续流程的数据库名称
const FieldFailedFlowID = "failed_flow_id"

//FieldUnknownFlowID 字段未知后续流程的数据库名称
const FieldUnknownFlowID = "unknown_flow_id"

//FieldQueueName 字段队列名称的数据库名称
const FieldQueueName = "queue_name"

//FieldScanInterval 字段超时时长的数据库名称
const FieldScanInterval = "scan_interval"

//FieldDelay 字段延后处理时长的数据库名称
const FieldDelay = "delay"

//FieldTimeout 字段超时时长的数据库名称
const FieldTimeout = "timeout"

//FieldMaxCount 字段最大执行次数的数据库名称
const FieldMaxCount = "max_count"

//TradeOrder 订单记录----------------------------

//ots_trade_order 订单记录的字段信息------------------------------------

//FieldOrderID 字段订单编号的数据库名称
const FieldOrderID = "order_id"

//FieldMerNo 字段商户编号的数据库名称
const FieldMerNo = "mer_no"

//FieldMerOrderNo 字段商户订单编号的数据库名称
const FieldMerOrderNo = "mer_order_no"

//FieldMerProductID 字段商品编号的数据库名称
const FieldMerProductID = "mer_product_id"

//FieldMerShelfID 字段货架编号的数据库名称
const FieldMerShelfID = "mer_shelf_id"

//FieldMerProductNo 字段外部商品编号的数据库名称
const FieldMerProductNo = "mer_product_no"

//FieldBrandNo 字段品牌的数据库名称
const FieldBrandNo = "brand_no"

//FieldProvinceNo 字段省份的数据库名称
const FieldProvinceNo = "province_no"

//FieldCityNo 字段城市的数据库名称
const FieldCityNo = "city_no"

//FieldFace 字段商品面值的数据库名称
const FieldFace = "face"

//FieldNum 字段商品数量的数据库名称
const FieldNum = "num"

//FieldTotalFace 字段商品总面值的数据库名称
const FieldTotalFace = "total_face"

//FieldAccountName 字段用户账户信息的数据库名称
const FieldAccountName = "account_name"

//FieldInvoiceType 字段开票方式（1.不开发票）的数据库名称
const FieldInvoiceType = "invoice_type"

//FieldSellDiscount 字段销售折扣的数据库名称
const FieldSellDiscount = "sell_Discount"

//FieldSellAmount 字段总销售金额的数据库名称
const FieldSellAmount = "sell_amount"

//FieldMerFeeAmount 字段商户佣金金额的数据库名称
const FieldMerFeeAmount = "mer_fee_amount"

//FieldTradeFeeAmount 字段交易服务费金额的数据库名称
const FieldTradeFeeAmount = "trade_fee_amount"

//FieldPaymentFeeAmount 字段支付手续费金额的数据库名称
const FieldPaymentFeeAmount = "payment_fee_amount"

//FieldCanSplitOrder 字段是否拆单（0.是，1否）的数据库名称
const FieldCanSplitOrder = "can_split_order"

//FieldCreateTime 字段创建时间的数据库名称
const FieldCreateTime = "create_time"

//FieldOrderTimeout 字段订单超时时间的数据库名称
const FieldOrderTimeout = "order_timeout"

//FieldPaymentTimeout 字段支付超时时间的数据库名称
const FieldPaymentTimeout = "payment_timeout"

//FieldDeliveryPause 字段发货暂停（0.是，1否）的数据库名称
const FieldDeliveryPause = "delivery_pause"

//FieldOrderStatus 字段订单状态的数据库名称
const FieldOrderStatus = "order_status"

//FieldPaymentStatus 字段支付状态的数据库名称
const FieldPaymentStatus = "payment_status"

//FieldDeliveryStatus 字段发货状态的数据库名称
const FieldDeliveryStatus = "delivery_status"

//FieldRefundStatus 字段退款状态的数据库名称
const FieldRefundStatus = "refund_status"

//FieldNotifyStatus 字段通知状态的数据库名称
const FieldNotifyStatus = "notify_status"

//FieldIsRefund 字段用户退款（0.是，1否）的数据库名称
const FieldIsRefund = "is_refund"

//FieldBindFace 字段成功绑定总面值的数据库名称
const FieldBindFace = "bind_face"

//FieldSuccessFace 字段实际成功总面值的数据库名称
const FieldSuccessFace = "success_face"

//FieldSuccessUserAmount 字段实际成功总销售金额 （1）的数据库名称
const FieldSuccessUserAmount = "success_user_amount"

//FieldSuccessMerFee 字段实际成功总佣金金额 （2）的数据库名称
const FieldSuccessMerFee = "success_mer_fee"

//FieldSuccessTradeFee 字段实际成功总服务费金额 （3）的数据库名称
const FieldSuccessTradeFee = "success_trade_fee"

//FieldSuccessPaymentFee 字段实际成功总手续费金额 （4）的数据库名称
const FieldSuccessPaymentFee = "success_payment_fee"

//FieldSuccessCostAmount 字段实际发货成功总成本 （5）的数据库名称
const FieldSuccessCostAmount = "success_cost_amount"

//FieldSuccessSppFee 字段实际发货成功总供货商佣金 （6）的数据库名称
const FieldSuccessSppFee = "success_spp_fee"

//FieldSuccessSppTradeFee 字段实际发货成功总供货商服务费 （7）的数据库名称
const FieldSuccessSppTradeFee = "success_spp_trade_fee"

//FieldProfit 字段利润（1-2-3-4-5add6-7）的数据库名称
const FieldProfit = "profit"

//FieldName 字段名称
const FieldName = "name"

//ots_supplier_product 供货商商品的字段信息------------------------------------

//FieldSppProductID 字段商品编号的数据库名称
const FieldSppProductID = "spp_product_id"

//FieldSppShelfID 字段货架编号的数据库名称
const FieldSppShelfID = "spp_shelf_id"

//FieldSppNo 字段供货商编号的数据库名称
const FieldSppNo = "spp_no"

//FieldSppProductNo 字段供货商商品编号的数据库名称
const FieldSppProductNo = "spp_product_no"

//FieldCostDiscount 字段成本折扣的数据库名称
const FieldCostDiscount = "cost_discount"

//FieldStatus 字段状态的数据库名称
const FieldStatus = "status"

//FieldAssign 指定上游
const FieldAssign = "assign"

//FieldNotifyUrl 字段通知地址的数据库名称
const FieldNotifyUrl = "notify_url"

//ots_trade_delivery 订单发货表的字段信息------------------------------------

//FieldDeliveryID 字段发货编号的数据库名称
const FieldDeliveryID = "delivery_id"

//FieldSppDeliveryNo 字段供货商发货编号的数据库名称
const FieldSppDeliveryNo = "spp_delivery_no"

//FieldCostAmount 字段发货成本的数据库名称
const FieldCostAmount = "cost_amount"

//FieldRealDiscount 字段供货商实际折扣的数据库名称
const FieldRealDiscount = "real_discount"

//FieldSppFeeAmount 字段供货商佣金的数据库名称
const FieldSppFeeAmount = "spp_fee_amount"

//FieldSuccTotalFace 字段成功面值的数据库名称
const FieldSuccTotalFace = "succ_total_face"

//FieldSuccFace 字段成功面值的数据库名称
const FieldSuccFace = "succ_face"

//FieldStartTime 字段开始时间的数据库名称
const FieldStartTime = "start_time"

//FieldEndTime 字段结束时间的数据库名称
const FieldEndTime = "end_time"

//FieldReturnMsg 字段发货返回信息的数据库名称
const FieldReturnMsg = "return_msg"

//FieldRequestParams 字段发货信息参数json的数据库名称
const FieldRequestParams = "request_params"

//FieldResultSource 字段发货结果来源（1：通知，2：查询，3：同步返回）的数据库名称
const FieldResultSource = "result_source"

//FieldResultCode 字段发货结果码的数据库名称
const FieldResultCode = "result_code"

//FieldSppProductCost 字段供货商实际折扣的数据库名称
const FieldSppProductCost = "spp_product_cost"

//FieldLastUpdateTime 字段最后更新时间的数据库名称
const FieldLastUpdateTime = "last_update_time"

//ots_supplier_error_code 供货商错误码的字段信息------------------------------------

//FieldDealCode 字段处理码的数据库名称
const FieldDealCode = "deal_code"

//FieldErrorCode 字段错误码的数据库名称
const FieldErrorCode = "error_code"

//FieldCategory 字段产品线的数据库名称
const FieldCategory = "category"

//FieldNotifyMsg 字段通知结果的数据库名称
const FieldNotifyMsg = "notify_msg"
