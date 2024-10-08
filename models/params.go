/*
*

	@author:ExiaYe
	@data:2024/4/15
	@note:

*
*/
package models

/**
 * @Author ExiaYe
 * @Description //TODO magic string
 * @Date 21:56 2024/4/15
 **/
const (
	OrderTime  = "time"
	OrderScore = "score"
)

/**
 * @Author ExiaYe
 * @Description //TODO 获取帖子列表query string参数
 * @Date 21:51 2024/4/15
 **/
type ParamPostList struct {
	CommunityID uint64 `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page"`                   // 页码
	Size        int64  `json:"size" form:"size"`                   // 每页数量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

/**
 * @Author ExiaYe
 * @Description //TODO 按社区获取帖子列表query string参数
 * @Date 22:47 2024/4/16
 **/
//type ParamCommunityPostList struct {
//	* ParamPostList
//	CommunityID uint64 `json:"community_id" form:"community_id"`
//}
