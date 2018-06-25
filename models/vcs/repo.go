package vcs

type Repository struct {
	VCS           string
	Name          string
	NameWithOwner string
	URL           string
	Description   string
	Readme        string
	Topics        []string
	License       License
}

type License struct {
	Name        string
	Body        string
	Description string
}
