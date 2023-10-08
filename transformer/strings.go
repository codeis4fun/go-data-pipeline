package transformer

import "strings"

type ToUpper struct {
	Input  string
	Output string
}

type JoinStrings struct {
	Input  []string
	Output string
}

func NewToUpper(input string) *ToUpper {
	return &ToUpper{Input: input}
}

func NewJoinStrings(inputs ...string) *JoinStrings {
	return &JoinStrings{Input: inputs}
}

func (t *ToUpper) Transform() {
	t.Output = strings.ToUpper(t.Input)
}

func (t *JoinStrings) Transform() {
	t.Output = strings.Join(t.Input, " ")
}
