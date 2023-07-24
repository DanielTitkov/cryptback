package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/DanielTitkov/cryptgame/internal/domain"
	"github.com/google/uuid"
)

// Challenge holds the schema definition for the Challenge entity.
type Challenge struct {
	ent.Schema
}

// Fields of the Challenge.
func (Challenge) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("challenge").NotEmpty().Unique(),
		field.String("text").NotEmpty(),
		field.String("key").NotEmpty(),
		field.String("author").NotEmpty(),
		field.Enum("type").Values(domain.ChallengeTypeFirst, domain.ChallengeTypeName, domain.ChallengeTypeFrequency).Immutable(),
	}
}

// Edges of the Challenge.
func (Challenge) Edges() []ent.Edge {
	return nil
}

func (Challenge) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
