CREATE VIEW alerts_sum_by_month as select
	DATE_TRUNC('month', alerts.created_at) as date,
	count(alerts.id) as total
from alerts
group by DATE_TRUNC('month', alerts.created_at)
order by date desc;


CREATE VIEW alerts_priority_by_month as select
	DATE_TRUNC('month', alerts.created_at) as date,
	alerts.priority,
	count(alerts.id) as total
from alerts
group by DATE_TRUNC('month', alerts.created_at), alerts.priority
order by date desc, alerts.priority asc;

CREATE VIEW alerts_integration_by_month as select
	info.date,
	integrations."name",
	integrations."type",
	info.total,
	integrations.enabled
from (
	select
		DATE_TRUNC('month', alerts.created_at) as date,
		alerts.integration_id,
		count(alerts.id) as total
	from alerts
	group by DATE_TRUNC('month', alerts.created_at), alerts.integration_id
) as info
left join integrations on info.integration_id = integrations.id
order by info.date desc;
