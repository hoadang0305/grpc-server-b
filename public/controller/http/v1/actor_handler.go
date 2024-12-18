package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hoadang0305/grpc-server-b/public/domain/entity"
	httpcommon "github.com/hoadang0305/grpc-server-b/public/domain/http_common"
	"github.com/hoadang0305/grpc-server-b/public/domain/model"
	"github.com/hoadang0305/grpc-server-b/public/service"
	"github.com/hoadang0305/grpc-server-b/public/utils/validation"
	"net/http"
	"strconv"
)

type ActorHandler struct {
	actorService service.ActorService
}

func NewActorHandler(actorService service.ActorService) *ActorHandler {
	return &ActorHandler{actorService: actorService}
}

func (handler *ActorHandler) GetAll(c *gin.Context) {
	actors := handler.actorService.GetAllActor(c.Request.Context())
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[[]entity.Actor](&actors))
}

// @Summary Get an actor
// @Description Get an actor
// @Tags Actor
// @Produce  json
// @Param id path int true "actorId" example(1)
// @Router /actors/{id} [get]
// @Success 200 {object} httpcommon.HttpResponse[entity.Actor]
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *ActorHandler) Get(c *gin.Context) {
	id := c.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "ID", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
	}

	actor, err := handler.actorService.GetActorById(c.Request.Context(), parsedId)
	if err != nil {
		if err.Error() == httpcommon.ErrorMessage.GormRecordNotFound {
			c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
				Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.RecordNotFound,
			}))
			return
		}
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[entity.Actor](actor))
}

// @Summary Create an actor
// @Description Create an actor
// @Tags Actor
// @Accept json
// @Param params body model.ActorRequest true "Actor payload"
// @Produce  json
// @Router /actors [post]
// @Success 201 {object} httpcommon.HttpResponse[entity.Actor]
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *ActorHandler) Create(c *gin.Context) {
	var actorRequest model.ActorRequest

	if err := validation.BindJsonAndValidate(c, &actorRequest); err != nil {
		return
	}

	actor, err := handler.actorService.CreateActor(c.Request.Context(), actorRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusCreated, httpcommon.NewSuccessResponse[entity.Actor](actor))
}

// @Summary Update an actor
// @Description Update an actor
// @Tags Actor
// @Accept json
// @Param id path int true "actorId" example(1)
// @Param request body model.ActorRequest true "Actor payload"
// @Produce  json
// @Router /actors/{id} [put]
// @Success 200 {object} httpcommon.HttpResponse[entity.Actor]
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *ActorHandler) Update(c *gin.Context) {
	//check param id
	id, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: "", Field: "ID", Code: httpcommon.ErrorResponseCode.MissingIdParameter,
		}))
		return
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "ID", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}

	var actorRequest model.ActorRequest
	//binding request
	if err = validation.BindJsonAndValidate(c, &actorRequest); err != nil {
		return
	}
	//update
	updatedActor, err := handler.actorService.UpdateActor(c.Request.Context(), actorRequest, parsedId)
	if err != nil {
		if err.Error() == httpcommon.ErrorMessage.GormRecordNotFound {
			c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
				Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.RecordNotFound,
			}))
			return
		}
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[entity.Actor](updatedActor))
}

// @Summary Delete an actor
// @Description Delete an actor with the given ID
// @Tags Actor
// @Produce json
// @Param id path int true "actorId" example(1)
// @Router /actors/{id} [delete]
// @Success 200 "Actor deleted successfully"
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *ActorHandler) Delete(c *gin.Context) {
	//check param id
	id := c.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "ID", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}

	err = handler.actorService.DeleteActor(c.Request.Context(), parsedId)
	if err != nil {
		if err.Error() == httpcommon.ErrorMessage.GormRecordNotFound {
			c.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
				Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.RecordNotFound,
			}))
			return
		}
		c.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusOK, httpcommon.NewSuccessResponse[any](nil))
}
