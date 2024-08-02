
-- used to check which projects have no domain
select * from projects p left join project_domains pd on pd.project_id  = p.id where pd.id is null ;

-- used to check which skills have no responsibilities
select s.name, s.parent_id from skills s left join responsibilities r on r.skill_id = s.id where r.id is null;

-- used to check which skills have more responsibilities
select s.name as skill_name, count(r.id) as responsibilities_amount from skills s left join responsibilities r on r.skill_id = s.id group by s.name ORDER BY responsibilities_amount DESC;

-- check amount of skills without responsibilities
select count(s.name) from skills s left join responsibilities r on r.skill_id = s.id where r is null;