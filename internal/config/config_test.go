package config

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

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
