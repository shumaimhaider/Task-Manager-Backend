CREATE TABLE public.task (
  task_id SERIAL NOT NULL PRIMARY KEY,
  details TEXT not null,
  is_completed boolean DEFAULT False,
  task_type_id int not null,
  CONSTRAINT fk_task_type_id
  Foreign KEY(task_type_id)References public.task_type(task_type_id) on Delete Cascade,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);