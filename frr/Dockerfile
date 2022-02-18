FROM scratch AS frr
COPY --from=frrouting/frr:latest / /usr/local/lib/containers/frr/
COPY frr.yaml /rootfs/usr/local/etc/containers/frr.yaml
COPY manifest.yaml /manifest.yaml
