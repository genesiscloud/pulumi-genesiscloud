# Configure

1. Create a directory at the root of your repo called .github/workflows

2. Place the release.yml from this directory there

3. Place the .goreleaser.yml from this directory at the root of your repo

4. The example shown uses Pulumi's [Publishing Action](https://github.com/pulumi/pulumi-package-publisher) to publish the language SDKS.
   Refer to the README for any environment secrets you need to set up.

5. Customize .goreleaser.yml for your provider, paying special attention that the ldlflags are set to match your provider/go.mod exactly:

   `-X github.com/pulumi/pulumi-aws/provider/v5/pkg/version.Version={{.Tag}}`

6. Delete this directory if desired

# Deploy

7. Push a tag to your repo in the format "v0.0.0" to initiate a release

8. IMPORTANT: also add a tag in the format "sdk/v0.0.0" for the Go SDK

# Upgrade

Pulumi provides a [GitHub action](https://github.com/pulumi/pulumi-upgrade-provider-action) to automate upgrading your provider upon a new upstream version release. An example workflow that runs the upgrade action on a cron job, as well as whenever an issue is created with a title prefix of 'Upgrade terraform-provider', is [provided here](./upgrade-provider.yml).
