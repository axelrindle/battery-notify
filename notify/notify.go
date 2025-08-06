package notify

import (
	"fmt"

	"github.com/axelrindle/battery-notifier/config"
	"github.com/expr-lang/expr"
	"go.uber.org/zap"
)

type Notifier struct {
	Config *config.Config
	Logger *zap.Logger
}

func (n *Notifier) Notify(charge float64, config config.NotifyConfig) (bool, error) {
	env := map[string]any{
		"charge": charge,
	}

	program, err := expr.Compile(config.Condition, expr.Env(env))
	if err != nil {
		n.Logger.Error("failed to evaluate notification condition",
			zap.String("notification", config.ID),
			zap.String("condition", config.Condition),
			zap.Error(err),
		)
		return false, err
	}

	output, err := expr.Run(program, env)
	if err != nil {
		n.Logger.Error("failed to evaluate notification condition",
			zap.String("notification", config.ID),
			zap.String("condition", config.Condition),
			zap.Error(err),
		)
		return false, err
	}

	b, ok := output.(bool)
	if !ok {
		n.Logger.Error("notification condition did not evaluate to a boolean",
			zap.String("notification", config.ID),
			zap.String("condition", config.Condition),
		)
	}

	// skip if condition is false
	if !b {
		return false, nil
	}

	n.Logger.Debug("calling notifier",
		zap.String("type", config.Type),
		zap.String("notification", config.ID),
	)
	if config.Type == "http" {
		return n.notifyHttp(config)
	}

	// should never happen due to config validation
	return false, fmt.Errorf("invalid notification type %s", config.Type)
}
