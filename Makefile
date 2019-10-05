all: env-down env-clean env-up

.PHONY: env-up
env-up:
	@echo "Starting environment...."
	@docker-compose -f fixtures/docker-compose-cli.yaml up -d
	@echo "Environment up"
	@docker exec -it cli bash


.PHONY:env-down
env-down:
	 @echo "Stoping environment...."
	 @docker-compose -f fixtures/docker-compose-cli.yaml down
	 @echo "Environment down"

.PHONY:env-clean
env-clean:
	@echo "Clean up ..."
	@docker network prune -f # 来清理没有再被任何容器引用的networks
	@docker volume prune -f  # 清理挂载卷
	@docker rm -f -v `docker ps -a --no-trunc | grep "mycc" | cut -d ' ' -f 1` 2>/dev/null || true  #mycc 是链码的名字
	@docker rmi `docker images --no-trunc | grep "mycc" | cut -d ' ' -f 1` 2>/dev/null || true
	@echo "Clean up done"


