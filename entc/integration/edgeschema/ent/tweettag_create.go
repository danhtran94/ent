// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TweetTagCreate is the builder for creating a TweetTag entity.
type TweetTagCreate struct {
	config
	mutation *TweetTagMutation
	hooks    []Hook
}

// SetAddedAt sets the "added_at" field.
func (ttc *TweetTagCreate) SetAddedAt(t time.Time) *TweetTagCreate {
	ttc.mutation.SetAddedAt(t)
	return ttc
}

// SetNillableAddedAt sets the "added_at" field if the given value is not nil.
func (ttc *TweetTagCreate) SetNillableAddedAt(t *time.Time) *TweetTagCreate {
	if t != nil {
		ttc.SetAddedAt(*t)
	}
	return ttc
}

// SetTagID sets the "tag_id" field.
func (ttc *TweetTagCreate) SetTagID(i int) *TweetTagCreate {
	ttc.mutation.SetTagID(i)
	return ttc
}

// SetTweetID sets the "tweet_id" field.
func (ttc *TweetTagCreate) SetTweetID(i int) *TweetTagCreate {
	ttc.mutation.SetTweetID(i)
	return ttc
}

// SetID sets the "id" field.
func (ttc *TweetTagCreate) SetID(u uuid.UUID) *TweetTagCreate {
	ttc.mutation.SetID(u)
	return ttc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ttc *TweetTagCreate) SetNillableID(u *uuid.UUID) *TweetTagCreate {
	if u != nil {
		ttc.SetID(*u)
	}
	return ttc
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttc *TweetTagCreate) SetTag(t *Tag) *TweetTagCreate {
	return ttc.SetTagID(t.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (ttc *TweetTagCreate) SetTweet(t *Tweet) *TweetTagCreate {
	return ttc.SetTweetID(t.ID)
}

// Mutation returns the TweetTagMutation object of the builder.
func (ttc *TweetTagCreate) Mutation() *TweetTagMutation {
	return ttc.mutation
}

// Save creates the TweetTag in the database.
func (ttc *TweetTagCreate) Save(ctx context.Context) (*TweetTag, error) {
	var (
		err  error
		node *TweetTag
	)
	ttc.defaults()
	if len(ttc.hooks) == 0 {
		if err = ttc.check(); err != nil {
			return nil, err
		}
		node, err = ttc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TweetTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ttc.check(); err != nil {
				return nil, err
			}
			ttc.mutation = mutation
			if node, err = ttc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ttc.hooks) - 1; i >= 0; i-- {
			if ttc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ttc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TweetTag)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TweetTagMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ttc *TweetTagCreate) SaveX(ctx context.Context) *TweetTag {
	v, err := ttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttc *TweetTagCreate) Exec(ctx context.Context) error {
	_, err := ttc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttc *TweetTagCreate) ExecX(ctx context.Context) {
	if err := ttc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ttc *TweetTagCreate) defaults() {
	if _, ok := ttc.mutation.AddedAt(); !ok {
		v := tweettag.DefaultAddedAt()
		ttc.mutation.SetAddedAt(v)
	}
	if _, ok := ttc.mutation.ID(); !ok {
		v := tweettag.DefaultID()
		ttc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttc *TweetTagCreate) check() error {
	if _, ok := ttc.mutation.AddedAt(); !ok {
		return &ValidationError{Name: "added_at", err: errors.New(`ent: missing required field "TweetTag.added_at"`)}
	}
	if _, ok := ttc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag_id", err: errors.New(`ent: missing required field "TweetTag.tag_id"`)}
	}
	if _, ok := ttc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet_id", err: errors.New(`ent: missing required field "TweetTag.tweet_id"`)}
	}
	if _, ok := ttc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag", err: errors.New(`ent: missing required edge "TweetTag.tag"`)}
	}
	if _, ok := ttc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet", err: errors.New(`ent: missing required edge "TweetTag.tweet"`)}
	}
	return nil
}

func (ttc *TweetTagCreate) sqlSave(ctx context.Context) (*TweetTag, error) {
	_node, _spec := ttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ttc *TweetTagCreate) createSpec() (*TweetTag, *sqlgraph.CreateSpec) {
	var (
		_node = &TweetTag{config: ttc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tweettag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tweettag.FieldID,
			},
		}
	)
	if id, ok := ttc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ttc.mutation.AddedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tweettag.FieldAddedAt,
		})
		_node.AddedAt = value
	}
	if nodes := ttc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TagTable,
			Columns: []string{tweettag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TagID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ttc.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TweetTable,
			Columns: []string{tweettag.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tweet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TweetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TweetTagCreateBulk is the builder for creating many TweetTag entities in bulk.
type TweetTagCreateBulk struct {
	config
	builders []*TweetTagCreate
}

// Save creates the TweetTag entities in the database.
func (ttcb *TweetTagCreateBulk) Save(ctx context.Context) ([]*TweetTag, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ttcb.builders))
	nodes := make([]*TweetTag, len(ttcb.builders))
	mutators := make([]Mutator, len(ttcb.builders))
	for i := range ttcb.builders {
		func(i int, root context.Context) {
			builder := ttcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TweetTagMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ttcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ttcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ttcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ttcb *TweetTagCreateBulk) SaveX(ctx context.Context) []*TweetTag {
	v, err := ttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ttcb *TweetTagCreateBulk) Exec(ctx context.Context) error {
	_, err := ttcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttcb *TweetTagCreateBulk) ExecX(ctx context.Context) {
	if err := ttcb.Exec(ctx); err != nil {
		panic(err)
	}
}
