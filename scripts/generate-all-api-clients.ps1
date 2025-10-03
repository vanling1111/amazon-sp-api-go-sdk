# Amazon SP-API Go SDK - Auto Generate API Clients
# Parse all 57 OpenAPI JSON files and generate client.go for each API

param([switch]$DryRun)
$ErrorActionPreference = "Stop"

$MODELS_DIR = "C:\Users\Administrator\selling-partner-api-models\models"
$SDK_DIR = "C:\Users\Administrator\amazon-sp-api-go-sdk\pkg\spapi"

Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "Generating API Clients from OpenAPI specs" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

# Load API config
. "$PSScriptRoot\api-config.ps1"

$SuccessCount = 0
$FailedCount = 0
$TotalOperations = 0

foreach ($API in $APIs) {
    $ApiName = $API.Name
    $Version = $API.Version
    $JsonFile = $API.JsonFile
    
    $DirName = "$ApiName-$Version"
    $PackageName = ($ApiName -replace '-','_') + "_" + ($Version -replace '-','_' -replace '\.','_')
    
    Write-Host "[$($SuccessCount + $FailedCount + 1)/$($APIs.Count)] $DirName" -ForegroundColor Yellow
    
    # Find JSON file
    if ($ApiName -like "*-model") {
        $JsonPath = "$MODELS_DIR\$ApiName\$JsonFile"
    } else {
        $JsonPath = "$MODELS_DIR\$($ApiName)-api-model\$JsonFile"
    }
    
    if (-not (Test-Path $JsonPath)) {
        Write-Host "  X JSON not found" -ForegroundColor Red
        $FailedCount++
        continue
    }
    
    try {
        # Parse JSON
        $apiSpec = Get-Content $JsonPath -Raw -Encoding UTF8 | ConvertFrom-Json
        
        # Extract operations
        $operations = @()
        if ($apiSpec.paths) {
            foreach ($pathProp in $apiSpec.paths.PSObject.Properties) {
                $path = $pathProp.Name
                $pathObj = $pathProp.Value
                
                foreach ($methodProp in $pathObj.PSObject.Properties) {
                    $method = $methodProp.Name.ToUpper()
                    if ($method -in @('GET','POST','PUT','DELETE','PATCH')) {
                        $operation = $methodProp.Value
                        $operations += @{
                            Id = if($operation.operationId){$operation.operationId}else{"Op"}
                            Method = $method
                            Path = $path
                            Summary = if($operation.summary){$operation.summary}else{""}
                        }
                    }
                }
            }
        }
        
        if ($operations.Count -eq 0) {
            Write-Host "  ! No operations found" -ForegroundColor Yellow
            continue
        }
        
        Write-Host "  + Found $($operations.Count) operations" -ForegroundColor Green
        $TotalOperations += $operations.Count
        
        # Generate code
        $TargetDir = Join-Path $SDK_DIR $DirName
        if (-not (Test-Path $TargetDir)) {
            Write-Host "  X Dir not found" -ForegroundColor Red
            $FailedCount++
            continue
        }
        
        if (-not $DryRun) {
            $ClientFile = Join-Path $TargetDir "client.go"
            . "$PSScriptRoot\client-generator-helper.ps1"
            Generate-ClientCode -OutputPath $ClientFile -PackageName $PackageName `
                -ApiName $ApiName -Version $Version -DirName $DirName -Operations $operations
        }
        
        Write-Host "  + Generated client.go" -ForegroundColor Green
        $SuccessCount++
        
    } catch {
        Write-Host "  X Error: $_" -ForegroundColor Red
        $FailedCount++
    }
}

Write-Host ""
Write-Host "Success: $SuccessCount" -ForegroundColor Green
Write-Host "Failed: $FailedCount" -ForegroundColor Red  
Write-Host "Total Operations: $TotalOperations" -ForegroundColor Yellow
