package model

// ConfigJSON is a reflection of config.json
type ConfigJSON struct {
	AutoSave                 bool
	WindowWidth              int
	WindowHeight             int
	Simulation               string
	OutputSomaAPFastFiles    string
	DataPath                 string
	OutputPoissonFiles       string
	OutputSynapseSurgeFiles  string
	OutputSynapseSpikeFiles  string
	Scroll                   float64
	DataOutputPath           string
	RangeEnd                 int
	SourceStimulus           string
	StimExpander             int
	Duration                 int
	OutputDendriteAvgFiles   string
	RangeStart               int
	OutputSynapseWeightFiles string
	OutputSomaSpikeFiles     string
	Frequency                int
	OutputSomaPSPFiles       string
	Spans                    int
	OutputSomaAPSlowFiles    string
	OutputSynapsePspFiles    string
	TimeScale                int
	OutputStimulusFiles      string
}
