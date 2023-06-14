#!/bin/bash
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.

while (( "$#" )); do
    case "$1" in
        -s|--skip_apply)
            skip_apply="yes"
            shift
            ;;
        -l|--run_live)
            run_live="yes"
            shift
            ;;
        -m|--marketplace)
            marketplace="yes"
            shift
            ;;

        *)
            echo "Invalid arg '$1'"
            exit
            ;;
    esac
done


script_dir="$(dirname "$(realpath "$0")")"

mkdir "$script_dir/apply_logs"
echo "Apply scripts" 1>&2
touch "$script_dir/fail.txt"
touch "$script_dir/success.txt"
touch $script_dir/failure_details.txt


for script in $(find "$script_dir/rhel8" -name '*.sh' | sort -u); do
    scriptname="$(basename "${script}")"
    prunedname=$(echo "${scriptname#*-}" | cut -d'.' -f1)



    echo "checking '${prunedname}'"  1>&2
    if grep -q -E "^${prunedname}\$" "$script_dir/skip_list.txt" ; then
        # If we are running live scripts, run those anyways
        if [[ "${run_live}" == "yes" ]]; then
            if ! grep -q -E "^${prunedname}\$" "$script_dir/live_machine_only.txt" ; then
                echo "Skipping ${script} since its in skip_list.txt but not also in live_machine_only.txt" 1>&2
                continue
            fi
        else
            echo "Skipping ${script} since its in skip_list.txt" 1>&2
            continue
        fi
    fi

    if [[ "${marketplace}" == "yes" ]]; then
    if grep -q -E "^${prunedname}\$" "$script_dir/marketplace_skip_list.txt" ; then
        echo "Skipping ${script} since its in marketplace_skip_list.txt" 1>&2
        continue
    fi
    fi

done

