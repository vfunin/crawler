package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type fields struct {
	needHelp     bool
	url          string
	maxDepth     uint64
	timeout      int
	depthIncStep int
	output       string
	jsonLog      bool
	withPanic    bool
	logLevel     zerolog.Level
}

func TestNew(t *testing.T) {
	var (
		err error
		c   Configuration
	)

	if err = os.Setenv("URL", "http://go.test"); err != nil {
		assert.NotNil(t, err, "create config: can't set env url")
	}

	c, err = New()

	assert.Nil(t, err, "create config: error is not nil")
	assert.NotNil(t, c, "create config: config is not nil")
}

func TestDepthIncStep(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 99,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    99,
			wantErr: false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if got := c.DepthIncStep(); got != tt.want && !tt.wantErr {
				t.Errorf("DepthIncStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONLog(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 99,
				output:       "",
				jsonLog:      true,
				withPanic:    false,
				logLevel:     0,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    true,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if got := c.JSONLog(); got != tt.want && !tt.wantErr {
				t.Errorf("JSONLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogLevel(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    zerolog.Level
		wantErr bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     zerolog.PanicLevel,
			},
			want:    zerolog.PanicLevel,
			wantErr: false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 99,
				output:       "",
				jsonLog:      true,
				withPanic:    false,
				logLevel:     zerolog.InfoLevel,
			},
			want:    zerolog.InfoLevel,
			wantErr: false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     5,
			},
			want:    zerolog.InfoLevel,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if got := c.LogLevel(); got != tt.want && !tt.wantErr {
				t.Errorf("LogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     2,
				timeout:      0,
				depthIncStep: 99,
				output:       "",
				jsonLog:      true,
				withPanic:    false,
				logLevel:     0,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    2,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if got := c.MaxDepth(); got != tt.want && !tt.wantErr {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeedHelp(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     true,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 99,
				output:       "",
				jsonLog:      true,
				withPanic:    false,
				logLevel:     0,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    true,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if got := c.NeedHelp(); got != tt.want && !tt.wantErr {
				t.Errorf("NeedHelp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutput(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 99,
				output:       "test",
				jsonLog:      true,
				withPanic:    false,
				logLevel:     0,
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    "test",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if got := c.Output(); got != tt.want && !tt.wantErr {
				t.Errorf("Output() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutputToFile(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 99,
				output:       "test",
				jsonLog:      true,
				withPanic:    false,
				logLevel:     0,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    true,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if got := c.OutputToFile(); got != tt.want && !tt.wantErr {
				t.Errorf("OutputToFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleConfiguration_ShowHelp() {
	c := &configuration{
		needHelp:     true,
		url:          "",
		maxDepth:     0,
		timeout:      0,
		depthIncStep: 99,
		output:       "test",
		jsonLog:      true,
		withPanic:    false,
		logLevel:     0,
	}
	c.ShowHelp()
	//Output:
	//About:
	//Crawler is a program for find titles in URLs.
	//
	//Examples of using:
	//./bin/crawler -u https://ya.ru # Finds all links and title up to the second level of nesting and outputs to the console
	//./bin/crawler -u https://ya.ru -o result.csv # Same but output to file
	//./bin/crawler -u https://ya.ru -l # Use for change log level (string debug/info/error etc)
	//./bin/crawler -u https://ya.ru -p # Fires panic and recover in first link
}

func TestString(t *testing.T) {
	tests := []struct {
		name       string
		fields     fields
		wantResult string
		wantErr    bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     true,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			wantResult: "config.configuration{needHelp:true, url:\"\", maxDepth:0x0, timeout:0, depthIncStep:0, output:\"\", jsonLog:false, withPanic:false, logLevel:0}",
			wantErr:    false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 99,
				output:       "test",
				jsonLog:      true,
				withPanic:    false,
				logLevel:     0,
			},
			wantResult: "config.configuration{needHelp:false, url:\"\", maxDepth:0x0, timeout:0, depthIncStep:99, output:\"test\", jsonLog:true, withPanic:false, logLevel:0}",
			wantErr:    false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			wantResult: "test",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if gotResult := c.String(); gotResult != tt.wantResult && !tt.wantErr {
				t.Errorf("String() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTimeout(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     2,
				timeout:      10,
				depthIncStep: 99,
				output:       "",
				jsonLog:      true,
				withPanic:    false,
				logLevel:     0,
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    2,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if got := c.Timeout(); got != tt.want && !tt.wantErr {
				t.Errorf("Timeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURL(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "test",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "test2",
				maxDepth:     2,
				timeout:      0,
				depthIncStep: 99,
				output:       "",
				jsonLog:      true,
				withPanic:    false,
				logLevel:     0,
			},
			want:    "test2",
			wantErr: false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    "error",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if got := c.URL(); got != tt.want && !tt.wantErr {
				t.Errorf("URL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPanic(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "success result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     2,
				timeout:      0,
				depthIncStep: 99,
				output:       "",
				jsonLog:      true,
				withPanic:    true,
				logLevel:     0,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "error result",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			want:    true,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if got := c.WithPanic(); got != tt.want && !tt.wantErr {
				t.Errorf("WithPanic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleConfiguration_configureBaseLogger() {
	c := &configuration{
		needHelp:     false,
		url:          "",
		maxDepth:     0,
		timeout:      0,
		depthIncStep: 0,
		output:       "",
		jsonLog:      false,
		withPanic:    false,
		logLevel:     0,
	}

	c.configureBaseLogger()

	out := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true}
	out.FormatLevel = func(i interface{}) string { return strings.ToUpper(fmt.Sprintf("%-6s|", i)) }
	out.FormatFieldName = func(i interface{}) string { return fmt.Sprintf("%s:", i) }
	out.FormatFieldValue = func(i interface{}) string { return strings.ToUpper(fmt.Sprintf("%s", i)) }
	log := zerolog.New(out)

	log.Info().Msg("test")
	// Output: <nil> INFO  | test
}

func Test_configuration_fillFromEnv(t *testing.T) {
	c := &configuration{
		needHelp:     false,
		url:          "test",
		maxDepth:     0,
		timeout:      0,
		depthIncStep: 0,
		output:       "",
		jsonLog:      false,
		withPanic:    false,
		logLevel:     0,
	}

	err := c.fillFromEnv()

	assert.Nil(t, err, "error is not nil")

	tests := []struct {
		name    string
		envName string
		value   string
		wantErr bool
	}{
		{
			name:    "check timeout success",
			envName: "TIMEOUT",
			value:   "timeout",
			wantErr: true,
		},
		{
			name:    "check timeout error",
			envName: "TIMEOUT",
			value:   "2",
			wantErr: false,
		},
		{
			name:    "check URL success",
			envName: "URL",
			value:   "URL",
			wantErr: false,
		},
		{
			name:    "check OUTPUT success",
			envName: "OUTPUT",
			value:   "test",
			wantErr: false,
		},
		{
			name:    "check MAX_DEPTH success",
			envName: "MAX_DEPTH",
			value:   "2",
			wantErr: false,
		},
		{
			name:    "check MAX_DEPTH error",
			envName: "MAX_DEPTH",
			value:   "MAX_DEPTH",
			wantErr: true,
		},
		{
			name:    "check DEPTH_INC_STEP success",
			envName: "DEPTH_INC_STEP",
			value:   "2",
			wantErr: false,
		},
		{
			name:    "check DEPTH_INC_STEP error",
			envName: "DEPTH_INC_STEP",
			value:   "DEPTH_INC_STEP",
			wantErr: true,
		},
		{
			name:    "check LOG_LEVEL success",
			envName: "LOG_LEVEL",
			value:   "info",
			wantErr: false,
		},
		{
			name:    "check LOG_LEVEL error",
			envName: "LOG_LEVEL",
			value:   "test_level",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			err = os.Setenv(tt.envName, tt.value)
			assert.Nil(t, err, tt.envName+" - can't set env")
			err = c.fillFromEnv()
			if tt.wantErr {
				assert.NotNil(t, err, tt.envName)

				return
			}

			assert.Nil(t, err, tt.envName)
		})
	}

	os.Clearenv()
}

func Test_configuration_loadFromEnv(t *testing.T) {
	c := &configuration{
		needHelp:     false,
		url:          "test",
		maxDepth:     0,
		timeout:      0,
		depthIncStep: 0,
		output:       "",
		jsonLog:      false,
		withPanic:    false,
		logLevel:     0,
	}

	err := os.Setenv("TIMEOUT", "2")
	assert.Nil(t, err, "Timeout - can't set env")
	err = c.loadFromEnv()
	assert.Nil(t, err, "error is not nil")
	os.Clearenv()
	err = os.Setenv("TIMEOUT", "test")
	assert.Nil(t, err, "Timeout - can't set env")
	err = c.loadFromEnv()
	assert.NotNil(t, err, "error is not nil")
}

func Test_configuration_loadFromFile(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal file",
			args: args{
				path: "test.yaml",
			},
			wantErr: true,
		},
		{
			name: "wrong file type",
			args: args{
				path: "test.cfg",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     false,
				url:          "test",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			}
			if err := c.loadFromFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("loadFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_configuration_loadFromFlags(t *testing.T) {
	type args struct {
		fc configuration
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "normal case",
			fields: fields{
				needHelp:     false,
				url:          "",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			args: args{
				fc: configuration{
					needHelp:     true,
					url:          "test",
					maxDepth:     2,
					timeout:      3,
					depthIncStep: 4,
					output:       "test",
					jsonLog:      true,
					withPanic:    true,
					logLevel:     1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			c.loadFromFlags(tt.args.fc)
		})
	}
}

func Test_configuration_validate(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "normal url",
			fields: fields{
				needHelp:     false,
				url:          "https://test.ru",
				maxDepth:     0,
				timeout:      0,
				depthIncStep: 0,
				output:       "",
				jsonLog:      false,
				withPanic:    false,
				logLevel:     0,
			},
			wantErr: false,
		},
		{
			name: "wrong url",
			fields: fields{
				needHelp:     true,
				url:          "not-url",
				maxDepth:     1,
				timeout:      1,
				depthIncStep: 1,
				output:       "none",
				jsonLog:      true,
				withPanic:    true,
				logLevel:     2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &configuration{
				needHelp:     tt.fields.needHelp,
				url:          tt.fields.url,
				maxDepth:     tt.fields.maxDepth,
				timeout:      tt.fields.timeout,
				depthIncStep: tt.fields.depthIncStep,
				output:       tt.fields.output,
				jsonLog:      tt.fields.jsonLog,
				withPanic:    tt.fields.withPanic,
				logLevel:     tt.fields.logLevel,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fillEnv(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "normal case",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := fillEnv(); (err != nil) != tt.wantErr {
				t.Errorf("fillEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_remove(t *testing.T) {
	type args struct {
		s []string
		i int
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "normal case",
			args: args{
				s: []string{"1", "2", "3"},
				i: 1,
			},
			want:    []string{"1", "3"},
			wantErr: false,
		},
		{
			name: "error case",
			args: args{
				s: []string{"1", "2", "3"},
				i: 2,
			},
			want:    []string{"2"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
