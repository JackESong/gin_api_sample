package v1
import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"gin_api_sample/common"
	"gin_api_sample/project/dao"
	"gin_api_sample/framework/setting"
)
//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := common.SUCCESS
	data["lists"] = dao.GetTags(common.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = dao.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : common.GetMsg(code),
		"data" : data,
	})
}
//新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := common.INVALID_PARAMS
	if ! valid.HasErrors() {
		if ! dao.ExistTagByName(name) {
			code = common.SUCCESS
			dao.AddTag(name, state, createdBy)
		} else {
			code = common.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : common.GetMsg(code),
		"data" : make(map[string]string),
	})
}
//修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")
	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	code := common.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = common.SUCCESS
		if dao.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			dao.EditTag(id, data)
		} else {
			code = common.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : common.GetMsg(code),
		"data" : make(map[string]string),
	})
}
//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := common.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = common.SUCCESS
		if dao.ExistTagByID(id) {
			dao.DeleteTag(id)
		} else {
			code = common.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : common.GetMsg(code),
		"data" : make(map[string]string),
	})
}
