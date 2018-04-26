package orm

import "github.com/go-xorm/xorm"

// XormPlus xorm 的扩展版本
type XormPlus struct {
	// xorm
	*xorm.Session

	// 分页每页数量
	pageNum int

	// 字段映射规则
	fieldMap func(string) string
}

// NewXormPlus 初始化
func NewXormPlus(xorm *xorm.Session) (*XormPlus, error) {
	return &XormPlus{Session: xorm}, nil
}

// Fetch 获取单条数据
func (orm *XormPlus) Fetch(inst interface{}, filter interface{}, conds ...interface{}) error {
	_, err := orm.Where(filter, conds...).Get(inst)
	return err
}

// List 获取数据列表
func (orm *XormPlus) List(inst interface{}, filter interface{}, conds ...interface{}) error {
	return orm.Where(filter, conds...).Find(inst)
}

// Update 更新
func (orm *XormPlus) Update(inst interface{}, cols []string, filter interface{}, conds ...interface{}) (int64, error) {
	return orm.Cols(cols...).Where(filter, conds...).Update(inst)
}

// SetPageNavi 设置分页参数
func (orm *XormPlus) SetPageNavi(setting map[string]int) *XormPlus {
	// 设置每页数量
	if pageNum, ok := setting["pageNum"]; ok {
		if pageNum > 0 {
			orm.pageNum = pageNum
		}
	}

	return orm
}

// Page 分页
func (orm *XormPlus) Page(n int, perNum ...int) *XormPlus {
	limit := orm.pageNum
	if len(perNum) > 0 && perNum[0] > 0 {
		limit = perNum[0]
	}

	offset := (n - 1) * limit
	if n < 1 {
		offset = 0
	}

	orm.Limit(limit, offset)

	return orm
}

// SetFieldMap 设置字段转换方式
// 默认方式是 hyphen2Hump
func (orm *XormPlus) SetFieldMap(fieldMap func(string) string) *XormPlus {
	orm.fieldMap = fieldMap
	return orm
}

// Collention 将给定的 ds 结果集 按照 mapKeys 的规则映射
func (orm *XormPlus) Collention(sqlStr string, args ...interface{}) (map[string]string, error) {
	list, err := orm.Collentions(sqlStr, args...)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list[0], nil
}

// Collentions 将给定的 ds 结果集 按照 mapKeys 的规则映射
func (orm *XormPlus) Collentions(sqlStr string, args ...interface{}) ([]map[string]string, error) {
	if orm.fieldMap == nil {
		orm.fieldMap = hyphen2Hump
	}

	sqlOrArgs := []interface{}{
		sqlStr,
	}

	for _, arg := range args {
		sqlOrArgs = append(sqlOrArgs, arg)
	}

	rawRes, err := orm.QueryString(sqlOrArgs...)

	if err != nil {
		return rawRes, err
	}

	res := make([]map[string]string, 0)
	for _, mp := range rawRes {
		it := make(map[string]string)
		for key, val := range mp {
			newKey := orm.fieldMap(key)
			it[newKey] = val
		}
		res = append(res, it)
	}

	return res, nil
}

// RawCollention 将给定的 ds 结果集 按照 mapKeys 的规则映射
func (orm *XormPlus) RawCollention(sqlStr string, args ...interface{}) (map[string]interface{}, error) {
	list, err := orm.RawCollentions(sqlStr, args...)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	return list[0], nil
}

// RawCollentions 将给定的 ds 结果集 按照 mapKeys 的规则映射
func (orm *XormPlus) RawCollentions(sqlStr string, args ...interface{}) ([]map[string]interface{}, error) {
	if orm.fieldMap == nil {
		orm.fieldMap = hyphen2Hump
	}

	sqlOrArgs := []interface{}{
		sqlStr,
	}

	for _, arg := range args {
		sqlOrArgs = append(sqlOrArgs, arg)
	}

	rawRes, err := orm.QueryInterface(sqlOrArgs...)

	if err != nil {
		return rawRes, err
	}

	res := make([]map[string]interface{}, 0)
	for _, mp := range rawRes {
		it := make(map[string]interface{})
		for key, val := range mp {
			newKey := orm.fieldMap(key)
			it[newKey] = val
		}
		res = append(res, it)
	}

	return res, nil
}
