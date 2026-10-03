#!/bin/sh
set -e

AWG_INTERFACE="${AWG_INTERFACE:-awg0}"
AWG_CONFIG="${AWG_CONFIG:-/etc/amneziawg/${AWG_INTERFACE}.conf}"
AWG_LOG_LEVEL="${AWG_LOG_LEVEL:-info}"

export WG_SOCKET_DIR=/run/amneziawg

log() { echo "[amneziawg] $*"; }

if [ ! -f "$AWG_CONFIG" ]; then
  log "ERROR: config file not found: $AWG_CONFIG"
  log "Provide config via ExtensionServiceConfig:"
  log ""
  log "  apiVersion: v1alpha1"
  log "  kind: ExtensionServiceConfig"
  log "  name: amneziawg"
  log "  configFiles:"
  log "    - content: |"
  log "        [Interface]"
  log "        PrivateKey = <key>"
  log "        ListenPort = 51820"
  log "        Address = 10.0.0.1/24"
  log "        Jc = 5"
  log "        Jmin = 50"
  log "        Jmax = 1000"
  log "        S1 = 0"
  log "        S2 = 0"
  log "        H1 = 1"
  log "        H2 = 2"
  log "        H3 = 3"
  log "        H4 = 4"
  log "        [Peer]"
  log "        PublicKey = <key>"
  log "        Endpoint = <host>:51820"
  log "        AllowedIPs = 0.0.0.0/0"
  log "      mountPath: /etc/amneziawg/awg0.conf"
  exit 1
fi

parse_field() {
  grep -i "^${1} *=" "$2" | head -1 | sed "s/^[^ ]* *= *//"
}

collect_field() {
  grep -i "^${1} *=" "$2" | sed "s/^[^ ]* *= *//" | tr '\n' ' '
}

ADDRESSES=$(collect_field Address "$AWG_CONFIG")
MTU=$(parse_field MTU "$AWG_CONFIG")

STRIPPED_CONF=$(mktemp)
grep -iv "^\(Address\|DNS\|MTU\|Table\|PreUp\|PostUp\|PreDown\|PostDown\|SaveConfig\) *=" "$AWG_CONFIG" > "$STRIPPED_CONF"

log "Starting amneziawg-go on interface $AWG_INTERFACE"
export LOG_LEVEL="$AWG_LOG_LEVEL"

mkdir -p /run/amneziawg
amneziawg-go "$AWG_INTERFACE" &
AWG_PID=$!

SOCK="/run/amneziawg/${AWG_INTERFACE}.sock"
for i in $(seq 1 50); do
  [ -S "$SOCK" ] && break
  sleep 0.1
done

if [ ! -S "$SOCK" ]; then
  log "ERROR: UAPI socket did not appear at $SOCK"
  kill $AWG_PID 2>/dev/null
  exit 1
fi

awg setconf "$AWG_INTERFACE" "$STRIPPED_CONF"
rm -f "$STRIPPED_CONF"
log "Configuration applied from $AWG_CONFIG"

for addr in $ADDRESSES; do
  ip addr add "$addr" dev "$AWG_INTERFACE" 2>/dev/null || true
  log "Address $addr added"
done

if [ -n "$MTU" ]; then
  ip link set mtu "$MTU" dev "$AWG_INTERFACE"
  log "MTU set to $MTU"
fi

ip link set "$AWG_INTERFACE" up
log "Interface $AWG_INTERFACE is UP"

POST_UP=$(grep -i "^PostUp *=" "$AWG_CONFIG" | sed "s/^[^ ]* *= *//" || true)
if [ -n "$POST_UP" ]; then
  log "Running PostUp"
  eval "$POST_UP" || true
fi

cleanup() {
  log "Shutting down $AWG_INTERFACE"
  POST_DOWN=$(grep -i "^PostDown *=" "$AWG_CONFIG" | sed "s/^[^ ]* *= *//" || true)
  if [ -n "$POST_DOWN" ]; then
    log "Running PostDown"
    eval "$POST_DOWN" || true
  fi
  ip link set "$AWG_INTERFACE" down 2>/dev/null || true
  ip link delete "$AWG_INTERFACE" 2>/dev/null || true
  kill $AWG_PID 2>/dev/null
}
trap cleanup TERM INT

log "AmneziaWG is running:"
awg show "$AWG_INTERFACE"

wait $AWG_PID
