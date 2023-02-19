# Sample

## Project Layout
```
.
├── build                        docker configurations
├── cmd                          main application of the project
│     └── Sample      the API server application
├── deployments                  docker-compose and other deploy configurations
│     └── data                      directory for pg persist volume
├── internal                     private application and library code
│     ├── application               application instance code
│     ├── config                    configuration library
│     ├── db                        db access logic
│     ├── entity                    entity definitions and domain logic
│     ├── errors                    error types and handling
│     ├── exithandler               handler for server stop events
│     ├── handler                   API handlers logic
│     ├── logger                    logger
│     ├── middleware                mux middleware logic
│     ├── repository                repository logic for data access
│     ├── router                    API routes
│     ├── server                    mux server implementation
│     ├── service                   business logic layer
│     ├── test                      helpers for testing (?)
│     └── utils                     useful utils
├── migrations                   database migrations
├── pkg                          public library code
│     └── pagination                paginated list
└── testdata                     test data scripts
```

### Updating Database Schema

```shell
# Execute new migrations made by you or other team members.
# Usually you should run this command each time after you pull new code from the code repo. 
make migrate

# Create a new database migration.
# In the generated `migrations/*.up.sql` file, write the SQL statements that implement the schema changes.
# In the `*.down.sql` file, write the SQL statements that revert the schema changes.
make migrate-new

# Revert the last database migration.
# This is often used when a migration has some issues and needs to be reverted.
make migrate-down

# Clean up the database and rerun the migrations from the very beginning.
# Note that this command will first erase all data and tables in the database, and then
# run all migrations. 
make migrate-reset
```

### Managing Configurations

The application configuration is represented in `internal/config/config.go`. Configurations
specified in environment variables should be named with the `APP_` prefix and in upper case. When a configuration
is specified in both a configuration file and an environment variable, the latter takes precedence.
