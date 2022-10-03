package models

import (
	"github.com/wilburt/wilburx9.dev/backend/api/internal"
	"github.com/wilburt/wilburx9.dev/backend/api/internal/database"
	"time"
)

// GitHub is a container foe GitHub response data
type GitHub struct {
	Data data `json:"data"`
}

type data struct {
	Viewer viewer `json:"viewer"`
}

type viewer struct {
	Repositories repositories `json:"repositories"`
}

type repositories struct {
	Nodes []nodeElement `json:"nodes"`
}

type nodeElement struct {
	CreatedAt      string       `json:"createdAt"`
	Description    *string      `json:"description"`
	ForkCount      int          `json:"forkCount"`
	ID             string       `json:"id"`
	LicenseInfo    *licenseInfo `json:"licenseInfo"`
	Name           string       `json:"name"`
	StargazerCount int          `json:"stargazerCount"`
	URL            string       `json:"url"`
	UpdatedAt      string       `json:"updatedAt"`
	Languages      languages    `json:"languages"`
}

type languages struct {
	Edges []edge `json:"edges"`
}

type edge struct {
	Node edgeNode `json:"node"`
}

type edgeNode struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type licenseInfo struct {
	Name string `json:"name"`
}

// ToResult creates RepoResult by mapping GitHub to a slice of Repo
func (m GitHub) ToResult(source string) []database.Model {
	nodes := m.Data.Viewer.Repositories.Nodes
	var repos = make([]database.Model, len(nodes))

	mapLanguages := func(node nodeElement) []language {
		edges := node.Languages.Edges
		var languages = make([]language, len(edges))
		for j, edge := range edges {
			languages[j] = language{
				Name:  edge.Node.Name,
				Color: edge.Node.Color,
			}
		}
		return languages
	}

	getLicense := func(node nodeElement) string {
		license := node.LicenseInfo
		if license == nil {
			return ""
		} else {
			return license.Name
		}
	}

	for i, node := range nodes {
		repos[i] = Repo{
			ID:          internal.MakeId(source, node.ID),
			Name:        node.Name,
			Stars:       node.StargazerCount,
			Forks:       node.ForkCount,
			Url:         node.URL,
			Description: node.Description,
			CreatedOn:   internal.StringToTime(time.RFC3339, node.CreatedAt, "Git"),
			UpdatedOn:   internal.StringToTime(time.RFC3339, node.UpdatedAt, "Git"),
			License:     getLicense(node),
			Languages:   mapLanguages(node),
			Score:       node.StargazerCount + node.ForkCount,
			Source:      source,
		}
	}
	return repos
}
