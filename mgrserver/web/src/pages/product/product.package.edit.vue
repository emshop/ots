<template>
	<el-dialog title="编辑产品包" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="包名称:" prop="pg_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.pg_name" placeholder="请输入包名称">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="产品线:" prop="pl_id">
				<el-select size="medium" style="width: 100%;"	v-model="editData.pl_id" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="商品面值:" prop="face">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.face" placeholder="请输入商品面值">
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
      plID: this.$enum.get("product_line"),
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				pg_name: [
					{ required: true, message: "请输入包名称", trigger: "blur" }
				],
				pl_id: [
					{ required: true, message: "请输入产品线", trigger: "blur" }
				],
				face: [
					{ required: true, message: "请输入商品面值", trigger: "blur" }
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
			var pg_id = this.editData.pg_id
			this.editData = this.$http.xget("/product/package", { pg_id: pg_id })
			this.editData.pg_id = pg_id
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/product/package", this.editData, {}, true, true)
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
