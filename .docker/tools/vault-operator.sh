#!/bin/bash

SECRET_PATH=$SECRET_PATH
POLICY_PATH=$POLICY_PATH
STATUS="000"
SECRETS=()
ROOT_TOKEN=""
TOKEN=""

# check whether $SECRET_PATH variable exists or not
if [ -z "$SECRET_PATH" ]; then
    SECRET_PATH=$(pwd)/.docker/volumes/vault/secrets
fi

# check whether SECRET_PATH exists or not
# if didn't exist, create it
if [ ! -d "$SECRET_PATH" ]; then
    mkdir -p $SECRET_PATH
fi


if [ -z "$POLICY_PATH" ]; then
    POLICY_PATH=$(pwd)/.docker/volumes/vault/policies
fi

if [ ! -d "$POLICY_PATH" ]; then
    mkdir -p $POLICY_PATH
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
    ROOT_TOKEN=${SECRETS[5]}
}

#
# unseal vault from secret.keys
#
function unseal() {
    vault operator unseal -address=$1 ${SECRETS[0]}
    vault operator unseal -address=$1 ${SECRETS[1]}
    vault operator unseal -address=$1 ${SECRETS[2]}
}

function login() {
    vault login -address=$1 $2
    TOKEN=$(vault token create -address=$1 -format=json | jq -r '.auth.client_token ' -c)
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
    # vault operator init -address=$1 > $SECRET_PATH/output.raw
    # cat $SECRET_PATH/output.raw | awk -F':' '{ if ($2 != "") { gsub(/[ \t]+/, "", $2); print $2 } }' > "$SECRET_PATH/secret.keys"
    vault operator init -address=$1 | awk -F':' '{ if ($2 != "") { gsub(/[ \t]+/, "", $2); print $2 } }' > "$SECRET_PATH/secret.keys"
}

function policy-init() {
    POLICES=($(find $POLICY_PATH -name "*.hcl" -type f))
    for path in "${POLICES[@]}"; do
        echo $path
        policy=$(echo $path | awk -F"/" '{ print $NF }' | tr "_" " " | tr "." " " | awk -F" " '{ print $2 }')
        policy-exists $1 $policy
        if [ "$POLICY_EXISTS" = "404" ] || [ "$POLICY_EXISTS" = "400" ]; then
            vault policy write -address $1 $policy $path
        fi
    done
}

function policy-exists() {
    TOKEN=$(vault print token)
    POLICY_EXISTS=$(curl -s -o /dev/null -w "%{http_code}" -H "X-Vault-Token: $TOKEN" $1/v1/sys/policies/acl/$2)
}

function auto-init() {
    status $1

    case $STATUS in
        "503")
            secrets
            root
            unseal $1 $SECRETS
            login $1 $ROOT_TOKEN
            policy-init $1
            ;;
        "501")
            init $1
            secrets
            root
            unseal $1 $SECRETS
            echo "root-token" $ROOT_TOKEN
            login $1 $ROOT_TOKEN
            policy-init $1
            ;;
        *)
            echo -n "everything is good :))"
    esac
}

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
    login)
        root
        login $2 $ROOT_TOKEN
        ;;
    policy-init)
        root
        policy-init $2
        ;;
    root)
        root
        echo $ROOT_TOKEN
        ;;
    policy-exists)
        root
        policy-exists $2 $3
        echo $POLICY_EXISTS
        ;;
    unseal)
        secrets
        unseal $2 $SECRETS
        ;;
    auto-init)
        auto-init $2
        ;;
    help)
        echo -n "vault-operator [status|secrets|init|unseal|auto-init]"
        ;;
    *)
        echo -n "unknown action"
esac
