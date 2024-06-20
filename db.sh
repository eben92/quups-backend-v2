#!/bin/bash

echo "running script ${APP_ENV}"
make db-migrate-up
echo "i am done"