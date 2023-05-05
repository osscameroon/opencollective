package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osscameroon/opencollective/billing/internal/config"
	"github.com/osscameroon/opencollective/billing/internal/graphql"
)

// PostTierResponse
type PostTierResponse struct {
	CreateTier struct {
		ID          string `json:"id"`
		LegacyID    int    `json:"legacyId"`
		Slug        string `json:"slug"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Amount      struct {
			Value        float64 `json:"value"`
			Currency     string  `json:"currency"`
			ValueInCents int     `json:"valueInCents"`
		} `json:"amount"`
		Button interface{} `json:"button"`
		Goal   struct {
			Value        interface{} `json:"value"`
			Currency     string      `json:"currency"`
			ValueInCents interface{} `json:"valueInCents"`
		} `json:"goal"`
		Type              string      `json:"type"`
		Interval          string      `json:"interval"`
		Frequency         string      `json:"frequency"`
		Presets           interface{} `json:"presets"`
		MaxQuantity       int         `json:"maxQuantity"`
		AvailableQuantity int         `json:"availableQuantity"`
		CustomFields      interface{} `json:"customFields"`
		AmountType        string      `json:"amountType"`
		MinimumAmount     struct {
			Value        interface{} `json:"value"`
			Currency     string      `json:"currency"`
			ValueInCents interface{} `json:"valueInCents"`
		} `json:"minimumAmount"`
		EndsAt            interface{} `json:"endsAt"`
		InvoiceTemplate   string      `json:"invoiceTemplate"`
		UseStandalonePage bool        `json:"useStandalonePage"`
		SingleTicket      bool        `json:"singleTicket"`
	} `json:"createTier"`
}

// PostTier creates a new tier
// POST /tier
// it accepts an organization name as a query parameter
func PostTier(c *gin.Context) {
	url := config.GetEnv().OCURL
	key := config.GetEnv().OCKey
	client := graphql.NewClient(url, key)

	query := graphql.Query(`
	mutation (
  	  $tier: TierCreateInput!
  	  $account: AccountReferenceInput!
	) {
  	  createTier(tier: $tier, account: $account) {
    	id
    	legacyId
    	slug
    	name
    	description
   	   amount {
      	  value
      	  currency
      	  valueInCents
    	}
    	button
    	goal {
      	  value
      	  currency
      	  valueInCents
    	}
    	type
    	interval
    	frequency
    	presets
    	maxQuantity
    	availableQuantity
    	customFields
    	amountType
    	minimumAmount {
      	  value
      	  currency
      	  valueInCents
    	}
    	endsAt
    	invoiceTemplate
    	useStandalonePage
    	singleTicket
  	  }
	}
`)

	variables := map[string]interface{}{
		"tier": map[string]interface{}{
			"amount": map[string]interface{}{
				"value":        30.7,
				"currency":     "USD",
				"valueInCents": 42,
			},
			"name":        "non-empty-string",
			"description": "A description",
			"type":        "TIER",
			"amountType":  "FIXED",
			"frequency":   "MONTHLY",
			"presets": []interface{}{
				42,
			},
			"maxQuantity": 42,
			"minimumAmount": map[string]interface{}{
				"value":        30.7,
				"currency":     "USD",
				"valueInCents": 42,
			},
			"useStandalonePage": true,
			"invoiceTemplate":   "invoiceTemplate",
			"singleTicket":      true,
		},
		"account": map[string]interface{}{
			"id":   "4rxg0j35-lzkwm6vz-bxbqvoe9-8n47daby",
			"slug": "osscameroon",
		},
	}

	response := PostTierResponse{}

	if err := client.Run(query, variables, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTierResponse
type DeleteTierResponse struct {
	DeleteTier struct {
		ID          string `json:"id"`
		LegacyID    int    `json:"legacyId"`
		Slug        string `json:"slug"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Amount      struct {
			Value        float64 `json:"value"`
			Currency     string  `json:"currency"`
			ValueInCents int     `json:"valueInCents"`
		} `json:"amount"`
		Button interface{} `json:"button"`
		Goal   struct {
			Value        interface{} `json:"value"`
			Currency     string      `json:"currency"`
			ValueInCents interface{} `json:"valueInCents"`
		} `json:"goal"`
		Type              string      `json:"type"`
		Interval          string      `json:"interval"`
		Frequency         string      `json:"frequency"`
		Presets           interface{} `json:"presets"`
		MaxQuantity       int         `json:"maxQuantity"`
		AvailableQuantity interface{} `json:"availableQuantity"`
		CustomFields      interface{} `json:"customFields"`
		AmountType        string      `json:"amountType"`
		MinimumAmount     struct {
			Value        interface{} `json:"value"`
			Currency     string      `json:"currency"`
			ValueInCents interface{} `json:"valueInCents"`
		} `json:"minimumAmount"`
		EndsAt            interface{} `json:"endsAt"`
		InvoiceTemplate   string      `json:"invoiceTemplate"`
		UseStandalonePage bool        `json:"useStandalonePage"`
		SingleTicket      bool        `json:"singleTicket"`
	} `json:"deleteTier"`
}

// PostTier creates a new tier
// DELETE /tier
// it accepts an organization name as a query parameter
func DeleteTier(c *gin.Context) {
	url := config.GetEnv().OCURL
	key := config.GetEnv().OCKey
	client := graphql.NewClient(url, key)

	query := graphql.Query(`
	mutation (
  	  $tier: TierReferenceInput!
	) {
  	  deleteTier(
    	tier: $tier
  	  ) {
    	id
    	legacyId
    	slug
    	name
    	description
    	amount {
      	  value
      	  currency
      	  valueInCents
    	}
    	button
    	goal {
      	  value
      	  currency
      	  valueInCents
    	}
    	type
    	interval
    	frequency
    	presets
    	maxQuantity
    	availableQuantity
    	customFields
    	amountType
    	minimumAmount {
      	  value
      	  currency
      	  valueInCents
    	}
    	endsAt
    	invoiceTemplate
    	useStandalonePage
    	singleTicket
  	  }
	}
`)

	variables := map[string]interface{}{
		"tier": map[string]interface{}{
			//Note: this is the hardcoded value of the tier to be deleted
			"id": "53kzxy4v-07wlr6m9-vkmpmj9n-o8agdbe5",
		},
	}

	response := DeleteTierResponse{}

	if err := client.Run(query, variables, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
