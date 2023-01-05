package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/kubemanage/dao/model"
	"github.com/noovertime7/kubemanage/pkg"
	"gorm.io/gorm"
)

type CMDBHostCreateInput struct {
	InstanceID      int64  `json:"instanceID" `
	Address         string `json:"address" validate:"required" `
	Port            string `json:"port" validate:"required"`
	HostUserName    string `json:"hostUserName" `
	HostPassword    string `json:"hostPassword" `
	PrivateKey      string `json:"privateKey" `
	SecretID        uint   `json:"secretID"`
	CMDBHostGroupID uint   `json:"cmdbHostGroupID" validate:"required" `
}

func (p *CMDBHostCreateInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

type PageCMDBHostOut struct {
	Total int64            `json:"total"`
	List  []model.CMDBHost `json:"list"`
}

type PageListCMDBHostInput struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

func (p *PageListCMDBHostInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

func (p *PageListCMDBHostInput) GetPage() int {
	return p.Page
}

func (p *PageListCMDBHostInput) GetPageSize() int {
	return p.PageSize
}

func (p *PageListCMDBHostInput) IsFitter() bool {
	return p.Keyword != ""
}

func (p *PageListCMDBHostInput) Do(tx *gorm.DB) {
	tx.Where("name like ? or address like ?", "%"+p.Keyword+"%", "%"+p.Keyword+"%")
}