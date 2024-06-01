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

// StoreTransportType is the golang structure for table store_transport_type.
type StoreTransportType struct {
	TransportTypeId            uint        `json:"transport_type_id"             ` // 物流及售卖区域编号
	TransportTypeName          string      `json:"transport_type_name"           ` // 模板名称
	StoreId                    uint        `json:"store_id"                      ` // 所属店铺
	ChainId                    uint        `json:"chain_id"                      ` // 門店编号
	TransportTypePricingMethod uint        `json:"transport_type_pricing_method" ` // 计费规则(ENUM):1-按件数;2-按重量;3-按体积
	TransportTypeFreightFree   float64     `json:"transport_type_freight_free"   ` // 免运费额度
	TransportTypeBuildin       bool        `json:"transport_type_buildin"        ` // 系统内置(BOOL):0-非内置;1-内置
	TransportTypeFree          bool        `json:"transport_type_free"           ` // 全免运费(BOOL):0-不全免;1-全免（不限制地区且免运费）
	UpdateTime                 *gtime.Time `json:"update_time"                   ` // 编辑时间
}
