//go:build ignore
// +build ignore

package main

import (
    "log"
    "runtime"
    "strings"

    "entgo.io/contrib/entgql"
    "entgo.io/ent/entc"
    "entgo.io/ent/entc/gen"
    "github.com/go-keg/keg/contrib/ent/annotations"
    enttemp "github.com/go-keg/keg/contrib/ent/template"
    "github.com/go-keg/keg/contrib/gql/hooks"
    gqltemp "github.com/go-keg/keg/contrib/gql/template"
)

func main() {
    ex, err := entgql.NewExtension(
        entgql.WithConfigPath("./gqlgen.yml"),
        entgql.WithSchemaGenerator(),
        entgql.WithSchemaPath("./ent.graphql"),
        entgql.WithSchemaHook(annotations.EnumsGQLSchemaHook, hooks.GenerateList),
        entgql.WithNodeDescriptor(true),
        entgql.WithTemplates(append(entgql.AllTemplates, entgql.WhereTemplate, entgql.NodeDescriptorTemplate, gqltemp.Template())...),
        entgql.WithWhereInputs(true),
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
