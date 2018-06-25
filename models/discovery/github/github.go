package github

import (
	"context"

	"github.com/gobuffalo/toolkit/models/discovery"
	"github.com/pkg/errors"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var _ discovery.Engine = &Github{}

type Github struct {
	client *githubv4.Client
}

func (Github) Name() string {
	return "github"
}

func (g Github) Search(ctx context.Context, topics ...string) ([]discovery.Project, error) {
	repos := []discovery.Project{}

	var q struct {
		Search search `graphql:"search(type: REPOSITORY, query: \"topic:gobuffalo\", first: 100, after: $repositoryCursor)"`
	}
	variables := map[string]interface{}{
		"repositoryCursor": (*githubv4.String)(nil), // Null after argument to get first page.
	}

	for {
		err := g.client.Query(ctx, &q, variables)
		if err != nil {
			return repos, errors.WithStack(err)
		}

		for _, e := range q.Search.Edges {
			repos = append(repos, processRepo(e.Node.Repository))
		}

		if !q.Search.PageInfo.HasNextPage {
			break
		}
		variables["repositoryCursor"] = githubv4.NewString(q.Search.PageInfo.EndCursor)
	}
	return repos, nil
}

func processRepo(r repository) discovery.Project {
	rl := r.LicenseInfo
	repo := discovery.Project{
		Engine:        "github",
		Name:          r.Name,
		NameWithOwner: r.NameWithOwner,
		URL:           r.URL,
		Description:   r.Description,
		Readme:        r.Object.Text.Text,
		Stars:         r.Stargazers.TotalCount,
		License: discovery.License{
			Name:        rl.Name,
			Body:        rl.Body,
			Description: rl.Description,
		},
		Topics: []string{},
	}
	for _, t := range r.Topics.Nodes {
		repo.Topics = append(repo.Topics, t.Topic.Name)
	}
	return repo
}

func New(ctx context.Context, token string) (*Github, error) {
	g := &Github{}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(ctx, src)
	g.client = githubv4.NewClient(httpClient)
	return g, nil
}

// var q struct {
// 	Search struct {
// 		RepositoryCount int
// 		Edges           []struct {
// 			Cursor string
// 			Node   struct {
// 				Repository struct {
// 					Name         string
// 					ResourcePath string
// 					URL          string
// 					Topics       struct {
// 						Nodes []struct {
// 							Topic struct {
// 								Name string
// 							}
// 						}
// 					} `graphql:"repositoryTopics(first: 100)"`
// 				} `graphql:"...on Repository"`
// 			}
// 		}
// 	} `graphql:"search(type: REPOSITORY, query: \"topic:gobuffalo\", first: 100)"`
// }
