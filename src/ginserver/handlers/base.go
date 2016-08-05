package handlers

import (
	"ginserver/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	SetContext(*gin.Context)
	Handle()
}

func NewHandler(handler Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		handler.SetContext(context)
		handler.Handle()
	}
}

type BaseHandler struct {
	Context *gin.Context
}

func (this *BaseHandler) SetContext(context *gin.Context) {
	this.Context = context
}

func (this *BaseHandler) GetParam(key string) string {
	return this.Context.Param(key)
}

func (this *BaseHandler) GetHeader(key string) string {
	return this.Context.Request.Header.Get(key)
}

func (this *BaseHandler) ParseRequest(request interface{}) error {
	return this.Context.Bind(request)
}

func (this *BaseHandler) ResponseSuccess(code int, data interface{}) {
	this.Response(code, responses.Success{
		Data: data,
	})
}

func (this *BaseHandler) ResponseError(statusCode int, message string) {
	this.Response(statusCode, responses.Error{
		Error: message,
	})
}

func (this *BaseHandler) ResponseForbidden() {
	this.ResponseError(http.StatusForbidden, "Forbidden")
}

func (this *BaseHandler) ResponseEmpty() {
	this.Context.AbortWithStatus(http.StatusNoContent)
}

func (this *BaseHandler) Response(code int, data interface{}) {
	this.Context.JSON(code, data)
}
