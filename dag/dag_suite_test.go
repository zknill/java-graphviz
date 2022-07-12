package dag_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDag(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dag Suite")
}
