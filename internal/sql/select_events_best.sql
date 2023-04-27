SELECT id, name, owner, description, start_date, end_date, min_participants, max_participants
FROM fsp.events
WHERE start_date > $2 AND end_date < $1
ORDER BY rating ASC