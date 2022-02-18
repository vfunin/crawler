package config

import (
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

type Configuration interface {
	OutputToFile() bool
	String() (result string)
	ShowHelp()
	NeedHelp() bool
	URL() string
	MaxDepth() uint64
	Timeout() int
	DepthIncStep() int
	Output() string
	WithPanic() bool
}

type configuration struct {
	needHelp     bool
	url          string        `yaml:"url"`
	maxDepth     uint64        `yaml:"max_depth"`
	timeout      int           `yaml:"timeout"`
	depthIncStep int           `yaml:"depth_inc_step"`
	output       string        `yaml:"output"`
	jsonLog      bool          `yaml:"json_log"`
	withPanic    bool          `yaml:"with_panic"`
	logLevel     zerolog.Level `yaml:"log_level"`
}

func (c *configuration) NeedHelp() bool {
	return c.needHelp
}

func (c *configuration) URL() string {
	return c.url
}

func (c *configuration) MaxDepth() uint64 {
	return c.maxDepth
}

func (c *configuration) Timeout() int {
	return c.timeout
}

func (c *configuration) DepthIncStep() int {
	return c.depthIncStep
}

func (c *configuration) Output() string {
	return c.output
}

func (c *configuration) JSONLog() bool {
	return c.jsonLog
}

func (c *configuration) WithPanic() bool {
	return c.withPanic
}

func (c *configuration) LogLevel() zerolog.Level {
	return c.logLevel
}

func (c *configuration) fillFromEnv() (err error) {
	if value := os.Getenv("URL"); value != "" {
		c.url = value
	}

	if value := os.Getenv("OUTPUT"); value != "" {
		c.output = value
	}

	var v int

	if value := os.Getenv("MAX_DEPTH"); value != "" {
		if v, err = strconv.Atoi(value); err != nil {
			return
		}

		c.maxDepth = uint64(v)
	}

	if value := os.Getenv("DEPTH_INC_STEP"); value != "" {
		if v, err = strconv.Atoi(value); err != nil {
			return
		}

		c.depthIncStep = v
	}

	if value := os.Getenv("TIMEOUT"); value != "" {
		if v, err = strconv.Atoi(value); err != nil {
			return
		}

		c.timeout = v
	}

	if value := os.Getenv("LOG_LEVEL"); value != "" {
		if c.logLevel, err = zerolog.ParseLevel(value); err != nil {
			return err
		}
	}

	return err
}

func (c *configuration) validate() (err error) {
	if _, err = url.ParseRequestURI(c.url); err != nil {
		return errors.New("wrong url (" + c.url + "); error: " + err.Error())
	}

	return
}

func (c *configuration) loadFromEnv() (err error) {
	if err = fillEnv(); err != nil {
		return
	}

	if err = c.fillFromEnv(); err != nil {
		return
	}

	return
}

func (c *configuration) loadFromFile(path string) (err error) {
	var content []byte

	if filepath.Ext(path) != ".yaml" {
		return
	}

	if content, err = os.ReadFile(path); err != nil {
		return
	}

	err = yaml.Unmarshal(content, &c)

	return
}

func (c *configuration) loadFromFlags(fc configuration) {
	if fc.url != "" {
		c.url = fc.url
	}

	if fc.output != "" {
		c.output = fc.output
	}

	if fc.maxDepth != 0 {
		c.maxDepth = fc.maxDepth
	}

	if fc.timeout != 0 {
		c.timeout = fc.timeout
	}

	if fc.depthIncStep != 0 {
		c.depthIncStep = fc.depthIncStep
	}

	if fc.needHelp {
		c.needHelp = fc.needHelp
	}

	if fc.jsonLog {
		c.jsonLog = fc.jsonLog
	}

	if fc.withPanic {
		c.withPanic = fc.withPanic
	}

	if fc.logLevel != 0 {
		c.logLevel = fc.logLevel
	}
}

func (c *configuration) configureBaseLogger() {
	zerolog.SetGlobalLevel(c.logLevel)

	if c.jsonLog {
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

// New Reads configuration from files (yaml, json), .env, environment variables or arguments and returns.
func New() (Configuration, error) {
	fConf, cfgPath, err := readFlags()
	if err != nil {
		return nil, err
	}

	c := &configuration{} //nolint:exhaustivestruct

	if err = c.loadFromEnv(); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	if err = c.loadFromFile(cfgPath); err != nil {
		return nil, err
	}

	c.loadFromFlags(fConf)
	c.configureBaseLogger()

	if err = c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *configuration) OutputToFile() bool {
	return c.output != ""
}

func (c configuration) String() (result string) {
	return fmt.Sprintf("%#v", c)
}

// ShowHelp - prints app help info
func (c *configuration) ShowHelp() {
	fmt.Println(`
About:
Crawler is a program for find titles in URLs.

Examples of using:
./bin/crawler -u https://ya.ru # Finds all links and title up to the second level of nesting and outputs to the console
./bin/crawler -u https://ya.ru -o result.csv # Same but output to file
./bin/crawler -u https://ya.ru -l # Use for change log level (string debug/info/error etc)
./bin/crawler -u https://ya.ru -p # Fires panic and recover in first link`)
}

func readFlags() (c configuration, path string, err error) {
	var ll string

	flag.StringVar(&c.url, "u", DefaultURL, "URL for parsing")
	flag.Uint64Var(&c.maxDepth, "m", DefaultMaxDepth, "Max depth")
	flag.IntVar(&c.timeout, "t", DefaultTimeout, "Timeout for request")
	flag.IntVar(&c.depthIncStep, "s", DefaultDepthStep, "Depth step")
	flag.StringVar(&c.output, "o", DefaultOutput, "Output - file path or empty for log in console")
	flag.BoolVar(&c.needHelp, "h", DefaultNeedHelp, "Print help")
	flag.BoolVar(&c.jsonLog, "j", DefaultJSONLog, "JSON log format")
	flag.BoolVar(&c.withPanic, "p", DefaultWithPanic, "Show test panic")
	flag.StringVar(&ll, "l", DefaultLogLevel, "Log level")
	flag.StringVar(&path, "c", DefaultConfigFilePath, "Config file path")
	flag.Parse()

	if c.logLevel, err = zerolog.ParseLevel(ll); err != nil {
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

	if len(matches) == 0 {
		return
	}

	return godotenv.Overload(matches...)
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
