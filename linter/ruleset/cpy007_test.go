package ruleset_test

import (
	"testing"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/stretchr/testify/assert"

	RuleSet "github.com/cremindes/whalelint/linter/ruleset"
)

func TestValidateCpy007(t *testing.T) {
	t.Parallel()

	// nolint:gofmt,gofumpt,goimports
	testCases := []struct {
		IsViolation bool
		ExampleName string
		CopyNode    mockCopyNode
		DocsContext string
	}{
		{
			IsViolation: true,
			ExampleName: "`COPY` command with src: '.' and dest: '.' paths.",
			CopyNode: mockCopyNode{
				Original: "COPY . .",
				flags:    []string{},
				next:     []string{".", "."},
			},
			DocsContext: "{{ .CopyNode.Original }}",
		},
		{
			IsViolation: true,
			ExampleName: "`COPY` command with src: '.' and dest: 'dst/' paths.",
			CopyNode: mockCopyNode{
				Original: "COPY . dst/",
				flags:    []string{},
				next:     []string{".", "dst/"},
			},
			DocsContext: "{{ .CopyNode.Original }}",
		},
		{
			IsViolation: true,
			ExampleName: "`COPY` command with src: '.' and dest: './dst/' paths.",
			CopyNode: mockCopyNode{
				Original: "COPY . ./dst/",
				flags:    []string{},
				next:     []string{".", "./dst/"},
			},
			DocsContext: "{{ .CopyNode.Original }}",
		},
		{
			IsViolation: false,
			ExampleName: "`COPY` command with src: './src' and dest: 'dst/' paths.",
			CopyNode: mockCopyNode{
				Original: "COPY ./src dst/",
				flags:    []string{},
				next:     []string{"./src", "dst/"},
			},
			DocsContext: "{{ .CopyNode.Original }}",
		},
		{
			IsViolation: false,
			ExampleName: "`COPY` command with src: '../src' and dest: 'dst/' paths.",
			CopyNode: mockCopyNode{
				Original: "COPY ../src dst/",
				flags:    []string{},
				next:     []string{"../src", "dst/"},
			},
			DocsContext: "{{ .CopyNode.Original }}",
		},
		{
			IsViolation: true,
			ExampleName: "`COPY` command with `--chmod` flag, src: '.' and dest: '.' paths.",
			CopyNode: mockCopyNode{
				Original: "COPY --chmod=7780 . .",
				flags:    []string{"--chmod=7780"},
				next:     []string{".", "."},
			},
			DocsContext: "{{ .CopyNode.Original }}",
		},
		{
			IsViolation: true,
			ExampleName: "`COPY` command with `-chmod` flag, src: '.' and dest: '.' paths.",
			CopyNode: mockCopyNode{
				Original: "COPY -chmod=7780 . .",
				flags:    []string{},
				next:     []string{"-chmod=7780", ".", "."},
			},
			DocsContext: "{{ .CopyNode.Original }}",
		},
		{
			IsViolation: false,
			ExampleName: "`COPY` command with `-chmod` flag, src: '/src' and dest: '.' paths.",
			CopyNode: mockCopyNode{
				Original: "COPY -chmod=7780 /src .",
				flags:    []string{},
				next:     []string{"-chmod=7780", "/src", "."},
			},
			DocsContext: "{{ .CopyNode.Original }}",
		},
		{
			IsViolation: false,
			ExampleName: "`COPY` command with `--chmod` flag, src: '/src' and dest: '.' paths.",
			CopyNode: mockCopyNode{
				Original: "COPY --chmod=7780 /src .",
				flags:    []string{"--chmod=7780"},
				next:     []string{"/src", "."},
			},
			DocsContext: "{{ .CopyNode.Original }}",
		},
		{
			IsViolation: false,
			ExampleName: "`COPY` command with src: 'src.txt' and dest: '/dst' paths.",
			CopyNode: mockCopyNode{
				Original: "COPY src.txt /dst",
				flags:    []string{"--chmod=7780"},
				next:     []string{"src.txt", "/dst"},
			},
			DocsContext: "{{ .CopyNode.Original }}",
		},
	}

	RuleSet.RegisterTestCaseDocs("CPY007", testCases)

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.ExampleName, func(t *testing.T) {
			t.Parallel()

			// nolint:exhaustivestruct
			node := &parser.Node{
				Value:    "copy",
				Original: testCase.CopyNode.Original,
				Flags:    testCase.CopyNode.flags,
			}

			// fill out Next values down the tree
			nextPointer := &node.Next
			for _, next := range testCase.CopyNode.next {
				*nextPointer = &parser.Node{Value: next} // nolint:exhaustivestruct
				nextPointer = &(*nextPointer).Next
			}

			instruction, err := instructions.ParseInstruction(node)
			if err != nil {
				t.Error("cannot parse mock CopyCommand AST Node")
			}
			command, ok := instruction.(*instructions.CopyCommand)
			if !ok {
				t.Error("cannot type assert instruction to *instructions.CopyCommand")
			}

			assert.Equal(t, testCase.IsViolation, RuleSet.ValidateCpy007(command).IsViolated())
		})
	}
}
