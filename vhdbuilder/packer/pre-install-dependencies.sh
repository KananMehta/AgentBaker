#!/bin/bash

script_start_timestamp=$(date +%H:%M:%S)
section_start_timestamp=$(date +%H:%M:%S)

script_start_stopwatch=$(date +%s)
section_start_stopwatch=$(date +%s)

declare -a benchmarks=()
declare -a jsonBenchmarks=()

OS=$(sort -r /etc/*-release | gawk 'match($0, /^(ID_LIKE=(coreos)|ID=(.*))$/, a) { print toupper(a[2] a[3]); exit }')
OS_VERSION=$(sort -r /etc/*-release | gawk 'match($0, /^(VERSION_ID=(.*))$/, a) { print toupper(a[2] a[3]); exit }' | tr -d '"')
THIS_DIR="$(cd "$(dirname ${BASH_SOURCE[0]})" && pwd)"

#the following sed removes all comments of the format {{/* */}}
sed -i 's/{{\/\*[^*]*\*\/}}//g' /home/packer/provision_source.sh
sed -i 's/{{\/\*[^*]*\*\/}}//g' /home/packer/tool_installs_distro.sh

source /home/packer/provision_installs.sh
source /home/packer/provision_installs_distro.sh
source /home/packer/provision_source.sh
source /home/packer/provision_source_distro.sh
source /home/packer/tool_installs.sh
source /home/packer/tool_installs_distro.sh
source /home/packer/packer_source.sh

sudo snap install jq
export PATH=$PATH:/snap/bin
jq --version
capture_benchmarks false "declare_variables_remove_comments_and_execute_packer_files"
start_watch

CPU_ARCH=$(getCPUArch)  #amd64 or arm64
VHD_LOGS_FILEPATH=/opt/azure/vhd-install.complete
COMPONENTS_FILEPATH=/opt/azure/components.json
MANIFEST_FILEPATH=/opt/azure/manifest.json
KUBE_PROXY_IMAGES_FILEPATH=/opt/azure/kube-proxy-images.json
#this is used by post build test to check whether the compoenents do indeed exist
cat components.json > ${COMPONENTS_FILEPATH}
cat manifest.json > ${MANIFEST_FILEPATH}
cat ${THIS_DIR}/kube-proxy-images.json > ${KUBE_PROXY_IMAGES_FILEPATH}
echo "Starting build on " $(date) > ${VHD_LOGS_FILEPATH}
capture_benchmarks false "create_post_build_test"
start_watch

if [[ $OS == $MARINER_OS_NAME ]]; then
  chmod 755 /opt
  chmod 755 /opt/azure
  chmod 644 ${VHD_LOGS_FILEPATH}
fi
capture_benchmarks false "set_permissions_if_mariner"
start_watch

copyPackerFiles
systemctlEnableAndStart disk_queue || exit 1
capture_benchmarks false "copy_packer_files"
start_watch

mkdir /opt/certs
chmod 1666 /opt/certs
systemctlEnableAndStart update_certs.path || exit 1
capture_benchmarks false "update_certs_make_directory_and_set_permissions"
start_watch

systemctlEnableAndStart ci-syslog-watcher.path || exit 1
systemctlEnableAndStart ci-syslog-watcher.service || exit 1

# enable AKS log collector
echo -e "\n# Disable WALA log collection because AKS Log Collector is installed.\nLogs.Collect=n" >> /etc/waagent.conf || exit 1
systemctlEnableAndStart aks-log-collector.timer || exit 1
capture_benchmarks false "start_sys_logs_and_aks_log_collector"
start_watch

# enable the modified logrotate service and remove the auto-generated default logrotate cron job if present
systemctlEnableAndStart logrotate.timer || exit 1
rm -f /etc/cron.daily/logrotate
capture_benchmarks false "enable_modified_log_rotate_service"
start_watch

systemctlEnableAndStart sync-container-logs.service || exit 1
capture_benchmarks false "sync_container_logs"
start_watch

# First handle Mariner + FIPS
if [[ ${OS} == ${MARINER_OS_NAME} ]]; then
  dnf_makecache || exit $ERR_APT_UPDATE_TIMEOUT
  dnf_update || exit $ERR_APT_DIST_UPGRADE_TIMEOUT
  if [[ "${ENABLE_FIPS,,}" == "true" ]]; then
    # This is FIPS install for Mariner and has nothing to do with Ubuntu Advantage
    echo "Install FIPS for Mariner SKU"
    installFIPS
  fi
else
  # Enable ESM for 18.04 and FIPS only
  if [[ "${UBUNTU_RELEASE}" == "18.04" ]] || [[ "${ENABLE_FIPS,,}" == "true" ]]; then
    autoAttachUA
  fi

  # Run apt get update to refresh repo list
  # Run apt dist get upgrade to install packages/kernels

  # CVM breaks on kernel image updates due to nullboot package post-install.
  # it relies on boot measurements from real tpm hardware.
  # building on a real CVM would solve this, but packer doesn't support it.
  # we could make upstream changes but that takes time, and we are broken now.
  # so we just hold the kernel image packages for now on CVM.
  # this still allows us base image and package updates on a weekly cadence.
  if [[ "$IMG_SKU" != "20_04-lts-cvm" ]]; then
    apt_get_update || exit $ERR_APT_UPDATE_TIMEOUT
    apt_get_dist_upgrade || exit $ERR_APT_DIST_UPGRADE_TIMEOUT    
  fi

  if [[ "${ENABLE_FIPS,,}" == "true" ]]; then
    # This is FIPS Install for Ubuntu, it purges non FIPS Kernel and attaches UA FIPS Updates
    echo "Install FIPS for Ubuntu SKU"
    installFIPS
  fi
fi
capture_benchmarks false "handle_mariner_and_fips_configurations"
start_watch

# Handle Azure Linux + CgroupV2
if [[ ${OS} == ${MARINER_OS_NAME} ]] && [[ "${ENABLE_CGROUPV2,,}" == "true" ]]; then
  enableCgroupV2forAzureLinux
fi

if [[ "${UBUNTU_RELEASE}" == "22.04" && "${ENABLE_FIPS,,}" != "true" ]]; then
  echo "Logging the currently running kernel: $(uname -r)"
  echo "Before purging kernel, here is a list of kernels/headers installed:"; dpkg -l 'linux-*azure*'

  # Purge all current kernels and dependencies
  DEBIAN_FRONTEND=noninteractive apt-get remove --purge -y $(dpkg-query -W 'linux-*azure*' | awk '$2 != "" { print $1 }' | paste -s)
  echo "After purging kernel, dpkg list should be empty"; dpkg -l 'linux-*azure*'

  # Install lts-22.04 kernel
  DEBIAN_FRONTEND=noninteractive apt-get install -y linux-image-azure-lts-22.04 linux-cloud-tools-azure-lts-22.04 linux-headers-azure-lts-22.04 linux-modules-extra-azure-lts-22.04 linux-tools-azure-lts-22.04
  echo "After installing new kernel, here is a list of kernels/headers installed"; dpkg -l 'linux-*azure*'
  
  update-grub
fi
capture_benchmarks false "handle_azureLinux_and_cgroupV2"

echo "pre-install-dependencies step finished successfully"