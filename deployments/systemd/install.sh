#!/bin/bash

set -e -o pipefail

print_sep() {
    local cols line
    cols=${COLUMNS:-$(tput cols 2>/dev/null || echo 80)}
    printf -v line '%*s' "$cols" ''
    printf '%s\n' "${line// /-}"
}

echo "Starting TeleIRC installation script..."
print_sep

echo "Checking if systemd is the init system..."
check_systemd_is_init() {
    if [ "$(basename "$(readlink /proc/1/exe)")" = "systemd" ]; then
        return 0
    else
        return 1
    fi
}
if ! check_systemd_is_init; then
    echo "Error: This script requires systemd as the init system." >&2
    exit 1
else
    echo "systemd is the init system. Continuing..."
fi
print_sep

echo "Checking if sudo and curl is installed..."
if ! command -v sudo &> /dev/null; then
    echo "sudo could not be found. Please install sudo and re-run this script." >&2
    exit 1
else
    echo "sudo is installed. Continuing..."
fi
if ! command -v curl &> /dev/null; then
    echo "curl could not be found. Please install curl and re-run this script." >&2
    exit 1
else
    echo "curl is installed. Continuing..."
fi
print_sep

echo "Fetching the latest TeleIRC version tag from Anitya API..."
ANITYA_API_URL="https://release-monitoring.org/api/v2/versions/?project_id=386778"
LATEST_VERSION=$(curl -sL "$ANITYA_API_URL" | sed -n 's/.*"latest_version":"\([^"]*\)".*/\1/p')
if [[ -z "$LATEST_VERSION" ]]; then
    echo "Error: Could not retrieve or parse the latest version tag of TeleIRC from Anitya API." >&2
    exit 1
fi
VER=$LATEST_VERSION
echo "Selecting TeleIRC version ${VER}..."
print_sep

echo "Creating temporary working directory..."
TMP_DIR=$(mktemp -d)
if [ ! -d "$TMP_DIR" ]; then
    echo "Error: Failed to create temporary directory." >&2
    exit 1
else 
    echo "Temporary directory created at ${TMP_DIR}."
fi
print_sep

cleanup() {
    if [ -n "${TMP_DIR:-}" ] && [ -d "$TMP_DIR" ]; then
        echo "Cleaning up temporary files from ${TMP_DIR} ..."
        rm -rvf "$TMP_DIR"
    fi
}
trap cleanup EXIT INT

GITHUB_RELEASE_URL="https://github.com/RITlug/teleirc/releases/download/v${VER}"
GITHUB_RAW_URL="https://raw.githubusercontent.com/RITlug/teleirc/v${VER}"
echo "Set GITHUB_RELEASE_URL to ${GITHUB_RELEASE_URL}"
echo "Set GITHUB_RAW_URL to ${GITHUB_RAW_URL}" 
print_sep

echo "Downloading TeleIRC deployment assets from GitHub..."
curl --verbose --location --output "${TMP_DIR}"/teleirc "${GITHUB_RELEASE_URL}"/teleirc-${VER}-linux-x86_64
curl --verbose --location --output "${TMP_DIR}"/teleirc.sysusers "${GITHUB_RAW_URL}"/deployments/systemd/teleirc.sysusers
curl --verbose --location --output "${TMP_DIR}"/teleirc.tmpfiles "${GITHUB_RAW_URL}"/deployments/systemd/teleirc.tmpfiles
curl --verbose --location --output "${TMP_DIR}"/teleirc@.service "${GITHUB_RAW_URL}"/deployments/systemd/teleirc@.service
curl --verbose --location --output "${TMP_DIR}"/teleirc.env "${GITHUB_RAW_URL}"/env.example
print_sep

if [ -z "$EDITOR" ]; then
    if command -v vim &> /dev/null; then
        EDITOR="vim"
    else
        EDITOR="nano"
    fi
fi
echo "Selected EDITOR as ${EDITOR} ..."
print_sep

echo "Please edit the TeleIRC environment file to configure your bridge settings..."
$EDITOR "${TMP_DIR}"/teleirc.env
print_sep

echo "Please make sure you configured TeleIRC correctly in ${TMP_DIR}/teleirc.env before proceeding!"
echo "Press any key to continue..."
read -n 1 -s -r
print_sep

echo "Installing TeleIRC files and user..."
sudo install -v -Dm755 -o root -g root "${TMP_DIR}"/teleirc /usr/local/bin/teleirc
sudo install -v -Dm644 -o root -g root "${TMP_DIR}"/teleirc.sysusers /etc/sysusers.d/teleirc.conf
sudo install -v -Dm644 -o root -g root "${TMP_DIR}"/teleirc.tmpfiles /etc/tmpfiles.d/teleirc.conf
sudo install -v -Dm644 -o root -g root "${TMP_DIR}"/teleirc@.service /etc/systemd/system/teleirc@.service
sudo systemd-sysusers /etc/sysusers.d/teleirc.conf
sudo systemd-tmpfiles --create /etc/tmpfiles.d/teleirc.conf
sudo install -v -Dm644 -o root -g root "${TMP_DIR}"/teleirc.env /etc/teleirc/bridge
print_sep

echo "Checking if SELinux is installed and enforced..."
check_selinux_status() {
    if ! command -v getenforce &> /dev/null; then
        return 1
    fi
    SELINUX_STATUS=$(getenforce)
    case "$SELINUX_STATUS" in
        "Enforcing")
            return 0
            ;;
        *)
            return 1
            ;;
    esac
}
if check_selinux_status; then
    echo "SELinux is enforced. Setting SELinux context for TeleIRC files..."
    sudo chcon --type bin_t --user system_u /usr/local/bin/teleirc
    sudo chcon --type etc_t --user system_u -R /etc/teleirc/
    sudo chcon --type etc_t --user system_u /etc/sysusers.d/teleirc.conf
    sudo chcon --type etc_t --user system_u /etc/tmpfiles.d/teleirc.conf
    sudo chcon --type systemd_unit_file_t --user system_u /etc/systemd/system/teleirc@.service
else
    echo "SELinux is not enforcing or not installed. We have no need to set it up."
fi
print_sep

echo "Enabling TeleIRC systemd service unit..."
sudo systemctl enable teleirc@bridge.service
echo "Enabled TeleIRC systemd service unit."
print_sep

read -p "Do you want to start TeleIRC now? (y/n): " response
if [[ "$response" =~ ^[Yy]$ ]]; then
    echo "Starting TeleIRC..."
    sudo systemctl start teleirc@bridge.service
else
    echo "Skipping starting TeleIRC..."
fi
print_sep

echo "Installation of TeleIRC version ${VER} is completed."
