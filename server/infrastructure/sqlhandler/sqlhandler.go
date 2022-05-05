package sqlhandler

import (
    "cloud.google.com/go/datastore"
    "context"
    "os"
)

type SqlHandler struct {
    Client *datastore.Client
}

func NewSqlHandler() *SqlHandler {
    sqlHandler := new(SqlHandler)
    ctx := context.Background()
    projectID := os.Getenv("DATASTORE_PROJECT_ID")
    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		return sqlHandler
	}
    sqlHandler.Client = client
    return sqlHandler
}