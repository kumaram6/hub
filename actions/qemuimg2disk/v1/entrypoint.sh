#!/bin/sh
set -o errexit
set -o xtrace
# We actually do want EXTRA_ARGS to be split up
#shellcheck disable=SC2086
curl "${IMG_URL:?}" -O jammy-server-cloudimg-amd64.img
qemu-img convert jammy-server-cloudimg-amd64.img -O "${FORMAT:-host_device}" "${DEST_DISK:?}" ${EXTRA_ARGS}
partprobe || true
