#!/bin/bash

set -e

BACKUP_DIR="dunmain_backup"
rm -rf $BACKUP_DIR || true

mv dunmain/ ${BACKUP_DIR}/

cp -r sing-box/cmd/sing-box/ dunmain

DIR="dunmain"

copy_back() {
    cp "${BACKUP_DIR}/${1}" "${DIR}/"
}

for file in "$DIR"/*; do
    # 检查文件是否为普通文件
    if [ -f "$file" ]; then
        first_line=$(head -n 1 "$file")

        if [[ "$first_line" == *main* ]]; then
            sed -i '1s/main/dunmain/' "$file"
            echo "Changed: $file"
        else
            echo "$file not belongs to 'main'"
        fi
    fi
done

# IGNORE_LIST=(
#     "main.go"
#     "color.go"
#     "cmd_run.go"
# )

# for IGNORE_FILE in "${IGNORE_LIST[@]}"; do
#     copy_back $IGNORE_FILE
# done

echo ""
echo "CHECK IT PLEASE!"
