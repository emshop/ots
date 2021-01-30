<template>
  <!-- Add Form -->
  <el-dialog title="添加供货商错误码" width="25%" :visible.sync="dialogAddVisible">
    <el-form :model="addData"  :rules="rules" ref="addForm" label-width="110px">
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
			rules: {                    //数据验证规则
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
