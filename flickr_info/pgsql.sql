-- Table: public.image_info

-- DROP TABLE public.image_info;

CREATE TABLE public.image_info
(
    image_uid uuid NOT NULL DEFAULT uuid_generate_v4(),
    image_id character varying COLLATE pg_catalog."default",
    image_source character varying(10) COLLATE pg_catalog."default",
    obtain_time timestamp without time zone DEFAULT now(),
    taken_time timestamp without time zone,
    width integer NOT NULL,
    height integer NOT NULL,
    tags text COLLATE pg_catalog."default",
    image_url character varying COLLATE pg_catalog."default" NOT NULL,
    location geometry NOT NULL,
    owner character varying(20) COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    lat double precision,
    lon double precision,
    CONSTRAINT image_info_pkey PRIMARY KEY (image_uid)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.image_info
    OWNER to junyi;
COMMENT ON TABLE public.image_info
    IS 'images with geotags from flickr';