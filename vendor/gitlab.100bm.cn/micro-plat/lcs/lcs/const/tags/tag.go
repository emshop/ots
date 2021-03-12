package tags

import "github.com/micro-plat/lib4go/types"

const (
	//Start 成对开始
	Start = "start"

	//End 成对结束
	End = "end"

	//Single 当个记录
	Single = "single"
)

//MessageFormats 生命周期消息格式
var MessageFormats = types.XMap{
	Start:  "【%s】开始: %s",
	End:    "【%s】结束: %s",
	Single: "【%s】%s",
}
