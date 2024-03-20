# Deployment

1. Update the version in the `Makefile` file, e.g. `vX.Y.Z`.
2. Run `make build` to build the binaries and SDKs locally.
3. Create a new release on GitHub with a tag `vX.Y.Z` and `sdk/vX.Y.Z`. Make sure to only tag `vX.Y.Z` as latest.
4. The release action will be automatically triggered by the tag and will publish the binaries and SDKs to the GitHub release.
