# InkspanDocument SDK feature factory

from inkspandocument_sdk.feature.base_feature import InkspanDocumentBaseFeature
from inkspandocument_sdk.feature.ratelimit_feature import InkspanDocumentRatelimitFeature
from inkspandocument_sdk.feature.retry_feature import InkspanDocumentRetryFeature
from inkspandocument_sdk.feature.test_feature import InkspanDocumentTestFeature
from inkspandocument_sdk.feature.timeout_feature import InkspanDocumentTimeoutFeature


_FEATURES = {
    "base": lambda: InkspanDocumentBaseFeature(),
    "ratelimit": lambda: InkspanDocumentRatelimitFeature(),
    "retry": lambda: InkspanDocumentRetryFeature(),
    "test": lambda: InkspanDocumentTestFeature(),
    "timeout": lambda: InkspanDocumentTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
