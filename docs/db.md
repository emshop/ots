PK:主键
SEQ:SEQ[(起始值,[大小])]，mysql自增，oracle序列
UNQ:UNQ[(名称[,位置])],唯一约束
IDX:索引,IDX[(名称[,位置])]
C: 创建数据时的字段
R: 单条数据读取时的字段 
U: 修改数据时需要的字段
D: 删除，仅在主键字段使用
Q: 查询条件的字段
L：(前端页面)列表里列出的字段
OB：查询时的order by字段；默认降序； OB[(顺序)]，越小越先排序
DI: 字典编号，数据表作为字典数据时的id字段
DN: 字典名称，数据表作为字典数据时的name字段
IN:        "input"       //表单输入框
SL: "select"      //表单下拉框,默认使用dds字典表枚举,指定表名的SL[(字典数据表)]
CB: "checkbox"    //表单复选框,默认使用dds字典表枚举,指定表名的CB[(字典数据表)]
RB: "radio"       //表单单选框,默认使用dds字典表枚举,指定表名的RB[(字典数据表)]
TA: "textarea"    //表单文本域
DT: "date-picker" //表单日期选择器
After: After(字段名) //在某个字段后面


## 一、商户信息

###  1. 商户信息[ots_merchant_info]

| 字段名      | 类型         | 默认值  | 为空  |           约束            | 描述     |
| ----------- | ------------ | :-----: | :---: | :-----------------------: | :------- |
| mer_no      | varchar2(32) |         |  否   |        PK,l,r,c,DI        | 编号     |
| mer_name    | varchar2(64) |         |  否   |       q,l,r,u,c,DN        | 商户名称 |
| mer_crop    | varchar2(64) |         |  是   |          l,u,c,r          | 公司名称 |
| mer_type    | number(1)    |         |  否   | l,u,c,r,sl(merchant_type) | 类型     |
| bd_uid      | number(20)   |    0    |  是   |    r,u,c,sl(user_info)    | 商务人员 |
| status      | number(1)    |    0    |  否   |    q,l,r,u,c,sl,cc,rd     | 状态     |
| create_time | date         | sysdate |  否   |            r,l            | 创建时间 |


###  2. 商户货架[ots_merchant_shelf]

| 字段名               | 类型         | 默认值  | 为空  |             约束              | 描述                   |
| -------------------- | ------------ | :-----: | :---: | :---------------------------: | :--------------------- |
| mer_shelf_id         | number(10)   |   100   |  否   |         PK,SEQ,DI,l,r         | 编号                   |
| mer_shelf_name       | varchar2(64) |         |  否   |         DN,q,l,u,r,c          | 货架名称               |
| mer_no               | varchar2(32) |         |  否   | q,l,r,sl(ots_merchant_info),c | 商户名称               |
| mer_fee_discount     | number(10,5) |    0    |  否   |      l,u,r,decimal(5),c       | 商户佣金               |
| trade_fee_discount   | number(10,5) |    0    |  否   |      l,u,r,decimal(5),c       | 交易服务费             |
| payment_fee_discount | number(10,5) |    0    |  否   |      l,u,r,decimal(5),c       | 支付手续费             |
| order_timeout        | number(10)   |         |  否   |            l,u,r,c            | 订单超时时长           |
| payment_timeout      | number(10)   |         |  否   |             u,r,c             | 支付超时时长           |
| invoice_type         | number(2)    |    1    |  否   |  r ,u,cc,sl(invoice_type),c   | 开票方式（1.不开发票） |
| can_refund           | number(1)    |    1    |  否   |       r,u,sl(bool),cc,c       | 允许退款(0.是,1否)     |
| limit_count          | number(10)   |    1    |  否   |             u,r,c             | 单次购买数量           |
| can_split_order      | number(1)    |    0    |  否   |      l,r,u,sl(bool),cc,c      | 允许拆单               |
| status               | number(1)    |    0    |  否   |         l,r,u,cc,sl,c         | 状态                   |
| create_time          | date         | sysdate |  否   |              l,r              | 创建时间               |

###  3.商户商品[ots_merchant_product]

