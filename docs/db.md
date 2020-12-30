
## 一、商户信息

###  1. 商户信息[ots_merchant_info]

| 字段名      | 类型         | 默认值  | 为空  |     约束     | 描述     |
| ----------- | ------------ | :-----: | :---: | :----------: | :------- |
| mer_no      | varchar2(32) |         |  否   | PK,SEQ,RL,DI | 编号     |
| mer_name    | varchar2(64) |         |  否   |   CRUQL,DN   | 名称     |
| status      | number(1)    |    0    |  否   |   RUQL,SL    | 状态     |
| create_time | date         | sysdate |  否   |    RL,DT     | 创建时间 |


###  2. 商户货架[ots_merchant_shelf]

| 字段名         | 类型         | 默认值  | 为空  |            约束             | 描述                |
| -------------- | ------------ | :-----: | :---: | :-------------------------: | :------------------ |
| mer_shelf_id   | number(10)   |    1    |  否   |        PK,SEQ,RL,DI         | 货架编号            |
| mer_shelf_name | varchar2(64) |         |  否   |          CRUQL,DN           | 货架名称            |
| mer_no         | varchar2(32) |         |  否   | CRUQL,SL(ots_down_supplier) | 商户编号            |
| can_refund     | number(1)    |         |  否   |          CRUQL,SL           | 支持退款(0.是,1否)  |
| limit_count    | number(10)   |    1    |  否   |            CRUL             | 单次最大购买数量    |
| can_split      | number(1)    |         |  否   |          CRULQ,SL           | 是否拆单(0.是,1.否) |
| split_face     | number(20,5) |         |  否   |            CRUL             | 拆单面值            |
| order_timeout  | number(10)   |         |  否   |            RUCL             | 订单超时时长        |
| refund_timeout | number(10)   |         |  否   |            RCUL             | 退款超时时长        |
| status         | number(1)    |    0    |  否   |           RUQL,SL           | 状态                |
| create_time    | date         | sysdate |  否   |             RL              | 创建时间            |

###  3.商户商品[ots_merchant_product]

| 字段名               | 类型         | 默认值  | 为空  |            约束            | 描述                   |
| -------------------- | ------------ | :-----: | :---: | :------------------------: | :--------------------- |
| mer_product_id       | number(10)   |   300   |  否   |         PK,SEQ,LR          | 商品编号               |
| mer_shelf_id         | number(10)   |         |  否   |  CRUQL,SL(ots_down_shelf)  | 货架编号               |
| pl_id                | number(10)   |         |  否   | CRULQ,SL(ots_product_line) | 产品线                 |
| mer_product_no       | varchar2(32) |         |  是   |           CRUQL            | 商户商品编号           |
| carrier_no           | varchar2(8)  |         |  否   |          CRUQL,SL          | 运营商                 |
| province_no          | varchar2(8)  |    -    |  否   | CRUQL,SL(ots_canton_info)  | 省份                   |
| city_no              | varchar2(8)  |    -    |  否   | CRUQL,SL(ots_canton_info)  | 城市                   |
| invoice_type         | number(2)    |         |  否   |          CRUQL,SL          | 开票方式（1.不开发票） |
| face                 | number(20,5) |         |  否   |            CRUL            | 面值                   |
| user_discount        | number(10,5) |         |  否   |            CRUL            | 用户折扣（以面值算）   |
| mer_fee_discount     | number(10,5) |         |  否   |            CRUL            | 商户佣金               |
| trade_fee_discount   | number(10,5) |         |  否   |            CRUL            | 交易服务费             |
| payment_fee_discount | number(10,5) |         |  否   |            CRUL            | 支付手续费             |
| status               | number(1)    |    0    |  否   |          RUQL,SL           | 状态(0.是,1.否)        |
| create_time          | date         | sysdate |  否   |             RL             | 创建时间               |


## 二、供货商信息

###  1. 供货商信息[ots_supplier_info]

| 字段名      | 类型         | 默认值  | 为空  |     约束     | 描述     |
| :---------- | :----------- | :-----: | :---: | :----------: | :------- |
| spp_no      | varchar2(32) |         |  否   | PK,SEQ,RL,DI | 编号     |
| spp_name    | varchar2(64) |         |  否   |   CRUQL,DN   | 名称     |
| status      | number(1)    |    0    |  否   |   RUQL,SL    | 状态     |
| create_time | date         | sysdate |  否   |    RL,DT     | 创建时间 |


