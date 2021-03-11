package service

import (
	"encoding/json"
	"fmt"
	"time"

	model "project/app/admin/models"
	Bo "project/app/admin/models/bo"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/cache"
	"project/common/global"
	"project/pkg/jwt"
	"project/utils"
	"project/utils/config"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
	"go.uber.org/zap"
)

type Login struct {
}

func (login *Login) LoginService(c *gin.Context, info oauth.UserInfo) (data bo.MarchLoginData, err error) {
	user := new(models.Model)
	status, sysUser, err := user.LoginDao(info)
	if err != nil {
		return
	}
	var token string
	if status == 2 {
		keys := new([]string)
		*keys = append(*keys, cache.KeyUserJob, cache.KeyUserRole, cache.KeyUserMenu, cache.KeyUserDept, cache.KeyUserDataScope)
		cacheMap := cache.GetUserCache(keys, sysUser.ID)
		cacheJob, jobErr := cacheMap[cache.KeyUserJob].Result()
		cacheRole, rolesErr := cacheMap[cache.KeyUserRole].Result()
		cacheMenu, menuErr := cacheMap[cache.KeyUserMenu].Result()
		cacheDept, deptErr := cacheMap[cache.KeyUserDept].Result()
		cacheDataScopes, dataScopesErr := cacheMap[cache.KeyUserDataScope].Result()
		jobs := new([]model.SysJob)
		if err = GetUserJobData(cacheJob, jobErr, jobs, sysUser.ID); err != nil {
			return
		}
		roles := new([]model.SysRole)
		if err = GetUserRoleData(cacheRole, rolesErr, roles, sysUser.ID); err != nil {
			return
		}
		menuPermission := new([]string)
		if err = GetUserMenuData(cacheMenu, menuErr, sysUser.ID, menuPermission, roles); err != nil {
			return
		}

		dept := new(model.SysDept)
		if err = GetUserDeptData(cacheDept, deptErr, dept, sysUser.ID); err != nil {
			return
		}

		dataScopes := new([]int)
		if err = GetUserDataScopes(cacheDataScopes, dataScopesErr, dataScopes, sysUser.ID, dept.ID, roles); err != nil {
			return
		}

		token, err = jwt.GenToken(sysUser.ID, sysUser.Username)
		if err != nil {
			return
		}

		var jobMessage []*Bo.Job
		for _, job := range *jobs {
			jobMessage = append(jobMessage, &Bo.Job{
				ID:   job.ID,
				Name: job.Name,
			})
		}
		var roleMessage []*Bo.Role
		for _, role := range *roles {
			roleMessage = append(roleMessage, &Bo.Role{
				ID:        role.ID,
				Level:     role.Level,
				Name:      role.Name,
				DataScope: role.DataScope,
			})
		}
		deptMessage := new(Bo.DeptCommon)
		deptMessage.ID = dept.ID
		deptMessage.Name = dept.Name

		data := new(Bo.LoginData)
		loginUser := new(Bo.LoginUser)
		recordUser := new(Bo.RecordUser)
		recordUserHalf := new(Bo.RecordUserHalf)
		roleDeptJobBool := new(Bo.RoleDeptJobBool)
		recordUserHalf.Id = sysUser.ID
		recordUserHalf.DeptId = sysUser.DeptId
		recordUserHalf.CreateBy = sysUser.CreateBy
		recordUserHalf.UpdateBy = sysUser.UpdateBy
		recordUserHalf.PwdResetTime = sysUser.PwdResetTime
		recordUserHalf.CreateTime = sysUser.CreateTime
		recordUserHalf.UpdateTime = sysUser.UpdateTime
		recordUserHalf.AvatarName = sysUser.Avatar
		recordUserHalf.AvatarPath = sysUser.AvatarPath
		recordUserHalf.Email = sysUser.Email
		recordUserHalf.NickName = sysUser.NickName
		recordUserHalf.Phone = sysUser.Phone
		recordUserHalf.Username = sysUser.Username
		recordUser.RecordUserHalf = recordUserHalf

		roleDeptJobBool.Enabled = utils.ByteIntoBool(sysUser.Enabled)
		roleDeptJobBool.Gender = utils.ByteIntoBool(sysUser.Gender)
		roleDeptJobBool.Jobs = jobMessage
		roleDeptJobBool.Role = roleMessage
		roleDeptJobBool.Dept = deptMessage
		recordUser.RoleDeptJobBool = roleDeptJobBool

		loginUser.User = recordUser
		loginUser.DataScopes = *dataScopes
		loginUser.Roles = *menuPermission

		data.User = loginUser
		data.Token = "Bearer " + token
		token = "Bearer " + token
		err = RedisUserMessage(c, data, token)
	}
	data = bo.MarchLoginData{
		OpenId: info.OpenID,
		Token:  token,
		Status: status,
	}
	return data, err
}

