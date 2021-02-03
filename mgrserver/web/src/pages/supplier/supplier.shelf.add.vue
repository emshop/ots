<template>
  <!-- Add Form -->
  <el-dialog title="添加供货商货架" width="65%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="货架编号" prop="spp_shelf_id">
				<el-input maxlength="10" clearable v-model="addData.spp_shelf_id" placeholder="请输入货架编号">
				</el-input>
      </el-form-item>
      
      <el-form-item label="货架名称" prop="spp_shelf_name">
				<el-input maxlength="64" clearable v-model="addData.spp_shelf_name" placeholder="请输入货架名称">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="供货商:" prop="spp_no">
				<el-select  placeholder="---请选择---" clearable v-model="addData.spp_no" style="width: 100%;">
					<el-option v-for="(item, index) in sppNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="请求地址" prop="req_url">
				<el-input maxlength="128" clearable v-model="addData.req_url" placeholder="请输入请求地址">
				</el-input>
      </el-form-item>
      
      <el-form-item label="查询地址" prop="query_url">
				<el-input maxlength="128" clearable v-model="addData.query_url" placeholder="请输入查询地址">
				</el-input>
      </el-form-item>
      
      <el-form-item label="回调地址" prop="notify_url">
				<el-input maxlength="128" clearable v-model="addData.notify_url" placeholder="请输入回调地址">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="开票:" prop="invoice_type">
				<el-select  placeholder="---请选择---" clearable v-model="addData.invoice_type" style="width: 100%;">
					<el-option v-for="(item, index) in invoiceType" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="商户佣金" prop="spp_fee_discount">
				<el-input maxlength="0" clearable v-model="addData.spp_fee_discount" placeholder="请输入商户佣金">
				</el-input>
      </el-form-item>
      
      <el-form-item label="交易服务费" prop="trade_fee_discount">
				<el-input maxlength="0" clearable v-model="addData.trade_fee_discount" placeholder="请输入交易服务费">
				</el-input>
      </el-form-item>
      
      <el-form-item label="支付手续费" prop="payment_fee_discount">
				<el-input maxlength="0" clearable v-model="addData.payment_fee_discount" placeholder="请输入支付手续费">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="支持退货:" prop="can_refund">
				<el-select  placeholder="---请选择---" clearable v-model="addData.can_refund" style="width: 100%;">
					<el-option v-for="(item, index) in canRefund" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="货架状态:" prop="status">
				<el-select  placeholder="---请选择---" clearable v-model="addData.status" style="width: 100%;">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="单次最大发货数量" prop="limit_count">
				<el-input maxlength="10" clearable v-model="addData.limit_count" placeholder="请输入单次最大发货数量">
				</el-input>
      </el-form-item>
      
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button size="small" @click="resetForm('addForm')">取 消</el-button>
      <el-button size="small" type="success" @click="add('addForm')">确 定</el-button>
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
      sppNo: this.$enum.get("supplier_info"),
      invoiceType: this.$enum.get("invoice_type"),
      canRefund: this.$enum.get("bool"),
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				spp_shelf_id: [{ required: true, message: "请输入货架编号", trigger: "blur" }],
				spp_shelf_name: [{ required: true, message: "请输入货架名称", trigger: "blur" }],
				spp_no: [{ required: true, message: "请输入供货商", trigger: "blur" }],
				req_url: [{ required: true, message: "请输入请求地址", trigger: "blur" }],
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
					this.$http.post("/supplier/shelf", this.addData, {}, true, true)
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
