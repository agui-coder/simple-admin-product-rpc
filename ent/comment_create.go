// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/agui-coder/simple-admin-product-rpc/ent/comment"
	entType "github.com/agui-coder/simple-admin-product-rpc/ent/schema/enttype"
	"github.com/agui-coder/simple-admin-product-rpc/ent/sku"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CommentCreate) SetCreatedAt(t time.Time) *CommentCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableCreatedAt(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CommentCreate) SetUpdatedAt(t time.Time) *CommentCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableUpdatedAt(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CommentCreate) SetDeletedAt(t time.Time) *CommentCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableDeletedAt(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetUserID sets the "user_id" field.
func (cc *CommentCreate) SetUserID(u uint64) *CommentCreate {
	cc.mutation.SetUserID(u)
	return cc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableUserID(u *uint64) *CommentCreate {
	if u != nil {
		cc.SetUserID(*u)
	}
	return cc
}

// SetUserNickname sets the "user_nickname" field.
func (cc *CommentCreate) SetUserNickname(s string) *CommentCreate {
	cc.mutation.SetUserNickname(s)
	return cc
}

// SetNillableUserNickname sets the "user_nickname" field if the given value is not nil.
func (cc *CommentCreate) SetNillableUserNickname(s *string) *CommentCreate {
	if s != nil {
		cc.SetUserNickname(*s)
	}
	return cc
}

// SetUserAvatar sets the "user_avatar" field.
func (cc *CommentCreate) SetUserAvatar(s string) *CommentCreate {
	cc.mutation.SetUserAvatar(s)
	return cc
}

// SetNillableUserAvatar sets the "user_avatar" field if the given value is not nil.
func (cc *CommentCreate) SetNillableUserAvatar(s *string) *CommentCreate {
	if s != nil {
		cc.SetUserAvatar(*s)
	}
	return cc
}

// SetAnonymous sets the "anonymous" field.
func (cc *CommentCreate) SetAnonymous(b bool) *CommentCreate {
	cc.mutation.SetAnonymous(b)
	return cc
}

// SetNillableAnonymous sets the "anonymous" field if the given value is not nil.
func (cc *CommentCreate) SetNillableAnonymous(b *bool) *CommentCreate {
	if b != nil {
		cc.SetAnonymous(*b)
	}
	return cc
}

// SetOrderID sets the "order_id" field.
func (cc *CommentCreate) SetOrderID(u uint64) *CommentCreate {
	cc.mutation.SetOrderID(u)
	return cc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableOrderID(u *uint64) *CommentCreate {
	if u != nil {
		cc.SetOrderID(*u)
	}
	return cc
}

// SetOrderItemID sets the "order_item_id" field.
func (cc *CommentCreate) SetOrderItemID(u uint64) *CommentCreate {
	cc.mutation.SetOrderItemID(u)
	return cc
}

// SetNillableOrderItemID sets the "order_item_id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableOrderItemID(u *uint64) *CommentCreate {
	if u != nil {
		cc.SetOrderItemID(*u)
	}
	return cc
}

// SetSpuID sets the "spu_id" field.
func (cc *CommentCreate) SetSpuID(u uint64) *CommentCreate {
	cc.mutation.SetSpuID(u)
	return cc
}

// SetNillableSpuID sets the "spu_id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableSpuID(u *uint64) *CommentCreate {
	if u != nil {
		cc.SetSpuID(*u)
	}
	return cc
}

// SetSpuName sets the "spu_name" field.
func (cc *CommentCreate) SetSpuName(s string) *CommentCreate {
	cc.mutation.SetSpuName(s)
	return cc
}

// SetNillableSpuName sets the "spu_name" field if the given value is not nil.
func (cc *CommentCreate) SetNillableSpuName(s *string) *CommentCreate {
	if s != nil {
		cc.SetSpuName(*s)
	}
	return cc
}

// SetSkuID sets the "sku_id" field.
func (cc *CommentCreate) SetSkuID(u uint64) *CommentCreate {
	cc.mutation.SetSkuID(u)
	return cc
}

// SetNillableSkuID sets the "sku_id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableSkuID(u *uint64) *CommentCreate {
	if u != nil {
		cc.SetSkuID(*u)
	}
	return cc
}

// SetSkuPicURL sets the "sku_pic_url" field.
func (cc *CommentCreate) SetSkuPicURL(s string) *CommentCreate {
	cc.mutation.SetSkuPicURL(s)
	return cc
}

