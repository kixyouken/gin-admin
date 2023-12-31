package service

import (
	"gin-admin/database"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sBaseService struct{}

var BaseService = sBaseService{}

var db = database.InitMysql()

// GetAll 查询所有
//
//	@receiver s
//	@param out
//	@param search
//	@return error
func (s *sBaseService) GetAll(out interface{}, search func(db *gorm.DB) *gorm.DB) error {
	return db.Scopes(search).Find(out).Error
}

// GetPage 分页查询（不关联表）
//
//	@receiver s
//	@param c
//	@param out
//	@param search
//	@param column
//	@param order
//	@return error
func (s *sBaseService) GetPage(c *gin.Context, out interface{}, search func(db *gorm.DB) *gorm.DB, column interface{}, order string) error {
	return db.Scopes(search, s.Paginate(c), s.Order(order)).Select(column).Find(out).Error
}

// GetPagePreload 分页查询（With关联表）
//
//	@receiver s
//	@param c
//	@param out
//	@param search
//	@param column
//	@param order
//	@param preload
//	@return error
func (s *sBaseService) GetPagePreload(c *gin.Context, out interface{}, search func(db *gorm.DB) *gorm.DB, column interface{}, order string, preload ...string) error {
	return db.Scopes(search, s.Paginate(c), s.Order(order), s.Preload(preload...)).Select(column).Find(out).Error
}

// GetPageJoins 分页查询（Join关联表）
//
//	@receiver s
//	@param c
//	@param out
//	@param search
//	@param column
//	@param order
//	@param joins
//	@return error
func (s *sBaseService) GetPageJoins(c *gin.Context, out interface{}, search func(db *gorm.DB) *gorm.DB, column interface{}, order string, joins ...string) error {
	return db.Scopes(search, s.Paginate(c), s.Order(order), s.Joins(joins...)).Select(column).Find(out).Error
}

// GetCount 查询数量
//
//	@receiver s
//	@param out
//	@param search
//	@return int64
func (s *sBaseService) GetCount(out interface{}, search func(db *gorm.DB) *gorm.DB) int64 {
	var count int64
	err := db.Model(out).Scopes(search).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// GetCountJoins 查询数量（Join关联表）
//
//	@receiver s
//	@param out
//	@param search
//	@param joins
//	@return int64
func (s *sBaseService) GetCountJoins(out interface{}, search func(db *gorm.DB) *gorm.DB, joins ...string) int64 {
	var count int64
	err := db.Model(out).Scopes(search, s.Joins(joins...)).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

// GetDistinct 去重（Distinct）
//
//	@receiver s
//	@param out
//	@param column
//	@return error
func (s *sBaseService) GetDistinct(out interface{}, column interface{}) error {
	return db.Distinct(column).Find(out).Error
}

// GetByID 根据ID获取
//
//	@receiver s
//	@param out
//	@param id
//	@return error
func (s *sBaseService) GetByID(out interface{}, id uint, preload ...string) error {
	return db.Scopes(s.Preload(preload...)).Limit(1).Where("id = ?", id).Find(out).Error
}

// GetByIDUnscoped 根据ID获取（不带deleted_at）
//
//	@receiver s
//	@param out
//	@param id
//	@return error
func (s *sBaseService) GetByIDUnscoped(out interface{}, id uint, preload ...string) error {
	return db.Scopes(s.Preload(preload...)).Unscoped().Limit(1).Where("id = ?", id).Find(out).Error
}

// UpdateByID 根据ID更新
//
//	@receiver s
//	@param model
//	@param id
//	@param updates
//	@param column
//	@return error
func (s *sBaseService) UpdateByID(model interface{}, id uint, updates interface{}, column interface{}) error {
	return db.Model(model).Select(column).Where("id = ?", id).Updates(updates).Error
}

// Paginate 分页条件
//
//	@receiver s
//	@param c
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sBaseService) Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(c.Query("limit"))
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

// Order 排序
//
//	@receiver s
//	@param order
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sBaseService) Order(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if order == "" {
			order = "id DESC"
		}
		return db.Order(order)
	}
}

// Preload 嵌套预加载
//
//	@receiver s
//	@param preload
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sBaseService) Preload(preload ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, v := range preload {
			if v != "" {
				// 使用strings.Contains()检查字符串是否包含"."符号
				if strings.Contains(v, ".") {
					preloads := strings.Split(v, ".")
					switch preloads[1] {
					case "Unscoped":
						db.Preload(preloads[0], func(db *gorm.DB) *gorm.DB {
							return db.Unscoped()
						})
					}
				} else {
					db.Preload(v)
				}
			}
		}
		return db
	}
}

// Joins join 关联
//
//	@receiver s
//	@param joins
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sBaseService) Joins(joins ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, v := range joins {
			if v != "" {
				db.Joins(v)
			}
		}
		return db
	}
}
