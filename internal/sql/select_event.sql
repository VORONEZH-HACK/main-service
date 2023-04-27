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
WHERE id=$1