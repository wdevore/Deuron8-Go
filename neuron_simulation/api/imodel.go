package api

// IModel is app config data
type IModel interface {
	Data() interface{}

	Load()
	Save()

	Changed()
	Clean()

	SetActiveSynapse(id int)

	Samples() ISamples
}