###  2. 供货商货架[ots_supplier_shelf]

| 字段名                 | 类型          | 默认值  | 为空  |            约束            | 描述                 |
| ---------------------- | ------------- | :-----: | :---: | :------------------------: | :------------------- |
| spp_shelf_id           | number(10)    |    1    |  否   |        PK,SEQ,RL,DI        | 货架编号             |
| spp_shelf_name         | varchar2(64)  |         |  否   |          CRUQL,DN          | 货架名称             |
| spp_no                 | varchar2(32)  |         |  否   | CRQUL,SL(ots_spp_supplier) | 供货商编号           |
| req_url                | varchar2(128) |         |  否   |            CRUL            | 请求地址             |
| query_url              | varchar2(128) |         |  是   |            RCUL            | 查询地址             |
| notify_url             | varchar2(128) |         |  否   |            CRUL            | 回调地址             |
| can_refund             | number(1)     |         |  否   |          CRUQL,SL          | 支持退货 (0.是,1.否) |
| request_replenish_time | number(10)    |         |  否   |            CRUL            | 发货后补间隔时间     |
| can_query              | number(1)     |    0    |  否   |            RUL             | 是否支持查询         |
| first_query_time       | number(10)    |         |  是   |            RCUL            | 首次查询时间         |
| query_replenish_time   | number(10)    |         |  是   |            RCUL            | 查询后补间隔时间     |
| status                 | number(1)     |    0    |  否   |          RUQL,SL           | 货架状态             |
| delivery_timeout       | number(10)    |   300   |  否   |            RUCL            | 发货超时时间         |
| return_timeout         | number(10)    |   300   |  否   |            CULR            | 退货超时时间         |
| limit_count            | number(10)    |    1    |  否   |            CRUL            | 单次最大发货数量     |
| create_time            | date          | sysdate |  否   |             RL             | 创建时间             |

###  3. 供货商商品[ots_supplier_product]

| 字段名               | 类型         | 默认值  | 为空  |            约束            | 描述                                   |
| -------------------- | ------------ | :-----: | :---: | :------------------------: | :------------------------------------- |
| spp_product_id       | number(10)   |   300   |  否   |         PK,SEQ,LR          | 商品编号                               |
| spp_shelf_id         | number(10)   |         |  否   |  CRUQL,SL(ots_spp_shelf)   | 货架编号                               |
| pl_id                | number(10)   |         |  否   | CRUQL,SL(ots_product_line) | 产品线                                 |
| spp_no               | varchar2(32) |         |  否   | CRQUL,SL(ots_spp_supplier) | 供货商编号                             |
| spp_product_no       | varchar2(32) |         |  是   |            CRUL            | 供货商商品编号                         |
| carrier_no           | varchar2(8)  |         |  否   |          CRUQL,SL          | 运营商                                 |
| province_no          | varchar2(8)  |    -    |  否   | CRUQL,SL(ots_canton_info)  | 省份                                   |
| city_no              | varchar2(8)  |    -    |  否   | CRUQL,SL(ots_canton_info)  | 城市                                   |
| invoice_type         | number(2)    |         |  否   |          CRUQL,SL          | 开票方式（1.不开发票，2.供货商开发票） |
| face                 | number(20,5) |         |  否   |            CRUL            | 面值                                   |
| cost_discount        | number(10,5) |         |  否   |            CRUL            | 成本折扣（以面值算）                   |
| spp_fee_discount     | number(10,5) |         |  否   |            CRUL            | 商户佣金                               |
| trade_fee_discount   | number(10,5) |         |  否   |            CRUL            | 交易服务费                             |
| payment_fee_discount | number(10,5) |         |  否   |            CRUL            | 支付手续费                             |
| status               | number(1)    |    0    |  否   |          RUQL,SL           | 状态                                   |
| create_time          | date         | sysdate |  否   |             RL             | 创建时间                               |



###  4. 供货商错误码[ots_supplier_error_code]

