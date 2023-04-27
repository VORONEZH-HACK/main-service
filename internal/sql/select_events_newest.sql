SELECT 
  id, 
  name, 
  owner, 
  place,
  description, 
  start_date, 
  end_date, 
  min_participants, 
  max_participants,
  rating
FROM fsp.events
WHERE start_date > extract(epoch from now())
ORDER BY start_date DESC LIMIT 100