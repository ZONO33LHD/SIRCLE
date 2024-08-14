# スクリプトのディレクトリを取得
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
$manifestsDir = Join-Path $scriptDir ".." "manifests"

Write-Host "Current directory: $PWD"
Write-Host "Script directory: $scriptDir"
Write-Host "Manifests directory: $manifestsDir"

# マニフェストディレクトリの存在確認
if (-not (Test-Path $manifestsDir)) {
    Write-Host "Error: Manifests directory does not exist: $manifestsDir"
    exit 1
}

# Kubernetesクラスタに接続していることを確認
try {
    $clusterInfo = kubectl cluster-info
    if ($LASTEXITCODE -ne 0) {
        throw "Failed to get cluster info"
    }
    Write-Host "Successfully connected to Kubernetes cluster"
} catch {
    Write-Host "Error: Unable to connect to Kubernetes cluster. Please check your connection and authentication."
    Write-Host "Error details: $_"
    exit 1
}

# 各マニフェストファイルの適用順序
$manifestOrder = @(
    "namespace.yaml",
    "configmaps",
    "secrets",
    "storage",
    "deployments\postgres-deployment.yaml",
    "services\postgres-service.yaml",
    "deployments\kakeibo-user-service-deployment.yaml",
    "services\kakeibo-user-service-service.yaml"
)

foreach ($item in $manifestOrder) {
    $fullPath = Join-Path $manifestsDir $item
    if (Test-Path $fullPath -PathType Container) {
        # ディレクトリの場合、すべての.yamlファイルを適用
        Get-ChildItem $fullPath -Filter *.yaml | ForEach-Object {
            Write-Host "Applying $($_.FullName)"
            kubectl apply -f $_.FullName
            if ($LASTEXITCODE -ne 0) {
                Write-Host "Error applying $($_.FullName)"
                exit 1
            }
        }
    } elseif (Test-Path $fullPath -PathType Leaf) {
        # 単一のファイルの場合
        Write-Host "Applying $fullPath"
        kubectl apply -f $fullPath
        if ($LASTEXITCODE -ne 0) {
            Write-Host "Error applying $fullPath"
            exit 1
        }
    } else {
        Write-Host "Warning: Path not found: $fullPath"
    }
}

Write-Host "Deployment completed successfully!"
