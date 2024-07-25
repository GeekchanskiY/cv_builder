create table if not exists employees (
    id serial primary key,
    name varchar(255) not null unique,
    about_me text,
    image_url varchar(255),
    real_experience int
);

create table if not exists domains (
    id serial primary key,
    name varchar(255) not null unique,
    description text
);

create table if not exists companies (
    id serial primary key,
    name varchar(255) not null unique,
    description text,
    homepage varchar(255),
    is_trusted boolean
);

create table if not exists projects(
    id serial primary key,
    name varchar(255) not null unique,
    description text
);

/* We can create a same project with different prefix,
   which will be the same, but with different architecture for example */
create table if not exists project_services(
    id serial primary key,
    project_id int references projects(id),
    name varchar(255),
    description text
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
    name varchar(255) not null unique,
    vacancy_id int references vacancies(id),
    employee_id int references employees(id),
    is_real boolean
);

create table if not exists skills (
    id serial primary key,
    name varchar(255) not null unique,
    description text,
    parent_id int references skills(id)
);

create table if not exists skill_domains (
    id serial primary key,
    domain_id int references domains(id),
    skill_id int references skills(id),
    comments text,
    priority int
);


create table if not exists vacancy_skills (
    id serial primary key,
    vacancy_id int references vacancies(id),
    skill_id int references skills(id),
    priority int
);

create table if not exists project_domains (
    id serial primary key,
    project_id int references projects(id),
    domain_id int references domains(id),
    comments text
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

create table if not exists responsibility_synonyms(
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

create table if not exists cv_projects(
    id serial primary key,
    cv_id int references cvs(id),
    project_id int references projects(id),
    company_id int references companies(id),
    end_time date,
    start_time date
);

create table if not exists cv_project_services(
    id serial primary key,
    cv_project_id int references cv_projects(id),
    project_service_id int references project_services(id),
    order_num int
);

create table if not exists cv_service_responsibilities(
    id serial primary key,
    cv_service_id int references cv_project_services(id),
    responsibility_id int references responsibilities(id),
    order_num int
);

create table if not exists vacancy_domains(
    id serial primary key,
    vacancy_id int references vacancies(id),
    domain_id int references domains(id),
    priority int
)
