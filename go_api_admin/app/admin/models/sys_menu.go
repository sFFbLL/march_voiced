package models

import (
	"encoding/json"

	"project/app/admin/models/bo"
	"project/app/admin/models/cache"
	"project/app/admin/models/dto"
	"project/common/global"
	"project/utils"
	"strconv"
)

const ForeNeed string = "menu::userNeed:"

type SysMenu struct {
	*BaseModel
	SubCount   int    `json:"sub_count"`  //子菜单数目
	Pid        int    `json:"pid"`        //上级菜单ID
	Type       int    `json:"type"`       //菜单类型
	MenuSort   int    `json:"menu_sort"`  //排序
	CreateBy   int    `json:"create_by"`  //
	UpdateBy   int    `json:"update_by"`  //
	Address    string `json:"address"`    //请求地址
	Action     string `json:"action"`     //请求方式
	Icon       string `json:"icon"`       //图标
	Path       string `json:"path"`       //链接地址
	Title      string `json:"title"`      //菜单标题
	Name       string `json:"name"`       //组件名称
	Component  string `json:"component"`  //组件
	Permission string `json:"permission"` //权限
	IFrame     []byte `json:"i_frame"`    //是否外链
	Cache      []byte `json:"cache"`      //缓存
	Hidden     []byte `json:"hidden"`     //隐藏
}

type ChildMenu struct {
	Id  int `json:"id"`
	Pid int `json:"pid"`
}

type RedisForeNeedMenu struct {
	SelectForeNeedMenuList []*bo.SelectForeNeedMenuBo
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}

