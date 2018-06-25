package models

import (
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/toolkit/models/vcs"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
)

func ProcessRepos(repos []vcs.Repository) error {
	return DB.Transaction(func(tx *pop.Connection) error {
		for _, r := range repos {
			pretty.Println("### r.URL ->", r.URL)
			t := &Tool{}
			exists, err := tx.Where("url = ?", r.URL).Exists(t)
			if err != nil {
				return errors.WithStack(err)
			}
			if exists {
				if err := tx.Where("url = ?", r.URL).First(t); err != nil {
					return errors.WithStack(err)
				}
				if err := tx.RawQuery("delete from licenses where tool_id = ?", t.ID).Exec(); err != nil {
					return errors.WithStack(err)
				}
			}
			t.Name = r.Name
			t.NameWithOwner = r.NameWithOwner
			t.URL = r.URL
			if r.Description != "" {
				t.Description = nulls.NewString(r.Description)
			}
			if r.Readme != "" {
				t.Readme = nulls.NewString(r.Readme)
			}
			t.Topics = slices.String(r.Topics)
			rl := r.License
			t.License = License{
				Name: rl.Name,
				Body: nulls.String{
					String: rl.Body,
					Valid:  rl.Body != "",
				},
				Description: nulls.String{
					String: rl.Description,
					Valid:  rl.Description != "",
				},
			}
			if exists {
				if err := tx.Eager().Update(t); err != nil {
					return errors.WithStack(err)
				}
				l := t.License
				l.ToolID = t.ID
				if err := tx.Create(&l); err != nil {
					return errors.WithStack(err)
				}
			} else {
				if err := tx.Eager().Create(t); err != nil {
					return errors.WithStack(err)
				}
			}
		}
		return nil
	})
}
