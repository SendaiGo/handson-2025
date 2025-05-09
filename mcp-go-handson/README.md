# mcp-goの基本

READMEのサンプルを確認してみましょう。

```go
package main

import (
    "context"
    "errors"
    "fmt"

    "github.com/mark3labs/mcp-go/mcp"
    "github.com/mark3labs/mcp-go/server"
)

func main() {
    // MCPサーバーを作成
    s := server.NewMCPServer(
        "Demo 🚀",
        "1.0.0",
    )

    // ツールを追加
    tool := mcp.NewTool("hello_world",
        mcp.WithDescription("Say hello to someone"),
        mcp.WithString("name",
            mcp.Required(),
            mcp.Description("Name of the person to greet"),
        ),
    )

    // ツールハンドラーを追加
    s.AddTool(tool, helloHandler)

    // Stdioサーバーを起動
    if err := server.ServeStdio(s); err != nil {
        fmt.Printf("Server error: %v\n", err)
    }
}

func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    name, ok := request.Params.Arguments["name"].(string)
    if !ok {
        return nil, errors.New("name must be a string")
    }

    return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}
```

https://github.com/mark3labs/mcp-go/blob/v0.23.1/README.md

まずは、実際に動作する様子を確認するため、MCPホストの設定にこのMCPサーバーを登録してみましょう。

VS Codeをご利用の場合は、以下の手順を実行します。

1. `Cmd + Shift +P` からコマンドパレットを開く
2. `MCP: Add Server` を実行
3. `Command: (stdio)` を選択
4. `Enter Command` と聞かれるので `go` と入力
5. `Enter Server ID` と聞かれるので `hello` と入力
6. 表示されるJSONの `servers.hello` を以下の内容に修正する

```json
"servers": {
  "hello": {
    "type": "stdio",
    "command": "go",
    "args": [
      "run",
      "github.com/Senoue/handson-2025/mcp-go-handson/hello@latest"
    ]
  }
}
```

ここまで設定が完了したら、 `hello` のキーの上に表示される `Start` をクリックして、GitHub CopilotのAgentを開いてください。

チャットに「fooにhelloと言ってください」と入力すると、 `hello_world` のツールが選択され、呼ばれることが確認できます。

