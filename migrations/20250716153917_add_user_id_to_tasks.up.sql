alter table tasks
add column user_id UUID references users(id) on delete cascade;