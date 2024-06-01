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

package user

import (
	"context"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

type sUserInvoice struct{}

func init() {
	service.RegisterUserInvoice(NewUserInvoice())
}

func NewUserInvoice() *sUserInvoice {
	return &sUserInvoice{}
}

// Get 读取兑换码
func (s *sUserInvoice) Get(ctx context.Context, id any) (out *entity.UserInvoice, err error) {
	var list []*entity.UserInvoice
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条兑换码
func (s *sUserInvoice) Gets(ctx context.Context, id any) (list []*entity.UserInvoice, err error) {
	err = dao.UserInvoice.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sUserInvoice) Find(ctx context.Context, in *do.UserInvoiceListInput) (out []*entity.UserInvoice, err error) {
	out, err = dao.UserInvoice.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserInvoice) List(ctx context.Context, in *do.UserInvoiceListInput) (out *do.UserInvoiceListOutput, err error) {
	out, err = dao.UserInvoice.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserInvoice) Add(ctx context.Context, in *do.UserInvoice) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserInvoice.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserInvoice) Edit(ctx context.Context, in *do.UserInvoice) (affected int64, err error) {
	_, err = dao.UserInvoice.Edit(ctx, in.UserInvoiceId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserInvoice) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.UserInvoice.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}
