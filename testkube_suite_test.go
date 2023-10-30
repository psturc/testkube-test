package testkube_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestkube(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testkube Suite")
}
