# 変数定義
ANSIBLE_DIR = ansible
PLAYBOOK = $(ANSIBLE_DIR)/playbook.yml
INVENTORY = $(ANSIBLE_DIR)/inventory

# Ansibleプレイブックの実行
.PHONY: deploy
deploy:
	@echo "Running Ansible playbook..."
	ansible-playbook -i $(INVENTORY) $(PLAYBOOK) --ask-become-pass


# Ansibleプレイブックの実行
.PHONY: deploy-debug
deploy-debug:
	@echo "Running Ansible playbook..."
	ansible-playbook -i $(INVENTORY) $(PLAYBOOK) --ask-become-pass -vvv


# ローカル開発環境のセットアップ
.PHONY: up-build
up-build:
	@echo "Starting local development environment with build..."
	docker-compose up --build

# ローカル開発環境のセットアップ
.PHONY: up
up:
	@echo "Starting local development environment..."
	docker-compose up -f 
	
# ローカル開発環境のセットアップ
.PHONY: down
down:
	@echo "Stop local development environment..."
	docker-compose down


# ssh-keygen -R tk2-412-47176.vs.sakura.ne.jp
