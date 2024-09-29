package handlers

import (
	"github.com/brianxor/nudata-solver/payload"
	"github.com/gofiber/fiber/v2"
)

func HandleNudataSolver(ctx *fiber.Ctx) error {
	var reqBody struct {
		WebsiteName string `json:"websiteName"`
		Proxy       string `json:"proxy"`
	}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed parsing body",
		})
	}

	websiteName := reqBody.WebsiteName
	proxy := reqBody.Proxy

	if websiteName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "missing website name",
		})
	}

	solver, err := payload.NewSolver(websiteName, proxy)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed creating solver",
		})
	}

	solution, err := solver.Solve()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed solving",
		})
	}

	respBody := fiber.Map{
		"success": true,
		"solution": fiber.Map{
			"payload":   solution.NdsPmd,
			"sessionId": solution.Sid,
			"solveTime": solution.SolveTime,
		},
	}

	return ctx.Status(fiber.StatusOK).JSON(respBody)
}
