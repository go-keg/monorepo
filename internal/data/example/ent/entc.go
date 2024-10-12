//go:build ignore
// +build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/go-keg/keg/contrib/ent/annotations"
	enttemp "github.com/go-keg/keg/contrib/ent/template"
	gqltemp "github.com/go-keg/keg/contrib/gql/template"
	"log"
	"runtime"
	"strings"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./ent.graphql"),
		entgql.WithWhereInputs(true),
		entgql.WithNodeDescriptor(true),
		entgql.WithSchemaHook(annotations.EnumsGQLSchemaHook),
		entgql.WithTemplates(gqltemp.Template()),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	_, filename, _, _ := runtime.Caller(0)
	entPath := strings.TrimSuffix(filename, "entc.go")
	if err = entc.Generate(entPath+"schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureIntercept,
			gen.FeatureSnapshot,
			gen.FeatureModifier,
			gen.FeatureExecQuery,
			gen.FeatureUpsert,
		},
	},
		entc.Extensions(ex),
		enttemp.Template(),
	); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
