<template>
	<el-dialog title="编辑订单发货表" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      
			<el-form-item label="人工发货:" prop="is_mf">
				<el-select size="medium" style="width: 100%;"	v-model="editData.is_mf" clearable filterable class="input-cos" placeholder="---请选择---"	>
					<el-option v-for="(item, index) in isMf" :key="index" :value="item.value" :label="item.name"></el-option>
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
      isMf: this.$enum.get("bool"),
			rules: {                    //数据验证规则
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
			var delivery_id = this.editData.delivery_id
			this.editData = this.$http.xget("/trade/delivery", { delivery_id: delivery_id })
			this.editData.delivery_id = delivery_id
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/trade/delivery", this.editData, {}, true, true)
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
