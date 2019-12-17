package webhook

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

var HookFile string

type Hooks struct {
	Entries []struct {
		Param   string `yaml:"param"`
		Command string `yaml:"command"`
	} `yaml:"hooks"`
}

func Exec(param string) {
	hs := loadConf(HookFile)
	for _, hook := range hs.Entries {
		if hook.Param == param {
			cmd := exec.Command("bash", "-c", hook.Command)
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				logrus.Infof("command %s execute failed", hook.Command)
			}
			logrus.Infof("command %s executed", hook.Command)
		}
	}

}

func loadConf(file string) *Hooks {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Fatalf("open webhook yaml failed .")
	}
	hs := Hooks{}
	if err := yaml.Unmarshal(bs, &hs); err != nil {
		logrus.Fatalf("unmarshal webhook failed .")
	}
	return &hs
}
