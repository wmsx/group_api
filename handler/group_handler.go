package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	groupProto "github.com/wmsx/group_svc/proto/group"
	mygin "github.com/wmsx/pkg/gin"
	"strconv"
)

type GroupHandler struct {
	groupClient groupProto.GroupService
}

func (h *GroupHandler) GetAllDiscussGroupList(ctx *gin.Context) {
	var (
		mengerId                   int64
		getAllDiscussGroupRequest  *groupProto.GetAllDiscussGroupRequest
		getAllDiscussGroupResponse *groupProto.GetAllDiscussGroupResponse
		err                        error
	)
	app := mygin.Gin{C: ctx}
	discussGroupResults := make([]*DiscussGroupResult, 0)

	mengerId, err = strconv.ParseInt(ctx.GetHeader("uid"), 10, 64)
	if err != nil {
		log.Error("获取用户id失败 err: ", err)
		app.Response(discussGroupResults)
		return
	}
	if mengerId  <= 0 {
		log.Error("获取用户id <= 0, mengerId: ", mengerId)
		app.Response(discussGroupResults)
		return
	}

	getAllDiscussGroupRequest = &groupProto.GetAllDiscussGroupRequest{
		MengerId: mengerId,
	}
	getAllDiscussGroupResponse, err = h.groupClient.GetAllDiscussGroup(ctx, getAllDiscussGroupRequest)
	if err != nil {
		log.Error("【rpc】【group_svc】调用失败 err:", err)
		app.ServerErrorResponse()
		return
	}
	if getAllDiscussGroupResponse.ErrorMsg != nil {
		app.LogicErrorResponse(getAllDiscussGroupResponse.ErrorMsg.Msg)
		return
	}

	for _, protoDiscussGroup := range getAllDiscussGroupResponse.DiscussGroups {
		discussGroupResult := &DiscussGroupResult{
			Id:    protoDiscussGroup.Id,
			Name:  protoDiscussGroup.Name,
			Cover: protoDiscussGroup.Cover,
		}
		discussGroupResults = append(discussGroupResults, discussGroupResult)
	}

	app.Response(discussGroupResults)
}

func NewGroupHandler(client client.Client) *GroupHandler {
	return &GroupHandler{
		groupClient: groupProto.NewGroupService(mengerSvcName, client),
	}
}
