#!/usr/bin/env bash

for x in fullbogons dshield spamhaus_drop spamhaus_edrop
do
    ipset create ${x} hash:net || true
    update-ipsets enable ${x} || true
done

update-ipsets

for x in fullbogons dshield spamhaus_drop spamhaus_edrop
do
    # ipset-apply /etc/firehol/ipsets/${x}
    ipset-apply ${x} || true
done
