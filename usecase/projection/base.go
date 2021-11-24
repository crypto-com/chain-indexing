package projection

import (
	"fmt"
	"time"

	"github.com/ettle/strcase"
	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	filesystem_mh "github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper/filesystem"
	github_mh "github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper/github"
)

type Base struct {
	projectionId string

	config interface{}
}

type Options struct {
	// Pointer to the configuration
	MaybeConfigPtr interface{}
	// Pointer to MigrationHelper
	MaybeMigrationHelper migrationhelper.MigrationHelper
}

func NewBase(projectionId string) Base {
	return NewBaseWithOptions(
		projectionId,
		Options{
			MaybeConfigPtr:       nil,
			MaybeMigrationHelper: nil,
		},
	)
}

func NewBaseWithOptions(projectionId string, options Options) Base {
	base := Base{
		projectionId: projectionId,

		config: nil,
	}

	if options.MaybeConfigPtr != nil {
		decoderConfig := &mapstructure.DecoderConfig{
			WeaklyTypedInput: true,
			DecodeHook: mapstructure.ComposeDecodeHookFunc(
				mapstructure.StringToTimeDurationHookFunc(),
				mapstructure.StringToTimeHookFunc(time.RFC3339),
			),
			Result: options.MaybeConfigPtr,
		}
		decoder, decoderErr := mapstructure.NewDecoder(decoderConfig)
		if decoderErr != nil {
			panic(fmt.Errorf("error creating projection config decoder: %v", decoderErr))
		}

		projectionConfigId := strcase.ToSnake(projectionId)
		projectionConfig, projectionConfigFound := GlobalConfig.Projection[projectionConfigId]
		if !projectionConfigFound {
			panic(fmt.Sprintf("no projection config found for projection id: %s", projectionConfigId))
		}
		if err := decoder.Decode(projectionConfig); err != nil {
			panic(fmt.Errorf("error decoding projection %s config: %v", projectionId, err))
		}

		base.config = options.MaybeConfigPtr
	}

	if options.MaybeMigrationHelper != nil {

		switch migrationHelper := options.MaybeMigrationHelper.(type) {
		case *github_mh.GithubMigrationHelper:

			if migrationHelper.MaybeSourceURL == nil || migrationHelper.MaybeDatabaseURL == nil {
				setupDefaultGithubMigration(projectionId, migrationHelper)
			}

			migrationHelper.InitAndRunMigrate()

		case *filesystem_mh.FilesystemMigrationHelper:
			migrationHelper.InitAndRunMigrate()
		default:
			panic("unknown MigrationHelper type")
		}

	}

	return base
}

// Implements projection.Id()
func (base *Base) Id() string {
	return base.projectionId
}

func (base *Base) Config() interface{} {
	return base.config
}

func setupDefaultGithubMigration(
	projectionId string,
	migrationHelper *github_mh.GithubMigrationHelper,
) {
	projectionSnakeId := strcase.ToSnake(projectionId)

	if migrationHelper.MaybeSourceURL == nil {
		migrationDirectory := fmt.Sprintf(github_mh.MIGRATION_DIRECTORY_FORMAT, projectionSnakeId)

		sourceURL := github_mh.GenerateSourceURL(
			github_mh.MIGRATION_GITHUB_URL_FORMAT,
			migrationHelper.Config.GithubAPIUser,
			migrationHelper.Config.GithubAPIToken,
			migrationDirectory,
			migrationHelper.Config.MigrationRepoRef,
		)

		migrationHelper.MaybeSourceURL = &sourceURL
	}
	if migrationHelper.MaybeSourceURL == nil {
		migrationTableName := fmt.Sprintf(github_mh.MIGRATION_TABLE_NAME_FORMAT, projectionSnakeId)

		databaseURL := migrationhelper.GenerateDatabaseURL(migrationHelper.Config.ConnString, migrationTableName)

		migrationHelper.MaybeDatabaseURL = &databaseURL
	}
	return
}
