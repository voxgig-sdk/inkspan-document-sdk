# InkspanDocument SDK exists test

import pytest
from inkspandocument_sdk import InkspanDocumentSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = InkspanDocumentSDK.test(None, None)
        assert testsdk is not None
