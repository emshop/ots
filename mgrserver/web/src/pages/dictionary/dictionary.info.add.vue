<template>
  <!-- Add Form -->
  <el-dialog title="添加字典配置" width="25%" :visible.sync="dialogAddVisible">
    <el-form :model="addData"  :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="名称" prop="name">
				<el-input size="medium" maxlength="64"
				 clearable v-model="addData.name" placeholder="请输入名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="值" prop="value">
				<el-input size="medium" maxlength="32"
				 clearable v-model="addData.value" placeholder="请输入值">
				</el-input>
      </el-form-item>
      
      <el-form-item label="类型" prop="type">
				<el-input size="medium" maxlength="32"
				 clearable v-model="addData.type" placeholder="请输入类型">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select size="medium" style="width: 100%;"	v-model="addData.status"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="排序值" prop="sort_no">
				<el-input size="medium" maxlength="2"
				 clearable v-model="addData.sort_no" placeholder="请输入排序值">
				</el-input>
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
			status:this.$enum.get("status"),
			rules: {                    //数据验证规则
				name: [{ required: true, message: "请输入名称", trigger: "blur" }],
				value: [{ required: true, message: "请输入值", trigger: "blur" }],
				type: [{ required: true, message: "请输入类型", trigger: "blur" }],
				status: [{ required: true, message: "请输入状态", trigger: "blur" }],
				sort_no: [{ required: true, message: "请输入排序值", trigger: "blur" }],
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
					this.$http.post("/dictionary/info", this.addData, {}, true, true)
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
