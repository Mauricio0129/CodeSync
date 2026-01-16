-- Optional: Indexes for performance
CREATE INDEX idx_projects_owner ON projects(owner_id);
CREATE INDEX idx_directories_project ON directories(project_id);
CREATE INDEX idx_directories_parent ON directories(parent_directory_id);
CREATE INDEX idx_files_project ON files(project_id);
CREATE INDEX idx_files_directory ON files(directory_id);