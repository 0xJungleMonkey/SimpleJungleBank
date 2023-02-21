# Note

[sqlc](https://docs.sqlc.dev/en/stable/reference/config.html#version-2)

One important thing when working with sqlc is
We should not modify the content of the generated files,
Because everytime we run make sqlc,
all of those files will be regenerated,
and our changes will be lost.
So make sure to create new files if you want to add more codes to the db package.
