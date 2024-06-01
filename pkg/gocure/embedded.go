package gocure

func (e *Embedded) AddToStep() (err error) {
	return e.Config.AddEmbeddedFileToStep()
}

func (e *Embedded) AddToScenario() (err error) {
	return e.Config.AddEmbeddedFileToScenario()
}

func (e *Embedded) AddToFeature() (err error) {
	return e.Config.AddEmbeddedFileToFeature()
}

func (e *Embedded) Add() (embeddedModel interface{}, err error) {
	return e.Config.AddEmbeddedFile()
}
