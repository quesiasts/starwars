CREATE DATABASE desafioStarWars

CREATE TABLE planets (
	id SERIAL PRIMARY KEY,
	name varchar(100) not null,
	climate varchar(100) not null,
	terrain varchar(100) not null
)

SELECT * FROM planets