package server

import (
	"context"
	"errors"

	companyname "github.com/Senoue/handson-2025/mcp-go-handson/mcp-compname/name"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func NewServer() *server.MCPServer {
	s := server.NewMCPServer(
		"Company name Cipher",
		"1.0.0",
	)

	corpname := mcp.NewTool(
		"Company name",
		mcp.WithDescription("Generate a Company name."),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the Company"),
		),
	)

	s.AddTool(corpname, corpnameHandler)
	return s
}

func corpnameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.Params.Arguments["name"].(string)
	if !ok {
		return nil, errors.New("company name must be a string")
	}

	result := companyname.CompanyName(len(name))
	return mcp.NewToolResultText(result), nil
}
