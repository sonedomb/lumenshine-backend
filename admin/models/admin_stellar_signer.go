// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// AdminStellarSigner is an object representing the database table.
type AdminStellarSigner struct {
	ID                        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	StellarAccountPublicKeyID string    `boil:"stellar_account_public_key_id" json:"stellar_account_public_key_id" toml:"stellar_account_public_key_id" yaml:"stellar_account_public_key_id"`
	Name                      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description               string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	SignerPublicKey           string    `boil:"signer_public_key" json:"signer_public_key" toml:"signer_public_key" yaml:"signer_public_key"`
	SignerSecretSeed          string    `boil:"signer_secret_seed" json:"signer_secret_seed" toml:"signer_secret_seed" yaml:"signer_secret_seed"`
	Type                      string    `boil:"type" json:"type" toml:"type" yaml:"type"`
	CreatedAt                 time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt                 time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy                 string    `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`

	R *adminStellarSignerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L adminStellarSignerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AdminStellarSignerColumns = struct {
	ID                        string
	StellarAccountPublicKeyID string
	Name                      string
	Description               string
	SignerPublicKey           string
	SignerSecretSeed          string
	Type                      string
	CreatedAt                 string
	UpdatedAt                 string
	UpdatedBy                 string
}{
	ID: "id",
	StellarAccountPublicKeyID: "stellar_account_public_key_id",
	Name:             "name",
	Description:      "description",
	SignerPublicKey:  "signer_public_key",
	SignerSecretSeed: "signer_secret_seed",
	Type:             "type",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	UpdatedBy:        "updated_by",
}

// AdminStellarSignerRels is where relationship names are stored.
var AdminStellarSignerRels = struct {
	StellarAccountPublicKey string
}{
	StellarAccountPublicKey: "StellarAccountPublicKey",
}

// adminStellarSignerR is where relationships are stored.
type adminStellarSignerR struct {
	StellarAccountPublicKey *AdminStellarAccount
}

// NewStruct creates a new relationship struct
func (*adminStellarSignerR) NewStruct() *adminStellarSignerR {
	return &adminStellarSignerR{}
}

// adminStellarSignerL is where Load methods for each relationship are stored.
type adminStellarSignerL struct{}

var (
	adminStellarSignerColumns               = []string{"id", "stellar_account_public_key_id", "name", "description", "signer_public_key", "signer_secret_seed", "type", "created_at", "updated_at", "updated_by"}
	adminStellarSignerColumnsWithoutDefault = []string{"stellar_account_public_key_id", "name", "description", "signer_public_key", "signer_secret_seed", "type", "updated_by"}
	adminStellarSignerColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	adminStellarSignerPrimaryKeyColumns     = []string{"id"}
)

type (
	// AdminStellarSignerSlice is an alias for a slice of pointers to AdminStellarSigner.
	// This should generally be used opposed to []AdminStellarSigner.
	AdminStellarSignerSlice []*AdminStellarSigner
	// AdminStellarSignerHook is the signature for custom AdminStellarSigner hook methods
	AdminStellarSignerHook func(boil.Executor, *AdminStellarSigner) error

	adminStellarSignerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	adminStellarSignerType                 = reflect.TypeOf(&AdminStellarSigner{})
	adminStellarSignerMapping              = queries.MakeStructMapping(adminStellarSignerType)
	adminStellarSignerPrimaryKeyMapping, _ = queries.BindMapping(adminStellarSignerType, adminStellarSignerMapping, adminStellarSignerPrimaryKeyColumns)
	adminStellarSignerInsertCacheMut       sync.RWMutex
	adminStellarSignerInsertCache          = make(map[string]insertCache)
	adminStellarSignerUpdateCacheMut       sync.RWMutex
	adminStellarSignerUpdateCache          = make(map[string]updateCache)
	adminStellarSignerUpsertCacheMut       sync.RWMutex
	adminStellarSignerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var adminStellarSignerBeforeInsertHooks []AdminStellarSignerHook
var adminStellarSignerBeforeUpdateHooks []AdminStellarSignerHook
var adminStellarSignerBeforeDeleteHooks []AdminStellarSignerHook
var adminStellarSignerBeforeUpsertHooks []AdminStellarSignerHook

var adminStellarSignerAfterInsertHooks []AdminStellarSignerHook
var adminStellarSignerAfterSelectHooks []AdminStellarSignerHook
var adminStellarSignerAfterUpdateHooks []AdminStellarSignerHook
var adminStellarSignerAfterDeleteHooks []AdminStellarSignerHook
var adminStellarSignerAfterUpsertHooks []AdminStellarSignerHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AdminStellarSigner) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarSignerBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AdminStellarSigner) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarSignerBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AdminStellarSigner) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarSignerBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AdminStellarSigner) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarSignerBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AdminStellarSigner) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarSignerAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AdminStellarSigner) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarSignerAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AdminStellarSigner) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarSignerAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AdminStellarSigner) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarSignerAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AdminStellarSigner) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminStellarSignerAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAdminStellarSignerHook registers your hook function for all future operations.
func AddAdminStellarSignerHook(hookPoint boil.HookPoint, adminStellarSignerHook AdminStellarSignerHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		adminStellarSignerBeforeInsertHooks = append(adminStellarSignerBeforeInsertHooks, adminStellarSignerHook)
	case boil.BeforeUpdateHook:
		adminStellarSignerBeforeUpdateHooks = append(adminStellarSignerBeforeUpdateHooks, adminStellarSignerHook)
	case boil.BeforeDeleteHook:
		adminStellarSignerBeforeDeleteHooks = append(adminStellarSignerBeforeDeleteHooks, adminStellarSignerHook)
	case boil.BeforeUpsertHook:
		adminStellarSignerBeforeUpsertHooks = append(adminStellarSignerBeforeUpsertHooks, adminStellarSignerHook)
	case boil.AfterInsertHook:
		adminStellarSignerAfterInsertHooks = append(adminStellarSignerAfterInsertHooks, adminStellarSignerHook)
	case boil.AfterSelectHook:
		adminStellarSignerAfterSelectHooks = append(adminStellarSignerAfterSelectHooks, adminStellarSignerHook)
	case boil.AfterUpdateHook:
		adminStellarSignerAfterUpdateHooks = append(adminStellarSignerAfterUpdateHooks, adminStellarSignerHook)
	case boil.AfterDeleteHook:
		adminStellarSignerAfterDeleteHooks = append(adminStellarSignerAfterDeleteHooks, adminStellarSignerHook)
	case boil.AfterUpsertHook:
		adminStellarSignerAfterUpsertHooks = append(adminStellarSignerAfterUpsertHooks, adminStellarSignerHook)
	}
}

