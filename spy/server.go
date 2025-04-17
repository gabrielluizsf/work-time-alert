package spy

import (
	"context"
	"sync"

	server "github.com/i9si-sistemas/nine/pkg/server"
)

// Server implements the nine.Server interface and records all method calls
type Server struct {
	mu *sync.Mutex

	// Recorded method calls
	UseCalls        []UseCall
	GetCalls        []RouteCall
	PostCalls       []RouteCall
	PutCalls        []RouteCall
	PatchCalls      []RouteCall
	DeleteCalls     []RouteCall
	RouteCalls      []RouteCall
	GroupCalls      []GroupCall
	ServeFilesCalls []ServeFilesCall
	TestCalls       int
	ListenCalls     int
	ShutdownCalls   []context.Context
}

type UseCall struct {
	Middlewares any
	Err         error
}

type RouteCall struct {
	Path     string
	Handlers []any
	Err      error
}

type GroupCall struct {
	Prefix      string
	Middlewares []any
	ReturnGroup *server.RouteGroup
}

type ServeFilesCall struct {
	Prefix string
	Root   string
}

// NewServer creates a new server Spy instance
func NewServer() *Server {
	return &Server{
		mu:              new(sync.Mutex),
		UseCalls:        []UseCall{},
		GetCalls:        []RouteCall{},
		PostCalls:       []RouteCall{},
		PutCalls:        []RouteCall{},
		PatchCalls:      []RouteCall{},
		DeleteCalls:     []RouteCall{},
		RouteCalls:      []RouteCall{},
		GroupCalls:      []GroupCall{},
		ServeFilesCalls: []ServeFilesCall{},
		TestCalls:       0,
		ListenCalls:     0,
		ShutdownCalls:   []context.Context{},
	}
}

func (s *Server) Use(middlewares any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := error(nil)
	s.UseCalls = append(s.UseCalls, UseCall{
		Middlewares: middlewares,
		Err:         err,
	})
	return err
}

func (s *Server) Get(path string, handlers ...any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := error(nil)
	s.GetCalls = append(s.GetCalls, RouteCall{
		Path:     path,
		Handlers: handlers,
		Err:      err,
	})
	return err
}

func (s *Server) Post(path string, handlers ...any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := error(nil)
	s.PostCalls = append(s.PostCalls, RouteCall{
		Path:     path,
		Handlers: handlers,
		Err:      err,
	})
	return err
}

func (s *Server) Put(path string, handlers ...any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := error(nil)
	s.PutCalls = append(s.PutCalls, RouteCall{
		Path:     path,
		Handlers: handlers,
		Err:      err,
	})
	return err
}

func (s *Server) Patch(path string, handlers ...any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := error(nil)
	s.PatchCalls = append(s.PatchCalls, RouteCall{
		Path:     path,
		Handlers: handlers,
		Err:      err,
	})
	return err
}

func (s *Server) Delete(path string, handlers ...any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := error(nil)
	s.DeleteCalls = append(s.DeleteCalls, RouteCall{
		Path:     path,
		Handlers: handlers,
		Err:      err,
	})
	return err
}

func (s *Server) Route(prefix string, fn func(*server.RouteGroup)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.RouteCalls = append(s.RouteCalls, RouteCall{
		Path: prefix,
	})
}

func (s *Server) Group(prefix string, middlewares ...any) *server.RouteGroup {
	s.mu.Lock()
	defer s.mu.Unlock()

	group := server.NewRouteGroup(s, prefix, middlewares...)
	s.GroupCalls = append(s.GroupCalls, GroupCall{
		Prefix:      prefix,
		Middlewares: middlewares,
		ReturnGroup: group,
	})
	return group
}

func (s *Server) ServeFiles(prefix, root string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ServeFilesCalls = append(s.ServeFilesCalls, ServeFilesCall{
		Prefix: prefix,
		Root:   root,
	})
}

func (s *Server) Test() *server.TestServer {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.TestCalls++
	return &server.TestServer{}
}

func (s *Server) Listen() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ListenCalls++
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ShutdownCalls = append(s.ShutdownCalls, ctx)
	return nil
}
