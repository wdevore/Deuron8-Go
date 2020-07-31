package main

type Preset struct {
	Enabled bool
	Center  float64
	Min     float64
	Max     float64
}

type Presets struct {
	TaoP             Preset
	TaoN             Preset
	Mu               Preset
	Distance         Preset
	Lambda           Preset
	Amb              Preset
	W                Preset
	Alpha            Preset
	LearningRateSlow Preset
	LearningRateFast Preset
	TaoI             Preset
	Ama              Preset
}

// ConfigJSON is a reflection of config.json
type ConfigJSON struct {
	OutputFile   string
	TargetPath   string
	SynapseCount int
	Presets      Presets
}
