package api

// IModel is app config data
type IModel interface {
	Data() interface{}

	SetActiveSynapse(id int)

	Samples() ISamples
}
