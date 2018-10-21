#!/usr/bin/env sh

psql -c "create database ipsets"

psql -c "create schema ipsets"

psql ipsets -c "CREATE TABLE provider (
    last_updated_time TIMESTAMP WITH TIME ZONE default current_timestamp,
    pname text NOT NULL,
    metadata jsonb,
    PRIMARY KEY (pname)
)"

# https://www.postgresql.org/docs/9.1/static/datatype-net-types.html

psql ipsets -c "CREATE TABLE ipsets (
    last_updated_time TIMESTAMP WITH TIME ZONE default current_timestamp,
    address CIDR,
    provider text REFERENCES provider(pname),
    PRIMARY KEY (address, provider)
)"

psql ipsets -c "CREATE INDEX idx_address ON ipsets USING GIST(address inet_ops)"