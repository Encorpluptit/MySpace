
API_DIR					=	api

api-bkgr:
	cd $(API_DIR) && docker-compose up -d --build

api-test:
	cd $(API_DIR) && docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

clean-api-bkgr:
	cd $(API_DIR) && docker-compose down --remove-orphans --volumes

clean: clean-api-bkgr

fclean:
	docker system prune -a --volumes

.PHONY: api-bkgr api-test clean-api-bkgr
