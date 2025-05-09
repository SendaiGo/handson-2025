module github.com/Senoue/handson-2025/mcp-go-handson/mcp-compname

go 1.24.1

require github.com/mark3labs/mcp-go v0.23.1

replace github.com/Senoue/handson-2025/mcp-go-handson/mcp-compname/server => ./server

replace github.com/Senoue/handson-2025/mcp-go-handson/mcp-compname/name => ./name

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
)
