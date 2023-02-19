-- +goose Up
-- +goose StatementBegin
CREATE TABLE samples.services (
    id BIGSERIAL NOT NULL,
    master_id BIGINT NOT NULL,
    name VARCHAR(200) NOT NULL,
    price INT,
    time TIME,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP(0) with time zone,
    CONSTRAINT services_pkey PRIMARY KEY(id)
)
    WITH (oids = false);

COMMENT ON TABLE samples.services IS 'Список услуг';
COMMENT ON COLUMN samples.services.id IS 'ID услуги';
COMMENT ON COLUMN samples.services.master_id IS 'ID мастера';
COMMENT ON COLUMN samples.services.name IS 'Наименование услуги';
COMMENT ON COLUMN samples.services.price IS 'Цена услуги';
COMMENT ON COLUMN samples.services.time IS 'Время выполнения услуги';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS samples.services;
-- +goose StatementEnd
