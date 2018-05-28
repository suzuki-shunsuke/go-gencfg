if [ -n "$TRAVIS_TAG" ]; then
  gox -output="dist/${TRAVIS_TAG}/gencfg_${TRAVIS_TAG}_{{.OS}}_{{.Arch}}" -osarch="darwin/amd64 linux/amd64 windows/amd64" .
fi
