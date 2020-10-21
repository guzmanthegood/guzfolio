package model

type GraphQLResponse struct {
	Errors []Errors    `json:"errors"`
	Data   interface{} `json:"data"`
}
type Errors struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}

func GraphQLError(message string, path string) GraphQLResponse {
	return GraphQLResponse{
		Errors: []Errors{{
			Message: message,
			Path:    []string{path},
		}},
		Data:   nil,
	}
}