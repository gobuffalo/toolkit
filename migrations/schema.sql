--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.5
-- Dumped by pg_dump version 9.6.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: licenses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE licenses (
    id uuid NOT NULL,
    tool_id uuid NOT NULL,
    name character varying(255) NOT NULL,
    body text,
    description text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE licenses OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE schema_migration OWNER TO postgres;

--
-- Name: tools; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE tools (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    name_with_owner character varying(255) NOT NULL,
    url character varying(255) NOT NULL,
    discovery_engine character varying(255) DEFAULT 'github'::character varying NOT NULL,
    description text,
    readme text,
    authors character varying[] NOT NULL,
    topics character varying[] NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    stars integer DEFAULT 0 NOT NULL
);


ALTER TABLE tools OWNER TO postgres;

--
-- Name: licenses licenses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY licenses
    ADD CONSTRAINT licenses_pkey PRIMARY KEY (id);


--
-- Name: tools tools_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY tools
    ADD CONSTRAINT tools_pkey PRIMARY KEY (id);


--
-- Name: licenses_tool_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX licenses_tool_id_idx ON licenses USING btree (tool_id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON schema_migration USING btree (version);


--
-- Name: tools_stars_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX tools_stars_idx ON tools USING btree (stars);


--
-- Name: tools_url_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX tools_url_idx ON tools USING btree (url);


--
-- PostgreSQL database dump complete
--

