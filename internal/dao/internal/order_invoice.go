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

package internal

import (
	"context"
	"database/sql"
	"math"

	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// OrderInvoiceDao is the data access object for table trade_order_invoice.
type OrderInvoiceDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns OrderInvoiceColumns // columns contains all the column names of Table for convenient usage.
}

// OrderInvoiceColumns defines and stores column names for table trade_order_invoice.
type OrderInvoiceColumns struct {
	OrderInvoiceId        string // 发票编号
	OrderId               string // 订单编号
	UserId                string // 所属用户
	StoreId               string // 店铺编号
	InvoiceTitle          string // 发票抬头
	InvoiceContent        string // 发票内容
	InvoiceAmount         string // 开票金额
	InvoiceCompanyCode    string // 纳税人识别号
	InvoiceIsCompany      string // 公司开票(BOOL):0-个人;1-公司
	InvoiceIsElectronic   string // 电子发票(ENUM):0-纸质发票;1-电子发票
	InvoiceType           string // 发票类型(ENUM):1-普通发票;2-增值税专用发票
	InvoiceStatus         string // 开票状态(BOOL): 0-未开票; 1-已开票;
	InvoiceImg            string // 电子发票图片
	InvoiceDatetime       string // 开票时间
	InvoiceAddress        string // 单位地址
	InvoicePhone          string // 单位电话
	InvoiceBankname       string // 开户银行
	InvoiceBankaccount    string // 银行账号
	InvoiceContactName    string // 收票人
	InvoiceContactArea    string // 收票人地区
	InvoiceContactAddress string // 收票详细地址
	UserIntl              string // 国家编码
	UserMobile            string // 手机号码(mobile)
	UserEmail             string // 用户邮箱(email)
	OrderIsPaid           string // 是否支付(BOOL): 0-未支付; 1-已支付;
	InvoiceTime           string // 创建时间
	InvoiceCancel         string // 是否取消(BOOL): 0-未取消; 1-取消;
	PrimaryKey            string // 主键
}

// orderInvoiceColumns holds the columns for table trade_order_invoice.
var orderInvoiceColumns = OrderInvoiceColumns{
	OrderInvoiceId:        "order_invoice_id",
	OrderId:               "order_id",
	UserId:                "user_id",
	StoreId:               "store_id",
	InvoiceTitle:          "invoice_title",
	InvoiceContent:        "invoice_content",
	InvoiceAmount:         "invoice_amount",
	InvoiceCompanyCode:    "invoice_company_code",
	InvoiceIsCompany:      "invoice_is_company",
	InvoiceIsElectronic:   "invoice_is_electronic",
	InvoiceType:           "invoice_type",
	InvoiceStatus:         "invoice_status",
	InvoiceImg:            "invoice_img",
	InvoiceDatetime:       "invoice_datetime",
	InvoiceAddress:        "invoice_address",
	InvoicePhone:          "invoice_phone",
	InvoiceBankname:       "invoice_bankname",
	InvoiceBankaccount:    "invoice_bankaccount",
	InvoiceContactName:    "invoice_contact_name",
	InvoiceContactArea:    "invoice_contact_area",
	InvoiceContactAddress: "invoice_contact_address",
	UserIntl:              "user_intl",
	UserMobile:            "user_mobile",
	UserEmail:             "user_email",
	OrderIsPaid:           "order_is_paid",
	InvoiceTime:           "invoice_time",
	InvoiceCancel:         "invoice_cancel",
	PrimaryKey:            "order_invoice_id",
}

// NewOrderInvoiceDao creates and returns a new DAO object for table data access.
func NewOrderInvoiceDao() *OrderInvoiceDao {
	return &OrderInvoiceDao{
		group:   "trade",
		table:   "trade_order_invoice",
		columns: orderInvoiceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderInvoiceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderInvoiceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderInvoiceDao) Columns() OrderInvoiceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderInvoiceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderInvoiceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderInvoiceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *OrderInvoiceDao) Get(ctx context.Context, id any) (one *entity.OrderInvoice, err error) {
	var entitys []*entity.OrderInvoice
	entitys, err = dao.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(entitys) > 0 {
		one = entitys[0]
	}

	return one, err
}

// Gets 读取多条记录
func (dao *OrderInvoiceDao) Gets(ctx context.Context, id any) (entitys []*entity.OrderInvoice, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *OrderInvoiceDao) Find(ctx context.Context, in *do.OrderInvoiceListInput) (out []*entity.OrderInvoice, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	// 排序
	query = ml.BuildOrder(query, in.Sidx, in.Sort)
	if len(in.Order) > 0 {
		for _, it := range in.Order {
			query = ml.BuildOrder(query, it.Sidx, it.Sort)
		}
	}

	// 对象转换
	if err := query.Scan(&out); err != nil {
		return out, err
	}

	return out, nil
}

// FindOne 查询一条数据
func (dao *OrderInvoiceDao) FindOne(ctx context.Context, in *do.OrderInvoiceListInput) (one *entity.OrderInvoice, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.OrderInvoice
	entitys, err = dao.Find(ctx, in)

	if err != nil {
		return nil, err
	}

	if len(entitys) > 0 {
		one = entitys[0]
	}

	return one, err
}

// Find 查询字段数据
func (dao *OrderInvoiceDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.OrderInvoiceListInput) (out gdb.Result, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Fields(fieldNamesOrMapStruct).Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	// 排序
	query = ml.BuildOrder(query, in.Sidx, in.Sort)
	if len(in.Order) > 0 {
		for _, it := range in.Order {
			query = ml.BuildOrder(query, it.Sidx, it.Sort)
		}
	}

	out, err = query.All()

	if err != nil {
		return out, err
	}

	return out, nil
}

// FindKey 查询主键数据
func (dao *OrderInvoiceDao) FindKey(ctx context.Context, in *do.OrderInvoiceListInput) (out []interface{}, err error) {
	idRes, err := dao.FindFields(ctx, dao.Columns().PrimaryKey, in)

	if err != nil {
		return nil, err
	}

	for _, record := range idRes {
		if !record[dao.Columns().PrimaryKey].IsEmpty() {
			out = append(out, record[dao.Columns().PrimaryKey])
		}
	}

	return out, err
}

// List 分页读取
func (dao *OrderInvoiceDao) List(ctx context.Context, in *do.OrderInvoiceListInput) (out *do.OrderInvoiceListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.OrderInvoiceListOutput{}
	out.Page = in.Page
	out.Size = in.Size

	// 查询记录总数
	count, err1 := query.Count()
	if err1 != nil {
		return nil, err1
	}

	out.Records = count
	out.Total = int(math.Ceil(float64(count) / float64(out.Size)))

	// 排序
	query = ml.BuildOrder(query, in.Sidx, in.Sort)
	if len(in.Order) > 0 {
		for _, it := range in.Order {
			query = ml.BuildOrder(query, it.Sidx, it.Sort)
		}
	}

	// 分页
	query = query.Page(in.Page, in.Size)

	// 对象转换
	if err := query.Scan(&out.Items); err != nil {
		return out, err
	}

	return out, nil
}

// Add 新增
func (dao *OrderInvoiceDao) Add(ctx context.Context, in *do.OrderInvoice) (lastInsertId int64, err error) {
	data := do.OrderInvoice{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *OrderInvoiceDao) Edit(ctx context.Context, id any, in *do.OrderInvoice) (int64, error) {
	data := do.OrderInvoice{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *OrderInvoiceDao) EditWhere(ctx context.Context, where *do.OrderInvoiceListInput, in *do.OrderInvoice) (int64, error) {
	data := do.OrderInvoice{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *OrderInvoiceDao) Save(ctx context.Context, in *do.OrderInvoice) (affected int64, err error) {
	data := do.OrderInvoice{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	res, err := dao.Ctx(ctx).Data(data).OmitNil().Save()

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Saves 批量保存
func (dao *OrderInvoiceDao) Saves(ctx context.Context, in []*do.OrderInvoice) (affected int64, err error) {
	data := []do.OrderInvoice{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	res, err := dao.Ctx(ctx).Data(data).OmitNil().Save()

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Increment 增加
func (dao *OrderInvoiceDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *OrderInvoiceDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *OrderInvoiceDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *OrderInvoiceDao) RemoveWhere(ctx context.Context, where *do.OrderInvoiceListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *OrderInvoiceDao) Count(ctx context.Context, in *do.OrderInvoiceListInput) (count int, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	//记录数
	count, err = query.Count()

	if err != nil {
		return 0, err
	}

	return count, nil
}
