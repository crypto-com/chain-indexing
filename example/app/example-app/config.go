package main

type CustomConfig struct {
	ServerGithubAPI ServerGithubAPIConfig `toml:"server_github_api"`
}

type ServerGithubAPIConfig struct {
	MigrationRepoRef string `toml:"migration_repo_ref"`
}
