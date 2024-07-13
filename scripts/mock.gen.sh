#!/usr/bin/env bash

set -euo pipefail

if ! [[ "$0" =~ scripts/mock.gen.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

# https://github.com/uber-go/mock
go install -v go.uber.org/mock/mockgen@v0.4.0

source ./scripts/constants.sh

outputted_files=()

# tuples of (source interface import path, comma-separated interface names, output file path)
input="scripts/mocks.mockgen.txt"
while IFS= read -r line
do
  IFS='=' read -r src_import_path interface_name output_path <<< "${line}"
  package_name="$(basename "$(dirname "$output_path")")"
  echo "Generating ${output_path}..."
  outputted_files+=("${output_path}")
  mockgen -package="${package_name}" -destination="${output_path}" "${src_import_path}" "${interface_name}"

done < "$input"

# tuples of (source import path, comma-separated interface names to exclude, output file path)
input="scripts/mocks.mockgen.source.txt"
while IFS= read -r line
do
  IFS='=' read -r source_path exclude_interfaces output_path <<< "${line}"
  package_name=$(basename "$(dirname "$output_path")")
  outputted_files+=("${output_path}")
  echo "Generating ${output_path}..."

  mockgen \
    -source="${source_path}" \
    -destination="${output_path}" \
    -package="${package_name}" \
    -exclude_interfaces="${exclude_interfaces}"

done < "$input"

mapfile -t all_generated_files < <(grep -Rl 'Code generated by MockGen. DO NOT EDIT.')

# Exclude certain files
outputted_files+=('scripts/mock.gen.sh') # This file
outputted_files+=('vms/components/lux/mock_transferable_out.go') # Embedded verify.IsState
outputted_files+=('vms/platformvm/fx/mock_fx.go') # Embedded verify.IsNotState

mapfile -t diff_files < <(echo "${all_generated_files[@]}" "${outputted_files[@]}" | tr ' ' '\n' | sort | uniq -u)

if (( ${#diff_files[@]} )); then
  printf "\nFAILURE\n"
  echo "Detected MockGen generated files that are not in scripts/mocks.mockgen.source.txt or scripts/mocks.mockgen.txt:"
  printf "%s\n" "${diff_files[@]}"
  exit 255
fi

echo "SUCCESS"
