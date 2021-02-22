package service

import (
	"errors"
	"sort"
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

	go models.AddSysMessage(1, i, uint(userId), uint(userId))
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
	go models.AddSysMessage(1, *p.Status, uint(userId), p.Id)
	return
}

// 获取三月圈基本信息
func (e *Marchsoft) GetMarchMsg(userName string, id int) (marchMsg *bo.MarchSoftInfo, err error) {
	march := new(models.MarchSoftInfo)
	user := new(models.User)

	// 三月圈基本信息
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

func (e *Marchsoft) InsertMarchSoft(marchMsg *dto.InsertMarchSoft, userId uint) (err error) {
	m := new(models.MarchSoft)

	m.Content = marchMsg.Content
	m.Image = marchMsg.Image
	m.CreateBy = userId
	m.UpdateBy = userId
	err = m.InsertMarchSoft(marchMsg.ImageList)

	go models.AddMessage(1, 0, uint(m.ID), userId, "")
	return
}

func (e *Marchsoft) DeleteMarchSoft(id, userId int) (err error) {
	m := new(models.MarchSoft)

	m.IsDeleted = 1
	m.ID = id
	m.UpdateBy = uint(userId)
	err = m.DeleteMarchSoft()
	return
}

func (e *Marchsoft) SelectMarchList(paging dto.Paging, userId uint) (marchList *[]bo.March, err error) {
	// 声明所需变量
	m := new(models.MarchSoft)
	marchList = new([]bo.March)
	var goMarch bo.GoMarchMsg
	var keys []int

	// 取出要返回文章信息
	marchArray, err := m.SelectMarchSoftList(paging)
	if err != nil {
		zap.L().Error("SelectArticleListIndex Select ArticleList filed", zap.Error(err))
		return
	}

	// 组装进map
	marchMapList := make(map[int]*bo.March, len(*marchArray))
	for _, i := range *marchArray {
		var marchBo bo.March
		marchBo = i
		marchMapList[int(i.CreateTime)] = &marchBo
		keys = append(keys, int(i.CreateTime))
	}

	// 数据拼接
	var wg sync.WaitGroup
	marchCh := make(chan *bo.GoMarchMsg, len(*marchArray))
	for _, v := range *marchArray {
		wg.Add(1)
		goMarch.MarchId = uint(v.ID)
		goMarch.MarchUserId = v.CreateBy
		goMarch.UserId = userId
		goMarch.CreateTime = v.CreateTime
		go goMarchMsg(&marchCh, &wg, goMarch)
	}
	wg.Wait()
	close(marchCh)

	//articleMap 排序 遍历
	for i := range marchCh {
		goMarch = *i
		marchMapList[int(goMarch.CreateTime)].MarchTotal = goMarch.MarchTotal
	}
	// 排序
	sort.Ints(keys)
	// 遍历
	for _, key := range keys {
		*marchList = append(*marchList, *marchMapList[key])
	}
	return
}

// 用户三月圈
func (e *Marchsoft) SelectMarchListById(paging dto.SelectMarchListById, userId uint) (marchList *[]bo.MarchByUserId, err error) {
	// 声明所需变量
	m := new(models.MarchSoft)
	marchList = new([]bo.MarchByUserId)
	var goArticle bo.GoMarchMsg
	var keys []int

	// 取出要返回文章信息
	marchArray, err := m.SelectMarchSoftListByUserId(paging)
	if err != nil {
		zap.L().Error("SelectMarchListById Select ArticleList filed", zap.Error(err))
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
		marchMapList[int(i.CreateTime)] = marchBo
		keys = append(keys, int(i.CreateTime))
	}

	// 数据拼接
	var wg sync.WaitGroup
	marchCh := make(chan *bo.GoMarchMsg, len(*marchArray))
	for _, v := range *marchArray {
		wg.Add(1)
		goArticle.MarchId = uint(v.ID)
		goArticle.MarchUserId = v.CreateBy
		goArticle.UserId = userId
		goArticle.CreateTime = v.CreateTime
		go goMarchMsg(&marchCh, &wg, goArticle)
	}
	wg.Wait()
	close(marchCh)

	//articleMap 排序 遍历
	for i := range marchCh {
		goArticle = *i
		marchMapList[int(goArticle.CreateTime)].MarchTotal = goArticle.MarchTotal
	}
	// 排序
	sort.Ints(keys)
	// 遍历
	for _, key := range keys {
		*marchList = append(*marchList, *marchMapList[key])
	}
	return
}

func goMarchMsg(marchCh *chan *bo.GoMarchMsg, wg *sync.WaitGroup, marchMsg bo.GoMarchMsg) {
	marchComment := new(models.MarchsoftComment)
	marchFavour := new(models.MarchSoftFavour)
	marchImage := new(models.MarchSoftImage)
	//follow := new(models.Follow)
	marchFavourTotal := new([]bo.MarchSoftFavourTotal)
	var err error

	// 获取数据
	marchComment.MarchsoftId = marchMsg.MarchId
	marchMsg.CommentTotal, err = marchComment.GetMarchSoftCommentCount()
	if err != nil {
		zap.L().Error("goMarchMsg GetMarchSoftCommentCount failed", zap.Error(err))
	}
	marchFavour.MarchsoftId = marchMsg.MarchId
	marchFavourTotal, err = marchFavour.MarchFavourCountByType()
	if err != nil {
		zap.L().Error("goMarchMsg MarchFavourCountByType failed", zap.Error(err))
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
	marchMsg.Type, err = marchFavour.GetMarchFavourType()
	if err != nil {
		zap.L().Error("goMarchMsg GetMarchFavourType failed", zap.Error(err))
	}
	marchImage.MarchsoftId = marchMsg.MarchId
	marchMsg.ImageList, err = marchImage.GetImageListById()
	if err != nil {
		zap.L().Error("goMarchMsg GetImageListById failed", zap.Error(err))
	}
	//// 是否关注
	//if marchMsg.UserId == marchMsg.MarchUserId {
	//	marchMsg.IsFollow = 0
	//} else {
	//	marchMsg.IsFollow, err = follow.IsFollow(int(marchMsg.UserId), int(marchMsg.MarchUserId))
	//	if err != nil {
	//		zap.L().Error("goArticleMsg IsFollow failed", zap.Error(err))
	//		err = nil
	//	}
	//}
	// 管道放数据
	*marchCh <- &marchMsg
	wg.Done()
	return
}
