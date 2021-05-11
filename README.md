<p align="center">
  <a href="https://www.bitmovin.com">
    <img alt="Bitmovin Go API SDK Header" src="https://cdn.bitmovin.com/frontend/encoding/openapi-clients/readme-headers/ReadmeHeader_Go.png" >
  </a>

  <h4 align="center">
    Go API SDK which enables you to seamlessly integrate the Bitmovin API into your projects.
  </h4>

  <p align="center">
    <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></img></a>
  </p>
</p>

***Important! This is an alpha version, there may be breaking changes.***

Using this API client requires an active account.

> Don't have an account yet? [Sign up for a free Bitmovin trial plan](https://dashboard.bitmovin.com/signup)!

---

## Documentation & Release Notes
+ Full Bitmovin API reference documentation: https://bitmovin.com/docs/encoding/api-reference
+ Release Notes: https://bitmovin.com/docs/encoding/changelogs/rest

## Support
If you have any questions regarding the SDK, provided examples or our services, please log in to your Bitmovin Dashboard at https://bitmovin.com/dashboard and [create a support ticket](https://bitmovin.com/dashboard/support/cases/create?tab=encoding). Our team will get back to you as soon as possible :+1:

---

## Installation:

```bash
go get github.com/bitmovin/bitmovin-api-sdk-go
```

## Initializiation

```go
package main

import (
    "github.com/bitmovin/bitmovin-api-sdk-go"
    "github.com/bitmovin/bitmovin-api-sdk-go/apiclient"
)

func main() {
    bitmovinAPI, err := bitmovin.NewBitmovinAPI(apiclient.WithAPIKey("<YOUR_API_KEY>"))
    if err != nil {
        panic(err)
    }

   // ...
}
```
