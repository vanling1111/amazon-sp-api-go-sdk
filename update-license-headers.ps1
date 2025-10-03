# Batch update Go file license headers
# Change from Apache 2.0 to AGPL-3.0 + Commercial Dual License

$oldHeader1 = "// Copyright 2025 Amazon SP-API Go SDK Authors.`n// Licensed under the Apache License, Version 2.0."
$oldHeader2 = "// Copyright 2025 Amazon SP-API Go SDK Authors."

$newHeader = @"
// Copyright 2025 Amazon SP-API Go SDK Authors.
//
// This file is part of Amazon SP-API Go SDK.
//
// Amazon SP-API Go SDK is dual-licensed:
//
// 1. GNU Affero General Public License v3.0 (AGPL-3.0) for open source use
//    - Free for personal, educational, and open source projects
//    - Your project must also be open sourced under AGPL-3.0
//    - See: https://www.gnu.org/licenses/agpl-3.0.html
//
// 2. Commercial License for proprietary/commercial use
//    - Required for any commercial, enterprise, or proprietary use
//    - Allows closed source distribution
//    - Contact: vanling1111@gmail.com
//
// Unless you have obtained a commercial license, this file is licensed
// under AGPL-3.0. By using this software, you agree to comply with the
// terms of the applicable license.
"@

$count = 0
$updated = 0

Write-Host "Scanning Go files..." -ForegroundColor Yellow

Get-ChildItem -Path "." -Recurse -Filter "*.go" | 
    Where-Object { $_.FullName -notlike "*\.git*" } | 
    ForEach-Object {
        $count++
        $file = $_
        $content = Get-Content $file.FullName -Raw -Encoding UTF8
        
        if ($content -match "Licensed under the Apache License") {
            $content = $content -replace [regex]::Escape($oldHeader1), $newHeader
            $content = $content -replace [regex]::Escape($oldHeader2), $newHeader
            
            [System.IO.File]::WriteAllText($file.FullName, $content, [System.Text.UTF8Encoding]::new($false))
            
            $updated++
            
            if ($updated % 100 -eq 0) {
                Write-Host "Updated $updated files..." -ForegroundColor Cyan
            }
        }
    }

Write-Host ""
Write-Host "Complete!" -ForegroundColor Green
Write-Host "Total files scanned: $count"
Write-Host "Total files updated: $updated" -ForegroundColor Green
