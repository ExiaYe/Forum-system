/*
*

	@author:ExiaYe
	@data:2024/4/12
	@note:

*
*/
package logic

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
)

/**
 * @Author ExiaYe
 * @Description //TODO 查询分类社区列表
 * @Date 16:42 2024/4/12
 **/
func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

/**
 * @Author ExiaYe
 * @Description //TODO 根据ID查询分类社区详情
 * @Date 17:08 2024/4/12
 **/
func GetCommunityDetailByID(id uint64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityByID(id)
}
