package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-rate-limit-plugin] Starting Rate Limit plugin...")
	plugin := sdk.Init("hc-rate-limit-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("RateLimitStatus", "Rate limit plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getRateLimitStatus",
		sdk.ComplexObjectField("Get rate limit plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-rate-limit-plugin] Plugin ready")
	plugin.Serve()
}
