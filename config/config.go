package config

type NotifyConfig struct {
	Condition string `key:"condition" validate:"required"`
	Type      string `key:"type" validate:"required,oneof=exec,http"`
	ID        string `key:"id" validate:"required"`

	Url     string            `key:"url" validate:"excluded_unless=type http,required_if=type http"`
	Method  string            `key:"method" validate:"excluded_unless=type http,required_if=type http,oneof=post put"`
	Headers map[string]string `key:"headers" validate:"excluded_unless=type http"`
	Body    string            `key:"body" validate:"excluded_unless=type http"`
}

type Config struct {
	Environment string `key:"env" default:"production" validate:"oneof=production development"`

	DevicePath string `key:"device" validate:"required"`
	Interval   int64  `key:"refresh" default:"30"`

	Notifications []NotifyConfig `key:"notifications"`
}
