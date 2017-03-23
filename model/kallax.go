// IMPORTANT! This is auto generated code by https://github.com/src-d/go-kallax
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package model

import (
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

// NewRepository returns a new instance of Repository.
func NewRepository() (record *Repository) {
	return newRepository()
}

// GetID returns the primary key of the model.
func (r *Repository) GetID() kallax.Identifier {
	return (*kallax.ULID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Repository) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.ULID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "endpoints":
		return types.Slice(&r.Endpoints), nil
	case "status":
		return &r.Status, nil
	case "fetched_at":
		return types.Nullable(&r.FetchedAt), nil
	case "fetch_error_at":
		return types.Nullable(&r.FetchErrorAt), nil
	case "last_commit_at":
		return types.Nullable(&r.LastCommitAt), nil
	case "_references":
		return types.JSON(&r.References), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Repository: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Repository) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "endpoints":
		return types.Slice(r.Endpoints), nil
	case "status":
		return (string)(r.Status), nil
	case "fetched_at":
		if r.FetchedAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.FetchedAt, nil
	case "fetch_error_at":
		if r.FetchErrorAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.FetchErrorAt, nil
	case "last_commit_at":
		if r.LastCommitAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.LastCommitAt, nil
	case "_references":
		return types.JSON(r.References), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Repository: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Repository) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Repository has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Repository) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Repository has no relationships")
}

// RepositoryStore is the entity to access the records of the type Repository
// in the database.
type RepositoryStore struct {
	*kallax.Store
}

// NewRepositoryStore creates a new instance of RepositoryStore
// using a SQL database.
func NewRepositoryStore(db *sql.DB) *RepositoryStore {
	return &RepositoryStore{kallax.NewStore(db)}
}

// Insert inserts a Repository in the database. A non-persisted object is
// required for this operation.
func (s *RepositoryStore) Insert(record *Repository) error {

	if err := record.BeforeSave(); err != nil {
		return err
	}

	return s.Store.Insert(Schema.Repository.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *RepositoryStore) Update(record *Repository, cols ...kallax.SchemaField) (updated int64, err error) {

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	return s.Store.Update(Schema.Repository.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *RepositoryStore) Save(record *Repository) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *RepositoryStore) Delete(record *Repository) error {

	return s.Store.Delete(Schema.Repository.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *RepositoryStore) Find(q *RepositoryQuery) (*RepositoryResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewRepositoryResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *RepositoryStore) MustFind(q *RepositoryQuery) *RepositoryResultSet {
	return NewRepositoryResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *RepositoryStore) Count(q *RepositoryQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *RepositoryStore) MustCount(q *RepositoryQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *RepositoryStore) FindOne(q *RepositoryQuery) (*Repository, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *RepositoryStore) MustFindOne(q *RepositoryQuery) *Repository {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Repository with the data in the database and
// makes it writable.
func (s *RepositoryStore) Reload(record *Repository) error {
	return s.Store.Reload(Schema.Repository.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *RepositoryStore) Transaction(callback func(*RepositoryStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&RepositoryStore{store})
	})
}

// RepositoryQuery is the object used to create queries for the Repository
// entity.
type RepositoryQuery struct {
	*kallax.BaseQuery
}

// NewRepositoryQuery returns a new instance of RepositoryQuery.
func NewRepositoryQuery() *RepositoryQuery {
	return &RepositoryQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Repository.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *RepositoryQuery) Select(columns ...kallax.SchemaField) *RepositoryQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *RepositoryQuery) SelectNot(columns ...kallax.SchemaField) *RepositoryQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *RepositoryQuery) Copy() *RepositoryQuery {
	return &RepositoryQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *RepositoryQuery) Order(cols ...kallax.ColumnOrder) *RepositoryQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *RepositoryQuery) BatchSize(size uint64) *RepositoryQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *RepositoryQuery) Limit(n uint64) *RepositoryQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *RepositoryQuery) Offset(n uint64) *RepositoryQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *RepositoryQuery) Where(cond kallax.Condition) *RepositoryQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values, it will do nothing
