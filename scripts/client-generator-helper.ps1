# 客户端代码生成辅助函数

function Generate-ClientCode {
    param(
        [string]$OutputPath,
        [string]$PackageName,
        [string]$ApiName,
        [string]$Version,
        [string]$DirName,
        [array]$Operations
    )
    
    $lines = @()
    
    # 检查是否需要 strings 包
    $needStrings = $false
    foreach ($op in $Operations) {
        if ($op.Path -match '\{[^}]+\}') {
            $needStrings = $true
            break
        }
    }
    
    # 文件头
    $lines += "// Copyright 2025 Amazon SP-API Go SDK Authors."
    $lines += "// Licensed under the Apache License, Version 2.0."
    $lines += ""
    $lines += "package $PackageName"
    $lines += ""
    $lines += "import ("
    $lines += "`t`"context`""
    $lines += "`t`"fmt`""
    if ($needStrings) {
        $lines += "`t`"strings`""
    }
    $lines += "`t"
    $lines += "`t`"github.com/vanling1111/amazon-sp-api-go-sdk/pkg/spapi`""
    $lines += ")"
    $lines += ""
    
    # Client 结构
    $lines += "// Client $ApiName API $Version"
    $lines += "type Client struct {"
    $lines += "`tbaseClient *spapi.Client"
    $lines += "}"
    $lines += ""
    
    # NewClient 构造函数
    $lines += "// NewClient creates API client"
    $lines += "func NewClient(baseClient *spapi.Client) *Client {"
    $lines += "`treturn &Client{baseClient: baseClient}"
    $lines += "}"
    $lines += ""
    
    # 去重操作ID
    $uniqueOps = @{}
    foreach ($op in $Operations) {
        $opId = $op.Id
        if (-not $uniqueOps.ContainsKey($opId)) {
            $uniqueOps[$opId] = $op
        }
    }
    
    # 为每个操作生成方法
    foreach ($op in $uniqueOps.Values) {
        $opId = $op.Id
        $method = $op.Method
        $path = $op.Path
        $summary = $op.Summary
        
        # 函数名
        $funcName = $opId.Substring(0,1).ToUpper() + $opId.Substring(1)
        
        # 提取路径参数
        $pathParams = [regex]::Matches($path, '\{([^}]+)\}') | ForEach-Object { $_.Groups[1].Value }
        
        # 构建参数
        $methodParams = ""
        if ($pathParams.Count -gt 0) {
            foreach ($param in $pathParams) {
                $pName = $param.Substring(0,1).ToLower() + $param.Substring(1)
                $methodParams += ", $pName string"
            }
        }
        
        if ($method -eq "GET") {
            $methodParams += ", query map[string]string"
        } elseif ($method -in @("POST","PUT","PATCH")) {
            $methodParams += ", body interface{}"
        }
        
        # 方法注释
        $lines += "// $funcName $summary"
        $lines += "// Method: $method | Path: $path"
        $lines += "func (c *Client) $funcName(ctx context.Context$methodParams) (interface{}, error) {"
        $lines += "`tpath := `"$path`""
        
        # 替换路径参数
        if ($pathParams.Count -gt 0) {
            foreach ($param in $pathParams) {
                $pName = $param.Substring(0,1).ToLower() + $param.Substring(1)
                $lines += "`tpath = strings.Replace(path, `"{$param}`", $pName, 1)"
            }
        }
        
        # 调用方法
        $lines += "`tvar result interface{}"
        if ($method -eq "GET") {
            $lines += "`terr := c.baseClient.Get(ctx, path, query, &result)"
        } elseif ($method -eq "POST") {
            $lines += "`terr := c.baseClient.Post(ctx, path, body, &result)"
        } elseif ($method -eq "PUT") {
            $lines += "`terr := c.baseClient.Put(ctx, path, body, &result)"
        } elseif ($method -eq "DELETE") {
            $lines += "`terr := c.baseClient.Delete(ctx, path, &result)"
        } elseif ($method -eq "PATCH") {
            $lines += "`terr := c.baseClient.DoRequest(ctx, `"PATCH`", path, nil, body, &result)"
        }
        
        $errorMsg = "$funcName" + ": %w"
        $lines += "`tif err != nil { return nil, fmt.Errorf(`"" + $errorMsg + "`", err) }"
        $lines += "`treturn result, nil"
        $lines += "}"
        $lines += ""
    }
    
    # 写入文件（无 BOM 的 UTF-8）
    $content = $lines -join "`n"
    $utf8NoBom = New-Object System.Text.UTF8Encoding $false
    [System.IO.File]::WriteAllText($OutputPath, $content, $utf8NoBom)
}

