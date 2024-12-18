CREATE TABLE IF NOT EXISTS subsidiaries
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
);


CREATE TABLE IF NOT EXISTS clients
(
    id SERIAL PRIMARY KEY,
    num1 VARCHAR(12) NOT NULL,
    num2 VARCHAR(12) NOT NULL,
    num3 VARCHAR(12) NOT NULL,
    subsidiary FOREIGN KEY (subsidiary_id) REFERENCES subsidiaries(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO subsidiaries (name) VALUES ('Subsidiary 1');
INSERT INTO subsidiaries (name) VALUES ('Subsidiary 2');
INSERT INTO subsidiaries (name) VALUES ('Subsidiary 3');

CREATE INDEX IF NOT EXISTS idx_clients_subsidiary_id ON clients