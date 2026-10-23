# InkspanDocument SDK utility: make_context

from projectname_sdk.core.context import InkspanDocumentContext


def make_context_util(ctxmap, basectx):
    return InkspanDocumentContext(ctxmap, basectx)
