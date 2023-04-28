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
WHERE start_date > $1 AND start_date < $2