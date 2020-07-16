package api

// IGraph is an imgui graph
type IGraph interface {
	Draw(config, sim IModel, samples ISamples, vertPos int)
}