| 字段名          | 类型         | 默认值  | 为空  |            约束            | 描述       |
| --------------- | ------------ | :-----: | :---: | :------------------------: | :--------- |
| id              | number(20)   |         |  否   |         PK,SEQ,LR          | 编号       |
| spp_no          | varchar2(32) |         |  否   | CRUQL,SL(ots_spp_supplier) | 编号       |
| pl_id           | number(10)   |         |  否   | CRUQL,SL(ots_product_line) | 产品线     |
| error_code      | varchar2(32) |         |  否   |            CRUL            | 错误码     |
| deal_code       | number(2)    |         |  否   |            CRUL            | 处理码     |
| error_code_desc | varchar2(64) |         |  否   |            CRUL            | 错误码描述 |
| create_time     | date         | sysdate |  否   |             RL             | 创建时间   |


# 三、交易类

###  1. 订单记录[ots_trade_order]

| 字段名                | 类型         |   默认值   | 为空  |           约束            | 描述                             |
| --------------------- | ------------ | :--------: | :---: | :-----------------------: | :------------------------------- |
| order_id              | number(20)   | 1100000000 |  否   |           PK,RL           | 订单编号                         |
| mer_no                | varchar2(32) |            |  否   | QRL,SL(ots_down_supplier) | 商户编号                         |
| mer_order_no          | varchar2(64) |            |  否   |            QRL            | 商户订单编号                     |
| mer_shelf_id          | number(10)   |            |  否   |  QRL,SL(ots_down_shelf)   | 商户货架编号                     |
| mer_product_id        | number(10)   |            |  否   |            QRL            | 商户商品编号                     |
| mer_product_no        | varchar2(32) |            |  是   |            QRL            | 外部商品编号                     |
| pl_id                 | number(10)   |            |  否   | QRL,SL(ots_product_line)  | 产品线                           |
| carrier_no            | varchar2(8)  |            |  否   |          QRL,SL           | 运营商                           |
| province_no           | varchar2(8)  |            |  否   |  QRL,SL(ots_canton_info)  | 省份                             |
| city_no               | varchar2(8)  |            |  否   |  QRL,SL(ots_canton_info)  | 城市                             |
| face                  | number(20,5) |            |  否   |            RL             | 商品面值                         |
| num                   | number(10)   |            |  否   |            RL             | 商品数量                         |
| total_face            | number(20,5) |            |  否   |            RL             | 商品总面值                       |
| invoice_type          | number(2)    |            |  否   |          QRL,SL           | 开票方式（1.不开发票）           |
| sell_amount           | number(20,5) |            |  否   |            RL             | 总销售金额                       |
| mer_fee_amount        | number(20,5) |            |  否   |            RL             | 商务佣金金额                     |
| trade_fee_amount      | number(20,5) |            |  否   |            RL             | 商务服务费金额                   |
| payment_fee_amount    | number(20,5) |            |  否   |            RL             | 支付手续费金额                   |
| can_split_order       | number(1)    |            |  否   |          QRL,SL           | 是否拆单（0.是，1否）            |
| split_order_face      | number(20,5) |            |  否   |            RL             | 拆单面值                         |
| create_time           | date         |  sysdate   |  否   |           RL,DT           | 创建时间                         |
| order_timeout         | date         |            |  否   |           RL,DT           | 订单超时时间                     |
| delivery_pause        | number(1)    |     1      |  否   |          QRL,SL           | 发货暂停（0.是，1否）            |
| order_status          | number(3)    |     10     |  否   |          QRL,SL           | 订单状态                         |
| payment_status        | number(3)    |     20     |  否   |          QRL,SL           | 支付状态                         |
| delivery_status       | number(3)    |     10     |  否   |          QRL,SL           | 发货状态                         |
| refund_status         | number(3)    |     10     |  否   |          QRL,SL           | 退款状态                         |
| notify_status         | number(3)    |     10     |  否   |          QRL,SL           | 通知状态                         |
| is_refund             | number(1)    |     1      |  否   |          QRL,SL           | 用户退款（0.是，1否）            |
| bind_face             | number(20,5) |     0      |  否   |            RL             | 成功绑定总面值                   |
| success_face          | number(20,5) |     0      |  否   |            RL             | 实际成功总面值                   |
| success_user_amount   | number(20,5) |     0      |  否   |            RL             | 实际成功总销售金额 （1）         |
| success_mer_fee       | number(20,5) |     0      |  否   |            RL             | 实际成功总佣金金额 （2）         |
| success_trade_fee     | number(20,5) |     0      |  否   |            RL             | 实际成功总服务费金额 （3）       |
| success_payment_fee   | number(20,5) |     0      |  否   |            RL             | 实际成功总手续费金额 （4）       |
| success_cost_amount   | number(20,5) |     0      |  否   |            RL             | 实际发货成功总成本 （5）         |
| success_spp_fee       | number(20,5) |     0      |  否   |            RL             | 实际发货成功总供货商佣金 （6）   |
| success_spp_trade_fee | number(20,5) |     0      |  否   |            RL             | 实际发货成功总供货商服务费 （7） |
| profit                | number(20,5) |     0      |  否   |            RL             | 利润（1-2+3-4-5+6+7）            |



