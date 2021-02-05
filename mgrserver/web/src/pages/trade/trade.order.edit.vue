<template>
	<el-dialog title="编辑订单记录[ots_trade_order][(row(btn:处理/deal)),tab(发货/ots_trade_delivery)]" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
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
			this.editData = this.$http.xget("/trade/order", { order_id: this.editData.order_id })
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/trade/order", this.editData, {}, true, true)
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
