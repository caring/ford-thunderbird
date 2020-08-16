package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/caring/go-packages/pkg/errors"
	"github.com/google/uuid"

	"github.com/caring/ford-thunderbird/pb"
)



// thunderbirdService provides an API for interacting with the thunderbirds table
type thunderbirdService struct {
	db    *sql.DB
	stmts map[string]*sql.Stmt
}

// Thunderbird is a struct representation of a row in the thunderbirds table
type Thunderbird struct {
	ID  	uuid.UUID
	Name  string
}

// protoThunderbird is an interface that most proto thunderbird objects will satisfy
type protoThunderbird interface {
	GetName() string
}

// NewThunderbird is a convenience helper cast a proto thunderbird to it's DB layer struct
func NewThunderbird(ID string, proto protoThunderbird) (*Thunderbird, error) {
	mID, err := ParseUUID(ID)
	if err != nil {
		return nil, err
	}

	return &Thunderbird{
		ID:  	mID,
		Name: proto.GetName(),
	}, nil
}

// ToProto casts a db thunderbird into a proto response object
func (m *Thunderbird) ToProto() *pb.ThunderbirdResponse {
	return &pb.ThunderbirdResponse{
		Id:  				m.ID.String(),
		Name:       m.Name,
	}
}

// Get fetches a single thunderbird from the db
func (svc *thunderbirdService) Get(ctx context.Context, ID uuid.UUID) (*Thunderbird, error) {
	return svc.get(ctx, false, ID)
}

// GetTx fetches a single thunderbird from the db inside of a tx from ctx
func (svc *thunderbirdService) GetTx(ctx context.Context, ID uuid.UUID) (*Thunderbird, error) {
	return svc.get(ctx, true, ID)
}

// get fetches a single thunderbird from the db
func (svc *thunderbirdService) get(ctx context.Context, useTx bool, ID uuid.UUID) (*Thunderbird, error) {
	errMsg := func() string { return "Error executing get thunderbird - " + fmt.Sprint(ID) }

	var (
		stmt *sql.Stmt
		err  error
		tx   *sql.Tx
	)

	if useTx {

		if tx, err = FromCtx(ctx); err != nil {
			return nil, err
		}

		stmt = tx.Stmt(svc.stmts["get-thunderbird"])
	} else {
		stmt = svc.stmts["get-thunderbird"]
	}

	p := Thunderbird{}

	err = stmt.QueryRowContext(ctx, ID).
		Scan(&m.ThunderbirdID, &m.Name)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(ErrNotFound, errMsg())
		}

		return nil, errors.Wrap(err, errMsg())
	}

	return &p, nil
}

// Create a new thunderbird
func (svc *thunderbirdService) Create(ctx context.Context, input *Thunderbird) error {
	return svc.create(ctx, false, input)
}

// CreateTx creates a new thunderbird withing a tx from ctx
func (svc *thunderbirdService) CreateTx(ctx context.Context, input *Thunderbird) error {
	return svc.create(ctx, true, input)
}

// create a new thunderbird. if useTx = true then it will attempt to create the thunderbird within a transaction
// from context.
func (svc *thunderbirdService) create(ctx context.Context, useTx bool, input *Thunderbird) error {
	errMsg := func() string { return "Error executing create thunderbird - " + fmt.Sprint(input) }

	var (
		stmt *sql.Stmt
		err  error
		tx   *sql.Tx
	)

	if useTx {

		if tx, err = FromCtx(ctx); err != nil {
			return err
		}

		stmt = tx.Stmt(svc.stmts["create-thunderbird"])
	} else {
		stmt = svc.stmts["create-thunderbird"]
	}

	result, err := stmt.ExecContext(ctx, input.ThunderbirdID, input.Name)
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	if rowCount == 0 {
		return errors.Wrap(ErrNotCreated, errMsg())
	}

	return nil
}

// Update updates a single thunderbird row in the DB
func (svc *thunderbirdService) Update(ctx context.Context, input *Thunderbird) error {
	return svc.update(ctx, false, input)
}

// UpdateTx updates a single thunderbird row in the DB within a tx from ctx
func (svc *thunderbirdService) UpdateTx(ctx context.Context, input *Thunderbird) error {
	return svc.update(ctx, true, input)
}

// update a thunderbird. if useTx = true then it will attempt to update the thunderbird within a transaction
// from context.
func (svc *thunderbirdService) update(ctx context.Context, useTx bool, input *Thunderbird) error {
	errMsg := func() string { return "Error executing update thunderbird - " + fmt.Sprint(input) }

	var (
		stmt *sql.Stmt
		err  error
		tx   *sql.Tx
	)

	if useTx {

		if tx, err = FromCtx(ctx); err != nil {
			return err
		}

		stmt = tx.Stmt(svc.stmts["update-thunderbird"])
	} else {
		stmt = svc.stmts["update-thunderbird"]
	}

	result, err := stmt.ExecContext(ctx, input.Name, input.ThunderbirdID)
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	if rowCount == 0 {
		return errors.Wrap(ErrNoRowsAffected, errMsg())
	}

	return nil
}

// Delete sets deleted_at for a single thunderbirds row
func (svc *thunderbirdService) Delete(ctx context.Context, ID uuid.UUID) error {
	return svc.delete(ctx, false, ID)
}

// DeleteTx sets deleted_at for a single thunderbirds row within a tx from ctx
func (svc *thunderbirdService) DeleteTx(ctx context.Context, ID uuid.UUID) error {
	return svc.delete(ctx, true, ID)
}

// delete a thunderbird by setting deleted at. if useTx = true then it will attempt to delete the thunderbird within a transaction
// from context.
func (svc *thunderbirdService) delete(ctx context.Context, useTx bool, ID uuid.UUID) error {
	errMsg := func() string { return "Error executing delete thunderbird - " + ID.String() }

	var (
		stmt *sql.Stmt
		err  error
		tx   *sql.Tx
	)

	if useTx {

		if tx, err = FromCtx(ctx); err != nil {
			return err
		}

		stmt = tx.Stmt(svc.stmts["delete-thunderbird"])
	} else {
		stmt = svc.stmts["delete-thunderbird"]
	}

	result, err := stmt.ExecContext(ctx, ID)
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, errMsg())
	}

	if rowCount == 0 {
		return errors.Wrap(ErrNotFound, errMsg())
	}

	return nil
}

