SELECT id, name, owner, description, start_date, end_date, min_participants, max_participants
FROM fsp.events
ORDER BY start_date DESC LIMIT 100