func (m *SysMenu) InsertMenu() error {
	tx := global.Eloquent.Begin()
	err := tx.Create(&m).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//清除缓存
	var keys []string
	keys, err = global.Rdb.Keys("menu::user:*").Result()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = cache.DeleteAllMenuIdCache()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = cache.DeleteAllUserMenuCache(keys)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (m *SysMenu) SelectMenu(p *dto.SelectMenuDto) (data []*SysMenu, err error) {
	//排序条件
	var orderJson []bo.Order
	orderJson, err = utils.OrderJson(p.Orders)
	orderRule := utils.GetOrderRule(orderJson)
	//模糊条件
	blurry := "%" + p.Blurry + "%"
	var total int64
	//查询所有菜单
	allMenu := make([]*SysMenu, 0)
	table := global.Eloquent.Table("sys_menu").Where("is_deleted=?", []byte{0})
	table = table.Where("pid=? AND title LIKE ?", p.Pid, blurry)
	err = table.Find(&allMenu).Error
	if err != nil {
		return nil, err
	}
	if p.EndTime != 0 && p.StartTime != 0 {
		table = table.Where("create_time > ? AND create_time < ?", p.StartTime, p.EndTime)
	}
	table = table.Limit(p.Size).Offset(p.Current - 1*p.Size).Count(&total)
	if orderRule != "" {
		table = table.Order(orderRule)
	}
	err = table.Find(&data).Error
	//返回数据
	return data, err
}

//删除菜单
func (m *SysMenu) DeleteMenu(ids []int) (err error) {
	for _, v := range ids {
		err = utils.DeleteChild(m.TableName(), v)
		if err != nil {
			return err
		}
	}
	//删除bo缓存
	if err := cache.DeleteMenuByIdCache(ids); err != nil {
		return err
	}
	//删除userMenu 缓存
	if err := cache.DelAllUserMenuCache(); err != nil {
		return err
	}
	return nil
}

//更新菜单
func (m *SysMenu) UpdateMenu(p *dto.UpdateMenuDto, userId int) (err error) {
	tx := global.Eloquent.Begin()
	err = tx.Table("sys_menu").Where("id=?", p.ID).Updates(map[string]interface{}{
		"pid":        p.Pid,
		"sub_count":  p.SubCount,
		"type":       p.Type,
		"title":      p.Title,
		"name":       p.Name,
		"component":  p.Component,
		"menu_sort":  p.MenuSort,
		"icon":       p.Icon,
		"permission": p.Permission,
		"path":       p.Path,
		"update_by":  userId,
		"i_frame":    utils.BoolIntoByte(p.Iframe),
		"cache":      utils.BoolIntoByte(p.Cache),
		"hidden":     utils.BoolIntoByte(p.Iframe),
	}).Error
	if err != nil {
		tx.Rollback()
	}
	var keys []string
	keys, err = global.Rdb.Keys("menu::user:*").Result()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = cache.DeleteMenuByIdCache([]int{p.ID})
	if err != nil {
		tx.Rollback()
		return err
	}
	err = cache.DeleteAllUserMenuCache(keys)
	if err != nil {
		tx.Rollback()
		return err
	}
	//删除userMenu 缓存
	if err := cache.DelAllUserMenuCache(); err != nil {
		return err
	}
	return tx.Commit().Error
}

//查找前端所需菜单
func (m *SysMenu) SelectForeNeedMenu(user *ModelUserMessage) (data []*bo.SelectForeNeedMenuBo, err error) {
	//检查缓存有无,有的话从缓存中读取
	forNeedKey := ForeNeed + strconv.Itoa(user.UserId)
	if global.Rdb.Exists(forNeedKey).Val() == 1 {
		var val []byte
		val, err = global.Rdb.Get(forNeedKey).Bytes()
		redisForeNeedMenu := new(RedisForeNeedMenu)
		err = json.Unmarshal(val, redisForeNeedMenu)
		if err != nil {
			return nil, err
		}
		data = redisForeNeedMenu.SelectForeNeedMenuList
		if data != nil && len(data) != 0 {
			return data, nil
		}
	}
	//查找角色Id
	parentIds := make([]int, 0)
	err = global.Eloquent.Table("sys_roles_menus").Select("menu_id").Where("role_id=?", 1).Joins("left join sys_menu "+
		"on sys_menu.id=sys_roles_menus.menu_id").Where("sys_menu.type=? AND sys_menu.is_deleted=?", 0, []byte{0}).Find(&parentIds).Error
	if err != nil {
		return
	}
	var results []*bo.SelectForeNeedMenuBo
	for _, pid := range parentIds {
		parentMenu := new(SysMenu)
		//查找父亲信息
		childMenus := make([]*SysMenu, 0)
		childS := make([]*bo.Children, 0)
		err = global.Eloquent.Table("sys_menu").Where("id=? AND is_deleted=?", pid, []byte{0}).First(parentMenu).Error
		if err != nil {
			return
		}
		result := &bo.SelectForeNeedMenuBo{
			Hidden:    utils.ByteIntoBool(parentMenu.Hidden),
			Component: parentMenu.Component,
			Name:      parentMenu.Name,
			Path:      "/" + parentMenu.Path,
			Meta: &bo.Meta{
				Icon:    parentMenu.Icon,
				NoCache: !utils.ByteIntoBool(parentMenu.Cache),
				Title:   parentMenu.Title,
			},
		}
		//查找多个父级菜单的子信息
		err = global.Eloquent.Table("sys_menu").Where("pid=? AND is_deleted=?", pid, []byte{0}).Find(&childMenus).Error
		if err != nil {
			return
		}
		result.AlwaysShow = false
		result.Redirect = "index"
		if len(childMenus) != 0 {
			result.AlwaysShow = true
			result.Redirect = "noredirect"
		}
		//查询下一级
		for _, menu := range childMenus {
			child := &bo.Children{
				Component: menu.Component,
				Hidden:    utils.ByteIntoBool(menu.Hidden),
				Name:      menu.Name,
				Path:      menu.Path,
				Meta: &bo.Meta{
					Icon:    menu.Icon,
					NoCache: !utils.ByteIntoBool(menu.Cache),
					Title:   menu.Title,
				},
			}
			childS = append(childS, child)
		}
		result.Children = childS
		results = append(results, result)
	}
	//序列化结果
	var foreNeedMenu []byte
	redisForeNeedMenu := new(RedisForeNeedMenu)
	redisForeNeedMenu.SelectForeNeedMenuList = results
	foreNeedMenu, err = json.Marshal(redisForeNeedMenu)
	if err != nil {
		return nil, err
	}
	//添加缓存
	err = global.Rdb.Set(forNeedKey, foreNeedMenu, 0).Err()
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (m *SysMenu) ReturnToAllMenus(pid int) (data []*SysMenu, err error) {
	err = global.Eloquent.
		Table("sys_menu").
		Where("is_deleted=? AND pid = ?", []byte{0}, pid).
		Find(&data).
		Error
	return data, nil
}

// 多条件查询
func (m *SysMenu) DownloadMenu(p dto.DownloadMenuDto, orderData []bo.Order) (sysMenu []SysMenu, err error) {
	var order string
	for key, value := range orderData {
		order += value.Column + " "
		if value.Asc == "true" {
			if key == len(orderData)-1 {
				order += "asc "
			} else {
				order += "asc, "
			}
		} else {
			if key == len(orderData)-1 {
				order += "desc "
			} else {
				order += "desc, "
			}
		}
	}
	// 查询pid
	if err := global.Eloquent.
		Where("pid = ? or is_deleted = ?", m.Pid, []byte{0}).Order(order).
		Find(&sysMenu).Error; err != nil {
		return nil, err
	}
	return sysMenu, nil
}

func (m *SysMenu) SuperiorMenu(p *dto.DataMenuDto) (data, child []*SysMenu, id int, err error) {
	var dataMeau []*SysMenu
	var children []*SysMenu
	var superID int
	//根据id查出本行信息pid
	if err := global.Eloquent.Table("sys_menu").Where("is_deleted=? AND id=?", []byte{0}, p.Ids[0]).
		First(&dataMeau).Error; err != nil {
		return nil, nil, 0, err
	}
	//根据本行的pid查询出所有信息
	if err := global.Eloquent.Table("sys_menu").Where("is_deleted=? AND pid=?", []byte{0}, dataMeau[0].Pid).
		Find(&dataMeau).Error; err != nil {
		return nil, nil, 0, err
	}
	//判断pid是否为0，如果为0，则此不用查上级，直接返回数据
	if dataMeau[0].Pid == 0 {
		return dataMeau, nil, 0, nil
	} else {
		//如果不为零，则此级为children.根据pid查询上级
		if err := global.Eloquent.Table("sys_menu").Where("is_deleted=? AND id=?", []byte{0}, dataMeau[0].Pid).
			First(&children).Error; err != nil {
			return nil, nil, 0, err
		}
		superID = children[0].ID
		if err := global.Eloquent.Table("sys_menu").Where("is_deleted=? AND pid=? ", []byte{0}, children[0].Pid).
			Find(&children).Error; err != nil {
			return nil, nil, 0, err
		}
		return children, dataMeau, superID, err
	}
}

func (m *SysMenu) ChildMenu(p int) (data []int, err error) {
	data = append(data, p)
	var childMeau []*ChildMenu
	for i := 0; i < len(data); i++ {
		if err := global.Eloquent.Table("sys_menu").Where("is_deleted=? AND pid=?", []byte{0}, data[i]).
			Find(&childMeau).Error; err != nil {
			return nil, err
		}
		for _, j := range childMeau {
			data = append(data, j.Id)
		}
	}
	return data, nil
}