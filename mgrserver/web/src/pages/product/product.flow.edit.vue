<template>
	<el-dialog title="编辑业务流程" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData" :inline="true" :rules="rules" ref="editForm" label-width="110px">
      
			<el-form-item label="流程名称:" prop="flow_tag">
				<el-select size="medium" style="width: 100%;"	v-model="editData.flow_tag" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in flowTag" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="产品线:" prop="pl_id">
				<el-select size="medium" style="width: 100%;"	v-model="editData.pl_id" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="队列名称:" prop="queue_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.queue_name" placeholder="请输入队列名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="执行间隔:" prop="scan_interval">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.scan_interval" placeholder="请输入执行间隔">
				</el-input>
      </el-form-item>
      
      <el-form-item label="延后时长:" prop="delay">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.delay" placeholder="请输入延后时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="超时时长:" prop="timeout">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.timeout" placeholder="请输入超时时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="最大次数:" prop="max_count">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.max_count" placeholder="请输入最大次数">
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
      flowTag: this.$enum.get("flow_tag"),
      plID: this.$enum.get("product_line"),
      status: this.$enum.get("status"),
			rules: {                    //数据验证规则
				flow_tag: [
					{ required: true, message: "请输入流程名称", trigger: "blur" }
				],
				pl_id: [
					{ required: true, message: "请输入产品线", trigger: "blur" }
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
			var flow_id = this.editData.flow_id
			this.editData = this.$http.xget("/product/flow", { flow_id: flow_id })
			this.editData.flow_id = flow_id
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/product/flow", this.editData, {}, true, true)
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
