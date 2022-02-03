package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"math"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/asaskevich/govalidator"
	"github.com/joho/godotenv"
)

const (
	EnvPattern       = "*.env"
	EnvExclude       = "example"
	DefaultDepthStep = 2
)

type Configuration struct {
	NeedHelp     bool
	URL          string `yaml:"db_url" json:"db_url"`
	MaxDepth     int    `yaml:"max_depth" json:"max_depth"`
	Timeout      int    `yaml:"timeout" json:"timeout"`
	DepthIncStep int    `yaml:"depth_inc_step" json:"depth_inc_step"`
	Output       string `yaml:"output" json:"output"`
}

func (c *Configuration) fillFromEnv() (err error) {
	if value := os.Getenv("URL"); value != "" {
		c.URL = value
	}

	if value := os.Getenv("OUTPUT"); value != "" {
		c.Output = value
	}

	var v int

	if value := os.Getenv("MAX_DEPTH"); value != "" {
		if v, err = strconv.Atoi(value); err != nil {
			return
		}

		c.MaxDepth = v
	}

	if value := os.Getenv("DEPTH_INC_STEP"); value != "" {
		if v, err = strconv.Atoi(value); err != nil {
			return
		}

		c.DepthIncStep = v
	}

	if value := os.Getenv("TIMEOUT"); value != "" {
		if v, err = strconv.Atoi(value); err != nil {
			return
		}

		c.Timeout = v
	}

	return err
}

func (c *Configuration) validate() (err error) {
	if _, err = url.ParseRequestURI(c.URL); err != nil {
		return fmt.Errorf("wrong url; %w", err)
	}

	if c.Output != "" {
		isPath, _ := govalidator.IsFilePath(c.Output)
		if !isPath {
			return errors.New("wrong output path")
		}
	}

	return
}

func (c *Configuration) loadFromEnv() (err error) {
	if err = fillEnv(); err != nil {
		return
	}

	if err = c.fillFromEnv(); err != nil {
		return
	}

	return
}

func (c *Configuration) loadFromFile(path string) (err error) {
	var content []byte

	if filepath.Ext(path) == ".yaml" {
		if content, err = os.ReadFile(path); err != nil {
			return
		}

		if err = yaml.Unmarshal(content, &c); err != nil {
			return
		}

		return
	}

	if filepath.Ext(path) == ".json" {
		if content, err = os.ReadFile(path); err != nil {
			return
		}

		if err = json.Unmarshal(content, &c); err != nil {
			return
		}

		return
	}

	return
}

func (c *Configuration) loadFromFlags(fc Configuration) {
	if fc.URL != "" {
		c.URL = fc.URL
	}

	if fc.Output != "" {
		c.Output = fc.Output
	}

	if fc.MaxDepth != 0 {
		c.MaxDepth = fc.MaxDepth
	}

	if fc.Timeout != 0 {
		c.Timeout = fc.Timeout
	}

	if fc.DepthIncStep != 0 {
		c.DepthIncStep = fc.DepthIncStep
	}

	if fc.NeedHelp {
		c.NeedHelp = fc.NeedHelp
	}
}

// Load Reads configuration from files (yaml, json), .env, environment variables or arguments and returns.
func Load() (c Configuration, err error) {
	fConf, cfgPath := readFlags()

	if err = c.loadFromEnv(); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return
		}
	}

	if err = c.loadFromFile(cfgPath); err != nil {
		return
	}

	c.loadFromFlags(fConf)

	if err = c.validate(); err != nil {
		return
	}

	return
}

func readFlags() (c Configuration, path string) {
	flag.StringVar(&c.URL, "u", "", "URL")
	flag.IntVar(&c.MaxDepth, "m", math.MaxInt, "Max depth")
	flag.IntVar(&c.Timeout, "t", 0, "Timeout")
	flag.IntVar(&c.DepthIncStep, "s", DefaultDepthStep, "Depth step")
	flag.StringVar(&c.Output, "o", "", "Output - file path or empty for log in console")
	flag.BoolVar(&c.NeedHelp, "h", false, "Print help")
	flag.StringVar(&path, "c", "", "Config file path")
	flag.Parse()

	return
}

func fillEnv() (err error) {
	var matches []string

	if matches, err = filepath.Glob(EnvPattern); err != nil {
		return
	}

	for i, fileName := range matches {
		if strings.Contains(fileName, EnvExclude) {
			matches = remove(matches, i)
		}
	}

	return godotenv.Overload(matches...)
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

// ShowHelp - prints app help info
func ShowHelp() {
	fmt.Println(`
About:
Crawler is a program for find titles in URLs.

Examples of using:`)
}
