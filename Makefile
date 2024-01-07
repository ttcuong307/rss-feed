migrate-up:
	goose mysql "root:password@/rss_feed?parseTime=true" up 

migrate-down:
	goose mysql "root:password@/rss_feed?parseTime=true" down 

generate-sql:
	sqlc generate