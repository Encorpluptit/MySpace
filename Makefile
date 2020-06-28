
API_DIR					=	api

api-bkgr:
	cd $(API_DIR) && docker-compose up -d --build

clean-api-bkgr:
	cd $(API_DIR) && docker-compose down --remove-orphans --volumes

api-test:
	cd $(API_DIR) && docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit