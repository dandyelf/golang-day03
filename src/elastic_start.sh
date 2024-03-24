#!/bin/bash

# OS := uname

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux
        $HOME/elasticsearch-8.12.2/bin/elasticsearch -E xpack.security.enabled=false
elif [[ "$OSTYPE" == "darwin"* ]]; then
        # Mac OSX
        $HOME/goinfre/elasticsearch-8.4.2/bin/elasticsearch -E xpack.security.enabled=false
else
        echo "Unknown OS type."
fi
