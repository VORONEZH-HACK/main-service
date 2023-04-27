SELECT 
  id, 
  email, 
  password, 
  edu, 
  name, 
  patronymic, 
  surname, 
  rating
FROM fsp.users WHERE id=$1