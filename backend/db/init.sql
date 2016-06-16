CREATE TABLE IF NOT EXISTS users (
	username VARCHAR(32) PRIMARY KEY,
	passhash TEXT,
	profile  jsonb
);

CREATE TABLE IF NOT EXISTS friend (
	owner       VARCHAR(32),
	group_name  VARCHAR(32),
	friend      VARCHAR(32),
	PRIMARY KEY (owner, friend)
);

CREATE TABLE IF NOT EXISTS chat (
	sender      VARCHAR(32),
	receiver    VARCHAR(32),
	send_time   TIMESTAMP NOT NULL,
	recv_time   TIMESTAMP,
	message     TEXT NOT NULL,
	received    BOOLEAN NOT NULL,
	PRIMARY KEY (sender, receiver, send_time)
);
