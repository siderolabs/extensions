#!/bin/bash
# FRR Startup Script for Talos
# This script sets up the MetalLB VRF, veth pair, and generates frr.conf before starting FRR.

set -e

log() { echo "$(date '+%Y-%m-%d %H:%M:%S') [frr-startup] $*" >&2; }
error() { echo "$(date '+%Y-%m-%d %H:%M:%S') [frr-startup] ERROR: $*" >&2; exit 1; }

# Default values
ASN_VRF_METALLB="${ASN_VRF_METALLB:-4000099998}"
ASN_METALLB_SPEAKER="${ASN_METALLB_SPEAKER:-4000099999}"
VRF_METALLB="${VRF_METALLB:-metallb}"
METALLB_VRF_ROUTE_TABLE_ID="${METALLB_VRF_ROUTE_TABLE_ID:-88}"
PEER_IP_FRR="${PEER_IP_FRR:-fda1:b2c3:d4e5:0001::0}"
PEER_IP_METALLB="${PEER_IP_METALLB:-fda1:b2c3:d4e5:0001::1}"
PEER_IP_PREFIX="${PEER_IP_PREFIX:-127}"
ENABLE_BFD="${ENABLE_BFD:-true}"
FRR_PROFILE="${FRR_PROFILE:-datacenter}"
# INTERFACE_MTU - auto-detected from first fabric interface if not set

# Validate mandatory variables
[ -z "$NODE_IP" ] && error "NODE_IP is required"
[ -z "$ASN_LOCAL" ] && error "ASN_LOCAL is required"

# Validate: either FE_MACS or FE_PORT_NAlog syslogMES is required (not both)
if [ -n "$FE_MACS" ] && [ -n "$FE_PORT_NAMES" ]; then
    error "Specify either FE_MACS or FE_PORT_NAMES, not both"
elif [ -z "$FE_MACS" ] && [ -z "$FE_PORT_NAMES" ]; then
    error "Either FE_MACS or FE_PORT_NAMES is required"
fi

# Validate FRR_PROFILE
case "$FRR_PROFILE" in
    datacenter|traditional) ;;
    *) error "FRR_PROFILE must be 'datacenter' or 'traditional', got: $FRR_PROFILE" ;;
esac

# Resolve MAC addresses to interface names
# Input: comma-separated MAC addresses
# Output: comma-separated interface names
resolve_interfaces() {
    local macs="$1"
    local interfaces=""

    IFS=',' read -ra MAC_ARRAY <<< "$macs"
    for mac in "${MAC_ARRAY[@]}"; do
        mac=$(echo "$mac" | tr '[:upper:]' '[:lower:]' | xargs)
        iface=$(/sbin/ip -o link | awk -v mac="$mac" 'tolower($0) ~ mac {print $2}' | tr -d ':')
        if [ -n "$iface" ]; then
            interfaces="${interfaces:+$interfaces,}$iface"
            log "Resolved MAC $mac -> $iface"
        else
            log "WARNING: No interface found for MAC $mac"
        fi
    done
    echo "$interfaces"
}

# Create MetalLB VRF
create_metallb_vrf() {
    log "Creating VRF $VRF_METALLB with table $METALLB_VRF_ROUTE_TABLE_ID..."

    if ! /sbin/ip link show "$VRF_METALLB" &>/dev/null; then
        /sbin/ip link add "$VRF_METALLB" type vrf table "$METALLB_VRF_ROUTE_TABLE_ID"
    fi
    /sbin/ip link set "$VRF_METALLB" up
    log "VRF $VRF_METALLB is up"
}

# Create veth pair for MetalLB
# frr-metallb: in VRF metallb (FRR listens here)
# metallb-frr: in host namespace (MetalLB speaker connects from here)
create_metallb_veth() {
    local veth_frr="veth-frr"
    local veth_metallb="veth-metallb"

    log "Creating veth pair: $veth_frr <-> $veth_metallb..."
    /sbin/ip link add "$veth_frr" type veth peer name "$veth_metallb" 2>/dev/null || log "Veth pair $veth_frr <-> $veth_metallb already exists"

    log "Assining $veth_frr to VRF $VRF_METALLB..."
    /sbin/ip link set "$veth_frr" master "$VRF_METALLB"

    # Set MTU with explicit error handling
    log "Setting MTU $INTERFACE_MTU on $veth_frr..."
    /sbin/ip link set "$veth_frr" mtu "$INTERFACE_MTU"

    log "Setting MTU $INTERFACE_MTU on $veth_metallb..."
    /sbin/ip link set "$veth_metallb" mtu "$INTERFACE_MTU"

    # Assign IPv6 addresses
    # PEER_IP_FRR is on frr-metallb (in VRF) - FRR listens here
    # PEER_IP_METALLB is on metallb-frr (in host ns) - MetalLB binds here
    log "Assigning IPv6 addresses..."
    /sbin/ip -6 addr add "${PEER_IP_FRR}/${PEER_IP_PREFIX}" dev "$veth_frr" 2>/dev/null || log "IPv6 $PEER_IP_FRR already assigned to $veth_frr"
    /sbin/ip -6 addr add "${PEER_IP_METALLB}/${PEER_IP_PREFIX}" dev "$veth_metallb" 2>/dev/null || log "IPv6 $PEER_IP_METALLB already assigned to $veth_metallb"

    # Bring up interfaces
    log "Bringing up interfaces..."
    /sbin/ip link set "$veth_frr" up
    /sbin/ip link set "$veth_metallb" up

    log "MetalLB veth pair configured:"
    log "  $veth_frr ($PEER_IP_FRR) in VRF $VRF_METALLB (FRR listens)"
    log "  $veth_metallb ($PEER_IP_METALLB) in host ns (MetalLB connects from)"
}

