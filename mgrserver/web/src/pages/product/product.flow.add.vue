<template>
  <!-- Add Form -->
  <el-dialog title="添加业务流程" width="65%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="流程名称" prop="flow_name">
				<el-input maxlength="64" clearable v-model="addData.flow_name" placeholder="请输入流程名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="tag标签" prop="tag_name">
				<el-input maxlength="64" clearable v-model="addData.tag_name" placeholder="请输入tag标签">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="产品线:" prop="pl_id">
				<el-select  placeholder="---请选择---" clearable v-model="addData.pl_id" style="width: 100%;">
					<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="成功流程:" prop="success_flow_id">
				<el-select  placeholder="---请选择---" clearable v-model="addData.success_flow_id" style="width: 100%;">
					<el-option v-for="(item, index) in successFlowID" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="失败流程:" prop="failed_flow_id">
				<el-select  placeholder="---请选择---" clearable v-model="addData.failed_flow_id" style="width: 100%;">
					<el-option v-for="(item, index) in failedFlowID" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="未知流程:" prop="unknown_flow_id">
				<el-select  placeholder="---请选择---" clearable v-model="addData.unknown_flow_id" style="width: 100%;">
					<el-option v-for="(item, index) in unknownFlowID" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="队列名称" prop="queue_name">
				<el-input maxlength="64" clearable v-model="addData.queue_name" placeholder="请输入队列名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="超时时长" prop="scan_interval">
				<el-input maxlength="10" clearable v-model="addData.scan_interval" placeholder="请输入超时时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="延后处理时长" prop="delay">
				<el-input maxlength="10" clearable v-model="addData.delay" placeholder="请输入延后处理时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="超时时长" prop="timeout">
				<el-input maxlength="10" clearable v-model="addData.timeout" placeholder="请输入超时时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="最大执行次数" prop="max_count">
				<el-input maxlength="10" clearable v-model="addData.max_count" placeholder="请输入最大执行次数">
				</el-input>
      </el-form-item>
      
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button size="small" @click="resetForm('addForm')">取 消</el-button>
      <el-button size="small" type="success" @click="add('addForm')">确 定</el-button>
    </div>
  </el-dialog>
  <!--Add Form -->
</template>

<script>
export default {
	data() {
		return {
			addData: {},
			dialogAddVisible: false,
      plID: this.$enum.get("product_line"),
      successFlowID: this.$enum.get("product_flow"),
      failedFlowID: this.$enum.get("product_flow"),
      unknownFlowID: this.$enum.get("product_flow"),
			rules: {                    //数据验证规则
				flow_name: [{ required: true, message: "请输入流程名称", trigger: "blur" }],
				tag_name: [{ required: true, message: "请输入tag标签", trigger: "blur" }],
				pl_id: [{ required: true, message: "请输入产品线", trigger: "blur" }],
				success_flow_id: [{ required: true, message: "请输入成功流程", trigger: "blur" }],
				failed_flow_id: [{ required: true, message: "请输入失败流程", trigger: "blur" }],
				unknown_flow_id: [{ required: true, message: "请输入未知流程", trigger: "blur" }],
				queue_name: [{ required: true, message: "请输入队列名称", trigger: "blur" }],
				scan_interval: [{ required: true, message: "请输入超时时长", trigger: "blur" }],
				delay: [{ required: true, message: "请输入延后处理时长", trigger: "blur" }],
				timeout: [{ required: true, message: "请输入超时时长", trigger: "blur" }],
				max_count: [{ required: true, message: "请输入最大执行次数", trigger: "blur" }],
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
		resetForm(formName) {
			this.dialogAddVisible = false;
			this.$refs[formName].resetFields();
		},
		show(){
			this.dialogAddVisible = true;
		},
		add(formName) {
			this.$refs[formName].validate((valid) => {
				if (valid) {
					this.$http.post("/product/flow", this.addData, {}, true, true)
						.then(res => {
							this.$refs[formName].resetFields()
							this.dialogAddVisible = false
							this.refresh()
						})
						.catch(err => {
							this.$message({
								type: "error",
								message: err.response.data
							});
						})
				} else {
						console.log("error submit!!");
						return false;
				}
			});
		},
	}

}
</script>

<style scoped>
</style>
