WITH teams_with_participants_count AS (
  SELECT 
    fsp.users_x_teams.team AS team, 
    count(*) AS participants 
  FROM fsp.users_x_teams 
  GROUP BY fsp.users_x_teams.team
), events_with_participants_count AS (
  SELECT 
    fsp.events_x_teams.event AS event, 
    SUM(teams_with_participants_count.participants) AS participants
  FROM fsp.events_x_teams
  INNER JOIN teams_with_participants_count 
  ON fsp.events_x_teams.team = teams_with_participants_count.team
  GROUP BY fsp.events_x_teams.event
)
SELECT 
  fsp.events.id, 
  fsp.events.name, 
  fsp.events.owner, 
  fsp.events.place,
  fsp.events.description, 
  fsp.events.start_date, 
  fsp.events.end_date, 
  fsp.events.min_participants, 
  fsp.events.max_participants,
  fsp.events.rating,
  events_with_participants_count.participants
FROM fsp.events
INNER JOIN events_with_participants_count
ON events_with_participants_count.event = fsp.events.id
WHERE start_date > $1 AND end_date < $2
ORDER BY events_with_participants_count.participants
