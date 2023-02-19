-- +goose Up
-- +goose StatementBegin
CREATE TABLE samples.appointments (
    id BIGSERIAL NOT NULL,
    master_id BIGINT NOT NULL,
    client_id BIGINT,
    service_id BIGINT,
    status_id INT NOT NULL DEFAULT 1,
    appointment_date DATE NOT NULL,
    appointment_time TIME NOT NULL,
    note TEXT,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP(0) with time zone,
    CONSTRAINT appointments_pkey PRIMARY KEY(id)
);

COMMENT ON TABLE samples.appointments IS 'Список мастеров';
COMMENT ON COLUMN samples.appointments.id IS 'ID встречи';
COMMENT ON COLUMN samples.appointments.master_id IS 'ID мастера';
COMMENT ON COLUMN samples.appointments.client_id IS 'ID клиента';
COMMENT ON COLUMN samples.appointments.service_id IS 'ID услуги';
COMMENT ON COLUMN samples.appointments.status_id IS 'ID статуса';
COMMENT ON COLUMN samples.appointments.appointment_date IS 'Назначенная дата';
COMMENT ON COLUMN samples.appointments.appointment_time IS 'Назначенное время';
COMMENT ON COLUMN samples.appointments.note IS 'Примечание';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS samples.appointments;
-- +goose StatementEnd
