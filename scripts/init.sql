-- Crear usuario y base de datos para Waiter API
-- Ejecutar como superusuario (postgres)

CREATE USER waiter_db WITH PASSWORD '-Q73cyvfsLp#E:S*';
CREATE DATABASE waiter OWNER waiter_db;

\c waiter

-- Otorgar permisos
GRANT ALL PRIVILEGES ON DATABASE waiter TO waiter_db;
GRANT ALL PRIVILEGES ON SCHEMA public TO waiter_db;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO waiter_db;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO waiter_db;
