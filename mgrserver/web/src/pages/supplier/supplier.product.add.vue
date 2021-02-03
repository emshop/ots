<template>
  <!-- Add Form -->
  <el-dialog title="添加供货商商品" width="65%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="商品编号" prop="spp_product_id">
				<el-input maxlength="10" clearable v-model="addData.spp_product_id" placeholder="请输入商品编号">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="货架名称:" prop="spp_shelf_id">
				<el-select  placeholder="---请选择---" clearable v-model="addData.spp_shelf_id" style="width: 100%;">
					<el-option v-for="(item, index) in sppShelfID" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="供货商:" prop="spp_no">
				<el-select  placeholder="---请选择---" clearable v-model="addData.spp_no" style="width: 100%;">
					<el-option v-for="(item, index) in sppNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="供货商商品编号" prop="spp_product_no">
				<el-input maxlength="32" clearable v-model="addData.spp_product_no" placeholder="请输入供货商商品编号">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="产品线:" prop="pl_id">
				<el-select  placeholder="---请选择---" clearable v-model="addData.pl_id" style="width: 100%;">
					<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="品牌:" prop="brand_no">
				<el-select  placeholder="---请选择---" clearable v-model="addData.brand_no" style="width: 100%;">
					<el-option v-for="(item, index) in brandNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="省份:" prop="province_no">
				<el-select  placeholder="---请选择---" clearable v-model="addData.province_no" style="width: 100%;">
					<el-option v-for="(item, index) in provinceNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="城市:" prop="city_no">
				<el-select  placeholder="---请选择---" clearable v-model="addData.city_no" style="width: 100%;">
					<el-option v-for="(item, index) in cityNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="面值" prop="face">
				<el-input maxlength="10" clearable v-model="addData.face" placeholder="请输入面值">
				</el-input>
      </el-form-item>
      
      <el-form-item label="成本折扣" prop="cost_discount">
				<el-input maxlength="0" clearable v-model="addData.cost_discount" placeholder="请输入成本折扣">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select  placeholder="---请选择---" clearable v-model="addData.status" style="width: 100%;">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
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
      sppShelfID: this.$enum.get("supplier_shelf"),
      sppNo: this.$enum.get("supplier_info"),
      plID: this.$enum.get("product_line"),
      brandNo: this.$enum.get("brand"),
      provinceNo: this.$enum.get("province"),
      cityNo: this.$enum.get("city"),
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				spp_product_id: [{ required: true, message: "请输入商品编号", trigger: "blur" }],
				spp_shelf_id: [{ required: true, message: "请输入货架名称", trigger: "blur" }],
				spp_no: [{ required: true, message: "请输入供货商", trigger: "blur" }],
				pl_id: [{ required: true, message: "请输入产品线", trigger: "blur" }],
				brand_no: [{ required: true, message: "请输入品牌", trigger: "blur" }],
				face: [{ required: true, message: "请输入面值", trigger: "blur" }],
				cost_discount: [{ required: true, message: "请输入成本折扣", trigger: "blur" }],
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
					this.$http.post("/supplier/product", this.addData, {}, true, true)
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
