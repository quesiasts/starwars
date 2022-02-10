package controller

import (
	"context"
)

type PlanetsClient interface {
	GetPlanetAppearances(context.Context, string) (int, error)
}
