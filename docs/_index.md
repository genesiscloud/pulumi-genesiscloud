---
title: Genesis Cloud
meta_desc: Provides an overview of the Genesis Cloud Provider for Pulumi.
layout: package
---

The Genesis Cloud provider for Pulumi can be used to provision any of the cloud resources available in [Genesis Cloud](https://www.genesiscloud.com).
The Genesis Cloud provider must be configured with credentials to deploy and update resources.

## Example

{{< chooser language "typescript,python,go,dotnet" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import { Instance } from "@genesiscloud/pulumi-genesiscloud";

const example = new Instance("example", {
  image: "my-image-id",
  region: "ARC-IS-HAF-1",
  sshKeyIds: ["my-ssh-key-id"],
  type: "vcpu-2_memory-4g_disk-80g",
});
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_genesiscloud as genesiscloud

example = genesiscloud.Instance("example",
  image="my-image-id",
  region="ARC-IS-HAF-1",
  ssh_key_ids=["my-ssh-key-id"],
  type="vcpu-2_memory-4g_disk-80g")
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"github.com/genesiscloud/pulumi-genesiscloud/sdk/go/genesiscloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    _, err := genesiscloud.NewInstance(ctx, "example", &genesiscloud.InstanceArgs{
      Image:  pulumi.String("my-image-id"),
      Region: pulumi.String("ARC-IS-HAF-1"),
      SshKeyIds: pulumi.StringArray{
        pulumi.String("my-ssh-key-id"),
      },
      Type: pulumi.String("vcpu-2_memory-4g_disk-80g"),
    })
    if err != nil {
      return err
    }
    return nil
  })
}
```

{{% /choosable %}}
{{% choosable language dotnet %}}

```dotnet
TODO: add this
```

{{% /choosable %}}

{{< /chooser >}}
