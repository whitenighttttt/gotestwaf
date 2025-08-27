# GoTestWAF Docker Running Examples

This document provides examples of how to run GoTestWAF using Docker with various configuration options.

## Basic Usage

```bash
sudo docker run --rm --network="host" \
  -v "$(pwd)/reports:/app/reports" \
  -u "$(id -u):$(id -g)" \
  gotestwaf \
  --url="http://example.com" \
  --email="user@example.com"
```

## With HTTP Request Details Enabled

To include HTTP request details (method, URL, headers, body) in reports and CSV exports:

```bash
sudo docker run --rm --network="host" \
  -v "$(pwd)/reports:/app/reports" \
  -u "$(id -u):$(id -g)" \
  gotestwaf \
  --url="http://aaa.wangyaoyang.top" \
  --email="yaoyangwang@tencent.com" \
  --blockStatusCodes=566,400 \
  --passStatusCodes=200,301,302,303,304,307,404,405,410 \
  --reportPath="/app/reports" \
  --reportName="eo-waf-2052" \
  --includeRequestDetails
```

## Advanced Configuration

```bash
sudo docker run --rm --network="host" \
  -v "$(pwd)/reports:/app/reports" \
  -v "$(pwd)/testcases:/app/testcases" \
  -v "$(pwd)/config.yaml:/app/config.yaml" \
  -u "$(id -u):$(id -g)" \
  gotestwaf \
  --url="https://target-website.com" \
  --email="admin@company.com" \
  --wafName="Cloudflare WAF" \
  --blockStatusCodes=403,429,444 \
  --passStatusCodes=200,301,302,404 \
  --reportPath="/app/reports" \
  --reportName="waf-test-$(date +%Y%m%d-%H%M%S)" \
  --reportFormat=html,pdf \
  --includePayloads \
  --includeRequestDetails \
  --workers=10 \
  --sendDelay=200
```

## With OpenAPI Specification

```bash
sudo docker run --rm --network="host" \
  -v "$(pwd)/reports:/app/reports" \
  -v "$(pwd)/api-spec.yaml:/app/api-spec.yaml" \
  -u "$(id -u):$(id -g)" \
  gotestwaf \
  --url="https://api.example.com" \
  --openapiFile="/app/api-spec.yaml" \
  --includeRequestDetails \
  --reportPath="/app/reports"
```

## Chrome-based Testing

For testing with Chrome browser (requires additional capabilities):

```bash
sudo docker run --rm --network="host" \
  --cap-add=SYS_ADMIN \
  -v "$(pwd)/reports:/app/reports" \
  -u "$(id -u):$(id -g)" \
  gotestwaf \
  --url="https://example.com" \
  --httpClient=chrome \
  --includeRequestDetails \
  --reportPath="/app/reports"
```

## Configuration File Example

You can also use a configuration file to set options:

```yaml
# config.yaml
url: "https://example.com"
email: "user@example.com"
wafName: "Generic WAF"
blockStatusCodes: [403, 429, 444]
passStatusCodes: [200, 301, 302, 404]
reportPath: "/app/reports"
reportName: "waf-test"
includePayloads: true
includeRequestDetails: true
workers: 10
sendDelay: 200
```

Then run with:

```bash
sudo docker run --rm --network="host" \
  -v "$(pwd)/reports:/app/reports" \
  -v "$(pwd)/config.yaml:/app/config.yaml" \
  -u "$(id -u):$(id -g)" \
  gotestwaf \
  --configPath="/app/config.yaml"
```

## Notes

- The `--includeRequestDetails` flag enables recording of HTTP request method, URL, headers, and body in reports and CSV exports
- This feature is useful for detailed analysis and debugging of WAF behavior
- When enabled, the CSV export will include additional columns for HTTP request details
- The feature works with both GoHTTP and Chrome HTTP clients
- Request body content-type detection is automatically performed for better accuracy
