#!/bin/bash
source .env

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${LOCAL_MIGRATION_DSN}" up -v