-- name: CreateUser :one
INSERT INTO users(
    userId,
    name,
    email,
  password,
  role,
  profilephoto,
  createdAt
) VALUES ($1,$2,$3,$4,$5,$6,now())RETURNING *;