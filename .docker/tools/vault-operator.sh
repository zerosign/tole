#!/bin/sh

# https://www.vaultproject.io/api/system/health.html
# /sys/health
# - 200 if initialized, unsealed, and active
# - 429 if unsealed and standby
# - 472 if data recovery mode replication secondary and active
# - 473 if performance standby
# - 501 if not initialized
# - 503 if sealed

function check() {
    vault read
}
