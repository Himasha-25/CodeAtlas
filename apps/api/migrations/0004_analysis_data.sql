CREATE TABLE files (
    id         BIGSERIAL PRIMARY KEY,
    run_id     BIGINT NOT NULL REFERENCES analysis_runs(id) ON DELETE CASCADE,
    path       TEXT NOT NULL,
    language   TEXT,
    line_count INT
);

CREATE TABLE symbols (
    id         BIGSERIAL PRIMARY KEY,
    file_id    BIGINT NOT NULL REFERENCES files(id) ON DELETE CASCADE,
    run_id     BIGINT NOT NULL REFERENCES analysis_runs(id) ON DELETE CASCADE,
    name       TEXT NOT NULL,
    kind       TEXT NOT NULL,
    line       INT,
    exported   BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE dependencies (
    id        BIGSERIAL PRIMARY KEY,
    run_id    BIGINT NOT NULL REFERENCES analysis_runs(id) ON DELETE CASCADE,
    from_id   BIGINT NOT NULL,
    to_id     BIGINT NOT NULL,
    edge_type TEXT NOT NULL
);

CREATE TABLE api_endpoints (
    id      BIGSERIAL PRIMARY KEY,
    run_id  BIGINT NOT NULL REFERENCES analysis_runs(id) ON DELETE CASCADE,
    method  TEXT NOT NULL,
    path    TEXT NOT NULL,
    file    TEXT,
    line    INT
);

CREATE INDEX idx_files_run_id ON files(run_id);
CREATE INDEX idx_symbols_run_id ON symbols(run_id);
CREATE INDEX idx_symbols_file_id ON symbols(file_id);
CREATE INDEX idx_dependencies_run_id ON dependencies(run_id);
