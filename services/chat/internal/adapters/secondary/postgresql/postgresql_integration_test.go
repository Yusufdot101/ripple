package postgresql

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

	_ "github.com/lib/pq"
)

type RepositoryTestSuite struct {
	suite.Suite
	dataSourceURL string
}

func (rts *RepositoryTestSuite) SetupSuite() {
	ctx := context.Background()
	ctr, err := postgres.Run(
		ctx,
		"postgres:18.3-alpine",
		postgres.WithDatabase("ripple_chat_service_test"),
		postgres.WithUsername("chat_service"),
		postgres.WithPassword("verystrongpassword"),
		postgres.BasicWaitStrategies(),
	)
	testcontainers.CleanupContainer(rts.T(), ctr)
	if err != nil {
		log.Fatalf("failed to start postgresql: %v", err)
	}

	rts.dataSourceURL, err = ctr.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("failed to obtain connection string: %v", err)
	}
}

func (rts *RepositoryTestSuite) SetupTest() {
	err := rts.truncateDB(context.Background())
	rts.Require().NoError(err)
}

func (rts *RepositoryTestSuite) truncateDB(ctx context.Context) error {
	db, err := sql.Open("postgres", rts.dataSourceURL)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	_, err = db.Exec(`
        DO $$ DECLARE
            r RECORD;
        BEGIN
            FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
                EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' CASCADE';
            END LOOP;
        END $$;
    `)

	return err
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
