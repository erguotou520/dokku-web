package routers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// GitInfo Git history info
type GitInfo struct {
	Sha     string `json:"sha"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Date    string `json:"date"`
	Avatar  string `json:"avatar"`
}

// app activities
func getAppActivities(c echo.Context) error {
	name := c.Param("name")
	pwd, _ := os.Getwd()
	err := os.Chdir("/home/dokku/" + name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Can't cd to git repo folder",
		})
	}
	out, e := exec.Command("git", "log", "--pretty=format:%h --- %ce --- %cd --- %s").Output()
	os.Chdir(pwd)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, e)
	}
	histories := make([]GitInfo, 0)
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		splited := strings.Split(line, " --- ")
		email := splited[1]
		h := md5.New()
		h.Write([]byte(email))
		cipherStr := h.Sum(nil)
		history := GitInfo{Sha: splited[0], Author: email, Date: splited[2], Message: splited[3], Avatar: "http://en.gravatar.com/avatar/" + hex.EncodeToString(cipherStr)}
		histories = append(histories, history)
	}
	return c.JSON(http.StatusOK, histories)
}