# Configure daemons based on ENABLE_BFD
configure_daemons() {
    if [ "$ENABLE_BFD" = "true" ]; then
        log "BFD enabled"
    else
        log "BFD disabled, modifying /etc/frr/daemons..."
        sed -i 's/^bfdd=yes/bfdd=no/' /etc/frr/daemons
    fi
}

# Generate frr.conf from Jinja2 template
generate_frr_conf() {
    local interfaces="$1"

    log "Generating /etc/frr/frr.conf from template..."
    log "FRR Profile: $FRR_PROFILE"

    # Build JSON for ports array
    local ports_json="["
    if [ -n "$interfaces" ]; then
        IFS=',' read -ra IFACE_ARRAY <<< "$interfaces"
        for i in "${!IFACE_ARRAY[@]}"; do
            [ $i -gt 0 ] && ports_json+=","
            ports_json+="\"${IFACE_ARRAY[$i]}\""
        done
    fi
    ports_json+="]"

    # Build JSON with all template variables
    cat > /frr-vars.json <<EOF
{
    "ports": $ports_json,
    "NODE_IP": "$NODE_IP",
    "ASN_LOCAL": "$ASN_LOCAL",
    "ASN_VRF_METALLB": "$ASN_VRF_METALLB",
    "ASN_METALLB_SPEAKER": "$ASN_METALLB_SPEAKER",
    "VRF_METALLB": "$VRF_METALLB",
    "PEER_IP_FRR": "$PEER_IP_FRR",
    "PEER_IP_METALLB": "$PEER_IP_METALLB",
    "ENABLE_BFD": "$ENABLE_BFD",
    "FRR_PROFILE": "$FRR_PROFILE"
}
EOF

    # Generate config using j2cli
    jinja2 /etc/frr/frr.conf.j2 /frr-vars.json > /etc/frr/frr.conf
    log "Generated /etc/frr/frr.conf"
}

# Main
main() {
    log "FRR Startup Script - Initializing..."

    # Get interfaces from FE_PORT_NAMES or resolve from FE_MACS
    local interfaces=""
    if [ -n "$FE_PORT_NAMES" ]; then
        interfaces="$FE_PORT_NAMES"
        log "Using fabric interfaces from FE_PORT_NAMES: $interfaces"
    else
        interfaces=$(resolve_interfaces "$FE_MACS")
        log "Resolved fabric interfaces from FE_MACS: $interfaces"
    fi

    # Use MTU from first fabric interface as default if not set
    if [ -z "$INTERFACE_MTU" ]; then
        local first_iface="${interfaces%%,*}"
        if [ -n "$first_iface" ]; then
            INTERFACE_MTU=$(cat "/sys/class/net/$first_iface/mtu" 2>/dev/null || echo "1500")
            log "Using MTU $INTERFACE_MTU from $first_iface"
        else
            INTERFACE_MTU="1500"
            log "No interfaces resolved, using default MTU 1500"
        fi
    else
        log "Using user-provided MTU $INTERFACE_MTU"
    fi

    # Create MetalLB VRF and veth pair
    create_metallb_vrf
    create_metallb_veth

    # Configure daemons based on ENABLE_BFD
    configure_daemons

    # Generate FRR configuration
    generate_frr_conf "$interfaces"

    ## create vtysh.conf
    [ -r /etc/frr/vtysh.conf ] || touch /etc/frr/vtysh.conf

    # Ensure FRR runtime directories exist
    log "mkdir -p /var/run/frr || true"
    mkdir -p /var/run/frr || true
    chown -R frr:frr /etc/frr || true
    rm -rf /var/run/frr/* || true
    chown -R frr:frr /var/run/frr || true
    log "Initialization complete, starting FRR..."

    # Start FRR using the standard docker-start script
    exec /usr/lib/frr/docker-start
}

main "$@"
