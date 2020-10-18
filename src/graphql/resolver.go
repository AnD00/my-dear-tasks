package graphql

import (
	"context"

	db "github.com/AnD00/my-dear-tasks/src/databases"
	"github.com/AnD00/my-dear-tasks/src/enum"
	"github.com/AnD00/my-dear-tasks/src/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

/*
mutation createTask {
createTask(input:{text:"graphql test"}) {
text
status
}
}
*/
func (r *mutationResolver) CreateTask(ctx context.Context, input NewTask) (*models.Task, error) {
	task := &models.Task{Text: input.Text, Status: enum.OPEN}
	db.Get().Create(task)
	return task, nil
}

type queryResolver struct{ *Resolver }

/*
query tasks {
tasks {
id
text
status
}
}
*/
// $ curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"# Write your query or mutation here\nquery tasks {\n  tasks {\n    id\n    text\n    status\n  }\n}\n"}' --compressed | json_pp
func (r *queryResolver) Tasks(ctx context.Context) ([]*models.Task, error) {
	var tasks []models.Task
	db.Get().Find(&tasks)
	response := make([]*models.Task, 0, len(tasks))

	for _, task := range tasks {
		response = append(response,
			&models.Task{
				ID:        task.ID,
				CreatedAt: task.CreatedAt,
				UpdatedAt: task.UpdatedAt,
				DeletedAt: task.DeletedAt,
				Text:      task.Text,
				Status:    task.Status,
			})

	}
	return response, nil
}
