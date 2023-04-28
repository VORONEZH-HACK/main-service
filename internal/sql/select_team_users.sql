WITH teams AS (
  SELECT userid FROM fsp.users_x_teams WHERE team=$1
)
SELECT 
  id, 
  name 
FROM fsp.users 
WHERE id=ANY(SELECT userid FROM teams)