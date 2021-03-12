<template>
  <!-- Add Form -->
  <el-dialog title="添加商户货架" width="65%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="货架名称:" prop="mer_shelf_name">
				<el-input size="medium" maxlength="64"
				 clearable v-model="addData.mer_shelf_name" placeholder="请输入货架名称">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="商户名称:" prop="mer_no">
				<el-select size="medium" style="width: 100%;"	v-model="addData.mer_no"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in merNo" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="商户佣金:" prop="mer_fee_discount">
				<el-input size="medium" maxlength="10" oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				 clearable v-model="addData.mer_fee_discount" placeholder="请输入商户佣金">
				</el-input>
      </el-form-item>
      
      <el-form-item label="交易服务费:" prop="trade_fee_discount">
				<el-input size="medium" maxlength="10" oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				 clearable v-model="addData.trade_fee_discount" placeholder="请输入交易服务费">
				</el-input>
      </el-form-item>
      
      <el-form-item label="支付手续费:" prop="payment_fee_discount">
				<el-input size="medium" maxlength="10" oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				 clearable v-model="addData.payment_fee_discount" placeholder="请输入支付手续费">
				</el-input>
      </el-form-item>
      
      <el-form-item label="订单超时时长:" prop="order_timeout">
				<el-input size="medium" maxlength="10"
				 clearable v-model="addData.order_timeout" placeholder="请输入订单超时时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="支付超时时长:" prop="payment_timeout">
				<el-input size="medium" maxlength="10"
				 clearable v-model="addData.payment_timeout" placeholder="请输入支付超时时长">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="开票方式:" prop="invoice_type">
				<el-select size="medium" style="width: 100%;"	v-model="addData.invoice_type"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in invoiceType" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="允许退款:" prop="can_refund">
				<el-select size="medium" style="width: 100%;"	v-model="addData.can_refund"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in canRefund" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="单次购买数量:" prop="limit_count">
				<el-input size="medium" maxlength="10"
				 clearable v-model="addData.limit_count" placeholder="请输入单次购买数量">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="允许拆单:" prop="can_split_order">
				<el-select size="medium" style="width: 100%;"	v-model="addData.can_split_order"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in canSplitOrder" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select size="medium" style="width: 100%;"	v-model="addData.status"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button size="medium" @click="resetForm('addForm')">取 消</el-button>
      <el-button size="medium" type="success" @click="add('addForm')">确 定</el-button>
    </div>
  </el-dialog>
  <!--Add Form -->
</template>

<script>
export default {
	data() {
		return {
			addData: {},
			dialogAddVisible: false,
			merNo:this.$enum.get("merchant_info"),
			invoiceType:this.$enum.get("invoice_type"),
			canRefund:this.$enum.get("bool"),
			canSplitOrder:this.$enum.get("bool"),
			status:this.$enum.get("status"),
			rules: {                    //数据验证规则
				mer_shelf_name: [{ required: true, message: "请输入货架名称", trigger: "blur" }],
				mer_no: [{ required: true, message: "请输入商户名称", trigger: "blur" }],
				mer_fee_discount: [{ required: true, message: "请输入商户佣金", trigger: "blur" }],
				trade_fee_discount: [{ required: true, message: "请输入交易服务费", trigger: "blur" }],
				payment_fee_discount: [{ required: true, message: "请输入支付手续费", trigger: "blur" }],
				order_timeout: [{ required: true, message: "请输入订单超时时长", trigger: "blur" }],
				payment_timeout: [{ required: true, message: "请输入支付超时时长", trigger: "blur" }],
				invoice_type: [{ required: true, message: "请输入开票方式", trigger: "blur" }],
				can_refund: [{ required: true, message: "请输入允许退款", trigger: "blur" }],
				limit_count: [{ required: true, message: "请输入单次购买数量", trigger: "blur" }],
				can_split_order: [{ required: true, message: "请输入允许拆单", trigger: "blur" }],
				status: [{ required: true, message: "请输入状态", trigger: "blur" }],
			},
		}
	},
	props: {
		refresh: {
			type: Function,
				default: () => {
				},
		}
	},
	created(){
	},
	methods: {
		closed() {
			this.refresh()
		},
		resetForm(formName) {
			this.dialogAddVisible = false;
			this.$refs[formName].resetFields();
		},
		show(){
			this.dialogAddVisible = true;
		},
		add(formName) {
			this.$refs[formName].validate((valid) => {
				if (valid) {
					this.$http.post("/merchant/shelf", this.addData, {}, true, true)
						.then(res => {
							this.$refs[formName].resetFields()
							this.dialogAddVisible = false
							this.refresh()
						})
						.catch(err => {
							this.$message({
								type: "error",
								message: err.response.data
							});
						})
				} else {
						console.log("error submit!!");
						return false;
				}
			});
		},
	}

}
</script>

<style scoped>
</style>
