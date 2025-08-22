package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/Done-0/jank/pkg/plugin/consts"

	jank "github.com/Done-0/jank/pkg/plugin"
)

type PluginConfig struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Repository  string `json:"repository"`
	Type        string `json:"type"`
}

type HelloPlugin struct {
	config *PluginConfig
}

func (p *HelloPlugin) Execute(ctx context.Context, method string, args map[string]any) (map[string]any, error) {
	switch method {
	case "greet":
		name, exists := args["name"]
		if !exists {
			return map[string]any{"error": "missing 'name' parameter"}, fmt.Errorf("name parameter is required")
		}

		return map[string]any{"message": fmt.Sprintf("Hello, %s!", name)}, nil
	case "info":
		if p.config == nil {
			return map[string]any{"error": "config not loaded"}, fmt.Errorf("config not loaded")
		}
		return map[string]any{
			"id":          p.config.ID,
			"name":        p.config.Name,
			"version":     p.config.Version,
			"author":      p.config.Author,
			"description": p.config.Description,
			"repository":  p.config.Repository,
			"type":        p.config.Type,
			"status":      "ready",
		}, nil
	case "echo":
		return args, nil
	default:
		return map[string]any{"error": "unknown method: " + method}, fmt.Errorf("unknown method: %s", method)
	}
}

func (p *HelloPlugin) HealthCheck(ctx context.Context) error {
	return nil
}

func main() {
	data, err := os.ReadFile("plugin.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read plugin.json: %v\n", err)
		os.Exit(1)
	}

	var config PluginConfig
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse plugin config: %v\n", err)
		os.Exit(1)
	}

	logger := hclog.New(&hclog.LoggerOptions{
		Name:   config.Name,
		Output: os.Stderr,
		Level:  hclog.Info,
	})

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: jank.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			consts.PluginTypeHandler: jank.NewGRPCPlugin(&HelloPlugin{config: &config}),
		},
		GRPCServer: plugin.DefaultGRPCServer,
		Logger:     logger,
	})
}
