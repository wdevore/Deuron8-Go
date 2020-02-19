package api

// IVisStimInput is stimulus from an image
type IVisStimInput interface {
	Configure(imageFile string) error

	SetSize(string)

	EnableExpand(bool)
	SetExpand(string)
	Expand(value int) string

	EnablePreCache(bool)

	SetStimulusAt(x, y int)
	SetStimulusFast(x, y int)

	GetStimulus() []int

	// Given an int returns a binary bit slice representation
	GetStimulusComp(value int) []int
}
