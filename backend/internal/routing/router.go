package routing

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
)

type ConnectionStatus[T comparable] struct {
	ID        T    `json:"robot_id"`
	Connected bool `json:"connected"`
}

type Router[T comparable] interface {
	StatusList() []ConnectionStatus[T]
	IDs() []T
	RegisterRobot(id T, con *websocket.Conn) (*Connection, error)
	ConnectRemote(id T, con *websocket.Conn) (*Connection, error)
	DisconnectRemote(id T) error
	DisconnectRobot(id T) error
}

func NewRouter[T comparable]() Router[T] {
	return &router[T]{
		routes: make(map[T]*Connection),
	}
}

var _ Router[int] = &router[int]{}

type router[T comparable] struct {
	routes map[T]*Connection
}

func (r *router[T]) IDs() []T {
	ids := make([]T, len(r.routes))

	count := 0

	for key := range r.routes {
		ids[count] = key
		count++
	}

	return ids
}

func (r *router[T]) StatusList() []ConnectionStatus[T] {

	status := make([]ConnectionStatus[T], len(r.routes))

	count := 0
	for id, route := range r.routes {
		status[count] = ConnectionStatus[T]{
			ID:        id,
			Connected: route.IsConnected(),
		}
		count++
	}

	return status
}

func (r *router[T]) exists(id T) bool {
	_, ok := r.routes[id]
	return ok
}

func (r *router[T]) RegisterRobot(id T, con *websocket.Conn) (*Connection, error) {
	if r.exists(id) {
		return nil, fmt.Errorf("id already registered")
	}

	connection := &Connection{
		robot: con,
	}

	r.routes[id] = connection

	return connection, nil
}

func (r *router[T]) ConnectRemote(id T, con *websocket.Conn) (*Connection, error) {

	if !r.exists(id) {
		return nil, fmt.Errorf("id not found")
	}

	route := r.routes[id]

	if route.IsConnected() {
		return nil, fmt.Errorf("already connected")
	}

	route.remote = con
	return route, nil
}

func (r *router[T]) DisconnectRemote(id T) error {
	route, ok := r.routes[id]

	if !ok {
		return fmt.Errorf("id not found")
	}

	if route.remote == nil {
		return fmt.Errorf("remote connection does not exist")
	}

	err := route.remote.Close()
	if err != nil {
		return err
	}

	route.remote = nil
	return nil
}

func (r *router[T]) DisconnectRobot(id T) error {
	route, ok := r.routes[id]

	if !ok {
		return fmt.Errorf("id not found")
	}

	if route.remote != nil {
		err := route.remote.Close()
		if err != nil {
			return err
		}
	}

	if route.robot == nil {
		return fmt.Errorf("robot connection does not exist")
	}

	err := route.robot.Close()
	if err != nil {
		return err
	}

	delete(r.routes, id)
	return nil
}
