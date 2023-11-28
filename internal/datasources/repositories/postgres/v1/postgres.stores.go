package v1

import (
	"context"
	"github.com/jmoiron/sqlx"
	V1Domains "github.com/snykk/go-rest-boilerplate/internal/business/domains/v1"
	"github.com/snykk/go-rest-boilerplate/internal/datasources/records"
)

type postgresStoreRepository struct {
	conn *sqlx.DB
}

func NewPostgresStoreRepository(conn *sqlx.DB) V1Domains.StoreRepository {
	return &postgresStoreRepository{conn: conn}
}

func (r *postgresStoreRepository) Create(ctx context.Context, inDom *V1Domains.StoreDomain) (err error) {
	storeRecord := records.FromStoresV1Domain(inDom)

	_, err = r.conn.NamedQueryContext(ctx, `INSERT INTO stores(id, name, address, contacts) VALUES (uuid_generate_v4(), :name, :address, :contacts)`, storeRecord)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresStoreRepository) FindAll(ctx context.Context) (outDom []V1Domains.StoreDomain, err error) {
	var storeRecords []records.Stores

	err = r.conn.SelectContext(ctx, &storeRecords, `SELECT * FROM stores`)
	if err != nil {
		return []V1Domains.StoreDomain{}, err
	}

	return records.ToArrayOfStoresV1Domain(&storeRecords), nil

}

func (r *postgresStoreRepository) FindById(ctx context.Context, id string) (outDom V1Domains.StoreDomain, err error) {
	var storeRecord records.Stores

	err = r.conn.GetContext(ctx, &storeRecord, `SELECT * FROM stores WHERE id = $1`, id)
	if err != nil {
		return V1Domains.StoreDomain{}, err
	}

	return storeRecord.ToV1Domain(), nil
}

func (r *postgresStoreRepository) UpdateById(ctx context.Context, inDom *V1Domains.StoreDomain) (err error) {
	storeRecord := records.FromStoresV1Domain(inDom)

	_, err = r.conn.NamedExecContext(ctx, `
		UPDATE stores
		SET name = :name, address = :address, contacts = :contacts
		WHERE id = :id
	`, storeRecord)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresStoreRepository) DeleteById(ctx context.Context, id string) (err error) {
	_, err = r.conn.ExecContext(ctx, `DELETE FROM stores WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
