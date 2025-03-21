BINARY_NAME=panoptes
CONFIG_DIR=$(HOME)/.config/panoptes

.PHONY: clan test install

###########
# Install #
###########

install:
	@echo "Building and installing..."
	@go build -ldflags "-X dlvlabs.net/panoptes-agent/cmd/cli.configPath=$(CONFIG_DIR)/config.toml" -o $(GOPATH)/bin/$(BINARY_NAME) main.go
	@echo "Installing config files..."
	@mkdir -p $(CONFIG_DIR)
	@if [ ! -f $(CONFIG_DIR)/config.toml ]; then \
		cp config/config.toml $(CONFIG_DIR)/config.toml; \
	fi

###########
# Clean   #
###########

clean:
	@echo "Cleaning..."
	@rm -f $(GOPATH)/bin/$(BINARY_NAME)
	@echo "Config files will remain in $(CONFIG_DIR)"
	@echo "To remove config files completely, run: make clean-all"

clean-all: clean
	@echo "Removing config files..."
	@rm -rf $(CONFIG_DIR)

###########
# Test    #
###########

test:
	@echo "Running tests..."
	@go test ./...

.DEFAULT_GOAL := install
