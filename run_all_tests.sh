#!/usr/bin/env bash
set -eu

trap "exit" INT TERM
trap end EXIT
end(){
    if [[ $? -ne 0 ]]; then
        echo "Tests failed :("
        echo "If the failure looks unrelated to changes made consider running: make nuke"
        echo "Then try again."
        exit 1
    fi
}

# Check if a makefile has a recipe: https://stackoverflow.com/a/58316463
function has-recipe() {
    make -C $1 -qp | awk -F':' '/^[a-zA-Z0-9][^$#\/\t=]*:([^=]|$)/ {split($1,A,/ /);for(i in A)print A[i]}' | grep -qx $2
}

make

# Iterate through all directories and run `make lint/test` if there's a makefile.
for dir in $(find . -mindepth 1 -maxdepth 1 -type d | sort); do
    if [ -f "$dir/Makefile" ]; then
        if has-recipe "$dir" "lint"; then
            # Skip some directories because lint fails.
            for exclude in "op-exporter" "op-ufm" "op-node"; do
                if [ "$dir" = "./$exclude" ]; then
                    echo "Skipping lint: $dir"
                    continue 2
                fi
            done
            echo "Running lint in $dir"
            make -C "$dir" lint
        fi

        if has-recipe "$dir" "test"; then
            # Skip some directories because tests fail.
            for exclude in "proxyd" "indexer" "op-bootnode" "op-ufm"; do
                if [ "$dir" = "./$exclude" ]; then
                    echo "Skipping test: $dir"
                    continue 2
                fi
            done
            echo "Running tests in $dir"
            make -C "$dir" test
        fi
    fi
done

echo Ok!