// SetSkuProperties sets the "sku_properties" field.
func (cc *CommentCreate) SetSkuProperties(etp []entType.SkuProperty) *CommentCreate {
	cc.mutation.SetSkuProperties(etp)
	return cc
}

// SetVisible sets the "visible" field.
func (cc *CommentCreate) SetVisible(b bool) *CommentCreate {
	cc.mutation.SetVisible(b)
	return cc
}

// SetNillableVisible sets the "visible" field if the given value is not nil.
func (cc *CommentCreate) SetNillableVisible(b *bool) *CommentCreate {
	if b != nil {
		cc.SetVisible(*b)
	}
	return cc
}

// SetScores sets the "scores" field.
func (cc *CommentCreate) SetScores(i int8) *CommentCreate {
	cc.mutation.SetScores(i)
	return cc
}

// SetNillableScores sets the "scores" field if the given value is not nil.
func (cc *CommentCreate) SetNillableScores(i *int8) *CommentCreate {
	if i != nil {
		cc.SetScores(*i)
	}
	return cc
}

// SetDescriptionScores sets the "description_scores" field.
func (cc *CommentCreate) SetDescriptionScores(i int8) *CommentCreate {
	cc.mutation.SetDescriptionScores(i)
	return cc
}

// SetNillableDescriptionScores sets the "description_scores" field if the given value is not nil.
func (cc *CommentCreate) SetNillableDescriptionScores(i *int8) *CommentCreate {
	if i != nil {
		cc.SetDescriptionScores(*i)
	}
	return cc
}

// SetBenefitScores sets the "benefit_scores" field.
func (cc *CommentCreate) SetBenefitScores(i int8) *CommentCreate {
	cc.mutation.SetBenefitScores(i)
	return cc
}

// SetNillableBenefitScores sets the "benefit_scores" field if the given value is not nil.
func (cc *CommentCreate) SetNillableBenefitScores(i *int8) *CommentCreate {
	if i != nil {
		cc.SetBenefitScores(*i)
	}
	return cc
}

// SetContent sets the "content" field.
func (cc *CommentCreate) SetContent(s string) *CommentCreate {
	cc.mutation.SetContent(s)
	return cc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cc *CommentCreate) SetNillableContent(s *string) *CommentCreate {
	if s != nil {
		cc.SetContent(*s)
	}
	return cc
}

// SetPicUrls sets the "pic_urls" field.
func (cc *CommentCreate) SetPicUrls(s string) *CommentCreate {
	cc.mutation.SetPicUrls(s)
	return cc
}

// SetNillablePicUrls sets the "pic_urls" field if the given value is not nil.
func (cc *CommentCreate) SetNillablePicUrls(s *string) *CommentCreate {
	if s != nil {
		cc.SetPicUrls(*s)
	}
	return cc
}

// SetReplyStatus sets the "reply_status" field.
func (cc *CommentCreate) SetReplyStatus(b bool) *CommentCreate {
	cc.mutation.SetReplyStatus(b)
	return cc
}

// SetNillableReplyStatus sets the "reply_status" field if the given value is not nil.
func (cc *CommentCreate) SetNillableReplyStatus(b *bool) *CommentCreate {
	if b != nil {
		cc.SetReplyStatus(*b)
	}
	return cc
}

// SetReplyUserID sets the "reply_user_id" field.
func (cc *CommentCreate) SetReplyUserID(i int) *CommentCreate {
	cc.mutation.SetReplyUserID(i)
	return cc
}

// SetNillableReplyUserID sets the "reply_user_id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableReplyUserID(i *int) *CommentCreate {
	if i != nil {
		cc.SetReplyUserID(*i)
	}
	return cc
}

// SetReplyContent sets the "reply_content" field.
func (cc *CommentCreate) SetReplyContent(s string) *CommentCreate {
	cc.mutation.SetReplyContent(s)
	return cc
}

// SetNillableReplyContent sets the "reply_content" field if the given value is not nil.
func (cc *CommentCreate) SetNillableReplyContent(s *string) *CommentCreate {
	if s != nil {
		cc.SetReplyContent(*s)
	}
	return cc
}

// SetReplyTime sets the "reply_time" field.
func (cc *CommentCreate) SetReplyTime(t time.Time) *CommentCreate {
	cc.mutation.SetReplyTime(t)
	return cc
}

