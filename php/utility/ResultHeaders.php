<?php
declare(strict_types=1);

// InkspanDocument SDK utility: result_headers

class InkspanDocumentResultHeaders
{
    public static function call(InkspanDocumentContext $ctx): ?InkspanDocumentResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
