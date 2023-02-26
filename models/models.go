package models

type GetPostListByFilterParams struct {
	Title       string `json:"title,omitempty" form:"title,omitempty" db:"title" sql:"like"`
	CommunityId string `json:"community_id,omitempty" form:"community_id,omitempty" db:"community_id"`
	AuthorId    string `json:"author_id,omitempty" form:"author_id,omitempty" db:"author_id"`
	PageSize    int    `form:"page_size" json:"page_size" binding:"required"`
	PageNumber  int    `form:"page_number" json:"page_number" binding:"required"`
	Sort        string `form:"sort" json:"sort" binding:"required"`
}
