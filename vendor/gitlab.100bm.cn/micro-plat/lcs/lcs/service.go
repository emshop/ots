package lcs

import (
	"fmt"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/global"
	"github.com/micro-plat/lib4go/logger"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs/const/sql"
	"gitlab.100bm.cn/micro-plat/lcs/lcs/const/tags"
	"gitlab.100bm.cn/micro-plat/lcs/lcs/lifetime"
)

func init() {
	//加载数据表结构创建sql
	hydra.OnReadying(func() error {
		sql.Install()
		return nil
	})
}

/*Client 生命周期
 *tagNote 标注说明:【%s】开始: xxxxx, %s为要填入的内容
 *orderNo 业务单号(可以存其他有意义的数据)
 *batchNo 批次号
 *extralParam 扩展参数
 */
type Client struct {
	tagNote     string
	orderNo     string
	batchNo     string
	extralParam string
	log         logger.ILogger
}

//New tagNote 业务标注说明如：创建订单 orderNo:业务单号, batchNo:批次号, extralParam:扩展参数
//(最多支持三个业务数据,第一个必须传),后面两个可以不传
func New(log logger.ILogger, tagNote string, orderNumber ...string) *Client {
	return &Client{
		log:         log,
		tagNote:     tagNote,
		orderNo:     types.GetStringByIndex(orderNumber, 0),
		batchNo:     types.GetStringByIndex(orderNumber, 1),
		extralParam: types.GetStringByIndex(orderNumber, 2),
	}
}

//Start 创建生命周期成对数据-开始
func (l *Client) Start(content string) *Client {
	lifetime.Create(l.constructData(tags.Start, content), l.log)
	return l
}

//End 创建生命周期成对数据-结束, info可以为字符串或者为error
func (l *Client) End(info interface{}) *Client {
	lifetime.Create(l.constructData(tags.End, info), l.log)
	return l
}

//Create 创建生命周期数据-单个
func (l *Client) Create(content string) *Client {
	lifetime.Create(l.constructData(tags.Single, content), l.log)
	return l
}

//constructData 构造数据实体
func (l *Client) constructData(lifeType string, content interface{}) *lifetime.LifeTimeInfo {
	tempContent := fmt.Sprintf(tags.MessageFormats.GetString(lifeType), l.tagNote, l.getInputContent(content))
	return &lifetime.LifeTimeInfo{
		OrderNo:     l.orderNo,
		BatchNo:     l.batchNo,
		ExtralParam: l.extralParam,
		Content:     tempContent,
		IP:          global.LocalIP(),
	}
}

//getInputContent 处理类型结果
func (l *Client) getInputContent(content interface{}) string {
	if content == nil {
		return "【success】"
	}
	data, flag := content.(error)
	if !flag {
		return types.GetString(content)
	}
	return "【error】" + data.Error()
}
