package pipeline

import (
	"fmt"

	"github.com/codeis4fun/go-data-pipeline/transformer"
)

type stage struct {
	name        string
	transformer transformer.Transformer
}

type pipeline struct {
	name   string
	stages []stage
}

func NewPipeline(name string) *pipeline {
	return &pipeline{
		name:   name,
		stages: []stage{},
	}
}

func (p *pipeline) AddStage(name string, transformation transformer.Transformer) *pipeline {
	p.stages = append(p.stages, stage{name: name, transformer: transformation})
	return p
}
func (p *pipeline) Run() {
	fmt.Println("Running pipeline:", p.name)
	for _, s := range p.stages {
		fmt.Println("Running stage:", s.name)
		s.transformer.Transform()
	}
	fmt.Print("Pipeline finished\n\n")
}
