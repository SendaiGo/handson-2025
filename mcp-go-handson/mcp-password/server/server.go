package server

import (
	"context"
	"errors"

	"github.com/Senoue/handson-2025/mcp-go-handson/mcp-password/random"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func NewServer() *server.MCPServer {
	s := server.NewMCPServer(
		"Random Password Cipher",
		"1.0.0",
	)

	length := mcp.NewTool(
		"random_password",
		mcp.WithDescription("Generate a random password."),
		mcp.WithNumber("length",
			mcp.Required(),
			mcp.Description("Length of the password"),
		),
	)

	s.AddTool(length, randomHandler)
	return s
}

func randomHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	length, ok := request.Params.Arguments["length"].(float64)
	if !ok {
		return nil, errors.New("length must be a number")
	}

	result := random.RandomPassword(int(length))
	return mcp.NewToolResultText(result), nil
}
