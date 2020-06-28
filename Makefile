
API_DIR					=	api

api-start:
	cd $(API_DIR) && docker-compose up -d --build

api-test:
	cd $(API_DIR) && docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

api-stop:
	cd $(API_DIR) && docker-compose down --remove-orphans --volumes

clean: api-stop
	echo "y" | docker system prune

fclean:
	docker system prune -a --volumes

.PHONY: api-main api-test clean-api-bkgr
