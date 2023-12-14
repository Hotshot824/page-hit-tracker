#! /bin/bash

# Create a user with the same UID as the host user
if ! id "$DOCKER_USER" &>/dev/null; then
    useradd -s /bin/bash $DOCKER_USER
    chown -R 1000:1000 /home/$DOCKER_USER > /dev/null 2>&1
fi

if ! which sqlite3 >/dev/null; then
    apt-get update && apt-get install -y sqlite3
fi

if ! test -f ./db/count.db; then
    (
        cd ./db
        su "$DOCKER_USER" --c "sqlite3 ./count.db < ./init.sql"
    )
fi

su "$DOCKER_USER" --c "cd /home/$DOCKER_USER && make build && make run"