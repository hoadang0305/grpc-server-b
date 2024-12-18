package serviceimplement

import (
	"context"
	"github.com/hoadang0305/grpc-server-b/public/domain/entity"
	"github.com/hoadang0305/grpc-server-b/public/domain/model"
	"github.com/hoadang0305/grpc-server-b/public/repository"
	"github.com/hoadang0305/grpc-server-b/public/service"
)

type ActorService struct {
	actorRepository repository.ActorRepository
}

func NewActorService(actorRepository repository.ActorRepository) service.ActorService {
	return &ActorService{actorRepository: actorRepository}
}

func (service *ActorService) GetAllActor(ctx context.Context) []entity.Actor {
	return service.actorRepository.GetAllActors(ctx)
}

func (service *ActorService) GetActorById(ctx context.Context, id int64) (*entity.Actor, error) {
	return service.actorRepository.GetActorByID(ctx, id)
}

func (service *ActorService) CreateActor(ctx context.Context, actorRequest model.ActorRequest) (*entity.Actor, error) {
	actor := &entity.Actor{
		FirstName: actorRequest.FirstName,
		LastName:  actorRequest.LastName,
	}
	err := service.actorRepository.CreateActor(ctx, actor)
	if err != nil {
		return nil, err
	}
	return actor, nil
}

func (service *ActorService) UpdateActor(ctx context.Context, actorRequest model.ActorRequest, actorId int64) (*entity.Actor, error) {
	actor := &entity.Actor{
		FirstName: actorRequest.FirstName,
		LastName:  actorRequest.LastName,
	}
	updatedActor, err := service.actorRepository.UpdateActor(ctx, actor, actorId)
	if err != nil {
		return nil, err
	}
	return updatedActor, nil
}

func (service *ActorService) DeleteActor(ctx context.Context, id int64) error {
	return service.actorRepository.DeleteActor(ctx, id)
}
