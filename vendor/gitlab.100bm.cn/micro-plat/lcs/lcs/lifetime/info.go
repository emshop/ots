package lifetime

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

/*LifeTimeInfo 生命周期实体信息
  Ident 系统标识(如sso)
  Custom1 子系统自定义字段，最好在子系统中唯一
  Content 要保存的内容
*/
type LifeTimeInfo struct {
	OrderNo     string `json:"order_no" valid:"required"`
	BatchNo     string `json:"batch_no"`
	ExtralParam string `json:"extral_param"`
	Content     string `json:"content" valid:"required"`
	IP          string `json:"ip"`
}

//Valid 验证实体数据
func (l *LifeTimeInfo) Valid() error {
	if b, err := govalidator.ValidateStruct(l); !b {
		return fmt.Errorf("参数传入有误有误:%v", err)
	}
	return nil
}
