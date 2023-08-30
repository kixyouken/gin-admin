package service

import (
	"gin-admin/database"
	"strconv"

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

// GetPage 分页查询
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