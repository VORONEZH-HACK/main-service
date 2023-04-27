INSERT INTO fsp.teams (name, lead) 
VALUES ($1, $2) RETURNING id