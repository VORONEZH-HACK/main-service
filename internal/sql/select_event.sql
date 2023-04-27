SELECT id, name, owner, description, start, end, min_participants, max_participants
FROM fsp.events 
WHERE id=$1