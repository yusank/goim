package v1

import (
	"github.com/gin-gonic/gin"

	messagev1 "github.com/yusank/goim/api/message/v1"
	"github.com/yusank/goim/apps/gateway/internal/service"
	"github.com/yusank/goim/pkg/mid"
	"github.com/yusank/goim/pkg/request"
	"github.com/yusank/goim/pkg/response"
	"github.com/yusank/goim/pkg/router"
)

type OfflineMessageRouter struct {
	router.Router
}

func NewOfflineMessageRouter() *OfflineMessageRouter {
	return &OfflineMessageRouter{
		Router: &router.BaseRouter{},
	}
}

func (r *OfflineMessageRouter) Load(g *gin.RouterGroup) {
	g.POST("/query", r.handleQueryOfflineMessage)
}

// @Summary 查询离线消息
// @Description 查询离线消息
// @Tags [gateway]offline_message
// @Accept  json
// @Produce  json
// @Param   req body messagev1.QueryOfflineMessageReq true "req"
// @Success 200 {object} messagev1.QueryOfflineMessageResp
// @Failure 200 {object} response.Response
// @Failure 401 {null} null
// @Router /gateway/v1/offline_message/query [post]
func (r *OfflineMessageRouter) handleQueryOfflineMessage(c *gin.Context) {
	req := new(messagev1.QueryOfflineMessageReq)
	if err := c.ShouldBindWith(req, request.PbJSONBinding{}); err != nil {
		response.ErrorResp(c, err)
		return
	}

	messages, err := service.GetOfflineMessageService().QueryOfflineMsg(mid.GetContext(c), req)
	if err != nil {
		response.ErrorResp(c, err)
		return
	}

	response.SuccessResp(c, messages, response.SetTotal(len(messages)))
}
