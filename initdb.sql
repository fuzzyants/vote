
USE vote;

CREATE OR REPLACE TABLE Votes (
	id SERIAL PRIMARY KEY,
	maxUsers MEDIUMINT UNSIGNED NOT NULL,
	slug VARCHAR(255),
	expires DATE,
	done BOOLEAN,
	INDEX slug_ix (slug),
	INDEX expires_done_ix (expires, done)
);

CREATE OR REPLACE TABLE Options (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	url VARCHAR(2083),
	INDEX title_ix (title)
);

CREATE OR REPLACE TABLE Votes_Options (
	voteId BIGINT UNSIGNED NOT NULL,
	optionId BIGINT UNSIGNED NOT NULL,
	PRIMARY KEY (voteId, optionId)	
);

CREATE OR REPLACE TABLE Ballots (
	id SERIAL PRIMARY KEY,
	userId BIGINT UNSIGNED NOT NULL,
	voteId BIGINT UNSIGNED NOT NULL,
	INDEX userId_ix (userId),
	INDEX voteId_ix (voteId)

);

CREATE OR REPLACE TABLE Ballots_Options (
	ballotId BIGINT UNSIGNED NOT NULL,
	optionId BIGINT UNSIGNED NOT NULL,
	rank BIGINT UNSIGNED,
	PRIMARY KEY (ballotId, optionId)
);

CREATE OR REPLACE TABLE Users (
	id SERIAL PRIMARY KEY,
	token CHAR(64),
	name VARCHAR(255),
	email VARCHAR(320),
	lastSeen DATE,
	INDEX email_ix (email),
	INDEX lastSeen_ix (lastSeen)
);

