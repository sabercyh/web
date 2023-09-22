package preheat

import (
	"web/pkg/setting"
	"web/service/article_service"
)

func Setup() {
	(&article_service.Article{PageSize: setting.AppSetting.PageSize}).GetAll()
}
