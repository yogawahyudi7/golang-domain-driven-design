#!/bin/bash
set -e

# Database initialization script with security best practices

# Create application user with limited privileges
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    -- Create read-only user for monitoring
    CREATE USER readonly_user WITH PASSWORD 'readonly_password';
    GRANT CONNECT ON DATABASE golang_domain_driven_design TO readonly_user;
    GRANT USAGE ON SCHEMA public TO readonly_user;
    GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly_user;
    ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO readonly_user;

    -- Create application user
    CREATE USER app_user WITH PASSWORD 'app_password';
    GRANT CONNECT ON DATABASE golang_domain_driven_design TO app_user;
    GRANT USAGE ON SCHEMA public TO app_user;
    GRANT CREATE ON SCHEMA public TO app_user;
    GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO app_user;
    GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO app_user;
    ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO app_user;
    ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT ON SEQUENCES TO app_user;

    -- Enable required extensions
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    CREATE EXTENSION IF NOT EXISTS "pg_stat_statements";
    
    -- Set secure password encryption
    SET password_encryption = 'scram-sha-256';
EOSQL

echo "Database initialization completed successfully"
