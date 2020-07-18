package api

// IBitStream is a base stream
type IBitStream interface {
	Reset()
	Output() int
	Step()
}
