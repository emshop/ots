<template>
  <!-- Add Form -->
  <el-dialog title="添加商户信息" width="65%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="编号" prop="mer_no">
				<el-input maxlength="32" 
				
				 clearable v-model="addData.mer_no" placeholder="请输入编号">
				</el-input>
      </el-form-item>
      
      <el-form-item label="商户名称" prop="mer_name">
				<el-input maxlength="64" 
				
				 clearable v-model="addData.mer_name" placeholder="请输入商户名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="公司名称" prop="mer_crop">
				<el-input maxlength="64" 
				
				 clearable v-model="addData.mer_crop" placeholder="请输入公司名称">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="类型:" prop="mer_type">
				<el-select style="width: 100%;"	v-model="addData.mer_type"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in merType" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="商务人员:" prop="bd_uid">
				<el-select style="width: 100%;"	v-model="addData.bd_uid"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in bdUid" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item  label="状态:" prop="status">
				<el-radio-group v-model="addData.status" style="margin-left:5px">
        	<el-radio v-for="(item, index) in status" :key="index" :label="item.value">{{item.name}}</el-radio>
				</el-radio-group>
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
			merType:this.$enum.get("merchant_type"),
			bdUid:this.$enum.get("user_info"),
			status:this.$enum.get("status"),
			rules: {                    //数据验证规则
				mer_no: [{ required: true, message: "请输入编号", trigger: "blur" }],
				mer_name: [{ required: true, message: "请输入商户名称", trigger: "blur" }],
				mer_type: [{ required: true, message: "请输入类型", trigger: "blur" }],
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
					this.$http.post("/merchant/info", this.addData, {}, true, true)
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
