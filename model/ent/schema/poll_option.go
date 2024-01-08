package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PollOption holds the schema definition for the PollOption entity.
type PollOption struct {
	ent.Schema
}

// Fields of the PollOption.
func (PollOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("vote_count"),
	}
}

// Edges of the PollOption.
func (PollOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users_voted", User.Type),

		edge.From("poll", Poll.Type).
			Ref("poll_options").
			Unique(),
	}
}
