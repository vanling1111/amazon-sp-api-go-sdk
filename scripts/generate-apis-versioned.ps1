# ========================================
# Amazon SP-API Go SDK - Versioned API Generation Script
# ========================================
# 特点：
# 1. 从JSON文件名提取版本号
# 2. 为每个API创建带版本号的包（如 orders_v0, catalog_items_v2022_04_01）
# 3. 只保留类型定义文件（model_*.go）
# 4. 支持同一API的多个版本共存
# ========================================

$ErrorActionPreference = "Continue"

# Configure paths
$JAVA_HOME = "C:\Program Files\Microsoft\jdk-21.0.8.9-hotspot"
$JAVA_EXE = "$JAVA_HOME\bin\java.exe"
$SWAGGER_JAR = "C:\Users\Administrator\swagger-codegen-cli.jar"
$MODELS_DIR = "C:\Users\Administrator\selling-partner-api-models\models"
$OUTPUT_DIR = "C:\Users\Administrator\amazon-sp-api-go-sdk\pkg\spapi"
$TEMP_DIR = "C:\Users\Administrator\temp-swagger-gen"

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Starting Versioned API Generation" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# API配置列表 (基于官方仓库的完整API列表 - 58个API版本)
# 格式: @{Name="api-name"; Version="vX"; JsonFile="file.json"}
$APIs = @(
    # Orders - 已验证
    @{Name="orders"; Version="v0"; JsonFile="ordersV0.json"},
    
    # Feeds - 已验证
    @{Name="feeds"; Version="v2021-06-30"; JsonFile="feeds_2021-06-30.json"},
    
    # Catalog Items - 已验证 (3个版本)
    @{Name="catalog-items"; Version="v0"; JsonFile="catalogItemsV0.json"},
    @{Name="catalog-items"; Version="v2020-12-01"; JsonFile="catalogItems_2020-12-01.json"},
    @{Name="catalog-items"; Version="v2022-04-01"; JsonFile="catalogItems_2022-04-01.json"},
    
    # Reports - 已验证
    @{Name="reports"; Version="v2021-06-30"; JsonFile="reports_2021-06-30.json"},
    
    # Finances - 已验证 (2个版本)
    @{Name="finances"; Version="v0"; JsonFile="financesV0.json"},
    @{Name="finances"; Version="v2024-06-19"; JsonFile="finances_2024-06-19.json"},
    
    # FBA Inventory - 已验证
    @{Name="fba-inventory"; Version="v1"; JsonFile="fbaInventory.json"},
    
    # FBA Eligibility - 已验证
    @{Name="fba-inbound-eligibility"; Version="v1"; JsonFile="fbaInbound.json"},
    
    # Fulfillment Inbound - 已验证 (2个版本)
    @{Name="fulfillment-inbound"; Version="v0"; JsonFile="fulfillmentInboundV0.json"},
    @{Name="fulfillment-inbound"; Version="v2024-03-20"; JsonFile="fulfillmentInbound_2024-03-20.json"},
    
    # Fulfillment Outbound - 已验证
    @{Name="fulfillment-outbound"; Version="v2020-07-01"; JsonFile="fulfillmentOutbound_2020-07-01.json"},
    
    # Listings Items - 已验证 (2个版本)
    @{Name="listings-items"; Version="v2020-09-01"; JsonFile="listingsItems_2020-09-01.json"},
    @{Name="listings-items"; Version="v2021-08-01"; JsonFile="listingsItems_2021-08-01.json"},
    
    # Listings Restrictions - 已验证
    @{Name="listings-restrictions"; Version="v2021-08-01"; JsonFile="listingsRestrictions_2021-08-01.json"},
    
    # Merchant Fulfillment - 已验证
    @{Name="merchant-fulfillment"; Version="v0"; JsonFile="merchantFulfillmentV0.json"},
    
    # Messaging - 已验证
    @{Name="messaging"; Version="v1"; JsonFile="messaging.json"},
    
    # Notifications - 已验证
    @{Name="notifications"; Version="v1"; JsonFile="notifications.json"},
    
    # Pricing - 已验证 (2个版本)
    @{Name="product-pricing"; Version="v0"; JsonFile="productPricingV0.json"},
    @{Name="product-pricing"; Version="v2022-05-01"; JsonFile="productPricing_2022-05-01.json"},
    
    # Product Fees - 已验证
    @{Name="product-fees"; Version="v0"; JsonFile="productFeesV0.json"},
    
    # Product Type Definitions - 已验证
    @{Name="product-type-definitions"; Version="v2020-09-01"; JsonFile="definitionsProductTypes_2020-09-01.json"},
    
    # Replenishment - 已验证
    @{Name="replenishment"; Version="v2022-11-07"; JsonFile="replenishment-2022-11-07.json"},
    
    # Sales - 已验证
    @{Name="sales"; Version="v1"; JsonFile="sales.json"},
    
    # Sellers - 已验证
    @{Name="sellers"; Version="v1"; JsonFile="sellers.json"},
    
    # Seller Wallet - 已验证
    @{Name="seller-wallet"; Version="v2024-03-01"; JsonFile="sellerWallet_2024-03-01.json"},
    
    # Services - 已验证
    @{Name="services"; Version="v1"; JsonFile="services.json"},
    
    # Shipping - 已验证
    @{Name="shipping"; Version="v2"; JsonFile="shippingV2.json"},
    
    # Solicitations - 已验证
    @{Name="solicitations"; Version="v1"; JsonFile="solicitations.json"},
    
    # Supply Sources - 已验证
    @{Name="supply-sources"; Version="v2020-07-01"; JsonFile="supplySources_2020-07-01.json"},
    
    # Tokens - 已验证
    @{Name="tokens"; Version="v2021-03-01"; JsonFile="tokens_2021-03-01.json"},
    
    # Transfers - 已验证 (在finances目录中)
    @{Name="finances"; Version="v2024-06-01-transfers"; JsonFile="transfers_2024-06-01.json"},
    
    # Uploads - 已验证
    @{Name="uploads"; Version="v2020-11-01"; JsonFile="uploads_2020-11-01.json"},
    
    # Vehicles - 已验证
    @{Name="vehicles"; Version="v2024-11-01"; JsonFile="vehicles_2024-11-01.json"},
    
    # A+ Content - 已验证
    @{Name="aplus-content"; Version="v2020-11-01"; JsonFile="aplusContent_2020-11-01.json"},
    
    # App Integrations - 已验证
    @{Name="application-integrations"; Version="v2024-04-01"; JsonFile="appIntegrations-2024-04-01.json"},
    
    # Applications - 已验证
    @{Name="application-management"; Version="v2023-11-30"; JsonFile="application_2023-11-30.json"},
    
    # Amazon Warehousing and Distribution - 已验证
    @{Name="amazon-warehousing-and-distribution-model"; Version="v2024-05-09"; JsonFile="awd_2024-05-09.json"},
    
    # Customer Feedback - 已验证
    @{Name="customer-feedback"; Version="v2024-06-01"; JsonFile="customerFeedback_2024-06-01.json"},
    
    # Data Kiosk - 已验证
    @{Name="data-kiosk"; Version="v2023-11-15"; JsonFile="dataKiosk_2023-11-15.json"},
    
    # Easy Ship - 已验证
    @{Name="easy-ship-model"; Version="v2022-03-23"; JsonFile="easyShip_2022-03-23.json"},
    
    # Invoices - 已验证
    @{Name="invoices"; Version="v2024-06-19"; JsonFile="InvoicesApiModel_2024-06-19.json"},
    
    # Invoicing - 已验证
    @{Name="shipment-invoicing"; Version="v0"; JsonFile="shipmentInvoicingV0.json"},
    
    # Vendor APIs - 已验证
    @{Name="vendor-direct-fulfillment-inventory"; Version="v1"; JsonFile="vendorDirectFulfillmentInventoryV1.json"},
    @{Name="vendor-direct-fulfillment-orders"; Version="v1"; JsonFile="vendorDirectFulfillmentOrdersV1.json"},
    @{Name="vendor-direct-fulfillment-orders"; Version="v2021-12-28"; JsonFile="vendorDirectFulfillmentOrders_2021-12-28.json"},
    @{Name="vendor-direct-fulfillment-payments"; Version="v1"; JsonFile="vendorDirectFulfillmentPaymentsV1.json"},
    @{Name="vendor-direct-fulfillment-sandbox-test-data"; Version="v2021-10-28"; JsonFile="vendorDirectFulfillmentSandboxData_2021-10-28.json"},
    @{Name="vendor-direct-fulfillment-shipping"; Version="v1"; JsonFile="vendorDirectFulfillmentShippingV1.json"},
    @{Name="vendor-direct-fulfillment-shipping"; Version="v2021-12-28"; JsonFile="vendorDirectFulfillmentShipping_2021-12-28.json"},
    @{Name="vendor-direct-fulfillment-transactions"; Version="v1"; JsonFile="vendorDirectFulfillmentTransactionsV1.json"},
    @{Name="vendor-direct-fulfillment-transactions"; Version="v2021-12-28"; JsonFile="vendorDirectFulfillmentTransactions_2021-12-28.json"},
    @{Name="vendor-invoices"; Version="v1"; JsonFile="vendorInvoices.json"},
    @{Name="vendor-orders"; Version="v1"; JsonFile="vendorOrders.json"},
    @{Name="vendor-shipments"; Version="v1"; JsonFile="vendorShipments.json"},
    @{Name="vendor-transaction-status"; Version="v1"; JsonFile="vendorTransactionStatus.json"}
)

