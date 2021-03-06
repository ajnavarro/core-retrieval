package core_retrieval

import (
	"database/sql"

	"gopkg.in/src-d/core-retrieval.v0/model"
	"gopkg.in/src-d/core-retrieval.v0/repository"

	"gopkg.in/src-d/core.v0"
	"gopkg.in/src-d/framework.v0/configurable"
	"gopkg.in/src-d/framework.v0/database"
	"gopkg.in/src-d/go-billy.v3/osfs"
)

const transactionerLocalDir = "transactioner"

type containerConfig struct {
	configurable.BasicConfiguration
	RootRepositoriesDir string `default:"/tmp/root-repositories"`
}

var config = &containerConfig{}

func init() {
	configurable.InitConfig(config)
}

var container struct {
	Database            *sql.DB
	ModelMentionStore   *model.MentionStore
	RootedTransactioner repository.RootedTransactioner
}

// Database returns a sql.DB for the default database. If it is not possible to
// connect to the database, this function will panic. Multiple calls will always
// return the same instance.
func Database() *sql.DB {
	if container.Database == nil {
		container.Database = database.Must(database.Default())
	}

	return container.Database
}

// ModelMentionStore returns the default *model.ModelMentionStore, using the
// default database. If it is not possible to connect to the database, this
// function will panic. Multiple calls will always return the same instance.
func ModelMentionStore() *model.MentionStore {
	if container.ModelMentionStore == nil {
		container.ModelMentionStore = model.NewMentionStore(Database())
	}

	return container.ModelMentionStore
}

// RootedTransactioner returns the default RootedTransactioner instance,
// using the default RootRepositories directory. The local filesystem used to
// create the transactioner is the default TemporaryFilesystem from core container.
func RootedTransactioner() repository.RootedTransactioner {
	if container.RootedTransactioner == nil {
		tmpFs, err := core.TemporaryFilesystem().Chroot(transactionerLocalDir)
		if err != nil {
			panic(err)
		}

		container.RootedTransactioner =
			repository.NewSivaRootedTransactioner(
				osfs.New(config.RootRepositoriesDir),
				tmpFs,
			)
	}

	return container.RootedTransactioner
}
