<template>
	<el-dialog title="编辑字典配置" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="名称" prop="name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.name" placeholder="请输入名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="值" prop="value">
				<el-input size="medium" maxlength="32"
				clearable v-model="editData.value" placeholder="请输入值">
				</el-input>
      </el-form-item>
      
      <el-form-item label="类型" prop="type">
				<el-input size="medium" maxlength="32"
				clearable v-model="editData.type" placeholder="请输入类型">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select size="medium" style="width: 100%;"	v-model="editData.status"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="排序值" prop="sort_no">
				<el-input size="medium" maxlength="2"
				clearable v-model="editData.sort_no" placeholder="请输入排序值">
				</el-input>
      </el-form-item>
      
    </el-form>
		<div slot="footer" class="dialog-footer">
			<el-button size="medium" @click="dialogFormVisible = false">取 消</el-button>
			<el-button type="success" size="medium" @click="edit">确 定</el-button>
		</div>
	</el-dialog>
</template>

<script>
export default {
	data() {
		return {
			dialogFormVisible: false,    //编辑表单显示隐藏
			editData: {},                //编辑数据对象
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				name: [
					{ required: true, message: "请输入名称", trigger: "blur" }
				],
				value: [
					{ required: true, message: "请输入值", trigger: "blur" }
				],
				type: [
					{ required: true, message: "请输入类型", trigger: "blur" }
				],
				status: [
					{ required: true, message: "请输入状态", trigger: "blur" }
				],
				sort_no: [
					{ required: true, message: "请输入排序值", trigger: "blur" }
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
			this.editData = this.$http.xget("/dictionary/info", { id: this.editData.id })
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/dictionary/info", this.editData, {}, true, true)
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
