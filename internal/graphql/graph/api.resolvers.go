package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/legion-zver/vss-brain-search/internal/graphql/graph/model"
	"github.com/mitchellh/mapstructure"
)

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, query string, useNlp *bool, isActive *bool) ([]*model.SearchResultObject, error) {
	resp, err := r.SearchEngine.Search(ctx, query, useNlp, isActive)
	if err != nil {
		return nil, err
	}
	result := make([]*model.SearchResultObject, len(resp.Hits), len(resp.Hits))
	for i, hit := range resp.Hits {
		obj := &model.SearchResultObject{
			ID:    hit.ID,
			Score: hit.Score,
		}
		_ = mapstructure.WeakDecode(hit.Fields, obj)
		switch obj.Service {
		case "premier.one":
			var url string
			if obj.Slug != nil {
				url = fmt.Sprintf("https://premier.one/show/%s", *obj.Slug)
			} else {
				url = fmt.Sprintf("https://premier.one/show/%s", obj.ID)
			}
			obj.URL = &url
		}
		result[i] = obj
	}
	return result, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