func (q *RepositoryQuery) FindByID(v ...kallax.ULID) *RepositoryQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Repository.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value
func (q *RepositoryQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *RepositoryQuery {
	return q.Where(cond(Schema.Repository.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value
func (q *RepositoryQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *RepositoryQuery {
	return q.Where(cond(Schema.Repository.UpdatedAt, v))
}

// FindByEndpoints adds a new filter to the query that will require that
// the Endpoints property contains all the passed values; if no passed values, it will do nothing
func (q *RepositoryQuery) FindByEndpoints(v ...string) *RepositoryQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.ArrayContains(Schema.Repository.Endpoints, values...))
}

// FindByStatus adds a new filter to the query that will require that
// the Status property is equal to the passed value
func (q *RepositoryQuery) FindByStatus(v FetchStatus) *RepositoryQuery {
	return q.Where(kallax.Eq(Schema.Repository.Status, v))
}

// FindByFetchedAt adds a new filter to the query that will require that
// the FetchedAt property is equal to the passed value
func (q *RepositoryQuery) FindByFetchedAt(cond kallax.ScalarCond, v time.Time) *RepositoryQuery {
	return q.Where(cond(Schema.Repository.FetchedAt, v))
}

// FindByFetchErrorAt adds a new filter to the query that will require that
// the FetchErrorAt property is equal to the passed value
func (q *RepositoryQuery) FindByFetchErrorAt(cond kallax.ScalarCond, v time.Time) *RepositoryQuery {
	return q.Where(cond(Schema.Repository.FetchErrorAt, v))
}

// FindByLastCommitAt adds a new filter to the query that will require that
// the LastCommitAt property is equal to the passed value
func (q *RepositoryQuery) FindByLastCommitAt(cond kallax.ScalarCond, v time.Time) *RepositoryQuery {
	return q.Where(cond(Schema.Repository.LastCommitAt, v))
}

// RepositoryResultSet is the set of results returned by a query to the
// database.
type RepositoryResultSet struct {
	ResultSet kallax.ResultSet
	last      *Repository
	lastErr   error
}

// NewRepositoryResultSet creates a new result set for rows of the type
// Repository.
func NewRepositoryResultSet(rs kallax.ResultSet) *RepositoryResultSet {
	return &RepositoryResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *RepositoryResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Repository.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Repository)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Repository")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *RepositoryResultSet) Get() (*Repository, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *RepositoryResultSet) ForEach(fn func(*Repository) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *RepositoryResultSet) All() ([]*Repository, error) {
	var result []*Repository
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *RepositoryResultSet) One() (*Repository, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *RepositoryResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *RepositoryResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Repository *schemaRepository
}

type schemaRepository struct {
	*kallax.BaseSchema
	ID           kallax.SchemaField
	CreatedAt    kallax.SchemaField
	UpdatedAt    kallax.SchemaField
	Endpoints    kallax.SchemaField
	Status       kallax.SchemaField
	FetchedAt    kallax.SchemaField
	FetchErrorAt kallax.SchemaField
	LastCommitAt kallax.SchemaField
	References   *schemaRepositoryReferences
}

type schemaRepositoryReferences struct {
	*kallax.BaseSchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
	Name      kallax.SchemaField
	Hash      kallax.SchemaField
	Init      kallax.SchemaField
	Roots     kallax.SchemaField
}

func (s *schemaRepositoryReferences) At(n int) *schemaRepositoryReferences {
	return &schemaRepositoryReferences{
		BaseSchemaField: kallax.NewSchemaField("_references").(*kallax.BaseSchemaField),
		CreatedAt:       kallax.NewJSONSchemaKey(kallax.JSONAny, "_references", "CreatedAt"),
		UpdatedAt:       kallax.NewJSONSchemaKey(kallax.JSONAny, "_references", "UpdatedAt"),
		Name:            kallax.NewJSONSchemaKey(kallax.JSONText, "_references", fmt.Sprint(n), "Name"),
		Hash:            kallax.NewJSONSchemaArray("_references", fmt.Sprint(n), "Hash"),
		Init:            kallax.NewJSONSchemaArray("_references", fmt.Sprint(n), "Init"),
		Roots:           kallax.NewJSONSchemaArray("_references", fmt.Sprint(n), "Roots"),
	}
}

type schemaRepositoryReferencesTimestamps struct {
	*kallax.BaseSchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
}

var Schema = &schema{
	Repository: &schemaRepository{
		BaseSchema: kallax.NewBaseSchema(
			"repositories",
			"__repository",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Repository)
			},
			false,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("endpoints"),
			kallax.NewSchemaField("status"),
			kallax.NewSchemaField("fetched_at"),
			kallax.NewSchemaField("fetch_error_at"),
			kallax.NewSchemaField("last_commit_at"),
			kallax.NewSchemaField("_references"),
		),
		ID:           kallax.NewSchemaField("id"),
		CreatedAt:    kallax.NewSchemaField("created_at"),
		UpdatedAt:    kallax.NewSchemaField("updated_at"),
		Endpoints:    kallax.NewSchemaField("endpoints"),
		Status:       kallax.NewSchemaField("status"),
		FetchedAt:    kallax.NewSchemaField("fetched_at"),
		FetchErrorAt: kallax.NewSchemaField("fetch_error_at"),
		LastCommitAt: kallax.NewSchemaField("last_commit_at"),
		References: &schemaRepositoryReferences{
			BaseSchemaField: kallax.NewSchemaField("_references").(*kallax.BaseSchemaField),
			CreatedAt:       kallax.NewJSONSchemaKey(kallax.JSONAny, "_references", "CreatedAt"),
			UpdatedAt:       kallax.NewJSONSchemaKey(kallax.JSONAny, "_references", "UpdatedAt"),
			Name:            kallax.NewJSONSchemaKey(kallax.JSONText, "_references", "Name"),
			Hash:            kallax.NewJSONSchemaArray("_references", "Hash"),
			Init:            kallax.NewJSONSchemaArray("_references", "Init"),
			Roots:           kallax.NewJSONSchemaArray("_references", "Roots"),
		},
	},
}
