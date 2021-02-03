<template>
  <!-- Add Form -->
  <el-dialog title="添加供货商错误码" width="65%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="编号" prop="id">
				<el-input maxlength="20" clearable v-model="addData.id" placeholder="请输入编号">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="供货商:" prop="spp_no">
				<el-select  placeholder="---请选择---" clearable v-model="addData.spp_no" style="width: 100%;">
					<el-option v-for="(item, index) in sppNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="产品线:" prop="pl_id">
				<el-select  placeholder="---请选择---" clearable v-model="addData.pl_id" style="width: 100%;">
					<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="分类:" prop="category">
				<el-select  placeholder="---请选择---" clearable v-model="addData.category" style="width: 100%;">
					<el-option v-for="(item, index) in category" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="处理码:" prop="deal_code">
				<el-select  placeholder="---请选择---" clearable v-model="addData.deal_code" style="width: 100%;">
					<el-option v-for="(item, index) in dealCode" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="错误码" prop="error_code">
				<el-input maxlength="32" clearable v-model="addData.error_code" placeholder="请输入错误码">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select  placeholder="---请选择---" clearable v-model="addData.status" style="width: 100%;">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="错误码描述" prop="error_desc">
				<el-input maxlength="64" clearable v-model="addData.error_desc" placeholder="请输入错误码描述">
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
      plID: this.$enum.get("product_line"),
      category: this.$enum.get("result_source"),
      dealCode: this.$enum.get("deal_code"),
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				id: [{ required: true, message: "请输入编号", trigger: "blur" }],
				spp_no: [{ required: true, message: "请输入供货商", trigger: "blur" }],
				pl_id: [{ required: true, message: "请输入产品线", trigger: "blur" }],
				category: [{ required: true, message: "请输入分类", trigger: "blur" }],
				deal_code: [{ required: true, message: "请输入处理码", trigger: "blur" }],
				error_code: [{ required: true, message: "请输入错误码", trigger: "blur" }],
				status: [{ required: true, message: "请输入状态", trigger: "blur" }],
				error_desc: [{ required: true, message: "请输入错误码描述", trigger: "blur" }],
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
					this.$http.post("/supplier/ecode", this.addData, {}, true, true)
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
