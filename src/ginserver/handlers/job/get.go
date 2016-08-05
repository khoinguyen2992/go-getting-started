package job

import (
	"ginserver/handlers"
	"ginserver/repository"
	"net/http"
)

type GetHandler struct {
	handlers.BaseHandler
}

func (this *GetHandler) Handle() {
	id := this.GetParam("id")

	jobRepo := repository.NewJobRepository()
	job, err := jobRepo.FindByID(id)

	if err != nil {
		this.ResponseError(http.StatusBadRequest, err.Error())
		return
	}

	this.ResponseSuccess(http.StatusOK, job)
}
