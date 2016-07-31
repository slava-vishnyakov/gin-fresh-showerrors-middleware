package showerrors

import (
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

// FromFresh is middleware to show errors on `fresh`` builds
func FromFresh(c *gin.Context) {
	b, err := ioutil.ReadFile("tmp/runner-build-errors.log")
	if err == nil && len(b) > 0 {
		c.String(500, "INTERNAL SERVER ERROR:\n\n"+string(b))
		c.Abort()
		return
	}
	c.Next()
}

// FromFreshAndGoGet is middleware to show errors on `fresh`` builds
func FromFreshAndGoGet(c *gin.Context) {
	b, err := ioutil.ReadFile("tmp/runner-build-errors.log")
	if err == nil && len(b) > 0 {
		if strings.Contains(string(b), "cannot find package") {
			b2, err := exec.Command("go", "get").Output()
			if err != nil {
				c.String(500, "INTERNAL SERVER ERROR on `go get`:\n\n"+err.Error())
				c.Abort()
				return
			}

			c.String(500, "You were missing a package, I installed it, change something to rebuild:\n\n"+string(b)+"\n\n"+string(b2))
		} else {
			c.String(500, "INTERNAL SERVER ERROR:\n\n"+string(b))
		}
		c.Abort()

		return
	}
	c.Next()
}
