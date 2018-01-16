CREATE KEYSPACE IF NOT EXISTS rtsupport WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE rtsupport;

CREATE TABLE IF NOT EXISTS user (
  id uuid,
  name text,PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS username_key ON user (name);

CREATE TABLE IF NOT EXISTS channel (
  id uuid,
  name text,PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS channelname_key ON channel (name);

CREATE TABLE IF NOT EXISTS message (
  id uuid,
  author text,
  createat text,
  body text,
  channelid uuid,PRIMARY KEY (channelid,createat)
);