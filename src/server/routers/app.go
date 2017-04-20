package routers

import (
	"github.com/labstack/echo"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

// Env env model
type Env struct {
	Key   string `json:"key" form:"key"`
	Value string `json:"value" form:"value"`
}

func (env Env) toString() string {
	return env.Key + "=" + env.Value
}

// EnvForm env form model
type EnvForm struct {
	Name   string   `json:"name" form:"name"`
	Envs   []Env    `json:"envs" form:"envs"`
	Remove []string `json:"remove" form:"remove"`
}

// App app list struct
type App struct {
	Name   string `json:"name"`
	Dir    string `json:"dir"`
	Sha    string `json:"sha"`
	CID    string `json:"cid"`
	Status string `json:"status"`
}

// app list
func appList(c echo.Context) error {
	out, err := exec.Command(DokkuPath, "apps:report").Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if len(out) == 0 {
		return c.JSON(http.StatusOK, []string{})
	}
	apps := []App{}
	l := regexp.MustCompile(`\n`)
	n := regexp.MustCompile(`=====> `)
	d := regexp.MustCompile(`App dir:`)
	g := regexp.MustCompile(`Git sha:`)
	cid := regexp.MustCompile(`App cid:`)
	s := regexp.MustCompile(`Status:`)
	nd := regexp.MustCompile(`not deployed`)
	lines := l.Split(string(out), -1)
	for _, line := range lines {
		if n.MatchString(line) {
			app := new(App)
			app.Name = n.ReplaceAllString(line, "")
			apps = append(apps, *app)
		} else {
			app := &apps[len(apps)-1]
			if d.MatchString(line) {
				app.Dir = strings.TrimSpace(d.ReplaceAllString(line, ""))
			} else if g.MatchString(line) {
				app.Sha = strings.TrimSpace(g.ReplaceAllString(line, ""))
			} else if cid.MatchString(line) {
				app.CID = strings.TrimSpace(cid.ReplaceAllString(line, ""))
			} else if s.MatchString(line) {
				app.Status = strings.TrimSpace(s.ReplaceAllString(line, ""))
			} else if nd.MatchString(line) {
				app.Status = `not deployed`
			}
		}
	}
	return c.JSON(http.StatusOK, apps)
}

// create app
func createApp(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	out, err := exec.Command(DokkuPath, "apps:create", name).Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	done, _ := regexp.MatchString("done[ ]?$", strings.TrimRight(string(out), "\r\n"))
	if done {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusInternalServerError)
}

// destroy app
func destroyApp(c echo.Context) error {
	name := c.Param("name")
	_, err := exec.Command(DokkuPath, "apps:destroy", "--force", name).Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}

// rename app
func renameApp(c echo.Context) error {
	name := c.Param("name")
	newName := c.FormValue("newName")
	if newName == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	deployed, e := exec.Command(DokkuPath, "apps:report", name, "--status").Output()
	if e != nil {
		return c.JSON(http.StatusInternalServerError, e)
	}
	if string(deployed) == "not deployed" {
		return c.JSON(http.StatusForbidden, "App not deployed.")
	}
	out, err := exec.Command(DokkuPath, "apps:rename", name, newName).Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	done, _ := regexp.MatchString("done$", strings.TrimRight(string(out), "\r\n"))
	if done {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusInternalServerError)
}

// start an app
func actionApp(c echo.Context) error {
	name := c.Param("name")
	action := c.FormValue("action")
	out, err := exec.Command(DokkuPath, "ps:"+action, name).Output()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	done, _ := regexp.MatchString("done$", strings.TrimRight(string(out), "\r\n"))
	if done {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(http.StatusInternalServerError, map[string]string{
		"message": string(out),
	})
}

// set env variables
func configEnv(c echo.Context) error {
	form := new(EnvForm)
	if err := c.Bind(form); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if form.Name == "" || (len(form.Envs) == 0 && len(form.Remove) == 0) {
		return c.NoContent(http.StatusBadRequest)
	}
	// set envs
	if len(form.Envs) > 0 {
		str := " "
		for _, env := range form.Envs {
			str = str + env.toString() + " "
		}
		_, e := exec.Command(DokkuPath, "config:set", form.Name, str, "--no-restart").Output()
		if e != nil {
			return c.JSON(http.StatusInternalServerError, e)
		}
	}
	// unset envs
	if len(form.Remove) > 0 {
		_, er := exec.Command(DokkuPath, "config:unset", form.Name, strings.Join(form.Remove, " "), "--no-restart").Output()
		if er != nil {
			return c.JSON(http.StatusInternalServerError, er)
		}
	}
	return c.NoContent(http.StatusOK)
}
