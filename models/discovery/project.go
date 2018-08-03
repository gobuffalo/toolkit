package discovery

type Project struct {
	Service       string
	Name          string
	NameWithOwner string
	URL           string
	Description   string
	Readme        string
	Stars         int
	Watchers      int
	Forks         int
	Topics        []string
	License       License
}