###  2. 订单发货表[ots_trade_delivery]

| 字段名             | 类型           | 默认值  | 为空  |           约束            | 描述                                          |
| ------------------ | -------------- | :-----: | :---: | :-----------------------: | :-------------------------------------------- |
| delivery_id        | number(20)     |  20000  |  否   |          PK,QRL           | 发货编号                                      |
| spp_spp_no         | varchar2(32)   |         |  否   |  RL,SL(ots_spp_supplier)  | 供货商编号                                    |
| spp_product_id     | number(10)     |         |  否   |            RL             | 供货商商品编号                                |
| spp_delivery_no    | varchar2(32)   |         |  是   |            RL             | 供货商发货编号                                |
| spp_product_no     | varchar2(32)   |         |  是   |            RL             | 供货商商品编号                                |
| order_id           | number(20)     |         |  否   |            RL             | 订单编号                                      |
| mer_no             | varchar2(32)   |         |  否   | RQL,SL(ots_down_supplier) | 商户编号                                      |
| mer_product_id     | number(10)     |         |  否   |            RL             | 商户商品编号                                  |
| pl_id              | number(10)     |         |  否   | RQL,SL(ots_product_line)  | 产品线                                        |
| carrier_no         | varchar2(8)    |         |  否   |          QRL,SL           | 运营商                                        |
| province_no        | varchar2(8)    |         |  否   |  QRL,SL(ots_canton_info)  | 省份                                          |
| city_no            | varchar2(8)    |         |  否   |  QRL,SL(ots_canton_info)  | 城市                                          |
| invoice_type       | number(3)      |         |  否   |          RQL,SL           | 开票方式（1.不开发票，2.供货商开发票）        |
| delivery_status    | number(3)      |   20    |  否   |          RQL,SL           | 发货状态                                      |
| payment_status     | number(3)      |   10    |  否   |          RQL,SL           | 支付状态                                      |
| create_time        | date           | sysdate |  否   |           RL,DT           | 创建时间                                      |
| face               | number(20,5)   |         |  否   |            RL             | 商品面值                                      |
| num                | number(10)     |         |  否   |            RL             | 发货数量                                      |
| total_face         | number(20,5)   |         |  否   |            RL             | 发货总面值                                    |
| cost_amount        | number(20,5)   |         |  否   |            RL             | 发货成本                                      |
| spp_fee_amount     | number(20,5)   |         |  否   |            RL             | 供货商佣金                                    |
| trade_fee_amount   | number(20,5)   |         |  否   |            RL             | 发货服务费                                    |
| payment_fee_amount | number(20,5)   |         |  否   |            RL             | 支付服务费                                    |
| succ_face          | number(10)     |         |  是   |            RL             | 成功面值                                      |
| start_time         | date           |         |  是   |            RL             | 开始时间                                      |
| end_time           | date           |         |  是   |            RL             | 结束时间                                      |
| return_msg         | varchar2(256)  |         |  是   |            RL             | 发货返回信息                                  |
| request_params     | varchar2(2000) |         |  是   |            RL             | 发货信息参数json                              |
| result_source      | number(1)      |         |  是   |          RQL,SL           | 发货结果来源（1：通知，2：查询，3：同步返回） |
| result_code        | varchar2(32)   |         |  是   |            LR             | 发货结果码                                    |
| spp_product_cost   | number(20,5)   |         |  是   |            LR             | 供货商实际折扣                                |
| spp_order_no       | varchar2(32)   |         |  是   |            RL             | 供货商发货订单号                              |
| last_update_time   | date           | sysdate |  否   |            RL             | 最后更新时间                                  |


###  3. 生命周期记录表[ots_trade_lifetime]

