package analysis

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/jedib0t/go-pretty/text"

	"github.com/jedib0t/go-pretty/table"
)

// Variable is the struct storing the variable should be stored in heap or struct
// according to the escape analysis
type Variable struct {
	Stacks []string
	Heaps  []string
}

// Analysis is the struct storing analysis results.
type Analysis struct {
	PathRgx  *regexp.Regexp
	StackRgx []*regexp.Regexp
	HeapRgx  []*regexp.Regexp
	Codes    []string
	Result   map[string]*Variable
}

// New returns Analysis initialized with path regexp.
func New() *Analysis {
	r, err := regexp.Compile(`^(.*):\d+:\d+`)
	if err != nil {
		log.Fatalf("regex failure %s\n", err)
		return nil
	}

	hr1, _ := regexp.Compile(`:[[:space:]](.*) escapes to heap$`)
	hr2, _ := regexp.Compile(`moved to heap:[[:space:]](.*)$`)
	sr1, _ := regexp.Compile(`:[[:space:]](.*) does not escape$`)

	return &Analysis{
		PathRgx:  r,
		HeapRgx:  []*regexp.Regexp{hr1, hr2},
		StackRgx: []*regexp.Regexp{sr1},
		Result:   make(map[string]*Variable),
	}
}

// Start starts parsing escape analysis of the output of go build -gcflags=-m
func (a *Analysis) Start(data string) {
	s := strings.Split(data, "\n")
	for _, val := range s {
		if a.PathRgx.MatchString(val) {
			path := a.PathRgx.FindStringSubmatch(val)[0]
			for _, stck := range a.StackRgx {
				if stck.MatchString(val) {
					v, ok := a.Result[path]
					if !ok {
						v = &Variable{
							Stacks: nil,
							Heaps:  nil,
						}
						a.Codes = append(a.Codes, path)
					}
					v.Stacks = append(v.Stacks, stck.FindStringSubmatch(val)[1])
					a.Result[path] = v
					continue
				}
			}

			for _, heap := range a.HeapRgx {
				if heap.MatchString(val) {
					v, ok := a.Result[path]
					if !ok {
						v = &Variable{
							Stacks: nil,
							Heaps:  nil,
						}
						a.Codes = append(a.Codes, path)
					}
					v.Heaps = append(v.Heaps, heap.FindStringSubmatch(val)[1])
					a.Result[path] = v
					continue
				}
			}
		}
	}
}

func (a *Analysis) String() string {
	tw := table.NewWriter()
	tw.SetTitle("GOESP")
	tw.AppendHeader(table.Row{"CODE", "STACK", "HEAP"})
	var ts []table.Row
	for _, code := range a.Codes {
		var strS string
		stcks := a.Result[code].Stacks
		for _, s := range stcks {
			strS += fmt.Sprintf("%s,", s)
		}
		var heapS string
		heaps := a.Result[code].Heaps
		for _, h := range heaps {
			heapS += fmt.Sprintf("%s,", h)
		}
		ts = append(ts, table.Row{
			code, strS, heapS,
		})
	}

	tw.AppendRows(ts)

	tw.SetStyle(table.StyleBold)
	tw.Style().Title.Align = text.AlignCenter
	tw.Style().Options.SeparateRows = true

	return tw.Render()
}
