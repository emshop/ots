<template>
	<el-dialog title="编辑商户商品" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="商户商品编号" prop="mer_product_no">
				<el-input maxlength="32" 
				
				clearable v-model="editData.mer_product_no" placeholder="请输入商户商品编号">
				</el-input>
      </el-form-item>
      
      <el-form-item label="销售折扣" prop="discount">
				<el-input maxlength="10" 
				 oninput="if(isNaN(value)) { value = null } if(value.indexOf('.')>0){value=value.slice(0,value.indexOf('.')+6)}"
				clearable v-model="editData.discount" placeholder="请输入销售折扣">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select style="width: 100%;"	v-model="editData.status"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
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
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				discount: [
					{ required: true, message: "请输入销售折扣", trigger: "blur" }
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
			this.editData = this.$http.xget("/merchant/product", { mer_product_id: this.editData.mer_product_id })
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/merchant/product", this.editData, {}, true, true)
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
