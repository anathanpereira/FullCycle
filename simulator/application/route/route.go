package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"positon"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"positon"`
	Finished bool      `json:"finished"`
}

func NewRoute() *Route {
	return &Route{}
}

func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("route id not informed")
	}

	file, err := os.Open("destinations/" + route.ID + ".txt")
	if err != nil {
		return err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		route.Positions = append(route.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}
	return nil
}

func (route *Route) ExportJsonPositions() ([]string, error) {
	var parsedRoute PartialRoutePosition
	var jsonRoutes []string
	totalPositions := len(route.Positions)

	for i, positon := range route.Positions {
		parsedRoute.ID = route.ID
		parsedRoute.ClientID = route.ClientID
		parsedRoute.Position = []float64{positon.Lat, positon.Long}
		parsedRoute.Finished = false

		if totalPositions-1 == i {
			parsedRoute.Finished = true
		}

		jsonRoute, err := json.Marshal(parsedRoute)
		if err != nil {
			return nil, err
		}

		jsonRoutes = append(jsonRoutes, string(jsonRoute))
	}
	return jsonRoutes, nil
}
