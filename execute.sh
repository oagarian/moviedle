#!/bin/bash

function select_environment() {
  local environment="$1"

  case "$environment" in
      -database)
          run_database;;
      *)
          echo "Invalid option for environment generation. Possible arguments: -environment <-development|-test|-all>"
          return 1;;
  esac
}

function run_database() {
    chmod +x ./tools/database.sh
    ./tools/database.sh
}

case "$1" in
    -environment)
        shift
        select_environment "$@";;
    *)
       echo "Invalid option. Possible arguments: ./execute.sh <-environment|-mockgen>";;
esac