func (login *Login) CreatSysUserService(c *gin.Context, p dto.InsertUserDto) (datas bo.LoginInfoData, err error) {
	user := new(models.Model)
	sysUser, err := user.CreatSysUserService(p)
	if sysUser.ID > 0 {
		keys := new([]string)
		*keys = append(*keys, cache.KeyUserJob, cache.KeyUserRole, cache.KeyUserMenu, cache.KeyUserDept, cache.KeyUserDataScope)
		cacheMap := cache.GetUserCache(keys, sysUser.ID)
		cacheJob, jobErr := cacheMap[cache.KeyUserJob].Result()
		cacheRole, rolesErr := cacheMap[cache.KeyUserRole].Result()
		cacheMenu, menuErr := cacheMap[cache.KeyUserMenu].Result()
		cacheDept, deptErr := cacheMap[cache.KeyUserDept].Result()
		cacheDataScopes, dataScopesErr := cacheMap[cache.KeyUserDataScope].Result()

		jobs := new([]model.SysJob)
		if err = GetUserJobData(cacheJob, jobErr, jobs, sysUser.ID); err != nil {
			return
		}

		roles := new([]model.SysRole)
		if err = GetUserRoleData(cacheRole, rolesErr, roles, sysUser.ID); err != nil {
			return
		}

		menuPermission := new([]string)
		if err = GetUserMenuData(cacheMenu, menuErr, sysUser.ID, menuPermission, roles); err != nil {
			return
		}

		dept := new(model.SysDept)
		if err = GetUserDeptData(cacheDept, deptErr, dept, sysUser.ID); err != nil {
			return
		}

		dataScopes := new([]int)
		if err = GetUserDataScopes(cacheDataScopes, dataScopesErr, dataScopes, sysUser.ID, dept.ID, roles); err != nil {
			return
		}

		datas.Token, err = jwt.GenToken(sysUser.ID, sysUser.Username)
		if err != nil {
			return
		}

		var jobMessage []*Bo.Job
		for _, job := range *jobs {
			jobMessage = append(jobMessage, &Bo.Job{
				ID:   job.ID,
				Name: job.Name,
			})
		}
		var roleMessage []*Bo.Role
		for _, role := range *roles {
			roleMessage = append(roleMessage, &Bo.Role{
				ID:        role.ID,
				Level:     role.Level,
				Name:      role.Name,
				DataScope: role.DataScope,
			})
		}
		deptMessage := new(Bo.DeptCommon)
		deptMessage.ID = dept.ID
		deptMessage.Name = dept.Name

		data := new(Bo.LoginData)
		loginUser := new(Bo.LoginUser)
		recordUser := new(Bo.RecordUser)
		recordUserHalf := new(Bo.RecordUserHalf)
		roleDeptJobBool := new(Bo.RoleDeptJobBool)
		recordUserHalf.Id = sysUser.ID
		recordUserHalf.DeptId = sysUser.DeptId
		recordUserHalf.CreateBy = sysUser.CreateBy
		recordUserHalf.UpdateBy = sysUser.UpdateBy
		recordUserHalf.PwdResetTime = sysUser.PwdResetTime
		recordUserHalf.CreateTime = sysUser.CreateTime
		recordUserHalf.UpdateTime = sysUser.UpdateTime
		recordUserHalf.AvatarName = sysUser.Avatar
		recordUserHalf.AvatarPath = sysUser.AvatarPath
		recordUserHalf.Email = sysUser.Email
		recordUserHalf.NickName = sysUser.NickName
		recordUserHalf.Phone = sysUser.Phone
		recordUserHalf.Username = sysUser.Username
		recordUser.RecordUserHalf = recordUserHalf

		roleDeptJobBool.Enabled = utils.ByteIntoBool(sysUser.Enabled)
		roleDeptJobBool.Gender = utils.ByteIntoBool(sysUser.Gender)
		roleDeptJobBool.Jobs = jobMessage
		roleDeptJobBool.Role = roleMessage
		roleDeptJobBool.Dept = deptMessage
		recordUser.RoleDeptJobBool = roleDeptJobBool

		loginUser.User = recordUser
		loginUser.DataScopes = *dataScopes
		loginUser.Roles = *menuPermission

		data.User = loginUser
		token := datas.Token
		data.Token = "Bearer " + datas.Token
		datas.Token = "Bearer " + datas.Token
		err = RedisUserMessage(c, data, token)
	}
	return
}

// GetUserJobData 获取用户岗位数据
func GetUserJobData(cacheJob string, jobErr error, jobs *[]model.SysJob, userId int) (err error) {
	if jobErr != nil {
		err = model.GetUserJob(jobs, userId)
		if err != nil {
			return
		}
		go cache.SetUserCache(userId, jobs, cache.KeyUserJob)

	} else {
		err = utils.JsonToStruct(cacheJob, jobs)
	}
	return
}

// GetUserRoleData 获取用户角色数据
func GetUserRoleData(cacheRole string, rolesErr error, roles *[]model.SysRole, userId int) (err error) {
	if rolesErr != nil {
		err = model.GetUserRole(roles, userId)
		if err != nil {
			return
		}
		go cache.SetUserCache(userId, roles, cache.KeyUserRole)
	} else {
		err = utils.JsonToStruct(cacheRole, roles)
	}
	return
}

