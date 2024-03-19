# Genesis Cloud Resource Provider

The Genesis Cloud Resource Provider lets you manage [Genesis Cloud](http://genesiscloud.com) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @genesiscloud/pulumi-genesiscloud
```

or `yarn`:

```bash
yarn add @genesiscloud/pulumi-genesiscloud
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-genesiscloud
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/genesiscloud/pulumi-genesiscloud/sdk/go/...
```

## Configuration

The following configuration points are available:

- `genesiscloud:token` - (Required) This is the Genesis Cloud API Token, can also be specified with the `GENESISCLOUD_TOKEN` environment variable.
- `genesiscloud:endpoint` - (Optional) Genesis Cloud API endpoint, can be used to override the default API Endpoint `https://api.genesiscloud.com`.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/genesiscloud/api-docs/).
