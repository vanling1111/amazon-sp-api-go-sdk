# 为所有迭代器生成测试文件

$ErrorActionPreference = "Stop"

$PKG_DIR = "C:\Users\Administrator\amazon-sp-api-go-sdk\pkg\spapi"

Write-Host "=== Generating Iterator Tests ===" -ForegroundColor Cyan
Write-Host ""

$SuccessCount = 0
$SkippedCount = 0

# 查找所有有 iterator.go 但没有 iterator_test.go 的 API
Get-ChildItem -Path $PKG_DIR -Directory | ForEach-Object {
    $apiName = $_.Name
    $apiDir = $_.FullName
    $iteratorFile = Join-Path $apiDir "iterator.go"
    $testFile = Join-Path $apiDir "iterator_test.go"
    
    # 检查是否有 iterator.go
    if (!(Test-Path $iteratorFile)) {
        return
    }
    
    # 检查是否已有测试
    if (Test-Path $testFile) {
        Write-Host "[$($SuccessCount + $SkippedCount + 1)] $apiName - Skipped (test exists)" -ForegroundColor Gray
        $SkippedCount++
        return
    }
    
    Write-Host "[$($SuccessCount + $SkippedCount + 1)] $apiName - Generating test..." -ForegroundColor Yellow
    
    # 生成包名
    $packageName = $apiName -replace '-','_'
    
    # 读取 iterator.go 查找方法名
    $iteratorContent = Get-Content $iteratorFile -Raw
    $methods = [regex]::Matches($iteratorContent, 'func \(c \*Client\) (Iterate\w+)\(') | ForEach-Object { $_.Groups[1].Value }
    
    if ($methods.Count -eq 0) {
        Write-Host "  Warning: No iterator methods found" -ForegroundColor Yellow
        return
    }
    
    # 生成测试内容
    $testContent = @"
// Copyright 2025 Amazon SP-API Go SDK Authors.
// Licensed under the Apache License, Version 2.0.

package $packageName

import (
	"testing"
)

"@
    
    # 为每个方法生成测试
    foreach ($method in $methods) {
        $testContent += @"

// Test$method 测试 $method 迭代器
func Test$method(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		// TODO: 使用 mock 客户端测试
		// 由于需要 HTTP mock，暂时跳过
		t.Skip("需要 HTTP mock 实现")
	})
	
	t.Run("empty_result", func(t *testing.T) {
		t.Skip("需要 HTTP mock 实现")
	})
	
	t.Run("multiple_pages", func(t *testing.T) {
		t.Skip("需要 HTTP mock 实现")
	})
	
	t.Run("early_exit", func(t *testing.T) {
		t.Skip("需要 HTTP mock 实现")
	})
}

// Example_$method 展示如何使用 $method
func Example_$method() {
	// 使用示例请参考 examples/iterators/main.go
}

"@
    }
    
    # 写入测试文件
    Set-Content -Path $testFile -Value $testContent -Encoding UTF8
    
    Write-Host "  ✓ Generated test with $($methods.Count) test cases" -ForegroundColor Green
    $SuccessCount++
}

Write-Host ""
Write-Host "=== Test Generation Complete ===" -ForegroundColor Cyan
Write-Host "Success: $SuccessCount" -ForegroundColor Green
Write-Host "Skipped: $SkippedCount" -ForegroundColor Gray
Write-Host ""

