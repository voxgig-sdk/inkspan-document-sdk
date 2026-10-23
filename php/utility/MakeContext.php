<?php
declare(strict_types=1);

// InkspanDocument SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class InkspanDocumentMakeContext
{
    public static function call(array $ctxmap, ?InkspanDocumentContext $basectx): InkspanDocumentContext
    {
        return new InkspanDocumentContext($ctxmap, $basectx);
    }
}
