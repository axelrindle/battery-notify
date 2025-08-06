package notify

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/axelrindle/battery-notifier/config"
	"github.com/axelrindle/battery-notifier/version"
	"go.uber.org/zap"
)

func (n *Notifier) notifyHttp(config config.NotifyConfig) (bool, error) {
	method := strings.ToUpper(config.Method)
	req, err := http.NewRequest(method, config.Url, bytes.NewBufferString(config.Body))
	if err != nil {
		return false, err
	}

	req.Header.Add("User-Agent", fmt.Sprintf("battery-notify/%s https://github.com/axelrindle/battery-notify", version.Version))
	for k, v := range config.Headers {
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	if res.StatusCode != 200 {
		n.Logger.Debug("http notifier response body",
			zap.String("notification", config.ID),
			zap.String("body", string(resBody)),
		)
		return false, fmt.Errorf("%s %s returned %s", method, config.Url, res.Status)
	}

	return true, nil
}
