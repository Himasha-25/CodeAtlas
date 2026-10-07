ALTER TABLE projects
    ADD COLUMN active_repository_id BIGINT REFERENCES repositories(id) ON DELETE SET NULL;
