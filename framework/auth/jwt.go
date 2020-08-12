package auth
import (
	"sample_api/common"
	"sample_api/common/util"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = common.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = common.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = common.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = common.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != common.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : common.GetMsg(code),
				"data" : data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}