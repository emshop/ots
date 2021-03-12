<template>
	<el-dialog title="编辑供货商信息" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="供货商名称:" prop="spp_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.spp_name" placeholder="请输入供货商名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="所属公司:" prop="mer_crop">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.mer_crop" placeholder="请输入所属公司">
				</el-input>
      </el-form-item>
      
      <el-form-item label="商务人员:" prop="bd_uid">
				<el-input size="medium" maxlength="20"
				clearable v-model="editData.bd_uid" placeholder="请输入商务人员">
				</el-input>
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
      status: this.$enum.get("status"),
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
			var spp_no = this.editData.spp_no
			this.editData = this.$http.xget("/supplier/info", { spp_no: spp_no })
			this.editData.spp_no = spp_no
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
