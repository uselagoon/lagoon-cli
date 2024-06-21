SHELL := /bin/bash

KCADM = /opt/jboss/keycloak/bin/kcadm.sh
KCADM_CONFIG = /tmp/.keycloak/kcadm.config
KCADM_STRING = --server http://$$(hostname -i):8080/auth --user $$KEYCLOAK_ADMIN_USER --password $$KEYCLOAK_ADMIN_PASSWORD --realm master
KCADM_LOGIN = $(KCADM) config credentials --config $(KCADM_CONFIG) $(KCADM_STRING)

.PHONY: build-lagoon
build-lagoon:
	docker compose up -d --quiet-pull
	$(MAKE) wait-for-keycloak
	@echo "\nYour Lagoon stack is now running locally - use 'GRAPHQL_API=http://0.0.0.0:33000/graphql KEYCLOAK_API=http://0.0.0.0:38088/auth' as the variables to access it\n"

.PHONY: wait-for-keycloak
wait-for-keycloak:
	$(info Waiting for Keycloak to be ready....)
	grep -m 1 "Config of Keycloak done." <(docker compose --compatibility logs -f keycloak 2>&1)

.PHONY: keycloak-setup
keycloak-setup:
	docker compose exec keycloak bash -c "/upload/configure-keycloak.sh"

.PHONY: reload-data
reload-data:
	docker compose up -d local-api-data-watcher-pusher

.PHONY: down
down:
	docker compose down --remove-orphans --volumes

.PHONY: up
up:
	$(MAKE) build-lagoon 
	$(MAKE) keycloak-setup
