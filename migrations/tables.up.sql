CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP SCHEMA IF EXISTS fsp CASCADE;
CREATE SCHEMA IF NOT EXISTS fsp;

CREATE TABLE fsp.users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  email TEXT NOT NULL,
  password TEXT NOT NULL,

  rating INTEGER DEFAULT 0,
  edu TEXT NOT NULL DEFAULT '',
  name TEXT NOT NULL,
  description TEXT NOT NULL DEFAULT '',
  patronymic TEXT NOT NULL,
  surname TEXT NOT NULL
);

CREATE TABLE fsp.users_x_tags (
  userid UUID NOT NULL,
  tag TEXT NOT NULL,

  FOREIGN KEY (userid) REFERENCES fsp.users(id)
);

CREATE TABLE fsp.organizations (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  description TEXT NOT NULL DEFAULT '',
  phone TEXT NOT NULL DEFAULT '',

  name TEXT NOT NULL
);

CREATE TABLE fsp.tokens (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  -- salt TEXT NOT NULL DEFAULT '',
  "type" TEXT NOT NULL DEFAULT 'USER',
  created TIMESTAMP DEFAULT now()
);

CREATE TABLE fsp.users_x_roles (
  user_id UUID NOT NULL,
  role_id UUID NOT NULL
);

CREATE TABLE fsp.roles (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL
);

CREATE TABLE fsp.teams (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL,
  lead UUID NOT NULL,

  FOREIGN KEY (lead) REFERENCES fsp.users(id)
);

CREATE TABLE fsp.users_x_teams (
  userid UUID NOT NULL,
  team UUID NOT NULL,

  FOREIGN KEY (team) REFERENCES fsp.teams(id),
  FOREIGN KEY (userid) REFERENCES fsp.users(id)
);

CREATE TABLE fsp.invites (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  userid UUID NOT NULL,
  team UUID NOT NULL,

  FOREIGN KEY (team) REFERENCES fsp.teams(id),
  FOREIGN KEY (userid) REFERENCES fsp.users(id)
);

CREATE TABLE fsp.events (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL,
  description TEXT NOT NULL DEFAULT '',
  owner UUID NOT NULL,
  place TEXT NOT NULL DEFAULT '',
  rating REAL NOT NULL DEFAULT 0,
  min_participants INTEGER NOT NULL DEFAULT 1,
  max_participants INTEGER NOT NULL DEFAULT 5,
  start_date BIGINT NOT NULL,
  end_date BIGINT NOT NULL

  -- FOREIGN KEY (owner) REFERENCES fsp.organizations(id)
);

CREATE TABLE fsp.events_x_teams (
  team UUID NOT NULL,
  event UUID NOT NULL,

  FOREIGN KEY (team) REFERENCES fsp.teams(id),
  FOREIGN KEY (event) REFERENCES fsp.events(id)
);