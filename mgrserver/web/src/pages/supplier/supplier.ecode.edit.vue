<template>
	<el-dialog title="编辑供货商错误码" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData" :inline="true" :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="编号" prop="id">
				<el-input maxlength="20" 
				
				clearable v-model="editData.id" placeholder="请输入编号">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="供货商:" prop="spp_no">
				<el-select style="width: 100%;"	v-model="editData.spp_no"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in sppNo" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="产品线:" prop="pl_id">
				<el-select style="width: 100%;"	v-model="editData.pl_id"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="分类:" prop="category">
				<el-select style="width: 100%;"	v-model="editData.category"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in category" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="处理码:" prop="deal_code">
				<el-select style="width: 100%;"	v-model="editData.deal_code"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in dealCode" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="错误码" prop="error_code">
				<el-input maxlength="32" 
				
				clearable v-model="editData.error_code" placeholder="请输入错误码">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select style="width: 100%;"	v-model="editData.status"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="错误码描述" prop="error_desc">
				<el-input maxlength="64" 
				
				clearable v-model="editData.error_desc" placeholder="请输入错误码描述">
				</el-input>
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
      sppNo: this.$enum.get("supplier_info"),
      plID: this.$enum.get("product_line"),
      category: this.$enum.get("result_source"),
      dealCode: this.$enum.get("deal_code"),
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				id: [
					{ required: true, message: "请输入编号", trigger: "blur" }
				],
				spp_no: [
					{ required: true, message: "请输入供货商", trigger: "blur" }
				],
				pl_id: [
					{ required: true, message: "请输入产品线", trigger: "blur" }
				],
				category: [
					{ required: true, message: "请输入分类", trigger: "blur" }
				],
				deal_code: [
					{ required: true, message: "请输入处理码", trigger: "blur" }
				],
				error_code: [
					{ required: true, message: "请输入错误码", trigger: "blur" }
				],
				status: [
					{ required: true, message: "请输入状态", trigger: "blur" }
				],
				error_desc: [
					{ required: true, message: "请输入错误码描述", trigger: "blur" }
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
			this.editData = this.$http.xget("/supplier/ecode", { id: this.editData.id })
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/supplier/ecode", this.editData, {}, true, true)
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