Write-Host "Total APIs to generate: $($APIs.Count)" -ForegroundColor Green
Write-Host ""

$SuccessCount = 0
$FailedCount = 0
$FailedAPIs = @()

foreach ($API in $APIs) {
    $ApiName = $API.Name
    $Version = $API.Version
    $JsonFile = $API.JsonFile
    
    # 创建版本化的包名
    # 例如: orders_v0, catalog_items_v2022_04_01
    $VersionSuffix = $Version -replace '-','_' -replace '\.','_'
    $PackageName = ($ApiName -replace '-','_') + "_" + $VersionSuffix
    
    # 创建版本化的目录名
    # 例如: orders-v0, catalog-items-v2022-04-01
    $DirName = $ApiName + "-" + $Version
    
    Write-Host "[$($SuccessCount + $FailedCount + 1)/$($APIs.Count)] $DirName (package: $PackageName)" -ForegroundColor Yellow
    
    # 构建本地JSON文件路径
    # 如果API名称已经包含-model，则不再添加-api-model
    if ($ApiName -like "*-model") {
        $JsonUrl = "$MODELS_DIR\$($ApiName)\$($JsonFile)"
    } else {
        $JsonUrl = "$MODELS_DIR\$($ApiName)-api-model\$($JsonFile)"
    }
    
    try {
        # 创建临时目录
        if (Test-Path $TEMP_DIR) {
            Remove-Item -Path $TEMP_DIR -Recurse -Force
        }
        New-Item -ItemType Directory -Path $TEMP_DIR -Force | Out-Null
        
        # 检查本地JSON文件是否存在
        if (-not (Test-Path $JsonUrl)) {
            throw "Local JSON file not found: $JsonUrl"
        }
        
        # 运行swagger-codegen
        $GenerateCmd = "& `"$JAVA_EXE`" -jar `"$SWAGGER_JAR`" generate -i `"$JsonUrl`" -l go --additional-properties packageName=$PackageName -o `"$TEMP_DIR`""
        $Output = Invoke-Expression $GenerateCmd 2>&1
        Write-Host "  swagger-codegen output: $Output" -ForegroundColor Gray
        
        if ($LASTEXITCODE -ne 0) {
            throw "swagger-codegen failed"
        }
        
        # 创建目标目录
        $TargetDir = Join-Path $OUTPUT_DIR $DirName
        if (Test-Path $TargetDir) {
            Remove-Item -Path $TargetDir -Recurse -Force
        }
        New-Item -ItemType Directory -Path $TargetDir -Force | Out-Null
        
        # 只复制 model_*.go 文件
        $ModelFiles = Get-ChildItem -Path $TEMP_DIR -Filter "model_*.go" -Recurse
        if ($ModelFiles.Count -gt 0) {
            foreach ($File in $ModelFiles) {
                Copy-Item -Path $File.FullName -Destination $TargetDir -Force
            }
            Write-Host "  ✓ Generated $($ModelFiles.Count) model files" -ForegroundColor Green
            $SuccessCount++
        } else {
            throw "No model files generated"
        }
        
    } catch {
        Write-Host "  ✗ Failed: $_" -ForegroundColor Red
        $FailedAPIs += @{API=$DirName; Error=$_.Exception.Message}
        $FailedCount++
    } finally {
        # 清理临时目录
        if (Test-Path $TEMP_DIR) {
            Remove-Item -Path $TEMP_DIR -Recurse -Force -ErrorAction SilentlyContinue
        }
    }
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Generation Complete" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Success: $SuccessCount" -ForegroundColor Green
Write-Host "Failed: $FailedCount" -ForegroundColor Red
Write-Host ""

if ($FailedAPIs.Count -gt 0) {
    Write-Host "Failed APIs:" -ForegroundColor Red
    foreach ($Failed in $FailedAPIs) {
        Write-Host "  - $($Failed.API): $($Failed.Error)" -ForegroundColor Red
    }
}


