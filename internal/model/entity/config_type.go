// +----------------------------------------------------------------------
// | ShopSuite商城系统 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 随商信息技术（上海）有限公司
// +----------------------------------------------------------------------
// | 未获商业授权前，不得将本软件用于商业用途。禁止整体或任何部分基础上以发展任何派生版本、
// | 修改版本或第三方版本用于重新分发。
// +----------------------------------------------------------------------
// | 官方网站: https://www.shopsuite.cn  https://www.golershop.cn
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本公司对该软件产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细见https://www.golershop.cn/policy
// +----------------------------------------------------------------------

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigType is the golang structure for table config_type.
type ConfigType struct {
	ConfigTypeId      uint        `json:"config_type_id"      ` // 分组编号
	ConfigTypeName    string      `json:"config_type_name"    ` // 分组名称
	ConfigTypeSort    uint        `json:"config_type_sort"    ` // 分组排序:从小到大
	ConfigTypeEnable  bool        `json:"config_type_enable"  ` // 是否有效(BOOL):0-禁用;1-启用
	ConfigTypeBuildin bool        `json:"config_type_buildin" ` // 系统内置(BOOL):1-是; 0-否
	UpdateTime        *gtime.Time `json:"update_time"         ` // 更新时间
	CreateTime        *gtime.Time `json:"create_time"         ` // 添加时间
	ConfigTypeModule  uint        `json:"config_type_module"  ` // 所属模块(ENUM):1001-站点设置;1002-上传设置;1003-运营设置;1004-财务设置;
}
