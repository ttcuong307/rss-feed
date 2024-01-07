migrate-up:
	cd "./sql/schema" && goose mysql "root:password@/rss_feed?parseTime=true" up 

migrate-down:
	cd "./sql/schema" && goose mysql "root:password@/rss_feed?parseTime=true" down 

generate-sql:
	sqlc generate