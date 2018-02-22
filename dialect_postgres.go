package schema

var postgres = dialect{
	queries: [3]string{
		// tableNames query.
		pack(`
			SELECT table_name
			FROM
				information_schema.tables
			WHERE
				table_type = 'BASE TABLE' AND
				table_schema = current_schema()
		`),
		// viewNames query.
		pack(`
			SELECT table_name
			FROM
				information_schema.tables
			WHERE
				table_type = 'VIEW' AND
				table_schema = current_schema()
		`),
		// columnTypes query.
		`SELECT * FROM %s LIMIT 0`,
	},
}
