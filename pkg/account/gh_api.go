package account

import (
	"context"

	"github.com/Matt-Gleich/lumber"
	"github.com/shurcooL/githubv4"
)

// Get account data from GitHub
func GetData(client *githubv4.Client) Outline {
	lumber.Info("Getting account data")
	var data Outline
	err := client.Query(context.Background(), &data, nil)
	lumber.Error(err, "Failed to get github account data")
	lumber.Success("Got account data!")
	return data
}
