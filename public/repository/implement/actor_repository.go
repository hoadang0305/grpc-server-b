package repositoryimplement

import (
	"context"
	"github.com/hoadang0305/grpc-server-b/public/database"
	"github.com/hoadang0305/grpc-server-b/public/domain/entity"
	"github.com/hoadang0305/grpc-server-b/public/repository"
	"gorm.io/gorm"
)

type ActorRepository struct {
	db *gorm.DB
}

func NewActorRepository(db database.Db) repository.ActorRepository {
	return &ActorRepository{db: db}
}

func (repo *ActorRepository) GetAllActors(ctx context.Context) []entity.Actor {
	var actors []entity.Actor
	result := repo.db.WithContext(ctx).Find(&actors)
	if result.Error != nil {
		return []entity.Actor{}
	}
	return actors
}

func (repo *ActorRepository) GetActorByID(ctx context.Context, id int64) (*entity.Actor, error) {
	var actor entity.Actor
	result := repo.db.WithContext(ctx).First(&actor, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &actor, nil
}

func (repo *ActorRepository) CreateActor(ctx context.Context, actor *entity.Actor) error {
	err := repo.db.WithContext(ctx).Create(actor).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ActorRepository) UpdateActor(ctx context.Context, actor *entity.Actor, actorId int64) (*entity.Actor, error) {
	//update all columns of an actor
	if err := repo.db.WithContext(ctx).Model(&entity.Actor{Id: actorId}).Updates(&actor).Error; err != nil {
		return nil, err
	}

	var updatedActor entity.Actor
	if err := repo.db.WithContext(ctx).First(&updatedActor, actorId).Error; err != nil {
		return nil, err
	}
	return &updatedActor, nil
}

func (repo *ActorRepository) DeleteActor(ctx context.Context, id int64) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		//check actor exists
		var actor entity.Actor
		if err := tx.WithContext(ctx).First(&actor, id).Error; err != nil {
			return err
		}

		// Delete from film_actor
		if err := tx.Exec(`
            DELETE fa FROM film_actor fa
            WHERE fa.actor_id = ?`, id).Error; err != nil {
			return err
		}

		//Delete actor
		if err := tx.Delete(&actor).Error; err != nil {
			return err
		}
		return nil
	})
}
