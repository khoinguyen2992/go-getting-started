package repository

import (
	"errors"
	"ginserver/models"

	"gopkg.in/mgo.v2/bson"
)

type JobRepository struct {
	BaseRepository
}

const (
	JOB_COLLECTION = "jobs"
)

func NewJobRepository() *JobRepository {
	return &JobRepository{
		BaseRepository: BaseRepository{
			Collection: db.C(JOB_COLLECTION),
		},
	}
}

func (this *JobRepository) FindByID(id string) (models.Job, error) {
	job := models.Job{}
	if !bson.IsObjectIdHex(id) {
		return job, errors.New("Invalid id")
	}

	err := this.Collection.Find(bson.M{
		"_id": bson.ObjectIdHex(id),
	}).One(&job)
	return job, err
}
