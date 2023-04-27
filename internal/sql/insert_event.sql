INSERT INTO fsp.events (
  name, 
  description, 
  owner, 
  min_participants, 
  max_participants, 
  start_date,
  end_date)
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING id