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

// WarehouseBaseDao is the data access object for table invoicing_warehouse_base.
type WarehouseBaseDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns WarehouseBaseColumns // columns contains all the column names of Table for convenient usage.
}

// WarehouseBaseColumns defines and stores column names for table invoicing_warehouse_base.
type WarehouseBaseColumns struct {
	WarehouseId           string // 仓库编号
	WarehouseNumber       string // 仓库编号
	WarehouseName         string // 仓库名称
	WarehouseAddress      string // 仓库地址
	WarehouseTypeId       string // 仓库类型编号
	WarehouseEnable       string // 仓库是否启用(BOOL):0-禁用;1-启用
	WarehouseIsDefault    string // 仓库是否默认(BOOL):0-非默认;1-默认仓库
	WarehouseContact      string // 仓库联系人
	WarehousePhone        string // 仓库联系方式
	WarehouseDeliveryType string // 发货属性(ENUM):1-发货仓库; 0-备货仓库
	WarehouseDeliveryIds  string // 备货仓库编号(DOT)如果为发货仓库绑定的备货仓库编号
	StoreId               string // 所属店铺
	PrimaryKey            string // 主键
}

// warehouseBaseColumns holds the columns for table invoicing_warehouse_base.
var warehouseBaseColumns = WarehouseBaseColumns{
	WarehouseId:           "warehouse_id",
	WarehouseNumber:       "warehouse_number",
	WarehouseName:         "warehouse_name",
	WarehouseAddress:      "warehouse_address",
	WarehouseTypeId:       "warehouse_type_id",
	WarehouseEnable:       "warehouse_enable",
	WarehouseIsDefault:    "warehouse_is_default",
	WarehouseContact:      "warehouse_contact",
	WarehousePhone:        "warehouse_phone",
	WarehouseDeliveryType: "warehouse_delivery_type",
	WarehouseDeliveryIds:  "warehouse_delivery_ids",
	StoreId:               "store_id",
	PrimaryKey:            "warehouse_id",
}

// NewWarehouseBaseDao creates and returns a new DAO object for table data access.
func NewWarehouseBaseDao() *WarehouseBaseDao {
	return &WarehouseBaseDao{
		group:   "invoicing",
		table:   "invoicing_warehouse_base",
		columns: warehouseBaseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WarehouseBaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WarehouseBaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WarehouseBaseDao) Columns() WarehouseBaseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WarehouseBaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WarehouseBaseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WarehouseBaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *WarehouseBaseDao) Get(ctx context.Context, id any) (one *entity.WarehouseBase, err error) {
	var entitys []*entity.WarehouseBase
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
func (dao *WarehouseBaseDao) Gets(ctx context.Context, id any) (entitys []*entity.WarehouseBase, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *WarehouseBaseDao) Find(ctx context.Context, in *do.WarehouseBaseListInput) (out []*entity.WarehouseBase, err error) {
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
func (dao *WarehouseBaseDao) FindOne(ctx context.Context, in *do.WarehouseBaseListInput) (one *entity.WarehouseBase, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.WarehouseBase
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
func (dao *WarehouseBaseDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.WarehouseBaseListInput) (out gdb.Result, err error) {
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
func (dao *WarehouseBaseDao) FindKey(ctx context.Context, in *do.WarehouseBaseListInput) (out []interface{}, err error) {
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
func (dao *WarehouseBaseDao) List(ctx context.Context, in *do.WarehouseBaseListInput) (out *do.WarehouseBaseListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.WarehouseBaseListOutput{}
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
func (dao *WarehouseBaseDao) Add(ctx context.Context, in *do.WarehouseBase) (lastInsertId int64, err error) {
	data := do.WarehouseBase{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *WarehouseBaseDao) Edit(ctx context.Context, id any, in *do.WarehouseBase) (int64, error) {
	data := do.WarehouseBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *WarehouseBaseDao) EditWhere(ctx context.Context, where *do.WarehouseBaseListInput, in *do.WarehouseBase) (int64, error) {
	data := do.WarehouseBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *WarehouseBaseDao) Save(ctx context.Context, in *do.WarehouseBase) (affected int64, err error) {
	data := do.WarehouseBase{}
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
func (dao *WarehouseBaseDao) Saves(ctx context.Context, in []*do.WarehouseBase) (affected int64, err error) {
	data := []do.WarehouseBase{}
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
func (dao *WarehouseBaseDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *WarehouseBaseDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *WarehouseBaseDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *WarehouseBaseDao) RemoveWhere(ctx context.Context, where *do.WarehouseBaseListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *WarehouseBaseDao) Count(ctx context.Context, in *do.WarehouseBaseListInput) (count int, err error) {
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
