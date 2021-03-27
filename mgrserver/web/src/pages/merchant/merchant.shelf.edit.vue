<template>
	<el-dialog title="编辑商户货架" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData" :inline="true" :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="货架名称:" prop="mer_shelf_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.mer_shelf_name" placeholder="请输入货架名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="商户佣金:" prop="mer_fee_discount">
				<el-input size="medium" maxlength="10" oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				clearable v-model="editData.mer_fee_discount" placeholder="请输入商户佣金">
				</el-input>
      </el-form-item>
      
      <el-form-item label="交易服务费:" prop="trade_fee_discount">
				<el-input size="medium" maxlength="10" oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				clearable v-model="editData.trade_fee_discount" placeholder="请输入交易服务费">
				</el-input>
      </el-form-item>
      
      <el-form-item label="支付手续费:" prop="payment_fee_discount">
				<el-input size="medium" maxlength="10" oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				clearable v-model="editData.payment_fee_discount" placeholder="请输入支付手续费">
				</el-input>
      </el-form-item>
      
      <el-form-item label="订单超时时长:" prop="order_timeout">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.order_timeout" placeholder="请输入订单超时时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="支付超时时长:" prop="payment_timeout">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.payment_timeout" placeholder="请输入支付超时时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="单次最大数量:" prop="limit_count">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.limit_count" placeholder="请输入单次最大数量">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="允许开票:" prop="invoice_type">
				<el-select size="medium" style="width: 100%;"	v-model="editData.invoice_type" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in invoiceType" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="允许退款:" prop="can_refund">
				<el-select size="medium" style="width: 100%;"	v-model="editData.can_refund" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in canRefund" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="指定上游:" prop="assign_upstream">
				<el-select size="medium" style="width: 100%;"	v-model="editData.assign_upstream" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in assignUpstream" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select size="medium" style="width: 100%;"	v-model="editData.status" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
    </el-form>
		<div slot="footer" class="dialog-footer">
			<el-button size="medium" @click="dialogFormVisible = false">取 消</el-button>
			<el-button type="success" size="medium" @click="edit">确 定</el-button>
		</div>
	</el-dialog>
</template>

<script>
export default {
	data() {
		return {
			dialogFormVisible: false,    //编辑表单显示隐藏
			editData: {},                //编辑数据对象
      invoiceType: this.$enum.get("invoice_type"),
      canRefund: this.$enum.get("bool"),
      assignUpstream: this.$enum.get("bool"),
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				mer_shelf_name: [
					{ required: true, message: "请输入货架名称", trigger: "blur" }
				],
				mer_fee_discount: [
					{ required: true, message: "请输入商户佣金", trigger: "blur" }
				],
				trade_fee_discount: [
					{ required: true, message: "请输入交易服务费", trigger: "blur" }
				],
				payment_fee_discount: [
					{ required: true, message: "请输入支付手续费", trigger: "blur" }
				],
				order_timeout: [
					{ required: true, message: "请输入订单超时时长", trigger: "blur" }
				],
				payment_timeout: [
					{ required: true, message: "请输入支付超时时长", trigger: "blur" }
				],
				limit_count: [
					{ required: true, message: "请输入单次最大数量", trigger: "blur" }
				],
				invoice_type: [
					{ required: true, message: "请输入允许开票", trigger: "blur" }
				],
				can_refund: [
					{ required: true, message: "请输入允许退款", trigger: "blur" }
				],
				assign_upstream: [
					{ required: true, message: "请输入指定上游", trigger: "blur" }
				],
				status: [
					{ required: true, message: "请输入状态", trigger: "blur" }
				],
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
		show() {
			var mer_shelf_id = this.editData.mer_shelf_id
			this.editData = this.$http.xget("/merchant/shelf", { mer_shelf_id: mer_shelf_id })
			this.editData.mer_shelf_id = mer_shelf_id
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/merchant/shelf", this.editData, {}, true, true)
			.then(res => {			
				this.dialogFormVisible = false;
				this.refresh()
			})
		},
	}
}
</script>

<style scoped>
</style>
