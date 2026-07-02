#!/bin/bash
# Wrapper script to invoke frr-startup.sh
# This allows overriding frr-startup.sh via extension service config for testing
exec /bin/bash /usr/local/bin/frr-startup.sh
