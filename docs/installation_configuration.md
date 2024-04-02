---
title: Genesis Cloud Installation & Configuration
meta_desc: Information on how to install the Genesis Cloud provider.
layout: package
---

## Installation

The Pulumi Genesis Cloud provider is available as a package in the following languages:

- JavaScript/TypeScript: [`@genesiscloud/pulumi-genesiscloud`](https://www.npmjs.com/package/@genesiscloud/pulumi-genesiscloud)
- Python: [`pulumi-genesiscloud`](https://pypi.org/project/pulumi-genesiscloud/)
- Go: [`github.com/genesiscloud/pulumi-genesiscloud/sdk/go/genesiscloud`](https://pkg.go.dev/github.com/pulumi/pulumi-genesiscloud/sdk)

## Setup

To provision resources with the Genesis Cloud provider, you need to have a Genesis Cloud token.

## Configuration Options

Use `pulumi config set genesiscloud:<option>`.

| Option     | Required/Optional | Description                                                             |
| ---------- | ----------------- | ----------------------------------------------------------------------- |
| `token`    | Required          | Genesis Cloud access token                                              |
| `endpoint` | Optional          | The endpoint to use in the provider. Defaults to `api.genesiscloud.com` |
