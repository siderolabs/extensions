# syntax=docker/dockerfile:1.3

FROM scratch AS iscsid
COPY --from=ghcr.io/talos-systems/kmod:v0.9.0 / /rootfs
COPY --from=ghcr.io/talos-systems/libressl:v0.9.0 / /rootfs
COPY --from=ghcr.io/talos-systems/musl:v0.9.0 / /rootfs
COPY --from=ghcr.io/talos-systems/open-isns:v0.9.0 / /rootfs
COPY --from=ghcr.io/talos-systems/util-linux:v0.9.0 / /rootfs
COPY --from=ghcr.io/talos-systems/open-iscsi:v0.9.0 / /rootfs
COPY iscsid.yaml /rootfs/usr/local/etc/containers/iscsid.yaml
COPY manifest.yaml /manifest.yaml