#!/bin/bash
set -e

# Connect to the default postgres database and create the roles and databases
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    -- Create test_admin user with full privileges
    CREATE USER test_admin WITH PASSWORD 'admin_password';
    CREATE DATABASE test_db;
    GRANT ALL PRIVILEGES ON DATABASE test_db TO test_admin;

    -- Create test_user with limited privileges
    CREATE USER test_user WITH PASSWORD 'test_password';
    GRANT CONNECT ON DATABASE test_db TO test_user;
    \c test_db
    GRANT USAGE ON SCHEMA public TO test_user;
    GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA public TO test_user;
    ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE ON TABLES TO test_user;
EOSQL
