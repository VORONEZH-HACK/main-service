WITH teams AS (
  SELECT team FROM fsp.events_x_teams WHERE event=$1
)
SELECT 
  id, 
  name 
FROM fsp.teams 
WHERE id=ANY(SELECT team FROM teams)