package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// GorpMigration represents a row from 'paper-news_local.gorp_migrations'.
type GorpMigration struct {
	ID        string       `json:"id" db:"id"`                 // id
	AppliedAt sql.NullTime `json:"applied_at" db:"applied_at"` // applied_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [GorpMigration] exists in the database.
func (gm *GorpMigration) Exists() bool {
	return gm._exists
}

// Deleted returns true when the [GorpMigration] has been marked for deletion
// from the database.
func (gm *GorpMigration) Deleted() bool {
	return gm._deleted
}

// Insert inserts the [GorpMigration] to the database.
func (gm *GorpMigration) Insert(ctx context.Context, db DB) error {
	switch {
	case gm._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case gm._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO paper-news_local.gorp_migrations (` +
		`id, applied_at` +
		`) VALUES (` +
		`?, ?` +
		`)`
	// run
	logf(sqlstr, gm.ID, gm.AppliedAt)
	if _, err := db.ExecContext(ctx, sqlstr, gm.ID, gm.AppliedAt); err != nil {
		return logerror(err)
	}
	// set exists
	gm._exists = true
	return nil
}

// Update updates a [GorpMigration] in the database.
func (gm *GorpMigration) Update(ctx context.Context, db DB) error {
	switch {
	case !gm._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case gm._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE paper-news_local.gorp_migrations SET ` +
		`applied_at = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, gm.AppliedAt, gm.ID)
	if _, err := db.ExecContext(ctx, sqlstr, gm.AppliedAt, gm.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [GorpMigration] to the database.
func (gm *GorpMigration) Save(ctx context.Context, db DB) error {
	if gm.Exists() {
		return gm.Update(ctx, db)
	}
	return gm.Insert(ctx, db)
}

// Upsert performs an upsert for [GorpMigration].
func (gm *GorpMigration) Upsert(ctx context.Context, db DB) error {
	switch {
	case gm._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO paper-news_local.gorp_migrations (` +
		`id, applied_at` +
		`) VALUES (` +
		`?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`id = VALUES(id), applied_at = VALUES(applied_at)`
	// run
	logf(sqlstr, gm.ID, gm.AppliedAt)
	if _, err := db.ExecContext(ctx, sqlstr, gm.ID, gm.AppliedAt); err != nil {
		return logerror(err)
	}
	// set exists
	gm._exists = true
	return nil
}

// Delete deletes the [GorpMigration] from the database.
func (gm *GorpMigration) Delete(ctx context.Context, db DB) error {
	switch {
	case !gm._exists: // doesn't exist
		return nil
	case gm._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM paper-news_local.gorp_migrations ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, gm.ID)
	if _, err := db.ExecContext(ctx, sqlstr, gm.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	gm._deleted = true
	return nil
}

// GorpMigrationByID retrieves a row from 'paper-news_local.gorp_migrations' as a [GorpMigration].
//
// Generated from index 'gorp_migrations_id_pkey'.
func GorpMigrationByID(ctx context.Context, db DB, id string) (*GorpMigration, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, applied_at ` +
		`FROM paper-news_local.gorp_migrations ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	gm := GorpMigration{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&gm.ID, &gm.AppliedAt); err != nil {
		return nil, logerror(err)
	}
	return &gm, nil
}