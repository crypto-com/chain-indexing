package rdbprojectionbase_test

import (
	"fmt"

	. "github.com/crypto-com/chain-indexing/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	"github.com/crypto-com/chain-indexing/internal/primptr"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
)

var _ = Describe("RdbStoreImpl", func() {
	WithTestPgxConn(func(pgxConn *pg.PgxConn, pgMigrate *pg.Migrate) {
		BeforeEach(func() {
			_ = pgMigrate.Reset()
			pgMigrate.MustUp()
		})

		AfterEach(func() {
			_ = pgMigrate.Reset()
		})

		Describe("UpdateLastHandledEventHeight", func() {
			It("should insert projection record when the projection id does not have record", func() {
				store := rdbprojectionbase.NewStore(rdbprojectionbase.DEFAULT_TABLE)

				anyNonExistingProjectionId := "not_exist"
				anyHeight := int64(100)

				Expect(IsProjectionRowExist(pgxConn, anyNonExistingProjectionId)).To(BeFalse())

				err := store.UpdateLastHandledEventHeight(pgxConn.ToHandle(), anyNonExistingProjectionId, anyHeight)
				Expect(err).To(BeNil())

				Expect(IsProjectionRowExist(pgxConn, anyNonExistingProjectionId)).To(BeTrue())
			})
		})

		Describe("GetLastHandledEventHeight", func() {
			It("should return nil when the projection id does not have record", func() {
				store := rdbprojectionbase.NewStore(rdbprojectionbase.DEFAULT_TABLE)

				anyNonExistingProjectionId := "not_exist"

				actual, err := store.GetLastHandledEventHeight(pgxConn.ToHandle(), anyNonExistingProjectionId)
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(primptr.Int64Nil()))
			})
		})

		It("should update projection last handled height when record already exist", func() {
			var err error

			store := rdbprojectionbase.NewStore(rdbprojectionbase.DEFAULT_TABLE)

			anyProjectionId := "projection"
			initHeight := int64(1)

			err = store.UpdateLastHandledEventHeight(pgxConn.ToHandle(), anyProjectionId, initHeight)
			Expect(err).To(BeNil())

			newHeight := int64(2)
			err = store.UpdateLastHandledEventHeight(pgxConn.ToHandle(), anyProjectionId, newHeight)
			Expect(err).To(BeNil())

			var actualHeight *int64
			actualHeight, err = store.GetLastHandledEventHeight(pgxConn.ToHandle(), anyProjectionId)
			Expect(err).To(BeNil())
			Expect(actualHeight).To(Equal(primptr.Int64(newHeight)))
		})
	})
})

func IsProjectionRowExist(pgxConn *pg.PgxConn, projectionId string) bool {
	var rowCount int64
	if err := pgxConn.QueryRow(
		// nolint:gosec
		fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE id='%s'", rdbprojectionbase.DEFAULT_TABLE, projectionId),
	).Scan(&rowCount); err != nil {
		panic(err)
	}

	return rowCount == 1
}
