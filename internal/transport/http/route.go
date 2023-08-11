package http

func (s *Server) setupRoutes() {
	v1 := s.app.Group("/api/v1")
	v1.GET("/health/ready", s.h.Ready)
	v1.GET("/health/live", s.h.Live)
}
