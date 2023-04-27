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
WHERE start_date > $2 AND end_date < $1
ORDER BY rating ASC