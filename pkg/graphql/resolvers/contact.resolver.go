package resolvers

import (
	//"fmt"
	"graphql-api/internal/contact"

	"github.com/graphql-go/graphql"
)

func GetContactResolve(params graphql.ResolveParams) (interface{}, error) {
	// Get the query from the context

	// Update limit and offset if provided
	limit, ok := params.Args["limit"].(int)
	if !ok {
		limit = 10
	}

	offset, ok := params.Args["offset"].(int)
	if !ok {
		offset = 0
	}

	searchText, ok := params.Args["searchText"].(string)
	if !ok {
		searchText = ""
	}
	contactRepo := contact.NewContactRepo()

	// Fetch contacts from the database
	contacts, err := contactRepo.GetContactsBySearchText(searchText, limit, offset)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}
