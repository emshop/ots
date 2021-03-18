#!/bin/sh

time=$(date "+%Y%m%d%H%M%S")

max=1
while(($max < 100000000))
do
max=$max+1
url="http://localhost:8081/order/request?mer_no=qx001&mer_order_no=m${time}${k}&mer_product_id=56&num=1&account_name=15828680877"
echo $url
done