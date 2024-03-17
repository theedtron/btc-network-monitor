package resource

import "btc-network-monitor/internal/core/services"

type HTTPHandler struct {
	userService    *services.UserService
	authService    *services.AuthService
	monitorService *services.MonitorService
}

func NewHTTPHandler(options ...interface{}) *HTTPHandler {
	handler := &HTTPHandler{
		userService:    services.NewUserService(),
		authService:    services.NewAuthService(),
		monitorService: services.NewMonitorService(),
	}
	return handler
}
