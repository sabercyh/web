package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
)

const (
	ERROR_EXIST_TAG = 10001 + iota
	ERROR_EXIST_TAG_FAIL
	ERROR_NOT_EXIST_TAG
	ERROR_GET_TAGS_FAIL
	ERROR_COUNT_TAG_FAIL
	ERROR_ADD_TAG_FAIL
	ERROR_EDIT_TAG_FAIL
	ERROR_DELETE_TAG_FAIL

	ERROR_EXIST_ARTICLE
	ERROR_EXIST_ARTICLE_FAIL
	ERROR_CHECK_EXIST_ARTICLE_FAIL
	ERROR_NOT_EXIST_ARTICLE
	ERROR_GET_ARTICLES_FAIL
	ERROR_COUNT_ARTICLE_FAIL
	ERROR_ADD_ARTICLE_FAIL
	ERROR_GET_ARTICLE_FAIL
	ERROR_EDIT_ARTICLE_FAIL
	ERROR_DELETE_ARTICLE_FAIL
)

const (
	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001 + iota
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	ERROR_AUTH_TOKEN
	ERROR_AUTH
)
