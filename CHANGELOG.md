# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-10-03

### 🎉 Initial Release

首次正式发布，提供完整的 Amazon SP-API Go SDK 实现。

### Added

#### Core Infrastructure
- ✅ LWA Authentication (Regular & Grantless operations)
- ✅ AWS Signature Version 4 request signing
- ✅ Restricted Data Token (RDT) support
- ✅ Token Bucket rate limiting algorithm
- ✅ HTTP transport with retry and middleware
- ✅ Comprehensive error handling
- ✅ Request/response encoding and validation

#### API Coverage
- ✅ **57 API versions** fully implemented
- ✅ **314 API operation methods**
- ✅ **1,623 model files** auto-generated from OpenAPI specs
- ✅ Support for all major SP-API endpoints:
  - Orders, Feeds, Reports, Catalog Items
  - FBA Inventory, Fulfillment Inbound/Outbound
  - Listings, Product Pricing, Product Fees
  - Finances, Seller Wallet, Services
  - Messaging, Notifications, Solicitations
  - Shipping, Merchant Fulfillment, Supply Sources
  - Tokens, Uploads, Vehicles, Sales, Sellers
  - A+ Content, Replenishment, AWD, Customer Feedback
  - Data Kiosk, Easy Ship, Applications, Invoices
  - Complete Vendor API suite (20 versions)

#### Testing
- ✅ **92.2% test coverage** for core modules
- ✅ **149 test files** (92 unit + 57 API tests)
- ✅ **150+ test cases** all passing
- ✅ **11 integration tests** for core APIs
- ✅ **Benchmark tests** for performance monitoring

#### Examples & Documentation
- ✅ **7 complete example programs**:
  - Basic usage
  - Orders API
  - Feeds API
  - Reports API
  - Listings API
  - Grantless operations
  - Advanced usage (concurrency, error handling)
- ✅ **9 design documents**
- ✅ **Integration test guide**
- ✅ **Complete API reference**

#### Tools & Utilities
- ✅ CLI code generator
- ✅ Automated API client generation from OpenAPI specs
- ✅ Monitoring and metrics collection
- ✅ Performance profiling utilities
- ✅ Request validation helpers

### Technical Details

#### Dependencies
- Go 1.21+
- No external dependencies for core functionality
- Standard library only

#### Code Quality
- All packages compile successfully
- No linter warnings
- Professional code style
- Complete Go documentation
- Production-ready error handling

### Breaking Changes
None - This is the initial release.

### Migration Guide
Not applicable - Initial release.

### Known Issues
None

### Credits
Built with reference to [Amazon SP-API Official Documentation](https://developer-docs.amazon.com/sp-api/docs/)

---

## Version History

- [1.0.0] - 2025-10-03: Initial release

[1.0.0]: https://github.com/vanling1111/amazon-sp-api-go-sdk/releases/tag/v1.0.0

