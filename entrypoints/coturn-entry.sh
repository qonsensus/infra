#!/bin/sh
set -e

# Substitute env vars into config and write to a temp location
sed \
  -e "s|\${COTURN_DOMAIN}|${COTURN_DOMAIN}|g" \
  -e "s|\${TURN_SECRET}|${TURN_SECRET}|g" \
  -e "s|\${TURN_MIN_PORT}|${TURN_MIN_PORT}|g" \
  -e "s|\${TURN_MAX_PORT}|${TURN_MAX_PORT}|g" \
  /etc/coturn/turnserver.conf > /tmp/turnserver.conf

exec turnserver -c /tmp/turnserver.conf
