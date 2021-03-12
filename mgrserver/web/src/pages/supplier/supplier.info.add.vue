<template>
  <!-- Add Form -->
  <el-dialog title="添加供货商信息" width="25%" :visible.sync="dialogAddVisible">
    <el-form :model="addData"  :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="编号:" prop="spp_no">
				<el-input size="medium" maxlength="32"
				 clearable v-model="addData.spp_no" placeholder="请输入编号">
				</el-input>
      </el-form-item>
      
      <el-form-item label="供货商名称:" prop="spp_name">
				<el-input size="medium" maxlength="64"
				 clearable v-model="addData.spp_name" placeholder="请输入供货商名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="所属公司:" prop="mer_crop">
				<el-input size="medium" maxlength="64"
				 clearable v-model="addData.mer_crop" placeholder="请输入所属公司">
				</el-input>
      </el-form-item>
      
      <el-form-item label="商务人员:" prop="bd_uid">
				<el-input size="medium" maxlength="20"
				 clearable v-model="addData.bd_uid" placeholder="请输入商务人员">
				</el-input>
      </el-form-item>
      
      
			<el-form-item  label="状态:" prop="status">
				<el-radio-group size="medium" v-model="addData.status" style="margin-left:5px">
        	<el-radio v-for="(item, index) in status" :key="index" :label="item.value">{{item.name}}</el-radio>
				</el-radio-group>
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
				spp_no: [{ required: true, message: "请输入编号", trigger: "blur" }],
				spp_name: [{ required: true, message: "请输入供货商名称", trigger: "blur" }],
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
					this.$http.post("/supplier/info", this.addData, {}, true, true)
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
