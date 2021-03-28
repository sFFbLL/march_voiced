package service

import (
	"errors"
	"sync"

	models2 "project/app/admin/models"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"
	"project/utils"

	"go.uber.org/zap"
)

type Marchsoft struct {
}

// GetMarchApplyUser 查询申请三月圈审核中用户业务方法
func (e *Marchsoft) GetMarchApplyUser(p *dto.ApplyMarchPaginator, userId int) (applyMarchUser *bo.ApplyMarchUser, err error) {
	marchSoft := new(models.MarchSoft)
	applyMarchUser = new(bo.ApplyMarchUser)
	applyMarchUserMessage := new([]bo.ApplyMarchUserMessage)
	applyMarchUser.Records = applyMarchUserMessage
	err = marchSoft.GetMarchApplyUser(applyMarchUser, p)
	applyMarchUser.Pages = utils.PagesCount(int(applyMarchUser.Total), int(p.Size))
	applyMarchUser.Current = p.Current
	applyMarchUser.Size = p.Size
	return
}

// ApplyMarch 申请加入三月圈业务方法
func (e *Marchsoft) ApplyMarch(userId int) (err error) {
	var count int64
	table := global.Eloquent.Table("sys_user").
		Where("id=? and is_march=0 and is_deleted=0", userId)
	err = table.Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		return errors.New("用户已申请或已加入")
	}
	var i uint8 = 2
	err = table.Updates(&models2.SysUser{IsMarch: &i, MarchUpdateTime: utils.NowUnix()}).Error
	if err != nil {
		return
	}

	go models.AddSysMessage(1, i, uint(userId), uint(userId), "")
	return
}

// MarchPass （后台）三月圈审核通过或驳回业务方法
func (e *Marchsoft) MarchPass(p *dto.MarchPass, userId int) (err error) {
	var count int64
	table := global.Eloquent.Table("sys_user").
		Where("id=? and is_march=2 and is_deleted=0", p.Id)
	err = table.Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		return errors.New("该用户未申请三月圈")
	}
	err = table.Updates(&models2.SysUser{IsMarch: p.Status, MarchUpdateTime: utils.NowUnix()}).Error
	if err != nil {
		return
	}
	go models.AddSysMessage(1, *p.Status, uint(userId), p.Id, "")
	return
}

// 获取三月圈基本信息
func (e *Marchsoft) GetMarchMsg(userName string, id int) (marchMsg *bo.MarchSoftInfo, err error) {
	march := new(models.MarchSoftInfo)
	user := new(models.User)

	// 三月圈基本信息
	zap.L().Info("Call GetMarchSoftInfo GetMarchSoftInfo", zap.String("Username", userName))
	err = march.GetMarchSoftInfo()
	if err != nil {
		zap.L().Error("GetMarchSoftInfo Failed", zap.String("Username", userName), zap.Error(err))
		return
	}
	marchMsg = &bo.MarchSoftInfo{
		Title:        march.Title,
		Brief:        march.Brief,
		PersonTotal:  march.PersonTotal,
		ArticleTotal: march.ArticleTotal,
	}

	// 查是否为三月圈
	user.ID = id
	marchMsg.IsMarch, err = user.IsMarchByUserId()
	if err != nil {
		zap.L().Error("IsMarchByUserId Failed", zap.String("Username", userName), zap.Error(err))
		return
	}
	return
}

