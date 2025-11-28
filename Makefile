.PHONY: init
init: check-signing-key
	@git config --local commit.gpgsign true
	@git config --local tag.gpgsign true
	@echo "Local commit and tag signing enabled (see .git/config)"

check-signing-key:
	@if ! git config --get user.signingkey >/dev/null 2>&1; then \
		echo "ERROR: user.signingkey is not configured"; \
		echo "Configure it with:"; \
		echo "  git config --local user.signingkey <GPG_KEY_ID>"; \
		exit 1; \
	else \
		echo "Using signing key: $$(git config --get user.signingkey)"; \
	fi
