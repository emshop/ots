<template>
	<el-dialog title="编辑业务流程" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData" :inline="true" :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="流程名称" prop="flow_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.flow_name" placeholder="请输入流程名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="tag标签" prop="tag_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.tag_name" placeholder="请输入tag标签">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="产品线:" prop="pl_id">
				<el-select size="medium" style="width: 100%;"	v-model="editData.pl_id"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="成功流程:" prop="success_flow_id">
				<el-select size="medium" placeholder="---请选择---" clearable filterable v-model="successFlowIDArray" multiple style="width: 100%;">
					<el-option v-for="(item, index) in successFlowID" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="失败流程:" prop="failed_flow_id">
				<el-select size="medium" placeholder="---请选择---" clearable filterable v-model="failedFlowIDArray" multiple style="width: 100%;">
					<el-option v-for="(item, index) in failedFlowID" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="未知流程:" prop="unknown_flow_id">
				<el-select size="medium" placeholder="---请选择---" clearable filterable v-model="unknownFlowIDArray" multiple style="width: 100%;">
					<el-option v-for="(item, index) in unknownFlowID" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="队列名称" prop="queue_name">
				<el-input size="medium" maxlength="64"
				clearable v-model="editData.queue_name" placeholder="请输入队列名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="执行间隔" prop="scan_interval">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.scan_interval" placeholder="请输入执行间隔">
				</el-input>
      </el-form-item>
      
      <el-form-item label="延后处理时长" prop="delay">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.delay" placeholder="请输入延后处理时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="超时时长" prop="timeout">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.timeout" placeholder="请输入超时时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="最大执行次数" prop="max_count">
				<el-input size="medium" maxlength="10"
				clearable v-model="editData.max_count" placeholder="请输入最大执行次数">
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
      plID: this.$enum.get("product_line"),
			successFlowID: this.$enum.get("product_flow"),
			successFlowIDArray: [],
			failedFlowID: this.$enum.get("product_flow"),
			failedFlowIDArray: [],
			unknownFlowID: this.$enum.get("product_flow"),
			unknownFlowIDArray: [],
			rules: {                    //数据验证规则
				flow_name: [
					{ required: true, message: "请输入流程名称", trigger: "blur" }
				],
				tag_name: [
					{ required: true, message: "请输入tag标签", trigger: "blur" }
				],
				pl_id: [
					{ required: true, message: "请输入产品线", trigger: "blur" }
				],
				success_flow_id: [
					{ required: true, message: "请输入成功流程", trigger: "blur" }
				],
				failed_flow_id: [
					{ required: true, message: "请输入失败流程", trigger: "blur" }
				],
				unknown_flow_id: [
					{ required: true, message: "请输入未知流程", trigger: "blur" }
				],
				queue_name: [
					{ required: true, message: "请输入队列名称", trigger: "blur" }
				],
				scan_interval: [
					{ required: true, message: "请输入执行间隔", trigger: "blur" }
				],
				delay: [
					{ required: true, message: "请输入延后处理时长", trigger: "blur" }
				],
				timeout: [
					{ required: true, message: "请输入超时时长", trigger: "blur" }
				],
				max_count: [
					{ required: true, message: "请输入最大执行次数", trigger: "blur" }
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
			this.editData = this.$http.xget("/product/flow", { flow_id: this.editData.flow_id })
			this.successFlowIDArray = this.editData.success_flow_id.split(",")
			this.failedFlowIDArray = this.editData.failed_flow_id.split(",")
			this.unknownFlowIDArray = this.editData.unknown_flow_id.split(",")
			this.dialogFormVisible = true;
		},
		edit() {
			this.editData.success_flow_id = this.successFlowIDArray.toString()
			this.editData.failed_flow_id = this.failedFlowIDArray.toString()
			this.editData.unknown_flow_id = this.unknownFlowIDArray.toString()
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
