#!/bin/bash

if [ $# -eq 0 ]
then
    echo "usage: ./dev.sh [app]"
    echo "  Example: ./dev.sh admin"
    exit 1
fi

fresh -c ./scripts/${1}_api_runner.conf