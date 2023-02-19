-- +goose Up
-- +goose StatementBegin
CREATE TABLE samples.clients (
    id BIGSERIAL NOT NULL,
    phone VARCHAR(15) NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP(0) with time zone,
    CONSTRAINT clients_pkey PRIMARY KEY(id)
)
    WITH (oids = false);

COMMENT ON TABLE samples.clients IS 'Список пользователей';
COMMENT ON COLUMN samples.clients.id IS 'ID клиента';
COMMENT ON COLUMN samples.clients.phone IS 'Номер телефона клиента';
COMMENT ON COLUMN samples.clients.name IS 'Имя клиента';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS samples.clients;
-- +goose StatementEnd
