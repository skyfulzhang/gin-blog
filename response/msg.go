package response

var CodeMsg = map[int]string{
	SUCCESS:                     "OK",
	PAGE_NOT_FOUND:              "请求路径错误",
	METHOD_NOT_ALLOW:            "请求方法不允许",
	INTERNAL_SERVER_ERROR:       "服务器错误",
	UNKNOWN_ERROR:               "未知错误",
	ACCESS_LIMIT:                "访问频率受限",
	REQUEST_TIMEOUT:             "请求超时",
	INVALID_PARAMS:              "请求参数错误",
	ERROR_NO_TOKEN:              "未发现Token",
	ERROR_INVALID_TOKEN:         "无效Token",
	ERROR_INVALID_USER:          "用户名或密码错误",
	ERROR_USER_EXIST:            "用户已存在",
	ERROR_CATEGORY_EXIST:        "分类已存在",
	ERROR_USER_NOT_ACTIVE:       "用户未激活",
	ERROR_USER_NOT_EXIST:        "用户不存在",
	ERROR_TOKEN_GENERATE_FAILED: "Token生成失败",
	ERROR_UNMARSHAL_FAILED:      "数据序列化失败",
	ERROR_ARTICLE_NOT_EXIST:     "文章不存在",
	ERROR_CATEGORY_NOT_EXIST:    "分类不存在",
	ERROR_COMMENT_NOT_EXIST:     "评论不存在",
	ERROR_QUERY_FAILED:          "查询失败",
	ERROR_CREATE_FAILED:         "创建失败",
	ERROR_UPDATE_FAILED:         "更新失败",
	ERROR_DELETE_FAILED:         "删除失败",
}

func GetMsg(code int) string {
	msg, ok := CodeMsg[code]
	if ok {
		return msg
	}
	return CodeMsg[UNKNOWN_ERROR]
}
