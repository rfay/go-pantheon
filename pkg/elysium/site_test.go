package elysium

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSiteList ensures SiteLists can be retrieved as expected.
func TestSiteList(t *testing.T) {
	assert := assert.New(t)
	sl := NewSiteList()
	mux.HandleFunc(sl.Path("GET", *session), func(w http.ResponseWriter, r *http.Request) {
		// Ensure a HTTP GET request was made with the proper authorization headers.
		testMethod(t, r, "GET")
		assert.Contains(r.Header.Get("Authorization"), session.Session)

		// Send JSON response back.
		contents, err := ioutil.ReadFile("test/sites.json")
		assert.NoError(err)
		w.Write(contents)
	})

	session.Request("GET", sl)

	// Ensure we got a valid response and were able to unmarshal it as expected.
	assert.Equal(len(sl.Sites), 2)
	assert.Equal(sl.Sites[0].Site.Framework, "wordpress")
	assert.Equal(sl.Sites[1].Site.Framework, "drupal")
	assert.Equal(sl.Sites[0].SiteID, "site-id1")
	assert.Equal(sl.Sites[1].SiteID, "site-id2")
	assert.Equal(sl.Sites[0].Site.Name, "sitename1")
	assert.Equal(sl.Sites[1].Site.Name, "sitename2")
}
