// Code generated by fabricator-generate-plugin-go
//
// Modifications in code regions will be lost during regeneration!

package fabricatorgenerateprojectgo_test

import (
	"os"

	fabricatorgenerateprojectgo "code.cestus.io/tools/fabricator-generate-project-go/pkg/fabricator-generate-project-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	It("Deserializes a fabricator config", func() {
		file, err := os.Open("./testdata/deserialize.yml")
		Expect(err).ToNot(HaveOccurred())
		config, err := fabricatorgenerateprojectgo.LoadPluginConfig(file)
		Expect(err).ToNot(HaveOccurred())
		expected := fabricatorgenerateprojectgo.PluginConfig{
			ApiVersion: "fabricator.cestus.io/v1alpha1",
			Kind:       "Config",
			Components: []fabricatorgenerateprojectgo.PluginComponent{
				{
					Name:      "fabricator-generate-project-go",
					Generator: "fabricator-generate-project-go",
					Spec: fabricatorgenerateprojectgo.Spec{
						GoModule: "code.cestus.io/testvalue",
						Minimal:  false,
						RepoURL:  "code.cestus.io/repo",
					},
				},
			},
		}
		Expect(config).To(Equal(expected))
	})
})
