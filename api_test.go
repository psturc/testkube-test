package testkube_test

import (
	"encoding/json"

	"io"

	"net/http"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

var _ = Describe("API Test", func() {

	It("There should be user entries", func() {

		resp, err := http.Get("https://jsonplaceholder.typicode.com/users")

		Expect(err).To(BeNil())

		body, err := io.ReadAll(resp.Body)

		Expect(err).To(BeNil())

		var data []map[string]interface{}

		err = json.Unmarshal(body, &data)

		Expect(err).To(BeNil())

		Expect(len(data)).To(BeNumerically(">", 0))

	})

})
