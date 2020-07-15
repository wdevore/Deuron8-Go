package api

// IModel is app config data
type IModel interface {
	Data() interface{}

	Save()

	Changed()
	Clean()

	SetActiveSynapse(id int)

	Samples() ISamples
}
