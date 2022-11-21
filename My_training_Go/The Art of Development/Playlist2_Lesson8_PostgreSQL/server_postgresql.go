package Playlist2_Lesson8_PostgreSQL

CREATE TABLE public.author
(
id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR (100) NOT NULL,
author_id UUID NOT NULL ,
CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
);

INSERT INTO author (name) VALUES ('Народ');
INSERT INTO author (name) VALUES ('Джофн Роулинг');
INSERT INTO author (name) VALUES ('Джек Лондон');