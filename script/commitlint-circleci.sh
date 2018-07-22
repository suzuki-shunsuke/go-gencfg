if [ -n "$CIRCLE_PULL_REQUEST" ]; then
  npx commitlint --from master --to $CIRCLE_BRANCH
else
  npx commitlint --from HEAD~10 --to HEAD
fi