// GetUserMenuData 获取用户菜单权限
func GetUserMenuData(cacheMenu string, menuErr error, userId int, menuPermission *[]string, roles *[]model.SysRole) (err error) {
	if menuErr != nil {
		a := new(model.Admin)
		if err = a.GetIsAdmin(userId); err != nil {
			return
		}

		if utils.ByteIntoInt(a.IsAdmin) == 1 {
			*menuPermission = []string{`admin`}
			go cache.SetUserCache(userId, menuPermission, cache.KeyUserMenu)
		} else {
			menus := new([]model.SysMenu)
			//if err = model.SelectUserMenuPermission(menus, roles); err != nil {
			//	return
			//}
			//for _, menu := range *menus {
			//	if menu.Permission != "" {
			//		*menuPermission = append(*menuPermission, menu.Permission)
			//	}
			//}
			go cache.SetUserCache(userId, menus, cache.KeyUserMenu)
		}
	} else {
		if cacheMenu == `["admin"]` {
			*menuPermission = []string{`admin`}
		} else {
			menus := new([]model.SysMenu)
			if err = utils.JsonToStruct(cacheMenu, menus); err != nil {
				return
			}
			for _, menu := range *menus {
				if menu.Permission != "" {
					*menuPermission = append(*menuPermission, menu.Permission)
				}
			}
		}
	}
	return
}

// GetUserDeptData 获取用户部门数据
func GetUserDeptData(cacheDept string, deptErr error, dept *model.SysDept, userId int) (err error) {
	if deptErr != nil {
		err = model.SelectUserDept(dept, userId)
		if err != nil {
			return
		}
		go cache.SetUserCache(userId, dept, cache.KeyUserDept)
	} else {
		err = utils.JsonToStruct(cacheDept, dept)
	}
	return
}

// GetUserDataScopes 获取用户数据权限
func GetUserDataScopes(cacheDataScopes string, dataScopesErr error, dataScopes *[]int, userId int, deptId int, roles *[]model.SysRole) (err error) {
	if dataScopesErr != nil {
		var dataScopesRoleIds []int
		var allScopes bool
		for _, role := range *roles {
			switch role.DataScope {
			case `全部`:
				allScopes = true
				*dataScopes = []int{}
				break
			case `本级`:
				*dataScopes = append(*dataScopes, deptId)
			default:
				dataScopesRoleIds = append(dataScopesRoleIds, role.ID)
			}
		}

		if !allScopes {
			deptIds, err := model.SelectUserDeptIdByRoleId(dataScopesRoleIds)
			if err != nil {
				return err
			}
			*dataScopes = append(*dataScopes, deptIds...)
		}
		go cache.SetUserCache(userId, dataScopes, cache.KeyUserDataScope)
	} else {
		err = utils.JsonToStruct(cacheDataScopes, dataScopes)
	}
	return err
}

func RedisUserMessage(c *gin.Context, l *Bo.LoginData, token string) (err error) {
	online := new(model.OnlineUser)
	ua := user_agent.New(c.Request.UserAgent())
	browserName, browserVersion := ua.Browser()
	online.LoginTime = utils.NowUnix()
	online.LoginLocation = utils.GetLocation(c.ClientIP())
	online.Browser = browserName + " " + browserVersion
	online.Dept = l.User.User.Dept.Name
	online.Ip = c.ClientIP()
	online.Nickname = l.User.User.NickName
	online.Username = l.User.User.Username
	online.Token = token
	userOnline, err := json.Marshal(online)
	if err != nil {
		zap.L().Error("RedisUserOnline Marshal failed", zap.Error(err))
		return
	}
	//添加缓存
	if err = global.Rdb.Set(fmt.Sprintf("%s%s%s", config.JwtConfig.RedisHeader, "-", token), userOnline, time.Second*time.Duration(config.JwtConfig.Timeout)).Err(); err != nil {
		zap.L().Error("用户缓存错误", zap.Error(err))
		return
	}
	return
}

func (login *Login) SearchUsername(username string) error {
	user := new(models.Model)
	err := user.SearchUsername(username)
	return err
}

func (login *Login) SearchUserInfo(id int) (bo.UserInfo, error) {
	user := new(models.Model)
	var userinfo bo.UserInfo
	var sex int
	info, err := user.SearchUserInfo(id)
	if info.Sex == "男" {
		sex = 0
	} else if info.Sex == "女" {
		sex = 1
	}
	userinfo = bo.UserInfo{
		NickName:   info.NickName,
		Sex:        sex,
		Signature:  info.Signature,
		AvatarPath: info.AvatarPath,
	}
	//strconv.Atoi(info.Gender),
	return userinfo, err
}

func (login *Login) ModInformation(userinfo dto.ModInformationDto, id string) error {
	user := new(models.Model)
	err := user.ModInformation(userinfo, id)
	return err
}