func (e *Marchsoft) InsertMarchSoft(marchMsg *dto.InsertMarchSoft, userId uint) (isMarch uint8, err error) {
	m := new(models.MarchSoft)
	u := new(models.User)

	// 查是否为三月圈
	u.ID = int(userId)
	isMarch, err = u.IsMarchByUserId()
	if err != nil || isMarch == 0 {
		zap.L().Error("InsertMarchSoft IsMarchByUserId Failed", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		return
	}

	m.Content = marchMsg.Content
	m.Image = marchMsg.Image
	m.CreateBy = userId
	m.UpdateBy = userId
	zap.L().Info("Call InsertMarchSoft InsertMarchSoft", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
	err = m.InsertMarchSoft(marchMsg.ImageList)

	zap.L().Info("Call InsertMarchSoft AddMessage", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
	go models.AddMessage(1, 0, uint(m.ID), userId, "")
	return
}

func (e *Marchsoft) DeleteMarchSoft(id, userId int) (err error) {
	m := new(models.MarchSoft)

	m.ID = id
	zap.L().Info("Call DeleteMarchSoft GetMarchById", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	err = m.GetMarchById()
	if err != nil {
		zap.L().Error("DeleteMarchSoft GetMarchById failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 判断是否有权利删除该文章
	if uint(userId) != m.CreateBy {
		err = errors.New("不能够删除他人三月圈！")
		zap.L().Error("DeleteMarchSoft check CreateBy and userId failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	m.IsDeleted = 1
	m.UpdateBy = uint(userId)
	zap.L().Info("Call DeleteMarchSoft DeleteMarchSoft", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	err = m.DeleteMarchSoft()
	return
}

func (e *Marchsoft) SelectMarchList(paging dto.Paging, userId uint) (marchList *[]bo.March, err error) {
	// 声明所需变量
	m := new(models.MarchSoft)
	marchList = new([]bo.March)
	var goMarch bo.GoMarchMsg

	// 取出要返回文章信息
	zap.L().Info("Call SelectMarchList SelectMarchSoftList", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
	marchArray, err := m.SelectMarchSoftList(paging)
	if err != nil {
		zap.L().Error("SelectArticleListIndex Select ArticleList filed", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		return
	}

	// 组装进map
	marchMapList := make(map[int]*bo.March, len(*marchArray))
	for _, i := range *marchArray {
		var marchBo bo.March
		marchBo = i
		marchMapList[i.ID] = &marchBo
	}

	// 数据拼接
	var wg sync.WaitGroup
	marchCh := make(chan *bo.GoMarchMsg, len(*marchArray))
	for _, v := range *marchArray {
		wg.Add(1)
		goMarch.MarchId = uint(v.ID)
		goMarch.MarchUserId = v.CreateBy
		goMarch.UserId = userId
		zap.L().Info("Call SelectMarchList goMarchMsg", zap.String("MarchId", utils.IntToString(v.ID)), zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		go goMarchMsg(&marchCh, &wg, goMarch)
	}
	wg.Wait()
	close(marchCh)

	//articleMap 排序 遍历
	for i := range marchCh {
		goMarch = *i
		marchMapList[int(goMarch.MarchId)].MarchTotal = goMarch.MarchTotal
	}

	for _, v := range *marchArray {
		*marchList = append(*marchList, *marchMapList[v.ID])
	}

	return
}

// 用户三月圈
func (e *Marchsoft) SelectMarchListById(paging dto.SelectMarchListById, userId uint) (marchList *[]bo.MarchByUserId, err error) {
	// 声明所需变量
	m := new(models.MarchSoft)
	marchList = new([]bo.MarchByUserId)
	var goArticle bo.GoMarchMsg

	// 取出要返回文章信息
	zap.L().Info("Call SelectMarchListById SelectMarchSoftListByUserId", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
	marchArray, err := m.SelectMarchSoftListByUserId(paging)
	if err != nil {
		zap.L().Error("SelectMarchListById Select ArticleList filed", zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		return
	}

	// 组装进map
	marchMapList := make(map[int]*bo.MarchByUserId, len(*marchArray))
	for _, i := range *marchArray {
		marchBo := &bo.MarchByUserId{
			MarchMsg: bo.MarchMsg{
				ID:         i.ID,
				Content:    i.Content,
				Image:      i.Image,
				UpdateBy:   i.UpdateBy,
				CreateBy:   i.CreateBy,
				CreateTime: i.CreateTime,
				UpdateTime: i.UpdateTime,
			},
		}
		marchMapList[i.ID] = marchBo

	}

	// 数据拼接
	var wg sync.WaitGroup
	marchCh := make(chan *bo.GoMarchMsg, len(*marchArray))
	for _, v := range *marchArray {
		wg.Add(1)
		goArticle.MarchId = uint(v.ID)
		goArticle.MarchUserId = v.CreateBy
		goArticle.UserId = userId
		zap.L().Info("Call SelectMarchListById goMarchMsg", zap.String("MarchId", utils.IntToString(v.ID)), zap.String("UserID", utils.UIntToString(userId)), zap.Error(err))
		go goMarchMsg(&marchCh, &wg, goArticle)
	}
	wg.Wait()
	close(marchCh)

	//articleMap 排序 遍历
	for i := range marchCh {
		goArticle = *i
		marchMapList[int(goArticle.MarchId)].MarchTotal = goArticle.MarchTotal
	}

	// 遍历
	for _, v := range *marchArray {
		*marchList = append(*marchList, *marchMapList[v.ID])
	}
	return
}

// 三月圈详情页
func (e *Marchsoft) MarchDetail(id int, userId int) (march *bo.March, err error) {
	//初始化
	m := new(models.MarchSoft)
	marchComment := new(models.MarchsoftComment)
	marchFavour := new(models.MarchSoftFavour)
	marchImage := new(models.MarchSoftImage)
	marchFavourTotal := new([]bo.MarchSoftFavourTotal)
	follow := new(models.Follow)
	march = new(bo.March)

	// 获取数据
	m.ID = id
	zap.L().Info("Call MarchDetail MarchDetail", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	march, err = m.MarchDetail()
	if err != nil {
		zap.L().Error("MarchDetail MarchDetail failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 文章图片列表
	marchImage.MarchsoftId = uint(id)
	zap.L().Info("Call MarchDetail GetImageListById", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	march.ImageList, err = marchImage.GetImageListById()
	if err != nil {
		zap.L().Error("MarchDetail GetImageListById failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 文章评论数
	marchComment.MarchsoftId = uint(id)
	zap.L().Info("Call MarchDetail GetMarchSoftCommentCount", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	march.CommentTotal, err = marchComment.GetMarchSoftCommentCount()
	if err != nil {
		zap.L().Error("MarchDetail GetMarchSoftCommentCount count failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 文章喜欢数
	marchFavour.MarchsoftId = uint(id)
	zap.L().Info("Call MarchDetail MarchFavourCountByType", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	marchFavourTotal, err = marchFavour.MarchFavourCountByType()
	if err != nil {
		zap.L().Error("MarchDetail MarchFavourCountByType failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}
	for _, v := range *marchFavourTotal {
		switch v.Type {
		case 1:
			march.FaceTotal = v.Total
		case 2:
			march.LikeTotal = v.Total
		case 3:
			march.FavourTotal = v.Total
		}
	}

	// 文章喜欢类型
	zap.L().Info("Call MarchDetail GetMarchFavourType", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
	march.Type, err = marchFavour.GetMarchFavourType()
	if err != nil {
		zap.L().Error("MarchDetail GetMarchFavourType failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		return
	}

	// 是否关注
	if userId == int(march.CreateBy) {
		march.IsFollow = 0
	} else {
		zap.L().Info("Call MarchDetail IsFollow", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
		march.IsFollow, err = follow.IsFollow(userId, int(march.CreateBy))
		if err != nil {
			zap.L().Error("MarchDetail IsFollow failed", zap.String("UserID", utils.IntToString(userId)), zap.Error(err))
			return
		}
	}
	return
}

func goMarchMsg(marchCh *chan *bo.GoMarchMsg, wg *sync.WaitGroup, marchMsg bo.GoMarchMsg) {
	// 异常捕获
	defer func() {
		if err := recover(); err != nil {
			zap.L().Error("Call goMarchMsg defer recover", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.String("error", string(utils.Stack())))
		}

		// 管道放数据
		*marchCh <- &marchMsg
		wg.Done()
		zap.L().Info("Call goMarchMsg Done", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)))
	}()

	marchComment := new(models.MarchsoftComment)
	marchFavour := new(models.MarchSoftFavour)
	marchImage := new(models.MarchSoftImage)
	follow := new(models.Follow)
	marchFavourTotal := new([]bo.MarchSoftFavourTotal)
	var err error

	// 获取评论数
	marchComment.MarchsoftId = marchMsg.MarchId
	zap.L().Info("Call goMarchMsg GetMarchSoftCommentCount", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.Error(err))
	marchMsg.CommentTotal, err = marchComment.GetMarchSoftCommentCount()
	if err != nil {
		zap.L().Error("goMarchMsg GetMarchSoftCommentCount failed", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.Error(err))
	}

	// 获取喜欢数
	marchFavour.MarchsoftId = marchMsg.MarchId
	zap.L().Info("Call goMarchMsg MarchFavourCountByType", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.Error(err))
	marchFavourTotal, err = marchFavour.MarchFavourCountByType()
	if err != nil {
		zap.L().Error("goMarchMsg MarchFavourCountByType failed", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.Error(err))
	}

	for _, v := range *marchFavourTotal {
		switch v.Type {
		case 1:
			marchMsg.FaceTotal = v.Total
		case 2:
			marchMsg.LikeTotal = v.Total
		case 3:
			marchMsg.FavourTotal = v.Total
		}
	}
	zap.L().Info("Call goMarchMsg GetMarchFavourType", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.Error(err))
	marchMsg.Type, err = marchFavour.GetMarchFavourType()
	if err != nil {
		zap.L().Error("goMarchMsg GetMarchFavourType failed", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.Error(err))
	}

	marchImage.MarchsoftId = marchMsg.MarchId
	zap.L().Info("Call goMarchMsg GetImageListById", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.Error(err))
	marchMsg.ImageList, err = marchImage.GetImageListById()
	if err != nil {
		zap.L().Error("goMarchMsg GetImageListById failed", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.Error(err))
	}

	// 是否关注
	if marchMsg.UserId == marchMsg.MarchUserId {
		marchMsg.IsFollow = 0
	} else {
		zap.L().Info("Call goMarchMsg IsFollow", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.Error(err))
		marchMsg.IsFollow, err = follow.IsFollow(int(marchMsg.UserId), int(marchMsg.MarchUserId))
		if err != nil {
			zap.L().Error("goMarchMsg IsFollow failed", zap.String("MarchId", utils.UIntToString(marchMsg.MarchId)), zap.String("UserID", utils.UIntToString(marchMsg.UserId)), zap.Error(err))
		}
	}

	return
}
