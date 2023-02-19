-- +goose Up
-- +goose StatementBegin
CREATE TABLE samples.masters (
    id BIGSERIAL NOT NULL,
    phone VARCHAR(15) NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP(0) with time zone,
    CONSTRAINT masters_pkey PRIMARY KEY(id)
)
    WITH (oids = false);

COMMENT ON TABLE samples.masters IS 'Список мастеров';
COMMENT ON COLUMN samples.masters.id IS 'ID мастера';
COMMENT ON COLUMN samples.masters.phone IS 'Номер телефона мастера';
COMMENT ON COLUMN samples.masters.name IS 'Имя мастера';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS samples.masters;
-- +goose StatementEnd