| 字段名      | 类型           | 默认值  | 为空  | 约束  | 描述     |
| ----------- | -------------- | :-----: | :---: | :---: | :------- |
| id          | number(20)     |         |  否   | PK,RL | 编号     |
| create_time | date           | sysdate |  否   | RL,DT | 创建时间 |
| ip          | varchar2(20)   |         |  是   |  RQL  | 服务器ip |
| order_id    | number(20)     |         |  否   |  RL   | 订单编号 |
| delivery_id | varchar2(30)   |         |  是   |  RQL  | 发货编号 |
| content     | varchar2(1000) |         |  否   |  RQL  | 操作内容 |



###  4. 退款申请[ots_refund_apply]

| 字段名        | 类型         | 默认值  | 为空  |           约束            | 描述       |
| ------------- | ------------ | :-----: | :---: | :-----------------------: | :--------- |
| apply_id      | number(20)   |         |  否   |            QRL            | 申请编号   |
| order_id      | number(20)   |         |  否   |            QRL            | 订单编号   |
| mer_no        | varchar2(32) |         |  否   | QRL,SL(ots_down_supplier) | 商户编号   |
| mer_order_no  | varchar2(32) |         |  否   |            QRL            | 商户订单号 |
| refund_cause  | number(3)    |   10    |  否   |          QRL,SL           | 退款原因   |
| apply_status  | number(2)    |   10    |  否   |          QRL,SL           | 申请状态   |
| refund_amount | number(20,5) |         |  否   |            RL             | 退款金额   |
| create_time   | date         | sysdate |  否   |           RL,DT           | 创建时间   |



###  5.订单通知表[ots_notify_info]

| 字段名         | 类型          | 默认值  | 为空  |           约束            | 描述                                             |
| -------------- | ------------- | :-----: | :---: | :-----------------------: | :----------------------------------------------- |
| order_id       | number(20)    |         |  否   |            LQR            | 订单编号                                         |
| mer_no         | varchar2(32)  |         |  否   | QRL,SL(ots_down_supplier) | 商户编号                                         |
| mer_order_no   | varchar2(64)  |         |  否   |            QRL            | 商户订单编号                                     |
| mer_product_id | number(10)    |         |  否   |            QRL            | 商户商品编号                                     |
| notify_url     | varchar2(128) |         |  否   |            LR             | 通知地址                                         |
| order_status   | number(3)     |   10    |  否   |          QRL,SL           | 订单状态                                         |
| notify_status  | number(3)     |   10    |  否   |          LQR,SL           | 通知状态（0成功,10未开始,20等待通知,30正在通知） |
| max_count      | number(3)     |         |  否   |            LR             | 最大通知次数                                     |
| notify_count   | number(3)     |    0    |  否   |            LR             | 通知次数                                         |
| create_time    | date          | sysdate |  否   |          LRQ,DT           | 创建时间                                         |
| start_time     | date          |         |  是   |            LR             | 开始时间                                         |
| end_time       | date          |         |  是   |            LR             | 结束时间                                         |
| notify_msg     | varchar2(256) |         |  是   |            LR             | 通知结果信息                                     |


###  6. 发货人工审核表[ots_audit_info]

| 字段名           | 类型          | 默认值  | 为空  |  约束   | 描述           |
| ---------------- | ------------- | :-----: | :---: | :-----: | :------------- |
| delivery_id      | number(20)    |         |  是   |   LQR   | 发货编号       |
| order_id         | number(20)    |         |  否   |   LQR   | 订单编号       |
| create_time      | date          | sysdate |  否   |   LQR   | 创建时间       |
| delivery_status  | number(2)     |         |  否   |   LQR   | 发货结果       |
| end_order        | number(1)     |         |  否   |   LQR   | 是否终结订单   |
| add_to_blacklist | number(1)     |         |  否   |   LQR   | 是否加入黑名单 |
| audit_status     | number(3)     |         |  否   | LRUQ,SL | 审核状态       |
| audit_by         | number(10)    |         |  是   |   LRU   | 审核人         |
| audit_time       | date          |         |  是   |   LRU   | 审核时间       |
| audit_msg        | varchar2(256) |         |  是   |   LRU   | 审核信息       |



# 四、基础信息类
###  1. 产品线[ots_product_line]

