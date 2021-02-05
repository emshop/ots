package sql
//GetSystemTaskByTaskID 查询单条数据任务表
const GetSystemTaskByTaskID = `
select
	t.task_id,
	t.order_no,
	t.name,
	t.create_time,
	t.last_execute_time,
	t.next_execute_time,
	t.max_execute_time,
	t.next_interval,
	t.delete_interval,
	t.delete_time,
	t.count,
	t.max_count,
	t.status,
	t.batch_id,
	t.queue_name,
	t.msg_content
from tsk_system_task t
where
	&task_id`

//GetSystemTaskListCount 获取任务表列表条数
const GetSystemTaskListCount = `
select count(1)
from tsk_system_task t
where
	&t.order_no
	and t.create_time >= @create_time 
	and t.create_time < date_add(@create_time, interval 1 day)`

//GetSystemTaskList 查询任务表列表数据
const GetSystemTaskList = `
select
	t.task_id,
	t.order_no,
	t.name,
	t.create_time,
	t.last_execute_time,
	t.next_execute_time,
	t.count,
	t.status,
	t.queue_name 
from tsk_system_task t
where
	&t.order_no
	and t.create_time >= @create_time 
	and t.create_time < date_add(@create_time, interval 1 day)
order by t.task_id desc
limit @ps offset @offset
`
