<template>
  <!-- Add Form -->
  <el-dialog title="添加账户信息" width="65%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="帐户名称" prop="account_name">
				<el-input size="medium" maxlength="32"
				 clearable v-model="addData.account_name" placeholder="请输入帐户名称">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="系统标识:" prop="ident">
				<el-select size="medium" style="width: 100%;"	v-model="addData.ident"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in ident" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="用户分组:" prop="groups">
				<el-select size="medium" style="width: 100%;"	v-model="addData.groups"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in groups" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="商户信息:" prop="eid">
				<el-select size="medium" style="width: 100%;"	v-model="addData.eid"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in eid" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="信用余额" prop="credit">
				<el-input size="medium" maxlength="20" oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				 clearable v-model="addData.credit" placeholder="请输入信用余额">
				</el-input>
      </el-form-item>
      
      <el-form-item label="状态" prop="status">
				<el-input size="medium" maxlength="1"
				 clearable v-model="addData.status" placeholder="请输入状态">
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
			ident:this.$enum.get("ident"),
			groups:this.$enum.get("account_group"),
			eid:this.$enum.get("merchant_info"),
			rules: {                    //数据验证规则
				account_name: [{ required: true, message: "请输入帐户名称", trigger: "blur" }],
				ident: [{ required: true, message: "请输入系统标识", trigger: "blur" }],
				groups: [{ required: true, message: "请输入用户分组", trigger: "blur" }],
				eid: [{ required: true, message: "请输入商户信息", trigger: "blur" }],
				credit: [{ required: true, message: "请输入信用余额", trigger: "blur" }],
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
					this.$http.post("/account/info", this.addData, {}, true, true)
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
