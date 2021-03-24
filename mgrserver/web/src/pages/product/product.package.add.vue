<template>
  <!-- Add Form -->
  <el-dialog title="添加产品包" width="65%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="包名称:" prop="pg_name">
				<el-input size="medium" maxlength="64"
				 clearable v-model="addData.pg_name" placeholder="请输入包名称">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="产品线:" prop="pl_id">
				<el-select size="medium" style="width: 100%;"	v-model="addData.pl_id"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="父级商品:" prop="parent_id">
				<el-select size="medium" style="width: 100%;"	v-model="addData.parent_id"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in parentID" :key="index" :value="item.value" :label="item.name"></el-option>
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
      
      <el-form-item label="商品面值:" prop="face">
				<el-input size="medium" maxlength="10"
				 clearable v-model="addData.face" placeholder="请输入商品面值">
				</el-input>
      </el-form-item>
      
      <el-form-item label="数量:" prop="num">
				<el-input size="medium" maxlength="2"
				 clearable v-model="addData.num" placeholder="请输入数量">
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
			plID:this.$enum.get("product_line"),
			parentID:this.$enum.get("product_package"),
			brandNo:this.$enum.get("brand"),
			provinceNo:this.$enum.get("province"),
			cityNo:this.$enum.get("city"),
			status:this.$enum.get("status"),
			rules: {                    //数据验证规则
				pg_name: [{ required: true, message: "请输入包名称", trigger: "blur" }],
				pl_id: [{ required: true, message: "请输入产品线", trigger: "blur" }],
				parent_id: [{ required: true, message: "请输入父级商品", trigger: "blur" }],
				brand_no: [{ required: true, message: "请输入品牌", trigger: "blur" }],
				province_no: [{ required: true, message: "请输入省份", trigger: "blur" }],
				city_no: [{ required: true, message: "请输入城市", trigger: "blur" }],
				face: [{ required: true, message: "请输入商品面值", trigger: "blur" }],
				num: [{ required: true, message: "请输入数量", trigger: "blur" }],
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
					this.$http.post("/product/package", this.addData, {}, true, true)
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
