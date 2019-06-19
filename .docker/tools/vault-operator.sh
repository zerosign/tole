#!/bin/bash

STATUS="000"
SECRETS=()
SECRET_PATH=""
ROOT_TOKEN=""

# check whether $SECRET_PATH variable exists or not
if [ -z "$SECRET_PATH" ]; then
    SECRET_PATH=$(pwd)/.docker/volumes/vault/secrets
fi

# check whether SECRET_PATH exists or not
# if didn't exist, create it
if [ ! -d "$SECRET_PATH" ]; then
    mkdir -p $SECRET_PATH
fi

# https://www.vaultproject.io/api/system/health.html
# /sys/health
# - 200 if initialized, unsealed, and active
# - 429 if unsealed and standby
# - 472 if data recovery mode replication secondary and active
# - 473 if performance standby
# - 501 if not initialized
# - 503 if sealed
#
# return status code only
#
function status() {
    STATUS=$(curl -s -o /dev/null -w "%{http_code}" "$1/v1/sys/health")
}

#
# fetch secrets from files
#
function secrets() {
    SECRETS=($(cat $SECRET_PATH/secret.keys))
}

function root() {
    secrets
    ROOT_TOKEN=${SECRETS[7]}
}

#
# unseal vault from secret.keys
#
function unseal() {
    vault operator unseal -address=$1 ${SECRETS[0]}
    vault operator unseal -address=$1 ${SECRETS[1]}
    vault operator unseal -address=$1 ${SECRETS[2]}
}

#
# call vault operator init then dump
# secrets into /secrets/secret.keys
#
# formats (separated by newline) :
# secrets[0]
# secrets[1]
# ...
# root_token
#
function init() {
    vault operator init -address=$1 | awk -F':' '{ if ($2 != "") { gsub(/[ \\t]+/, "", $2); print $2 } }' > \
                                          "$SECRET_PATH/secret.keys"
}

# function auto-init() {
#     # init if only vault aren't being intialized
#     local status =''
#     status = status()
#     if [ status = "501" ]; then
#         init()
#         unseal()
#     fi

#     if [ status = "503" ]; then

#     fi
# }

case $1 in
    status)
        status $2
        echo $STATUS
        ;;
    secrets)
        secrets
        printf '%s\n' "${SECRETS[@]}"
        ;;
    init)
        init $2
        ;;
    unseal)
        secrets
        unseal $2 $SECRETS
        ;;
    auto-init)
        auto-init
        ;;
    help)
        echo -n "vault-operator [status|secrets|init|unseal|auto-init]"
        ;;
    *)
        echo -n "unknown action"
esac