| 字段名         | 类型         | 默认值  | 为空  |                          约束                           | 描述                 |
| -------------- | ------------ | :-----: | :---: | :-----------------------------------------------------: | :------------------- |
| mer_product_id | number(10)   |  10000  |  否   |                       PK,SEQ,l,r                        | 商品编号             |
| mer_shelf_id   | number(10)   |         |  否   | q,l,r,UNQ(unq_mer_prod,7),sl(ots_merchant_shelf),c,sort | 货架名称             |
| mer_no         | varchar2(32) |         |  否   | q,l,r,UNQ(unq_mer_prod,1),sl(ots_merchant_info),c,sort  | 商户名称             |
| pl_id          | number(10)   |         |  否   |  q,l,r,UNQ(unq_mer_prod,2),sl(ots_product_line),c,sort  | 产品线               |
| brand_no       | varchar2(8)  |         |  否   |       q,l,r,UNQ(unq_mer_prod,3),sl(brand),c,sort        | 品牌                 |
| province_no    | varchar2(8)  |   '-'   |  否   |       l,r,UNQ(unq_mer_prod,4),sl(province),c,sort       | 省份                 |
| city_no        | varchar2(8)  |   '-'   |  否   |         l,r,UNQ(unq_mer_prod,5),sl(city),c,sort         | 城市                 |
| face           | number(10)   |         |  否   |             l,r,UNQ(unq_mer_prod,6),c,sort              | 面值                 |
| mer_product_no | varchar2(32) |         |  是   |                          u,r,c                          | 商户商品编号         |
| discount       | number(10,5) |         |  否   |                 l,u,r,decimal(5),c,sort                 | 销售折扣（以面值算） |
| status         | number(1)    |    0    |  否   |                   l,r,u,sl,cc,c,sort                    | 状态(0.是,1.否)      |
| create_time    | date         | sysdate |  否   |             l,r(f:yyyy-MM-dd HH:mm:ss),sort             | 创建时间             |


## 二、供货商信息

###  1. 供货商信息[ots_supplier_info]

| 字段名      | 类型         | 默认值  | 为空  |        约束        | 描述       |
| :---------- | :----------- | :-----: | :---: | :----------------: | :--------- |
| spp_no      | varchar2(32) |         |  否   |    PK,l,r,DI,c     | 编号       |
| spp_name    | varchar2(64) |         |  否   |    q,l,r,DN,c,u    | 供货商名称 |
| mer_crop    | varchar2(64) |         |  是   |      l,r,c,u       | 所属公司   |
| bd_uid      | number(20)   |    0    |  是   |       r,c,u        | 商务人员   |
| status      | number(1)    |    0    |  是   | q,l,r,sl,cc,c,u,rd | 状态       |
| create_time | date         | sysdate |  是   |        l,r         | 创建时间   |


###  2. 供货商货架[ots_supplier_shelf]

| 字段名               | 类型          | 默认值  | 为空  |             约束              | 描述                 |
| -------------------- | ------------- | :-----: | :---: | :---------------------------: | :------------------- |
| spp_shelf_id         | number(10)    |   600   |  否   |        PK,SEQ,l,r,DI,c        | 货架编号             |
| spp_shelf_name       | varchar2(64)  |         |  否   |         q,l,r,DN,c,u          | 货架名称             |
| spp_no               | varchar2(32)  |         |  否   | l,r,q,sl(ots_supplier_info),c | 供货商               |
| req_url              | varchar2(128) |         |  否   |             r,c,u             | 请求地址             |
| query_url            | varchar2(128) |         |  是   |             r,c,u             | 查询地址             |
| notify_url           | varchar2(128) |         |  是   |             r,c,u             | 回调地址             |
| invoice_type         | number(2)     |    1    |  是   | l,r,sl(invoice_type),cc ,c,u  | 开票                 |
| spp_fee_discount     | number(10,5)  |    0    |  是   |      l,r,decimal(5),c,u       | 商户佣金             |
| trade_fee_discount   | number(10,5)  |    0    |  是   |      l,r,decimal(5),c,u       | 交易服务费           |
| payment_fee_discount | number(10,5)  |    0    |  是   |      l,r,decimal(5),c,u       | 支付手续费           |
| can_refund           | number(1)     |    1    |  是   |      l,r,sl(bool),cc,c,u      | 支持退货 (0.是,1.否) |
| status               | number(1)     |    0    |  是   |        q,l,r,sl,cc,c,u        | 货架状态             |
| limit_count          | number(10)    |    1    |  是   |             r,c,u             | 单次最大发货数量     |
| create_time          | date          | sysdate |  是   |              l,r              | 创建时间             |

