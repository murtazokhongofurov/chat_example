package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kafka_example/api-gateway/api/handler/v1/http"
	chpb "github.com/kafka_example/api-gateway/genproto/chat_service"
	"strconv"
	"time"
)

// ShowAccount godoc
// @Summary      Add an user info
// @Description  post string user info
// @Security     BearerAuth
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        phone query string true "PhoneNumber"
// @Success      201  {object}  models.UserLoginResponse
// @Failure      400  {object}  http.Response{}
// @Failure      404  {object}  http.Response{}
// @Failure      500  {object}  http.Response{}
// @Router       /user [GET]
// func (h *handlerV1) Login(c *gin.Context) {
// 	phone := c.Query("phone")
	
// 	data, err := h.serviceManager.ChatService().FindUser()
// }


// ShowAccount godoc
// @Summary      Add an user info
// @Description  post string user info
// @Security     BearerAuth
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user body chat_service.UserReq true "Post user"
// @Success      201  {object}  chat_service.UserRes
// @Failure      400  {object}  http.Response{}
// @Failure      404  {object}  http.Response{}
// @Failure      500  {object}  http.Response{}
// @Router       /user [POST]
func (h *handlerV1) PostUser(c *gin.Context) {
	var body chpb.UserReq
	err := c.ShouldBindJSON(&body)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	data, err := h.serviceManager.ChatService().AddUser(ctx, &body)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, data)

}

// ShowAccount godoc
// @Summary      Get an user info
// @Description  get string user info
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id path int true "UserID"
// @Success      200  {object}  chat_service.UserRes
// @Failure      400  {object}  http.Response{}
// @Failure      404  {object}  http.Response{}
// @Failure      500  {object}  http.Response{}
// @Router       /user/{id} [GET]
func (h *handlerV1) GetUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.ParseInt(id, 64, 10)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()
	data, err := h.serviceManager.ChatService().FindUser(ctx, &chpb.UserId{
		UserId: userId,
	})

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, data)

}
