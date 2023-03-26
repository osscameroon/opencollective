package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/osscameroon/opencollective/billing/internal/config"
	"github.com/osscameroon/opencollective/billing/internal/graphql"
)

// Contributors is the struct that will hold the Contributors response from the graphql query
type Contributors struct {
	Account struct {
		Name         string `json:"name"`
		Slug         string `json:"slug"`
		Transactions struct {
			TotalCount int `json:"totalCount"`
			Nodes      []struct {
				Type        string `json:"type"`
				FromAccount struct {
					Name string `json:"name"`
					Slug string `json:"slug"`
				} `json:"fromAccount"`
				Amount struct {
					Value    int    `json:"value"`
					Currency string `json:"currency"`
				} `json:"amount"`
				CreatedAt time.Time `json:"createdAt"`
			} `json:"nodes"`
		} `json:"transactions"`
	} `json:"account"`
}

//GetContributors returns the list of contributors
// GET /contributors
// it accepts an organization name as a query parameter
func GetContributors(c *gin.Context) {
	url := config.GetEnv().OCURL
	key := config.GetEnv().OCKey
	client := graphql.NewClient(url, key)

	query := graphql.Query(`
		query account($slug: String) {
		account(slug: $slug) {
			name
			slug
			transactions(limit: 100, type: CREDIT) {
				totalCount
				nodes {
					type
					fromAccount {
						name
						slug
					}
					amount {
						value
						currency
					}
					createdAt
					}
				}
			}
		}
	`)

	variables := map[string]interface{}{
		"slug": "osscameroon",
	}

	response := Contributors{}

	if err := client.Run(query, variables, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