###  3. 供货商商品[ots_supplier_product]

| 字段名         | 类型         | 默认值  | 为空  |              约束              | 描述           |
| -------------- | ------------ | :-----: | :---: | :----------------------------: | :------------- |
| spp_product_id | number(10)   |  60000  |  否   |           PK,SEQ,l,r           | 商品编号       |
| spp_shelf_id   | number(10)   |         |  否   | l,r,q,sl(ots_supplier_shelf),c | 货架名称       |
| spp_no         | varchar2(32) |         |  否   | l,r,q,sl(ots_supplier_info),c  | 供货商         |
| spp_product_no | varchar2(32) |         |  是   |             r,c,u              | 供货商商品编号 |
| pl_id          | number(10)   |         |  否   |  l,r,q,sl(ots_product_line),c  | 产品线         |
| brand_no       | varchar2(8)  |         |  否   |       l,q,r,sl(brand),c        | 品牌           |
| province_no    | varchar2(8)  |    *    |  是   |      l,r,q,sl(province),c      | 省份           |
| city_no        | varchar2(8)  |    *    |  是   |        l,r,sl(city) ,c         | 城市           |
| face           | number(10)   |         |  否   |             l,r,c              | 面值           |
| cost_discount  | number(10,5) |         |  否   |       l,r,c,decimal(5),u       | 成本折扣       |
| status         | number(1)    |    0    |  否   |         l,r,sl,cc,c,u          | 状态           |
| create_time    | date         | sysdate |  是   |              l,r               | 创建时间       |



###  4. 供货商错误码[ots_supplier_ecode]

| 字段名      | 类型         | 默认值  | 为空  |                          约束                           | 描述       |
| ----------- | ------------ | :-----: | :---: | :-----------------------------------------------------: | :--------- |
| id          | number(20)   |   100   |  否   |                     PK,SEQ,l,r,u,c                      | 编号       |
| spp_no      | varchar2(32) |         |  否   | q,l,r,UNQ(unq_spp_error_code),sl(ots_supplier_info),u,c | 供货商     |
| pl_id       | number(10)   |         |  否   | q,l,r,UNQ(unq_spp_error_code),sl(ots_product_line),u,c  | 产品线     |
| category    | varchar2(32) |         |  否   | q,l,r,UNQ(unq_spp_error_code),sl(result_source),cc,u,c  | 分类       |
| deal_code   | number(1)    |         |  否   |                l,r,u,c,sl(deal_code),cc                 | 处理码     |
| error_code  | varchar2(32) |         |  否   |             l,r,UNQ(unq_spp_error_code),u,c             | 错误码     |
| status      | number(1)    |    0    |  否   |                    l,r,s,q,cc,sl,u,c                    | 状态       |
| error_desc  | varchar2(64) |         |  否   |                         l,r,u,c                         | 错误码描述 |
| create_time | date         | sysdate |  否   |                           l,r                           | 创建时间   |


# 三、交易类

###  1. 订单记录[ots_trade_order]{tab(ots_trade_delivery|)}

