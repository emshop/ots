<template>
  <!-- Add Form -->
  <el-dialog title="添加商户商品" width="65%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="110px">
      
			<el-form-item label="货架名称:" prop="mer_shelf_id">
				<el-select size="medium" style="width: 100%;"	v-model="addData.mer_shelf_id"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in merShelfID" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="商户名称:" prop="mer_no">
				<el-select size="medium" style="width: 100%;"	v-model="addData.mer_no"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in merNo" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="产品线:" prop="pl_id">
				<el-select size="medium" style="width: 100%;"	v-model="addData.pl_id"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="品牌:" prop="brand_no">
				<el-select size="medium" style="width: 100%;"	v-model="addData.brand_no"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in brandNo" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="省份:" prop="province_no">
				<el-select size="medium" style="width: 100%;"	v-model="addData.province_no"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in provinceNo" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="城市:" prop="city_no">
				<el-select size="medium" style="width: 100%;"	v-model="addData.city_no"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in cityNo" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="面值:" prop="face">
				<el-input size="medium" maxlength="10"
				 clearable v-model="addData.face" placeholder="请输入面值">
				</el-input>
      </el-form-item>
      
      <el-form-item label="商户商品编号:" prop="mer_product_no">
				<el-input size="medium" maxlength="32"
				 clearable v-model="addData.mer_product_no" placeholder="请输入商户商品编号">
				</el-input>
      </el-form-item>
      
      <el-form-item label="销售折扣:" prop="discount">
				<el-input size="medium" maxlength="10" oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				 clearable v-model="addData.discount" placeholder="请输入销售折扣">
				</el-input>
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
			merShelfID:this.$enum.get("merchant_shelf"),
			merNo:this.$enum.get("merchant_info"),
			plID:this.$enum.get("product_line"),
			brandNo:this.$enum.get("brand"),
			provinceNo:this.$enum.get("province"),
			cityNo:this.$enum.get("city"),
			status:this.$enum.get("status"),
			rules: {                    //数据验证规则
				mer_shelf_id: [{ required: true, message: "请输入货架名称", trigger: "blur" }],
				mer_no: [{ required: true, message: "请输入商户名称", trigger: "blur" }],
				pl_id: [{ required: true, message: "请输入产品线", trigger: "blur" }],
				brand_no: [{ required: true, message: "请输入品牌", trigger: "blur" }],
				province_no: [{ required: true, message: "请输入省份", trigger: "blur" }],
				city_no: [{ required: true, message: "请输入城市", trigger: "blur" }],
				face: [{ required: true, message: "请输入面值", trigger: "blur" }],
				discount: [{ required: true, message: "请输入销售折扣", trigger: "blur" }],
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
					this.$http.post("/merchant/product", this.addData, {}, true, true)
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
