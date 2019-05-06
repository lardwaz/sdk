package resolver_test

import (
	"testing"

	"go.lsl.digital/lardwaz/sdk/resolver"
)

type resolvable struct {
	single  bool
	listing bool
	create  bool
	update  bool
	delete  bool
}

func (r resolvable) same(r2 resolver.Resolvable) bool {
	if r.single != r2.SingleResolver() {
		return false
	}

	if r.listing != r2.ListingResolver() {
		return false
	}

	if r.create != r2.CreateResolver() {
		return false
	}

	if r.update != r2.UpdateResolver() {
		return false
	}

	if r.delete != r2.DeleteResolver() {
		return false
	}

	return true
}

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		resolvers uint8
		want      resolvable
	}{
		{
			name:      "Only Single",
			resolvers: resolver.Single,
			want: resolvable{
				single: true,
			},
		},
		{
			name:      "Only Listing",
			resolvers: resolver.Listing,
			want: resolvable{
				listing: true,
			},
		},
		{
			name:      "Only Create",
			resolvers: resolver.Create,
			want: resolvable{
				create: true,
			},
		},
		{
			name:      "Only Update",
			resolvers: resolver.Update,
			want: resolvable{
				update: true,
			},
		},
		{
			name:      "Only Delete",
			resolvers: resolver.Delete,
			want: resolvable{
				delete: true,
			},
		},
		{
			name:      "Single Listing",
			resolvers: resolver.Single | resolver.Listing,
			want: resolvable{
				single:  true,
				listing: true,
			},
		},
		{
			name:      "Create Update Delete",
			resolvers: resolver.Create | resolver.Update | resolver.Delete,
			want: resolvable{
				create: true,
				update: true,
				delete: true,
			},
		},
		{
			name:      "ALL",
			resolvers: resolver.ALL,
			want: resolvable{
				single:  true,
				listing: true,
				create:  true,
				update:  true,
				delete:  true,
			},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolver.New(tt.name, tt.resolvers); !tt.want.same(got) || tt.name != got.ResolverSummary() {
				t.Errorf("%d: New() = %v, want %v", i, got, tt.want)
			}
		})
	}
}