| 字段名                  | 类型         |   默认值   | 为空  |                             约束                             | 描述                  |
| ----------------------- | ------------ | :--------: | :---: | :----------------------------------------------------------: | :-------------------- |
| order_id                | varchar2(32) | 1100000000 |  否   |                     PK,l,q,r,sort,fixed                      | 订单编号              |
| mer_no                  | varchar2(32) |            |  否   |                 l,Q,sl(ots_merchant_info),r                  | 商户编号              |
| mer_order_no            | varchar2(64) |            |  否   |                             r,l                              | 商户订单编号          |
| mer_product_id          | number(10)   |    300     |  否   |                  r,sl(ots_merchant_product)                  | 商户商品              |
| mer_shelf_id            | number(10)   |            |  否   |                   r,sl(ots_merchant_shelf)                   | 商户货架              |
| mer_product_no          | varchar2(32) |            |  是   |                              r                               | 外部商品编号          |
| pl_id                   | number(10)   |            |  否   |                  l,Q,sl(ots_product_line),r                  | 产品线                |
| brand_no                | varchar2(8)  |            |  否   |                        l,r,sl(brand)                         | 品牌                  |
| province_no             | varchar2(8)  |            |  否   |                      q,l,r,sl(province)                      | 省份                  |
| city_no                 | varchar2(8)  |            |  否   |                 r,q(e:#province_no),sl(city)                 | 城市                  |
| face                    | number(10)   |            |  否   |                             l,r                              | 面值                  |
| num                     | number(10)   |            |  否   |                             l,r                              | 数量                  |
| total_face              | number(10)   |            |  否   |                              r                               | 商品总面值            |
| account_name            | varchar2(64) |            |  否   |                             l,r                              | 用户账户              |
| invoice_type            | number(2)    |            |  否   |                    r,cc,sl(invoice_type)                     | 发票（1.不支持）      |
| sell_discount           | number(20,5) |            |  否   |                        l,r,decimal(5)                        | 销售折扣              |
| sell_amount             | number(20,5) |            |  否   |                              r                               | 总销售金额            |
| mer_fee_discount        | number(20,5) |            |  否   |                         r,decimal(5)                         | 商户佣金折扣          |
| trade_fee_discount      | number(20,5) |            |  否   |                         r,decimal(5)                         | 交易服务折扣          |
| payment_fee_discount    | number(20,5) |            |  否   |                         r,decimal(5)                         | 支付手续费折扣        |
| can_split_order         | number(1)    |     1      |  否   |                        r,sl(bool),cc                         | 是否拆单（0.是，1否） |
| create_time             | date         |  sysdate   |  否   | q(f:yyyy-MM-dd),l(f:MM/dd HH:mm:ss),r(f:yyyy-MM-dd HH:mm:ss) | 创建时间              |
| finish_time             | date         |  sysdate   |  否   |                   r(f:yyyy-MM-dd HH:mm:ss)                   | 完成时间              |
| order_timeout           | date         |            |  否   |                  r,r(f:yyyy-MM-dd HH:mm:ss)                  | 订单超时时间          |
| payment_timeout         | date         |            |  否   |                  r,r(f:yyyy-MM-dd HH:mm:ss)                  | 支付超时时间          |
| bind_face               | number(10)   |     0      |  否   |                             r,l                              | 绑定面值              |
| delivery_pause          | number(1)    |     1      |  否   |                        r,sl(bool),cc                         | 发货暂停（0.是，1否） |
| order_status            | number(3)    |     10     |  否   |                          l,r,sl,cc                           | 订单状态              |
| payment_status          | number(3)    |     20     |  否   |                   r,sl(process_status),cc                    | 支付状态              |
| delivery_status         | number(3)    |     10     |  否   |                   r,sl(process_status),cc                    | 发货状态              |
| refund_status           | number(3)    |     10     |  否   |                   r,sl(process_status),cc                    | 退款状态              |
| notify_status           | number(3)    |     10     |  否   |                   r,sl(process_status),cc                    | 通知状态              |
| is_refund               | number(1)    |     1      |  否   |                        r,sl(bool),cc                         | 用户退款（0.是，1否） |
| success_face            | number(10)   |     0      |  否   |                              r                               | 成功总面值            |
| success_sell_amount     | number(20,5) |     0      |  否   |                              r                               | 成功销售金额 （1）    |
| success_mer_fee         | number(20,5) |     0      |  否   |                         r,decimal(3)                         | 商户佣金金额 （2）    |
| success_mer_trade_fee   | number(20,5) |     0      |  否   |                         r,decimal(3)                         | 商户服务费金额 （3）  |
| success_mer_payment_fee | number(20,5) |     0      |  否   |                         r,decimal(3)                         | 商户手续费金额 （4）  |
| success_cost_amount     | number(20,5) |     0      |  否   |                         r,decimal(3)                         | 成本金额 （5）        |
| success_spp_fee         | number(20,5) |     0      |  否   |                         r,decimal(3)                         | 供货商佣金 （6）      |
| success_spp_trade_fee   | number(20,5) |     0      |  否   |                         r,decimal(3)                         | 供货商服务费 （7）    |
| success_spp_payment_fee | number(20,5) |     0      |  否   |                         r,decimal(3)                         | 供货商手续费 （8）    |
| profit                  | number(20,5) |     0      |  否   |                         r,decimal(3)                         | 利润(1-2-3-4-5+6-7-8) |



###  2. 订单发货表[ots_trade_delivery]
| 字段名               | 类型           | 默认值  | 为空  |             约束             | 描述                                      |
| -------------------- | -------------- | :-----: | :---: | :--------------------------: | :---------------------------------------- |
| delivery_id          | varchar2(32)   |  20000  |  否   |            PK,r,l            | 发货编号                                  |
| order_id             | varchar2(32)   |         |  否   |             r,l              | 订单编号                                  |
| spp_no               | varchar2(32)   |         |  否   | r,q,l,sl(ots_supplier_info)  | 供货商                                    |
| spp_product_id       | number(10)     |         |  否   |              r               | 供货商商品编号                            |
| mer_no               | varchar2(32)   |         |  否   |   q,sl(ots_merchant_info)    | 商户编号                                  |
| mer_product_id       | number(10)     |         |  否   |              r               | 商户商品编号                              |
| pl_id                | number(10)     |         |  否   | q,r,l,sl(ots_product_line),r | 产品线                                    |
| brand_no             | varchar2(8)    |         |  否   |       q,r,l,sl(brand)        | 品牌                                      |
| province_no          | varchar2(8)    |         |  否   |       r,l,sl(province)       | 省份                                      |
| city_no              | varchar2(8)    |         |  否   |          r,sl(city)          | 城市                                      |
| invoice_type         | number(3)      |         |  否   |    r,cc,sl(invoice_type)     | 支持开票（1.不支持）                      |
| account_name         | varchar2(64)   |         |  否   |             r,l              | 用户账户                                  |
| delivery_status      | number(3)      |   20    |  否   |  r,l,sl(delivery_status),cc  | 发货状态                                  |
| payment_status       | number(3)      |   10    |  否   |   r,sl(process_status),cc    | 支付状态                                  |
| create_time          | date           | sysdate |  否   |     r,l,q(f:yyyy-MM-dd)      | 创建时间                                  |
| face                 | number(10)     |         |  否   |             r,l              | 商品面值                                  |
| num                  | number(10)     |         |  否   |              r               | 发货数量                                  |
| total_face           | number(10)     |         |  否   |              r               | 发货总面值                                |
| cost_discount        | number(20,5)   |         |  否   |         r,decimal(5)         | 扣款折扣                                  |
| real_discount        | number(20,5)   |         |  是   |         r,decimal(5)         | 实际折扣                                  |
| spp_fee_discount     | number(10,5)   |    0    |  否   |         r,decimal(5)         | 商户佣金                                  |
| trade_fee_discount   | number(10,5)   |    0    |  否   |         r,decimal(5)         | 交易服务费                                |
| payment_fee_discount | number(10,5)   |    0    |  否   |         r,decimal(5)         | 支付手续费                                |
| cost_amount          | number(20,5)   |    0    |  否   |         r,decimal(3)         | 发货成本                                  |
| spp_fee_amount       | number(20,5)   |    0    |  否   |        r ,decimal(3)         | 供货商佣金                                |
| trade_fee_amount     | number(20,5)   |    0    |  否   |         r,decimal(3)         | 供货商服务费                              |
| payment_fee_amount   | number(20,5)   |    0    |  否   |        r ,decimal(3)         | 供货商手续费                              |
| succ_face            | number(10)     |    0    |  是   |              r               | 成功面值                                  |
| start_time           | date           |         |  是   |             r,l              | 开始时间                                  |
| end_time             | date           |         |  是   |              r               | 结束时间                                  |
| spp_delivery_no      | varchar2(32)   |         |  是   |              r               | 供货商发货编号                            |
| spp_product_no       | varchar2(32)   |         |  是   |              r               | 供货商商品编号                            |
| return_msg           | varchar2(256)  |         |  是   |             r,l              | 发货结果                                  |
| request_params       | varchar2(2000) |         |  是   |              r               | 扩展参数json                              |
| result_source        | varchar2(32)   |         |  是   |    r,sl(result_source),cc    | 结果来源（1：通知，2：查询，3：同步返回） |
| result_code          | varchar2(32)   |         |  是   |              r,              | 发货结果码                                |
| last_update_time     | date           | sysdate |  否   |              r               | 最后更新时间                              |



###  4. 退款申请[ots_refund_apply]

| 字段名        | 类型         | 默认值  | 为空  |            约束             | 描述       |
| ------------- | ------------ | :-----: | :---: | :-------------------------: | :--------- |
| apply_id      | number(20)   |   100   |  否   |           PK,r,l            | 申请编号   |
| order_id      | varchar2(32) |         |  否   |            q,r,l            | 订单编号   |
| mer_no        | varchar2(32) |         |  否   | q,r,l,sl(ots_merchant_info) | 商户编号   |
| mer_order_no  | varchar2(32) |         |  否   |            q,r,l            | 商户订单号 |
| refund_cause  | number(3)    |   10    |  否   |             r,l             | 退款原因   |
| apply_status  | number(2)    |   10    |  否   |    r,sl(process_status)     | 申请状态   |
| refund_amount | number(20,5) |         |  否   |             r,l             | 退款金额   |
| create_time   | date         | sysdate |  否   |             r,l             | 创建时间   |



###  5.订单通知表[ots_notify_info]

| 字段名        | 类型          | 默认值  | 为空  |            约束             | 描述                                             |
| ------------- | ------------- | :-----: | :---: | :-------------------------: | :----------------------------------------------- |
| order_id      | varchar2(32)  |         |  否   |           PK,l,r            | 订单编号                                         |
| mer_no        | varchar2(32)  |         |  否   | l,r,q,sl(ots_merchant_info) | 商户名称                                         |
| mer_order_no  | varchar2(64)  |         |  否   |             l,r             | 订单编号                                         |
| notify_url    | varchar2(128) |         |  否   |              r              | 通知地址                                         |
| notify_status | number(3)     |   10    |  否   | l,r,q,sl(process_status),cc | 通知状态（0成功,10未开始,20等待通知,30正在通知） |
| max_count     | number(3)     |   10    |  否   |            r,cc             | 最大通知次数                                     |
| notify_count  | number(3)     |    0    |  否   |             l,r             | 通知次数                                         |
| create_time   | date          | sysdate |  否   |           l,r,dp            | 创建时间                                         |
| start_time    | date          |         |  是   |             l,r             | 开始时间                                         |
| end_time      | date          |         |  是   |             l,r             | 结束时间                                         |
| notify_msg    | varchar2(256) |         |  是   |             l,r             | 通知结果                                         |


###  6. 发货人工审核表[ots_audit_info]

| 字段名           | 类型          | 默认值  | 为空  |        约束         | 描述           |
| ---------------- | ------------- | :-----: | :---: | :-----------------: | :------------- |
| delivery_id      | varchar2(32)  |         |  是   |       PK,l,r        | 发货编号       |
| order_id         | varchar2(32)  |         |  否   |        q,l,r        | 订单编号       |
| create_time      | date          | sysdate |  否   |         l,r         | 创建时间       |
| delivery_status  | number(2)     |         |  否   |         l,r         | 发货结果       |
| end_order        | number(1)     |         |  否   |    r,cc,sl(bool)    | 是否终结订单   |
| add_to_blacklist | number(1)     |         |  否   |   r ,sl(bool),cc    | 是否加入黑名单 |
| audit_status     | number(3)     |   10    |  否   | l,r,q,sl(status),cc | 审核状态       |
| audit_by         | number(10)    |         |  是   |         l,r         | 审核人         |
| audit_time       | date          |         |  是   |         l,r         | 审核时间       |
| audit_msg        | varchar2(256) |         |  是   |         l,r         | 审核信息       |



# 四、基础信息类
###  1. 产品线[ots_product_line]

| 字段名      | 类型         | 默认值  | 为空  |       约束        | 描述       |
| ----------- | ------------ | :-----: | :---: | :---------------: | :--------- |
| pl_id       | number(10)   |   100   |  否   |   PK,SEQ,l,r,DI   | 产品线编号 |
| pl_name     | varchar2(64) |         |  否   |   q,l,r,DN,u,c    | 产品线名称 |
| status      | number(1)    |    0    |  否   | l,q,r,sl,cc,u,c,r | 状态       |
| create_time | date         | sysdate |  否   |     l,r(dtp)      | 创建时间   |

###  2. 业务流程[ots_product_flow]
| 字段名        | 类型         | 默认值 | 为空  |                       约束                       | 描述     |
| ------------- | ------------ | :----: | :---: | :----------------------------------------------: | :------- |
| flow_id       | number(10)   |  200   |  否   |                  PK,SEQ,l,r,DI                   | 流程编号 |
| flow_tag      | varchar2(32) |        |  否   |     q,l,r,UNQ(unq_flow_tag),u,c,sl(flow_tag)     | 流程名称 |
| pl_id         | number(10)   |        |  否   | l,r,UNQ(unq_flow_tag),q,u,c,sl(ots_product_line) | 产品线   |
| queue_name    | varchar2(64) |  '-'   |  是   |                     r,u,c,l                      | 队列名称 |
| scan_interval | number(10)   |   0     |  是   |                     r,u,c,l                      | 执行间隔 |
| delay         | number(10)   |   0    |  是   |                     r,u,c,l                      | 延后时长 |
| timeout       | number(10)   |    0    |  是   |                     r,u,c,l                      | 超时时长 |
| max_count     | number(10)   |   0     |  是   |                     r,u,c,l                      | 最大次数 |
| status        | number(1)    |   0    |  是   |                l,q,r,sl,cc,u,c,r                 | 状态     |

### 任务表[^tsk_system_task]

| 字段名            | 类型          | 默认值  | 为空  |                          约束                           | 描述                                        |
| ----------------- | ------------- | :-----: | :---: | :-----------------------------------------------------: | :------------------------------------------ |
| task_id           | number(20)    |         |  否   |                       PK,SEQ,l,r                        | 编号                                        |
| order_no          | varchar2(32)  |         |  是   |                          q,l,r                          | 订单号                                      |
| name              | varchar2(32)  |         |  否   |                           l,r                           | 流程名称                                    |
| create_time       | date          | sysdate |  否   | q(f:yyyy-MM-dd),l(f:MM/dd HH:mm:ss),r(f:MM/dd HH:mm:ss) | 创建时间                                    |
| last_execute_time | date          |         |  是   |                  l(f:MM/dd HH:mm:ss),r                  | 上次执行时间                                |
| next_execute_time | date          |         |  否   |         l(f:MM/dd HH:mm:ss),r(f:MM/dd HH:mm:ss)         | 下次执行时间                                |
| max_execute_time  | date          |         |  否   |                         r(dtp)                          | 执行期限(此时间前的任务可以被执行)          |
| next_interval     | number(10)    |         |  否   |                            r                            | 时间间隔,秒数                               |
| delete_interval   | number(10)    |         |  是   |                            r                            | 删除间隔,秒数                               |
| delete_time       | date          |         |  是   |                         r(dtp)                          | 删除期限                                    |
| count             | number(10)    |    0    |  否   |                           l,r                           | 执行次数                                    |
| max_count         | number(10)    |    0    |  否   |                            r                            | 最大执行次数                                |
| status            | number(2)     |         |  否   |                l,r,sl(process_status),cc                | 状态(20 等待，30 正在,0 已处理,90 处理失败) |
| batch_id          | number(20)    |         |  是   |                            r                            | 执行批次号                                  |
| queue_name        | varchar2(64)  |         |  否   |                           l,r                           | 消息队列                                    |
| msg_content       | varchar2(256) |         |  是   |                            r                            | 消息内容                                    |


### 1. 账户信息[^beanpay_account_info]

| 字段名       | 类型         | 默认值  | 为空  |             约束              | 描述                |
| ------------ | ------------ | :-----: | :---: | :---------------------------: | :------------------ |
| account_id   | number(20)   |  86000  |  否   |         PK,SEQ,l,r,DI         | 帐户编号            |
| account_name | varchar2(32) |         |  否   |          l,r,u,c,DN           | 帐户名称            |
| ident        | varchar2(32) |         |  否   |         r,sl(ident),c         | 系统标识            |
| groups       | varchar2(32) |         |  否   |    l,r,sl(account_group),c    | 用户分组            |
| eid          | varchar2(32) |         |  否   | l,r,q,sl(ots_merchant_info),c | 商户信息            |
| balance      | number(20,5) |    0    |  否   |               r               | 帐户余额，单位：元  |
| credit       | number(20,5) |    0    |  否   |            r ,u,c             | 信用余额，单位：元  |
| status       | number(1)    |    0    |  否   |    l,r,cc,sl(status) ,u,c     | 状态 0：正常 1:锁定 |
| create_time  | date         | sysdate |  否   |              l,r              | 创建时间            |

### 2. 账户余额变动信息[^beanpay_account_record]

| 字段名      | 类型           | 默认值  | 为空  |              约束              | 描述                                                           |
| ----------- | -------------- | :-----: | :---: | :----------------------------: | :------------------------------------------------------------- |
| record_id   | number(20)     | 100000  |  否   |           PK,SEQ,l,r           | 变动编号                                                       |
| account_id  | number(20)     |         |  否   | l,r,q,sl(beanpay_account_info) | 帐户编号                                                       |
| trade_no    | varchar2(32)   |         |  否   |             q,l,r,sl()              | 交易编号                                                       |
| ext_no      | varchar2(32)   |    0    |  是   |               r                | 拓展编号                                                       |
| trade_type  | number(1)      |    1    |  否   |      l,r,q,sl(trade_type)      | 交易类型 1:交易 2：手续费 3:佣金 4:红冲 5:平账                 |
| change_type | number(1)      |         |  否   |     l,r,q,sl(change_type)      | 变动类型 1:加款 2:提款 3：扣款 4：退款 5: 交易平账 6: 余额平账 |
| amount      | number(20,5)   |         |  否   |              l,r               | 变动金额 单位：元                                              |
| balance     | number(20,5)   |         |  否   |              l,r               | 帐户余额 单位：元                                              |
| create_time | date           | sysdate |  否   |              l,r               | 创建时间                                                       |
| memo        | varchar2(1024) |         |  是   |               r                | 交易说明                                                       |
| ext         | varchar2(1024) |         |  是   |               r                | 扩展字段                                                       |


### 字典配置[^dds_dictionary_info]

| 字段名  | 类型         | 默认值 | 为空  |       约束       | 描述   |
| ------- | ------------ | :----: | :---: | :--------------: | :----- |
| id      | number(10)   |        |  否   | PK,IS,SEQ,l,r,DI | 编号   |
| name    | varchar2(64) |        |  否   |   q,c,u,l,r,DN   | 名称   |
| value   | varchar2(32) |        |  否   |     c,u,l,r      | 值     |
| type    | varchar2(32) |        |  否   |   q,c,u,l,r,DT   | 类型   |
| status  | number(1)    |   0    |  否   | q,c,u,l,r,sl,cc  | 状态   |
| sort_no | number(2)    |   0    |  否   |     c,u,l,r      | 排序值 |

### 生命周期记录表[^lcs_life_time]

| 字段名       | 类型          |      默认值       | 为空  |    约束     | 描述           |
| ------------ | ------------- | :---------------: | :---: | :---------: | :------------- |
| id           | bigint        |                   |  否   | PK, IS, SEQ | id             |
| order_no     | varchar(32)   |                   |  否   |   IS, IDX   | 子系统唯一标识 |
| batch_no     | varchar(32)   |                   |  是   |     IS      | 自定义字段     |
| extral_param | varchar(32)   |                   |  是   |     IS      | 子系统唯一标识 |
| ip           | varchar(32)   |                   |  是   |     IS      | 用户ip         |
| content      | varchar(1000) |                   |  否   |     IS      | 内容           |
| create_time  | datetime      | current_timestamp |  否   |     IS      | 创建时间       |

* 生成DB gitcli db create ../docs/db.md  ./modules/const/db/scheme --gofile --drop --cover --seqfile
* 生成代码 gitcli md code entity ./docs/db.md -t 
* 生成SQL: gitcli md sql select ./docs/db.md -t ots_trade_order