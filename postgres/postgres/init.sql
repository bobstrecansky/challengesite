SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;
CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
SET search_path = public, pg_catalog;
SET default_tablespace = '';
SET default_with_oids = false;
CREATE TABLE users (
   ID  SERIAL PRIMARY KEY,
   NAME           TEXT      NOT NULL,
   EMAIL          CHAR(50)  NOT NULL,
);

ALTER TABLE users OWNER TO postgres;
COPY users (NAME, EMAIL) FROM stdin;
1	bob bob.strecansky@gmail.com
\.
