# Vultr Resource Provider

![Vultr](img/logo.png)

The Vultr Resource Provider lets you manage [Vultr](https://www.vultr.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)


To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @ediri/vultr
```

or `yarn`:

```bash
yarn add @ediri/vultr
```

### Python

To use from Python, install using `pip`:
``
```bash
pip install ediri-vultr
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/dirien/pulumi-vultr/sdk/v2
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package ediri.Vultr
```

## Configuration

The following configuration points are available for the `vultr` provider:

- `vultr:apiKey` (environment: `VULTR_API_KEY`) -  This is the Vultr API key. It can be found in the [Vultr API section](https://my.vultr.com/settings/#settingsapi).
- `vultr:rateLimit` - Vultr limits API calls to 30 calls per second. This field lets you configure how the rate limit using milliseconds. The default value if this field is omitted is 500 milliseconds per call.
- `vultr:retryLimit` - This field lets you configure how many retries should be attempted on a failed call. The default value if this field is omitted is 3 retries.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/vultr/api-docs/).
