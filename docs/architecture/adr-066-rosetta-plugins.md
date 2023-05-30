# ADR-066: Rosetta Plugins (multi zone=

## Changelog

* May 30, 2023: Initial Draft: {changelog}

## Status

{DRAFT | PROPOSED} Not Implemented

## Abstract

Enabling Rosetta in multiple zones requires adding interfaces and codec types.
Forking the project and implementing it would require a lot of effort. The proposed approach
aims to load such interfaces and types in a simple and easy way instead of implementing
Roseta as a whole

## Context

The Cosmos ecosystem contains multiple zones with different message types and interfaces.
Rosetta is a tool that allows chains to be listed by Coinbase while also providing a standard
access chain information.

In order for Rosetta to work in different zones, it requires it to register the types and
interfaces specific to the zones.

## Alternatives

There are a few alternatives:
- **Implementation**: It requires the zone developers to come up with their own implementation of
Rosetta, from top to bottom
- **Fork**: It takes the Cosmos SDK implementation but modifies it to register the different interfaces.
and keeping it updated by the zone developers.
- **Reflection**: Cosmos-sdk v0.47.x provides a mechanism to reflect over interfaces and types.
through RPc endpoints (see the HUDL tool). It requires zone owners to implement a service on
their nodes in order to enable this feature to be used.
- **Plugin**: Each zone just implements the configurations required (like prefixes) and loads the
required interfaces.

## Decision
In order for Rosetta to access different zones and decode the chain data, it must have
registered chain-specific interfaces; this involves importing such interfaces and types.
In most cases, these types depend on the same libraries as Rosetta does, but with
different version.

Onboarding a new chain involves spending a lot of time tuning up the conflicting packages.
Versioning (also known as dependency hell), in some cases, also means hitting a wall.
making it impossible to onboard a new chain or, in the long term, adding more complexity as the
amount of supported chains grows

The only thing missing on the Cosmos-SDK Rosetta implementation is the ability to work with other chains.
is just missing interfaces. Decoding the interface load process from Rosetta allows
developers to use whatever version of the packages they want without interference from Rosetta.
main dependencies.

Using reflection has been the main idea, like it is being used on the hubl tool in Cosmos-sdk v0.47.x.
There is a service that, if implemented in a node, allows users to reflect on the node.
interfaces through the node's RPC endpoint. Since most of the zones at this time run on v0.45.x,
and this service is not mandatory. Using reflection would still depend on chain developers to
upgrade and implement such a node.

On the one hand, using the go plugins feature allows you to solve the problem of interface decoupling.
load process while getting rid of the dependency issue, since it allows the main function
to execute pre-compiled Go code and call its functions without the need to compile it. On the
On the other side, it only requires chain dev teams to implement two functions in a main.go file:
and place it in a folder under the Rosetta plugins folder; developers should implement `InitZone`.
to configure the SDK configurations, i.e: setting prefixes. And implementing `RegisterInterfaces`.
which receives an interface registry and, through the `RegisterInterface` function, populates
with the chain-specific interfaces.

Once someone aims to start a Rosetta server and specifies the **--blockchain** flag, Rosetta
will look inside its own plug-in folder for such a blockchain name, and it will proceed to
execute such functions, populating the registry, and starting the Rosetta server.

Example of the **InitZone** function

``` go
func InitZone() {
    config := sdk.GetConfig()
    prefix := "osmo"
    config.SetBech32PrefixForAccount(prefix, prefix+"pub")
    config.SetBech32PrefixForValidator(prefix+"valoper", prefix+"valoperpub")
    config.SetBech32PrefixForConsensusNode(prefix+"valcons", prefix+"valconspub")
}
```

Example of the **RegisterInterfaces** function

``` go
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
    ibcclienttypes.RegisterInterfaces(registry)
    ibcLightClient.RegisterInterfaces(registry)
    ibcChannelClient.RegisterInterfaces(registry)
    cosmosWasmClient.RegisterInterfaces(registry)
    sdk.RegisterInterfaces(registry)
    txtypes.RegisterInterfaces(registry)
    cryptocodec.RegisterInterfaces(registry)
}
```

### Positive

- Simple implementation
- Easy to maintain
- Allows to switch to different zones

### Negative

- The Zones team would be responsible for maintaining their code and keeping it up-to-date
- The Cosmos team would be responsible for cleaning up unmaintained zones.

## Test Cases

The implementation has been tested to onboard:
- cosmos-hub
- osmosis

## References

- [Discussion on migrating Rosetta into a standalone repo](https://github.com/cosmos/cosmos-sdk/issues/16276)
- [Zondax Implementation Test](https://github.com/Zondax/cosmos-sdk/pull/653)
