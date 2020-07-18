package api

// IGraph is an imgui graph
type IGraph interface {
	Draw(environment IEnvironment, vertPos int)
}
