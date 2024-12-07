param(
    [string]$venvPath,
    [string]$projectName
)

# Check if either path is empty
if ([string]::IsNullOrEmpty($venvPath) -or [string]::IsNullOrEmpty($projectName)) {
    Write-Error "Error: Both 'venvPath' and 'projectName' parameters are required."
    exit 1
}

& "$venvPath\$projectName\Scripts\activate.ps1"

$requirementsPath = "../requirements.txt"
Write-Host $requirementsPath

pip install -r $requirementsPath