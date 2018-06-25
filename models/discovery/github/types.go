package github

import "github.com/shurcooL/githubv4"

type search struct {
	RepositoryCount int
	PageInfo        pageInfo
	Edges           edges
}

type edges []struct {
	Cursor string
	Node   repoNode
}

type pageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
}

type topic struct {
	Name string
}

type topicNodes []struct {
	Topic topic
}

type topics struct {
	Nodes topicNodes
}

type license struct {
	Name        string
	Body        string
	Description string
}

type repository struct {
	Name          string
	NameWithOwner string
	ResourcePath  string
	URL           string
	Description   string
	LicenseInfo   license
	Stargazers    stars
	Topics        topics `graphql:"repositoryTopics(first: 100)"`
	Object        object `graphql:"object(expression: \"master:README.md\")"`
}

type stars struct {
	TotalCount int
}

type object struct {
	Text struct {
		Text string
	} `graphql:"... on Blob"`
}

type repoNode struct {
	Repository repository `graphql:"...on Repository"`
}
