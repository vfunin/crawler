package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	"github.com/joho/godotenv"
)

type key int

const (
	LoggerCtxKey key = iota
)

const (
	EnvPattern            = "*.env"
	EnvExclude            = "example"
	DefaultDepthStep      = 2
	DefaultMaxDepth       = 2
	SkipFrameCount        = 2
	DefaultURL            = ""
	DefaultTimeout        = 10
	DefaultOutput         = ""
	DefaultNeedHelp       = false
	DefaultJSONLog        = false
	DefaultWithPanic      = false
	DefaultLogLevel       = "error"
	DefaultConfigFilePath = ""
)

type Configuration struct {
	NeedHelp     bool
	URL          string        `yaml:"db_url" json:"db_url"`
	MaxDepth     uint64        `yaml:"max_depth" json:"max_depth"`
	Timeout      int           `yaml:"timeout" json:"timeout"`
	DepthIncStep int           `yaml:"depth_inc_step" json:"depth_inc_step"`
	Output       string        `yaml:"output" json:"output"`
	JSONLog      bool          `yaml:"json_log" json:"json_log"`
	WithPanic    bool          `yaml:"with_panic" json:"with_panic"`
	LogLevel     zerolog.Level `yaml:"log_level" json:"log_level"`
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

		c.MaxDepth = uint64(v)
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

	if value := os.Getenv("LOG_LEVEL"); value != "" {
		if c.LogLevel, err = zerolog.ParseLevel(value); err != nil {
			return err
		}
	}

	return err
}

func (c *Configuration) validate() (err error) {
	if _, err = url.ParseRequestURI(c.URL); err != nil {
		return fmt.Errorf("wrong url; %w", err)
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

	if fc.JSONLog {
		c.JSONLog = fc.JSONLog
	}

	if fc.WithPanic {
		c.WithPanic = fc.WithPanic
	}

	if fc.LogLevel != 0 {
		c.LogLevel = fc.LogLevel
	}
}

func (c *Configuration) configureBaseLogger() {
	zerolog.SetGlobalLevel(c.LogLevel)

	if c.JSONLog {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

		return
	}

	output := zerolog.ConsoleWriter{ //nolint:exhaustivestruct
		Out: os.Stdout,
		FormatTimestamp: func(i interface{}) string {
			parse, _ := time.Parse(time.RFC3339, i.(string))

			return parse.Format("2006-01-02 15:04:05")
		},
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf(" %-6s ", i))
		},
	}

	log.Logger = zerolog.New(output).
		With().Timestamp().
		CallerWithSkipFrameCount(SkipFrameCount).
		Logger()
}

// Load Reads configuration from files (yaml, json), .env, environment variables or arguments and returns.
func Load() (c Configuration, err error) {
	fConf, cfgPath, err := readFlags()
	if err != nil {
		return
	}

	if err = c.loadFromEnv(); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return
		}
	}

	if err = c.loadFromFile(cfgPath); err != nil {
		return
	}

	c.loadFromFlags(fConf)
	c.configureBaseLogger()

	if err = c.validate(); err != nil {
		return
	}

	return
}

func (c *Configuration) OutputToFile() bool {
	return c.Output != ""
}

func (c Configuration) String() (result string) {
	var (
		confMap map[string]interface{}
		err     error
	)

	if confMap, err = c.Map(); err != nil {
		return
	}

	for key, value := range confMap {
		result += key + ": " + fmt.Sprint(value) + "\n"
	}

	return
}

func (c *Configuration) Map() (result map[string]interface{}, err error) {
	var confJSON []byte

	if confJSON, err = json.Marshal(c); err != nil {
		return
	}

	err = json.Unmarshal(confJSON, &result)

	return
}

// ShowHelp - prints app help info
func ShowHelp() {
	fmt.Println(`
About:
Crawler is a program for find titles in URLs.

Examples of using:
./bin/crawler -u https://ya.ru # Finds all links and title up to the second level of nesting and outputs to the console
./bin/crawler -u https://ya.ru -o result.csv # Same but output to file
./bin/crawler -u https://ya.ru -l # Use for change log level (string debug/info/error etc)
./bin/crawler -u https://ya.ru -p # Fires panic and recover in first link`)
}

func readFlags() (c Configuration, path string, err error) {
	var ll string

	flag.StringVar(&c.URL, "u", DefaultURL, "URL for parsing")
	flag.Uint64Var(&c.MaxDepth, "m", DefaultMaxDepth, "Max depth")
	flag.IntVar(&c.Timeout, "t", DefaultTimeout, "Timeout for request")
	flag.IntVar(&c.DepthIncStep, "s", DefaultDepthStep, "Depth step")
	flag.StringVar(&c.Output, "o", DefaultOutput, "Output - file path or empty for log in console")
	flag.BoolVar(&c.NeedHelp, "h", DefaultNeedHelp, "Print help")
	flag.BoolVar(&c.JSONLog, "j", DefaultJSONLog, "JSON log format")
	flag.BoolVar(&c.WithPanic, "p", DefaultWithPanic, "Show test panic")
	flag.StringVar(&ll, "l", DefaultLogLevel, "Log level")
	flag.StringVar(&path, "c", DefaultConfigFilePath, "Config file path")
	flag.Parse()

	if c.LogLevel, err = zerolog.ParseLevel(ll); err != nil {
		return
	}

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
