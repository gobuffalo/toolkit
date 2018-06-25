package actions

import (
	"context"
	"time"

	"github.com/gobuffalo/buffalo/worker"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/toolkit/models"
	"github.com/gobuffalo/toolkit/models/discovery/github"
	"github.com/pkg/errors"
)

func addRepoWorkers(ctx context.Context, w worker.Worker) error {
	if err := w.Register("findGitHubRepos", findGitHubRepos(ctx, w)); err != nil {
		return errors.WithStack(err)
	}
	if err := w.Perform(githubJob); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

var githubJob = worker.Job{
	Queue:   "default",
	Handler: "findGitHubRepos",
	Args:    worker.Args{},
}

func findGitHubRepos(ctx context.Context, w worker.Worker) worker.Handler {
	return func(args worker.Args) error {
		defer w.PerformIn(githubJob, 60*time.Minute)
		token, err := envy.MustGet("GITHUB_TOKEN")
		if err != nil {
			return errors.WithStack(err)
		}

		g, err := github.New(ctx, token)
		if err != nil {
			return errors.WithStack(err)
		}
		repos, err := g.Search(ctx, "gobuffalo")
		if err != nil {
			return errors.WithStack(err)
		}

		return models.ProcessProjects(repos)
	}
}
