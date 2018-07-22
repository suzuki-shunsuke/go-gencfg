if [ -n "$CODECOV_TOKEN" ]; then
  bash <(curl -s https://codecov.io/bash) || exit 1
fi
