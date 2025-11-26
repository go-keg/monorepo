package types

import (
	"encoding/json"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-viper/mapstructure/v2"
)

type File struct {
	FileName string `json:"fileName"`
	FilePath string `json:"filePath"`
}

func MarshalFile(f File) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		marshal, _ := json.Marshal(f)
		_, _ = io.WriteString(w, string(marshal))
	})
}

func UnmarshalFile(v any) (File, error) {
	var file File
	switch v := v.(type) {
	case string:
		err := json.Unmarshal([]byte(v), &file)
		if err != nil {
			return file, err
		}
		return file, nil
	case map[string]any:
		err := mapstructure.Decode(v, &file)
		if err != nil {
			return file, err
		}
	}
	return file, nil
}

type Files []File

func MarshalFiles(f Files) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		marshal, _ := json.Marshal(f)
		_, _ = io.WriteString(w, string(marshal))
	})
}

func UnmarshalFiles(v any) (Files, error) {
	switch v := v.(type) {
	case []any:
		var files Files
		marshal, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(marshal, &files)
		if err != nil {
			return Files{}, err
		}
		return files, nil
	}
	return Files{}, nil
}
