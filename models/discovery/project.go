package discovery

type Project struct {
	Engine        string
	Name          string
	NameWithOwner string
	URL           string
	Description   string
	Readme        string
	Stars         int
	Topics        []string
	License       License
}
