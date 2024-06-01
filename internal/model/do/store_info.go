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

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// StoreInfo is the golang structure of table shop_store_info for DAO operations like Where/Data.
type StoreInfo struct {
	g.Meta            `orm:"table:shop_store_info, do:true"`
	StoreId           interface{} // 店铺编号
	StoreSlogan       interface{} // 店铺口号
	StoreBanner       interface{} // 店铺banner
	StoreTel          interface{} // 卖家电话
	StoreFreeShipping interface{} // 免运费额度
	StoreTemplate     interface{} // 店铺绑定模板
	StoreStateId      interface{} // 店铺资料信息状态(ENUM):3210-待完善资料;   3220-等待审核 ;  3230-资料审核没有通过;     3240-资料审核通过,待付款
	StorePaymentState interface{} // 付款状态(ENUM):0-未付款;   1-已付款待审核;   2-审核通过
	StoreStartTime    *gtime.Time // 本次开始时间
	StoreEndTime      *gtime.Time // 有效期截止时间
	StoreCloseReason  interface{} // 关闭原因
	StoreNotice       interface{} // 购买须知
	StoreOpeningHours interface{} // 营业时间
	StoreCloseHours   interface{} // 打烊时间
	StoreArea         interface{} // 所在区域
	StoreDistrictId   interface{} // 所属地区(DOT)
	StoreAddress      interface{} // 店铺详细地址
}

type StoreInfoListInput struct {
	ml.BaseList
	Where StoreInfo // 查询条件
}

type StoreInfoListOutput struct {
	Items   []*entity.StoreInfo // 列表
	Page    int                 // 分页号码
	Total   int                 // 总页数
	Records int                 // 数据总数
	Size    int                 // 单页数量
}

type StoreInfoListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
