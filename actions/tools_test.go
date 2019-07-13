package actions

import (
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/toolkit/models"
	"github.com/markbates/going/randx"
)

func (as *ActionSuite) BuildTool() *models.Tool {
	n := randx.String(20)
	t := &models.Tool{
		Name:             n,
		NameWithOwner:    "gobuffalo/" + n,
		URL:              "https://github.com/gobuffalo/" + n,
		DiscoveryService: "github",
		Stars:            3000,
		License: models.License{
			Name:        "MIT",
			Body:        nulls.NewString("body"),
			Description: nulls.NewString("description"),
		},
	}
	return t
}

func (as *ActionSuite) CreateTool() *models.Tool {
	t := as.BuildTool()
	as.NoError(as.DB.Eager().Create(t))
	return t
}

func (as *ActionSuite) Test_ToolsResource_List() {
	t1 := as.CreateTool()
	t2 := as.CreateTool()

	res := as.HTML("/tools").Get()
	as.Equal(200, res.Code)

	body := res.Body.String()
	as.Contains(body, t1.NameWithOwner)
	as.Contains(body, t2.NameWithOwner)
}

func (as *ActionSuite) Test_ToolsResource_Show() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_ToolsResource_New() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_ToolsResource_Create() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_ToolsResource_Edit() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_ToolsResource_Update() {
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_ToolsResource_Destroy() {
	as.Fail("Not Implemented!")
}
