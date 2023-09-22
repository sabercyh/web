package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"web/pkg/setting"

	"web/pkg/app"
	"web/pkg/e"
	"web/service/article_service"
	"web/service/tag_service"
)

// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [get]
func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}

// @Summary Get multiple articles
// @Produce  json
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	//后用
	// state := -1
	// if arg := c.PostForm("state"); arg != "" {
	// 	state = com.StrTo(arg).MustInt()
	// 	valid.Range(state, 0, 1, "state")
	// }

	// tagId := -1
	// if arg := c.PostForm("tag_id"); arg != "" {
	// 	tagId = com.StrTo(arg).MustInt()
	// 	valid.Min(tagId, 1, "tag_id")
	// }

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{
		// TagID:    tagId,
		// State:    state,
		PageNum:  0,
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	articles, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type GetArticlesByYearForm struct {
	Theyear string `form:"theyear" valid:"Required;MaxSize(10)"`
}

// @Summary Get multiple articles
// @Produce  json
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles [get]
func GetArticlesByYear(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	var form GetArticlesByYearForm
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	//后用
	// state := -1
	// if arg := c.PostForm("state"); arg != "" {
	// 	state = com.StrTo(arg).MustInt()
	// 	valid.Range(state, 0, 1, "state")
	// }

	// tagId := -1
	// if arg := c.PostForm("tag_id"); arg != "" {
	// 	tagId = com.StrTo(arg).MustInt()
	// 	valid.Min(tagId, 1, "tag_id")
	// }

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{
		// TagID:    tagId,
		// State:    state,
		Theyear:  form.Theyear,
		PageNum:  0,
		PageSize: setting.AppSetting.PageSize,
	}

	articles, err := articleService.GetAllByYear()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = len(articles)

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddArticleForm struct {
	TagID int    `form:"tag_id" valid:"Required;Min(1)"`
	Title string `form:"title" valid:"Required;MaxSize(100)"`

	Journal   string `form:"journal" valid:"Required;MaxSize(100)"`
	Author    string `form:"author" valid:"Required;MaxSize(100)"`
	Authors   string `form:"authors" valid:"Required;MaxSize(100)"`
	Date      string `form:"date" valid:"Required;MaxSize(100)"`
	Link      string `form:"link" valid:"Required;MaxSize(100)"`
	Papercode string `form:"papercode" valid:"Required;MaxSize(100)"`
	Abstract  string `form:"abstract" valid:"Required;MaxSize(65535)"`
	Theyear   string `form:"theyear" valid:"Required;MaxSize(10)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Add article
// @Produce  json
// @Param tag_id body int true "TagID"
// @Param title body string true "Title"
// @Param desc body string true "Desc"
// @Param content body string true "Content"
// @Param created_by body string true "CreatedBy"
// @Param state body int true "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddArticleForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{ID: form.TagID}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	articleService := article_service.Article{
		TagID:     form.TagID,
		Title:     form.Title,
		Journal:   form.Journal,
		Author:    form.Author,
		Authors:   form.Authors,
		Date:      form.Date,
		Link:      form.Link,
		Papercode: form.Papercode,
		Abstract:  form.Abstract,
		Theyear:   form.Theyear,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}
	if err := articleService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditArticleForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	TagID      int    `form:"tag_id" valid:"Required;Min(1)"`
	Title      string `form:"title" valid:"Required;MaxSize(100)"`
	Journal    string `form:"journal" valid:"Required;MaxSize(100)"`
	Author     string `form:"author" valid:"Required;MaxSize(100)"`
	Authors    string `form:"authors" valid:"Required;MaxSize(100)"`
	Date       string `form:"date" valid:"Required;MaxSize(100)"`
	Link       string `form:"link" valid:"Required;MaxSize(100)"`
	Papercode  string `form:"papercode" valid:"Required;MaxSize(100)"`
	Abstract   string `form:"abstract" valid:"Required;MaxSize(65535)"`
	Theyear    string `form:"theyear" valid:"Required;MaxSize(10)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Update article
// @Produce  json
// @Param id path int true "ID"
// @Param tag_id body string false "TagID"
// @Param title body string false "Title"
// @Param desc body string false "Desc"
// @Param content body string false "Content"
// @Param modified_by body string true "ModifiedBy"
// @Param state body int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [put]
func EditArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditArticleForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	articleService := article_service.Article{
		ID:         form.ID,
		TagID:      form.TagID,
		Title:      form.Title,
		Journal:    form.Journal,
		Author:     form.Author,
		Authors:    form.Authors,
		Date:       form.Date,
		Link:       form.Link,
		Papercode:  form.Papercode,
		Abstract:   form.Abstract,
		Theyear:    form.Theyear,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	tagService := tag_service.Tag{ID: form.TagID}
	exists, err = tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = articleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Delete article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = articleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
