NAME := saga-app-infra

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(patsubst %/,%,$(dir $(MKFILE_PATH)))

INFRA_DIR = $(MKFILE_DIR)
ROOT_DIR = $(abspath $(MKFILE_DIR)/../../../)

TF_VERSION ?= 1.6.6
TG_VERSION ?= 0.53.0

# here mostly for an example of using compatible versions that can be used instead of the standard ones
ifeq ($(TF_VERSION),1.4.6)
	TG_VERSION=0.45.1
endif

TF_PLUGIN_CACHE := $(ROOT_DIR)/.tf_plugin_cache
BREAD_CRUMB := .$(NAME)-breadcrumb-$(TF_VERSION)

IMAGE_NAME := $(NAME)_$(TF_VERSION)
DOCKER_ARGS := -v $(ROOT_DIR):/workspace \
	-v ~/.aws:/root/.aws \
	-v ~/.ssh:/root/.ssh \
	--env-file .env

$(TF_PLUGIN_CACHE):
	@mkdir $(TF_PLUGIN_CACHE)

$(BREAD_CRUMB):
	@touch $@
	docker build -t $(IMAGE_NAME) \
	--build-arg TF_VERSION=$(TF_VERSION) \
	--build-arg TG_VERSION=$(TG_VERSION) .

shell: $(BREAD_CRUMB) $(TF_PLUGIN_CACHE)
	docker run -it \
	$(DOCKER_ARGS) \
	$(IMAGE_NAME) \
	/bin/bash --login

fmt:
	docker run \
	$(DOCKER_ARGS) \
	$(IMAGE_NAME) \
	/bin/sh -c 'terragrunt hclfmt && terraform fmt -recursive -diff'
