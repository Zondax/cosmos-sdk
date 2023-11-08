# Cosmos SDK - Cryptography Module v2

## Analysis

* Currently, there is no ADR providing a comprehensive description of the cryptographic module in the Cosmos SDK.
* There have been multiple requests for a more flexible and extensible approach to cryptography, address management, and more.
* Several open issues require significant changes for resolution.
* Similar efforts have been undertaken in the past concerning runtime modules.
* Existing signing types outside of the crypto module may pose challenges to backward compatibility while striving for a clean interface.
* Security implications must be considered during the module's redesign.

## Objectives

Modular Design Philosophy

* Establish a flexible and extensible foundation using interfaces to enable the seamless integration of various cryptographic methods and tools.

* Restructure, Refactor, and Decouple: Update the cryptography module to ensure modularity and future adaptability.

Documentation & Community Engagement

* Cryptography v2 ADR: Draft a new Architecture Decision Record to guide and document the evolution of the module.

* Enhance documentation to ensure clarity and promote community engagement, providing a platform for feedback and collaborative growth.

Backward Compatibility & Migration

* Prioritize compatibility with previous module versions to avoid disruptions for existing users.

* Design and propose a suitable migration path, ensuring transitions are as seamless as possible.

* Evaluate and decide on the relevance of existing systems and tools, incorporating or deprecating them based on their alignment with the module's new vision.

Developer-Centric Approach

* Prioritize clear, intuitive interfaces and best-practice design principles.
* Improve Developer Experience: Provide tools, samples, and best practices to foster an efficient and user-friendly development environment.

Leverage Extensibility

* Utilize the module's modular design to support a wide range of cryptographic tools, key types, and methods, ensuring adaptability for future technological advancements.
* Integrate support for advanced cryptographic features, ensuring the module's position at the forefront of cryptographic technologies.

Quality Assurance

* Enhanced Test Coverage: Improve testing methodologies to ensure the robustness and reliability of the module.
* Conduct an Audit: After implementation, perform a comprehensive audit to identify potential vulnerabilities and ensure the module's security and stability.

## Technical Goals

Hardware Device & Cloud-based HSM Interface Design:

* Design a foundational interface for various hardware devices (Ledger, YubiKey, Thales, etc.) and cloud-based HSMs (Amazon, Azure) to cater to both current and future implementations.

TPM 2.0 Interface Consideration:

* Integrate design considerations for Trusted Platform Module (TPM) 2.0 support to anticipate future enhancements.

PKCS#11 Interface Blueprint:

* Incorporate the Cryptographic Token Interface Standard (PKCS#11) into the design, ensuring seamless future interactions with cryptographic tokens.

Plugin Architecture and Dependency Injection:

* Establish the architectural foundation for an extensible plugin system and integrate a dependency injection framework, ensuring modularity, testability, and third-party integrations.

Plugin Sandbox Environment Blueprint:

* Design an environment for plugin testing, ensuring developers can validate integrations without compromising system integrity.

Extensibility for Cryptographic Techniques:

* Design the system with extensibility in mind to accommodate a broad spectrum of cryptographic techniques such as:
* Various signature types
* Different key types (elliptic curve, RSA, etc.)
* Post-Quantum Cryptography (PQC) methods
* Threshold signatures and encryption

Community Engagement Infrastructure:

* Structure the design with tools and documentation interfaces in mind, enabling a seamless future rollout of resources for developer engagement.

## Work Plan
This work plan describes the expected timelines for the architectural project milestones. We estimate it to be worked on in an estimated period of between 7 to 8 months. The plan offers clear milestones, deadlines, dependencies and interactions, ensuring coordinated efforts and systematic architectural changes by moving forward safely and reviewing each version after each milestone is completed. The goal is to achieve efficient implementation while meeting project objectives. It should be noted that each estimated milestone length is not serial, this means that some milestones are executed overlapped with their next or previous milestone at the beginning or end of the milestone.

### Phase 1: Foundation & Compatibility

#### Milestone 1: Design & Documentation

Finalize the Cryptography v2 ADR.
Publish an initial version of the updated module documentation.

Estimated time: 1 and 1/2 work month.

#### Milestone 2: Modularization & Refactoring

Complete the modular design based on the identified interface standards.
Successfully refactor and decouple the existing cryptography module.

Estimated time: 2 work months.

#### Milestone 3: Compatibility & Migration

Validate full backward compatibility with previous module versions.
Develop and conduct initial testing of the migration tools to ensure they are robust and user-friendly.

Estimated time: 2 and 1/2 work months.

#### Milestone 4: Developer Interfaces & Legacy Systems

Release new developer interfaces with sample applications/tools.
Decide on the incorporation, refactoring, or deprecation of all identified legacy elements.

Estimated time: 1 work month.

### Phase 2: Extensibility & Enhancement

#### Milestone 5: Extensibility Implementation

Successfully integrate the initial set of plugins and cryptographic tools.
Release a beta version of the sandbox environment for extensibility testing.

Estimated time: 1 work month.

#### Milestone 6: Advanced Cryptography Integration

Fully integrate and validate post-quantum cryptographic methods.
Report pilot testing results for the newly integrated cryptographic features.

Estimated time: 2 work months.

#### Milestone 7: Quality Assurance & Audit

Achieve targeted test coverage metrics for the new implementations.
Complete the external security and functionality audit with any necessary revisions made.

Estimated time: 1 work month.

#### Milestone 8: Full-Scale Deployment & Community Engagement

Officially release the Cryptography Module v2.
Organize community events, such as webinars or workshops, and gather feedback for potential future enhancements.

Estimated time: 1 work month.

## References

* Existing ADRs https://docs.cosmos.network/main/architecture

* https://docs.cosmos.network/main/architecture/adr-006-secret-store-replacement

* Runtime standalone modules https://github.com/cosmos/cosmos-sdk/issues/11899

* Relevant Issues ​


    * https://github.com/cosmos/cosmos-sdk/issues?q=is%3Aopen+is%3Aissue+label%3AC%3ACrypto

    * https://github.com/cosmos/cosmos-sdk/labels/C%3ALedger

    * https://github.com/cosmos/cosmos-sdk/labels/C%3AKeys

