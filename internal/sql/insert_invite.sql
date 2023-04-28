INSERT INTO fsp.invites (userid, team)
VALUES ($1, $2) RETURNING id