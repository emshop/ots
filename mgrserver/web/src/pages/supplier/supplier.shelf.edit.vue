<template>
	<el-dialog title="编辑供货商货架" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData" :inline="true" :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="货架名称:" prop="spp_shelf_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.spp_shelf_name" placeholder="请输入货架名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="请求地址:" prop="req_url">
				<el-input size="medium" maxlength="128"
				clearable v-model="editData.req_url" placeholder="请输入请求地址">
				</el-input>
      </el-form-item>
      
      <el-form-item label="查询地址:" prop="query_url">
				<el-input size="medium" maxlength="128"
				clearable v-model="editData.query_url" placeholder="请输入查询地址">
				</el-input>
      </el-form-item>
      
      <el-form-item label="回调地址:" prop="notify_url">
				<el-input size="medium" maxlength="128"
				clearable v-model="editData.notify_url" placeholder="请输入回调地址">
				</el-input>
      </el-form-item>
      
      <el-form-item label="商户佣金:" prop="spp_fee_discount">
				<el-input size="medium" maxlength="10" oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				clearable v-model="editData.spp_fee_discount" placeholder="请输入商户佣金">
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
      
      <el-form-item label="单次最大发货数量:" prop="limit_count">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.limit_count" placeholder="请输入单次最大发货数量">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="支持开票:" prop="invoice_type">
				<el-select size="medium" style="width: 100%;"	v-model="editData.invoice_type" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in invoiceType" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="支持退货:" prop="can_refund">
				<el-select size="medium" style="width: 100%;"	v-model="editData.can_refund" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in canRefund" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="人工发货:" prop="is_mf">
				<el-select size="medium" style="width: 100%;"	v-model="editData.is_mf" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in isMf" :key="index" :value="item.value" :label="item.name"></el-option>
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
      isMf: this.$enum.get("bool"),
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				spp_shelf_name: [
					{ required: true, message: "请输入货架名称", trigger: "blur" }
				],
				req_url: [
					{ required: true, message: "请输入请求地址", trigger: "blur" }
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
			var spp_shelf_id = this.editData.spp_shelf_id
			this.editData = this.$http.xget("/supplier/shelf", { spp_shelf_id: spp_shelf_id })
			this.editData.spp_shelf_id = spp_shelf_id
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/supplier/shelf", this.editData, {}, true, true)
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
