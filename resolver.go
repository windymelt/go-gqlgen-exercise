package main

import (
	"context"
	"fmt"
	"math/rand"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	ssss []*Sixsixsix
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Create666(ctx context.Context, input NewSixsixsix) (*Sixsixsix, error) {
	s := &Sixsixsix{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
	}
	r.ssss = append(r.ssss, s)
	return s, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ssss(ctx context.Context) ([]*Sixsixsix, error) {
	return r.ssss, nil
}
