create table if not exists domains (
    id serial primary key,
    name varchar(255) not null unique
);

create table if not exists companies (
    id serial primary key,
    name varchar(255) not null unique
);

create table if not exists vacancies (
    id serial primary key,
    name varchar(255) not null unique,
    company_id int references companies(id),
    link varchar(255) unique,
    description text,
    published_at timestamp,
    experience int
);

create table if not exists cvs (
    id serial primary key,
    vacancy_id int references vacancies(id),
    employee_id int references employees(id)
);

create table if not exists skills (
    id serial primary key,
    name varchar(255) not null unique,
    description text,
    parent_id int references skills(id)
);


create table if not exists vacancy_skills (
    id serial primary key,
    vacancy_id int references vacancies(id),
    skill_id int references skills(id)
);

create table if not exists cv_domains ( /* Project domains */
    id serial primary key,
    cv_id int references cvs(id),
    domain_id int references domains(id)
);

create table if not exists skill_conflicts (
    id serial primary key,
    skill_1_id int references skills(id),
    skill_2_id int references skills(id),
    comment text,
    priority int
);

create table if not exists responsibilities (
    id serial primary key,
    name varchar(255) not null unique,
    priority int,
    skill_id int references skills(id),
    experience int,
    comments text
);

create table if not exists responsibility_synonims(
    id serial primary key,
    responsibility_id int references responsibilities(id),
    name varchar(255)
);

create table if not exists responsibility_conflicts(
    id serial primary key,
    responsibility_1_id int references responsibilities(id),
    responsibility_2_id int references responsibilities(id),
    comment text,
    priority int
);

create table if not exists projects(
    id serial primary key,
    name varchar(255) not null unique,
    description text,
    company_id int references companies(id)
);

create table if not exists cv_project(
    id serial primary key,
    cv_id int references cvs(id),
    project_id int references projects(id),
    years int
);

create table if not exists project_responsibilities(
    id serial primary key,
    cv_project_id int references cv_project(id),
    responsibility_id int references responsibilities(id),
    years int
);


create table if not exists vacancy_domains(
    id serial primary key,
    vacancy_id int references vacancies(id),
    domain_id int references domains(id)
)
