package routers

import (
	"github.com/labstack/echo"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

// const (
// 	CouchDB       = "https://github.com/dokku/dokku-couchdb.git"
// 	Elasticsearch = "https://github.com/dokku/dokku-elasticsearch.git"
// 	Graphite      = "https://github.com/dokku/dokku-graphite-grafana.git"
// 	MariaDB       = "https://github.com/dokku/dokku-mariadb.git"
// 	Memcached     = "https://github.com/dokku/dokku-memcached.git.git"
//   Mongo         = "https://github.com/dokku/dokku-mongo.git"
//   MySQL         = "https://github.com/dokku/dokku-mysql.git"
//   Nats          = "https://github.com/dokku/dokku-nats.git"
//   Postgres      = "https://github.com/dokku/dokku-postgres.git"
//   RabbitMQ      = "https://github.com/dokku/dokku-rabbitmq.git"
//   Redis         = "https://github.com/dokku/dokku-redis.git"
//   RethinkDB     = "https://github.com/dokku/dokku-rethinkdb.git"
// )

// add plugin
func addPlugin(c echo.Context) error {
	plugin := c.FormValue("name")
	if plugin == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	_, err := exec.Command(DokkuPath, "plugin:install", plugin).Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}

// remove plugin
func removePlugin(c echo.Context) error {
	plugin := c.Param("name")
	out, err := exec.Command(DokkuPath, "plugin:uninstall", plugin).Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	done, _ := regexp.MatchString("uninstalled$", strings.TrimRight(string(out), "\r\n"))
	if done {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusInternalServerError)
}

// update plugin
func updatePlugin(c echo.Context) error {
	plugin := c.Param("name")
	action := c.FormValue("action")
	out, err := exec.Command(DokkuPath, "plugin:"+action, plugin).Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	done, _ := regexp.MatchString("updated$", strings.TrimRight(string(out), "\r\n"))
	if done {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusInternalServerError)
}