| 字段名                 | 类型         |            默认值            | 为空  |     约束     | 描述             |
| ---------------------- | ------------ | :--------------------------: | :---: | :----------: | :--------------- |
| pl_id                | number(10)   |              1               |  否   | PK,SEQ,LR,DI | 产品线编号       |
| line_name              | varchar2(64) |                              |  否   |   LCRUQ,DN   | 产品线名称       |
| can_package_delivery   | number(1)    |              1               |  否   |   LCRUQ,SL   | 支持打包发货     |
| payment_queue          | varchar2(32) |        oms:order:pay         |  是   |     LCRU     | 支付队列         |
| bind_queue             | varchar2(32) |        oms:order:bind        |  是   |     LCRU     | 绑定队列         |
| refund_queue           | varchar2(32) |        oms:refund:pay        |  是   |     LCRU     | 退款队列         |
| notify_queue           | varchar2(32) |       oms:order:notify       |  是   |     LCRU     | 通知队列         |
| spp_payment_queue      | varchar2(32) |      oms:order:spp_pay       |  是   |     LCRU     | 供货商支付队列   |
| spp_refund_queue       | varchar2(32) |      oms:refund:spp_pay      |  是   |     LCRU     | 供货商退款队列   |
| refund_notify_queue    | varchar2(32) |      oms:refund:notify       |  是   |     LCRU     | 退款通知队列     |
| order_refund_queue     | varchar2(32) |      oms:timeout:refund      |  是   |     LCRU     | 订单失败退款队列 |
| order_timeout_queue    | varchar2(32) |    oms:timeout:order_deal    |  是   |     LCRU     | 订单超时处理队列 |
| refund_timeout_queue   | varchar2(32) |   oms:timeout:refund_deal    |  是   |     LCRU     | 退款超时处理队列 |
| delivery_unknown_queue | varchar2(32) | oms:timeout:delivery_unknown |  是   |     LCRU     | 发货未知处理队列 |
| return_unknown_queue   | varchar2(32) |  oms:timeout:return_unknown  |  是   |     LCRU     | 退货未知处理队列 |
| delivery_start_queue   | varchar2(32) |      oms:order:delivery      |  是   |     LCRU     | 发货开始队列     |
| delivery_finish_queue  | varchar2(32) |  oms:order:delivery_finish   |  是   |     LCRU     | 发货结束队列     |
| return_queue           | varchar2(32) |      oms:refund:return       |  是   |     LCRU     | 退货队列         |
| return_finish_queue    | varchar2(32) |  oms:refund:return_complete  |  是   |     LCRU     | 退货结束队列     |


###  2.字典表[dds_dictionary_info]

| 字段名 | 类型         | 默认值 | 为空  |                 约束                  | 描述   |
| ------ | ------------ | :----: | :---: | :-----------------------------------: | :----- |
| id     | number(10)   |        |  否   |               PK,SEQ,LR               | 编号   |
| name   | varchar2(64) |        |  否   |                 LCRUQ                 | 名称   |
| value  | varchar2(32) |        |  否   |                 LCRU                  | 值     |
| type   | varchar2(32) |        |  否   | LCRUQ,IDX(IDX_DICTIONARY_INFO_TYPE,1) | 类型   |
| sort   | number(3)    |   0    |  否   |                 LCRU                  | 排序值 |
| status | number(1)    |        |  否   |                LRUQ,SL                | 状态   |



###  3. 省市信息[ots_canton_info]

| 字段名        | 类型         | 默认值 | 为空  |   约束   | 描述         |
| ------------- | ------------ | :----: | :---: | :------: | :----------- |
| canton_code   | varchar2(32) |        |  否   | PK,CRULQ | 区域编号     |
| chinese_name  | varchar2(32) |        |  否   | CRULQ,DN | 中文名称     |
| spell         | varchar2(64) |        |  否   |   CRUL   | 英文或全拼   |
| grade         | number(1)    |        |  否   |   CRUL   | 行政级别     |
| parent        | varchar2(32) |        |  否   |  CRULQ   | 父级         |
| simple_spell  | varchar2(8)  |        |  否   | CRULQ,DI | 简拼         |
| area_code     | varchar2(8)  |        |  否   |   CRUL   | 区号         |
| standard_code | number(6)    |        |  否   |   CRUL   | 标准行政编码 |


* 文件生成方式 gitcli create gofile db.md  ../modules/const/sql/mysql 