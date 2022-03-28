DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS public.users(
	ID serial8 NOT NULL UNIQUE,
	Username VARCHAR(127) NOT NULL UNIQUE,
	Password VARCHAR(127) NOT NULL,
	PRIMARY KEY(ID)
);