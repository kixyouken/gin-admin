package service

import (
	"gin-admin/database"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
)

type sMenuService struct{}

var MenuService = sMenuService{}

var (
	modelList []model.MenuModel
	// modelInfo model.MenuModel
	// paramList []param.MenuInfo
	// paramInfo param.MenuInfo
	err error
)

// MenuInfo 详情
//
//	@receiver s
//	@param in
//	@return out
func (s *sMenuService) MenuInfo(in model.MenuModel) (out param.MenuInfo) {
	out.ID = in.ID
	out.ParentID = in.ParentID
	out.Order = in.Order
	out.Title = in.Title
	out.Icon = in.Icon
	out.Uri = in.Uri
	out.Extension = in.Extension
	return out
}

// MenuList 列表
//
//	@receiver s
//	@param in
//	@return out
func (s *sMenuService) MenuList(in []model.MenuModel) (out []param.MenuInfo) {
	for _, v := range in {
		info := s.MenuInfo(v)
		out = append(out, info)
	}
	return out
}

// GetAllMenu 获取菜单
//
//	@receiver s
//	@param c
//	@return []model.MenuModel
//	@return error
func (s *sMenuService) GetAllMenu(c *gin.Context) ([]model.MenuModel, error) {
	db := database.InitMysql()
	err = db.Unscoped().Order("`order`").Where("`show` = ?", 1).Find(&modelList).Error
	if err != nil {
		return nil, err
	}
	return modelList, nil
}

// BuildTree 处理无限级菜单
//
//	@receiver s
//	@param menuItems
//	@param parentID
//	@return []*param.MenuInfo
func (s *sMenuService) BuildTree(menuItems []param.MenuInfo, parentID uint) []*param.MenuInfo {
	var tree []*param.MenuInfo

	for i := range menuItems {
		if menuItems[i].ParentID == parentID {
			children := s.BuildTree(menuItems, menuItems[i].ID)
			if len(children) > 0 {
				menuItems[i].Children = children
			}
			tree = append(tree, &menuItems[i])
		}
	}

	return tree
}
