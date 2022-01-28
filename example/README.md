# chain-indexing-example-app

Example implementation for [crypto.com/chain-indexing](https://github.com/crypto-com/chain-indexing)

## How to Run

### 1. Prerequisite

- Crypto.org Chain full node
- Postgres Database

### 2 Configuration file

A sample configuration is available under `config/config.toml`.

**Reminder**
Please generate your own `Personal access tokens` in Github. And then input your own username and token under `github_api.username` and `github_api.token`. This is for pulling migration files on `crypto.com/chain-indexing`.
### 3 Example file structure

```
app/
    example-app/
    
        // Custom configuartion
        config.go
        
        // Containing functions to initialize cronjobs
        cronjobs.go
        
        // App runner
        main.go
        
        // Containing functions to initialize projections
        projections.go
        
        // CLI runner
        run.go
        
        routes/
            // Implementation for `Route` interface
            routeregistry.go
            
            // Containing functions to initialize routes
            routes.go
config/

    // Configuration file
    config.yaml
    
httpapi/
    handlers/
    
        // Custom HTTP handlers
        example.go
    
internal/
    filereader/
        yaml/
        
            // YAML configuration reader
            ...
            
projections/

    // Custom example projection
    example/
        migration/
        
            // Projection database migration files (Both up and down script)
            ...
        
        view/
        
            // Projection database accessor
            examples.go
            
        // Projection implementation
        example.go
    
    // projections utility
    projection.go
```


### 4 Run the Service

#### Manual build

```bash
make build
```

#### Run binary

```bash
env DB_PASSWORD=your_postgresql_password ./example-app
```
