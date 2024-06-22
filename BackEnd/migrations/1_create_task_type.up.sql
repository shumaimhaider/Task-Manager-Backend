CREATE TABLE public.task_type(
    task_type_id SERIAL NOT NULL PRIMARY KEY,
    task_type_name VARCHAR (50) UNIQUE NOT NULL
)