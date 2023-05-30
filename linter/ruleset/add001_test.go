package ruleset_test

import (
	"testing"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/stretchr/testify/assert"

	RuleSet "github.com/cremindes/whalelint/linter/ruleset"
)

type mockAddNode struct {
	Original string
	flags    []string
	next     []string
}

func TestValidateAdd001(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		IsViolation bool
		ExampleName string
		AddNode     mockAddNode
		DocsContext string
	}{
		{
			IsViolation: true,
			ExampleName: "`ADD` command with src: 'src' and dest: 'dst/' paths.",
			AddNode: mockAddNode{
				Original: "ADD src dst/",
				flags:    []string{},
				next:     []string{"src", "dst/"},
			},
			DocsContext: "{{ .AddNode.Original }}",
		},
		{
			IsViolation: false,
			ExampleName: "`ADD` command to add a private git repository.",
			AddNode: mockAddNode{
				Original: "ADD git@git.example.com:foo/bar.git /dst",
				flags:    []string{},
				next:     []string{"git@git.example.com:foo/bar.git", "/dst"},
			},
			DocsContext: "{{ .AddNode.Original }}",
		},
		{
			IsViolation: false,
			ExampleName: "`ADD` command with remote URL.",
			AddNode: mockAddNode{
				Original: "ADD https://github.com/docker/src /dst/",
				flags:    []string{},
				next:     []string{"https://github.com/docker/src", "/dst/"},
			},
			DocsContext: "{{ .AddNode.Original }}",
		},
		{
			IsViolation: false,
			ExampleName: "`ADD` command with flag = '--checksum=...' and remote URL.",
			AddNode: mockAddNode{
				Original: "ADD --checksum=sha256:24454f830cdb571e2c4ad15481119c43b3 https://github.com/docker/src /dst/",
				flags:    []string{"--checksum=sha256:24454f830cdb571e2c4ad15481119c43b3"},
				next:     []string{"https://github.com/docker/src", "/dst/"},
			},
			DocsContext: "{{ .AddNode.Original }}",
		},
		{
			IsViolation: false,
			ExampleName: "`ADD` command with flag = '--keep-git-dir=true' and remote URL.",
			AddNode: mockAddNode{
				Original: "ADD --keep-git-dir=true https://github.com/docker/src.git#v0.10.1 /dst/",
				flags:    []string{"--keep-git-dir=true"},
				next:     []string{"https://github.com/docker/src.git#v0.10.1", "/dst/"},
			},
			DocsContext: "{{ .AddNode.Original }}",
		},
		{
			IsViolation: true,
			ExampleName: "`ADD` command with flag = '--chown=...' src: '/src' and dest: '/dst/'.",
			AddNode: mockAddNode{
				Original: "ADD --chown=55:mygroup /src /dst/",
				flags:    []string{"--chown=55:mygroup"},
				next:     []string{"/src", "/dst/"},
			},
			DocsContext: "{{ .AddNode.Original }}",
		},
		//{
		//	IsViolation: true,
		//	ExampleName: "`ADD` command with src: 'src.txt' and dest: 'dst/' paths.",
		//	AddNode:    mockAddNode{
		//		Original: "ADD src.txt dst/",
		//		flags:    []string{},
		//		next:     []string{"src.txt", "dst/"},
		//	},
		//	DocsContext: "{{ .AddNode.Original }}",
		//},
		//{
		//	IsViolation: false,
		//	ExampleName: "`ADD` command with src: 'src.zip' and dest: 'dst/' paths.",
		//	AddNode:    mockAddNode{
		//		Original: "ADD src.zip dst/",
		//		flags:    []string{},
		//		next:     []string{"src.zip", "dst/"},
		//	},
		//	DocsContext: "{{ .AddNode.Original }}",
		//},
	}

	RuleSet.RegisterTestCaseDocs("ADD001", testCases)

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.ExampleName, func(t *testing.T) {
			t.Parallel()

			// nolint:exhaustivestruct
			node := &parser.Node{
				Value:    "add",
				Original: testCase.AddNode.Original,
				Flags:    testCase.AddNode.flags,
			}

			// fill out Next values down the tree
			nextPointer := &node.Next
			for _, next := range testCase.AddNode.next {
				*nextPointer = &parser.Node{Value: next} // nolint:exhaustivestruct
				nextPointer = &(*nextPointer).Next
			}

			instruction, err := instructions.ParseInstruction(node)
			if err != nil {
				t.Error("cannot parse mock AddCommand AST Node")
			}
			command, ok := instruction.(*instructions.AddCommand)
			if !ok {
				t.Error("cannot type assert instruction to *instructions.AddCommand")
			}

			assert.Equal(t, testCase.IsViolation, RuleSet.ValidateAdd001(command).IsViolated())
		})
	}
}
