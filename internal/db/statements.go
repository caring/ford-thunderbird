package db

var statements = map[string]string{
  // inserts a new row into the thunderbirds table
  "create-thunderbird": `
  INSERT INTO thunderbirds (thunderbird_id, name)
    values(UUID_TO_BIN(?), ?)
  `,
  // soft deletes a thunderbird by id
  "delete-thunderbird": `
  UPDATE
    thunderbirds
  SET
    deleted_at = NOW()
  WHERE
    thunderbird_id = UUID_TO_BIN(?)
    AND deleted_at IS NULL
  `,
  // gets a single thunderbird row by id
  "get-thunderbird": `
  SELECT
    thunderbird_id, name
  FROM
    thunderbirds
  WHERE
    thunderbird_id = UUID_TO_BIN(?)
    AND deleted_at IS NULL
  `,
  // update a single thunderbird row by ID
  "update-thunderbird": `
  UPDATE
    thunderbirds
  SET
    name = ?
  WHERE
    thunderbird_id = UUID_TO_BIN(?)
    AND deleted_at IS NULL
  `,
}
