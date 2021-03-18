<template>
	<el-dialog title="编辑供货商商品" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="供货商商品:" prop="spp_product_no">
				<el-input size="medium" maxlength="32"
				clearable v-model="editData.spp_product_no" placeholder="请输入供货商商品">
				</el-input>
      </el-form-item>
      
      <el-form-item label="成本折扣:" prop="cost_discount">
				<el-input size="medium" maxlength="10" oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				clearable v-model="editData.cost_discount" placeholder="请输入成本折扣">
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
				cost_discount: [
					{ required: true, message: "请输入成本折扣", trigger: "blur" }
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
			var spp_product_id = this.editData.spp_product_id
			this.editData = this.$http.xget("/supplier/product", { spp_product_id: spp_product_id })
			this.editData.spp_product_id = spp_product_id
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/supplier/product", this.editData, {}, true, true)
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
