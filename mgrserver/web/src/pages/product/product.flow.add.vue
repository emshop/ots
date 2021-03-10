<template>
  <!-- Add Form -->
  <el-dialog title="添加业务流程" width="65%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="110px">
      
			<el-form-item label="流程名称:" prop="flow_tag">
				<el-select size="medium" style="width: 100%;"	v-model="addData.flow_tag"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in flowTag" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="产品线:" prop="pl_id">
				<el-select size="medium" style="width: 100%;"	v-model="addData.pl_id"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in plID" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="队列名称:" prop="queue_name">
				<el-input size="medium" maxlength="64"
				 clearable v-model="addData.queue_name" placeholder="请输入队列名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="执行间隔:" prop="scan_interval">
				<el-input size="medium" maxlength="10"
				 clearable v-model="addData.scan_interval" placeholder="请输入执行间隔">
				</el-input>
      </el-form-item>
      
      <el-form-item label="延后时长:" prop="delay">
				<el-input size="medium" maxlength="10"
				 clearable v-model="addData.delay" placeholder="请输入延后时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="超时时长:" prop="timeout">
				<el-input size="medium" maxlength="10"
				 clearable v-model="addData.timeout" placeholder="请输入超时时长">
				</el-input>
      </el-form-item>
      
      <el-form-item label="最大次数:" prop="max_count">
				<el-input size="medium" maxlength="10"
				 clearable v-model="addData.max_count" placeholder="请输入最大次数">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select size="medium" style="width: 100%;"	v-model="addData.status"	clearable filterable class="input-cos" placeholder="---请选择---">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
				</el-select>
			</el-form-item>
      
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button size="medium" @click="resetForm('addForm')">取 消</el-button>
      <el-button size="medium" type="success" @click="add('addForm')">确 定</el-button>
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
			flowTag:this.$enum.get("flow_tag"),
			plID:this.$enum.get("product_line"),
			status:this.$enum.get("status"),
			rules: {                    //数据验证规则
				flow_tag: [{ required: true, message: "请输入流程名称", trigger: "blur" }],
				pl_id: [{ required: true, message: "请输入产品线", trigger: "blur" }],
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
