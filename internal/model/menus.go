package model

import "github.com/liuzhengtao/auth-common-backend/internal/model/entity"

type RouteBO struct {
	ID         int64    `json:"id"`
	ParentId   int64    `json:"parentId"`
	Name       string   `json:"name"`
	Type       int      `json:"type"`
	Path       string   `json:"path"`
	Component  string   `json:"component"`
	Perm       string   `json:"perm"`
	Visible    int      `json:"visible"`
	Sort       int      `json:"sort"`
	Icon       string   `json:"icon"`
	Redirect   string   `json:"redirect"`
	Roles      []string `json:"roles"`
	AlwaysShow int      `json:"alwaysShow"`
	KeepAlive  int      `json:"keepAlive"`
}
type MenuFormOutput struct {
	*entity.SysMenu
	Type string `json:"type"`
}
