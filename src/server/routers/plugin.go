package routers

import (
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

// Plugin plugin struct
type Plugin struct {
	Name          string `json:"name"`
	URL           string `yaml:"url" json:"url"`
	Author        string `json:"author"`
	Compatibility string `json:"compatibility"`
	CMD           string `yaml:"cmd" json:"cmd"`
	Link          bool   `json:"link"`
}

// InstalledPlugin installed plugin
type InstalledPlugin struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Enabled     bool   `json:"enabled"`
	Author      string `json:"author"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// Plugins plugins struct
type Plugins struct {
	Official  []Plugin `json:"official"`
	Community []Plugin `json:"community"`
	Others    []Plugin `json:"others"`
}

var (
	_plugins = Plugins{}
)

// get plugin list
func getPlugins(c echo.Context) error {
	if _plugins.Official != nil && len(_plugins.Official) > 0 {
		return c.JSON(http.StatusOK, _plugins)
	}
	b, err := ioutil.ReadFile("plugins.yml")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	plugins := Plugins{}
	e := yaml.Unmarshal(b, &plugins)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, e)
	}
	_plugins = plugins
	return c.JSON(http.StatusOK, plugins)
}

// get installed plugin list
func getInstalledPlugins(c echo.Context) error {
	out, err := exec.Command(DokkuPath, "plugin:list").Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	lines := strings.Split(strings.TrimRight(string(out), "\r\n"), "\n")
	plugins := []InstalledPlugin{}
	for index, plugin := range lines {
		if index == 0 {
			continue
		}
		tabs := regexp.MustCompile(" +").Split(plugin, 7)
		if tabs[4] == "dokku" && tabs[5] != "core" {
			plugin := InstalledPlugin{Name: tabs[1], Version: tabs[2], Enabled: tabs[3] == "enabled", Author: tabs[4], Type: tabs[5], Description: tabs[6]}
			plugins = append(plugins, plugin)
		}
	}
	return c.JSON(http.StatusOK, plugins)
}

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
	out, err := exec.Command(DokkuPath, "plugin:update", plugin).Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	done, _ := regexp.MatchString("updated$", strings.TrimRight(string(out), "\r\n"))
	return c.JSON(http.StatusOK, done)
}

// toggle plugin enabled status
func togglePluginStatus(c echo.Context) error {
	plugin := c.Param("name")
	action := c.FormValue("action")
	if action == "" || (action != "enable" && action != "disable") {
		return c.NoContent(http.StatusBadRequest)
	}
	fmt.Println("1", action, plugin)
	out, err := exec.Command(DokkuPath, "plugin:"+action, plugin).Output()
	fmt.Println("2", string(out), err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	done, _ := regexp.MatchString(action+"d$", strings.TrimRight(string(out), "\r\n"))
	return c.JSON(http.StatusOK, done)
}

// link plugin to app
// func linkPlugin(c echo.Context) error {
// }
