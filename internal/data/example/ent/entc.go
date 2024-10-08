//go:build ignore
// +build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/go-keg/keg/contrib/ent/annotations"
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
		entgql.WithNodeDescriptor(false),
		entgql.WithSchemaHook(annotations.EnumsGQLSchemaHook),
		entgql.WithTemplates(gqltemp.Template()),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	_, filename, _, _ := runtime.Caller(0)
	entPath := strings.TrimSuffix(filename, "ent/entc.go")
	if err = entc.Generate(entPath+"ent/schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureIntercept,
			gen.FeatureSnapshot,
			gen.FeatureModifier,
			gen.FeatureExecQuery,
		},
	},
		entc.Extensions(ex),
		entc.TemplateDir(entPath+"../template"),
	); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
