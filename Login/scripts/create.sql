CREATE TABLE IF NOT EXISTS public.usuario
(
    id serial NOT NULL,
    nome character varying COLLATE pg_catalog."default",
    email character varying COLLATE pg_catalog."default",
    squad integer,
    permissao character varying COLLATE pg_catalog."default",
    CONSTRAINT usuario_pkey PRIMARY KEY (id)
)