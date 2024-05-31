#!/bin/bash
set -eux

TRIVY_SCRIPT_PATH="trivy-scan.sh"
EXE_SCRIPT_PATH="vhd-scanning-exe-on-vm.sh"
TEST_RESOURCE_PREFIX="vhd-scanning"
VM_NAME="$TEST_RESOURCE_PREFIX-vm"
VHD_IMAGE="$MANAGED_SIG_ID"

SIG_CONTAINER_NAME="vhd-scans"
TEST_VM_ADMIN_USERNAME="azureuser"
STORAGE_ACCOUNT_NAME=$(echo $CLASSIC_BLOB | cut -d '/' -f 3 | cut -d '.' -f 1)

set +x
TEST_VM_ADMIN_PASSWORD="TestVM@$(date +%s)"
set -x


RESOURCE_GROUP_NAME="$TEST_RESOURCE_PREFIX-$(date +%s)-$RANDOM"
az group create --name $RESOURCE_GROUP_NAME --location ${AZURE_LOCATION} --tags 'source=AgentBaker'

# 18.04 VMs don't have access to new enough 'az' versions to be able to run the az commands in vhd-scanning-vm-exe.sh
if [ "$OS_VERSION" == "18.04" ]; then
    echo "Skipping scanning for 18.04"
    exit 0
fi

function cleanup() {
    echo "Deleting resource group ${RESOURCE_GROUP_NAME}"
    az group delete --name $RESOURCE_GROUP_NAME --yes --no-wait
}
trap cleanup EXIT

VM_OPTIONS="--size Standard_DS1_v2"
if [[ "${ARCHITECTURE,,}" == "arm64" ]]; then
    VM_OPTIONS="--size Standard_D2pds_v5"
elif [[ "${FEATURE_FLAGS,,}" == "kata" ]]; then
    VM_OPTIONS="--size Standard_D4ds_v5"
fi

if [[ "${OS_TYPE}" == "Linux" && "${ENABLE_TRUSTED_LAUNCH}" == "True" ]]; then
    VM_OPTIONS+=" --security-type TrustedLaunch --enable-secure-boot true --enable-vtpm true"
fi

az vm create --resource-group $RESOURCE_GROUP_NAME \
    --name $VM_NAME \
    --image $VHD_IMAGE \
    --admin-username $TEST_VM_ADMIN_USERNAME \
    --admin-password $TEST_VM_ADMIN_PASSWORD \
    --os-disk-size-gb 50 \
    --assign-identity "[system]" \
    --scope "/subscriptions/${SUBSCRIPTION_ID}/resourceGroups/${AZURE_RESOURCE_GROUP_NAME}/providers/Microsoft.Storage/storageAccounts/${STORAGE_ACCOUNT_NAME}/blobServices/default/containers/${SIG_CONTAINER_NAME}" \
    --role "Storage Blob Data Contributor" \
    ${VM_OPTIONS}

FULL_PATH=$(realpath $0)
CDIR=$(dirname $FULL_PATH)
TRIVY_SCRIPT_PATH="$CDIR/$TRIVY_SCRIPT_PATH"
az vm run-command invoke \
    --command-id RunShellScript \
    --name $VM_NAME \
    --resource-group $RESOURCE_GROUP_NAME \
    --scripts @$TRIVY_SCRIPT_PATH


TIMESTAMP=$(date +%s%3N)
TRIVY_REPORT_NAME="trivy-report-${BUILD_ID}-${TIMESTAMP}.json"
TRIVY_TABLE_NAME="trivy-table-${BUILD_ID}-${TIMESTAMP}.txt"
EXE_SCRIPT_PATH="$CDIR/$EXE_SCRIPT_PATH"
az vm run-command invoke \
    --command-id RunShellScript \
    --name $VM_NAME \
    --resource-group $RESOURCE_GROUP_NAME \
    --scripts @$EXE_SCRIPT_PATH \
    --parameters "OS_SKU=${OS_SKU}" \
        "OS_VERSION=${OS_VERSION}" \
        "TEST_VM_ADMIN_USERNAME=${TEST_VM_ADMIN_USERNAME}" \
        "ARCHITECTURE=${ARCHITECTURE}" \
        "TRIVY_REPORT_NAME=${TRIVY_REPORT_NAME}" \
        "TRIVY_TABLE_NAME=${TRIVY_TABLE_NAME}" \
        "SIG_CONTAINER_NAME"=${SIG_CONTAINER_NAME} \
        "STORAGE_ACCOUNT_NAME"=${STORAGE_ACCOUNT_NAME} \
        ENABLE_TRUSTED_LAUNCH=${ENABLE_TRUSTED_LAUNCH}


az storage blob download --container-name ${SIG_CONTAINER_NAME} --name  ${TRIVY_REPORT_NAME} --file trivy-report.json --account-name ${STORAGE_ACCOUNT_NAME} --auth-mode login
az storage blob download --container-name ${SIG_CONTAINER_NAME} --name  ${TRIVY_TABLE_NAME} --file  trivy-images-table.txt --account-name ${STORAGE_ACCOUNT_NAME} --auth-mode login

az storage blob delete --account-name ${STORAGE_ACCOUNT_NAME} --container-name ${SIG_CONTAINER_NAME} --name ${TRIVY_REPORT_NAME} --auth-mode login
az storage blob delete --account-name ${STORAGE_ACCOUNT_NAME} --container-name ${SIG_CONTAINER_NAME} --name ${TRIVY_TABLE_NAME} --auth-mode login