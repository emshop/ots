<template>
  <!-- Add Form -->
  <el-dialog title="添加商户上游" width="25%" :visible.sync="dialogAddVisible">
    <el-form :model="addData"  :rules="rules" ref="addForm" label-width="110px">
      
			<el-form-item label="货架:" prop="mer_shelf_id">
				<el-select size="medium" style="width: 100%;"	v-model="addData.mer_shelf_id"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in merShelfID" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="供货商:" prop="spp_no">
				<el-select size="medium" style="width: 100%;"	v-model="addData.spp_no"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in sppNo" :key="index" :value="item.value" :label="item.name"></el-option>
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
			merShelfID:this.$enum.get("merchant_shelf"),
			sppNo:this.$enum.get("supplier_info"),
			status:this.$enum.get("status"),
			rules: {                    //数据验证规则
				mer_shelf_id: [{ required: true, message: "请输入货架", trigger: "blur" }],
				spp_no: [{ required: true, message: "请输入供货商", trigger: "blur" }],
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
					this.$http.post("/merchant/upstream", this.addData, {}, true, true)
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
