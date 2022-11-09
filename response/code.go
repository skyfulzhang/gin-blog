package response

const (
	SUCCESS                     = 200
	INVALID_PARAMS              = 400
	PAGE_NOT_FOUND              = 404
	METHOD_NOT_ALLOW            = 405
	INTERNAL_SERVER_ERROR       = 500
	UNKNOWN_ERROR               = 600
	ACCESS_LIMIT                = 700
	REQUEST_TIMEOUT             = 800
	ERROR_NO_TOKEN              = 10001
	ERROR_INVALID_TOKEN         = 10002
	ERROR_INVALID_USER          = 10003
	ERROR_USER_NOT_ACTIVE       = 10004
	ERROR_USER_EXIST            = 10005
	ERROR_CATEGORY_EXIST        = 10006
	ERROR_TOKEN_GENERATE_FAILED = 10007
	ERROR_UNMARSHAL_FAILED      = 10008
	ERROR_USER_NOT_EXIST        = 10011
	ERROR_ARTICLE_NOT_EXIST     = 10012
	ERROR_CATEGORY_NOT_EXIST    = 10013
	ERROR_COMMENT_NOT_EXIST     = 10014
	ERROR_QUERY_FAILED          = 20001
	ERROR_CREATE_FAILED         = 20002
	ERROR_UPDATE_FAILED         = 20003
	ERROR_DELETE_FAILED         = 20004
)