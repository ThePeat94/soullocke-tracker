.PHONY: generate-client boot

generate-client:
	$(MAKE) -C soullocke-backend export-openapi
	$(MAKE) -C soullocke-frontend generate

boot:
	$(MAKE) -C soullocke-backend run & \
	$(MAKE) -C soullocke-frontend run & \
	wait