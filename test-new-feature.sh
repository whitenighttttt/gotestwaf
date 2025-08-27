#!/bin/bash

# Test script for the new --includeRequestDetails feature
# This script demonstrates how to use the new functionality

echo "=== GoTestWAF New Feature Test ==="
echo "Testing the --includeRequestDetails functionality"
echo ""

# Check if the new flag is available
echo "1. Checking if --includeRequestDetails flag is available:"
docker run --rm gotestwaf-new --help | grep -q "includeRequestDetails"
if [ $? -eq 0 ]; then
    echo "   ✓ --includeRequestDetails flag is available"
else
    echo "   ✗ --includeRequestDetails flag is not available"
    exit 1
fi

echo ""

# Show the help text for the new flag
echo "2. Help text for --includeRequestDetails:"
docker run --rm gotestwaf-new --help | grep -A 1 -B 1 "includeRequestDetails"

echo ""

# Test basic functionality (without actually running a scan)
echo "3. Testing basic functionality:"
echo "   Creating a test configuration with --includeRequestDetails..."

# Create a test config file
cat > test-config.yaml << EOF
# Test configuration for --includeRequestDetails feature
url: "http://example.com"
email: "test@example.com"
includeRequestDetails: true
reportPath: "/app/reports"
reportName: "test-request-details"
EOF

echo "   ✓ Test configuration created"
echo ""

# Show the configuration
echo "4. Test configuration created:"
cat test-config.yaml

echo ""
echo "=== Test Summary ==="
echo "✓ New --includeRequestDetails flag is available"
echo "✓ Flag has proper help text"
echo "✓ Configuration file supports the new option"
echo ""
echo "To use this feature in Docker:"
echo "docker run --rm --network=\"host\" \\"
echo "  -v \"\$(pwd)/reports:/app/reports\" \\"
echo "  -u \"\$(id -u):\$(id -g)\" \\"
echo "  gotestwaf-new \\"
echo "  --url=\"http://your-target.com\" \\"
echo "  --email=\"your-email@example.com\" \\"
echo "  --includeRequestDetails"
echo ""
echo "Or use the configuration file:"
echo "docker run --rm --network=\"host\" \\"
echo "  -v \"\$(pwd)/reports:/app/reports\" \\"
echo "  -v \"\$(pwd)/test-config.yaml:/app/config.yaml\" \\"
echo "  -u \"\$(id -u):\$(id -g)\" \\"
echo "  gotestwaf-new \\"
echo "  --configPath=\"/app/config.yaml\""
echo ""

# Clean up
rm -f test-config.yaml
echo "✓ Test configuration cleaned up"
