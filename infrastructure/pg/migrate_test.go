package pg_test

import (
	"database/sql"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	. "github.com/crypto-com/chain-indexing/infrastructure/pg"
	. "github.com/crypto-com/chain-indexing/test"
)

type TestPgConn struct {
	*sql.DB
}

func MustNewTestPgConn(config *ConnConfig) *TestPgConn {
	db, err := sql.Open("postgres", config.ToURL())
	if err != nil {
		panic(err)
	}

	return &TestPgConn{
		db,
	}
}

func (conn *TestPgConn) IsTableExists(name string) bool {
	row := conn.QueryRow(`SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE  table_schema = 'public'
		AND    table_name   = $1
	);
	`, name)

	var result bool
	if err := row.Scan(&result); err != nil {
		panic(err)
	}

	return result
}

var _ = Describe("PgMigrate", func() {
	WithTestPgConnConfig(func(config *ConnConfig) {
		var pgConn *TestPgConn
		BeforeEach(func() {
			migrate := MustNewMigrate(config, "./test/valid")
			migrate.SetReCreateOnReset(false)
			_ = migrate.Reset()

			pgConn = MustNewTestPgConn(config)
		})

		AfterEach(func() {
			migrate := MustNewMigrate(config, "./test/valid")
			migrate.SetReCreateOnReset(false)
			_ = migrate.Reset()
		})

		Describe("NewPgMigrate", func() {
			It("should return Error when source folder is invalid", func() {
				_, err := NewMigrate(config, "./not/exist/folder")

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("no such file or directory"))
			})

			It("should return Error when database cannot be connected", func() {
				invalidConfig := ConnConfig{
					Host:          "127.0.0.1",
					Port:          12345,
					MaybeUsername: nil,
					MaybePassword: nil,
					Database:      "AnyNonExistDB",
					SSL:           true,
				}
				_, err := NewMigrate(&invalidConfig, "./test/valid")

				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("failed to open database, postgres://127.0.0.1:12344/AnyNonExistDB: dial tcp 127.0.0.1:12345: connect: connection refused"))
			})
		})

		Describe("PgMigrate", func() {
			It("should implement RDbMigrate interface", func() {
				migrate, _ := NewMigrate(config, "./test/valid")
				var _ rdb.Migrate = migrate
			})
		})

		Describe("Up", func() {
			It("should return Error when migration up file is invalid", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/invalid")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("migration failed: syntax error at or near \"INVALID\" (column 1)"))
			})

			It("should execute up migration", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				Expect(pgConn.IsTableExists("tests")).To(BeFalse())

				err = migrate.Up()
				Expect(err).To(BeNil())

				version, dirty, err := migrate.Version()
				Expect(version).To(Equal(uint(1)))
				Expect(dirty).To(BeFalse())
				Expect(err).To(BeNil())

				Expect(pgConn.IsTableExists("tests")).To(BeTrue())
			})

			It("should succeed when it is already latest migration", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())
			})

			It("should execute up migration from current version to the latest migration", func() {
				var err error

				// Perform migration to version 1
				var migrate *Migrate
				migrate, err = NewMigrate(config, "./test/step-1")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())

				Expect(pgConn.IsTableExists("tests")).To(BeTrue())

				// step-2 contains one more migrations
				// Perform migration to latest version
				migrate, err = NewMigrate(config, "./test/step-2")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())

				Expect(pgConn.IsTableExists("tests2")).To(BeTrue())
			})

			It("should execute all the up migrations in the way", func() {
				var err error

				var migrate *Migrate

				Expect(pgConn.IsTableExists("tests")).To(BeFalse())
				Expect(pgConn.IsTableExists("tests2")).To(BeFalse())

				migrate, err = NewMigrate(config, "./test/step-2")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())

				Expect(pgConn.IsTableExists("tests")).To(BeTrue())
				Expect(pgConn.IsTableExists("tests2")).To(BeTrue())
			})
		})

		Describe("Down", func() {
			It("should do nothing when no up migration has been executed", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Down()
				Expect(err).To(BeNil())
			})

			It("should execute down migration", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeTrue())

				err = migrate.Down()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeFalse())
			})

			It("should execute down migration from current version", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/step-2")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeTrue())
				Expect(pgConn.IsTableExists("tests2")).To(BeTrue())

				err = migrate.Steps(-1)
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeTrue())
				Expect(pgConn.IsTableExists("tests2")).To(BeFalse())

				err = migrate.Down()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeFalse())
				Expect(pgConn.IsTableExists("tests2")).To(BeFalse())
			})

			It("should execute down migration all the way down", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/step-2")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeTrue())
				Expect(pgConn.IsTableExists("tests2")).To(BeTrue())

				err = migrate.Down()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeFalse())
				Expect(pgConn.IsTableExists("tests2")).To(BeFalse())
			})
		})

		Describe("Steps", func() {
			It("should return error when positive delta exceed possible up migrations", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Steps(10)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("limit 9 short"))
			})

			Context("When delta is negative and exceed possible down migrations", func() {
				var migrate *Migrate
				BeforeEach(func() {
					var err error

					migrate, err = NewMigrate(config, "./test/valid")
					Expect(err).To(BeNil())

					err = migrate.Up()
					Expect(err).To(BeNil())
				})

				It("should return error", func() {
					err := migrate.Steps(-20)
					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(ContainSubstring("limit 19 short"))
				})

				It("should execute all the migrations in the way", func() {
					Expect(pgConn.IsTableExists("tests")).To(BeTrue())

					err := migrate.Steps(-20)
					Expect(err).NotTo(BeNil())
					Expect(pgConn.IsTableExists("tests")).To(BeFalse())
				})
			})

			It("should return error when negative delta and no migrations have been executed", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Steps(-10)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("file does not exist"))
			})

			It("should execute up migrations when delta is positive and there are un-migrated migrations", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/step-1")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())

				Expect(pgConn.IsTableExists("tests2")).To(BeFalse())

				migrate, err = NewMigrate(config, "./test/step-2")
				Expect(err).To(BeNil())
				err = migrate.Steps(1)
				Expect(err).To(BeNil())

				Expect(pgConn.IsTableExists("tests2")).To(BeTrue())
			})

			It("should execute down migrations when delta is negative and there are already executed up migrations", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeTrue())

				err = migrate.Steps(-1)
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeFalse())
			})
		})

		Describe("Reset", func() {
			It("should do nothing when no migrations have been executed", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Reset()
				Expect(err).To(BeNil())
			})

			It("should reset all the migrations", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeTrue())

				err = migrate.Reset()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeFalse())
			})

			It("should reset all the migrations in dirty state", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/dirty-down")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeTrue())

				err = migrate.Down()
				Expect(err).NotTo(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeTrue())

				err = migrate.Reset()
				Expect(err).To(BeNil())
				Expect(pgConn.IsTableExists("tests")).To(BeFalse())
			})
		})

		Describe("Version", func() {
			It("should return Error when no migrations have been executed", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				_, _, err = migrate.Version()
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("no migration"))
			})

			It("should return latest version after up migration", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())
				version, dirty, err := migrate.Version()
				Expect(err).To(BeNil())
				Expect(version).To(Equal(uint(1)))
				Expect(dirty).To(BeFalse())
			})

			It("should return current version after steps down migration", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/step-2")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())
				version, _, err := migrate.Version()
				Expect(err).To(BeNil())
				Expect(version).To(Equal(uint(2)))

				err = migrate.Steps(-1)
				Expect(err).To(BeNil())
				version, _, err = migrate.Version()
				Expect(err).To(BeNil())
				Expect(version).To(Equal(uint(1)))
			})

			It("should return no migration error after down migration", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())

				err = migrate.Down()
				Expect(err).To(BeNil())
				_, _, err = migrate.Version()
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("no migration"))
			})

			It("should return no migration error after reset", func() {
				var err error

				migrate, err := NewMigrate(config, "./test/valid")
				Expect(err).To(BeNil())

				err = migrate.Up()
				Expect(err).To(BeNil())
				version, _, err := migrate.Version()
				Expect(err).To(BeNil())
				Expect(version).To(Equal(uint(1)))

				err = migrate.Reset()
				Expect(err).To(BeNil())
				_, _, err = migrate.Version()
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("no migration"))
			})
		})
	})
})
