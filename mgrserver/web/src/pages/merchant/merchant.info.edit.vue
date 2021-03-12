<template>
	<el-dialog title="编辑商户信息" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="商户名称:" prop="mer_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.mer_name" placeholder="请输入商户名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="公司名称:" prop="mer_crop">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.mer_crop" placeholder="请输入公司名称">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="类型:" prop="mer_type">
				<el-select size="medium" style="width: 100%;"	v-model="editData.mer_type" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in merType" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="商务人员:" prop="bd_uid">
				<el-select size="medium" style="width: 100%;"	v-model="editData.bd_uid" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in bdUid" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item  label="状态:" prop="status">
				<el-radio-group  size="medium" v-model="editData.status" style="margin-left:5px">
        	<el-radio v-for="(item, index) in status" :key="index" :label="item.value">{{item.name}}</el-radio>
				</el-radio-group>
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
      merType: this.$enum.get("merchant_type"),
      bdUid: this.$enum.get("user_info"),
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				mer_name: [
					{ required: true, message: "请输入商户名称", trigger: "blur" }
				],
				mer_type: [
					{ required: true, message: "请输入类型", trigger: "blur" }
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
			var mer_no = this.editData.mer_no
			this.editData = this.$http.xget("/merchant/info", { mer_no: mer_no })
			this.editData.mer_no = mer_no
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
