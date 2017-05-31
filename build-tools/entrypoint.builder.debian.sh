#!/bin/bash

# Add local user
# Either use the LOCAL_USER_ID if passed in at runtime or
# fallback

USER_ID=${LOCAL_USER_ID:-9001}

echo "Starting with UID : $USER_ID"
useradd --create-home --shell /bin/bash --uid $USER_ID user
export HOME=/home/user

exec gosu user "$@"
