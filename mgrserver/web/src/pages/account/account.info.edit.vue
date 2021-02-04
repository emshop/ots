<template>
	<el-dialog title="编辑账户信息" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="帐户名称" prop="account_name">
				<el-input maxlength="32" clearable v-model="editData.account_name" placeholder="请输入帐户名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="信用余额" prop="credit">
				<el-input  clearable v-model="editData.credit" placeholder="请输入信用余额">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select  placeholder="---请选择---" clearable v-model="editData.status" style="width: 100%;">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
    </el-form>
		<div slot="footer" class="dialog-footer">
			<el-button size="small" @click="dialogFormVisible = false">取 消</el-button>
			<el-button type="success" size="small" @click="edit">确 定</el-button>
		</div>
	</el-dialog>
</template>

<script>
export default {
	data() {
		return {
			dialogFormVisible: false,    //编辑表单显示隐藏
			editData: {},                //编辑数据对象
      status:this.$enum.get("status"),
			rules: {                    //数据验证规则
				account_name: [
					{ required: true, message: "请输入帐户名称", trigger: "blur" }
				],
				credit: [
					{ required: true, message: "请输入信用余额", trigger: "blur" }
				],
				status: [
					{ required: true, message: "请输入状态", trigger: "blur" }
				],
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
		show() {
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/account/info", this.editData, {}, true, true)
			.then(res => {			
				this.dialogFormVisible = false;
				this.refresh()
			})
		},
	}
}
</script>

<style scoped>
</style>
