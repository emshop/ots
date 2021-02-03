<template>
	<el-dialog title="编辑供货商信息" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="供货商名称" prop="spp_name">
				<el-input maxlength="64" clearable v-model="editData.spp_name" placeholder="请输入供货商名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="所属公司" prop="mer_crop">
				<el-input maxlength="64" clearable v-model="editData.mer_crop" placeholder="请输入所属公司">
				</el-input>
      </el-form-item>
      
      <el-form-item label="商务人员" prop="bd_uid">
				<el-input maxlength="20" clearable v-model="editData.bd_uid" placeholder="请输入商务人员">
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
				spp_name: [
					{ required: true, message: "请输入供货商名称", trigger: "blur" }
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
			this.$http.put("/supplier/info", this.editData, {}, true, true)
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
