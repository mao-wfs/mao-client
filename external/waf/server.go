package waf

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/mao-wfs/mao-ctrl/adapters/controller"
	"github.com/mao-wfs/mao-ctrl/adapters/gateway/device"
	"github.com/mao-wfs/mao-ctrl/external/client"
	"github.com/mao-wfs/mao-ctrl/external/waf/handler"
)

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	return &Server{Echo: echo.New()}
}

func (s *Server) Run() {
	s.setRouter()
	s.run()
}

func (s *Server) run() {
	addr := s.getAddr()
	s.Echo.Logger.Fatal(s.Echo.Start(addr))
}

func (s *Server) setRouter() {
	s.Echo.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	corrAddr := os.Getenv("CORREATOR_ADDRESS")
	corrClt, err := client.NewTCPClient(corrAddr)
	if err != nil {
		s.Echo.Logger.Fatal(err)
	}
	swAddr := os.Getenv("SWITCH_ADDRESS")
	swClt, err := client.NewTCPClient(swAddr)
	if err != nil {
		s.Echo.Logger.Fatal(err)
	}

	wfsHandler := device.NewWFSHandler(corrClt, swClt)
	api := s.Echo.Group(s.getEndPointRoot())
	{
		wfsctrl := controller.NewWFSController(wfsHandler)
		api.POST("/initialize", handler.InitializeWFS(wfsctrl))
		api.POST("/finalize", handler.FinalizeWFS(wfsctrl))
		api.POST("/start", handler.StartWFS(wfsctrl))
		api.POST("/halt", handler.HaltWFS(wfsctrl))
	}
}

func (s *Server) getAddr() string {
	return os.Getenv("API_ADDRESS")
}

func (s *Server) getVersion() string {
	return os.Getenv("API_VERSION")
}

func (s *Server) getEndPointRoot() string {
	return fmt.Sprintf("api/%s", s.getVersion())
}
