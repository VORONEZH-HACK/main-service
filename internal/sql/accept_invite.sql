WITH accept AS (
  INSERT INTO fsp.users_x_teams (userid, team) 
  VALUES (
    (SELECT userid FROM fsp.invites WHERE id=$1),
    (SELECT team FROM fsp.invites WHERE id=$1),
  )
)
DELETE FROM fsp.invites WHERE id=$1