// OneG returns a single adminStellarSigner record from the query using the global executor.
func (q adminStellarSignerQuery) OneG() (*AdminStellarSigner, error) {
	return q.One(boil.GetDB())
}

// One returns a single adminStellarSigner record from the query.
func (q adminStellarSignerQuery) One(exec boil.Executor) (*AdminStellarSigner, error) {
	o := &AdminStellarSigner{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for admin_stellar_signer")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all AdminStellarSigner records from the query using the global executor.
func (q adminStellarSignerQuery) AllG() (AdminStellarSignerSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all AdminStellarSigner records from the query.
func (q adminStellarSignerQuery) All(exec boil.Executor) (AdminStellarSignerSlice, error) {
	var o []*AdminStellarSigner

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AdminStellarSigner slice")
	}

	if len(adminStellarSignerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all AdminStellarSigner records in the query, and panics on error.
func (q adminStellarSignerQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all AdminStellarSigner records in the query.
func (q adminStellarSignerQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count admin_stellar_signer rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q adminStellarSignerQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q adminStellarSignerQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if admin_stellar_signer exists")
	}

	return count > 0, nil
}

// StellarAccountPublicKey pointed to by the foreign key.
func (o *AdminStellarSigner) StellarAccountPublicKey(mods ...qm.QueryMod) adminStellarAccountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("public_key=?", o.StellarAccountPublicKeyID),
	}

	queryMods = append(queryMods, mods...)

	query := AdminStellarAccounts(queryMods...)
	queries.SetFrom(query.Query, "\"admin_stellar_account\"")

	return query
}

// LoadStellarAccountPublicKey allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (adminStellarSignerL) LoadStellarAccountPublicKey(e boil.Executor, singular bool, maybeAdminStellarSigner interface{}, mods queries.Applicator) error {
	var slice []*AdminStellarSigner
	var object *AdminStellarSigner

	if singular {
		object = maybeAdminStellarSigner.(*AdminStellarSigner)
	} else {
		slice = *maybeAdminStellarSigner.(*[]*AdminStellarSigner)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &adminStellarSignerR{}
		}
		args = append(args, object.StellarAccountPublicKeyID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &adminStellarSignerR{}
			}

			for _, a := range args {
				if a == obj.StellarAccountPublicKeyID {
					continue Outer
				}
			}

			args = append(args, obj.StellarAccountPublicKeyID)
		}
	}

	query := NewQuery(qm.From(`admin_stellar_account`), qm.WhereIn(`public_key in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AdminStellarAccount")
	}

	var resultSlice []*AdminStellarAccount
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AdminStellarAccount")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for admin_stellar_account")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for admin_stellar_account")
	}

	if len(adminStellarSignerAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.StellarAccountPublicKey = foreign
		if foreign.R == nil {
			foreign.R = &adminStellarAccountR{}
		}
		foreign.R.StellarAccountPublicKeyAdminStellarSigners = append(foreign.R.StellarAccountPublicKeyAdminStellarSigners, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StellarAccountPublicKeyID == foreign.PublicKey {
				local.R.StellarAccountPublicKey = foreign
				if foreign.R == nil {
					foreign.R = &adminStellarAccountR{}
				}
				foreign.R.StellarAccountPublicKeyAdminStellarSigners = append(foreign.R.StellarAccountPublicKeyAdminStellarSigners, local)
				break
			}
		}
	}

	return nil
}

// SetStellarAccountPublicKeyG of the adminStellarSigner to the related item.
// Sets o.R.StellarAccountPublicKey to related.
// Adds o to related.R.StellarAccountPublicKeyAdminStellarSigners.
// Uses the global database handle.
func (o *AdminStellarSigner) SetStellarAccountPublicKeyG(insert bool, related *AdminStellarAccount) error {
	return o.SetStellarAccountPublicKey(boil.GetDB(), insert, related)
}

// SetStellarAccountPublicKey of the adminStellarSigner to the related item.
// Sets o.R.StellarAccountPublicKey to related.
// Adds o to related.R.StellarAccountPublicKeyAdminStellarSigners.
func (o *AdminStellarSigner) SetStellarAccountPublicKey(exec boil.Executor, insert bool, related *AdminStellarAccount) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"admin_stellar_signer\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stellar_account_public_key_id"}),
		strmangle.WhereClause("\"", "\"", 2, adminStellarSignerPrimaryKeyColumns),
	)
	values := []interface{}{related.PublicKey, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StellarAccountPublicKeyID = related.PublicKey
	if o.R == nil {
		o.R = &adminStellarSignerR{
			StellarAccountPublicKey: related,
		}
	} else {
		o.R.StellarAccountPublicKey = related
	}

	if related.R == nil {
		related.R = &adminStellarAccountR{
			StellarAccountPublicKeyAdminStellarSigners: AdminStellarSignerSlice{o},
		}
	} else {
		related.R.StellarAccountPublicKeyAdminStellarSigners = append(related.R.StellarAccountPublicKeyAdminStellarSigners, o)
	}

	return nil
}

// AdminStellarSigners retrieves all the records using an executor.
func AdminStellarSigners(mods ...qm.QueryMod) adminStellarSignerQuery {
	mods = append(mods, qm.From("\"admin_stellar_signer\""))
	return adminStellarSignerQuery{NewQuery(mods...)}
}

// FindAdminStellarSignerG retrieves a single record by ID.
func FindAdminStellarSignerG(iD int, selectCols ...string) (*AdminStellarSigner, error) {
	return FindAdminStellarSigner(boil.GetDB(), iD, selectCols...)
}

// FindAdminStellarSigner retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAdminStellarSigner(exec boil.Executor, iD int, selectCols ...string) (*AdminStellarSigner, error) {
	adminStellarSignerObj := &AdminStellarSigner{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"admin_stellar_signer\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, adminStellarSignerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from admin_stellar_signer")
	}

	return adminStellarSignerObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AdminStellarSigner) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AdminStellarSigner) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no admin_stellar_signer provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(adminStellarSignerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	adminStellarSignerInsertCacheMut.RLock()
	cache, cached := adminStellarSignerInsertCache[key]
	adminStellarSignerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			adminStellarSignerColumns,
			adminStellarSignerColumnsWithDefault,
			adminStellarSignerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(adminStellarSignerType, adminStellarSignerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(adminStellarSignerType, adminStellarSignerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"admin_stellar_signer\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"admin_stellar_signer\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into admin_stellar_signer")
	}

	if !cached {
		adminStellarSignerInsertCacheMut.Lock()
		adminStellarSignerInsertCache[key] = cache
		adminStellarSignerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AdminStellarSigner record using the global executor.
// See Update for more documentation.
func (o *AdminStellarSigner) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the AdminStellarSigner.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AdminStellarSigner) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	adminStellarSignerUpdateCacheMut.RLock()
	cache, cached := adminStellarSignerUpdateCache[key]
	adminStellarSignerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			adminStellarSignerColumns,
			adminStellarSignerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update admin_stellar_signer, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"admin_stellar_signer\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, adminStellarSignerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(adminStellarSignerType, adminStellarSignerMapping, append(wl, adminStellarSignerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update admin_stellar_signer row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for admin_stellar_signer")
	}

	if !cached {
		adminStellarSignerUpdateCacheMut.Lock()
		adminStellarSignerUpdateCache[key] = cache
		adminStellarSignerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q adminStellarSignerQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for admin_stellar_signer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for admin_stellar_signer")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AdminStellarSignerSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AdminStellarSignerSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminStellarSignerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"admin_stellar_signer\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, adminStellarSignerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in adminStellarSigner slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all adminStellarSigner")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AdminStellarSigner) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AdminStellarSigner) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no admin_stellar_signer provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(adminStellarSignerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	adminStellarSignerUpsertCacheMut.RLock()
	cache, cached := adminStellarSignerUpsertCache[key]
	adminStellarSignerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			adminStellarSignerColumns,
			adminStellarSignerColumnsWithDefault,
			adminStellarSignerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			adminStellarSignerColumns,
			adminStellarSignerPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert admin_stellar_signer, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(adminStellarSignerPrimaryKeyColumns))
			copy(conflict, adminStellarSignerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"admin_stellar_signer\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(adminStellarSignerType, adminStellarSignerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(adminStellarSignerType, adminStellarSignerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert admin_stellar_signer")
	}

	if !cached {
		adminStellarSignerUpsertCacheMut.Lock()
		adminStellarSignerUpsertCache[key] = cache
		adminStellarSignerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single AdminStellarSigner record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AdminStellarSigner) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single AdminStellarSigner record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AdminStellarSigner) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AdminStellarSigner provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), adminStellarSignerPrimaryKeyMapping)
	sql := "DELETE FROM \"admin_stellar_signer\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from admin_stellar_signer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for admin_stellar_signer")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q adminStellarSignerQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no adminStellarSignerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from admin_stellar_signer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for admin_stellar_signer")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AdminStellarSignerSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AdminStellarSignerSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AdminStellarSigner slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(adminStellarSignerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminStellarSignerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"admin_stellar_signer\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminStellarSignerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from adminStellarSigner slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for admin_stellar_signer")
	}

	if len(adminStellarSignerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AdminStellarSigner) ReloadG() error {
	if o == nil {
		return errors.New("models: no AdminStellarSigner provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AdminStellarSigner) Reload(exec boil.Executor) error {
	ret, err := FindAdminStellarSigner(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminStellarSignerSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AdminStellarSignerSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminStellarSignerSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AdminStellarSignerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminStellarSignerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"admin_stellar_signer\".* FROM \"admin_stellar_signer\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminStellarSignerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AdminStellarSignerSlice")
	}

	*o = slice

	return nil
}

// AdminStellarSignerExistsG checks if the AdminStellarSigner row exists.
func AdminStellarSignerExistsG(iD int) (bool, error) {
	return AdminStellarSignerExists(boil.GetDB(), iD)
}

// AdminStellarSignerExists checks if the AdminStellarSigner row exists.
func AdminStellarSignerExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"admin_stellar_signer\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if admin_stellar_signer exists")
	}

	return exists, nil
}
