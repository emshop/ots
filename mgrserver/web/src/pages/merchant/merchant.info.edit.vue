<template>
	<el-dialog title="编辑商户信息" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="名称" prop="mer_name">
				<el-input clearable v-model="editData.mer_name" placeholder="请输入名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="公司" prop="mer_crop">
				<el-input clearable v-model="editData.mer_crop" placeholder="请输入公司">
				</el-input>
      </el-form-item>
      
      <el-form-item label="类型" prop="mer_type">
				<el-input clearable v-model="editData.mer_type" placeholder="请输入类型">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="商务人员:" prop="bd_uid">
				<el-select  placeholder="---请选择---" clearable v-model="editData.bd_uid" style="width: 100%;">
					<el-option v-for="(item, index) in bdUid" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
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
      bdUid:this.$enum.get("user_info"),
      status:this.$enum.get("status"),
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
		show() {
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/merchant/info", this.editData, {}, true, true)
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
