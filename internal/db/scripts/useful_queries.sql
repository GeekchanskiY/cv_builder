
-- used to check which projects have no domain
select * from projects p left join project_domains pd on pd.project_id  = p.id where pd.id is null ;