// SetNillableReplyTime sets the "reply_time" field if the given value is not nil.
func (cc *CommentCreate) SetNillableReplyTime(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetReplyTime(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CommentCreate) SetID(u uint64) *CommentCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetSkusID sets the "skus" edge to the Sku entity by ID.
func (cc *CommentCreate) SetSkusID(id uint64) *CommentCreate {
	cc.mutation.SetSkusID(id)
	return cc
}

// SetNillableSkusID sets the "skus" edge to the Sku entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableSkusID(id *uint64) *CommentCreate {
	if id != nil {
		cc = cc.SetSkusID(*id)
	}
	return cc
}

// SetSkus sets the "skus" edge to the Sku entity.
func (cc *CommentCreate) SetSkus(s *Sku) *CommentCreate {
	return cc.SetSkusID(s.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommentCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if comment.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := comment.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if comment.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := comment.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Comment.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Comment.updated_at"`)}
	}
	if _, ok := cc.mutation.SkuPicURL(); !ok {
		return &ValidationError{Name: "sku_pic_url", err: errors.New(`ent: missing required field "Comment.sku_pic_url"`)}
	}
	return nil
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(comment.Table, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint64))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(comment.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.UserID(); ok {
		_spec.SetField(comment.FieldUserID, field.TypeUint64, value)
		_node.UserID = value
	}
	if value, ok := cc.mutation.UserNickname(); ok {
		_spec.SetField(comment.FieldUserNickname, field.TypeString, value)
		_node.UserNickname = value
	}
	if value, ok := cc.mutation.UserAvatar(); ok {
		_spec.SetField(comment.FieldUserAvatar, field.TypeString, value)
		_node.UserAvatar = value
	}
	if value, ok := cc.mutation.Anonymous(); ok {
		_spec.SetField(comment.FieldAnonymous, field.TypeBool, value)
		_node.Anonymous = value
	}
	if value, ok := cc.mutation.OrderID(); ok {
		_spec.SetField(comment.FieldOrderID, field.TypeUint64, value)
		_node.OrderID = value
	}
	if value, ok := cc.mutation.OrderItemID(); ok {
		_spec.SetField(comment.FieldOrderItemID, field.TypeUint64, value)
		_node.OrderItemID = value
	}
	if value, ok := cc.mutation.SpuID(); ok {
		_spec.SetField(comment.FieldSpuID, field.TypeUint64, value)
		_node.SpuID = value
	}
	if value, ok := cc.mutation.SpuName(); ok {
		_spec.SetField(comment.FieldSpuName, field.TypeString, value)
		_node.SpuName = value
	}
	if value, ok := cc.mutation.SkuPicURL(); ok {
		_spec.SetField(comment.FieldSkuPicURL, field.TypeString, value)
		_node.SkuPicURL = value
	}
	if value, ok := cc.mutation.SkuProperties(); ok {
		_spec.SetField(comment.FieldSkuProperties, field.TypeJSON, value)
		_node.SkuProperties = value
	}
	if value, ok := cc.mutation.Visible(); ok {
		_spec.SetField(comment.FieldVisible, field.TypeBool, value)
		_node.Visible = value
	}
	if value, ok := cc.mutation.Scores(); ok {
		_spec.SetField(comment.FieldScores, field.TypeInt8, value)
		_node.Scores = value
	}
	if value, ok := cc.mutation.DescriptionScores(); ok {
		_spec.SetField(comment.FieldDescriptionScores, field.TypeInt8, value)
		_node.DescriptionScores = value
	}
	if value, ok := cc.mutation.BenefitScores(); ok {
		_spec.SetField(comment.FieldBenefitScores, field.TypeInt8, value)
		_node.BenefitScores = value
	}
	if value, ok := cc.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := cc.mutation.PicUrls(); ok {
		_spec.SetField(comment.FieldPicUrls, field.TypeString, value)
		_node.PicUrls = value
	}
	if value, ok := cc.mutation.ReplyStatus(); ok {
		_spec.SetField(comment.FieldReplyStatus, field.TypeBool, value)
		_node.ReplyStatus = value
	}
	if value, ok := cc.mutation.ReplyUserID(); ok {
		_spec.SetField(comment.FieldReplyUserID, field.TypeInt, value)
		_node.ReplyUserID = value
	}
	if value, ok := cc.mutation.ReplyContent(); ok {
		_spec.SetField(comment.FieldReplyContent, field.TypeString, value)
		_node.ReplyContent = value
	}
	if value, ok := cc.mutation.ReplyTime(); ok {
		_spec.SetField(comment.FieldReplyTime, field.TypeTime, value)
		_node.ReplyTime = value
	}
	if nodes := cc.mutation.SkusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.SkusTable,
			Columns: []string{comment.SkusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sku.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SkuID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	err      error
	builders []*CommentCreate
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
