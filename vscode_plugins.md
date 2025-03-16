# Go开发推荐VSCode插件列表

以下是针对Go语言开发（特别是API开发）推荐的VSCode插件列表，这些插件将帮助提高您的开发效率和代码质量。

## 核心插件

### 1. Go (官方插件)
- **插件ID**: `golang.go`
- **功能**: 提供Go语言的核心支持，包括代码补全、格式化、语法高亮、代码导航、调试等
- **必备程度**: ★★★★★
- **安装后配置**: 安装后会提示安装相关工具，建议全部安装

### 2. Go Test Explorer
- **插件ID**: `premparihar.gotestexplorer`
- **功能**: 可视化测试运行器，方便运行和调试Go测试
- **必备程度**: ★★★★☆

### 3. Go Doc
- **插件ID**: `msyrus.go-doc`
- **功能**: 快速查看Go文档
- **必备程度**: ★★★★☆

## 代码质量与效率

### 4. Go Outliner
- **插件ID**: `766b.go-outliner`
- **功能**: 提供Go文件的结构大纲视图
- **必备程度**: ★★★☆☆

### 5. Error Lens
- **插件ID**: `usernamehw.errorlens`
- **功能**: 直接在代码行内显示错误和警告
- **必备程度**: ★★★★☆

### 6. Code Spell Checker
- **插件ID**: `streetsidesoftware.code-spell-checker`
- **功能**: 检查代码中的拼写错误
- **必备程度**: ★★★☆☆

## API开发相关

### 7. REST Client
- **插件ID**: `humao.rest-client`
- **功能**: 在VSCode中直接发送HTTP请求，测试API
- **必备程度**: ★★★★★
- **特别适合**: 您的Silicon Flow API客户端开发和测试

### 8. Thunder Client
- **插件ID**: `rangav.vscode-thunder-client`
- **功能**: 轻量级的REST API客户端，类似Postman
- **必备程度**: ★★★★☆

### 9. YAML
- **插件ID**: `redhat.vscode-yaml`
- **功能**: YAML文件支持，对API配置文件很有用
- **必备程度**: ★★★☆☆

## Git与协作

### 10. GitLens
- **插件ID**: `eamodio.gitlens`
- **功能**: 增强Git功能，查看代码历史和作者
- **必备程度**: ★★★★☆

### 11. Git History
- **插件ID**: `donjayamanne.githistory`
- **功能**: 查看和搜索Git日志
- **必备程度**: ★★★☆☆

## 主题与界面

### 12. Material Icon Theme
- **插件ID**: `pkief.material-icon-theme`
- **功能**: 美化文件图标
- **必备程度**: ★★★☆☆

### 13. One Dark Pro
- **插件ID**: `zhuangtongfa.material-theme`
- **功能**: 流行的深色主题
- **必备程度**: ★★★☆☆

## 工具集成

### 14. Docker
- **插件ID**: `ms-azuretools.vscode-docker`
- **功能**: Docker集成，方便容器化应用开发
- **必备程度**: ★★★★☆

### 15. Remote - SSH
- **插件ID**: `ms-vscode-remote.remote-ssh`
- **功能**: 通过SSH连接到远程服务器进行开发
- **必备程度**: ★★★★☆

## 安装方法

1. 打开VSCode
2. 按下`Ctrl+Shift+X`(Windows/Linux)或`Cmd+Shift+X`(Mac)打开扩展面板
3. 搜索插件ID或名称
4. 点击安装

## Go开发环境配置建议

为了获得最佳的Go开发体验，建议在VSCode的`settings.json`中添加以下配置：

```json
{
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.lintFlags": ["--fast"],
  "go.formatTool": "goimports",
  "editor.formatOnSave": true,
  "[go]": {
    "editor.defaultFormatter": "golang.go",
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "go.coverOnSave": true,
  "go.coverageDecorator": {
    "type": "highlight"
  }
}
```

## 必要的Go工具

以下Go工具建议安装（可通过Go插件自动安装）：

- gopls: Go语言服务器
- golangci-lint: 代码质量检查工具
- dlv: 调试器
- goimports: 导入格式化工具

这些插件和配置将帮助您更高效地开发Silicon Flow API客户端，提高代码质量和开发体验。