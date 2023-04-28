WITH accept AS (
  INSERT INTO fsp.users_x_teams (userid, team) 
  VALUES (
    (SELECT userid FROM fsp.invites WHERE id=$1 AND userid=$2),
    (SELECT team FROM fsp.invites WHERE id=$1 AND userid=$2),
  )
)
DELETE FROM fsp.invites WHERE id=$1