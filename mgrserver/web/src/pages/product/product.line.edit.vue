<template>
	<el-dialog title="编辑产品线" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="产品线名称:" prop="pl_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.pl_name" placeholder="请输入产品线名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="数量:" prop="num">
				<el-input size="medium" maxlength="2"
				clearable v-model="editData.num" placeholder="请输入数量">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select size="medium" style="width: 100%;"	v-model="editData.status" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
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
				pl_name: [
					{ required: true, message: "请输入产品线名称", trigger: "blur" }
				],
				num: [
					{ required: true, message: "请输入数量", trigger: "blur" }
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
			var pl_id = this.editData.pl_id
			this.editData = this.$http.xget("/product/line", { pl_id: pl_id })
			this.editData.pl_id = pl_id
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/product/line", this.editData, {}, true, true)
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
